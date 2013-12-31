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
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"github.com/bellycard/procd/resources"
	"github.com/mozilla-services/heka/client"
	"github.com/mozilla-services/heka/message"
	"os"
)

type HekaClient struct {
	client   client.Client
	encoder  client.Encoder // I.e. protobufs
	sender   client.Sender  // I.e. tcp
	pid      int32
	hostname string
}

// NewHekaClient returns a new HekaClient with process ID, hostname, encoder and sender.
func NewHekaClient(h, s, hn string) (self *HekaClient, err error) {
	self = &HekaClient{}
	self.pid = int32(os.Getpid())
	if hn == "" {
		self.hostname, _ = os.Hostname()
	} else {
		self.hostname = hn
	}
	self.encoder = client.NewProtobufEncoder(nil)
	self.sender, err = client.NewNetworkSender(s, h)
	if err == nil {
		return self, nil
	}
	return
}

func (self HekaClient) SendCPUResources(pl *resources.Collection, ipl bool) (err error) {
	var stream []byte

	msg := &message.Message{}
	msg.SetTimestamp(pl.Timestamp.UnixNano())
	msg.SetUuid(uuid.NewRandom())
	msg.SetType("procd.resources.cpu")
	msg.SetLogger("procd")
	msg.SetPid(self.pid)
	msg.SetSeverity(int32(6))
	msg.SetHostname(self.hostname)

	if ipl == true {
		marshalledPayload, _ := json.MarshalIndent(pl, "", " ")
		msg.SetPayload(string(marshalledPayload))
	}

	// Super awful way to do this, should use reflection but need to test quickly
	f, _ := message.NewField("CpuUser", pl.CPU.User, "")
	msg.AddField(f)
	f, _ = message.NewField("CpuNice", pl.CPU.Nice, "")
	msg.AddField(f)
	f, _ = message.NewField("CpuIdle", pl.CPU.Idle, "")
	msg.AddField(f)
	f, _ = message.NewField("CpuIowait", pl.CPU.IOWait, "")
	msg.AddField(f)
	f, _ = message.NewField("CpuIrq", pl.CPU.IRQ, "")
	msg.AddField(f)
	f, _ = message.NewField("CpuSoftirq", pl.CPU.SoftIRQ, "")
	msg.AddField(f)
	f, _ = message.NewField("CpuSteal", pl.CPU.Steal, "")
	msg.AddField(f)
	f, _ = message.NewField("CpuGuest", pl.CPU.Guest, "")
	msg.AddField(f)
	f, _ = message.NewField("CpuGuestnice", pl.CPU.GuestNice, "")
	msg.AddField(f)

	err = self.encoder.EncodeMessageStream(msg, &stream)
	verifyErrorResponse(err, "output[heka] encode message error")
	err = self.sender.SendMessage(stream)
	verifyErrorResponse(err, "output[heka] send message error")
	return nil
}

func (self HekaClient) SendMemoryResources(pl *resources.Collection, ipl bool) (err error) {
	var stream []byte

	msg := &message.Message{}
	msg.SetTimestamp(pl.Timestamp.UnixNano())
	msg.SetUuid(uuid.NewRandom())
	msg.SetType("procd.resources.memory")
	msg.SetLogger("procd")
	msg.SetPid(self.pid)
	msg.SetSeverity(int32(6))
	msg.SetHostname(self.hostname)

	if ipl == true {
		marshalledPayload, _ := json.MarshalIndent(pl, "", " ")
		msg.SetPayload(string(marshalledPayload))
	}

	// Super awful way to do this, should use reflection but need to test quickly
	f, _ := message.NewField("MemoryTotal", pl.Memory.Total, "")
	msg.AddField(f)
	f, _ = message.NewField("MemoryFree", pl.Memory.Free, "")
	msg.AddField(f)
	f, _ = message.NewField("MemorySwapTotal", pl.Memory.SwapTotal, "")
	msg.AddField(f)
	f, _ = message.NewField("MemorySwapFree", pl.Memory.SwapFree, "")
	msg.AddField(f)
	f, _ = message.NewField("MemoryBuffers", pl.Memory.Buffers, "")
	msg.AddField(f)
	f, _ = message.NewField("MemoryCached", pl.Memory.Cached, "")
	msg.AddField(f)

	err = self.encoder.EncodeMessageStream(msg, &stream)
	verifyErrorResponse(err, "output[heka] encode message error")
	err = self.sender.SendMessage(stream)
	verifyErrorResponse(err, "output[heka] send message error")
	return nil
}

func (self HekaClient) SendDiskResources(pl *resources.Collection, ipl bool) (err error) {

	// Super awful way to do this, should use reflection but need to test quickly
	for _, v := range pl.Disk {
		var stream []byte

		msg := &message.Message{}
		msg.SetTimestamp(pl.Timestamp.UnixNano())
		msg.SetUuid(uuid.NewRandom())
		msg.SetType("procd.resources.disk")
		msg.SetLogger("procd")
		msg.SetPid(self.pid)
		msg.SetSeverity(int32(6))
		msg.SetHostname(self.hostname)

		if ipl == true {
			marshalledPayload, _ := json.MarshalIndent(pl, "", " ")
			msg.SetPayload(string(marshalledPayload))
		}

		f, _ := message.NewField("DiskMajorNumber", v.MajorNumber, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskMinorNumber", v.MinorNumber, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskDeviceName", v.DeviceName, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskReadsComplete", v.ReadsComplete, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskReadsMerged", v.ReadsMerged, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskSectorsRead", v.SectorsRead, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskTimeSpentReading", v.TimeSpentReading, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskWritesComplete", v.WritesComplete, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskWritesMerged", v.WritesMerged, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskSectorsWritten", v.SectorsWritten, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskTimeSpentWriting", v.TimeSpentWriting, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskIoCurrent", v.CurrentIO, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskIoTime", v.IOTime, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskIoTimeWeighted", v.WeightedIOTime, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskFilesystemType", v.FilesystemType, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskFilesystemBlocks", v.Blocks, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskFilesystemBlocksUsed", v.BlocksUsed, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskFilesystemBlocksFree", v.BlocksFree, "")
		msg.AddField(f)
		f, _ = message.NewField("DiskFilesystemMountpoint", v.MountPoint, "")
		msg.AddField(f)

		err = self.encoder.EncodeMessageStream(msg, &stream)
		verifyErrorResponse(err, "output[heka] encode message error")
		err = self.sender.SendMessage(stream)
		verifyErrorResponse(err, "output[heka] send message error")
	}
	return nil
}
