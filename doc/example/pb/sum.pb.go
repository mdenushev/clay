// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sum.proto

package sumpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SumRequest is a request for Summator service.
type SumRequest struct {
	// A is the number we're adding to. Can't be zero for the sake of example.
	A int64 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	// B is the number we're adding.
	B                    int64    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sum_f9b1650a56ecf8db, []int{0}
}
func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (dst *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(dst, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SumRequest) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

type SumResponse struct {
	Sum                  int64    `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_sum_f9b1650a56ecf8db, []int{1}
}
func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (dst *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(dst, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetSum() int64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *SumResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "sumpb.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "sumpb.SumResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SummatorClient is the client API for Summator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SummatorClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type summatorClient struct {
	cc *grpc.ClientConn
}

func NewSummatorClient(cc *grpc.ClientConn) SummatorClient {
	return &summatorClient{cc}
}

func (c *summatorClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/sumpb.Summator/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SummatorServer is the server API for Summator service.
type SummatorServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
}

func RegisterSummatorServer(s *grpc.Server, srv SummatorServer) {
	s.RegisterService(&_Summator_serviceDesc, srv)
}

func _Summator_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummatorServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sumpb.Summator/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummatorServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Summator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sumpb.Summator",
	HandlerType: (*SummatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Summator_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sum.proto",
}

func init() { proto.RegisterFile("sum.proto", fileDescriptor_sum_f9b1650a56ecf8db) }

var fileDescriptor_sum_f9b1650a56ecf8db = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xb1, 0xce, 0x02, 0x21,
	0x0c, 0x80, 0x7f, 0xfe, 0xcb, 0x19, 0xaf, 0x3a, 0x68, 0xe3, 0x70, 0x71, 0x32, 0x4c, 0x37, 0x31,
	0x68, 0x5c, 0x7c, 0x0c, 0x78, 0x02, 0x48, 0x18, 0x2b, 0x48, 0xe9, 0xfb, 0x9b, 0x03, 0x13, 0xdd,
	0xfa, 0x35, 0x5f, 0xf3, 0x15, 0x26, 0x16, 0x32, 0xb9, 0xa4, 0x9a, 0x70, 0x64, 0xa1, 0x1c, 0xf4,
	0x02, 0xe0, 0x84, 0x6c, 0x7c, 0x49, 0xe4, 0x8a, 0x7b, 0x50, 0x7e, 0x56, 0x17, 0xb5, 0x0c, 0x56,
	0xf9, 0x95, 0xc2, 0xfc, 0xdf, 0x29, 0xe8, 0x3b, 0xec, 0x9a, 0xc9, 0x39, 0x3d, 0x39, 0xe2, 0x01,
	0x06, 0x16, 0xfa, 0xc8, 0xeb, 0x88, 0x27, 0x18, 0x63, 0x29, 0xa9, 0xb4, 0x93, 0xc9, 0x76, 0xb8,
	0x3e, 0x60, 0xeb, 0x84, 0xc8, 0xd7, 0x54, 0xd0, 0xc0, 0xe0, 0x84, 0xf0, 0x68, 0x5a, 0xdb, 0x7c,
	0xc3, 0x67, 0xfc, 0x5d, 0xf5, 0x82, 0xfe, 0x0b, 0x9b, 0xf6, 0xea, 0xed, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x98, 0x2c, 0x07, 0x08, 0xb7, 0x00, 0x00, 0x00,
}
