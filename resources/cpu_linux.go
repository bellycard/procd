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
	"bufio"
	"os"
	"strings"
)

type CPUMetrics struct {
	User      int64 `json:",omitempty"`
	Nice      int64 `json:",omitempty"`
	System    int64 `json:",omitempty"`
	Idle      int64 `json:",omitempty"`
	IOWait    int64 `json:",omitempty"`
	IRQ       int64 `json:",omitempty"`
	SoftIRQ   int64 `json:",omitempty"`
	Steal     int64 `json:",omitempty"`
	Guest     int64 `json:",omitempty"`
	GuestNice int64 `json:",omitempty"`
}

func (self *Collection) CollectCPUMetrics() *Collection {
	self.updateTimestamp()
	stat, err := os.Open("/proc/stat")
	verifyErrorResponse(err, "reading /proc/stat")
	defer stat.Close()

	statScanner := bufio.NewScanner(stat)
	for statScanner.Scan() {
		scannedLine := strings.Fields(statScanner.Text())
		switch scannedLine[0] {
		case "cpu":
			self.CPU.User = atoi64(scannedLine[1])
			self.CPU.Nice = atoi64(scannedLine[2])
			self.CPU.System = atoi64(scannedLine[3])
			self.CPU.Idle = atoi64(scannedLine[4])
			self.CPU.IOWait = atoi64(scannedLine[5])
			self.CPU.IRQ = atoi64(scannedLine[6])
			self.CPU.SoftIRQ = atoi64(scannedLine[7])
			self.CPU.Steal = atoi64(scannedLine[8])
			self.CPU.Guest = atoi64(scannedLine[9])
			self.CPU.GuestNice = atoi64(scannedLine[10])
		case "ctxt":
			self.SystemContextSwitchesPerSecond = atoi64(scannedLine[1])
		case "processes":
			self.Procs = atoi64(scannedLine[1])
		case "procs_running":
			self.ProcsRunning = atoi64(scannedLine[1])
		case "procs_blocked":
			self.ProcsBlocked = atoi64(scannedLine[1])
		}
	}
	return self
}
