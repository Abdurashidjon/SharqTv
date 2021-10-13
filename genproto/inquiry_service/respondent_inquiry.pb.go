// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: respondent_inquiry.proto

package inquiry_service

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

type RespondentInquiry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RespondentId       string            `protobuf:"bytes,2,opt,name=respondent_id,json=respondentId,proto3" json:"respondent_id,omitempty"`
	InquiryId          string            `protobuf:"bytes,3,opt,name=inquiry_id,json=inquiryId,proto3" json:"inquiry_id,omitempty"`
	InquiryTimetableId string            `protobuf:"bytes,4,opt,name=inquiry_timetable_id,json=inquiryTimetableId,proto3" json:"inquiry_timetable_id,omitempty"`
	Link               string            `protobuf:"bytes,5,opt,name=link,proto3" json:"link,omitempty"`
	Note               string            `protobuf:"bytes,6,opt,name=note,proto3" json:"note,omitempty"`
	Rating             *RespondentRating `protobuf:"bytes,7,opt,name=rating,proto3" json:"rating,omitempty"`
	StatusId           string            `protobuf:"bytes,8,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
}

func (x *RespondentInquiry) Reset() {
	*x = RespondentInquiry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_inquiry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespondentInquiry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespondentInquiry) ProtoMessage() {}

func (x *RespondentInquiry) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_inquiry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespondentInquiry.ProtoReflect.Descriptor instead.
func (*RespondentInquiry) Descriptor() ([]byte, []int) {
	return file_respondent_inquiry_proto_rawDescGZIP(), []int{0}
}

func (x *RespondentInquiry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RespondentInquiry) GetRespondentId() string {
	if x != nil {
		return x.RespondentId
	}
	return ""
}

func (x *RespondentInquiry) GetInquiryId() string {
	if x != nil {
		return x.InquiryId
	}
	return ""
}

func (x *RespondentInquiry) GetInquiryTimetableId() string {
	if x != nil {
		return x.InquiryTimetableId
	}
	return ""
}

func (x *RespondentInquiry) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *RespondentInquiry) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *RespondentInquiry) GetRating() *RespondentRating {
	if x != nil {
		return x.Rating
	}
	return nil
}

func (x *RespondentInquiry) GetStatusId() string {
	if x != nil {
		return x.StatusId
	}
	return ""
}

type CreateRespondentInquiry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RespondentId       string            `protobuf:"bytes,1,opt,name=respondent_id,json=respondentId,proto3" json:"respondent_id,omitempty"`
	InquiryId          string            `protobuf:"bytes,2,opt,name=inquiry_id,json=inquiryId,proto3" json:"inquiry_id,omitempty"`
	InquiryTimetableId string            `protobuf:"bytes,3,opt,name=inquiry_timetable_id,json=inquiryTimetableId,proto3" json:"inquiry_timetable_id,omitempty"`
	Link               string            `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
	Note               string            `protobuf:"bytes,5,opt,name=note,proto3" json:"note,omitempty"`
	Rating             *RespondentRating `protobuf:"bytes,6,opt,name=rating,proto3" json:"rating,omitempty"`
	StatusId           string            `protobuf:"bytes,7,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
}

func (x *CreateRespondentInquiry) Reset() {
	*x = CreateRespondentInquiry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_inquiry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRespondentInquiry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRespondentInquiry) ProtoMessage() {}

func (x *CreateRespondentInquiry) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_inquiry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRespondentInquiry.ProtoReflect.Descriptor instead.
func (*CreateRespondentInquiry) Descriptor() ([]byte, []int) {
	return file_respondent_inquiry_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRespondentInquiry) GetRespondentId() string {
	if x != nil {
		return x.RespondentId
	}
	return ""
}

func (x *CreateRespondentInquiry) GetInquiryId() string {
	if x != nil {
		return x.InquiryId
	}
	return ""
}

func (x *CreateRespondentInquiry) GetInquiryTimetableId() string {
	if x != nil {
		return x.InquiryTimetableId
	}
	return ""
}

func (x *CreateRespondentInquiry) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *CreateRespondentInquiry) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *CreateRespondentInquiry) GetRating() *RespondentRating {
	if x != nil {
		return x.Rating
	}
	return nil
}

func (x *CreateRespondentInquiry) GetStatusId() string {
	if x != nil {
		return x.StatusId
	}
	return ""
}

type RespondentInquiryId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RespondentInquiryId) Reset() {
	*x = RespondentInquiryId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_inquiry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespondentInquiryId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespondentInquiryId) ProtoMessage() {}

func (x *RespondentInquiryId) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_inquiry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespondentInquiryId.ProtoReflect.Descriptor instead.
func (*RespondentInquiryId) Descriptor() ([]byte, []int) {
	return file_respondent_inquiry_proto_rawDescGZIP(), []int{2}
}

func (x *RespondentInquiryId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllRespondentInquiryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inquiries []*RespondentInquiry `protobuf:"bytes,1,rep,name=inquiries,proto3" json:"inquiries,omitempty"`
	Count     int32                `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllRespondentInquiryResponse) Reset() {
	*x = GetAllRespondentInquiryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_inquiry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRespondentInquiryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRespondentInquiryResponse) ProtoMessage() {}

func (x *GetAllRespondentInquiryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_inquiry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRespondentInquiryResponse.ProtoReflect.Descriptor instead.
func (*GetAllRespondentInquiryResponse) Descriptor() ([]byte, []int) {
	return file_respondent_inquiry_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllRespondentInquiryResponse) GetInquiries() []*RespondentInquiry {
	if x != nil {
		return x.Inquiries
	}
	return nil
}

func (x *GetAllRespondentInquiryResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type RespondentRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Communication float32 `protobuf:"fixed32,1,opt,name=communication,proto3" json:"communication,omitempty"`
	Experience    float32 `protobuf:"fixed32,2,opt,name=experience,proto3" json:"experience,omitempty"`
	Punctuality   float32 `protobuf:"fixed32,3,opt,name=punctuality,proto3" json:"punctuality,omitempty"`
}

func (x *RespondentRating) Reset() {
	*x = RespondentRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_inquiry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespondentRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespondentRating) ProtoMessage() {}

func (x *RespondentRating) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_inquiry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespondentRating.ProtoReflect.Descriptor instead.
func (*RespondentRating) Descriptor() ([]byte, []int) {
	return file_respondent_inquiry_proto_rawDescGZIP(), []int{4}
}

func (x *RespondentRating) GetCommunication() float32 {
	if x != nil {
		return x.Communication
	}
	return 0
}

func (x *RespondentRating) GetExperience() float32 {
	if x != nil {
		return x.Experience
	}
	return 0
}

func (x *RespondentRating) GetPunctuality() float32 {
	if x != nil {
		return x.Punctuality
	}
	return 0
}

type UpdateRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	InquiryId     string  `protobuf:"bytes,2,opt,name=inquiry_id,json=inquiryId,proto3" json:"inquiry_id,omitempty"`
	Communication float32 `protobuf:"fixed32,3,opt,name=communication,proto3" json:"communication,omitempty"`
	Experience    float32 `protobuf:"fixed32,4,opt,name=experience,proto3" json:"experience,omitempty"`
	Punctuality   float32 `protobuf:"fixed32,5,opt,name=punctuality,proto3" json:"punctuality,omitempty"`
}

func (x *UpdateRatingRequest) Reset() {
	*x = UpdateRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_inquiry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRatingRequest) ProtoMessage() {}

func (x *UpdateRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_inquiry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRatingRequest.ProtoReflect.Descriptor instead.
func (*UpdateRatingRequest) Descriptor() ([]byte, []int) {
	return file_respondent_inquiry_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRatingRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateRatingRequest) GetInquiryId() string {
	if x != nil {
		return x.InquiryId
	}
	return ""
}

func (x *UpdateRatingRequest) GetCommunication() float32 {
	if x != nil {
		return x.Communication
	}
	return 0
}

func (x *UpdateRatingRequest) GetExperience() float32 {
	if x != nil {
		return x.Experience
	}
	return 0
}

func (x *UpdateRatingRequest) GetPunctuality() float32 {
	if x != nil {
		return x.Punctuality
	}
	return 0
}

type RespondentId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RespondentId string `protobuf:"bytes,1,opt,name=respondent_id,json=respondentId,proto3" json:"respondent_id,omitempty"`
	StatusId     string `protobuf:"bytes,2,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
}

func (x *RespondentId) Reset() {
	*x = RespondentId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_inquiry_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespondentId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespondentId) ProtoMessage() {}

func (x *RespondentId) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_inquiry_proto_msgTypes[6]
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
	return file_respondent_inquiry_proto_rawDescGZIP(), []int{6}
}

func (x *RespondentId) GetRespondentId() string {
	if x != nil {
		return x.RespondentId
	}
	return ""
}

func (x *RespondentId) GetStatusId() string {
	if x != nil {
		return x.StatusId
	}
	return ""
}

type GetRespondentsByInquiryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InquiryId string `protobuf:"bytes,1,opt,name=inquiry_id,json=inquiryId,proto3" json:"inquiry_id,omitempty"`
	StatusId  string `protobuf:"bytes,2,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
	Offset    int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit     int64  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetRespondentsByInquiryRequest) Reset() {
	*x = GetRespondentsByInquiryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_inquiry_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRespondentsByInquiryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRespondentsByInquiryRequest) ProtoMessage() {}

func (x *GetRespondentsByInquiryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_inquiry_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRespondentsByInquiryRequest.ProtoReflect.Descriptor instead.
func (*GetRespondentsByInquiryRequest) Descriptor() ([]byte, []int) {
	return file_respondent_inquiry_proto_rawDescGZIP(), []int{7}
}

func (x *GetRespondentsByInquiryRequest) GetInquiryId() string {
	if x != nil {
		return x.InquiryId
	}
	return ""
}

func (x *GetRespondentsByInquiryRequest) GetStatusId() string {
	if x != nil {
		return x.StatusId
	}
	return ""
}

func (x *GetRespondentsByInquiryRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetRespondentsByInquiryRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetRespondentsByInquiryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids   []*RespondentId `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	Count int32           `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetRespondentsByInquiryResponse) Reset() {
	*x = GetRespondentsByInquiryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_inquiry_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRespondentsByInquiryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRespondentsByInquiryResponse) ProtoMessage() {}

func (x *GetRespondentsByInquiryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_inquiry_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRespondentsByInquiryResponse.ProtoReflect.Descriptor instead.
func (*GetRespondentsByInquiryResponse) Descriptor() ([]byte, []int) {
	return file_respondent_inquiry_proto_rawDescGZIP(), []int{8}
}

func (x *GetRespondentsByInquiryResponse) GetIds() []*RespondentId {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GetRespondentsByInquiryResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetInquiryByRespondentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RespondentId string `protobuf:"bytes,1,opt,name=respondent_id,json=respondentId,proto3" json:"respondent_id,omitempty"`
	StatusId     string `protobuf:"bytes,2,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
	Offset       int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit        int64  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetInquiryByRespondentRequest) Reset() {
	*x = GetInquiryByRespondentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_respondent_inquiry_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInquiryByRespondentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInquiryByRespondentRequest) ProtoMessage() {}

func (x *GetInquiryByRespondentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_respondent_inquiry_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInquiryByRespondentRequest.ProtoReflect.Descriptor instead.
func (*GetInquiryByRespondentRequest) Descriptor() ([]byte, []int) {
	return file_respondent_inquiry_proto_rawDescGZIP(), []int{9}
}

func (x *GetInquiryByRespondentRequest) GetRespondentId() string {
	if x != nil {
		return x.RespondentId
	}
	return ""
}

func (x *GetInquiryByRespondentRequest) GetStatusId() string {
	if x != nil {
		return x.StatusId
	}
	return ""
}

func (x *GetInquiryByRespondentRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetInquiryByRespondentRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_respondent_inquiry_proto protoreflect.FileDescriptor

var file_respondent_inquiry_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x71,
	0x75, 0x69, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x02, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x49, 0x64, 0x12, 0x30,
	0x0a, 0x14, 0x69, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6e,
	0x71, 0x75, 0x69, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x22, 0x88, 0x02, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x71, 0x75, 0x69, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e,
	0x71, 0x75, 0x69, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x71,
	0x75, 0x69, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x72, 0x0a, 0x1f, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x09, 0x69, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x52, 0x09,
	0x69, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x7a, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x75, 0x6e,
	0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x70, 0x75, 0x6e, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22, 0xb5, 0x01, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x75, 0x6e, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x70, 0x75, 0x6e, 0x63, 0x74, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x22, 0x50, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x49, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x71, 0x75,
	0x69, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e,
	0x71, 0x75, 0x69, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x61, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x71,
	0x75, 0x69, 0x72, 0x79, 0x42, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_respondent_inquiry_proto_rawDescOnce sync.Once
	file_respondent_inquiry_proto_rawDescData = file_respondent_inquiry_proto_rawDesc
)

func file_respondent_inquiry_proto_rawDescGZIP() []byte {
	file_respondent_inquiry_proto_rawDescOnce.Do(func() {
		file_respondent_inquiry_proto_rawDescData = protoimpl.X.CompressGZIP(file_respondent_inquiry_proto_rawDescData)
	})
	return file_respondent_inquiry_proto_rawDescData
}

var file_respondent_inquiry_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_respondent_inquiry_proto_goTypes = []interface{}{
	(*RespondentInquiry)(nil),               // 0: genproto.RespondentInquiry
	(*CreateRespondentInquiry)(nil),         // 1: genproto.CreateRespondentInquiry
	(*RespondentInquiryId)(nil),             // 2: genproto.RespondentInquiryId
	(*GetAllRespondentInquiryResponse)(nil), // 3: genproto.GetAllRespondentInquiryResponse
	(*RespondentRating)(nil),                // 4: genproto.RespondentRating
	(*UpdateRatingRequest)(nil),             // 5: genproto.UpdateRatingRequest
	(*RespondentId)(nil),                    // 6: genproto.RespondentId
	(*GetRespondentsByInquiryRequest)(nil),  // 7: genproto.GetRespondentsByInquiryRequest
	(*GetRespondentsByInquiryResponse)(nil), // 8: genproto.GetRespondentsByInquiryResponse
	(*GetInquiryByRespondentRequest)(nil),   // 9: genproto.GetInquiryByRespondentRequest
}
var file_respondent_inquiry_proto_depIdxs = []int32{
	4, // 0: genproto.RespondentInquiry.rating:type_name -> genproto.RespondentRating
	4, // 1: genproto.CreateRespondentInquiry.rating:type_name -> genproto.RespondentRating
	0, // 2: genproto.GetAllRespondentInquiryResponse.inquiries:type_name -> genproto.RespondentInquiry
	6, // 3: genproto.GetRespondentsByInquiryResponse.ids:type_name -> genproto.RespondentId
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_respondent_inquiry_proto_init() }
func file_respondent_inquiry_proto_init() {
	if File_respondent_inquiry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_respondent_inquiry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespondentInquiry); i {
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
		file_respondent_inquiry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRespondentInquiry); i {
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
		file_respondent_inquiry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespondentInquiryId); i {
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
		file_respondent_inquiry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRespondentInquiryResponse); i {
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
		file_respondent_inquiry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespondentRating); i {
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
		file_respondent_inquiry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRatingRequest); i {
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
		file_respondent_inquiry_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_respondent_inquiry_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRespondentsByInquiryRequest); i {
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
		file_respondent_inquiry_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRespondentsByInquiryResponse); i {
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
		file_respondent_inquiry_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInquiryByRespondentRequest); i {
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
			RawDescriptor: file_respondent_inquiry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_respondent_inquiry_proto_goTypes,
		DependencyIndexes: file_respondent_inquiry_proto_depIdxs,
		MessageInfos:      file_respondent_inquiry_proto_msgTypes,
	}.Build()
	File_respondent_inquiry_proto = out.File
	file_respondent_inquiry_proto_rawDesc = nil
	file_respondent_inquiry_proto_goTypes = nil
	file_respondent_inquiry_proto_depIdxs = nil
}
