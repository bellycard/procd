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

package resources

import (
	"os"
	"time"
)

type Collection struct {
	Timestamp                      time.Time
	Hostname                       string
	CPU                            CPUMetrics     `json:",omitempty"`
	Memory                         MemoryMetrics  `json:",omitempty"`
	Disk                           []DiskMetrics  `json:",omitempty"`
	Entropy                        EntropyMetrics `json:",omitempty"`
	Procs                          int64          `json:",omitempty"`
	ProcsRunning                   int64          `json:",omitempty"`
	ProcsBlocked                   int64          `json:",omitempty"`
	SystemContextSwitchesPerSecond int64          `json:",omitempty"`
}

// NewCollection returns a new Collection of metrics with a timestamp and hostnames.
func NewCollection() *Collection {
	timestamp := time.Now()
	hostname, _ := os.Hostname()
	return &Collection{Timestamp: timestamp, Hostname: hostname}
}

func (self *Collection) updateTimestamp() *Collection {
  self.Timestamp = time.Now()
  return self
}

func (self *Collection) CollectAllMetrics() *Collection {
	self.updateTimestamp()
	self.CollectCPUMetrics()
	self.CollectMemoryMetrics()
	self.CollectDiskMetrics()
	self.CollectEntropy()
	return self
}
