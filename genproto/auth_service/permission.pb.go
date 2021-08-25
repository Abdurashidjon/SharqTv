// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: permission.proto

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

type PermissionModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Id       string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PermissionModel) Reset() {
	*x = PermissionModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionModel) ProtoMessage() {}

func (x *PermissionModel) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionModel.ProtoReflect.Descriptor instead.
func (*PermissionModel) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionModel) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *PermissionModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PermissionModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type NestedPermissionModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId string                   `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name     string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Id       string                   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Childs   []*NestedPermissionModel `protobuf:"bytes,4,rep,name=childs,proto3" json:"childs,omitempty"`
}

func (x *NestedPermissionModel) Reset() {
	*x = NestedPermissionModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedPermissionModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedPermissionModel) ProtoMessage() {}

func (x *NestedPermissionModel) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedPermissionModel.ProtoReflect.Descriptor instead.
func (*NestedPermissionModel) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{1}
}

func (x *NestedPermissionModel) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *NestedPermissionModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NestedPermissionModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NestedPermissionModel) GetChilds() []*NestedPermissionModel {
	if x != nil {
		return x.Childs
	}
	return nil
}

type PermissionScopeModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PermissionId string `protobuf:"bytes,1,opt,name=permission_id,json=permissionId,proto3" json:"permission_id,omitempty"`
	Method       string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Url          string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *PermissionScopeModel) Reset() {
	*x = PermissionScopeModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionScopeModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionScopeModel) ProtoMessage() {}

func (x *PermissionScopeModel) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionScopeModel.ProtoReflect.Descriptor instead.
func (*PermissionScopeModel) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{2}
}

func (x *PermissionScopeModel) GetPermissionId() string {
	if x != nil {
		return x.PermissionId
	}
	return ""
}

func (x *PermissionScopeModel) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *PermissionScopeModel) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Scope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Url    string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Scope) Reset() {
	*x = Scope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scope) ProtoMessage() {}

func (x *Scope) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scope.ProtoReflect.Descriptor instead.
func (*Scope) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{3}
}

func (x *Scope) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Scope) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CreatePermissionModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId string   `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name     string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Scopes   []*Scope `protobuf:"bytes,4,rep,name=scopes,proto3" json:"scopes,omitempty"`
}

func (x *CreatePermissionModel) Reset() {
	*x = CreatePermissionModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePermissionModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePermissionModel) ProtoMessage() {}

func (x *CreatePermissionModel) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePermissionModel.ProtoReflect.Descriptor instead.
func (*CreatePermissionModel) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{4}
}

func (x *CreatePermissionModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreatePermissionModel) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *CreatePermissionModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePermissionModel) GetScopes() []*Scope {
	if x != nil {
		return x.Scopes
	}
	return nil
}

type PermissionFullModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId string                  `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name     string                  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Scopes   []*PermissionScopeModel `protobuf:"bytes,4,rep,name=scopes,proto3" json:"scopes,omitempty"`
}

func (x *PermissionFullModel) Reset() {
	*x = PermissionFullModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionFullModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionFullModel) ProtoMessage() {}

func (x *PermissionFullModel) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionFullModel.ProtoReflect.Descriptor instead.
func (*PermissionFullModel) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{5}
}

func (x *PermissionFullModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PermissionFullModel) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *PermissionFullModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PermissionFullModel) GetScopes() []*PermissionScopeModel {
	if x != nil {
		return x.Scopes
	}
	return nil
}

var File_permission_proto protoreflect.FileDescriptor

var file_permission_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0f,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x91, 0x01, 0x0a, 0x15, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x73, 0x22, 0x65, 0x0a, 0x14, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x31, 0x0a, 0x05, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x81,
	0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x13, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x46, 0x75, 0x6c, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x73, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_permission_proto_rawDescOnce sync.Once
	file_permission_proto_rawDescData = file_permission_proto_rawDesc
)

func file_permission_proto_rawDescGZIP() []byte {
	file_permission_proto_rawDescOnce.Do(func() {
		file_permission_proto_rawDescData = protoimpl.X.CompressGZIP(file_permission_proto_rawDescData)
	})
	return file_permission_proto_rawDescData
}

var file_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_permission_proto_goTypes = []interface{}{
	(*PermissionModel)(nil),       // 0: genproto.PermissionModel
	(*NestedPermissionModel)(nil), // 1: genproto.NestedPermissionModel
	(*PermissionScopeModel)(nil),  // 2: genproto.PermissionScopeModel
	(*Scope)(nil),                 // 3: genproto.Scope
	(*CreatePermissionModel)(nil), // 4: genproto.CreatePermissionModel
	(*PermissionFullModel)(nil),   // 5: genproto.PermissionFullModel
}
var file_permission_proto_depIdxs = []int32{
	1, // 0: genproto.NestedPermissionModel.childs:type_name -> genproto.NestedPermissionModel
	3, // 1: genproto.CreatePermissionModel.scopes:type_name -> genproto.Scope
	2, // 2: genproto.PermissionFullModel.scopes:type_name -> genproto.PermissionScopeModel
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_permission_proto_init() }
func file_permission_proto_init() {
	if File_permission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_permission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionModel); i {
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
		file_permission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedPermissionModel); i {
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
		file_permission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionScopeModel); i {
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
		file_permission_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scope); i {
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
		file_permission_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePermissionModel); i {
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
		file_permission_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionFullModel); i {
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
			RawDescriptor: file_permission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_permission_proto_goTypes,
		DependencyIndexes: file_permission_proto_depIdxs,
		MessageInfos:      file_permission_proto_msgTypes,
	}.Build()
	File_permission_proto = out.File
	file_permission_proto_rawDesc = nil
	file_permission_proto_goTypes = nil
	file_permission_proto_depIdxs = nil
}