// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

/*
Package storage is a generated protocol buffer package.

It is generated from these files:
	storage.proto

It has these top-level messages:
	UtxoEntry
*/
package storage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UtxoEntry struct {
	Type        uint32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	BlockHeight uint64 `protobuf:"varint,2,opt,name=blockHeight" json:"blockHeight,omitempty"`
	Spent       bool   `protobuf:"varint,3,opt,name=spent" json:"spent,omitempty"`
}

func (m *UtxoEntry) Reset()                    { *m = UtxoEntry{} }
func (m *UtxoEntry) String() string            { return proto.CompactTextString(m) }
func (*UtxoEntry) ProtoMessage()               {}
func (*UtxoEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UtxoEntry) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *UtxoEntry) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *UtxoEntry) GetSpent() bool {
	if m != nil {
		return m.Spent
	}
	return false
}

func init() {
	proto.RegisterType((*UtxoEntry)(nil), "chain.core.txdb.internal.storage.UtxoEntry")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0xce, 0x48, 0xcc, 0xcc,
	0xd3, 0x4b, 0xce, 0x2f, 0x4a, 0xd5, 0x2b, 0xa9, 0x48, 0x49, 0xd2, 0xcb, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0x83, 0xaa, 0x53, 0x0a, 0xe7, 0xe2, 0x0c, 0x2d, 0xa9, 0xc8, 0x77, 0xcd,
	0x2b, 0x29, 0xaa, 0x14, 0x12, 0xe2, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0d, 0x02, 0xb3, 0x85, 0x14, 0xb8, 0xb8, 0x93, 0x72, 0xf2, 0x93, 0xb3, 0x3d, 0x52, 0x33,
	0xd3, 0x33, 0x4a, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x82, 0x90, 0x85, 0x84, 0x44, 0xb8, 0x58,
	0x8b, 0x0b, 0x52, 0xf3, 0x4a, 0x24, 0x98, 0x15, 0x18, 0x35, 0x38, 0x82, 0x20, 0x1c, 0x27, 0xce,
	0x28, 0x76, 0xa8, 0x1d, 0x49, 0x6c, 0x60, 0xc7, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x61,
	0x0a, 0x61, 0x8f, 0x9d, 0x00, 0x00, 0x00,
}