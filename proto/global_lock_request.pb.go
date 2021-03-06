// Code generated by protoc-gen-go. DO NOT EDIT.
// source: global_lock_request.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// The request message containing the user's name.
type LockRequest struct {
	LockId               string   `protobuf:"bytes,1,opt,name=lock_id,json=lockId,proto3" json:"lock_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockRequest) Reset()         { *m = LockRequest{} }
func (m *LockRequest) String() string { return proto.CompactTextString(m) }
func (*LockRequest) ProtoMessage()    {}
func (*LockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c1f1b09c6e3ce14, []int{0}
}

func (m *LockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockRequest.Unmarshal(m, b)
}
func (m *LockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockRequest.Marshal(b, m, deterministic)
}
func (m *LockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockRequest.Merge(m, src)
}
func (m *LockRequest) XXX_Size() int {
	return xxx_messageInfo_LockRequest.Size(m)
}
func (m *LockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LockRequest proto.InternalMessageInfo

func (m *LockRequest) GetLockId() string {
	if m != nil {
		return m.LockId
	}
	return ""
}

// The response message containing the greetings
type LockReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockReply) Reset()         { *m = LockReply{} }
func (m *LockReply) String() string { return proto.CompactTextString(m) }
func (*LockReply) ProtoMessage()    {}
func (*LockReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c1f1b09c6e3ce14, []int{1}
}

func (m *LockReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockReply.Unmarshal(m, b)
}
func (m *LockReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockReply.Marshal(b, m, deterministic)
}
func (m *LockReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockReply.Merge(m, src)
}
func (m *LockReply) XXX_Size() int {
	return xxx_messageInfo_LockReply.Size(m)
}
func (m *LockReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LockReply.DiscardUnknown(m)
}

var xxx_messageInfo_LockReply proto.InternalMessageInfo

func (m *LockReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*LockRequest)(nil), "proto.LockRequest")
	proto.RegisterType((*LockReply)(nil), "proto.LockReply")
}

func init() { proto.RegisterFile("global_lock_request.proto", fileDescriptor_2c1f1b09c6e3ce14) }

var fileDescriptor_2c1f1b09c6e3ce14 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xcf, 0xc9, 0x4f,
	0x4a, 0xcc, 0x89, 0xcf, 0xc9, 0x4f, 0xce, 0x8e, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x6a, 0x5c, 0xdc, 0x3e, 0xf9, 0xc9,
	0xd9, 0x41, 0x10, 0x39, 0x21, 0x71, 0x2e, 0x76, 0xb0, 0xda, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x36, 0x10, 0xd7, 0x33, 0x45, 0x49, 0x95, 0x8b, 0x13, 0xa2, 0xae, 0x20, 0xa7,
	0x52, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x15, 0xaa, 0x0a, 0xc6, 0x35,
	0xaa, 0x86, 0x18, 0x17, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xca, 0xc5, 0xed, 0x98,
	0x5c, 0x58, 0x9a, 0x59, 0x94, 0x0a, 0x12, 0x15, 0x12, 0x82, 0xd8, 0xad, 0x87, 0x64, 0xa3, 0x94,
	0x00, 0x8a, 0x58, 0x41, 0x4e, 0xa5, 0x12, 0x03, 0x48, 0x5b, 0x50, 0x6a, 0x4e, 0x6a, 0x62, 0x31,
	0x49, 0xda, 0x9c, 0x0c, 0xb8, 0xa4, 0x33, 0xf3, 0xf5, 0xd2, 0x8b, 0x0a, 0x92, 0xf5, 0x52, 0x2b,
	0x12, 0x73, 0x0b, 0x72, 0x52, 0x8b, 0xf5, 0x32, 0x52, 0x73, 0x72, 0xf2, 0xcb, 0xf3, 0x8b, 0x72,
	0x52, 0x9c, 0xf8, 0x3d, 0x40, 0xec, 0x70, 0x10, 0x3b, 0x00, 0xa4, 0x37, 0x80, 0x31, 0x89, 0x0d,
	0x6c, 0x88, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x34, 0x0d, 0xee, 0x28, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LockServiceClient is the client API for LockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LockServiceClient interface {
	// Sends a greeting
	AcquireLock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockReply, error)
	ReleaseLock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockReply, error)
}

type lockServiceClient struct {
	cc *grpc.ClientConn
}

func NewLockServiceClient(cc *grpc.ClientConn) LockServiceClient {
	return &lockServiceClient{cc}
}

func (c *lockServiceClient) AcquireLock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockReply, error) {
	out := new(LockReply)
	err := c.cc.Invoke(ctx, "/proto.LockService/AcquireLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockServiceClient) ReleaseLock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockReply, error) {
	out := new(LockReply)
	err := c.cc.Invoke(ctx, "/proto.LockService/ReleaseLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LockServiceServer is the server API for LockService service.
type LockServiceServer interface {
	// Sends a greeting
	AcquireLock(context.Context, *LockRequest) (*LockReply, error)
	ReleaseLock(context.Context, *LockRequest) (*LockReply, error)
}

func RegisterLockServiceServer(s *grpc.Server, srv LockServiceServer) {
	s.RegisterService(&_LockService_serviceDesc, srv)
}

func _LockService_AcquireLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServiceServer).AcquireLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LockService/AcquireLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServiceServer).AcquireLock(ctx, req.(*LockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LockService_ReleaseLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServiceServer).ReleaseLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LockService/ReleaseLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServiceServer).ReleaseLock(ctx, req.(*LockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LockService",
	HandlerType: (*LockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AcquireLock",
			Handler:    _LockService_AcquireLock_Handler,
		},
		{
			MethodName: "ReleaseLock",
			Handler:    _LockService_ReleaseLock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "global_lock_request.proto",
}
