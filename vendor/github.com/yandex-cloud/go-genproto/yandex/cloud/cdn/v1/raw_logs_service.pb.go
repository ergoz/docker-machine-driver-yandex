// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: yandex/cloud/cdn/v1/raw_logs_service.proto

package cdn

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ActivateRawLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of CDN resource to switch logs storage for..
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// Raw logs settings.
	Settings *RawLogsSettings `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *ActivateRawLogsRequest) Reset() {
	*x = ActivateRawLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateRawLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateRawLogsRequest) ProtoMessage() {}

func (x *ActivateRawLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateRawLogsRequest.ProtoReflect.Descriptor instead.
func (*ActivateRawLogsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescGZIP(), []int{0}
}

func (x *ActivateRawLogsRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *ActivateRawLogsRequest) GetSettings() *RawLogsSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type ActivateRawLogsMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of resource with activated raw logs.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *ActivateRawLogsMetadata) Reset() {
	*x = ActivateRawLogsMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateRawLogsMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateRawLogsMetadata) ProtoMessage() {}

func (x *ActivateRawLogsMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateRawLogsMetadata.ProtoReflect.Descriptor instead.
func (*ActivateRawLogsMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescGZIP(), []int{1}
}

func (x *ActivateRawLogsMetadata) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type ActivateRawLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Raw logs status.
	Status RawLogsStatus `protobuf:"varint,1,opt,name=status,proto3,enum=yandex.cloud.cdn.v1.RawLogsStatus" json:"status,omitempty"`
	// Raw logs settings.
	Settings *RawLogsSettings `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *ActivateRawLogsResponse) Reset() {
	*x = ActivateRawLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateRawLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateRawLogsResponse) ProtoMessage() {}

func (x *ActivateRawLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateRawLogsResponse.ProtoReflect.Descriptor instead.
func (*ActivateRawLogsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescGZIP(), []int{2}
}

func (x *ActivateRawLogsResponse) GetStatus() RawLogsStatus {
	if x != nil {
		return x.Status
	}
	return RawLogsStatus_RAW_LOGS_STATUS_UNSPECIFIED
}

func (x *ActivateRawLogsResponse) GetSettings() *RawLogsSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type DeactivateRawLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of CDN resource to deactivate Raw Logs for.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *DeactivateRawLogsRequest) Reset() {
	*x = DeactivateRawLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeactivateRawLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateRawLogsRequest) ProtoMessage() {}

func (x *DeactivateRawLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateRawLogsRequest.ProtoReflect.Descriptor instead.
func (*DeactivateRawLogsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeactivateRawLogsRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type DeactivateRawLogsMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of CDN resource.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *DeactivateRawLogsMetadata) Reset() {
	*x = DeactivateRawLogsMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeactivateRawLogsMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateRawLogsMetadata) ProtoMessage() {}

func (x *DeactivateRawLogsMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateRawLogsMetadata.ProtoReflect.Descriptor instead.
func (*DeactivateRawLogsMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeactivateRawLogsMetadata) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type GetRawLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of CDN resource to request status and settings.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *GetRawLogsRequest) Reset() {
	*x = GetRawLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRawLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRawLogsRequest) ProtoMessage() {}

func (x *GetRawLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRawLogsRequest.ProtoReflect.Descriptor instead.
func (*GetRawLogsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetRawLogsRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type GetRawLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Raw logs status.
	Status RawLogsStatus `protobuf:"varint,1,opt,name=status,proto3,enum=yandex.cloud.cdn.v1.RawLogsStatus" json:"status,omitempty"`
	// Raw logs settings.
	Settings *RawLogsSettings `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *GetRawLogsResponse) Reset() {
	*x = GetRawLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRawLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRawLogsResponse) ProtoMessage() {}

func (x *GetRawLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRawLogsResponse.ProtoReflect.Descriptor instead.
func (*GetRawLogsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetRawLogsResponse) GetStatus() RawLogsStatus {
	if x != nil {
		return x.Status
	}
	return RawLogsStatus_RAW_LOGS_STATUS_UNSPECIFIED
}

func (x *GetRawLogsResponse) GetSettings() *RawLogsSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type UpdateRawLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of CDN resource.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// Raw logs settings.
	Settings *RawLogsSettings `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *UpdateRawLogsRequest) Reset() {
	*x = UpdateRawLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRawLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRawLogsRequest) ProtoMessage() {}

func (x *UpdateRawLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRawLogsRequest.ProtoReflect.Descriptor instead.
func (*UpdateRawLogsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRawLogsRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *UpdateRawLogsRequest) GetSettings() *RawLogsSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type UpdateRawLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Raw logs status.
	Status RawLogsStatus `protobuf:"varint,1,opt,name=status,proto3,enum=yandex.cloud.cdn.v1.RawLogsStatus" json:"status,omitempty"`
	// Raw logs settings.
	Settings *RawLogsSettings `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *UpdateRawLogsResponse) Reset() {
	*x = UpdateRawLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRawLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRawLogsResponse) ProtoMessage() {}

func (x *UpdateRawLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRawLogsResponse.ProtoReflect.Descriptor instead.
func (*UpdateRawLogsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateRawLogsResponse) GetStatus() RawLogsStatus {
	if x != nil {
		return x.Status
	}
	return RawLogsStatus_RAW_LOGS_STATUS_UNSPECIFIED
}

func (x *UpdateRawLogsResponse) GetSettings() *RawLogsSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type UpdateRawLogsMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of CDN resource.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *UpdateRawLogsMetadata) Reset() {
	*x = UpdateRawLogsMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRawLogsMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRawLogsMetadata) ProtoMessage() {}

func (x *UpdateRawLogsMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRawLogsMetadata.ProtoReflect.Descriptor instead.
func (*UpdateRawLogsMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateRawLogsMetadata) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

var File_yandex_cloud_cdn_v1_raw_logs_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x77, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x77, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a,
	0x16, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7,
	0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x3a, 0x0a, 0x17, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x17, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x22, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a, 0x08,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x49,
	0x0a, 0x18, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x61, 0x77, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x19, 0x44, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x61,
	0x77, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x92, 0x01, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x22, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40,
	0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x22, 0x87, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x77, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x15, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x4c, 0x6f,
	0x67, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x40, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x22, 0x38, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x77, 0x4c,
	0x6f, 0x67, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x32, 0xba, 0x05, 0x0a,
	0x0e, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xb5, 0x01, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x61, 0x77, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x59, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61,
	0x77, 0x4c, 0x6f, 0x67, 0x73, 0x3a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0xb2, 0xd2, 0x2a, 0x32, 0x0a, 0x17, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52,
	0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x17,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xbb, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f,
	0x22, 0x1d, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x77, 0x4c, 0x6f, 0x67,
	0x73, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0xb2,
	0xd2, 0x2a, 0x32, 0x0a, 0x19, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52,
	0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x15,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x7d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x26, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61,
	0x77, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0xb2, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x29, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x77, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x22, 0x32, 0x1d, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0xb2, 0xd2, 0x2a, 0x2e, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x77, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64,
	0x6e, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x64,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescData = file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDesc
)

func file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescData)
	})
	return file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDescData
}

var file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_yandex_cloud_cdn_v1_raw_logs_service_proto_goTypes = []interface{}{
	(*ActivateRawLogsRequest)(nil),    // 0: yandex.cloud.cdn.v1.ActivateRawLogsRequest
	(*ActivateRawLogsMetadata)(nil),   // 1: yandex.cloud.cdn.v1.ActivateRawLogsMetadata
	(*ActivateRawLogsResponse)(nil),   // 2: yandex.cloud.cdn.v1.ActivateRawLogsResponse
	(*DeactivateRawLogsRequest)(nil),  // 3: yandex.cloud.cdn.v1.DeactivateRawLogsRequest
	(*DeactivateRawLogsMetadata)(nil), // 4: yandex.cloud.cdn.v1.DeactivateRawLogsMetadata
	(*GetRawLogsRequest)(nil),         // 5: yandex.cloud.cdn.v1.GetRawLogsRequest
	(*GetRawLogsResponse)(nil),        // 6: yandex.cloud.cdn.v1.GetRawLogsResponse
	(*UpdateRawLogsRequest)(nil),      // 7: yandex.cloud.cdn.v1.UpdateRawLogsRequest
	(*UpdateRawLogsResponse)(nil),     // 8: yandex.cloud.cdn.v1.UpdateRawLogsResponse
	(*UpdateRawLogsMetadata)(nil),     // 9: yandex.cloud.cdn.v1.UpdateRawLogsMetadata
	(*RawLogsSettings)(nil),           // 10: yandex.cloud.cdn.v1.RawLogsSettings
	(RawLogsStatus)(0),                // 11: yandex.cloud.cdn.v1.RawLogsStatus
	(*operation.Operation)(nil),       // 12: yandex.cloud.operation.Operation
}
var file_yandex_cloud_cdn_v1_raw_logs_service_proto_depIdxs = []int32{
	10, // 0: yandex.cloud.cdn.v1.ActivateRawLogsRequest.settings:type_name -> yandex.cloud.cdn.v1.RawLogsSettings
	11, // 1: yandex.cloud.cdn.v1.ActivateRawLogsResponse.status:type_name -> yandex.cloud.cdn.v1.RawLogsStatus
	10, // 2: yandex.cloud.cdn.v1.ActivateRawLogsResponse.settings:type_name -> yandex.cloud.cdn.v1.RawLogsSettings
	11, // 3: yandex.cloud.cdn.v1.GetRawLogsResponse.status:type_name -> yandex.cloud.cdn.v1.RawLogsStatus
	10, // 4: yandex.cloud.cdn.v1.GetRawLogsResponse.settings:type_name -> yandex.cloud.cdn.v1.RawLogsSettings
	10, // 5: yandex.cloud.cdn.v1.UpdateRawLogsRequest.settings:type_name -> yandex.cloud.cdn.v1.RawLogsSettings
	11, // 6: yandex.cloud.cdn.v1.UpdateRawLogsResponse.status:type_name -> yandex.cloud.cdn.v1.RawLogsStatus
	10, // 7: yandex.cloud.cdn.v1.UpdateRawLogsResponse.settings:type_name -> yandex.cloud.cdn.v1.RawLogsSettings
	0,  // 8: yandex.cloud.cdn.v1.RawLogsService.Activate:input_type -> yandex.cloud.cdn.v1.ActivateRawLogsRequest
	3,  // 9: yandex.cloud.cdn.v1.RawLogsService.Deactivate:input_type -> yandex.cloud.cdn.v1.DeactivateRawLogsRequest
	5,  // 10: yandex.cloud.cdn.v1.RawLogsService.Get:input_type -> yandex.cloud.cdn.v1.GetRawLogsRequest
	7,  // 11: yandex.cloud.cdn.v1.RawLogsService.Update:input_type -> yandex.cloud.cdn.v1.UpdateRawLogsRequest
	12, // 12: yandex.cloud.cdn.v1.RawLogsService.Activate:output_type -> yandex.cloud.operation.Operation
	12, // 13: yandex.cloud.cdn.v1.RawLogsService.Deactivate:output_type -> yandex.cloud.operation.Operation
	6,  // 14: yandex.cloud.cdn.v1.RawLogsService.Get:output_type -> yandex.cloud.cdn.v1.GetRawLogsResponse
	12, // 15: yandex.cloud.cdn.v1.RawLogsService.Update:output_type -> yandex.cloud.operation.Operation
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_yandex_cloud_cdn_v1_raw_logs_service_proto_init() }
func file_yandex_cloud_cdn_v1_raw_logs_service_proto_init() {
	if File_yandex_cloud_cdn_v1_raw_logs_service_proto != nil {
		return
	}
	file_yandex_cloud_cdn_v1_raw_logs_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateRawLogsRequest); i {
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
		file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateRawLogsMetadata); i {
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
		file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateRawLogsResponse); i {
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
		file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeactivateRawLogsRequest); i {
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
		file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeactivateRawLogsMetadata); i {
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
		file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRawLogsRequest); i {
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
		file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRawLogsResponse); i {
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
		file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRawLogsRequest); i {
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
		file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRawLogsResponse); i {
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
		file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRawLogsMetadata); i {
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
			RawDescriptor: file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_cdn_v1_raw_logs_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_cdn_v1_raw_logs_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_cdn_v1_raw_logs_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_cdn_v1_raw_logs_service_proto = out.File
	file_yandex_cloud_cdn_v1_raw_logs_service_proto_rawDesc = nil
	file_yandex_cloud_cdn_v1_raw_logs_service_proto_goTypes = nil
	file_yandex_cloud_cdn_v1_raw_logs_service_proto_depIdxs = nil
}
