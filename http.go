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
	"fmt"
	"github.com/bellycard/procd/resources"
	"net/http"
)

const (
	PingResponse = "pong"
)

func rootFileHandler(p string, f string) {
	http.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, f)
	})
}

func HTTPServerStart(port string) {
	http.Handle("/", http.HandlerFunc(rootHandler))
	http.Handle("/v1/resources.json", http.HandlerFunc(resourcesHandler))
	http.Handle("/v1/resources/cpu.json", http.HandlerFunc(resourcesCpuHandler))
	http.Handle("/v1/resources/memory.json", http.HandlerFunc(resourcesMemoryHandler))
	http.Handle("/v1/resources/disk.json", http.HandlerFunc(resourcesDiskHandler))
	http.Handle("/ping", http.HandlerFunc(pingHandler))
	rootFileHandler("/favicon.ico", "public/favicon.ico")
	err := http.ListenAndServe(port, nil)
	verifyErrorResponse(err, "output[http] start failure")
}

func resourcesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := resources.NewCollection()
	collection.CollectAllMetrics()
	marshalledResources, _ := json.MarshalIndent(collection, "", "  ")
	fmt.Fprintf(w, string(marshalledResources))
}

func resourcesCpuHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := resources.NewCollection()
	collection.CollectCPUMetrics()
	marshalledResources, _ := json.MarshalIndent(collection, "", "  ")
	fmt.Fprintf(w, string(marshalledResources))
}

func resourcesMemoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := resources.NewCollection()
	collection.CollectMemoryMetrics()
	marshalledResources, _ := json.MarshalIndent(collection, "", "  ")
	fmt.Fprintf(w, string(marshalledResources))
}

func resourcesDiskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := resources.NewCollection()
	collection.CollectDiskMetrics()
	marshalledResources, _ := json.MarshalIndent(collection, "", "  ")
	fmt.Fprintf(w, string(marshalledResources))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, PingResponse)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Belly Procd v%s", Version)
}
