// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: company_service.proto

package user_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetAllCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Inn    string `protobuf:"bytes,4,opt,name=inn,proto3" json:"inn,omitempty"`
}

func (x *GetAllCompanyRequest) Reset() {
	*x = GetAllCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCompanyRequest) ProtoMessage() {}

func (x *GetAllCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCompanyRequest.ProtoReflect.Descriptor instead.
func (*GetAllCompanyRequest) Descriptor() ([]byte, []int) {
	return file_company_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllCompanyRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAllCompanyRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllCompanyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAllCompanyRequest) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

var File_company_service_proto protoreflect.FileDescriptor

var file_company_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6e, 0x6e, 0x32, 0xaf, 0x02, 0x0a, 0x0e, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x22, 0x00,
	0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x13, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x49, 0x64, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64,
	0x1a, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x1e, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_company_service_proto_rawDescOnce sync.Once
	file_company_service_proto_rawDescData = file_company_service_proto_rawDesc
)

func file_company_service_proto_rawDescGZIP() []byte {
	file_company_service_proto_rawDescOnce.Do(func() {
		file_company_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_company_service_proto_rawDescData)
	})
	return file_company_service_proto_rawDescData
}

var file_company_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_company_service_proto_goTypes = []interface{}{
	(*GetAllCompanyRequest)(nil),  // 0: genproto.GetAllCompanyRequest
	(*Company)(nil),               // 1: genproto.Company
	(*CompanyId)(nil),             // 2: genproto.CompanyId
	(*GetAllCompanyResponse)(nil), // 3: genproto.GetAllCompanyResponse
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_company_service_proto_depIdxs = []int32{
	1, // 0: genproto.CompanyService.Create:input_type -> genproto.Company
	1, // 1: genproto.CompanyService.Update:input_type -> genproto.Company
	2, // 2: genproto.CompanyService.Get:input_type -> genproto.CompanyId
	0, // 3: genproto.CompanyService.GetAll:input_type -> genproto.GetAllCompanyRequest
	2, // 4: genproto.CompanyService.Delete:input_type -> genproto.CompanyId
	2, // 5: genproto.CompanyService.Create:output_type -> genproto.CompanyId
	2, // 6: genproto.CompanyService.Update:output_type -> genproto.CompanyId
	1, // 7: genproto.CompanyService.Get:output_type -> genproto.Company
	3, // 8: genproto.CompanyService.GetAll:output_type -> genproto.GetAllCompanyResponse
	4, // 9: genproto.CompanyService.Delete:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_company_service_proto_init() }
func file_company_service_proto_init() {
	if File_company_service_proto != nil {
		return
	}
	file_company_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_company_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCompanyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_company_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_company_service_proto_goTypes,
		DependencyIndexes: file_company_service_proto_depIdxs,
		MessageInfos:      file_company_service_proto_msgTypes,
	}.Build()
	File_company_service_proto = out.File
	file_company_service_proto_rawDesc = nil
	file_company_service_proto_goTypes = nil
	file_company_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CompanyServiceClient is the client API for CompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompanyServiceClient interface {
	Create(ctx context.Context, in *Company, opts ...grpc.CallOption) (*CompanyId, error)
	Update(ctx context.Context, in *Company, opts ...grpc.CallOption) (*CompanyId, error)
	Get(ctx context.Context, in *CompanyId, opts ...grpc.CallOption) (*Company, error)
	GetAll(ctx context.Context, in *GetAllCompanyRequest, opts ...grpc.CallOption) (*GetAllCompanyResponse, error)
	Delete(ctx context.Context, in *CompanyId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type companyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyServiceClient(cc grpc.ClientConnInterface) CompanyServiceClient {
	return &companyServiceClient{cc}
}

func (c *companyServiceClient) Create(ctx context.Context, in *Company, opts ...grpc.CallOption) (*CompanyId, error) {
	out := new(CompanyId)
	err := c.cc.Invoke(ctx, "/genproto.CompanyService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) Update(ctx context.Context, in *Company, opts ...grpc.CallOption) (*CompanyId, error) {
	out := new(CompanyId)
	err := c.cc.Invoke(ctx, "/genproto.CompanyService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) Get(ctx context.Context, in *CompanyId, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, "/genproto.CompanyService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) GetAll(ctx context.Context, in *GetAllCompanyRequest, opts ...grpc.CallOption) (*GetAllCompanyResponse, error) {
	out := new(GetAllCompanyResponse)
	err := c.cc.Invoke(ctx, "/genproto.CompanyService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) Delete(ctx context.Context, in *CompanyId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.CompanyService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServiceServer is the server API for CompanyService service.
type CompanyServiceServer interface {
	Create(context.Context, *Company) (*CompanyId, error)
	Update(context.Context, *Company) (*CompanyId, error)
	Get(context.Context, *CompanyId) (*Company, error)
	GetAll(context.Context, *GetAllCompanyRequest) (*GetAllCompanyResponse, error)
	Delete(context.Context, *CompanyId) (*emptypb.Empty, error)
}

// UnimplementedCompanyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCompanyServiceServer struct {
}

func (*UnimplementedCompanyServiceServer) Create(context.Context, *Company) (*CompanyId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCompanyServiceServer) Update(context.Context, *Company) (*CompanyId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCompanyServiceServer) Get(context.Context, *CompanyId) (*Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCompanyServiceServer) GetAll(context.Context, *GetAllCompanyRequest) (*GetAllCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedCompanyServiceServer) Delete(context.Context, *CompanyId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterCompanyServiceServer(s *grpc.Server, srv CompanyServiceServer) {
	s.RegisterService(&_CompanyService_serviceDesc, srv)
}

func _CompanyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Company)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CompanyService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).Create(ctx, req.(*Company))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Company)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CompanyService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).Update(ctx, req.(*Company))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CompanyService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).Get(ctx, req.(*CompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CompanyService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetAll(ctx, req.(*GetAllCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CompanyService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).Delete(ctx, req.(*CompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

var _CompanyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.CompanyService",
	HandlerType: (*CompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CompanyService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CompanyService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CompanyService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _CompanyService_GetAll_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CompanyService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "company_service.proto",
}
