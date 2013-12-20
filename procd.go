// Copyright 2013, Belly, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bellycard/procd/resources"
	"github.com/bellycard/toml"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type procdConfig struct {
	TickerInterval int               `toml:"ticker_interval"`
	Outputs        map[string]output `toml:"output"`
}

type output struct {
	BindAddress string `toml:"bind_address"`
	Server      string
	Encoder     string
	Sender      string
	Payload     bool
	Hostname    string
}

func waitForSignals(exitStatus <-chan bool) {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, os.Kill)
	for {
		s := <-signalChan
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, os.Kill:
			fmt.Println("Procd: Shutdown initiated.")
			os.Exit(0)
		}
	}
}

func main() {
	flagVersion := flag.Bool("version", false, "Display Procd version.")
	flagConfigFile := flag.String("config", "procd.toml", "Procd configuration file.")
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	// Handle no command-line paramters
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Display Procd version
	if *flagVersion {
		fmt.Println("Procd v", Version)
		os.Exit(0)
	}

	// Decode configuration [https://github.com/mojombo/toml]
	var config procdConfig
	_, err := toml.DecodeFile(*flagConfigFile, &config)
	verifyErrorResponse(err, "could not decode TOML config")

	// Setup output channels
	for outputName, output := range config.Outputs {
		switch outputName {
		case "stdout":
			fmt.Printf("\nProcd: Output[stdout]\n")
		case "http":
			go HTTPServerStart(output.BindAddress)
			fmt.Printf("Procd: Output[http]: %s\n", output.BindAddress)
		case "heka":
			fmt.Printf("Procd: Output[heka]: %s [TCP / Protobufs]\n", output.Server)
		}
	}

	// Timing channels for polling interval
	collectionInterval := time.NewTicker(time.Second * time.Duration(config.TickerInterval))

	collection := resources.NewCollection()
	// Update statistics on interval until sigterm
	go func() {
		for _ = range collectionInterval.C {
			collection.CollectAllMetrics()
			for outputName, output := range config.Outputs {
				switch outputName {
				case "stdout":
					marshalledResources, _ := json.MarshalIndent(collection, "", " ")
					fmt.Println(string(marshalledResources))
				case "heka":
					hc, err := NewHekaClient(output.Server, output.Hostname)
					if err == nil {
						err := hc.SendCPUResources(collection, output.Payload)
						verifyErrorResponse(err, "generate CPU resources Heka message")
						err = hc.SendMemoryResources(collection, output.Payload)
						verifyErrorResponse(err, "generate memory resources Heka message")
						err = hc.SendDiskResources(collection, output.Payload)
						verifyErrorResponse(err, "generate disk resources Heka message")
						hc.sender.Close()
					}
				}
			}
		}
	}()

	fmt.Printf("Procd: Collection interval: %v seconds\n", config.TickerInterval)
	exitStatus := make(chan bool, 1)
	fmt.Println("Procd: Started.")
	waitForSignals(exitStatus)
}
