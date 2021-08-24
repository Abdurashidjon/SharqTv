// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: respondent.proto

package user_service

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

type Respondent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email  string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone  string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	SberId string `protobuf:"bytes,5,opt,name=sber_id,json=sberId,proto3" json:"sber_id,omitempty"`
}

func (x *Respondent) Reset() {
	*x = Respondent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Respondent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Respondent) ProtoMessage() {}

func (x *Respondent) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Respondent.ProtoReflect.Descriptor instead.
func (*Respondent) Descriptor() ([]byte, []int) {
	return file_respondent_proto_rawDescGZIP(), []int{0}
}

func (x *Respondent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Respondent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Respondent) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Respondent) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Respondent) GetSberId() string {
	if x != nil {
		return x.SberId
	}
	return ""
}

type GetAllRespondentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Respondents []*Respondent `protobuf:"bytes,1,rep,name=respondents,proto3" json:"respondents,omitempty"`
	Count       int32         `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllRespondentResponse) Reset() {
	*x = GetAllRespondentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRespondentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRespondentResponse) ProtoMessage() {}

func (x *GetAllRespondentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRespondentResponse.ProtoReflect.Descriptor instead.
func (*GetAllRespondentResponse) Descriptor() ([]byte, []int) {
	return file_respondent_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllRespondentResponse) GetRespondents() []*Respondent {
	if x != nil {
		return x.Respondents
	}
	return nil
}

func (x *GetAllRespondentResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_respondent_proto protoreflect.FileDescriptor

var file_respondent_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x0a,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x62, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x36, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x17, 0x5a,
	0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_respondent_proto_rawDescOnce sync.Once
	file_respondent_proto_rawDescData = file_respondent_proto_rawDesc
)

func file_respondent_proto_rawDescGZIP() []byte {
	file_respondent_proto_rawDescOnce.Do(func() {
		file_respondent_proto_rawDescData = protoimpl.X.CompressGZIP(file_respondent_proto_rawDescData)
	})
	return file_respondent_proto_rawDescData
}

var file_respondent_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_respondent_proto_goTypes = []interface{}{
	(*Respondent)(nil),               // 0: genproto.Respondent
	(*GetAllRespondentResponse)(nil), // 1: genproto.GetAllRespondentResponse
}
var file_respondent_proto_depIdxs = []int32{
	0, // 0: genproto.GetAllRespondentResponse.respondents:type_name -> genproto.Respondent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_respondent_proto_init() }
func file_respondent_proto_init() {
	if File_respondent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_respondent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Respondent); i {
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
		file_respondent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRespondentResponse); i {
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
			RawDescriptor: file_respondent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_respondent_proto_goTypes,
		DependencyIndexes: file_respondent_proto_depIdxs,
		MessageInfos:      file_respondent_proto_msgTypes,
	}.Build()
	File_respondent_proto = out.File
	file_respondent_proto_rawDesc = nil
	file_respondent_proto_goTypes = nil
	file_respondent_proto_depIdxs = nil
}
