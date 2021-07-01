// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pkg/proto/model/application.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Desc      string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	// Unix time of the last time when the cluster is updated.
	UpdatedAt   int64            `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Components  []*ComponentType `protobuf:"bytes,5,rep,name=components,proto3" json:"components,omitempty"`
	ClusterName string           `protobuf:"bytes,6,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	Events      []*AppEventType  `protobuf:"bytes,7,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_model_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_model_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_pkg_proto_model_application_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Application) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Application) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Application) GetComponents() []*ComponentType {
	if x != nil {
		return x.Components
	}
	return nil
}

func (x *Application) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *Application) GetEvents() []*AppEventType {
	if x != nil {
		return x.Events
	}
	return nil
}

type AppYaml struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yaml string `protobuf:"bytes,1,opt,name=yaml,proto3" json:"yaml,omitempty"`
}

func (x *AppYaml) Reset() {
	*x = AppYaml{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_model_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppYaml) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppYaml) ProtoMessage() {}

func (x *AppYaml) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_model_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppYaml.ProtoReflect.Descriptor instead.
func (*AppYaml) Descriptor() ([]byte, []int) {
	return file_pkg_proto_model_application_proto_rawDescGZIP(), []int{1}
}

func (x *AppYaml) GetYaml() string {
	if x != nil {
		return x.Yaml
	}
	return ""
}

type ComponentType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type       string           `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Namespace  string           `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Workload   string           `protobuf:"bytes,4,opt,name=workload,proto3" json:"workload,omitempty"`
	Desc       string           `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	Phase      string           `protobuf:"bytes,6,opt,name=phase,proto3" json:"phase,omitempty"`
	Health     bool             `protobuf:"varint,7,opt,name=health,proto3" json:"health,omitempty"`
	Properties *structpb.Struct `protobuf:"bytes,8,opt,name=properties,proto3" json:"properties,omitempty"`
	Traits     []*TraitType     `protobuf:"bytes,9,rep,name=traits,proto3" json:"traits,omitempty"`
}

func (x *ComponentType) Reset() {
	*x = ComponentType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_model_application_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentType) ProtoMessage() {}

func (x *ComponentType) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_model_application_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentType.ProtoReflect.Descriptor instead.
func (*ComponentType) Descriptor() ([]byte, []int) {
	return file_pkg_proto_model_application_proto_rawDescGZIP(), []int{2}
}

func (x *ComponentType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ComponentType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ComponentType) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ComponentType) GetWorkload() string {
	if x != nil {
		return x.Workload
	}
	return ""
}

func (x *ComponentType) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ComponentType) GetPhase() string {
	if x != nil {
		return x.Phase
	}
	return ""
}

func (x *ComponentType) GetHealth() bool {
	if x != nil {
		return x.Health
	}
	return false
}

func (x *ComponentType) GetProperties() *structpb.Struct {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *ComponentType) GetTraits() []*TraitType {
	if x != nil {
		return x.Traits
	}
	return nil
}

type TraitType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       string           `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Desc       string           `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Properties *structpb.Struct `protobuf:"bytes,3,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (x *TraitType) Reset() {
	*x = TraitType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_model_application_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraitType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraitType) ProtoMessage() {}

func (x *TraitType) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_model_application_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraitType.ProtoReflect.Descriptor instead.
func (*TraitType) Descriptor() ([]byte, []int) {
	return file_pkg_proto_model_application_proto_rawDescGZIP(), []int{3}
}

func (x *TraitType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TraitType) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *TraitType) GetProperties() *structpb.Struct {
	if x != nil {
		return x.Properties
	}
	return nil
}

type AppEventType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Reason  string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Age     string `protobuf:"bytes,3,opt,name=age,proto3" json:"age,omitempty"`
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AppEventType) Reset() {
	*x = AppEventType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_model_application_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppEventType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppEventType) ProtoMessage() {}

func (x *AppEventType) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_model_application_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppEventType.ProtoReflect.Descriptor instead.
func (*AppEventType) Descriptor() ([]byte, []int) {
	return file_pkg_proto_model_application_proto_rawDescGZIP(), []int{4}
}

func (x *AppEventType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AppEventType) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *AppEventType) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

func (x *AppEventType) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ApplicationListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Applications []*Application `protobuf:"bytes,1,rep,name=applications,proto3" json:"applications,omitempty"`
}

func (x *ApplicationListResponse) Reset() {
	*x = ApplicationListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_model_application_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationListResponse) ProtoMessage() {}

func (x *ApplicationListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_model_application_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationListResponse.ProtoReflect.Descriptor instead.
func (*ApplicationListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_model_application_proto_rawDescGZIP(), []int{5}
}

func (x *ApplicationListResponse) GetApplications() []*Application {
	if x != nil {
		return x.Applications
	}
	return nil
}

type ApplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Application *Application `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
}

func (x *ApplicationResponse) Reset() {
	*x = ApplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_model_application_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationResponse) ProtoMessage() {}

func (x *ApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_model_application_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationResponse.ProtoReflect.Descriptor instead.
func (*ApplicationResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_model_application_proto_rawDescGZIP(), []int{6}
}

func (x *ApplicationResponse) GetApplication() *Application {
	if x != nil {
		return x.Application
	}
	return nil
}

var File_pkg_proto_model_application_proto protoreflect.FileDescriptor

var file_pkg_proto_model_application_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x89, 0x02, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x65, 0x6c,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x1d, 0x0a,
	0x07, 0x41, 0x70, 0x70, 0x59, 0x61, 0x6d, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x61, 0x6d, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x22, 0x9f, 0x02, 0x0a,
	0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x74,
	0x72, 0x61, 0x69, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x65,
	0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61,
	0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x22, 0x6c,
	0x0a, 0x09, 0x54, 0x72, 0x61, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x66, 0x0a, 0x0c,
	0x41, 0x70, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x5a, 0x0a, 0x17, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x54, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76,
	0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x61, 0x6d, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x76, 0x65, 0x6c,
	0x61, 0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_model_application_proto_rawDescOnce sync.Once
	file_pkg_proto_model_application_proto_rawDescData = file_pkg_proto_model_application_proto_rawDesc
)

func file_pkg_proto_model_application_proto_rawDescGZIP() []byte {
	file_pkg_proto_model_application_proto_rawDescOnce.Do(func() {
		file_pkg_proto_model_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_model_application_proto_rawDescData)
	})
	return file_pkg_proto_model_application_proto_rawDescData
}

var file_pkg_proto_model_application_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_proto_model_application_proto_goTypes = []interface{}{
	(*Application)(nil),             // 0: vela.api.model.Application
	(*AppYaml)(nil),                 // 1: vela.api.model.AppYaml
	(*ComponentType)(nil),           // 2: vela.api.model.ComponentType
	(*TraitType)(nil),               // 3: vela.api.model.TraitType
	(*AppEventType)(nil),            // 4: vela.api.model.AppEventType
	(*ApplicationListResponse)(nil), // 5: vela.api.model.ApplicationListResponse
	(*ApplicationResponse)(nil),     // 6: vela.api.model.ApplicationResponse
	(*structpb.Struct)(nil),         // 7: google.protobuf.Struct
}
var file_pkg_proto_model_application_proto_depIdxs = []int32{
	2, // 0: vela.api.model.Application.components:type_name -> vela.api.model.ComponentType
	4, // 1: vela.api.model.Application.events:type_name -> vela.api.model.AppEventType
	7, // 2: vela.api.model.ComponentType.properties:type_name -> google.protobuf.Struct
	3, // 3: vela.api.model.ComponentType.traits:type_name -> vela.api.model.TraitType
	7, // 4: vela.api.model.TraitType.properties:type_name -> google.protobuf.Struct
	0, // 5: vela.api.model.ApplicationListResponse.applications:type_name -> vela.api.model.Application
	0, // 6: vela.api.model.ApplicationResponse.application:type_name -> vela.api.model.Application
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_pkg_proto_model_application_proto_init() }
func file_pkg_proto_model_application_proto_init() {
	if File_pkg_proto_model_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_model_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_pkg_proto_model_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppYaml); i {
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
		file_pkg_proto_model_application_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentType); i {
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
		file_pkg_proto_model_application_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraitType); i {
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
		file_pkg_proto_model_application_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppEventType); i {
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
		file_pkg_proto_model_application_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationListResponse); i {
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
		file_pkg_proto_model_application_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationResponse); i {
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
			RawDescriptor: file_pkg_proto_model_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_model_application_proto_goTypes,
		DependencyIndexes: file_pkg_proto_model_application_proto_depIdxs,
		MessageInfos:      file_pkg_proto_model_application_proto_msgTypes,
	}.Build()
	File_pkg_proto_model_application_proto = out.File
	file_pkg_proto_model_application_proto_rawDesc = nil
	file_pkg_proto_model_application_proto_goTypes = nil
	file_pkg_proto_model_application_proto_depIdxs = nil
}
