// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Msg
*/
package message

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

type Msg struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Msg) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Msg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Msg)(nil), "message.Msg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MessageService service

type MessageServiceClient interface {
	Message(ctx context.Context, opts ...grpc.CallOption) (MessageService_MessageClient, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) Message(ctx context.Context, opts ...grpc.CallOption) (MessageService_MessageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MessageService_serviceDesc.Streams[0], c.cc, "/message.MessageService/Message", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceMessageClient{stream}
	return x, nil
}

type MessageService_MessageClient interface {
	Send(*Msg) error
	Recv() (*Msg, error)
	grpc.ClientStream
}

type messageServiceMessageClient struct {
	grpc.ClientStream
}

func (x *messageServiceMessageClient) Send(m *Msg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageServiceMessageClient) Recv() (*Msg, error) {
	m := new(Msg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MessageService service

type MessageServiceServer interface {
	Message(MessageService_MessageServer) error
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_Message_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageServiceServer).Message(&messageServiceMessageServer{stream})
}

type MessageService_MessageServer interface {
	Send(*Msg) error
	Recv() (*Msg, error)
	grpc.ServerStream
}

type messageServiceMessageServer struct {
	grpc.ServerStream
}

func (x *messageServiceMessageServer) Send(m *Msg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageServiceMessageServer) Recv() (*Msg, error) {
	m := new(Msg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Message",
			Handler:       _MessageService_Message_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "message.proto",
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x74, 0xb9,
	0x98, 0x7d, 0x8b, 0xd3, 0x85, 0x84, 0xb8, 0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x4a, 0x62, 0x49, 0xa2, 0x04, 0x93, 0x02, 0xa3, 0x06,
	0x4f, 0x10, 0x98, 0x6d, 0x64, 0xcb, 0xc5, 0xe7, 0x0b, 0xd1, 0x19, 0x9c, 0x5a, 0x54, 0x96, 0x99,
	0x9c, 0x2a, 0xa4, 0xcd, 0xc5, 0x0e, 0x15, 0x11, 0xe2, 0xd1, 0x83, 0x59, 0xe2, 0x5b, 0x9c, 0x2e,
	0x85, 0xc2, 0x53, 0x62, 0xd0, 0x60, 0x34, 0x60, 0x4c, 0x62, 0x03, 0xdb, 0x6e, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x21, 0xf3, 0x1e, 0x6d, 0x8e, 0x00, 0x00, 0x00,
}