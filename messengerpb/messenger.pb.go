// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messengerpb/messenger.proto

package messengerpb

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

type Messenger struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Date                 string   `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Messenger) Reset()         { *m = Messenger{} }
func (m *Messenger) String() string { return proto.CompactTextString(m) }
func (*Messenger) ProtoMessage()    {}
func (*Messenger) Descriptor() ([]byte, []int) {
	return fileDescriptor_057616305f412db8, []int{0}
}

func (m *Messenger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Messenger.Unmarshal(m, b)
}
func (m *Messenger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Messenger.Marshal(b, m, deterministic)
}
func (m *Messenger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Messenger.Merge(m, src)
}
func (m *Messenger) XXX_Size() int {
	return xxx_messageInfo_Messenger.Size(m)
}
func (m *Messenger) XXX_DiscardUnknown() {
	xxx_messageInfo_Messenger.DiscardUnknown(m)
}

var xxx_messageInfo_Messenger proto.InternalMessageInfo

func (m *Messenger) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Messenger) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Messenger) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type CreateMessageRequest struct {
	Message              *Messenger `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateMessageRequest) Reset()         { *m = CreateMessageRequest{} }
func (m *CreateMessageRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMessageRequest) ProtoMessage()    {}
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_057616305f412db8, []int{1}
}

func (m *CreateMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessageRequest.Unmarshal(m, b)
}
func (m *CreateMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessageRequest.Marshal(b, m, deterministic)
}
func (m *CreateMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessageRequest.Merge(m, src)
}
func (m *CreateMessageRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMessageRequest.Size(m)
}
func (m *CreateMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessageRequest proto.InternalMessageInfo

func (m *CreateMessageRequest) GetMessage() *Messenger {
	if m != nil {
		return m.Message
	}
	return nil
}

type CreateMessageResponse struct {
	Message              *Messenger `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateMessageResponse) Reset()         { *m = CreateMessageResponse{} }
func (m *CreateMessageResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMessageResponse) ProtoMessage()    {}
func (*CreateMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_057616305f412db8, []int{2}
}

func (m *CreateMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessageResponse.Unmarshal(m, b)
}
func (m *CreateMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessageResponse.Marshal(b, m, deterministic)
}
func (m *CreateMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessageResponse.Merge(m, src)
}
func (m *CreateMessageResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMessageResponse.Size(m)
}
func (m *CreateMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessageResponse proto.InternalMessageInfo

func (m *CreateMessageResponse) GetMessage() *Messenger {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*Messenger)(nil), "messenger.Messenger")
	proto.RegisterType((*CreateMessageRequest)(nil), "messenger.CreateMessageRequest")
	proto.RegisterType((*CreateMessageResponse)(nil), "messenger.CreateMessageResponse")
}

func init() { proto.RegisterFile("messengerpb/messenger.proto", fileDescriptor_057616305f412db8) }

var fileDescriptor_057616305f412db8 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x4d, 0x2d, 0x2e,
	0x4e, 0xcd, 0x4b, 0x4f, 0x2d, 0x2a, 0x48, 0xd2, 0x87, 0xb3, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0x38, 0xe1, 0x02, 0x4a, 0xa1, 0x5c, 0x9c, 0xbe, 0x30, 0x8e, 0x90, 0x14, 0x17, 0x47, 0x69,
	0x71, 0x6a, 0x91, 0x5f, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0x2f,
	0x24, 0xc1, 0xc5, 0x0e, 0xd2, 0x95, 0x98, 0x9e, 0x2a, 0xc1, 0x04, 0x96, 0x82, 0x71, 0x85, 0x84,
	0xb8, 0x58, 0x52, 0x12, 0x4b, 0x52, 0x25, 0x98, 0xc1, 0xc2, 0x60, 0xb6, 0x92, 0x1b, 0x97, 0x88,
	0x73, 0x51, 0x6a, 0x62, 0x49, 0xaa, 0x2f, 0x44, 0x51, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89,
	0x90, 0x1e, 0xc2, 0x14, 0x90, 0x05, 0xdc, 0x46, 0x22, 0x7a, 0x08, 0xc7, 0xc1, 0x1d, 0x02, 0x37,
	0x5b, 0xc9, 0x9d, 0x4b, 0x14, 0xcd, 0x9c, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x52, 0x0d, 0x32,
	0x4a, 0xe3, 0xe2, 0x83, 0x1a, 0x11, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x14, 0xc2, 0xc5,
	0x8b, 0x62, 0xb4, 0x90, 0x3c, 0x92, 0x09, 0xd8, 0x1c, 0x2f, 0xa5, 0x80, 0x5b, 0x01, 0xc4, 0x55,
	0x4a, 0x0c, 0x4e, 0xbc, 0x51, 0xdc, 0x48, 0x21, 0x9f, 0xc4, 0x06, 0x0e, 0x70, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x8d, 0x82, 0xce, 0xbe, 0x8f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageServiceClient interface {
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error) {
	out := new(CreateMessageResponse)
	err := c.cc.Invoke(ctx, "/messenger.MessageService/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
type MessageServiceServer interface {
	CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error)
}

// UnimplementedMessageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (*UnimplementedMessageServiceServer) CreateMessage(ctx context.Context, req *CreateMessageRequest) (*CreateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messenger.MessageService/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messenger.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _MessageService_CreateMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messengerpb/messenger.proto",
}
