// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fix.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type VersionEnum int32

const (
	VersionEnum_VERSION_UNSPECIFIED VersionEnum = 0
	VersionEnum_VERSION_FIX_2_7     VersionEnum = 1
	VersionEnum_VERSION_FIX_3_0     VersionEnum = 2
	VersionEnum_VERSION_FIX_4_0     VersionEnum = 3
	VersionEnum_VERSION_FIX_4_1     VersionEnum = 4
	VersionEnum_VERSION_FIX_4_2     VersionEnum = 5
	VersionEnum_VERSION_FIX_4_3     VersionEnum = 6
	VersionEnum_VERSION_FIX_4_4     VersionEnum = 7
	VersionEnum_VERSION_FIX_5_0     VersionEnum = 8
	VersionEnum_VERSION_FIXT_1_1    VersionEnum = 9
	VersionEnum_VERSION_FIX_5_0SP1  VersionEnum = 10
	VersionEnum_VERSION_FIX_5_0SP2  VersionEnum = 11
)

var VersionEnum_name = map[int32]string{
	0:  "VERSION_UNSPECIFIED",
	1:  "VERSION_FIX_2_7",
	2:  "VERSION_FIX_3_0",
	3:  "VERSION_FIX_4_0",
	4:  "VERSION_FIX_4_1",
	5:  "VERSION_FIX_4_2",
	6:  "VERSION_FIX_4_3",
	7:  "VERSION_FIX_4_4",
	8:  "VERSION_FIX_5_0",
	9:  "VERSION_FIXT_1_1",
	10: "VERSION_FIX_5_0SP1",
	11: "VERSION_FIX_5_0SP2",
}

var VersionEnum_value = map[string]int32{
	"VERSION_UNSPECIFIED": 0,
	"VERSION_FIX_2_7":     1,
	"VERSION_FIX_3_0":     2,
	"VERSION_FIX_4_0":     3,
	"VERSION_FIX_4_1":     4,
	"VERSION_FIX_4_2":     5,
	"VERSION_FIX_4_3":     6,
	"VERSION_FIX_4_4":     7,
	"VERSION_FIX_5_0":     8,
	"VERSION_FIXT_1_1":    9,
	"VERSION_FIX_5_0SP1":  10,
	"VERSION_FIX_5_0SP2":  11,
}

func (x VersionEnum) String() string {
	return proto.EnumName(VersionEnum_name, int32(x))
}

func (VersionEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_09ed17ea1bedaf86, []int{0}
}

type DatatypeEnum int32

const (
	DatatypeEnum_DATATYPE_UNSPECIFIED           DatatypeEnum = 0
	DatatypeEnum_DATATYPE_CHAR                  DatatypeEnum = 1
	DatatypeEnum_DATATYPE_DATA                  DatatypeEnum = 2
	DatatypeEnum_DATATYPE_FLOAT                 DatatypeEnum = 3
	DatatypeEnum_DATATYPE_INT                   DatatypeEnum = 4
	DatatypeEnum_DATATYPE_DAY_OF_MONTH          DatatypeEnum = 5
	DatatypeEnum_DATATYPE_MONTH_YEAR            DatatypeEnum = 6
	DatatypeEnum_DATATYPE_AMT                   DatatypeEnum = 7
	DatatypeEnum_DATATYPE_BOOLEAN               DatatypeEnum = 8
	DatatypeEnum_DATATYPE_CURRENCY              DatatypeEnum = 9
	DatatypeEnum_DATATYPE_EXCHANGE              DatatypeEnum = 10
	DatatypeEnum_DATATYPE_LOCAL_MKT_DATE        DatatypeEnum = 11
	DatatypeEnum_DATATYPE_MULTIPLE_STRING_VALUE DatatypeEnum = 12
	DatatypeEnum_DATATYPE_PRICE                 DatatypeEnum = 13
	DatatypeEnum_DATATYPE_PRICE_OFFSET          DatatypeEnum = 14
	DatatypeEnum_DATATYPE_QTY                   DatatypeEnum = 15
	DatatypeEnum_DATATYPE_STRING                DatatypeEnum = 16
	DatatypeEnum_DATATYPE_UTC_TIMESTAMP         DatatypeEnum = 17
	DatatypeEnum_DATATYPE_UTC_TIME_ONLY         DatatypeEnum = 18
	DatatypeEnum_DATATYPE_LENGTH                DatatypeEnum = 19
	DatatypeEnum_DATATYPE_NUM_IN_GROUP          DatatypeEnum = 20
	DatatypeEnum_DATATYPE_PERCENTAGE            DatatypeEnum = 21
	DatatypeEnum_DATATYPE_SEQ_NUM               DatatypeEnum = 22
	DatatypeEnum_DATATYPE_TAG_NUM               DatatypeEnum = 23
	DatatypeEnum_DATATYPE_COUNTRY               DatatypeEnum = 24
	DatatypeEnum_DATATYPE_MULTIPLE_CHAR_VALUE   DatatypeEnum = 25
	DatatypeEnum_DATATYPE_PATTERN               DatatypeEnum = 26
	DatatypeEnum_DATATYPE_RESERVED1000PLUS      DatatypeEnum = 27
	DatatypeEnum_DATATYPE_RESERVED100PLUS       DatatypeEnum = 28
	DatatypeEnum_DATATYPE_RESERVED4000PLUS      DatatypeEnum = 29
	DatatypeEnum_DATATYPE_TENOR                 DatatypeEnum = 30
	DatatypeEnum_DATATYPE_TZ_TIMESTAMP          DatatypeEnum = 31
	DatatypeEnum_DATATYPE_TZ_TIME_ONLY          DatatypeEnum = 32
	DatatypeEnum_DATATYPE_UTC_DATE_ONLY         DatatypeEnum = 33
	DatatypeEnum_DATATYPE_XML_DATA              DatatypeEnum = 34
	DatatypeEnum_DATATYPE_LANGUAGE              DatatypeEnum = 35
)

var DatatypeEnum_name = map[int32]string{
	0:  "DATATYPE_UNSPECIFIED",
	1:  "DATATYPE_CHAR",
	2:  "DATATYPE_DATA",
	3:  "DATATYPE_FLOAT",
	4:  "DATATYPE_INT",
	5:  "DATATYPE_DAY_OF_MONTH",
	6:  "DATATYPE_MONTH_YEAR",
	7:  "DATATYPE_AMT",
	8:  "DATATYPE_BOOLEAN",
	9:  "DATATYPE_CURRENCY",
	10: "DATATYPE_EXCHANGE",
	11: "DATATYPE_LOCAL_MKT_DATE",
	12: "DATATYPE_MULTIPLE_STRING_VALUE",
	13: "DATATYPE_PRICE",
	14: "DATATYPE_PRICE_OFFSET",
	15: "DATATYPE_QTY",
	16: "DATATYPE_STRING",
	17: "DATATYPE_UTC_TIMESTAMP",
	18: "DATATYPE_UTC_TIME_ONLY",
	19: "DATATYPE_LENGTH",
	20: "DATATYPE_NUM_IN_GROUP",
	21: "DATATYPE_PERCENTAGE",
	22: "DATATYPE_SEQ_NUM",
	23: "DATATYPE_TAG_NUM",
	24: "DATATYPE_COUNTRY",
	25: "DATATYPE_MULTIPLE_CHAR_VALUE",
	26: "DATATYPE_PATTERN",
	27: "DATATYPE_RESERVED1000PLUS",
	28: "DATATYPE_RESERVED100PLUS",
	29: "DATATYPE_RESERVED4000PLUS",
	30: "DATATYPE_TENOR",
	31: "DATATYPE_TZ_TIMESTAMP",
	32: "DATATYPE_TZ_TIME_ONLY",
	33: "DATATYPE_UTC_DATE_ONLY",
	34: "DATATYPE_XML_DATA",
	35: "DATATYPE_LANGUAGE",
}

var DatatypeEnum_value = map[string]int32{
	"DATATYPE_UNSPECIFIED":           0,
	"DATATYPE_CHAR":                  1,
	"DATATYPE_DATA":                  2,
	"DATATYPE_FLOAT":                 3,
	"DATATYPE_INT":                   4,
	"DATATYPE_DAY_OF_MONTH":          5,
	"DATATYPE_MONTH_YEAR":            6,
	"DATATYPE_AMT":                   7,
	"DATATYPE_BOOLEAN":               8,
	"DATATYPE_CURRENCY":              9,
	"DATATYPE_EXCHANGE":              10,
	"DATATYPE_LOCAL_MKT_DATE":        11,
	"DATATYPE_MULTIPLE_STRING_VALUE": 12,
	"DATATYPE_PRICE":                 13,
	"DATATYPE_PRICE_OFFSET":          14,
	"DATATYPE_QTY":                   15,
	"DATATYPE_STRING":                16,
	"DATATYPE_UTC_TIMESTAMP":         17,
	"DATATYPE_UTC_TIME_ONLY":         18,
	"DATATYPE_LENGTH":                19,
	"DATATYPE_NUM_IN_GROUP":          20,
	"DATATYPE_PERCENTAGE":            21,
	"DATATYPE_SEQ_NUM":               22,
	"DATATYPE_TAG_NUM":               23,
	"DATATYPE_COUNTRY":               24,
	"DATATYPE_MULTIPLE_CHAR_VALUE":   25,
	"DATATYPE_PATTERN":               26,
	"DATATYPE_RESERVED1000PLUS":      27,
	"DATATYPE_RESERVED100PLUS":       28,
	"DATATYPE_RESERVED4000PLUS":      29,
	"DATATYPE_TENOR":                 30,
	"DATATYPE_TZ_TIMESTAMP":          31,
	"DATATYPE_TZ_TIME_ONLY":          32,
	"DATATYPE_UTC_DATE_ONLY":         33,
	"DATATYPE_XML_DATA":              34,
	"DATATYPE_LANGUAGE":              35,
}

func (x DatatypeEnum) String() string {
	return proto.EnumName(DatatypeEnum_name, int32(x))
}

func (DatatypeEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_09ed17ea1bedaf86, []int{1}
}

type Tenor struct {
	Days                 uint32   `protobuf:"varint,1,opt,name=days,proto3" json:"days,omitempty"`
	Weeks                uint32   `protobuf:"varint,2,opt,name=weeks,proto3" json:"weeks,omitempty"`
	Months               uint32   `protobuf:"varint,3,opt,name=months,proto3" json:"months,omitempty"`
	Years                uint32   `protobuf:"varint,4,opt,name=years,proto3" json:"years,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tenor) Reset()         { *m = Tenor{} }
func (m *Tenor) String() string { return proto.CompactTextString(m) }
func (*Tenor) ProtoMessage()    {}
func (*Tenor) Descriptor() ([]byte, []int) {
	return fileDescriptor_09ed17ea1bedaf86, []int{0}
}

func (m *Tenor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tenor.Unmarshal(m, b)
}
func (m *Tenor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tenor.Marshal(b, m, deterministic)
}
func (m *Tenor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tenor.Merge(m, src)
}
func (m *Tenor) XXX_Size() int {
	return xxx_messageInfo_Tenor.Size(m)
}
func (m *Tenor) XXX_DiscardUnknown() {
	xxx_messageInfo_Tenor.DiscardUnknown(m)
}

var xxx_messageInfo_Tenor proto.InternalMessageInfo

func (m *Tenor) GetDays() uint32 {
	if m != nil {
		return m.Days
	}
	return 0
}

func (m *Tenor) GetWeeks() uint32 {
	if m != nil {
		return m.Weeks
	}
	return 0
}

func (m *Tenor) GetMonths() uint32 {
	if m != nil {
		return m.Months
	}
	return 0
}

func (m *Tenor) GetYears() uint32 {
	if m != nil {
		return m.Years
	}
	return 0
}

type Decimal32 struct {
	Mantissa             int32    `protobuf:"fixed32,1,opt,name=mantissa,proto3" json:"mantissa,omitempty"`
	Exponent             int32    `protobuf:"fixed32,2,opt,name=exponent,proto3" json:"exponent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Decimal32) Reset()         { *m = Decimal32{} }
func (m *Decimal32) String() string { return proto.CompactTextString(m) }
func (*Decimal32) ProtoMessage()    {}
func (*Decimal32) Descriptor() ([]byte, []int) {
	return fileDescriptor_09ed17ea1bedaf86, []int{1}
}

func (m *Decimal32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Decimal32.Unmarshal(m, b)
}
func (m *Decimal32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Decimal32.Marshal(b, m, deterministic)
}
func (m *Decimal32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Decimal32.Merge(m, src)
}
func (m *Decimal32) XXX_Size() int {
	return xxx_messageInfo_Decimal32.Size(m)
}
func (m *Decimal32) XXX_DiscardUnknown() {
	xxx_messageInfo_Decimal32.DiscardUnknown(m)
}

var xxx_messageInfo_Decimal32 proto.InternalMessageInfo

func (m *Decimal32) GetMantissa() int32 {
	if m != nil {
		return m.Mantissa
	}
	return 0
}

func (m *Decimal32) GetExponent() int32 {
	if m != nil {
		return m.Exponent
	}
	return 0
}

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
	return fileDescriptor_09ed17ea1bedaf86, []int{2}
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

type Timestamp struct {
	Seconds              int64    `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos                int32    `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_09ed17ea1bedaf86, []int{3}
}

func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

type TimeOnly struct {
	Seconds              int64    `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos                int32    `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeOnly) Reset()         { *m = TimeOnly{} }
func (m *TimeOnly) String() string { return proto.CompactTextString(m) }
func (*TimeOnly) ProtoMessage()    {}
func (*TimeOnly) Descriptor() ([]byte, []int) {
	return fileDescriptor_09ed17ea1bedaf86, []int{4}
}

func (m *TimeOnly) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeOnly.Unmarshal(m, b)
}
func (m *TimeOnly) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeOnly.Marshal(b, m, deterministic)
}
func (m *TimeOnly) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeOnly.Merge(m, src)
}
func (m *TimeOnly) XXX_Size() int {
	return xxx_messageInfo_TimeOnly.Size(m)
}
func (m *TimeOnly) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeOnly.DiscardUnknown(m)
}

var xxx_messageInfo_TimeOnly proto.InternalMessageInfo

func (m *TimeOnly) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *TimeOnly) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

type LocalTimestamp struct {
	Date                 int32    `protobuf:"zigzag32,1,opt,name=date,proto3" json:"date,omitempty"`
	Hours                int32    `protobuf:"varint,2,opt,name=hours,proto3" json:"hours,omitempty"`
	Minutes              int32    `protobuf:"varint,3,opt,name=minutes,proto3" json:"minutes,omitempty"`
	Seconds              int32    `protobuf:"varint,4,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos                int32    `protobuf:"varint,5,opt,name=nanos,proto3" json:"nanos,omitempty"`
	HoursOffset          int32    `protobuf:"zigzag32,6,opt,name=hoursOffset,proto3" json:"hoursOffset,omitempty"`
	MinutesOffset        int32    `protobuf:"zigzag32,7,opt,name=minutesOffset,proto3" json:"minutesOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalTimestamp) Reset()         { *m = LocalTimestamp{} }
func (m *LocalTimestamp) String() string { return proto.CompactTextString(m) }
func (*LocalTimestamp) ProtoMessage()    {}
func (*LocalTimestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_09ed17ea1bedaf86, []int{5}
}

func (m *LocalTimestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalTimestamp.Unmarshal(m, b)
}
func (m *LocalTimestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalTimestamp.Marshal(b, m, deterministic)
}
func (m *LocalTimestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalTimestamp.Merge(m, src)
}
func (m *LocalTimestamp) XXX_Size() int {
	return xxx_messageInfo_LocalTimestamp.Size(m)
}
func (m *LocalTimestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalTimestamp.DiscardUnknown(m)
}

var xxx_messageInfo_LocalTimestamp proto.InternalMessageInfo

func (m *LocalTimestamp) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *LocalTimestamp) GetHours() int32 {
	if m != nil {
		return m.Hours
	}
	return 0
}

func (m *LocalTimestamp) GetMinutes() int32 {
	if m != nil {
		return m.Minutes
	}
	return 0
}

func (m *LocalTimestamp) GetSeconds() int32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *LocalTimestamp) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

func (m *LocalTimestamp) GetHoursOffset() int32 {
	if m != nil {
		return m.HoursOffset
	}
	return 0
}

func (m *LocalTimestamp) GetMinutesOffset() int32 {
	if m != nil {
		return m.MinutesOffset
	}
	return 0
}

type LocalTimeOnly struct {
	Hours                int32    `protobuf:"varint,1,opt,name=hours,proto3" json:"hours,omitempty"`
	Minutes              int32    `protobuf:"varint,2,opt,name=minutes,proto3" json:"minutes,omitempty"`
	Seconds              int32    `protobuf:"varint,3,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos                int32    `protobuf:"varint,4,opt,name=nanos,proto3" json:"nanos,omitempty"`
	HoursOffset          int32    `protobuf:"zigzag32,5,opt,name=hoursOffset,proto3" json:"hoursOffset,omitempty"`
	MinutesOffset        int32    `protobuf:"zigzag32,6,opt,name=minutesOffset,proto3" json:"minutesOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalTimeOnly) Reset()         { *m = LocalTimeOnly{} }
func (m *LocalTimeOnly) String() string { return proto.CompactTextString(m) }
func (*LocalTimeOnly) ProtoMessage()    {}
func (*LocalTimeOnly) Descriptor() ([]byte, []int) {
	return fileDescriptor_09ed17ea1bedaf86, []int{6}
}

func (m *LocalTimeOnly) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalTimeOnly.Unmarshal(m, b)
}
func (m *LocalTimeOnly) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalTimeOnly.Marshal(b, m, deterministic)
}
func (m *LocalTimeOnly) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalTimeOnly.Merge(m, src)
}
func (m *LocalTimeOnly) XXX_Size() int {
	return xxx_messageInfo_LocalTimeOnly.Size(m)
}
func (m *LocalTimeOnly) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalTimeOnly.DiscardUnknown(m)
}

var xxx_messageInfo_LocalTimeOnly proto.InternalMessageInfo

func (m *LocalTimeOnly) GetHours() int32 {
	if m != nil {
		return m.Hours
	}
	return 0
}

func (m *LocalTimeOnly) GetMinutes() int32 {
	if m != nil {
		return m.Minutes
	}
	return 0
}

func (m *LocalTimeOnly) GetSeconds() int32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *LocalTimeOnly) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

func (m *LocalTimeOnly) GetHoursOffset() int32 {
	if m != nil {
		return m.HoursOffset
	}
	return 0
}

func (m *LocalTimeOnly) GetMinutesOffset() int32 {
	if m != nil {
		return m.MinutesOffset
	}
	return 0
}

var E_Category = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         53002,
	Name:          "fix.category",
	Tag:           "bytes,53002,opt,name=category",
	Filename:      "fix.proto",
}

var E_MsgTypeValue = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         55001,
	Name:          "fix.msg_type_value",
	Tag:           "bytes,55001,opt,name=msg_type_value",
	Filename:      "fix.proto",
}

var E_Tag = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         56003,
	Name:          "fix.tag",
	Tag:           "fixed32,56003,opt,name=tag",
	Filename:      "fix.proto",
}

var E_Type = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*DatatypeEnum)(nil),
	Field:         56004,
	Name:          "fix.type",
	Tag:           "varint,56004,opt,name=type,enum=fix.DatatypeEnum",
	Filename:      "fix.proto",
}

var E_FieldAdded = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*VersionEnum)(nil),
	Field:         56005,
	Name:          "fix.field_added",
	Tag:           "varint,56005,opt,name=field_added,enum=fix.VersionEnum",
	Filename:      "fix.proto",
}

var E_FieldAddedEp = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         56006,
	Name:          "fix.field_added_ep",
	Tag:           "fixed32,56006,opt,name=field_added_ep",
	Filename:      "fix.proto",
}

var E_FieldDeprecated = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*VersionEnum)(nil),
	Field:         56007,
	Name:          "fix.field_deprecated",
	Tag:           "varint,56007,opt,name=field_deprecated,enum=fix.VersionEnum",
	Filename:      "fix.proto",
}

var E_GroupTag = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         56008,
	Name:          "fix.group_tag",
	Tag:           "fixed32,56008,opt,name=group_tag",
	Filename:      "fix.proto",
}

var E_EnumValue = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         72004,
	Name:          "fix.enum_value",
	Tag:           "bytes,72004,opt,name=enum_value",
	Filename:      "fix.proto",
}

var E_EnumAdded = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*VersionEnum)(nil),
	Field:         72005,
	Name:          "fix.enum_added",
	Tag:           "varint,72005,opt,name=enum_added,enum=fix.VersionEnum",
	Filename:      "fix.proto",
}

var E_EnumAddedEp = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         72006,
	Name:          "fix.enum_added_ep",
	Tag:           "fixed32,72006,opt,name=enum_added_ep",
	Filename:      "fix.proto",
}

var E_EnumDeprecated = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*VersionEnum)(nil),
	Field:         72007,
	Name:          "fix.enum_deprecated",
	Tag:           "varint,72007,opt,name=enum_deprecated,enum=fix.VersionEnum",
	Filename:      "fix.proto",
}

func init() {
	proto.RegisterEnum("fix.VersionEnum", VersionEnum_name, VersionEnum_value)
	proto.RegisterEnum("fix.DatatypeEnum", DatatypeEnum_name, DatatypeEnum_value)
	proto.RegisterType((*Tenor)(nil), "fix.Tenor")
	proto.RegisterType((*Decimal32)(nil), "fix.Decimal32")
	proto.RegisterType((*Decimal64)(nil), "fix.Decimal64")
	proto.RegisterType((*Timestamp)(nil), "fix.Timestamp")
	proto.RegisterType((*TimeOnly)(nil), "fix.TimeOnly")
	proto.RegisterType((*LocalTimestamp)(nil), "fix.LocalTimestamp")
	proto.RegisterType((*LocalTimeOnly)(nil), "fix.LocalTimeOnly")
	proto.RegisterExtension(E_Category)
	proto.RegisterExtension(E_MsgTypeValue)
	proto.RegisterExtension(E_Tag)
	proto.RegisterExtension(E_Type)
	proto.RegisterExtension(E_FieldAdded)
	proto.RegisterExtension(E_FieldAddedEp)
	proto.RegisterExtension(E_FieldDeprecated)
	proto.RegisterExtension(E_GroupTag)
	proto.RegisterExtension(E_EnumValue)
	proto.RegisterExtension(E_EnumAdded)
	proto.RegisterExtension(E_EnumAddedEp)
	proto.RegisterExtension(E_EnumDeprecated)
}

func init() { proto.RegisterFile("fix.proto", fileDescriptor_09ed17ea1bedaf86) }

var fileDescriptor_09ed17ea1bedaf86 = []byte{
	// 1234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0x4d, 0x6f, 0xdb, 0x46,
	0x13, 0xc7, 0x1f, 0x59, 0x92, 0x2d, 0x8d, 0x6d, 0x69, 0x4d, 0x3b, 0xb6, 0xe2, 0xbc, 0x39, 0xce,
	0x83, 0x22, 0xc8, 0xc1, 0xf1, 0x5b, 0x5b, 0x80, 0xed, 0x85, 0x96, 0x57, 0xb2, 0x50, 0x8a, 0x54,
	0x56, 0x2b, 0x23, 0x0a, 0x8a, 0x2e, 0x18, 0x69, 0xad, 0x08, 0x95, 0x48, 0x41, 0xa4, 0x5a, 0xfb,
	0x2b, 0xf4, 0x9c, 0x73, 0xcf, 0xbd, 0xf4, 0x52, 0xa0, 0xe8, 0x07, 0x68, 0x93, 0xf4, 0xd0, 0x53,
	0x6f, 0xfd, 0x36, 0xc5, 0x2e, 0x57, 0x14, 0x6d, 0xc9, 0xb1, 0x7b, 0xe3, 0xce, 0xff, 0x3f, 0xbf,
	0x99, 0x9d, 0xa1, 0x2d, 0x42, 0xf6, 0xac, 0x7b, 0xbe, 0x33, 0x18, 0x7a, 0x81, 0xa7, 0x25, 0xcf,
	0xba, 0xe7, 0x9b, 0x5b, 0x1d, 0xcf, 0xeb, 0xf4, 0xf8, 0x73, 0x19, 0x7a, 0x3d, 0x3a, 0x7b, 0xde,
	0xe6, 0x7e, 0x6b, 0xd8, 0x1d, 0x04, 0xde, 0x30, 0xb4, 0x6d, 0x33, 0x48, 0x53, 0xee, 0x7a, 0x43,
	0x4d, 0x83, 0x54, 0xdb, 0xb9, 0xf0, 0x0b, 0x89, 0xad, 0xc4, 0xd3, 0x65, 0x22, 0x9f, 0xb5, 0x35,
	0x48, 0x7f, 0xcf, 0xf9, 0xb7, 0x7e, 0x61, 0x4e, 0x06, 0xc3, 0x83, 0xb6, 0x0e, 0xf3, 0x7d, 0xcf,
	0x0d, 0xde, 0xf8, 0x85, 0xa4, 0x0c, 0xab, 0x93, 0x70, 0x5f, 0x70, 0x67, 0xe8, 0x17, 0x52, 0xa1,
	0x5b, 0x1e, 0xb6, 0x8b, 0x90, 0x3d, 0xe6, 0xad, 0x6e, 0xdf, 0xe9, 0x1d, 0xec, 0x6b, 0x9b, 0x90,
	0xe9, 0x3b, 0x6e, 0xd0, 0xf5, 0x7d, 0x47, 0x16, 0xca, 0x93, 0xe8, 0x2c, 0x34, 0x7e, 0x3e, 0xf0,
	0x5c, 0xee, 0x06, 0xb2, 0x5e, 0x9e, 0x44, 0xe7, 0x18, 0xe4, 0xb3, 0xc3, 0x29, 0x08, 0xba, 0x25,
	0xe4, 0x0b, 0xc8, 0xd2, 0x6e, 0x9f, 0xfb, 0x81, 0xd3, 0x1f, 0x68, 0x05, 0x58, 0xf0, 0x79, 0xcb,
	0x73, 0xdb, 0xe1, 0x8d, 0x93, 0x64, 0x7c, 0x14, 0xd7, 0x70, 0x1d, 0xd7, 0x0b, 0x2f, 0x9d, 0x26,
	0xe1, 0x61, 0x5b, 0x87, 0x8c, 0x48, 0xb6, 0xdd, 0xde, 0xc5, 0x7f, 0xce, 0xfd, 0x2b, 0x01, 0x39,
	0xd3, 0x6b, 0x39, 0xbd, 0x49, 0x79, 0x39, 0xed, 0x80, 0xcb, 0xfc, 0x15, 0x22, 0x9f, 0x45, 0xf2,
	0x1b, 0x6f, 0x34, 0x8c, 0x92, 0xe5, 0x41, 0x14, 0xeb, 0x77, 0xdd, 0x51, 0xc0, 0xc3, 0x71, 0xa7,
	0xc9, 0xf8, 0x18, 0x6f, 0x23, 0x15, 0x2a, 0x53, 0x6d, 0xa4, 0x63, 0x6d, 0x68, 0x5b, 0xb0, 0x28,
	0x91, 0xf6, 0xd9, 0x99, 0xcf, 0x83, 0xc2, 0xbc, 0x2c, 0x1d, 0x0f, 0x69, 0xff, 0x87, 0x65, 0x05,
	0x57, 0x9e, 0x05, 0xe9, 0xb9, 0x1c, 0xdc, 0xfe, 0x2d, 0x01, 0xcb, 0xd1, 0x75, 0xe4, 0x40, 0xa2,
	0xce, 0x13, 0xd7, 0x74, 0x3e, 0x77, 0x6d, 0xe7, 0xc9, 0x6b, 0x3a, 0x4f, 0x7d, 0xa4, 0xf3, 0xf4,
	0x2d, 0x3a, 0x9f, 0x9f, 0xd1, 0xf9, 0xb3, 0xb7, 0x73, 0xb0, 0x78, 0xca, 0x87, 0x7e, 0xd7, 0x73,
	0xb1, 0x3b, 0xea, 0x6b, 0x1b, 0xb0, 0x7a, 0x8a, 0x49, 0xbd, 0x62, 0x5b, 0xac, 0x61, 0xd5, 0x6b,
	0xb8, 0x58, 0x29, 0x55, 0xf0, 0x31, 0xfa, 0x9f, 0xb6, 0x0a, 0xf9, 0xb1, 0x50, 0xaa, 0xbc, 0x64,
	0xfb, 0xec, 0x73, 0x94, 0xb8, 0x1a, 0x3c, 0x60, 0xbb, 0x68, 0xee, 0x6a, 0xf0, 0x90, 0xed, 0xa2,
	0xe4, 0x74, 0x70, 0x0f, 0xa5, 0xa6, 0x83, 0xfb, 0x28, 0x3d, 0x1d, 0x3c, 0x40, 0xf3, 0xd3, 0xc1,
	0x43, 0xb4, 0x70, 0x35, 0xf8, 0x29, 0xdb, 0x45, 0x19, 0x6d, 0x0d, 0x50, 0x2c, 0x48, 0xd9, 0x1e,
	0xdb, 0x43, 0x59, 0x6d, 0x1d, 0xb4, 0x2b, 0xd6, 0x7a, 0x6d, 0x0f, 0xc1, 0xcc, 0xf8, 0x3e, 0x5a,
	0x7c, 0xf6, 0x77, 0x06, 0x96, 0x8e, 0x9d, 0xc0, 0x09, 0x2e, 0x06, 0x5c, 0xce, 0xa5, 0x00, 0x6b,
	0xc7, 0x06, 0x35, 0x68, 0xb3, 0x86, 0xaf, 0x0c, 0x66, 0x03, 0x96, 0x23, 0xa5, 0x78, 0x62, 0x10,
	0x94, 0xd8, 0x4c, 0xfd, 0xf4, 0xf6, 0x49, 0xe2, 0x92, 0x20, 0x1e, 0xd0, 0x9c, 0x12, 0x0a, 0x90,
	0x8b, 0x84, 0x92, 0x69, 0x1b, 0x14, 0x25, 0x95, 0xb2, 0x0e, 0x4b, 0x91, 0x52, 0xb1, 0x28, 0x4a,
	0xa9, 0xf8, 0x23, 0xb8, 0x13, 0x43, 0x35, 0x99, 0x5d, 0x62, 0x55, 0xdb, 0xa2, 0x27, 0x28, 0x2d,
	0x0d, 0x29, 0xed, 0x01, 0xac, 0x46, 0x06, 0xa9, 0xb0, 0x26, 0x36, 0x08, 0x9a, 0x57, 0x72, 0x9c,
	0x6b, 0x54, 0x29, 0x5a, 0x90, 0xf1, 0xb4, 0xb6, 0x09, 0x28, 0x8a, 0x1f, 0xd9, 0xb6, 0x89, 0x0d,
	0x0b, 0x65, 0x94, 0x76, 0x0f, 0x56, 0x26, 0xf7, 0x6a, 0x10, 0x82, 0xad, 0x62, 0x13, 0x65, 0x67,
	0x88, 0xf8, 0x65, 0xf1, 0xc4, 0xb0, 0xca, 0x18, 0x81, 0x12, 0x1f, 0xc3, 0x46, 0x24, 0x9a, 0x76,
	0xd1, 0x30, 0x59, 0xf5, 0x2b, 0x2a, 0x46, 0x80, 0xd1, 0xa2, 0xb2, 0x3c, 0x85, 0x87, 0x93, 0x7e,
	0x1b, 0x26, 0xad, 0xd4, 0x4c, 0xcc, 0xea, 0x94, 0x54, 0xac, 0x32, 0x3b, 0x35, 0xcc, 0x06, 0x46,
	0x4b, 0xca, 0x19, 0x1f, 0x56, 0x8d, 0x54, 0x8a, 0x18, 0x2d, 0x2b, 0x25, 0x3e, 0x14, 0xa9, 0x30,
	0xbb, 0x54, 0xaa, 0x63, 0x8a, 0x72, 0xca, 0x10, 0xbf, 0xf5, 0x0b, 0xda, 0x44, 0x79, 0x15, 0xbf,
	0x0b, 0xf9, 0x28, 0x1e, 0xd6, 0x44, 0x48, 0x49, 0x5b, 0xb0, 0x3e, 0x59, 0x33, 0x2d, 0x32, 0x5a,
	0xa9, 0xe2, 0x3a, 0x35, 0xaa, 0x35, 0xb4, 0xf2, 0x11, 0x07, 0xb3, 0x2d, 0xb3, 0x89, 0xb4, 0x19,
	0x78, 0x13, 0x5b, 0x65, 0x7a, 0x82, 0x56, 0xa5, 0x34, 0x7f, 0xa9, 0x65, 0xab, 0x51, 0x65, 0x15,
	0x8b, 0x95, 0x89, 0xdd, 0xa8, 0xa1, 0x35, 0x65, 0x88, 0xef, 0xb1, 0x86, 0x49, 0x11, 0x5b, 0xd4,
	0x28, 0x63, 0x74, 0x47, 0xc9, 0xf1, 0x7d, 0xd5, 0xf1, 0x0b, 0xc1, 0x40, 0xeb, 0x33, 0x34, 0x6a,
	0x94, 0xa5, 0xb6, 0x31, 0x43, 0x2b, 0xda, 0x0d, 0x8b, 0x92, 0x26, 0x2a, 0x48, 0x6d, 0x41, 0xfb,
	0x04, 0xee, 0x4f, 0xaf, 0x42, 0xbc, 0xc8, 0x6a, 0x11, 0x77, 0x95, 0x2f, 0xce, 0xa8, 0x19, 0x94,
	0x62, 0x62, 0xa1, 0x4d, 0xa5, 0x3d, 0x81, 0xbb, 0x91, 0x46, 0x70, 0x1d, 0x93, 0x53, 0x7c, 0xbc,
	0xb7, 0xbb, 0xbb, 0x5b, 0x33, 0x1b, 0x75, 0x74, 0x4f, 0x99, 0xb6, 0xa1, 0x30, 0xcb, 0x24, 0x3d,
	0xf7, 0x3f, 0x02, 0x3a, 0x1c, 0x83, 0x1e, 0x28, 0x53, 0xfc, 0x95, 0xa0, 0xd8, 0xb2, 0x09, 0x7a,
	0xa8, 0x94, 0xf8, 0x7c, 0xe9, 0xab, 0xd8, 0xf6, 0x1e, 0x5d, 0x6f, 0x08, 0x97, 0xb7, 0xa5, 0x0c,
	0x57, 0xd7, 0x2b, 0xde, 0xda, 0xd0, 0xf1, 0x58, 0x39, 0xe2, 0xaf, 0xfe, 0xcb, 0xaa, 0x19, 0xfe,
	0x69, 0x6f, 0x4b, 0x31, 0x73, 0x49, 0x34, 0x0d, 0xab, 0xdc, 0x10, 0xdb, 0x7b, 0x22, 0x45, 0xd0,
	0x75, 0xc8, 0xb4, 0x9c, 0x80, 0x77, 0xbc, 0xe1, 0x85, 0x76, 0x7f, 0x27, 0xfc, 0x0e, 0xd9, 0x19,
	0x7f, 0x87, 0xec, 0x94, 0xba, 0x3d, 0x6e, 0x0f, 0x82, 0xae, 0xe7, 0xfa, 0x85, 0x1f, 0x7e, 0x14,
	0x3f, 0x02, 0x59, 0x12, 0xf9, 0xf5, 0x32, 0xe4, 0xfa, 0x7e, 0x87, 0x89, 0xff, 0x47, 0xec, 0x3b,
	0xa7, 0x37, 0xe2, 0xda, 0xa3, 0x29, 0x42, 0x95, 0xfb, 0xbe, 0xd3, 0x89, 0x20, 0xff, 0xfc, 0x1c,
	0x42, 0x96, 0xfa, 0x7e, 0x87, 0x5e, 0x0c, 0xf8, 0xa9, 0x48, 0xd3, 0xf7, 0x20, 0x19, 0x38, 0x1d,
	0xed, 0xc1, 0x8c, 0xfa, 0xbc, 0xd7, 0x1e, 0xe7, 0xfe, 0xfe, 0xab, 0xc8, 0x5d, 0x20, 0xc2, 0xab,
	0x97, 0x20, 0x25, 0xea, 0xde, 0x94, 0xf3, 0x87, 0xcc, 0xc9, 0xed, 0xaf, 0xec, 0x88, 0x4f, 0xae,
	0xf8, 0x7f, 0x4f, 0x22, 0xf3, 0x75, 0x02, 0x8b, 0x67, 0x22, 0x81, 0x39, 0xed, 0x36, 0x6f, 0xdf,
	0x84, 0x7b, 0xa7, 0x70, 0x48, 0xe2, 0x62, 0xbf, 0x51, 0x04, 0x24, 0xc5, 0x10, 0x10, 0x1d, 0x43,
	0x2e, 0xc6, 0x64, 0x7c, 0x70, 0x13, 0xf6, 0xbd, 0xc4, 0xe6, 0xc9, 0xd2, 0x04, 0x82, 0x07, 0xfa,
	0xd7, 0x80, 0x42, 0x4c, 0x9b, 0x0f, 0x86, 0x5c, 0x4c, 0xfd, 0xc6, 0xfe, 0x3e, 0x5c, 0xdb, 0x5f,
	0x5e, 0xa2, 0x8e, 0x23, 0x92, 0xfe, 0x25, 0x64, 0x3b, 0x43, 0x6f, 0x34, 0x60, 0xb7, 0x98, 0xfc,
	0x9f, 0x6a, 0xf2, 0x19, 0x99, 0x41, 0x9d, 0x8e, 0x7e, 0x04, 0xc0, 0xdd, 0x51, 0x5f, 0xad, 0xfd,
	0xf1, 0x54, 0xba, 0xa8, 0x29, 0x77, 0x1b, 0x2d, 0xe2, 0x97, 0x94, 0x5c, 0x7c, 0x96, 0x8f, 0x15,
	0x9d, 0x2a, 0x46, 0x38, 0xf9, 0x5b, 0x30, 0xde, 0x49, 0xc6, 0xac, 0xdb, 0x49, 0x6a, 0x38, 0xfc,
	0x32, 0x2c, 0x4f, 0xa8, 0x62, 0xf6, 0xb7, 0x00, 0xbf, 0x97, 0xe0, 0x3c, 0x59, 0x8c, 0x30, 0x78,
	0xa0, 0x7f, 0x03, 0x79, 0x09, 0x8a, 0x4d, 0xff, 0x16, 0xa8, 0x0f, 0xd7, 0xf6, 0x98, 0x13, 0xb4,
	0xc9, 0x02, 0x8e, 0x9e, 0xc2, 0xa6, 0x37, 0xec, 0x08, 0x9b, 0xa4, 0xb5, 0xbc, 0xde, 0x4e, 0xcb,
	0xeb, 0x87, 0x1f, 0xc1, 0xfe, 0x51, 0xb2, 0xd4, 0x3d, 0x7f, 0x95, 0xee, 0x7b, 0x6d, 0xde, 0x7b,
	0x3d, 0x2f, 0x0d, 0x07, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xd9, 0x48, 0x1e, 0x37, 0x0c,
	0x00, 0x00,
}
