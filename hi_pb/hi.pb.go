// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hi.proto

package higrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type HiRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HiRequest) Reset()         { *m = HiRequest{} }
func (m *HiRequest) String() string { return proto.CompactTextString(m) }
func (*HiRequest) ProtoMessage()    {}
func (*HiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d092a8920edeec73, []int{0}
}

func (m *HiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiRequest.Unmarshal(m, b)
}
func (m *HiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiRequest.Marshal(b, m, deterministic)
}
func (m *HiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiRequest.Merge(m, src)
}
func (m *HiRequest) XXX_Size() int {
	return xxx_messageInfo_HiRequest.Size(m)
}
func (m *HiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HiRequest proto.InternalMessageInfo

func (m *HiRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HiReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HiReply) Reset()         { *m = HiReply{} }
func (m *HiReply) String() string { return proto.CompactTextString(m) }
func (*HiReply) ProtoMessage()    {}
func (*HiReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d092a8920edeec73, []int{1}
}

func (m *HiReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiReply.Unmarshal(m, b)
}
func (m *HiReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiReply.Marshal(b, m, deterministic)
}
func (m *HiReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiReply.Merge(m, src)
}
func (m *HiReply) XXX_Size() int {
	return xxx_messageInfo_HiReply.Size(m)
}
func (m *HiReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HiReply.DiscardUnknown(m)
}

var xxx_messageInfo_HiReply proto.InternalMessageInfo

func (m *HiReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HiRequest)(nil), "higrpc.HiRequest")
	proto.RegisterType((*HiReply)(nil), "higrpc.HiReply")
}

func init() { proto.RegisterFile("hi.proto", fileDescriptor_d092a8920edeec73) }

var fileDescriptor_d092a8920edeec73 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0xc8, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0xc8, 0x4c, 0x2f, 0x2a, 0x48, 0x56, 0x92, 0xe7, 0xe2,
	0xf4, 0xc8, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc,
	0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x94, 0xb9, 0xd8, 0x41, 0x0a,
	0x0a, 0x72, 0x2a, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x61, 0x2a, 0x60,
	0x5c, 0xa3, 0x14, 0x2e, 0x26, 0x8f, 0x4c, 0x21, 0x5d, 0x2e, 0xd6, 0xe0, 0xc4, 0x4a, 0x8f, 0x4c,
	0x21, 0x41, 0x3d, 0x88, 0xe9, 0x7a, 0x70, 0xa3, 0xa5, 0xf8, 0x91, 0x85, 0x0a, 0x72, 0x2a, 0x95,
	0x18, 0x84, 0xf4, 0xb9, 0xd8, 0x9d, 0xf3, 0x4b, 0xf3, 0x4a, 0x88, 0xd5, 0x90, 0xc4, 0x06, 0x76,
	0xba, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x32, 0xe4, 0x84, 0xe9, 0xc6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HiClient is the client API for Hi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HiClient interface {
	SayHi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiReply, error)
	CountHi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiReply, error)
}

type hiClient struct {
	cc *grpc.ClientConn
}

func NewHiClient(cc *grpc.ClientConn) HiClient {
	return &hiClient{cc}
}

func (c *hiClient) SayHi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiReply, error) {
	out := new(HiReply)
	err := c.cc.Invoke(ctx, "/higrpc.Hi/SayHi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hiClient) CountHi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiReply, error) {
	out := new(HiReply)
	err := c.cc.Invoke(ctx, "/higrpc.Hi/CountHi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HiServer is the server API for Hi service.
type HiServer interface {
	SayHi(context.Context, *HiRequest) (*HiReply, error)
	CountHi(context.Context, *HiRequest) (*HiReply, error)
}

// UnimplementedHiServer can be embedded to have forward compatible implementations.
type UnimplementedHiServer struct {
}

func (*UnimplementedHiServer) SayHi(ctx context.Context, req *HiRequest) (*HiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHi not implemented")
}
func (*UnimplementedHiServer) CountHi(ctx context.Context, req *HiRequest) (*HiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountHi not implemented")
}

func RegisterHiServer(s *grpc.Server, srv HiServer) {
	s.RegisterService(&_Hi_serviceDesc, srv)
}

func _Hi_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HiServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/higrpc.Hi/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HiServer).SayHi(ctx, req.(*HiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hi_CountHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HiServer).CountHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/higrpc.Hi/CountHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HiServer).CountHi(ctx, req.(*HiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "higrpc.Hi",
	HandlerType: (*HiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _Hi_SayHi_Handler,
		},
		{
			MethodName: "CountHi",
			Handler:    _Hi_CountHi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hi.proto",
}
