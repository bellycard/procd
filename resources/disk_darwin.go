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
	return self
}
