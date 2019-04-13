// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/profile.proto

package profile

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

type User struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_c50ceb3f9af5266e, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *User) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *User) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *User) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_c50ceb3f9af5266e, []int{1}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Response struct {
	Created              bool     `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c50ceb3f9af5266e, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "profile.User")
	proto.RegisterType((*Container)(nil), "profile.Container")
	proto.RegisterType((*Response)(nil), "profile.Response")
}

func init() { proto.RegisterFile("proto/profile.proto", fileDescriptor_c50ceb3f9af5266e) }

var fileDescriptor_c50ceb3f9af5266e = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4e, 0xb3, 0x40,
	0x14, 0x85, 0x7f, 0x28, 0x2d, 0x70, 0xc9, 0xdf, 0xc4, 0x31, 0xd1, 0x89, 0x2e, 0x44, 0x56, 0x6c,
	0xac, 0x09, 0x3e, 0x80, 0x0b, 0x17, 0x0d, 0x3b, 0x33, 0xc6, 0xb5, 0xa9, 0xcc, 0xb5, 0xde, 0xa4,
	0x65, 0xc8, 0xcc, 0xb4, 0x3e, 0x8f, 0x6f, 0x6a, 0x98, 0x01, 0x52, 0xe3, 0xf2, 0x3b, 0x67, 0x4e,
	0xce, 0xe1, 0x02, 0xe7, 0x9d, 0x56, 0x56, 0xdd, 0x77, 0x5a, 0x7d, 0xd0, 0x0e, 0x57, 0x8e, 0x58,
	0x3c, 0x60, 0xf1, 0x1d, 0x40, 0xf4, 0x6a, 0x50, 0xb3, 0x25, 0x84, 0x24, 0x79, 0x90, 0x07, 0x65,
	0x2a, 0x42, 0x92, 0x2c, 0x87, 0x4c, 0xa2, 0x69, 0x34, 0x75, 0x96, 0x54, 0xcb, 0x43, 0x67, 0x9c,
	0x4a, 0xec, 0x02, 0x16, 0x5f, 0x48, 0xdb, 0x4f, 0xcb, 0x67, 0x79, 0x50, 0xce, 0xc5, 0x40, 0xac,
	0x02, 0x68, 0x54, 0x6b, 0x37, 0xd4, 0xa2, 0x36, 0x3c, 0xca, 0x67, 0x65, 0x56, 0xb1, 0xd5, 0xd8,
	0xff, 0x34, 0x5a, 0xe2, 0xe4, 0x15, 0xbb, 0x86, 0xf4, 0x88, 0xc6, 0xe0, 0xee, 0x8d, 0x24, 0x9f,
	0xbb, 0xae, 0xc4, 0x0b, 0xb5, 0x2c, 0xf6, 0x90, 0x4e, 0xa9, 0x3f, 0x3b, 0x6f, 0x20, 0x6b, 0x0e,
	0xc6, 0xaa, 0x3d, 0xea, 0x3e, 0xeb, 0x77, 0xc2, 0x28, 0xd5, 0xb2, 0x9f, 0xa9, 0x34, 0x6d, 0xa9,
	0x75, 0x33, 0x53, 0x31, 0x10, 0xbb, 0x84, 0xf8, 0x60, 0x7c, 0x28, 0xf2, 0x46, 0x8f, 0xb5, 0x2c,
	0xd6, 0x90, 0x08, 0x34, 0x9d, 0x6a, 0x0d, 0x32, 0x0e, 0x71, 0xa3, 0x71, 0x63, 0xd1, 0x57, 0x26,
	0x62, 0x44, 0x76, 0x0b, 0x51, 0xff, 0xde, 0x15, 0x66, 0xd5, 0xff, 0xe9, 0xfb, 0xfa, 0x63, 0x0a,
	0x67, 0x55, 0x8f, 0xb0, 0x7c, 0xf6, 0xea, 0x0b, 0xea, 0x23, 0x35, 0xc8, 0xee, 0x20, 0x5e, 0xa3,
	0x75, 0xf7, 0xfe, 0x9d, 0xb8, 0x3a, 0x9b, 0x70, 0xec, 0x2e, 0xfe, 0xbd, 0x2f, 0xdc, 0xcf, 0x7a,
	0xf8, 0x09, 0x00, 0x00, 0xff, 0xff, 0x89, 0x8d, 0x65, 0xbf, 0xc3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileServiceClient interface {
	GetUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
}

type profileServiceClient struct {
	cc *grpc.ClientConn
}

func NewProfileServiceClient(cc *grpc.ClientConn) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) GetUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
type ProfileServiceServer interface {
	GetUser(context.Context, *User) (*Response, error)
}

// UnimplementedProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (*UnimplementedProfileServiceServer) GetUser(ctx context.Context, req *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func RegisterProfileServiceServer(s *grpc.Server, srv ProfileServiceServer) {
	s.RegisterService(&_ProfileService_serviceDesc, srv)
}

func _ProfileService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profile.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _ProfileService_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/profile.proto",
}
