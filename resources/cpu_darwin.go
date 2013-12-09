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
	return self
}
