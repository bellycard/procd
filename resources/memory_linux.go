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
//
// Contributor(s):
//   Christian Vozar (christian@bellycard.com)

package resources

import (
	"bufio"
	"os"
	"strings"
)

type MemoryMetrics struct {
	Total     int64 `json:",omitempty"`
	Free      int64 `json:",omitempty"`
	SwapTotal int64 `json:",omitempty"`
	SwapFree  int64 `json:",omitempty"`
	Buffers   int64 `json:",omitempty"`
	Cached    int64 `json:",omitempty"`
}

func (self *Collection) CollectMemoryMetrics() *Collection {
	self.updateTimestamp()
	memInfo, err := os.Open("/proc/meminfo")
	verifyErrorResponse(err, "reading /proc/meminfo")
	defer memInfo.Close()

	memInfoScanner := bufio.NewScanner(memInfo)
	for memInfoScanner.Scan() {
		scannedLine := strings.Fields(memInfoScanner.Text())
		switch scannedLine[0] {
		case "MemTotal:":
			self.Memory.Total = atoi64(scannedLine[1])
		case "MemFree:":
			self.Memory.Free = atoi64(scannedLine[1])
		case "SwapTotal:":
			self.Memory.SwapTotal = atoi64(scannedLine[1])
		case "SwapFree:":
			self.Memory.SwapFree = atoi64(scannedLine[1])
		case "Buffers:":
			self.Memory.Buffers = atoi64(scannedLine[1])
		case "Cached:":
			self.Memory.Cached = atoi64(scannedLine[1])
		}
	}
	return self
}
