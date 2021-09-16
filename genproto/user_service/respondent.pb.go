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

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email       string  `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone       string  `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	SberId      string  `protobuf:"bytes,5,opt,name=sber_id,json=sberId,proto3" json:"sber_id,omitempty"`
	Photo       string  `protobuf:"bytes,6,opt,name=photo,proto3" json:"photo,omitempty"`
	Company     string  `protobuf:"bytes,7,opt,name=company,proto3" json:"company,omitempty"`
	Position    string  `protobuf:"bytes,8,opt,name=position,proto3" json:"position,omitempty"`
	Description string  `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	Rating      *Rating `protobuf:"bytes,10,opt,name=rating,proto3" json:"rating,omitempty"`
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

func (x *Respondent) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *Respondent) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *Respondent) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *Respondent) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Respondent) GetRating() *Rating {
	if x != nil {
		return x.Rating
	}
	return nil
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

type RespondentId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RespondentId) Reset() {
	*x = RespondentId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespondentId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespondentId) ProtoMessage() {}

func (x *RespondentId) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespondentId.ProtoReflect.Descriptor instead.
func (*RespondentId) Descriptor() ([]byte, []int) {
	return file_respondent_proto_rawDescGZIP(), []int{2}
}

func (x *RespondentId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateRespondentPhoto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Photo  string `protobuf:"bytes,2,opt,name=photo,proto3" json:"photo,omitempty"`
}

func (x *UpdateRespondentPhoto) Reset() {
	*x = UpdateRespondentPhoto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRespondentPhoto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRespondentPhoto) ProtoMessage() {}

func (x *UpdateRespondentPhoto) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRespondentPhoto.ProtoReflect.Descriptor instead.
func (*UpdateRespondentPhoto) Descriptor() ([]byte, []int) {
	return file_respondent_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRespondentPhoto) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateRespondentPhoto) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Communication float32 `protobuf:"fixed32,1,opt,name=communication,proto3" json:"communication,omitempty"`
	Experience    float32 `protobuf:"fixed32,2,opt,name=experience,proto3" json:"experience,omitempty"`
	Punctuality   float32 `protobuf:"fixed32,3,opt,name=punctuality,proto3" json:"punctuality,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_respondent_proto_rawDescGZIP(), []int{4}
}

func (x *Rating) GetCommunication() float32 {
	if x != nil {
		return x.Communication
	}
	return 0
}

func (x *Rating) GetExperience() float32 {
	if x != nil {
		return x.Experience
	}
	return 0
}

func (x *Rating) GetPunctuality() float32 {
	if x != nil {
		return x.Punctuality
	}
	return 0
}

type UpdateRespondentRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Communication float32 `protobuf:"fixed32,2,opt,name=communication,proto3" json:"communication,omitempty"`
	Experience    float32 `protobuf:"fixed32,3,opt,name=experience,proto3" json:"experience,omitempty"`
	Punctuality   float32 `protobuf:"fixed32,4,opt,name=punctuality,proto3" json:"punctuality,omitempty"`
}

func (x *UpdateRespondentRating) Reset() {
	*x = UpdateRespondentRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRespondentRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRespondentRating) ProtoMessage() {}

func (x *UpdateRespondentRating) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRespondentRating.ProtoReflect.Descriptor instead.
func (*UpdateRespondentRating) Descriptor() ([]byte, []int) {
	return file_respondent_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRespondentRating) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateRespondentRating) GetCommunication() float32 {
	if x != nil {
		return x.Communication
	}
	return 0
}

func (x *UpdateRespondentRating) GetExperience() float32 {
	if x != nil {
		return x.Experience
	}
	return 0
}

func (x *UpdateRespondentRating) GetPunctuality() float32 {
	if x != nil {
		return x.Punctuality
	}
	return 0
}

var File_respondent_proto protoreflect.FileDescriptor

var file_respondent_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x02, 0x0a,
	0x0a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x62,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x68, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1e, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x22, 0x70,
	0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x75, 0x6e, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0b, 0x70, 0x75, 0x6e, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x22, 0x99, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x75,
	0x6e, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0b, 0x70, 0x75, 0x6e, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x17, 0x5a, 0x15,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_respondent_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_respondent_proto_goTypes = []interface{}{
	(*Respondent)(nil),               // 0: genproto.Respondent
	(*GetAllRespondentResponse)(nil), // 1: genproto.GetAllRespondentResponse
	(*RespondentId)(nil),             // 2: genproto.RespondentId
	(*UpdateRespondentPhoto)(nil),    // 3: genproto.UpdateRespondentPhoto
	(*Rating)(nil),                   // 4: genproto.Rating
	(*UpdateRespondentRating)(nil),   // 5: genproto.UpdateRespondentRating
}
var file_respondent_proto_depIdxs = []int32{
	4, // 0: genproto.Respondent.rating:type_name -> genproto.Rating
	0, // 1: genproto.GetAllRespondentResponse.respondents:type_name -> genproto.Respondent
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
		file_respondent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespondentId); i {
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
		file_respondent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRespondentPhoto); i {
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
		file_respondent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rating); i {
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
		file_respondent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRespondentRating); i {
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
			NumMessages:   6,
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
