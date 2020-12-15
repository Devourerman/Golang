// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.proto

package chatpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Channel struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SendersName          string   `protobuf:"bytes,2,opt,name=senders_name,json=sendersName,proto3" json:"senders_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Channel) Reset()         { *m = Channel{} }
func (m *Channel) String() string { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()    {}
func (*Channel) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}
func (m *Channel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Channel.Unmarshal(m, b)
}
func (m *Channel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Channel.Marshal(b, m, deterministic)
}
func (m *Channel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Channel.Merge(m, src)
}
func (m *Channel) XXX_Size() int {
	return xxx_messageInfo_Channel.Size(m)
}
func (m *Channel) XXX_DiscardUnknown() {
	xxx_messageInfo_Channel.DiscardUnknown(m)
}

var xxx_messageInfo_Channel proto.InternalMessageInfo

func (m *Channel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Channel) GetSendersName() string {
	if m != nil {
		return m.SendersName
	}
	return ""
}

type Message struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Channel              *Channel `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Message) GetChannel() *Channel {
	if m != nil {
		return m.Channel
	}
	return nil
}

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type MessageAck struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageAck) Reset()         { *m = MessageAck{} }
func (m *MessageAck) String() string { return proto.CompactTextString(m) }
func (*MessageAck) ProtoMessage()    {}
func (*MessageAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
}
func (m *MessageAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageAck.Unmarshal(m, b)
}
func (m *MessageAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageAck.Marshal(b, m, deterministic)
}
func (m *MessageAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageAck.Merge(m, src)
}
func (m *MessageAck) XXX_Size() int {
	return xxx_messageInfo_MessageAck.Size(m)
}
func (m *MessageAck) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageAck.DiscardUnknown(m)
}

var xxx_messageInfo_MessageAck proto.InternalMessageInfo

func (m *MessageAck) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Channel)(nil), "chatpb.Channel")
	proto.RegisterType((*Message)(nil), "chatpb.Message")
	proto.RegisterType((*MessageAck)(nil), "chatpb.MessageAck")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x6b, 0x40, 0x31, 0x9c, 0x91, 0x90, 0x6e, 0x40, 0x51, 0x27, 0xb0, 0x18, 0xca, 0xe2,
	0xa2, 0x22, 0x31, 0x53, 0x32, 0xc3, 0xd0, 0x6e, 0x2c, 0xc8, 0x75, 0x0f, 0x82, 0x4a, 0x9c, 0x2a,
	0x36, 0xf9, 0xfd, 0x88, 0xb3, 0x93, 0x21, 0xdb, 0xbd, 0xf7, 0xa2, 0xef, 0x53, 0x0c, 0xe0, 0x6a,
	0x1b, 0xcd, 0xb1, 0x6b, 0x63, 0x8b, 0xc5, 0xff, 0x7d, 0xdc, 0xe9, 0x67, 0x90, 0x55, 0x6d, 0xbd,
	0xa7, 0x1f, 0x44, 0x38, 0xf3, 0xb6, 0xa1, 0x52, 0xdc, 0x88, 0xc5, 0xc5, 0x86, 0x6f, 0xbc, 0x85,
	0xcb, 0x40, 0x7e, 0x4f, 0x5d, 0xf8, 0xe0, 0xed, 0x84, 0x37, 0x95, 0xbb, 0x37, 0xdb, 0x90, 0xfe,
	0x04, 0xf9, 0x4a, 0x21, 0xd8, 0x2f, 0xc2, 0x6b, 0x28, 0xd2, 0x92, 0x19, 0x39, 0xe1, 0x3d, 0x48,
	0x97, 0x24, 0x0c, 0x50, 0xab, 0x2b, 0x93, 0xf4, 0x26, 0xbb, 0x37, 0xc3, 0x8e, 0x25, 0xc8, 0x26,
	0xd1, 0xca, 0x53, 0x66, 0x0c, 0x51, 0xdf, 0x01, 0x64, 0xcf, 0xda, 0x1d, 0x58, 0x15, 0x6d, 0xfc,
	0x0d, 0xa3, 0x8a, 0xd3, 0xaa, 0x07, 0x55, 0xd5, 0x36, 0x6e, 0xa9, 0xeb, 0xbf, 0x1d, 0xe1, 0x12,
	0x64, 0xd5, 0x7a, 0x4f, 0x2e, 0xe2, 0xd4, 0x39, 0x1f, 0x8b, 0x8c, 0xd5, 0xb3, 0x07, 0x81, 0x4f,
	0xa0, 0xb6, 0xe4, 0xf7, 0xc3, 0x1f, 0x4d, 0xbf, 0x99, 0xe3, 0xa4, 0x58, 0xbb, 0x83, 0x9e, 0x2d,
	0xc4, 0x0b, 0xbc, 0x9f, 0x9b, 0x65, 0x9a, 0x76, 0x05, 0x3f, 0xf1, 0xe3, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x86, 0xcc, 0x40, 0x31, 0x70, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	Connect(ctx context.Context, in *Channel, opts ...grpc.CallOption) (ChatService_ConnectClient, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMessageClient, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Connect(ctx context.Context, in *Channel, opts ...grpc.CallOption) (ChatService_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/chatpb.ChatService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_ConnectClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatServiceConnectClient struct {
	grpc.ClientStream
}

func (x *chatServiceConnectClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[1], "/chatpb.ChatService/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSendMessageClient{stream}
	return x, nil
}

type ChatService_SendMessageClient interface {
	Send(*Message) error
	CloseAndRecv() (*MessageAck, error)
	grpc.ClientStream
}

type chatServiceSendMessageClient struct {
	grpc.ClientStream
}

func (x *chatServiceSendMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceSendMessageClient) CloseAndRecv() (*MessageAck, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessageAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	Connect(*Channel, ChatService_ConnectServer) error
	SendMessage(ChatService_SendMessageServer) error
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) Connect(req *Channel, srv ChatService_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedChatServiceServer) SendMessage(srv ChatService_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Channel)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).Connect(m, &chatServiceConnectServer{stream})
}

type ChatService_ConnectServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chatServiceConnectServer struct {
	grpc.ServerStream
}

func (x *chatServiceConnectServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).SendMessage(&chatServiceSendMessageServer{stream})
}

type ChatService_SendMessageServer interface {
	SendAndClose(*MessageAck) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatServiceSendMessageServer struct {
	grpc.ServerStream
}

func (x *chatServiceSendMessageServer) SendAndClose(m *MessageAck) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceSendMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chatpb.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _ChatService_Connect_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendMessage",
			Handler:       _ChatService_SendMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}
