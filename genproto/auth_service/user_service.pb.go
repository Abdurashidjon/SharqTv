// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: user_service.proto

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

var File_user_service_proto protoreflect.FileDescriptor

var file_user_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xe0, 0x06, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x19, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x19, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x1a, 0x19, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x54, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x19, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x46, 0x75, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x75, 0x6c, 0x6c, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1c, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0d, 0x4f, 0x6e, 0x65, 0x32, 0x4d, 0x61, 0x6e,
	0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4f, 0x6e, 0x65, 0x32, 0x4d, 0x61, 0x6e, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1c, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x1a, 0x1c, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x51, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1c, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x1c, 0x53, 0x65, 0x6e, 0x64, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x72, 0x6c, 0x54, 0x6f, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x2a, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x79,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_user_service_proto_goTypes = []interface{}{
	(*CreateUserModel)(nil),                  // 0: genproto.CreateUserModel
	(*UpdateUserModel)(nil),                  // 1: genproto.UpdateUserModel
	(*UpdateUserPasswordModel)(nil),          // 2: genproto.UpdateUserPasswordModel
	(*GetRequest)(nil),                       // 3: genproto.GetRequest
	(*DeleteRequest)(nil),                    // 4: genproto.DeleteRequest
	(*StandardLoginModel)(nil),               // 5: genproto.StandardLoginModel
	(*One2ManyLoginModel)(nil),               // 6: genproto.One2ManyLoginModel
	(*RegisterUserModel)(nil),                // 7: genproto.RegisterUserModel
	(*ConfirmPasscodeModel)(nil),             // 8: genproto.ConfirmPasscodeModel
	(*UserEmail)(nil),                        // 9: genproto.UserEmail
	(*UpdateUserPasswordByEmailRequest)(nil), // 10: genproto.UpdateUserPasswordByEmailRequest
	(*CreatedResponse)(nil),                  // 11: genproto.CreatedResponse
	(*UserFullModel)(nil),                    // 12: genproto.UserFullModel
	(*emptypb.Empty)(nil),                    // 13: google.protobuf.Empty
	(*LoginResponseModel)(nil),               // 14: genproto.LoginResponseModel
	(*RegisteredResponse)(nil),               // 15: genproto.RegisteredResponse
}
var file_user_service_proto_depIdxs = []int32{
	0,  // 0: genproto.UserService.CreateUser:input_type -> genproto.CreateUserModel
	1,  // 1: genproto.UserService.UpdateUser:input_type -> genproto.UpdateUserModel
	2,  // 2: genproto.UserService.UpdateUserPassword:input_type -> genproto.UpdateUserPasswordModel
	3,  // 3: genproto.UserService.GetUserFull:input_type -> genproto.GetRequest
	4,  // 4: genproto.UserService.DeleteUser:input_type -> genproto.DeleteRequest
	5,  // 5: genproto.UserService.StandardLogin:input_type -> genproto.StandardLoginModel
	6,  // 6: genproto.UserService.One2ManyLogin:input_type -> genproto.One2ManyLoginModel
	7,  // 7: genproto.UserService.RegisterUser:input_type -> genproto.RegisterUserModel
	8,  // 8: genproto.UserService.ConfirmRegister:input_type -> genproto.ConfirmPasscodeModel
	9,  // 9: genproto.UserService.SendUpdatePasswordUrlToEmail:input_type -> genproto.UserEmail
	10, // 10: genproto.UserService.UpdateUserPasswordByEmail:input_type -> genproto.UpdateUserPasswordByEmailRequest
	11, // 11: genproto.UserService.CreateUser:output_type -> genproto.CreatedResponse
	11, // 12: genproto.UserService.UpdateUser:output_type -> genproto.CreatedResponse
	11, // 13: genproto.UserService.UpdateUserPassword:output_type -> genproto.CreatedResponse
	12, // 14: genproto.UserService.GetUserFull:output_type -> genproto.UserFullModel
	13, // 15: genproto.UserService.DeleteUser:output_type -> google.protobuf.Empty
	14, // 16: genproto.UserService.StandardLogin:output_type -> genproto.LoginResponseModel
	14, // 17: genproto.UserService.One2ManyLogin:output_type -> genproto.LoginResponseModel
	15, // 18: genproto.UserService.RegisterUser:output_type -> genproto.RegisteredResponse
	14, // 19: genproto.UserService.ConfirmRegister:output_type -> genproto.LoginResponseModel
	13, // 20: genproto.UserService.SendUpdatePasswordUrlToEmail:output_type -> google.protobuf.Empty
	13, // 21: genproto.UserService.UpdateUserPasswordByEmail:output_type -> google.protobuf.Empty
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_user_service_proto_init() }
func file_user_service_proto_init() {
	if File_user_service_proto != nil {
		return
	}
	file_user_proto_init()
	file_auth_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_depIdxs,
	}.Build()
	File_user_service_proto = out.File
	file_user_service_proto_rawDesc = nil
	file_user_service_proto_goTypes = nil
	file_user_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserModel, opts ...grpc.CallOption) (*CreatedResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserModel, opts ...grpc.CallOption) (*CreatedResponse, error)
	UpdateUserPassword(ctx context.Context, in *UpdateUserPasswordModel, opts ...grpc.CallOption) (*CreatedResponse, error)
	GetUserFull(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*UserFullModel, error)
	DeleteUser(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StandardLogin(ctx context.Context, in *StandardLoginModel, opts ...grpc.CallOption) (*LoginResponseModel, error)
	One2ManyLogin(ctx context.Context, in *One2ManyLoginModel, opts ...grpc.CallOption) (*LoginResponseModel, error)
	RegisterUser(ctx context.Context, in *RegisterUserModel, opts ...grpc.CallOption) (*RegisteredResponse, error)
	ConfirmRegister(ctx context.Context, in *ConfirmPasscodeModel, opts ...grpc.CallOption) (*LoginResponseModel, error)
	SendUpdatePasswordUrlToEmail(ctx context.Context, in *UserEmail, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateUserPasswordByEmail(ctx context.Context, in *UpdateUserPasswordByEmailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserModel, opts ...grpc.CallOption) (*CreatedResponse, error) {
	out := new(CreatedResponse)
	err := c.cc.Invoke(ctx, "/genproto.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserModel, opts ...grpc.CallOption) (*CreatedResponse, error) {
	out := new(CreatedResponse)
	err := c.cc.Invoke(ctx, "/genproto.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserPassword(ctx context.Context, in *UpdateUserPasswordModel, opts ...grpc.CallOption) (*CreatedResponse, error) {
	out := new(CreatedResponse)
	err := c.cc.Invoke(ctx, "/genproto.UserService/UpdateUserPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserFull(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*UserFullModel, error) {
	out := new(UserFullModel)
	err := c.cc.Invoke(ctx, "/genproto.UserService/GetUserFull", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) StandardLogin(ctx context.Context, in *StandardLoginModel, opts ...grpc.CallOption) (*LoginResponseModel, error) {
	out := new(LoginResponseModel)
	err := c.cc.Invoke(ctx, "/genproto.UserService/StandardLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) One2ManyLogin(ctx context.Context, in *One2ManyLoginModel, opts ...grpc.CallOption) (*LoginResponseModel, error) {
	out := new(LoginResponseModel)
	err := c.cc.Invoke(ctx, "/genproto.UserService/One2ManyLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RegisterUser(ctx context.Context, in *RegisterUserModel, opts ...grpc.CallOption) (*RegisteredResponse, error) {
	out := new(RegisteredResponse)
	err := c.cc.Invoke(ctx, "/genproto.UserService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ConfirmRegister(ctx context.Context, in *ConfirmPasscodeModel, opts ...grpc.CallOption) (*LoginResponseModel, error) {
	out := new(LoginResponseModel)
	err := c.cc.Invoke(ctx, "/genproto.UserService/ConfirmRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SendUpdatePasswordUrlToEmail(ctx context.Context, in *UserEmail, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.UserService/SendUpdatePasswordUrlToEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserPasswordByEmail(ctx context.Context, in *UpdateUserPasswordByEmailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.UserService/UpdateUserPasswordByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserModel) (*CreatedResponse, error)
	UpdateUser(context.Context, *UpdateUserModel) (*CreatedResponse, error)
	UpdateUserPassword(context.Context, *UpdateUserPasswordModel) (*CreatedResponse, error)
	GetUserFull(context.Context, *GetRequest) (*UserFullModel, error)
	DeleteUser(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	StandardLogin(context.Context, *StandardLoginModel) (*LoginResponseModel, error)
	One2ManyLogin(context.Context, *One2ManyLoginModel) (*LoginResponseModel, error)
	RegisterUser(context.Context, *RegisterUserModel) (*RegisteredResponse, error)
	ConfirmRegister(context.Context, *ConfirmPasscodeModel) (*LoginResponseModel, error)
	SendUpdatePasswordUrlToEmail(context.Context, *UserEmail) (*emptypb.Empty, error)
	UpdateUserPasswordByEmail(context.Context, *UpdateUserPasswordByEmailRequest) (*emptypb.Empty, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserModel) (*CreatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserModel) (*CreatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUserServiceServer) UpdateUserPassword(context.Context, *UpdateUserPasswordModel) (*CreatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPassword not implemented")
}
func (*UnimplementedUserServiceServer) GetUserFull(context.Context, *GetRequest) (*UserFullModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFull not implemented")
}
func (*UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (*UnimplementedUserServiceServer) StandardLogin(context.Context, *StandardLoginModel) (*LoginResponseModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StandardLogin not implemented")
}
func (*UnimplementedUserServiceServer) One2ManyLogin(context.Context, *One2ManyLoginModel) (*LoginResponseModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method One2ManyLogin not implemented")
}
func (*UnimplementedUserServiceServer) RegisterUser(context.Context, *RegisterUserModel) (*RegisteredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (*UnimplementedUserServiceServer) ConfirmRegister(context.Context, *ConfirmPasscodeModel) (*LoginResponseModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmRegister not implemented")
}
func (*UnimplementedUserServiceServer) SendUpdatePasswordUrlToEmail(context.Context, *UserEmail) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUpdatePasswordUrlToEmail not implemented")
}
func (*UnimplementedUserServiceServer) UpdateUserPasswordByEmail(context.Context, *UpdateUserPasswordByEmailRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPasswordByEmail not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPasswordModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserService/UpdateUserPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserPassword(ctx, req.(*UpdateUserPasswordModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserFull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserFull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserService/GetUserFull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserFull(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_StandardLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StandardLoginModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).StandardLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserService/StandardLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).StandardLogin(ctx, req.(*StandardLoginModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_One2ManyLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(One2ManyLoginModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).One2ManyLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserService/One2ManyLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).One2ManyLogin(ctx, req.(*One2ManyLoginModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterUser(ctx, req.(*RegisterUserModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ConfirmRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmPasscodeModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ConfirmRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserService/ConfirmRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ConfirmRegister(ctx, req.(*ConfirmPasscodeModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SendUpdatePasswordUrlToEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SendUpdatePasswordUrlToEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserService/SendUpdatePasswordUrlToEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SendUpdatePasswordUrlToEmail(ctx, req.(*UserEmail))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserPasswordByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPasswordByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserPasswordByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserService/UpdateUserPasswordByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserPasswordByEmail(ctx, req.(*UpdateUserPasswordByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateUserPassword",
			Handler:    _UserService_UpdateUserPassword_Handler,
		},
		{
			MethodName: "GetUserFull",
			Handler:    _UserService_GetUserFull_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "StandardLogin",
			Handler:    _UserService_StandardLogin_Handler,
		},
		{
			MethodName: "One2ManyLogin",
			Handler:    _UserService_One2ManyLogin_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _UserService_RegisterUser_Handler,
		},
		{
			MethodName: "ConfirmRegister",
			Handler:    _UserService_ConfirmRegister_Handler,
		},
		{
			MethodName: "SendUpdatePasswordUrlToEmail",
			Handler:    _UserService_SendUpdatePasswordUrlToEmail_Handler,
		},
		{
			MethodName: "UpdateUserPasswordByEmail",
			Handler:    _UserService_UpdateUserPasswordByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}
