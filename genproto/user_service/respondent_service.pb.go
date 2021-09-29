// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: respondent_service.proto

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

type GetAllRespondentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email  string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone  string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *GetAllRespondentRequest) Reset() {
	*x = GetAllRespondentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRespondentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRespondentRequest) ProtoMessage() {}

func (x *GetAllRespondentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRespondentRequest.ProtoReflect.Descriptor instead.
func (*GetAllRespondentRequest) Descriptor() ([]byte, []int) {
	return file_respondent_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllRespondentRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAllRespondentRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllRespondentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAllRespondentRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetAllRespondentRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_respondent_service_proto protoreflect.FileDescriptor

var file_respondent_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x32, 0xa7, 0x05,
	0x0a, 0x11, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x1a, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x12, 0x21, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x64, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x20, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x55, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x6e, 0x12, 0x24, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x23, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_respondent_service_proto_rawDescOnce sync.Once
	file_respondent_service_proto_rawDescData = file_respondent_service_proto_rawDesc
)

func file_respondent_service_proto_rawDescGZIP() []byte {
	file_respondent_service_proto_rawDescOnce.Do(func() {
		file_respondent_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_respondent_service_proto_rawDescData)
	})
	return file_respondent_service_proto_rawDescData
}

var file_respondent_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_respondent_service_proto_goTypes = []interface{}{
	(*GetAllRespondentRequest)(nil),    // 0: genproto.GetAllRespondentRequest
	(*CreateRespondent)(nil),           // 1: genproto.CreateRespondent
	(*UpdateRespondent)(nil),           // 2: genproto.UpdateRespondent
	(*RespondentID)(nil),               // 3: genproto.RespondentID
	(*UpdateRespondentPhoto)(nil),      // 4: genproto.UpdateRespondentPhoto
	(*UpdateRespondentRating)(nil),     // 5: genproto.UpdateRespondentRating
	(*UpdateRespondentInnRequest)(nil), // 6: genproto.UpdateRespondentInnRequest
	(*GetRespondentsByIdRequest)(nil),  // 7: genproto.GetRespondentsByIdRequest
	(*Respondent)(nil),                 // 8: genproto.Respondent
	(*GetAllRespondentResponse)(nil),   // 9: genproto.GetAllRespondentResponse
	(*emptypb.Empty)(nil),              // 10: google.protobuf.Empty
}
var file_respondent_service_proto_depIdxs = []int32{
	1,  // 0: genproto.RespondentService.Create:input_type -> genproto.CreateRespondent
	2,  // 1: genproto.RespondentService.Update:input_type -> genproto.UpdateRespondent
	3,  // 2: genproto.RespondentService.Get:input_type -> genproto.RespondentID
	0,  // 3: genproto.RespondentService.GetAll:input_type -> genproto.GetAllRespondentRequest
	3,  // 4: genproto.RespondentService.Delete:input_type -> genproto.RespondentID
	4,  // 5: genproto.RespondentService.UpdatePhoto:input_type -> genproto.UpdateRespondentPhoto
	5,  // 6: genproto.RespondentService.UpdateRating:input_type -> genproto.UpdateRespondentRating
	6,  // 7: genproto.RespondentService.UpdateRespondentInn:input_type -> genproto.UpdateRespondentInnRequest
	7,  // 8: genproto.RespondentService.GetRespondentsById:input_type -> genproto.GetRespondentsByIdRequest
	3,  // 9: genproto.RespondentService.Create:output_type -> genproto.RespondentID
	3,  // 10: genproto.RespondentService.Update:output_type -> genproto.RespondentID
	8,  // 11: genproto.RespondentService.Get:output_type -> genproto.Respondent
	9,  // 12: genproto.RespondentService.GetAll:output_type -> genproto.GetAllRespondentResponse
	10, // 13: genproto.RespondentService.Delete:output_type -> google.protobuf.Empty
	10, // 14: genproto.RespondentService.UpdatePhoto:output_type -> google.protobuf.Empty
	10, // 15: genproto.RespondentService.UpdateRating:output_type -> google.protobuf.Empty
	10, // 16: genproto.RespondentService.UpdateRespondentInn:output_type -> google.protobuf.Empty
	9,  // 17: genproto.RespondentService.GetRespondentsById:output_type -> genproto.GetAllRespondentResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_respondent_service_proto_init() }
func file_respondent_service_proto_init() {
	if File_respondent_service_proto != nil {
		return
	}
	file_respondent_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_respondent_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRespondentRequest); i {
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
			RawDescriptor: file_respondent_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_respondent_service_proto_goTypes,
		DependencyIndexes: file_respondent_service_proto_depIdxs,
		MessageInfos:      file_respondent_service_proto_msgTypes,
	}.Build()
	File_respondent_service_proto = out.File
	file_respondent_service_proto_rawDesc = nil
	file_respondent_service_proto_goTypes = nil
	file_respondent_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RespondentServiceClient is the client API for RespondentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RespondentServiceClient interface {
	Create(ctx context.Context, in *CreateRespondent, opts ...grpc.CallOption) (*RespondentID, error)
	Update(ctx context.Context, in *UpdateRespondent, opts ...grpc.CallOption) (*RespondentID, error)
	Get(ctx context.Context, in *RespondentID, opts ...grpc.CallOption) (*Respondent, error)
	GetAll(ctx context.Context, in *GetAllRespondentRequest, opts ...grpc.CallOption) (*GetAllRespondentResponse, error)
	Delete(ctx context.Context, in *RespondentID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdatePhoto(ctx context.Context, in *UpdateRespondentPhoto, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateRating(ctx context.Context, in *UpdateRespondentRating, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateRespondentInn(ctx context.Context, in *UpdateRespondentInnRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetRespondentsById(ctx context.Context, in *GetRespondentsByIdRequest, opts ...grpc.CallOption) (*GetAllRespondentResponse, error)
}

type respondentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRespondentServiceClient(cc grpc.ClientConnInterface) RespondentServiceClient {
	return &respondentServiceClient{cc}
}

func (c *respondentServiceClient) Create(ctx context.Context, in *CreateRespondent, opts ...grpc.CallOption) (*RespondentID, error) {
	out := new(RespondentID)
	err := c.cc.Invoke(ctx, "/genproto.RespondentService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *respondentServiceClient) Update(ctx context.Context, in *UpdateRespondent, opts ...grpc.CallOption) (*RespondentID, error) {
	out := new(RespondentID)
	err := c.cc.Invoke(ctx, "/genproto.RespondentService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *respondentServiceClient) Get(ctx context.Context, in *RespondentID, opts ...grpc.CallOption) (*Respondent, error) {
	out := new(Respondent)
	err := c.cc.Invoke(ctx, "/genproto.RespondentService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *respondentServiceClient) GetAll(ctx context.Context, in *GetAllRespondentRequest, opts ...grpc.CallOption) (*GetAllRespondentResponse, error) {
	out := new(GetAllRespondentResponse)
	err := c.cc.Invoke(ctx, "/genproto.RespondentService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *respondentServiceClient) Delete(ctx context.Context, in *RespondentID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.RespondentService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *respondentServiceClient) UpdatePhoto(ctx context.Context, in *UpdateRespondentPhoto, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.RespondentService/UpdatePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *respondentServiceClient) UpdateRating(ctx context.Context, in *UpdateRespondentRating, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.RespondentService/UpdateRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *respondentServiceClient) UpdateRespondentInn(ctx context.Context, in *UpdateRespondentInnRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.RespondentService/UpdateRespondentInn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *respondentServiceClient) GetRespondentsById(ctx context.Context, in *GetRespondentsByIdRequest, opts ...grpc.CallOption) (*GetAllRespondentResponse, error) {
	out := new(GetAllRespondentResponse)
	err := c.cc.Invoke(ctx, "/genproto.RespondentService/GetRespondentsById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RespondentServiceServer is the server API for RespondentService service.
type RespondentServiceServer interface {
	Create(context.Context, *CreateRespondent) (*RespondentID, error)
	Update(context.Context, *UpdateRespondent) (*RespondentID, error)
	Get(context.Context, *RespondentID) (*Respondent, error)
	GetAll(context.Context, *GetAllRespondentRequest) (*GetAllRespondentResponse, error)
	Delete(context.Context, *RespondentID) (*emptypb.Empty, error)
	UpdatePhoto(context.Context, *UpdateRespondentPhoto) (*emptypb.Empty, error)
	UpdateRating(context.Context, *UpdateRespondentRating) (*emptypb.Empty, error)
	UpdateRespondentInn(context.Context, *UpdateRespondentInnRequest) (*emptypb.Empty, error)
	GetRespondentsById(context.Context, *GetRespondentsByIdRequest) (*GetAllRespondentResponse, error)
}

// UnimplementedRespondentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRespondentServiceServer struct {
}

func (*UnimplementedRespondentServiceServer) Create(context.Context, *CreateRespondent) (*RespondentID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedRespondentServiceServer) Update(context.Context, *UpdateRespondent) (*RespondentID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedRespondentServiceServer) Get(context.Context, *RespondentID) (*Respondent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRespondentServiceServer) GetAll(context.Context, *GetAllRespondentRequest) (*GetAllRespondentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedRespondentServiceServer) Delete(context.Context, *RespondentID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedRespondentServiceServer) UpdatePhoto(context.Context, *UpdateRespondentPhoto) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePhoto not implemented")
}
func (*UnimplementedRespondentServiceServer) UpdateRating(context.Context, *UpdateRespondentRating) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRating not implemented")
}
func (*UnimplementedRespondentServiceServer) UpdateRespondentInn(context.Context, *UpdateRespondentInnRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRespondentInn not implemented")
}
func (*UnimplementedRespondentServiceServer) GetRespondentsById(context.Context, *GetRespondentsByIdRequest) (*GetAllRespondentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRespondentsById not implemented")
}

func RegisterRespondentServiceServer(s *grpc.Server, srv RespondentServiceServer) {
	s.RegisterService(&_RespondentService_serviceDesc, srv)
}

func _RespondentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRespondent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RespondentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.RespondentService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RespondentServiceServer).Create(ctx, req.(*CreateRespondent))
	}
	return interceptor(ctx, in, info, handler)
}

func _RespondentService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRespondent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RespondentServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.RespondentService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RespondentServiceServer).Update(ctx, req.(*UpdateRespondent))
	}
	return interceptor(ctx, in, info, handler)
}

func _RespondentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RespondentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RespondentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.RespondentService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RespondentServiceServer).Get(ctx, req.(*RespondentID))
	}
	return interceptor(ctx, in, info, handler)
}

func _RespondentService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRespondentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RespondentServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.RespondentService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RespondentServiceServer).GetAll(ctx, req.(*GetAllRespondentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RespondentService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RespondentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RespondentServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.RespondentService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RespondentServiceServer).Delete(ctx, req.(*RespondentID))
	}
	return interceptor(ctx, in, info, handler)
}

func _RespondentService_UpdatePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRespondentPhoto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RespondentServiceServer).UpdatePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.RespondentService/UpdatePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RespondentServiceServer).UpdatePhoto(ctx, req.(*UpdateRespondentPhoto))
	}
	return interceptor(ctx, in, info, handler)
}

func _RespondentService_UpdateRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRespondentRating)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RespondentServiceServer).UpdateRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.RespondentService/UpdateRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RespondentServiceServer).UpdateRating(ctx, req.(*UpdateRespondentRating))
	}
	return interceptor(ctx, in, info, handler)
}

func _RespondentService_UpdateRespondentInn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRespondentInnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RespondentServiceServer).UpdateRespondentInn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.RespondentService/UpdateRespondentInn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RespondentServiceServer).UpdateRespondentInn(ctx, req.(*UpdateRespondentInnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RespondentService_GetRespondentsById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRespondentsByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RespondentServiceServer).GetRespondentsById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.RespondentService/GetRespondentsById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RespondentServiceServer).GetRespondentsById(ctx, req.(*GetRespondentsByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RespondentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.RespondentService",
	HandlerType: (*RespondentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RespondentService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RespondentService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RespondentService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _RespondentService_GetAll_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RespondentService_Delete_Handler,
		},
		{
			MethodName: "UpdatePhoto",
			Handler:    _RespondentService_UpdatePhoto_Handler,
		},
		{
			MethodName: "UpdateRating",
			Handler:    _RespondentService_UpdateRating_Handler,
		},
		{
			MethodName: "UpdateRespondentInn",
			Handler:    _RespondentService_UpdateRespondentInn_Handler,
		},
		{
			MethodName: "GetRespondentsById",
			Handler:    _RespondentService_GetRespondentsById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "respondent_service.proto",
}
