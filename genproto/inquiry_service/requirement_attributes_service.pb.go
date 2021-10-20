// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: requirement_attributes_service.proto

package inquiry_service

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

type GetAllAttributeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset    int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit     int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	InputType string `protobuf:"bytes,4,opt,name=input_type,json=inputType,proto3" json:"input_type,omitempty"`
	DataType  string `protobuf:"bytes,5,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
}

func (x *GetAllAttributeRequest) Reset() {
	*x = GetAllAttributeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requirement_attributes_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAttributeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAttributeRequest) ProtoMessage() {}

func (x *GetAllAttributeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requirement_attributes_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAttributeRequest.ProtoReflect.Descriptor instead.
func (*GetAllAttributeRequest) Descriptor() ([]byte, []int) {
	return file_requirement_attributes_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllAttributeRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAllAttributeRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllAttributeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAllAttributeRequest) GetInputType() string {
	if x != nil {
		return x.InputType
	}
	return ""
}

func (x *GetAllAttributeRequest) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

type GetAllAttributeValuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset                 int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	RequirementAttributeId string `protobuf:"bytes,3,opt,name=requirement_attribute_id,json=requirementAttributeId,proto3" json:"requirement_attribute_id,omitempty"`
	Value                  string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetAllAttributeValuesRequest) Reset() {
	*x = GetAllAttributeValuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requirement_attributes_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAttributeValuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAttributeValuesRequest) ProtoMessage() {}

func (x *GetAllAttributeValuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requirement_attributes_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAttributeValuesRequest.ProtoReflect.Descriptor instead.
func (*GetAllAttributeValuesRequest) Descriptor() ([]byte, []int) {
	return file_requirement_attributes_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllAttributeValuesRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAllAttributeValuesRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllAttributeValuesRequest) GetRequirementAttributeId() string {
	if x != nil {
		return x.RequirementAttributeId
	}
	return ""
}

func (x *GetAllAttributeValuesRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetAttributeValuesRequestByValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetAttributeValuesRequestByValue) Reset() {
	*x = GetAttributeValuesRequestByValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requirement_attributes_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttributeValuesRequestByValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttributeValuesRequestByValue) ProtoMessage() {}

func (x *GetAttributeValuesRequestByValue) ProtoReflect() protoreflect.Message {
	mi := &file_requirement_attributes_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttributeValuesRequestByValue.ProtoReflect.Descriptor instead.
func (*GetAttributeValuesRequestByValue) Descriptor() ([]byte, []int) {
	return file_requirement_attributes_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAttributeValuesRequestByValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_requirement_attributes_service_proto protoreflect.FileDescriptor

var file_requirement_attributes_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x38, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xc1, 0x05,
	0x0a, 0x11, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x1a, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x64, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49,
	0x64, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49,
	0x64, 0x1a, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x20, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x26, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1a,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1c, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x64, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a,
	0x1c, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x64, 0x22, 0x00, 0x12,
	0x65, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e,
	0x71, 0x75, 0x69, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_requirement_attributes_service_proto_rawDescOnce sync.Once
	file_requirement_attributes_service_proto_rawDescData = file_requirement_attributes_service_proto_rawDesc
)

func file_requirement_attributes_service_proto_rawDescGZIP() []byte {
	file_requirement_attributes_service_proto_rawDescOnce.Do(func() {
		file_requirement_attributes_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_requirement_attributes_service_proto_rawDescData)
	})
	return file_requirement_attributes_service_proto_rawDescData
}

var file_requirement_attributes_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_requirement_attributes_service_proto_goTypes = []interface{}{
	(*GetAllAttributeRequest)(nil),           // 0: genproto.GetAllAttributeRequest
	(*GetAllAttributeValuesRequest)(nil),     // 1: genproto.GetAllAttributeValuesRequest
	(*GetAttributeValuesRequestByValue)(nil), // 2: genproto.GetAttributeValuesRequestByValue
	(*Attribute)(nil),                        // 3: genproto.Attribute
	(*AttributeId)(nil),                      // 4: genproto.AttributeId
	(*RequirementValue)(nil),                 // 5: genproto.RequirementValue
	(*GetAllAttributeResponse)(nil),          // 6: genproto.GetAllAttributeResponse
	(*GetAllAttributeValuesResponse)(nil),    // 7: genproto.GetAllAttributeValuesResponse
	(*RequirementValueId)(nil),               // 8: genproto.RequirementValueId
	(*emptypb.Empty)(nil),                    // 9: google.protobuf.Empty
}
var file_requirement_attributes_service_proto_depIdxs = []int32{
	3, // 0: genproto.AttributesService.Create:input_type -> genproto.Attribute
	3, // 1: genproto.AttributesService.Update:input_type -> genproto.Attribute
	4, // 2: genproto.AttributesService.Get:input_type -> genproto.AttributeId
	0, // 3: genproto.AttributesService.GetAll:input_type -> genproto.GetAllAttributeRequest
	1, // 4: genproto.AttributesService.GetAllAttributeValues:input_type -> genproto.GetAllAttributeValuesRequest
	5, // 5: genproto.AttributesService.UpdateAttributeValues:input_type -> genproto.RequirementValue
	5, // 6: genproto.AttributesService.CreateAttributeValues:input_type -> genproto.RequirementValue
	2, // 7: genproto.AttributesService.GetAttributeValuesByValue:input_type -> genproto.GetAttributeValuesRequestByValue
	4, // 8: genproto.AttributesService.Delete:input_type -> genproto.AttributeId
	4, // 9: genproto.AttributesService.Create:output_type -> genproto.AttributeId
	4, // 10: genproto.AttributesService.Update:output_type -> genproto.AttributeId
	3, // 11: genproto.AttributesService.Get:output_type -> genproto.Attribute
	6, // 12: genproto.AttributesService.GetAll:output_type -> genproto.GetAllAttributeResponse
	7, // 13: genproto.AttributesService.GetAllAttributeValues:output_type -> genproto.GetAllAttributeValuesResponse
	8, // 14: genproto.AttributesService.UpdateAttributeValues:output_type -> genproto.RequirementValueId
	8, // 15: genproto.AttributesService.CreateAttributeValues:output_type -> genproto.RequirementValueId
	5, // 16: genproto.AttributesService.GetAttributeValuesByValue:output_type -> genproto.RequirementValue
	9, // 17: genproto.AttributesService.Delete:output_type -> google.protobuf.Empty
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_requirement_attributes_service_proto_init() }
func file_requirement_attributes_service_proto_init() {
	if File_requirement_attributes_service_proto != nil {
		return
	}
	file_requirement_attributes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_requirement_attributes_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAttributeRequest); i {
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
		file_requirement_attributes_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAttributeValuesRequest); i {
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
		file_requirement_attributes_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttributeValuesRequestByValue); i {
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
			RawDescriptor: file_requirement_attributes_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_requirement_attributes_service_proto_goTypes,
		DependencyIndexes: file_requirement_attributes_service_proto_depIdxs,
		MessageInfos:      file_requirement_attributes_service_proto_msgTypes,
	}.Build()
	File_requirement_attributes_service_proto = out.File
	file_requirement_attributes_service_proto_rawDesc = nil
	file_requirement_attributes_service_proto_goTypes = nil
	file_requirement_attributes_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AttributesServiceClient is the client API for AttributesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AttributesServiceClient interface {
	Create(ctx context.Context, in *Attribute, opts ...grpc.CallOption) (*AttributeId, error)
	Update(ctx context.Context, in *Attribute, opts ...grpc.CallOption) (*AttributeId, error)
	Get(ctx context.Context, in *AttributeId, opts ...grpc.CallOption) (*Attribute, error)
	GetAll(ctx context.Context, in *GetAllAttributeRequest, opts ...grpc.CallOption) (*GetAllAttributeResponse, error)
	GetAllAttributeValues(ctx context.Context, in *GetAllAttributeValuesRequest, opts ...grpc.CallOption) (*GetAllAttributeValuesResponse, error)
	UpdateAttributeValues(ctx context.Context, in *RequirementValue, opts ...grpc.CallOption) (*RequirementValueId, error)
	CreateAttributeValues(ctx context.Context, in *RequirementValue, opts ...grpc.CallOption) (*RequirementValueId, error)
	GetAttributeValuesByValue(ctx context.Context, in *GetAttributeValuesRequestByValue, opts ...grpc.CallOption) (*RequirementValue, error)
	Delete(ctx context.Context, in *AttributeId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type attributesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAttributesServiceClient(cc grpc.ClientConnInterface) AttributesServiceClient {
	return &attributesServiceClient{cc}
}

func (c *attributesServiceClient) Create(ctx context.Context, in *Attribute, opts ...grpc.CallOption) (*AttributeId, error) {
	out := new(AttributeId)
	err := c.cc.Invoke(ctx, "/genproto.AttributesService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributesServiceClient) Update(ctx context.Context, in *Attribute, opts ...grpc.CallOption) (*AttributeId, error) {
	out := new(AttributeId)
	err := c.cc.Invoke(ctx, "/genproto.AttributesService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributesServiceClient) Get(ctx context.Context, in *AttributeId, opts ...grpc.CallOption) (*Attribute, error) {
	out := new(Attribute)
	err := c.cc.Invoke(ctx, "/genproto.AttributesService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributesServiceClient) GetAll(ctx context.Context, in *GetAllAttributeRequest, opts ...grpc.CallOption) (*GetAllAttributeResponse, error) {
	out := new(GetAllAttributeResponse)
	err := c.cc.Invoke(ctx, "/genproto.AttributesService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributesServiceClient) GetAllAttributeValues(ctx context.Context, in *GetAllAttributeValuesRequest, opts ...grpc.CallOption) (*GetAllAttributeValuesResponse, error) {
	out := new(GetAllAttributeValuesResponse)
	err := c.cc.Invoke(ctx, "/genproto.AttributesService/GetAllAttributeValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributesServiceClient) UpdateAttributeValues(ctx context.Context, in *RequirementValue, opts ...grpc.CallOption) (*RequirementValueId, error) {
	out := new(RequirementValueId)
	err := c.cc.Invoke(ctx, "/genproto.AttributesService/UpdateAttributeValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributesServiceClient) CreateAttributeValues(ctx context.Context, in *RequirementValue, opts ...grpc.CallOption) (*RequirementValueId, error) {
	out := new(RequirementValueId)
	err := c.cc.Invoke(ctx, "/genproto.AttributesService/CreateAttributeValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributesServiceClient) GetAttributeValuesByValue(ctx context.Context, in *GetAttributeValuesRequestByValue, opts ...grpc.CallOption) (*RequirementValue, error) {
	out := new(RequirementValue)
	err := c.cc.Invoke(ctx, "/genproto.AttributesService/GetAttributeValuesByValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributesServiceClient) Delete(ctx context.Context, in *AttributeId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.AttributesService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttributesServiceServer is the server API for AttributesService service.
type AttributesServiceServer interface {
	Create(context.Context, *Attribute) (*AttributeId, error)
	Update(context.Context, *Attribute) (*AttributeId, error)
	Get(context.Context, *AttributeId) (*Attribute, error)
	GetAll(context.Context, *GetAllAttributeRequest) (*GetAllAttributeResponse, error)
	GetAllAttributeValues(context.Context, *GetAllAttributeValuesRequest) (*GetAllAttributeValuesResponse, error)
	UpdateAttributeValues(context.Context, *RequirementValue) (*RequirementValueId, error)
	CreateAttributeValues(context.Context, *RequirementValue) (*RequirementValueId, error)
	GetAttributeValuesByValue(context.Context, *GetAttributeValuesRequestByValue) (*RequirementValue, error)
	Delete(context.Context, *AttributeId) (*emptypb.Empty, error)
}

// UnimplementedAttributesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAttributesServiceServer struct {
}

func (*UnimplementedAttributesServiceServer) Create(context.Context, *Attribute) (*AttributeId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAttributesServiceServer) Update(context.Context, *Attribute) (*AttributeId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedAttributesServiceServer) Get(context.Context, *AttributeId) (*Attribute, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAttributesServiceServer) GetAll(context.Context, *GetAllAttributeRequest) (*GetAllAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedAttributesServiceServer) GetAllAttributeValues(context.Context, *GetAllAttributeValuesRequest) (*GetAllAttributeValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAttributeValues not implemented")
}
func (*UnimplementedAttributesServiceServer) UpdateAttributeValues(context.Context, *RequirementValue) (*RequirementValueId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAttributeValues not implemented")
}
func (*UnimplementedAttributesServiceServer) CreateAttributeValues(context.Context, *RequirementValue) (*RequirementValueId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAttributeValues not implemented")
}
func (*UnimplementedAttributesServiceServer) GetAttributeValuesByValue(context.Context, *GetAttributeValuesRequestByValue) (*RequirementValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttributeValuesByValue not implemented")
}
func (*UnimplementedAttributesServiceServer) Delete(context.Context, *AttributeId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterAttributesServiceServer(s *grpc.Server, srv AttributesServiceServer) {
	s.RegisterService(&_AttributesService_serviceDesc, srv)
}

func _AttributesService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Attribute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributesServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AttributesService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributesServiceServer).Create(ctx, req.(*Attribute))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributesService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Attribute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributesServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AttributesService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributesServiceServer).Update(ctx, req.(*Attribute))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributesService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttributeId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributesServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AttributesService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributesServiceServer).Get(ctx, req.(*AttributeId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributesService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributesServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AttributesService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributesServiceServer).GetAll(ctx, req.(*GetAllAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributesService_GetAllAttributeValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAttributeValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributesServiceServer).GetAllAttributeValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AttributesService/GetAllAttributeValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributesServiceServer).GetAllAttributeValues(ctx, req.(*GetAllAttributeValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributesService_UpdateAttributeValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequirementValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributesServiceServer).UpdateAttributeValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AttributesService/UpdateAttributeValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributesServiceServer).UpdateAttributeValues(ctx, req.(*RequirementValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributesService_CreateAttributeValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequirementValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributesServiceServer).CreateAttributeValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AttributesService/CreateAttributeValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributesServiceServer).CreateAttributeValues(ctx, req.(*RequirementValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributesService_GetAttributeValuesByValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttributeValuesRequestByValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributesServiceServer).GetAttributeValuesByValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AttributesService/GetAttributeValuesByValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributesServiceServer).GetAttributeValuesByValue(ctx, req.(*GetAttributeValuesRequestByValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributesService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttributeId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributesServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AttributesService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributesServiceServer).Delete(ctx, req.(*AttributeId))
	}
	return interceptor(ctx, in, info, handler)
}

var _AttributesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.AttributesService",
	HandlerType: (*AttributesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AttributesService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AttributesService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AttributesService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AttributesService_GetAll_Handler,
		},
		{
			MethodName: "GetAllAttributeValues",
			Handler:    _AttributesService_GetAllAttributeValues_Handler,
		},
		{
			MethodName: "UpdateAttributeValues",
			Handler:    _AttributesService_UpdateAttributeValues_Handler,
		},
		{
			MethodName: "CreateAttributeValues",
			Handler:    _AttributesService_CreateAttributeValues_Handler,
		},
		{
			MethodName: "GetAttributeValuesByValue",
			Handler:    _AttributesService_GetAttributeValuesByValue_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AttributesService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "requirement_attributes_service.proto",
}
