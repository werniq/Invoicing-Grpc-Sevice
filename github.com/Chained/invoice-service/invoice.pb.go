// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/invoice.proto

package invoice_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type InvoiceRequest struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity             int64                `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Amount               int64                `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Product              string               `protobuf:"bytes,4,opt,name=product,proto3" json:"product,omitempty"`
	Firstname            string               `protobuf:"bytes,5,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string               `protobuf:"bytes,6,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email                string               `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *InvoiceRequest) Reset()         { *m = InvoiceRequest{} }
func (m *InvoiceRequest) String() string { return proto.CompactTextString(m) }
func (*InvoiceRequest) ProtoMessage()    {}
func (*InvoiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_716db03bb6e01030, []int{0}
}

func (m *InvoiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceRequest.Unmarshal(m, b)
}
func (m *InvoiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceRequest.Marshal(b, m, deterministic)
}
func (m *InvoiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceRequest.Merge(m, src)
}
func (m *InvoiceRequest) XXX_Size() int {
	return xxx_messageInfo_InvoiceRequest.Size(m)
}
func (m *InvoiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceRequest proto.InternalMessageInfo

func (m *InvoiceRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InvoiceRequest) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *InvoiceRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *InvoiceRequest) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *InvoiceRequest) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *InvoiceRequest) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *InvoiceRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *InvoiceRequest) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type InvoiceResponse struct {
	Error                bool     `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvoiceResponse) Reset()         { *m = InvoiceResponse{} }
func (m *InvoiceResponse) String() string { return proto.CompactTextString(m) }
func (*InvoiceResponse) ProtoMessage()    {}
func (*InvoiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_716db03bb6e01030, []int{1}
}

func (m *InvoiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceResponse.Unmarshal(m, b)
}
func (m *InvoiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceResponse.Marshal(b, m, deterministic)
}
func (m *InvoiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceResponse.Merge(m, src)
}
func (m *InvoiceResponse) XXX_Size() int {
	return xxx_messageInfo_InvoiceResponse.Size(m)
}
func (m *InvoiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceResponse proto.InternalMessageInfo

func (m *InvoiceResponse) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *InvoiceResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*InvoiceRequest)(nil), "invoice.InvoiceRequest")
	proto.RegisterType((*InvoiceResponse)(nil), "invoice.InvoiceResponse")
}

func init() {
	proto.RegisterFile("proto/invoice.proto", fileDescriptor_716db03bb6e01030)
}

var fileDescriptor_716db03bb6e01030 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4b, 0xf3, 0x30,
	0x18, 0xc7, 0xdf, 0x6e, 0xef, 0x7e, 0x34, 0xc2, 0x06, 0x51, 0x34, 0x14, 0xc1, 0x51, 0x3c, 0xec,
	0x62, 0x0b, 0xf3, 0xe4, 0x71, 0x0a, 0x8a, 0xd7, 0xea, 0x41, 0xbc, 0x48, 0xd6, 0x3e, 0xeb, 0x02,
	0x4d, 0xd2, 0x25, 0x4f, 0x07, 0xfe, 0xe9, 0xde, 0x64, 0x49, 0x5b, 0x11, 0x8f, 0x9f, 0xe7, 0xfb,
	0xc0, 0x93, 0xcf, 0x37, 0xe4, 0xb4, 0x36, 0x1a, 0x75, 0x2a, 0xd4, 0x41, 0x8b, 0x1c, 0x12, 0x47,
	0x74, 0xd2, 0x62, 0x74, 0x55, 0x6a, 0x5d, 0x56, 0x90, 0xba, 0xf1, 0xa6, 0xd9, 0xa6, 0x28, 0x24,
	0x58, 0xe4, 0xb2, 0xf6, 0x9b, 0xf1, 0x57, 0x40, 0x66, 0xcf, 0x7e, 0x39, 0x83, 0x7d, 0x03, 0x16,
	0xe9, 0x8c, 0x0c, 0x44, 0xc1, 0x82, 0x45, 0xb0, 0x1c, 0x66, 0x03, 0x51, 0xd0, 0x88, 0x4c, 0xf7,
	0x0d, 0x57, 0x28, 0xf0, 0x93, 0x0d, 0xdc, 0xb4, 0x67, 0x7a, 0x4e, 0xc6, 0x5c, 0xea, 0x46, 0x21,
	0x1b, 0xba, 0xa4, 0x25, 0xca, 0xc8, 0xa4, 0x36, 0xba, 0x68, 0x72, 0x64, 0xff, 0x17, 0xc1, 0x32,
	0xcc, 0x3a, 0xa4, 0x97, 0x24, 0xdc, 0x0a, 0x63, 0x51, 0x71, 0x09, 0x6c, 0xe4, 0xb2, 0x9f, 0xc1,
	0xf1, 0x56, 0xc5, 0xdb, 0x70, 0xec, 0xc2, 0x9e, 0xe9, 0x19, 0x19, 0x81, 0xe4, 0xa2, 0x62, 0x13,
	0x17, 0x78, 0xa0, 0x77, 0x84, 0xe4, 0x06, 0x38, 0x42, 0xf1, 0xc1, 0x91, 0x4d, 0x17, 0xc1, 0xf2,
	0x64, 0x15, 0x25, 0x5e, 0x3b, 0xe9, 0xb4, 0x93, 0xd7, 0x4e, 0x3b, 0x0b, 0xdb, 0xed, 0x35, 0xc6,
	0x6b, 0x32, 0xef, 0xd5, 0x6d, 0xad, 0x95, 0xf5, 0x37, 0x8c, 0xd1, 0xc6, 0xe9, 0x4f, 0x33, 0x0f,
	0x47, 0x1b, 0x09, 0xd6, 0xf2, 0x12, 0x5c, 0x01, 0x61, 0xd6, 0xe1, 0xea, 0xad, 0x6f, 0xef, 0x05,
	0xcc, 0x41, 0xe4, 0x40, 0x1f, 0xc9, 0xfc, 0x09, 0x14, 0x18, 0x8e, 0xd0, 0x26, 0xf4, 0x22, 0xe9,
	0x7e, 0xe7, 0x77, 0xd3, 0x11, 0xfb, 0x1b, 0xf8, 0x77, 0xc4, 0xff, 0xee, 0xaf, 0xdf, 0xe3, 0x52,
	0xe0, 0xae, 0xd9, 0x24, 0xb9, 0x96, 0xe9, 0xc3, 0x8e, 0x0b, 0x05, 0x45, 0xf7, 0xcd, 0x37, 0xd6,
	0x5f, 0xdb, 0x8c, 0x9d, 0xe1, 0xed, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0xd1, 0x1b, 0x53,
	0x06, 0x02, 0x00, 0x00,
}
