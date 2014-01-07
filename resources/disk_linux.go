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
	"os/exec"
	"strings"
)

type DiskMetrics struct {
	MajorNumber      int64  `json:",omitempty"`
	MinorNumber      int64  `json:",omitempty"`
	DeviceName       string `json:",omitempty"`
	ReadsComplete    int64  `json:",omitempty"`
	ReadsMerged      int64  `json:",omitempty"`
	SectorsRead      int64  `json:",omitempty"`
	TimeSpentReading int64  `json:",omitempty"`
	WritesComplete   int64  `json:",omitempty"`
	WritesMerged     int64  `json:",omitempty"`
	SectorsWritten   int64  `json:",omitempty"`
	TimeSpentWriting int64  `json:",omitempty"`
	CurrentIO        int64  `json:",omitempty"`
	IOTime           int64  `json:",omitempty"`
	WeightedIOTime   int64  `json:",omitempty"`
	FilesystemType   string `json:",omitempty"`
	Blocks           int64  `json:",omitempty"`
	BlocksUsed       int64  `json:",omitempty"`
	BlocksFree       int64  `json:",omitempty"`
	MountPoint       string `json:",omitempty"`
}

func (self *Collection) CollectDiskMetrics() *Collection {
	self.updateTimestamp()
	diskStats, err := os.Open("/proc/diskstats")
	verifyErrorResponse(err, "reading /proc/diskstats")
	defer diskStats.Close()

	diskStatsScanner := bufio.NewScanner(diskStats)
	self.Disk = make([]DiskMetrics, 0, 10)

	df, err := exec.Command("df", "-T").Output()
	verifyErrorResponse(err, "executing df command")
	dfOutputLines := strings.Split(string(df), "\n")

	for diskStatsScanner.Scan() {
		scannedDiskMetrics := DiskMetrics{}
		scannedLine := strings.Fields(diskStatsScanner.Text())
		for _, dfOutputLine := range dfOutputLines {
			if strings.Contains(dfOutputLine, scannedLine[2]) == true {
				diskUsage := strings.Fields(dfOutputLine)
				scannedDiskMetrics.FilesystemType = diskUsage[1]
				scannedDiskMetrics.Blocks = atoi64(diskUsage[2])
				scannedDiskMetrics.BlocksUsed = atoi64(diskUsage[3])
				scannedDiskMetrics.BlocksFree = atoi64(diskUsage[4])
				scannedDiskMetrics.MountPoint = diskUsage[6]
			}
		}
		scannedDiskMetrics.MajorNumber = atoi64(scannedLine[0])
		scannedDiskMetrics.MinorNumber = atoi64(scannedLine[1])
		scannedDiskMetrics.DeviceName = scannedLine[2]
		scannedDiskMetrics.ReadsComplete = atoi64(scannedLine[3])
		scannedDiskMetrics.ReadsMerged = atoi64(scannedLine[4])
		scannedDiskMetrics.SectorsRead = atoi64(scannedLine[5])
		scannedDiskMetrics.TimeSpentReading = atoi64(scannedLine[6])
		scannedDiskMetrics.WritesComplete = atoi64(scannedLine[7])
		scannedDiskMetrics.WritesMerged = atoi64(scannedLine[8])
		scannedDiskMetrics.SectorsWritten = atoi64(scannedLine[9])
		scannedDiskMetrics.TimeSpentWriting = atoi64(scannedLine[10])
		scannedDiskMetrics.CurrentIO = atoi64(scannedLine[11])
		scannedDiskMetrics.IOTime = atoi64(scannedLine[12])
		scannedDiskMetrics.WeightedIOTime = atoi64(scannedLine[13])
		self.Disk = append(self.Disk, scannedDiskMetrics)
	}
	return self
}
