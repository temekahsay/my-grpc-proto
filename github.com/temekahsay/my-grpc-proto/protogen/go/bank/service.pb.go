// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/bank/service.proto

package bank

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

type CurrentBalanceRequest struct {
	AccountNumber        string   `protobuf:"bytes,1,opt,name=account_number,proto3" json:"account_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrentBalanceRequest) Reset()         { *m = CurrentBalanceRequest{} }
func (m *CurrentBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*CurrentBalanceRequest) ProtoMessage()    {}
func (*CurrentBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8fc67153972fff9, []int{0}
}

func (m *CurrentBalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentBalanceRequest.Unmarshal(m, b)
}
func (m *CurrentBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentBalanceRequest.Marshal(b, m, deterministic)
}
func (m *CurrentBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentBalanceRequest.Merge(m, src)
}
func (m *CurrentBalanceRequest) XXX_Size() int {
	return xxx_messageInfo_CurrentBalanceRequest.Size(m)
}
func (m *CurrentBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentBalanceRequest proto.InternalMessageInfo

func (m *CurrentBalanceRequest) GetAccountNumber() string {
	if m != nil {
		return m.AccountNumber
	}
	return ""
}

type CurrentBalanceResponse struct {
	Amount               float64  `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	CurrentDate          *Date    `protobuf:"bytes,2,opt,name=current_date,proto3" json:"current_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrentBalanceResponse) Reset()         { *m = CurrentBalanceResponse{} }
func (m *CurrentBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*CurrentBalanceResponse) ProtoMessage()    {}
func (*CurrentBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8fc67153972fff9, []int{1}
}

func (m *CurrentBalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentBalanceResponse.Unmarshal(m, b)
}
func (m *CurrentBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentBalanceResponse.Marshal(b, m, deterministic)
}
func (m *CurrentBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentBalanceResponse.Merge(m, src)
}
func (m *CurrentBalanceResponse) XXX_Size() int {
	return xxx_messageInfo_CurrentBalanceResponse.Size(m)
}
func (m *CurrentBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentBalanceResponse proto.InternalMessageInfo

func (m *CurrentBalanceResponse) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *CurrentBalanceResponse) GetCurrentDate() *Date {
	if m != nil {
		return m.CurrentDate
	}
	return nil
}

type Date struct {
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without
	// a year.
	Year int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a
	// month and day.
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
	// to specify a year by itself or a year and month where the day isn't
	// significant.
	Day                  int32    `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Date) Reset()         { *m = Date{} }
func (m *Date) String() string { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()    {}
func (*Date) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8fc67153972fff9, []int{2}
}

func (m *Date) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Date.Unmarshal(m, b)
}
func (m *Date) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Date.Marshal(b, m, deterministic)
}
func (m *Date) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Date.Merge(m, src)
}
func (m *Date) XXX_Size() int {
	return xxx_messageInfo_Date.Size(m)
}
func (m *Date) XXX_DiscardUnknown() {
	xxx_messageInfo_Date.DiscardUnknown(m)
}

var xxx_messageInfo_Date proto.InternalMessageInfo

func (m *Date) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Date) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Date) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func init() {
	proto.RegisterType((*CurrentBalanceRequest)(nil), "bank.CurrentBalanceRequest")
	proto.RegisterType((*CurrentBalanceResponse)(nil), "bank.CurrentBalanceResponse")
	proto.RegisterType((*Date)(nil), "bank.Date")
}

func init() {
	proto.RegisterFile("proto/bank/service.proto", fileDescriptor_f8fc67153972fff9)
}

var fileDescriptor_f8fc67153972fff9 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x8d, 0x4d, 0x0a, 0x4e, 0x45, 0x74, 0xd1, 0x12, 0xd4, 0x43, 0xc9, 0x41, 0x7a, 0x69,
	0x02, 0x55, 0xbc, 0x0a, 0x51, 0xf0, 0x2a, 0xeb, 0xcd, 0x4b, 0x9c, 0x6c, 0x87, 0xa4, 0xc4, 0xdd,
	0x8d, 0x9b, 0x8d, 0x90, 0x7f, 0x2f, 0xd9, 0xed, 0xa5, 0xa5, 0x97, 0x65, 0xe6, 0x9b, 0x9d, 0xe1,
	0xcd, 0x3c, 0x88, 0x5b, 0xa3, 0xad, 0xce, 0x4a, 0x54, 0x4d, 0xd6, 0x91, 0xf9, 0xdb, 0x0a, 0x4a,
	0x1d, 0x62, 0xe1, 0xc8, 0x92, 0x17, 0xb8, 0x79, 0xed, 0x8d, 0x21, 0x65, 0x73, 0xfc, 0x41, 0x25,
	0x88, 0xd3, 0x6f, 0x4f, 0x9d, 0x65, 0x0f, 0x70, 0x81, 0x42, 0xe8, 0x5e, 0xd9, 0x42, 0xf5, 0xb2,
	0x24, 0x13, 0x07, 0x8b, 0x60, 0x79, 0xc6, 0x0f, 0x68, 0xf2, 0x0d, 0xf3, 0xc3, 0x01, 0x5d, 0xab,
	0x55, 0x47, 0x6c, 0x0e, 0x53, 0x94, 0xe3, 0x57, 0xd7, 0x19, 0xf0, 0x5d, 0xc6, 0x52, 0x38, 0x17,
	0xbe, 0xa3, 0xd8, 0xa0, 0xa5, 0xf8, 0x74, 0x11, 0x2c, 0x67, 0x6b, 0x48, 0x47, 0x3d, 0xe9, 0x1b,
	0x5a, 0xe2, 0x7b, 0xf5, 0x24, 0x87, 0x70, 0xa4, 0x8c, 0x41, 0x38, 0x10, 0x7a, 0x1d, 0x11, 0x77,
	0x31, 0xbb, 0x86, 0x48, 0x6a, 0x65, 0x6b, 0x37, 0x24, 0xe2, 0x3e, 0x61, 0x97, 0x30, 0xd9, 0xe0,
	0x10, 0x4f, 0x1c, 0x1b, 0xc3, 0x75, 0x01, 0xb3, 0x1c, 0x55, 0xf3, 0xe9, 0x2f, 0xc0, 0x3e, 0xe0,
	0xea, 0x9d, 0xec, 0xbe, 0x6e, 0x76, 0xe7, 0x15, 0x1c, 0x3d, 0xc7, 0xed, 0xfd, 0xf1, 0xa2, 0x5f,
	0x35, 0x39, 0xc9, 0x9f, 0xbf, 0x9e, 0xaa, 0xad, 0xad, 0xfb, 0x32, 0x15, 0x5a, 0x66, 0x96, 0x24,
	0x35, 0x58, 0x77, 0x38, 0x64, 0x72, 0x58, 0x55, 0xa6, 0x15, 0x2b, 0xef, 0x83, 0x7b, 0x2b, 0x52,
	0x59, 0xe5, 0x3d, 0x29, 0xa7, 0x8e, 0x3c, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x00, 0x88, 0xec,
	0x7b, 0xa8, 0x01, 0x00, 0x00,
}
