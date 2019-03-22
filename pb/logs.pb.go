// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logs.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LogStore struct {
	LogId                string   `protobuf:"bytes,1,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	Timestamp            string   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Log                  string   `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Host                 string   `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogStore) Reset()         { *m = LogStore{} }
func (m *LogStore) String() string { return proto.CompactTextString(m) }
func (*LogStore) ProtoMessage()    {}
func (*LogStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_782e6d65c19305b4, []int{0}
}

func (m *LogStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogStore.Unmarshal(m, b)
}
func (m *LogStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogStore.Marshal(b, m, deterministic)
}
func (m *LogStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogStore.Merge(m, src)
}
func (m *LogStore) XXX_Size() int {
	return xxx_messageInfo_LogStore.Size(m)
}
func (m *LogStore) XXX_DiscardUnknown() {
	xxx_messageInfo_LogStore.DiscardUnknown(m)
}

var xxx_messageInfo_LogStore proto.InternalMessageInfo

func (m *LogStore) GetLogId() string {
	if m != nil {
		return m.LogId
	}
	return ""
}

func (m *LogStore) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *LogStore) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *LogStore) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func init() {
	proto.RegisterType((*LogStore)(nil), "pb.LogStore")
}

func init() { proto.RegisterFile("logs.proto", fileDescriptor_782e6d65c19305b4) }

var fileDescriptor_782e6d65c19305b4 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xc9, 0x4f, 0x2f,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x4a, 0xe5, 0xe2, 0xf0, 0xc9,
	0x4f, 0x0f, 0x2e, 0xc9, 0x2f, 0x4a, 0x15, 0x12, 0xe5, 0x62, 0xcb, 0xc9, 0x4f, 0x8f, 0xcf, 0x4c,
	0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xcd, 0xc9, 0x4f, 0xf7, 0x4c, 0x11, 0x92, 0xe1,
	0xe2, 0x2c, 0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x90, 0x60, 0x02, 0xcb, 0x20, 0x04,
	0x84, 0x04, 0xb8, 0x98, 0x73, 0xf2, 0xd3, 0x25, 0x98, 0xc1, 0xe2, 0x20, 0xa6, 0x90, 0x10, 0x17,
	0x4b, 0x46, 0x7e, 0x71, 0x89, 0x04, 0x0b, 0x58, 0x08, 0xcc, 0x4e, 0x62, 0x03, 0xdb, 0x68, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x9f, 0xdb, 0x78, 0x7f, 0x00, 0x00, 0x00,
}
