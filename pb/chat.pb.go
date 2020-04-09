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
	ResponseType_SUCCESS       ResponseType = 0
	ResponseType_FAIL          ResponseType = 1
	ResponseType_ALREADYEXISTS ResponseType = 2
	ResponseType_NOMATCH       ResponseType = 3
)

var ResponseType_name = map[int32]string{
	0: "SUCCESS",
	1: "FAIL",
	2: "ALREADYEXISTS",
	3: "NOMATCH",
}

var ResponseType_value = map[string]int32{
	"SUCCESS":       0,
	"FAIL":          1,
	"ALREADYEXISTS": 2,
	"NOMATCH":       3,
}

func (x ResponseType) String() string {
	return proto.EnumName(ResponseType_name, int32(x))
}

func (ResponseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

type UserInfo struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
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

func (m *UserInfo) GetUsername() string {
	if m != nil {
		return m.Username
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
	Response             ResponseType `protobuf:"varint,1,opt,name=response,proto3,enum=pb.ResponseType" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
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
	Response             ResponseType `protobuf:"varint,1,opt,name=response,proto3,enum=pb.ResponseType" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
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

type SignupResponse struct {
	Response             ResponseType `protobuf:"varint,1,opt,name=response,proto3,enum=pb.ResponseType" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SignupResponse) Reset()         { *m = SignupResponse{} }
func (m *SignupResponse) String() string { return proto.CompactTextString(m) }
func (*SignupResponse) ProtoMessage()    {}
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{3}
}

func (m *SignupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignupResponse.Unmarshal(m, b)
}
func (m *SignupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignupResponse.Marshal(b, m, deterministic)
}
func (m *SignupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupResponse.Merge(m, src)
}
func (m *SignupResponse) XXX_Size() int {
	return xxx_messageInfo_SignupResponse.Size(m)
}
func (m *SignupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignupResponse proto.InternalMessageInfo

func (m *SignupResponse) GetResponse() ResponseType {
	if m != nil {
		return m.Response
	}
	return ResponseType_SUCCESS
}

type Message struct {
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Username  string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*Message_ChatMessage
	//	*Message_InviteMessage_
	//	*Message_BroadcastMessage_
	Event                isMessage_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{4}
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

func (m *Message) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Message) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type isMessage_Event interface {
	isMessage_Event()
}

type Message_ChatMessage struct {
	ChatMessage *Message_Message `protobuf:"bytes,3,opt,name=chat_message,json=chatMessage,proto3,oneof"`
}

type Message_InviteMessage_ struct {
	InviteMessage *Message_InviteMessage `protobuf:"bytes,4,opt,name=invite_message,json=inviteMessage,proto3,oneof"`
}

type Message_BroadcastMessage_ struct {
	BroadcastMessage *Message_BroadcastMessage `protobuf:"bytes,5,opt,name=broadcast_message,json=broadcastMessage,proto3,oneof"`
}

func (*Message_ChatMessage) isMessage_Event() {}

func (*Message_InviteMessage_) isMessage_Event() {}

func (*Message_BroadcastMessage_) isMessage_Event() {}

func (m *Message) GetEvent() isMessage_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Message) GetChatMessage() *Message_Message {
	if x, ok := m.GetEvent().(*Message_ChatMessage); ok {
		return x.ChatMessage
	}
	return nil
}

func (m *Message) GetInviteMessage() *Message_InviteMessage {
	if x, ok := m.GetEvent().(*Message_InviteMessage_); ok {
		return x.InviteMessage
	}
	return nil
}

func (m *Message) GetBroadcastMessage() *Message_BroadcastMessage {
	if x, ok := m.GetEvent().(*Message_BroadcastMessage_); ok {
		return x.BroadcastMessage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_ChatMessage)(nil),
		(*Message_InviteMessage_)(nil),
		(*Message_BroadcastMessage_)(nil),
	}
}

type Message_Message struct {
	ChatGroup            uint64   `protobuf:"varint,2,opt,name=chat_group,json=chatGroup,proto3" json:"chat_group,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_Message) Reset()         { *m = Message_Message{} }
func (m *Message_Message) String() string { return proto.CompactTextString(m) }
func (*Message_Message) ProtoMessage()    {}
func (*Message_Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{4, 0}
}

func (m *Message_Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_Message.Unmarshal(m, b)
}
func (m *Message_Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_Message.Marshal(b, m, deterministic)
}
func (m *Message_Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Message.Merge(m, src)
}
func (m *Message_Message) XXX_Size() int {
	return xxx_messageInfo_Message_Message.Size(m)
}
func (m *Message_Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Message proto.InternalMessageInfo

func (m *Message_Message) GetChatGroup() uint64 {
	if m != nil {
		return m.ChatGroup
	}
	return 0
}

func (m *Message_Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Message_InviteMessage struct {
	TargetUsername       string   `protobuf:"bytes,2,opt,name=target_username,json=targetUsername,proto3" json:"target_username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_InviteMessage) Reset()         { *m = Message_InviteMessage{} }
func (m *Message_InviteMessage) String() string { return proto.CompactTextString(m) }
func (*Message_InviteMessage) ProtoMessage()    {}
func (*Message_InviteMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{4, 1}
}

func (m *Message_InviteMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_InviteMessage.Unmarshal(m, b)
}
func (m *Message_InviteMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_InviteMessage.Marshal(b, m, deterministic)
}
func (m *Message_InviteMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_InviteMessage.Merge(m, src)
}
func (m *Message_InviteMessage) XXX_Size() int {
	return xxx_messageInfo_Message_InviteMessage.Size(m)
}
func (m *Message_InviteMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_InviteMessage.DiscardUnknown(m)
}

var xxx_messageInfo_Message_InviteMessage proto.InternalMessageInfo

func (m *Message_InviteMessage) GetTargetUsername() string {
	if m != nil {
		return m.TargetUsername
	}
	return ""
}

type Message_BroadcastMessage struct {
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_BroadcastMessage) Reset()         { *m = Message_BroadcastMessage{} }
func (m *Message_BroadcastMessage) String() string { return proto.CompactTextString(m) }
func (*Message_BroadcastMessage) ProtoMessage()    {}
func (*Message_BroadcastMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{4, 2}
}

func (m *Message_BroadcastMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_BroadcastMessage.Unmarshal(m, b)
}
func (m *Message_BroadcastMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_BroadcastMessage.Marshal(b, m, deterministic)
}
func (m *Message_BroadcastMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_BroadcastMessage.Merge(m, src)
}
func (m *Message_BroadcastMessage) XXX_Size() int {
	return xxx_messageInfo_Message_BroadcastMessage.Size(m)
}
func (m *Message_BroadcastMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_BroadcastMessage.DiscardUnknown(m)
}

var xxx_messageInfo_Message_BroadcastMessage proto.InternalMessageInfo

func (m *Message_BroadcastMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
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
	return fileDescriptor_8c585a45e2093e54, []int{5}
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
	proto.RegisterType((*UserInfo)(nil), "pb.UserInfo")
	proto.RegisterType((*LoginResponse)(nil), "pb.LoginResponse")
	proto.RegisterType((*LogoutResponse)(nil), "pb.LogoutResponse")
	proto.RegisterType((*SignupResponse)(nil), "pb.SignupResponse")
	proto.RegisterType((*Message)(nil), "pb.Message")
	proto.RegisterType((*Message_Message)(nil), "pb.Message.Message")
	proto.RegisterType((*Message_InviteMessage)(nil), "pb.Message.InviteMessage")
	proto.RegisterType((*Message_BroadcastMessage)(nil), "pb.Message.BroadcastMessage")
	proto.RegisterType((*UserList)(nil), "pb.UserList")
	proto.RegisterMapType((map[string]string)(nil), "pb.UserList.UserNameActiveMapEntry")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdd, 0x6e, 0xda, 0x4c,
	0x10, 0xb5, 0x0d, 0x04, 0x18, 0x7e, 0x3e, 0xb3, 0x5f, 0x55, 0x51, 0xab, 0x55, 0x2b, 0x57, 0x6a,
	0xa3, 0x28, 0x72, 0x2a, 0x7a, 0x83, 0x2a, 0xb5, 0x12, 0x10, 0xda, 0xa0, 0x92, 0x54, 0xb5, 0x41,
	0x6a, 0xaf, 0xd0, 0x9a, 0x6c, 0x8c, 0x95, 0xe0, 0xb5, 0xbc, 0x6b, 0x2a, 0x9e, 0xa7, 0x0f, 0xd6,
	0x8b, 0xbe, 0x48, 0xb5, 0x6b, 0x0c, 0xb6, 0x93, 0xab, 0x5c, 0xd9, 0x33, 0x7b, 0xe6, 0xcc, 0xec,
	0xd9, 0x39, 0x00, 0xcb, 0x15, 0xe6, 0x56, 0x18, 0x51, 0x4e, 0x91, 0x16, 0xba, 0xc6, 0x4b, 0x8f,
	0x52, 0xef, 0x8e, 0x9c, 0xc9, 0x8c, 0x1b, 0xdf, 0x9c, 0x71, 0x7f, 0x4d, 0x18, 0xc7, 0xeb, 0x30,
	0x01, 0x99, 0x43, 0xa8, 0xcd, 0x19, 0x89, 0x26, 0xc1, 0x0d, 0x45, 0x06, 0xd4, 0x62, 0x46, 0xa2,
	0x00, 0xaf, 0x49, 0x57, 0x7d, 0xa5, 0x1e, 0xd7, 0xed, 0x7d, 0x2c, 0xce, 0x42, 0xcc, 0xd8, 0x2f,
	0x1a, 0x5d, 0x77, 0xb5, 0xe4, 0x2c, 0x8d, 0xcd, 0x8f, 0xd0, 0x9a, 0x52, 0xcf, 0x0f, 0x6c, 0xc2,
	0x42, 0x1a, 0x30, 0x82, 0x4e, 0xa1, 0x16, 0xed, 0xfe, 0x25, 0x51, 0xbb, 0xa7, 0x5b, 0xa1, 0x6b,
	0xa5, 0xe7, 0xb3, 0x6d, 0x48, 0xec, 0x3d, 0xc2, 0xfc, 0x04, 0xed, 0x29, 0xf5, 0x68, 0xcc, 0x1f,
	0x5f, 0xef, 0xf8, 0x5e, 0x10, 0x87, 0x8f, 0xac, 0xff, 0x53, 0x82, 0xea, 0x25, 0x61, 0x0c, 0x7b,
	0x04, 0xf5, 0xa1, 0xbe, 0x57, 0x48, 0x96, 0x36, 0x7a, 0x86, 0x95, 0x68, 0x68, 0xa5, 0x1a, 0x5a,
	0xb3, 0x14, 0x61, 0x1f, 0xc0, 0x39, 0xf1, 0xb4, 0x82, 0x78, 0x7d, 0x68, 0x8a, 0x77, 0x59, 0xac,
	0x93, 0x2e, 0xdd, 0x92, 0x24, 0xfe, 0x5f, 0xcc, 0xb4, 0x6b, 0x9c, 0x7e, 0x2f, 0x14, 0xbb, 0x21,
	0xa0, 0xe9, 0x3c, 0x43, 0x68, 0xfb, 0xc1, 0xc6, 0xe7, 0x64, 0x5f, 0x5b, 0x96, 0xb5, 0xcf, 0xb2,
	0xb5, 0x13, 0x89, 0x38, 0x30, 0xb4, 0xfc, 0x6c, 0x02, 0x7d, 0x85, 0x8e, 0x1b, 0x51, 0x7c, 0xbd,
	0xc4, 0xec, 0x30, 0x42, 0x45, 0xd2, 0x3c, 0xcf, 0xd2, 0x0c, 0x53, 0xd0, 0x81, 0x49, 0x77, 0x0b,
	0x39, 0x63, 0x78, 0xd0, 0xea, 0x45, 0xb2, 0x6d, 0x0b, 0x2f, 0xa2, 0x71, 0x28, 0xef, 0x5c, 0xb6,
	0xeb, 0x22, 0xf3, 0x45, 0x24, 0x50, 0x17, 0xaa, 0xd9, 0xfb, 0xd6, 0xed, 0x34, 0x34, 0xfa, 0xd0,
	0xca, 0x8d, 0x8c, 0xde, 0xc2, 0x7f, 0x1c, 0x47, 0x1e, 0xe1, 0x8b, 0x82, 0x84, 0xed, 0x24, 0x3d,
	0xdf, 0x65, 0x8d, 0x53, 0xd0, 0x8b, 0x53, 0x66, 0xfb, 0x68, 0xb9, 0x3e, 0xc3, 0x2a, 0x54, 0xc8,
	0x86, 0x04, 0xdc, 0xfc, 0xad, 0x26, 0x5b, 0x3e, 0xf5, 0x19, 0x47, 0xdf, 0xa1, 0x23, 0xfe, 0xaf,
	0xf0, 0x9a, 0x0c, 0x96, 0xdc, 0xdf, 0x90, 0x4b, 0x2c, 0x9e, 0xba, 0x74, 0xdc, 0xe8, 0xbd, 0x16,
	0x72, 0xa4, 0x40, 0xeb, 0x1e, 0x6a, 0x1c, 0xf0, 0x68, 0x6b, 0xdf, 0xaf, 0x36, 0xce, 0xe1, 0xe9,
	0xc3, 0x60, 0xa4, 0x43, 0xe9, 0x96, 0x6c, 0x77, 0x6e, 0x12, 0xbf, 0xe8, 0x09, 0x54, 0x36, 0xf8,
	0x2e, 0x4e, 0x87, 0x4d, 0x82, 0x0f, 0x5a, 0x5f, 0x3d, 0x19, 0x43, 0x33, 0xbb, 0xa1, 0xa8, 0x01,
	0x55, 0x67, 0x3e, 0x1a, 0x8d, 0x1d, 0x47, 0x57, 0x50, 0x0d, 0xca, 0x9f, 0x07, 0x93, 0xa9, 0xae,
	0xa2, 0x0e, 0xb4, 0x06, 0x53, 0x7b, 0x3c, 0x38, 0xff, 0x39, 0xfe, 0x31, 0x71, 0x66, 0x8e, 0xae,
	0x09, 0xe4, 0xd5, 0xb7, 0xcb, 0xc1, 0x6c, 0x74, 0xa1, 0x97, 0x7a, 0x7f, 0x55, 0x90, 0xaf, 0xb0,
	0xe0, 0x98, 0xdd, 0xa2, 0x53, 0x38, 0x12, 0xe6, 0x98, 0x87, 0xa8, 0x99, 0x5e, 0x4e, 0x78, 0xdd,
	0x40, 0x22, 0xca, 0xdb, 0xc6, 0x54, 0xd0, 0x09, 0x54, 0xa4, 0x93, 0x0b, 0xe0, 0x8e, 0x88, 0x72,
	0x16, 0x37, 0x15, 0xc1, 0x9c, 0xd8, 0xf6, 0x21, 0xe6, 0xbc, 0xa1, 0x4d, 0x05, 0xbd, 0x81, 0x23,
	0x87, 0xe0, 0x68, 0xb9, 0x2a, 0xa0, 0x9b, 0x59, 0xc9, 0x25, 0xae, 0x3c, 0x5a, 0x61, 0x8e, 0x1a,
	0x99, 0xcd, 0x34, 0xb2, 0x81, 0xa9, 0x1c, 0xab, 0xef, 0x54, 0xf7, 0x48, 0xba, 0xf1, 0xfd, 0xbf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xf8, 0x5b, 0x39, 0xf1, 0x04, 0x00, 0x00,
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
	SignUp(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*SignupResponse, error)
	Login(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*LogoutResponse, error)
	Search(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserList, error)
	Chat(ctx context.Context, opts ...grpc.CallOption) (ChatTask_ChatClient, error)
}

type chatTaskClient struct {
	cc *grpc.ClientConn
}

func NewChatTaskClient(cc *grpc.ClientConn) ChatTaskClient {
	return &chatTaskClient{cc}
}

func (c *chatTaskClient) SignUp(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, "/pb.chat_task/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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

func (c *chatTaskClient) Chat(ctx context.Context, opts ...grpc.CallOption) (ChatTask_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatTask_serviceDesc.Streams[0], "/pb.chat_task/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatTaskChatClient{stream}
	return x, nil
}

type ChatTask_ChatClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatTaskChatClient struct {
	grpc.ClientStream
}

func (x *chatTaskChatClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatTaskChatClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatTaskServer is the server API for ChatTask service.
type ChatTaskServer interface {
	SignUp(context.Context, *UserInfo) (*SignupResponse, error)
	Login(context.Context, *UserInfo) (*LoginResponse, error)
	Logout(context.Context, *UserInfo) (*LogoutResponse, error)
	Search(context.Context, *UserInfo) (*UserList, error)
	Chat(ChatTask_ChatServer) error
}

// UnimplementedChatTaskServer can be embedded to have forward compatible implementations.
type UnimplementedChatTaskServer struct {
}

func (*UnimplementedChatTaskServer) SignUp(ctx context.Context, req *UserInfo) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
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
func (*UnimplementedChatTaskServer) Chat(srv ChatTask_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}

func RegisterChatTaskServer(s *grpc.Server, srv ChatTaskServer) {
	s.RegisterService(&_ChatTask_serviceDesc, srv)
}

func _ChatTask_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatTaskServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.chat_task/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatTaskServer).SignUp(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
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

func _ChatTask_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatTaskServer).Chat(&chatTaskChatServer{stream})
}

type ChatTask_ChatServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatTaskChatServer struct {
	grpc.ServerStream
}

func (x *chatTaskChatServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatTaskChatServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatTask_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.chat_task",
	HandlerType: (*ChatTaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _ChatTask_SignUp_Handler,
		},
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
			StreamName:    "Chat",
			Handler:       _ChatTask_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}
