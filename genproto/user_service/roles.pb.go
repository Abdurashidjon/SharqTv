// Code generated by protoc-gen-go. DO NOT EDIT.
// source: roles.proto

package user_service

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

func init() { proto.RegisterFile("roles.proto", fileDescriptor_b96358c61fe6d5ae) }

var fileDescriptor_b96358c61fe6d5ae = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xca, 0xcf, 0x49,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x4f, 0xcd, 0x03, 0xb3, 0xa4, 0xb8,
	0x4a, 0x8b, 0x53, 0x8b, 0x20, 0xa2, 0x46, 0x5f, 0x18, 0xb9, 0x78, 0x82, 0x40, 0xaa, 0x82, 0x53,
	0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xf4, 0xb9, 0xd8, 0x9c, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x85,
	0xf8, 0xf4, 0x60, 0x3a, 0xf4, 0x40, 0x2a, 0xa4, 0x84, 0x11, 0x7c, 0xcf, 0x14, 0xdf, 0xd4, 0xe2,
	0xe2, 0xc4, 0xf4, 0x54, 0x25, 0x06, 0x21, 0x2b, 0x2e, 0x4e, 0x90, 0xb4, 0x7b, 0x6a, 0x89, 0xa7,
	0x8b, 0x90, 0x08, 0x42, 0x8d, 0x7b, 0x6a, 0x49, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x94,
	0x28, 0x9a, 0x68, 0x71, 0x41, 0x7e, 0x5e, 0x31, 0x48, 0xaf, 0x1d, 0x17, 0x57, 0x68, 0x41, 0x4a,
	0x62, 0x49, 0x2a, 0xc8, 0x04, 0x64, 0xcd, 0x10, 0x51, 0x9f, 0xcc, 0xe2, 0x12, 0x29, 0x09, 0x74,
	0x51, 0x24, 0xfd, 0xe6, 0x5c, 0x3c, 0x2e, 0xa9, 0x39, 0xa9, 0x10, 0xfd, 0x9e, 0x29, 0x38, 0xac,
	0xe7, 0x47, 0x88, 0xba, 0xe6, 0x16, 0x94, 0x54, 0x2a, 0x31, 0x38, 0x89, 0x47, 0x89, 0xc2, 0xc4,
	0xf4, 0x41, 0xa1, 0x11, 0x5f, 0x0c, 0xf1, 0x7e, 0x12, 0x1b, 0x58, 0xcc, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x30, 0x11, 0x3f, 0x8d, 0x3b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RolesServiceClient is the client API for RolesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RolesServiceClient interface {
	Create(ctx context.Context, in *Role, opts ...grpc.CallOption) (*IdMessage, error)
	RoleGetID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	UpdateRole(ctx context.Context, in *UpdateList, opts ...grpc.CallOption) (*UpdateResponse, error)
	DeleteRoleId(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Empty, error)
}

type rolesServiceClient struct {
	cc *grpc.ClientConn
}

func NewRolesServiceClient(cc *grpc.ClientConn) RolesServiceClient {
	return &rolesServiceClient{cc}
}

func (c *rolesServiceClient) Create(ctx context.Context, in *Role, opts ...grpc.CallOption) (*IdMessage, error) {
	out := new(IdMessage)
	err := c.cc.Invoke(ctx, "/genproto.RolesService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) RoleGetID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/genproto.RolesService/RoleGetID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) UpdateRole(ctx context.Context, in *UpdateList, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/genproto.RolesService/UpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) DeleteRoleId(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/genproto.RolesService/DeleteRoleId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolesServiceServer is the server API for RolesService service.
type RolesServiceServer interface {
	Create(context.Context, *Role) (*IdMessage, error)
	RoleGetID(context.Context, *GetRequest) (*GetResponse, error)
	UpdateRole(context.Context, *UpdateList) (*UpdateResponse, error)
	DeleteRoleId(context.Context, *GetRequest) (*Empty, error)
}

// UnimplementedRolesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRolesServiceServer struct {
}

func (*UnimplementedRolesServiceServer) Create(ctx context.Context, req *Role) (*IdMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedRolesServiceServer) RoleGetID(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleGetID not implemented")
}
func (*UnimplementedRolesServiceServer) UpdateRole(ctx context.Context, req *UpdateList) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (*UnimplementedRolesServiceServer) DeleteRoleId(ctx context.Context, req *GetRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoleId not implemented")
}

func RegisterRolesServiceServer(s *grpc.Server, srv RolesServiceServer) {
	s.RegisterService(&_RolesService_serviceDesc, srv)
}

func _RolesService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.RolesService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).Create(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_RoleGetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).RoleGetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.RolesService/RoleGetID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).RoleGetID(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.RolesService/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).UpdateRole(ctx, req.(*UpdateList))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_DeleteRoleId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).DeleteRoleId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.RolesService/DeleteRoleId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).DeleteRoleId(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RolesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.RolesService",
	HandlerType: (*RolesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RolesService_Create_Handler,
		},
		{
			MethodName: "RoleGetID",
			Handler:    _RolesService_RoleGetID_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _RolesService_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRoleId",
			Handler:    _RolesService_DeleteRoleId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "roles.proto",
}