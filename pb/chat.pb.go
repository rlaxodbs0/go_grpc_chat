// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ResponseType int32

const (
	ResponseType_SUCCESS ResponseType = 0
	ResponseType_FAIL    ResponseType = 1
)

var ResponseType_name = map[int32]string{
	0: "SUCCESS",
	1: "FAIL",
}

var ResponseType_value = map[string]int32{
	"SUCCESS": 0,
	"FAIL":    1,
}

func (x ResponseType) String() string {
	return proto.EnumName(ResponseType_name, int32(x))
}

func (ResponseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

type Message struct {
	Sender               string               `protobuf:"bytes,1,opt,name=Sender,proto3" json:"Sender,omitempty"`
	Receiver             string               `protobuf:"bytes,2,opt,name=Receiver,proto3" json:"Receiver,omitempty"`
	Text                 string               `protobuf:"bytes,3,opt,name=Text,proto3" json:"Text,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
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

func (m *Message) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Message) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type UserInfo struct {
	UserName             string   `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Response             ResponseType `protobuf:"varint,1,opt,name=Response,proto3,enum=pb.ResponseType" json:"Response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetResponse() ResponseType {
	if m != nil {
		return m.Response
	}
	return ResponseType_SUCCESS
}

type LogoutResponse struct {
	Response             ResponseType `protobuf:"varint,1,opt,name=Response,proto3,enum=pb.ResponseType" json:"Response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{3}
}

func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (m *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(m, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

func (m *LogoutResponse) GetResponse() ResponseType {
	if m != nil {
		return m.Response
	}
	return ResponseType_SUCCESS
}

type UserList struct {
	UserNameActiveMap    map[string]string `protobuf:"bytes,1,rep,name=UserNameActiveMap,proto3" json:"UserNameActiveMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{4}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetUserNameActiveMap() map[string]string {
	if m != nil {
		return m.UserNameActiveMap
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.ResponseType", ResponseType_name, ResponseType_value)
	proto.RegisterType((*Message)(nil), "pb.Message")
	proto.RegisterType((*UserInfo)(nil), "pb.UserInfo")
	proto.RegisterType((*LoginResponse)(nil), "pb.LoginResponse")
	proto.RegisterType((*LogoutResponse)(nil), "pb.LogoutResponse")
	proto.RegisterType((*UserList)(nil), "pb.UserList")
	proto.RegisterMapType((map[string]string)(nil), "pb.UserList.UserNameActiveMapEntry")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xe3, 0x76, 0xeb, 0xda, 0xd3, 0x32, 0xba, 0x23, 0x34, 0x45, 0xb9, 0xa1, 0x32, 0x02,
	0x55, 0xd3, 0x94, 0x41, 0xb9, 0x99, 0x90, 0x40, 0x2a, 0x65, 0x48, 0x93, 0x3a, 0x04, 0x69, 0x77,
	0x8d, 0xdc, 0xf6, 0x34, 0x8b, 0xb6, 0xc6, 0x91, 0xed, 0x16, 0xfa, 0x0c, 0x3c, 0x06, 0x2f, 0xc2,
	0xa3, 0x21, 0x27, 0x71, 0x09, 0x1b, 0x57, 0xdc, 0x9d, 0xef, 0xf8, 0xb3, 0xfd, 0x3b, 0x7f, 0x00,
	0xe6, 0x37, 0xc2, 0x84, 0x99, 0x92, 0x46, 0x62, 0x2d, 0x9b, 0x05, 0x4f, 0x63, 0x29, 0xe3, 0x3b,
	0x3a, 0xcb, 0x33, 0xb3, 0xf5, 0xf2, 0xcc, 0x24, 0x2b, 0xd2, 0x46, 0xac, 0xb2, 0xc2, 0xc4, 0x7f,
	0x30, 0x38, 0xb8, 0x22, 0xad, 0x45, 0x4c, 0x78, 0x0c, 0x8d, 0x09, 0xa5, 0x0b, 0x52, 0x3e, 0xeb,
	0xb1, 0x7e, 0x2b, 0x2a, 0x15, 0x06, 0xd0, 0x8c, 0x68, 0x4e, 0xc9, 0x86, 0x94, 0x5f, 0xcb, 0x4f,
	0x76, 0x1a, 0x11, 0xf6, 0xa6, 0xf4, 0xdd, 0xf8, 0xf5, 0x3c, 0x9f, 0xc7, 0x78, 0x0e, 0xad, 0x91,
	0x22, 0x61, 0x68, 0x31, 0x34, 0x7e, 0xab, 0xc7, 0xfa, 0xed, 0x41, 0x10, 0x16, 0x20, 0xa1, 0x03,
	0x09, 0xa7, 0x0e, 0x24, 0xfa, 0x63, 0xe6, 0xef, 0xa1, 0x79, 0xad, 0x49, 0x5d, 0xa6, 0x4b, 0x69,
	0x7f, 0xb5, 0xf1, 0x27, 0xb1, 0xa2, 0x92, 0x67, 0xa7, 0xed, 0xd9, 0x67, 0xa1, 0xf5, 0x37, 0xa9,
	0x16, 0x8e, 0xc8, 0x69, 0xfe, 0x16, 0x1e, 0x8d, 0x65, 0x9c, 0xa4, 0x11, 0xe9, 0x4c, 0xa6, 0x9a,
	0xf0, 0xd4, 0xe2, 0x17, 0x71, 0xfe, 0xd0, 0xe1, 0xa0, 0x1b, 0x66, 0xb3, 0xd0, 0xe5, 0xa6, 0xdb,
	0x8c, 0xa2, 0x9d, 0x83, 0xbf, 0x83, 0xc3, 0xb1, 0x8c, 0xe5, 0xda, 0xfc, 0xe7, 0xfd, 0x9f, 0xac,
	0xe0, 0x1e, 0x27, 0xda, 0xe0, 0x17, 0x38, 0x72, 0xcc, 0xc3, 0xb9, 0x49, 0x36, 0x74, 0x25, 0x32,
	0x9f, 0xf5, 0xea, 0xfd, 0xf6, 0xe0, 0x99, 0x7d, 0xc3, 0x19, 0xc3, 0x07, 0xae, 0x8b, 0xd4, 0xa8,
	0x6d, 0xf4, 0xf0, 0x76, 0xf0, 0x01, 0x8e, 0xff, 0x6d, 0xc6, 0x2e, 0xd4, 0x6f, 0x69, 0x5b, 0xf6,
	0xca, 0x86, 0xf8, 0x04, 0xf6, 0x37, 0xe2, 0x6e, 0x4d, 0x65, 0x8f, 0x0a, 0xf1, 0xa6, 0x76, 0xce,
	0x4e, 0x9e, 0x43, 0xa7, 0xca, 0x8f, 0x6d, 0x38, 0x98, 0x5c, 0x8f, 0x46, 0x17, 0x93, 0x49, 0xd7,
	0xc3, 0x26, 0xec, 0x7d, 0x1c, 0x5e, 0x8e, 0xbb, 0x6c, 0xf0, 0x8b, 0x41, 0xcb, 0x6e, 0xd4, 0x57,
	0x23, 0xf4, 0x2d, 0xbe, 0x82, 0xc7, 0x53, 0x25, 0x52, 0xbd, 0x24, 0xe5, 0x56, 0xa6, 0x6d, 0xab,
	0x28, 0x45, 0x50, 0x15, 0xdc, 0xeb, 0xb3, 0x97, 0x0c, 0x4f, 0x60, 0x3f, 0x1f, 0x06, 0x76, 0x5c,
	0xb9, 0x76, 0xb6, 0xc1, 0x91, 0x55, 0x7f, 0x4d, 0x89, 0x7b, 0x78, 0x0a, 0x8d, 0xa2, 0xf3, 0xf7,
	0xcc, 0x58, 0x9a, 0x2b, 0x33, 0xe1, 0x1e, 0xbe, 0xb0, 0xcb, 0x2a, 0xd4, 0xfc, 0xe6, 0x9e, 0xbb,
	0x53, 0xed, 0x2b, 0xf7, 0x66, 0x8d, 0x7c, 0xe3, 0x5e, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xe0,
	0x50, 0x64, 0xa7, 0x1a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatTaskClient is the client API for ChatTask service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatTaskClient interface {
	TransferMessage(ctx context.Context, opts ...grpc.CallOption) (ChatTask_TransferMessageClient, error)
	Login(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*LogoutResponse, error)
	Search(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserList, error)
}

type chatTaskClient struct {
	cc *grpc.ClientConn
}

func NewChatTaskClient(cc *grpc.ClientConn) ChatTaskClient {
	return &chatTaskClient{cc}
}

func (c *chatTaskClient) TransferMessage(ctx context.Context, opts ...grpc.CallOption) (ChatTask_TransferMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatTask_serviceDesc.Streams[0], "/pb.chat_task/TransferMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatTaskTransferMessageClient{stream}
	return x, nil
}

type ChatTask_TransferMessageClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatTaskTransferMessageClient struct {
	grpc.ClientStream
}

func (x *chatTaskTransferMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatTaskTransferMessageClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatTaskClient) Login(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/pb.chat_task/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatTaskClient) Logout(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/pb.chat_task/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatTaskClient) Search(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/pb.chat_task/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatTaskServer is the server API for ChatTask service.
type ChatTaskServer interface {
	TransferMessage(ChatTask_TransferMessageServer) error
	Login(context.Context, *UserInfo) (*LoginResponse, error)
	Logout(context.Context, *UserInfo) (*LogoutResponse, error)
	Search(context.Context, *UserInfo) (*UserList, error)
}

// UnimplementedChatTaskServer can be embedded to have forward compatible implementations.
type UnimplementedChatTaskServer struct {
}

func (*UnimplementedChatTaskServer) TransferMessage(srv ChatTask_TransferMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method TransferMessage not implemented")
}
func (*UnimplementedChatTaskServer) Login(ctx context.Context, req *UserInfo) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedChatTaskServer) Logout(ctx context.Context, req *UserInfo) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (*UnimplementedChatTaskServer) Search(ctx context.Context, req *UserInfo) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterChatTaskServer(s *grpc.Server, srv ChatTaskServer) {
	s.RegisterService(&_ChatTask_serviceDesc, srv)
}

func _ChatTask_TransferMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatTaskServer).TransferMessage(&chatTaskTransferMessageServer{stream})
}

type ChatTask_TransferMessageServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatTaskTransferMessageServer struct {
	grpc.ServerStream
}

func (x *chatTaskTransferMessageServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatTaskTransferMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatTask_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatTaskServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.chat_task/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatTaskServer).Login(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatTask_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatTaskServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.chat_task/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatTaskServer).Logout(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatTask_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatTaskServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.chat_task/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatTaskServer).Search(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatTask_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.chat_task",
	HandlerType: (*ChatTaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _ChatTask_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _ChatTask_Logout_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _ChatTask_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransferMessage",
			Handler:       _ChatTask_TransferMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}
