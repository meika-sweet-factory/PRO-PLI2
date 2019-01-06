// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package unicampus_account

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Credential struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Credential) Reset()         { *m = Credential{} }
func (m *Credential) String() string { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()    {}
func (*Credential) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *Credential) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Credential.Unmarshal(m, b)
}
func (m *Credential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Credential.Marshal(b, m, deterministic)
}
func (m *Credential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credential.Merge(m, src)
}
func (m *Credential) XXX_Size() int {
	return xxx_messageInfo_Credential.Size(m)
}
func (m *Credential) XXX_DiscardUnknown() {
	xxx_messageInfo_Credential.DiscardUnknown(m)
}

var xxx_messageInfo_Credential proto.InternalMessageInfo

func (m *Credential) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Credential) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Session struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SignInRequest struct {
	Credential           *Credential `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SignInRequest) Reset()         { *m = SignInRequest{} }
func (m *SignInRequest) String() string { return proto.CompactTextString(m) }
func (*SignInRequest) ProtoMessage()    {}
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *SignInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInRequest.Unmarshal(m, b)
}
func (m *SignInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInRequest.Marshal(b, m, deterministic)
}
func (m *SignInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInRequest.Merge(m, src)
}
func (m *SignInRequest) XXX_Size() int {
	return xxx_messageInfo_SignInRequest.Size(m)
}
func (m *SignInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignInRequest proto.InternalMessageInfo

func (m *SignInRequest) GetCredential() *Credential {
	if m != nil {
		return m.Credential
	}
	return nil
}

type SignInReply struct {
	Session              *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInReply) Reset()         { *m = SignInReply{} }
func (m *SignInReply) String() string { return proto.CompactTextString(m) }
func (*SignInReply) ProtoMessage()    {}
func (*SignInReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
}

func (m *SignInReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInReply.Unmarshal(m, b)
}
func (m *SignInReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInReply.Marshal(b, m, deterministic)
}
func (m *SignInReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInReply.Merge(m, src)
}
func (m *SignInReply) XXX_Size() int {
	return xxx_messageInfo_SignInReply.Size(m)
}
func (m *SignInReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInReply.DiscardUnknown(m)
}

var xxx_messageInfo_SignInReply proto.InternalMessageInfo

func (m *SignInReply) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type SignUpRequest struct {
	Credential           *Credential `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SignUpRequest) Reset()         { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()    {}
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{4}
}

func (m *SignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRequest.Unmarshal(m, b)
}
func (m *SignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRequest.Marshal(b, m, deterministic)
}
func (m *SignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRequest.Merge(m, src)
}
func (m *SignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignUpRequest.Size(m)
}
func (m *SignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRequest proto.InternalMessageInfo

func (m *SignUpRequest) GetCredential() *Credential {
	if m != nil {
		return m.Credential
	}
	return nil
}

type SignUpReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpReply) Reset()         { *m = SignUpReply{} }
func (m *SignUpReply) String() string { return proto.CompactTextString(m) }
func (*SignUpReply) ProtoMessage()    {}
func (*SignUpReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{5}
}

func (m *SignUpReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpReply.Unmarshal(m, b)
}
func (m *SignUpReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpReply.Marshal(b, m, deterministic)
}
func (m *SignUpReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpReply.Merge(m, src)
}
func (m *SignUpReply) XXX_Size() int {
	return xxx_messageInfo_SignUpReply.Size(m)
}
func (m *SignUpReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpReply.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpReply proto.InternalMessageInfo

type SignOutRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignOutRequest) Reset()         { *m = SignOutRequest{} }
func (m *SignOutRequest) String() string { return proto.CompactTextString(m) }
func (*SignOutRequest) ProtoMessage()    {}
func (*SignOutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{6}
}

func (m *SignOutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignOutRequest.Unmarshal(m, b)
}
func (m *SignOutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignOutRequest.Marshal(b, m, deterministic)
}
func (m *SignOutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignOutRequest.Merge(m, src)
}
func (m *SignOutRequest) XXX_Size() int {
	return xxx_messageInfo_SignOutRequest.Size(m)
}
func (m *SignOutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignOutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignOutRequest proto.InternalMessageInfo

func (m *SignOutRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SignOutReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignOutReply) Reset()         { *m = SignOutReply{} }
func (m *SignOutReply) String() string { return proto.CompactTextString(m) }
func (*SignOutReply) ProtoMessage()    {}
func (*SignOutReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{7}
}

func (m *SignOutReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignOutReply.Unmarshal(m, b)
}
func (m *SignOutReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignOutReply.Marshal(b, m, deterministic)
}
func (m *SignOutReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignOutReply.Merge(m, src)
}
func (m *SignOutReply) XXX_Size() int {
	return xxx_messageInfo_SignOutReply.Size(m)
}
func (m *SignOutReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SignOutReply.DiscardUnknown(m)
}

var xxx_messageInfo_SignOutReply proto.InternalMessageInfo

type CloseRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseRequest) Reset()         { *m = CloseRequest{} }
func (m *CloseRequest) String() string { return proto.CompactTextString(m) }
func (*CloseRequest) ProtoMessage()    {}
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{8}
}

func (m *CloseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseRequest.Unmarshal(m, b)
}
func (m *CloseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseRequest.Marshal(b, m, deterministic)
}
func (m *CloseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseRequest.Merge(m, src)
}
func (m *CloseRequest) XXX_Size() int {
	return xxx_messageInfo_CloseRequest.Size(m)
}
func (m *CloseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloseRequest proto.InternalMessageInfo

func (m *CloseRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CloseReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseReply) Reset()         { *m = CloseReply{} }
func (m *CloseReply) String() string { return proto.CompactTextString(m) }
func (*CloseReply) ProtoMessage()    {}
func (*CloseReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{9}
}

func (m *CloseReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseReply.Unmarshal(m, b)
}
func (m *CloseReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseReply.Marshal(b, m, deterministic)
}
func (m *CloseReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseReply.Merge(m, src)
}
func (m *CloseReply) XXX_Size() int {
	return xxx_messageInfo_CloseReply.Size(m)
}
func (m *CloseReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseReply.DiscardUnknown(m)
}

var xxx_messageInfo_CloseReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Credential)(nil), "unicampus_account.Credential")
	proto.RegisterType((*Session)(nil), "unicampus_account.Session")
	proto.RegisterType((*SignInRequest)(nil), "unicampus_account.SignInRequest")
	proto.RegisterType((*SignInReply)(nil), "unicampus_account.SignInReply")
	proto.RegisterType((*SignUpRequest)(nil), "unicampus_account.SignUpRequest")
	proto.RegisterType((*SignUpReply)(nil), "unicampus_account.SignUpReply")
	proto.RegisterType((*SignOutRequest)(nil), "unicampus_account.SignOutRequest")
	proto.RegisterType((*SignOutReply)(nil), "unicampus_account.SignOutReply")
	proto.RegisterType((*CloseRequest)(nil), "unicampus_account.CloseRequest")
	proto.RegisterType((*CloseReply)(nil), "unicampus_account.CloseReply")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xb7, 0xc2, 0x36, 0x7d, 0xdb, 0x0a, 0x06, 0x0f, 0xa3, 0xb0, 0x1f, 0xe6, 0xe4, 0x69,
	0x87, 0xe9, 0x55, 0x41, 0x76, 0x1a, 0x88, 0x83, 0x8e, 0x9d, 0xa5, 0xb6, 0x41, 0x82, 0x59, 0x12,
	0x9b, 0x54, 0xd9, 0x7f, 0xe8, 0x9f, 0x25, 0xcd, 0x8f, 0xae, 0x60, 0xbb, 0x93, 0xc7, 0xc7, 0xfb,
	0xbc, 0x4f, 0xbe, 0xfb, 0x8e, 0xc2, 0x38, 0x49, 0x53, 0x51, 0x70, 0xbd, 0x94, 0xb9, 0xd0, 0x02,
	0x5d, 0x15, 0x9c, 0xa6, 0xc9, 0x41, 0x16, 0xea, 0xd5, 0x2d, 0xf0, 0x23, 0xc0, 0x3a, 0x27, 0x19,
	0xe1, 0x9a, 0x26, 0x0c, 0x5d, 0x43, 0x8f, 0x1c, 0x12, 0xca, 0x26, 0xdd, 0x45, 0xf7, 0xf6, 0x32,
	0xb6, 0x03, 0x8a, 0xe0, 0x42, 0x26, 0x4a, 0x7d, 0x8b, 0x3c, 0x9b, 0x04, 0x66, 0x51, 0xcd, 0x78,
	0x0e, 0x83, 0x1d, 0x51, 0x8a, 0x0a, 0x5e, 0x1e, 0x6b, 0xf1, 0x41, 0xb8, 0x3f, 0x36, 0x03, 0x7e,
	0x81, 0xf1, 0x8e, 0xbe, 0xf3, 0x0d, 0x8f, 0xc9, 0x67, 0x41, 0x94, 0x46, 0x0f, 0x00, 0x69, 0xf5,
	0xa2, 0x61, 0x87, 0xab, 0xe9, 0xf2, 0x4f, 0xb2, 0xe5, 0x29, 0x56, 0x5c, 0x3b, 0xc0, 0x6b, 0x18,
	0x7a, 0x9f, 0x64, 0x47, 0x74, 0x0f, 0x03, 0x65, 0xdf, 0x77, 0xaa, 0xa8, 0x41, 0xe5, 0x12, 0xc6,
	0x1e, 0xf5, 0xa1, 0xf6, 0xf2, 0x9f, 0x42, 0x8d, 0x6d, 0xa8, 0xd2, 0x27, 0xd9, 0x11, 0x2f, 0x20,
	0x2c, 0xc7, 0x6d, 0xa1, 0xbd, 0x3f, 0x84, 0x80, 0x66, 0xae, 0x98, 0x80, 0x66, 0x38, 0x84, 0x51,
	0x45, 0x94, 0x17, 0x33, 0x18, 0xad, 0x99, 0x50, 0xa4, 0x8d, 0x1f, 0x01, 0xb8, 0xbd, 0x64, 0xc7,
	0xd5, 0x4f, 0x00, 0xe1, 0x93, 0x4d, 0xb4, 0x23, 0xf9, 0x17, 0x4d, 0x09, 0x7a, 0x86, 0xbe, 0xad,
	0x05, 0x2d, 0x9a, 0x0a, 0xa8, 0xff, 0x03, 0xd1, 0xec, 0x0c, 0x51, 0x86, 0xe9, 0x78, 0xdb, 0x5e,
	0xb6, 0xda, 0xaa, 0xea, 0x5a, 0x6d, 0xbe, 0x8c, 0x0e, 0xda, 0xc2, 0xc0, 0xfd, 0x58, 0x74, 0xd3,
	0x02, 0x9f, 0xaa, 0x8a, 0xe6, 0xe7, 0x10, 0x2b, 0xdc, 0x40, 0xcf, 0xb4, 0x81, 0x9a, 0xd8, 0x7a,
	0x8f, 0xd1, 0xb4, 0x1d, 0x30, 0xaa, 0xb7, 0xbe, 0xf9, 0x32, 0xee, 0x7e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x94, 0x8b, 0x55, 0x15, 0x2a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInReply, error)
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpReply, error)
	SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutReply, error)
	Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseReply, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInReply, error) {
	out := new(SignInReply)
	err := c.cc.Invoke(ctx, "/unicampus_account.AccountService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpReply, error) {
	out := new(SignUpReply)
	err := c.cc.Invoke(ctx, "/unicampus_account.AccountService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutReply, error) {
	out := new(SignOutReply)
	err := c.cc.Invoke(ctx, "/unicampus_account.AccountService/SignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseReply, error) {
	out := new(CloseReply)
	err := c.cc.Invoke(ctx, "/unicampus_account.AccountService/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	SignIn(context.Context, *SignInRequest) (*SignInReply, error)
	SignUp(context.Context, *SignUpRequest) (*SignUpReply, error)
	SignOut(context.Context, *SignOutRequest) (*SignOutReply, error)
	Close(context.Context, *CloseRequest) (*CloseReply, error)
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/unicampus_account.AccountService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/unicampus_account.AccountService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/unicampus_account.AccountService/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).SignOut(ctx, req.(*SignOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/unicampus_account.AccountService/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Close(ctx, req.(*CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "unicampus_account.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _AccountService_SignIn_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _AccountService_SignUp_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _AccountService_SignOut_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _AccountService_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}
