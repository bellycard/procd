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
func NewHekaClient(h, e, s, hn string) (self *HekaClient, err error) {
	pid := int32(os.Getpid())
	if hn == "" {
		hostname, _ := os.Hostname()
	} else {
		hostname := hn
	}
	encoder := client.NewProtobufEncoder(nil)
	sender, err := client.NewNetworkSender(s, h)
	if err == nil {
		self = &HekaClient{encoder: encoder, sender: sender, pid: pid, hostname: hostname}
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
	f, _ := message.NewField("cpu.user", pl.CPU.User, "")
	msg.AddField(f)
	f, _ = message.NewField("cpu.nice", pl.CPU.Nice, "")
	msg.AddField(f)
	f, _ = message.NewField("cpu.idle", pl.CPU.Idle, "")
	msg.AddField(f)
	f, _ = message.NewField("cpu.iowait", pl.CPU.IOWait, "")
	msg.AddField(f)
	f, _ = message.NewField("cpu.irq", pl.CPU.IRQ, "")
	msg.AddField(f)
	f, _ = message.NewField("cpu.softirq", pl.CPU.SoftIRQ, "")
	msg.AddField(f)
	f, _ = message.NewField("cpu.steal", pl.CPU.Steal, "")
	msg.AddField(f)
	f, _ = message.NewField("cpu.guest", pl.CPU.Guest, "")
	msg.AddField(f)
	f, _ = message.NewField("cpu.guestnice", pl.CPU.GuestNice, "")
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
	f, _ := message.NewField("memory.total", pl.Memory.Total, "")
	msg.AddField(f)
	f, _ = message.NewField("memory.free", pl.Memory.Free, "")
	msg.AddField(f)
	f, _ = message.NewField("memory.swap.total", pl.Memory.SwapTotal, "")
	msg.AddField(f)
	f, _ = message.NewField("memory.swap.free", pl.Memory.SwapFree, "")
	msg.AddField(f)
	f, _ = message.NewField("memory.buffers", pl.Memory.Buffers, "")
	msg.AddField(f)
	f, _ = message.NewField("memory.cached", pl.Memory.Cached, "")
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

		f, _ := message.NewField("disk.majornumber", v.MajorNumber, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.minornumber", v.MinorNumber, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.devicename", v.DeviceName, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.reads.complete", v.ReadsComplete, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.reads.merged", v.ReadsMerged, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.sectors.read", v.SectorsRead, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.time.spent.reading", v.TimeSpentReading, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.writes.complete", v.WritesComplete, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.writes.merged", v.WritesMerged, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.sectors.written", v.SectorsWritten, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.time.spent.writing", v.TimeSpentWriting, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.io.current", v.CurrentIO, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.io.time", v.IOTime, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.io.time.weighted", v.WeightedIOTime, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.filesystem.type", v.FilesystemType, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.filesystem.blocks", v.Blocks, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.filesystem.blocks.used", v.BlocksUsed, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.filesystem.blocks.free", v.BlocksFree, "")
		msg.AddField(f)
		f, _ = message.NewField("disk.filesystem.mountpoint", v.MountPoint, "")
		msg.AddField(f)

		err = self.encoder.EncodeMessageStream(msg, &stream)
		verifyErrorResponse(err, "output[heka] encode message error")
		err = self.sender.SendMessage(stream)
		verifyErrorResponse(err, "output[heka] send message error")
	}
	return nil
}
