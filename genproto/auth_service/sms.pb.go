// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: sms.proto

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

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_sms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_sms_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type StorageSms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipients []string `protobuf:"bytes,1,rep,name=recipients,proto3" json:"recipients,omitempty"`
	Text       string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *StorageSms) Reset() {
	*x = StorageSms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageSms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageSms) ProtoMessage() {}

func (x *StorageSms) ProtoReflect() protoreflect.Message {
	mi := &file_sms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageSms.ProtoReflect.Descriptor instead.
func (*StorageSms) Descriptor() ([]byte, []int) {
	return file_sms_proto_rawDescGZIP(), []int{1}
}

func (x *StorageSms) GetRecipients() []string {
	if x != nil {
		return x.Recipients
	}
	return nil
}

func (x *StorageSms) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SMS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Originator string   `protobuf:"bytes,1,opt,name=originator,proto3" json:"originator,omitempty"`
	Content    *Content `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SMS) Reset() {
	*x = SMS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SMS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SMS) ProtoMessage() {}

func (x *SMS) ProtoReflect() protoreflect.Message {
	mi := &file_sms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SMS.ProtoReflect.Descriptor instead.
func (*SMS) Descriptor() ([]byte, []int) {
	return file_sms_proto_rawDescGZIP(), []int{2}
}

func (x *SMS) GetOriginator() string {
	if x != nil {
		return x.Originator
	}
	return ""
}

func (x *SMS) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId     string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Recipient     string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	RouteId       string `protobuf:"bytes,3,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
	RegionCode    string `protobuf:"bytes,4,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	MessageId     string `protobuf:"bytes,5,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Type          string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	ScheduledDate int32  `protobuf:"varint,7,opt,name=scheduled_date,json=scheduledDate,proto3" json:"scheduled_date,omitempty"`
	StoreId       string `protobuf:"bytes,8,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	Sms           *SMS   `protobuf:"bytes,9,opt,name=sms,proto3" json:"sms,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_sms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_sms_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *Message) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *Message) GetRouteId() string {
	if x != nil {
		return x.RouteId
	}
	return ""
}

func (x *Message) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *Message) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Message) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Message) GetScheduledDate() int32 {
	if x != nil {
		return x.ScheduledDate
	}
	return 0
}

func (x *Message) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *Message) GetSms() *SMS {
	if x != nil {
		return x.Sms
	}
	return nil
}

type Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *Body) Reset() {
	*x = Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Body) ProtoMessage() {}

func (x *Body) ProtoReflect() protoreflect.Message {
	mi := &file_sms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Body.ProtoReflect.Descriptor instead.
func (*Body) Descriptor() ([]byte, []int) {
	return file_sms_proto_rawDescGZIP(), []int{4}
}

func (x *Body) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_sms_proto protoreflect.FileDescriptor

var file_sms_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x40, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53,
	0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x52, 0x0a, 0x03, 0x53, 0x4d, 0x53, 0x12, 0x1e, 0x0a,
	0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2b, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x98, 0x02, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x03, 0x73, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x4d, 0x53,
	0x52, 0x03, 0x73, 0x6d, 0x73, 0x22, 0x35, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x2d, 0x0a,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x17, 0x5a, 0x15,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sms_proto_rawDescOnce sync.Once
	file_sms_proto_rawDescData = file_sms_proto_rawDesc
)

func file_sms_proto_rawDescGZIP() []byte {
	file_sms_proto_rawDescOnce.Do(func() {
		file_sms_proto_rawDescData = protoimpl.X.CompressGZIP(file_sms_proto_rawDescData)
	})
	return file_sms_proto_rawDescData
}

var file_sms_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sms_proto_goTypes = []interface{}{
	(*Content)(nil),    // 0: genproto.Content
	(*StorageSms)(nil), // 1: genproto.StorageSms
	(*SMS)(nil),        // 2: genproto.SMS
	(*Message)(nil),    // 3: genproto.Message
	(*Body)(nil),       // 4: genproto.Body
}
var file_sms_proto_depIdxs = []int32{
	0, // 0: genproto.SMS.content:type_name -> genproto.Content
	2, // 1: genproto.Message.sms:type_name -> genproto.SMS
	3, // 2: genproto.Body.messages:type_name -> genproto.Message
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sms_proto_init() }
func file_sms_proto_init() {
	if File_sms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_sms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageSms); i {
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
		file_sms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SMS); i {
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
		file_sms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_sms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Body); i {
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
			RawDescriptor: file_sms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sms_proto_goTypes,
		DependencyIndexes: file_sms_proto_depIdxs,
		MessageInfos:      file_sms_proto_msgTypes,
	}.Build()
	File_sms_proto = out.File
	file_sms_proto_rawDesc = nil
	file_sms_proto_goTypes = nil
	file_sms_proto_depIdxs = nil
}
