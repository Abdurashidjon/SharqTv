// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: scope.proto

package auth_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ScopeModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method   string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Requests int32  `protobuf:"varint,3,opt,name=requests,proto3" json:"requests,omitempty"`
}

func (x *ScopeModel) Reset() {
	*x = ScopeModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scope_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopeModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeModel) ProtoMessage() {}

func (x *ScopeModel) ProtoReflect() protoreflect.Message {
	mi := &file_scope_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeModel.ProtoReflect.Descriptor instead.
func (*ScopeModel) Descriptor() ([]byte, []int) {
	return file_scope_proto_rawDescGZIP(), []int{0}
}

func (x *ScopeModel) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *ScopeModel) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ScopeModel) GetRequests() int32 {
	if x != nil {
		return x.Requests
	}
	return 0
}

type ScopeListModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int32       `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Scopes *ScopeModel `protobuf:"bytes,2,opt,name=scopes,proto3" json:"scopes,omitempty"`
}

func (x *ScopeListModel) Reset() {
	*x = ScopeListModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scope_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopeListModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeListModel) ProtoMessage() {}

func (x *ScopeListModel) ProtoReflect() protoreflect.Message {
	mi := &file_scope_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeListModel.ProtoReflect.Descriptor instead.
func (*ScopeListModel) Descriptor() ([]byte, []int) {
	return file_scope_proto_rawDescGZIP(), []int{1}
}

func (x *ScopeListModel) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ScopeListModel) GetScopes() *ScopeModel {
	if x != nil {
		return x.Scopes
	}
	return nil
}

type ScopeQueryParamModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ScopeQueryParamModel) Reset() {
	*x = ScopeQueryParamModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scope_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopeQueryParamModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeQueryParamModel) ProtoMessage() {}

func (x *ScopeQueryParamModel) ProtoReflect() protoreflect.Message {
	mi := &file_scope_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeQueryParamModel.ProtoReflect.Descriptor instead.
func (*ScopeQueryParamModel) Descriptor() ([]byte, []int) {
	return file_scope_proto_rawDescGZIP(), []int{2}
}

func (x *ScopeQueryParamModel) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ScopeQueryParamModel) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_scope_proto protoreflect.FileDescriptor

var file_scope_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0a, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x54, 0x0a, 0x0e, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x73, 0x22, 0x44, 0x0a, 0x14, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scope_proto_rawDescOnce sync.Once
	file_scope_proto_rawDescData = file_scope_proto_rawDesc
)

func file_scope_proto_rawDescGZIP() []byte {
	file_scope_proto_rawDescOnce.Do(func() {
		file_scope_proto_rawDescData = protoimpl.X.CompressGZIP(file_scope_proto_rawDescData)
	})
	return file_scope_proto_rawDescData
}

var file_scope_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_scope_proto_goTypes = []interface{}{
	(*ScopeModel)(nil),           // 0: genproto.ScopeModel
	(*ScopeListModel)(nil),       // 1: genproto.ScopeListModel
	(*ScopeQueryParamModel)(nil), // 2: genproto.ScopeQueryParamModel
}
var file_scope_proto_depIdxs = []int32{
	0, // 0: genproto.ScopeListModel.scopes:type_name -> genproto.ScopeModel
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_scope_proto_init() }
func file_scope_proto_init() {
	if File_scope_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scope_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopeModel); i {
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
		file_scope_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopeListModel); i {
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
		file_scope_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopeQueryParamModel); i {
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
			RawDescriptor: file_scope_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scope_proto_goTypes,
		DependencyIndexes: file_scope_proto_depIdxs,
		MessageInfos:      file_scope_proto_msgTypes,
	}.Build()
	File_scope_proto = out.File
	file_scope_proto_rawDesc = nil
	file_scope_proto_goTypes = nil
	file_scope_proto_depIdxs = nil
}
