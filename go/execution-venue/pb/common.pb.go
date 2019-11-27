// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

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

type Decimal64 struct {
	Mantissa             int64    `protobuf:"fixed64,1,opt,name=mantissa,proto3" json:"mantissa,omitempty"`
	Exponent             int32    `protobuf:"fixed32,2,opt,name=exponent,proto3" json:"exponent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Decimal64) Reset()         { *m = Decimal64{} }
func (m *Decimal64) String() string { return proto.CompactTextString(m) }
func (*Decimal64) ProtoMessage()    {}
func (*Decimal64) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Decimal64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Decimal64.Unmarshal(m, b)
}
func (m *Decimal64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Decimal64.Marshal(b, m, deterministic)
}
func (m *Decimal64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Decimal64.Merge(m, src)
}
func (m *Decimal64) XXX_Size() int {
	return xxx_messageInfo_Decimal64.Size(m)
}
func (m *Decimal64) XXX_DiscardUnknown() {
	xxx_messageInfo_Decimal64.DiscardUnknown(m)
}

var xxx_messageInfo_Decimal64 proto.InternalMessageInfo

func (m *Decimal64) GetMantissa() int64 {
	if m != nil {
		return m.Mantissa
	}
	return 0
}

func (m *Decimal64) GetExponent() int32 {
	if m != nil {
		return m.Exponent
	}
	return 0
}

func init() {
	proto.RegisterType((*Decimal64)(nil), "common.Decimal64")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 104 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x9c, 0xb9, 0x38,
	0x5d, 0x52, 0x93, 0x33, 0x73, 0x13, 0x73, 0xcc, 0x4c, 0x84, 0xa4, 0xb8, 0x38, 0x72, 0x13, 0xf3,
	0x4a, 0x32, 0x8b, 0x8b, 0x13, 0x25, 0x18, 0x15, 0x18, 0x35, 0x04, 0x82, 0xe0, 0x7c, 0x90, 0x5c,
	0x6a, 0x45, 0x41, 0x7e, 0x5e, 0x6a, 0x5e, 0x89, 0x04, 0x13, 0x50, 0x8e, 0x3f, 0x08, 0xce, 0x77,
	0x62, 0x89, 0x62, 0x2a, 0x48, 0x4a, 0x62, 0x03, 0x9b, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x83, 0x11, 0x65, 0x91, 0x69, 0x00, 0x00, 0x00,
}
