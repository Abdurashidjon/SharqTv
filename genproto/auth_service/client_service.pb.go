// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: client_service.proto

package auth_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_client_service_proto protoreflect.FileDescriptor

var file_client_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe6, 0x03, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x15, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x50, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x60, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12,
	0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17,
	0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_client_service_proto_goTypes = []interface{}{
	(*ClientModel)(nil),                       // 0: genproto.ClientModel
	(*GetRequest)(nil),                        // 1: genproto.GetRequest
	(*emptypb.Empty)(nil),                     // 2: google.protobuf.Empty
	(*GetClientListResponse)(nil),             // 3: genproto.GetClientListResponse
	(*GetClientTypeListResponse)(nil),         // 4: genproto.GetClientTypeListResponse
	(*GetRoleListByClientTypeIdResponse)(nil), // 5: genproto.GetRoleListByClientTypeIdResponse
	(*GetClientPlatformListResponse)(nil),     // 6: genproto.GetClientPlatformListResponse
}
var file_client_service_proto_depIdxs = []int32{
	0, // 0: genproto.ClientService.AddClient:input_type -> genproto.ClientModel
	1, // 1: genproto.ClientService.GetClientList:input_type -> genproto.GetRequest
	0, // 2: genproto.ClientService.RemoveClient:input_type -> genproto.ClientModel
	1, // 3: genproto.ClientService.GetClientTypeList:input_type -> genproto.GetRequest
	1, // 4: genproto.ClientService.GetRoleListByClientTypeID:input_type -> genproto.GetRequest
	1, // 5: genproto.ClientService.GetClientPlatformList:input_type -> genproto.GetRequest
	2, // 6: genproto.ClientService.AddClient:output_type -> google.protobuf.Empty
	3, // 7: genproto.ClientService.GetClientList:output_type -> genproto.GetClientListResponse
	2, // 8: genproto.ClientService.RemoveClient:output_type -> google.protobuf.Empty
	4, // 9: genproto.ClientService.GetClientTypeList:output_type -> genproto.GetClientTypeListResponse
	5, // 10: genproto.ClientService.GetRoleListByClientTypeID:output_type -> genproto.GetRoleListByClientTypeIdResponse
	6, // 11: genproto.ClientService.GetClientPlatformList:output_type -> genproto.GetClientPlatformListResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_client_service_proto_init() }
func file_client_service_proto_init() {
	if File_client_service_proto != nil {
		return
	}
	file_auth_proto_init()
	file_role_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_client_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_service_proto_goTypes,
		DependencyIndexes: file_client_service_proto_depIdxs,
	}.Build()
	File_client_service_proto = out.File
	file_client_service_proto_rawDesc = nil
	file_client_service_proto_goTypes = nil
	file_client_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientServiceClient interface {
	AddClient(ctx context.Context, in *ClientModel, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetClientList(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetClientListResponse, error)
	RemoveClient(ctx context.Context, in *ClientModel, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetClientTypeList(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetClientTypeListResponse, error)
	GetRoleListByClientTypeID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetRoleListByClientTypeIdResponse, error)
	GetClientPlatformList(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetClientPlatformListResponse, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) AddClient(ctx context.Context, in *ClientModel, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.ClientService/AddClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClientList(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetClientListResponse, error) {
	out := new(GetClientListResponse)
	err := c.cc.Invoke(ctx, "/genproto.ClientService/GetClientList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) RemoveClient(ctx context.Context, in *ClientModel, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.ClientService/RemoveClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClientTypeList(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetClientTypeListResponse, error) {
	out := new(GetClientTypeListResponse)
	err := c.cc.Invoke(ctx, "/genproto.ClientService/GetClientTypeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetRoleListByClientTypeID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetRoleListByClientTypeIdResponse, error) {
	out := new(GetRoleListByClientTypeIdResponse)
	err := c.cc.Invoke(ctx, "/genproto.ClientService/GetRoleListByClientTypeID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClientPlatformList(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetClientPlatformListResponse, error) {
	out := new(GetClientPlatformListResponse)
	err := c.cc.Invoke(ctx, "/genproto.ClientService/GetClientPlatformList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
type ClientServiceServer interface {
	AddClient(context.Context, *ClientModel) (*emptypb.Empty, error)
	GetClientList(context.Context, *GetRequest) (*GetClientListResponse, error)
	RemoveClient(context.Context, *ClientModel) (*emptypb.Empty, error)
	GetClientTypeList(context.Context, *GetRequest) (*GetClientTypeListResponse, error)
	GetRoleListByClientTypeID(context.Context, *GetRequest) (*GetRoleListByClientTypeIdResponse, error)
	GetClientPlatformList(context.Context, *GetRequest) (*GetClientPlatformListResponse, error)
}

// UnimplementedClientServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (*UnimplementedClientServiceServer) AddClient(context.Context, *ClientModel) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClient not implemented")
}
func (*UnimplementedClientServiceServer) GetClientList(context.Context, *GetRequest) (*GetClientListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientList not implemented")
}
func (*UnimplementedClientServiceServer) RemoveClient(context.Context, *ClientModel) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveClient not implemented")
}
func (*UnimplementedClientServiceServer) GetClientTypeList(context.Context, *GetRequest) (*GetClientTypeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientTypeList not implemented")
}
func (*UnimplementedClientServiceServer) GetRoleListByClientTypeID(context.Context, *GetRequest) (*GetRoleListByClientTypeIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleListByClientTypeID not implemented")
}
func (*UnimplementedClientServiceServer) GetClientPlatformList(context.Context, *GetRequest) (*GetClientPlatformListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientPlatformList not implemented")
}

func RegisterClientServiceServer(s *grpc.Server, srv ClientServiceServer) {
	s.RegisterService(&_ClientService_serviceDesc, srv)
}

func _ClientService_AddClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).AddClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ClientService/AddClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).AddClient(ctx, req.(*ClientModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClientList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClientList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ClientService/GetClientList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClientList(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_RemoveClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).RemoveClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ClientService/RemoveClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).RemoveClient(ctx, req.(*ClientModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClientTypeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClientTypeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ClientService/GetClientTypeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClientTypeList(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetRoleListByClientTypeID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetRoleListByClientTypeID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ClientService/GetRoleListByClientTypeID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetRoleListByClientTypeID(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClientPlatformList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClientPlatformList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ClientService/GetClientPlatformList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClientPlatformList(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddClient",
			Handler:    _ClientService_AddClient_Handler,
		},
		{
			MethodName: "GetClientList",
			Handler:    _ClientService_GetClientList_Handler,
		},
		{
			MethodName: "RemoveClient",
			Handler:    _ClientService_RemoveClient_Handler,
		},
		{
			MethodName: "GetClientTypeList",
			Handler:    _ClientService_GetClientTypeList_Handler,
		},
		{
			MethodName: "GetRoleListByClientTypeID",
			Handler:    _ClientService_GetRoleListByClientTypeID_Handler,
		},
		{
			MethodName: "GetClientPlatformList",
			Handler:    _ClientService_GetClientPlatformList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client_service.proto",
}
