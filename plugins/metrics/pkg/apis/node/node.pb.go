// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v1.0.0
// source: github.com/rancher/opni/plugins/metrics/pkg/apis/node/node.proto

package node

import (
	_ "github.com/rancher/opni/pkg/apis/capability/v1"
	v1beta1 "github.com/rancher/opni/pkg/config/v1beta1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConfigStatus int32

const (
	ConfigStatus_Unknown     ConfigStatus = 0
	ConfigStatus_UpToDate    ConfigStatus = 1
	ConfigStatus_NeedsUpdate ConfigStatus = 2
)

// Enum value maps for ConfigStatus.
var (
	ConfigStatus_name = map[int32]string{
		0: "Unknown",
		1: "UpToDate",
		2: "NeedsUpdate",
	}
	ConfigStatus_value = map[string]int32{
		"Unknown":     0,
		"UpToDate":    1,
		"NeedsUpdate": 2,
	}
)

func (x ConfigStatus) Enum() *ConfigStatus {
	p := new(ConfigStatus)
	*p = x
	return p
}

func (x ConfigStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_enumTypes[0].Descriptor()
}

func (ConfigStatus) Type() protoreflect.EnumType {
	return &file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_enumTypes[0]
}

func (x ConfigStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigStatus.Descriptor instead.
func (ConfigStatus) EnumDescriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDescGZIP(), []int{0}
}

type MetricsCapabilityConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// If enabled is false, conditions may contain a list of relevant status
	// messages describing why the capability is disabled.
	Conditions []string               `protobuf:"bytes,2,rep,name=conditions,proto3" json:"conditions,omitempty"`
	Spec       *MetricsCapabilitySpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *MetricsCapabilityConfig) Reset() {
	*x = MetricsCapabilityConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsCapabilityConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsCapabilityConfig) ProtoMessage() {}

func (x *MetricsCapabilityConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsCapabilityConfig.ProtoReflect.Descriptor instead.
func (*MetricsCapabilityConfig) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDescGZIP(), []int{0}
}

func (x *MetricsCapabilityConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *MetricsCapabilityConfig) GetConditions() []string {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *MetricsCapabilityConfig) GetSpec() *MetricsCapabilitySpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type MetricsCapabilitySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rules      *v1beta1.RulesSpec `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	Prometheus *PrometheusSpec    `protobuf:"bytes,2,opt,name=prometheus,proto3" json:"prometheus,omitempty"` // TODO: add config options for metrics capability here
}

func (x *MetricsCapabilitySpec) Reset() {
	*x = MetricsCapabilitySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsCapabilitySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsCapabilitySpec) ProtoMessage() {}

func (x *MetricsCapabilitySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsCapabilitySpec.ProtoReflect.Descriptor instead.
func (*MetricsCapabilitySpec) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDescGZIP(), []int{1}
}

func (x *MetricsCapabilitySpec) GetRules() *v1beta1.RulesSpec {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *MetricsCapabilitySpec) GetPrometheus() *PrometheusSpec {
	if x != nil {
		return x.Prometheus
	}
	return nil
}

type PrometheusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// default: quay.io/prometheus/prometheus:latest
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// can be one of: ["externalPromOperator"]
	DeploymentStrategy string `protobuf:"bytes,2,opt,name=deploymentStrategy,proto3" json:"deploymentStrategy,omitempty"`
}

func (x *PrometheusSpec) Reset() {
	*x = PrometheusSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrometheusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrometheusSpec) ProtoMessage() {}

func (x *PrometheusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrometheusSpec.ProtoReflect.Descriptor instead.
func (*PrometheusSpec) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDescGZIP(), []int{2}
}

func (x *PrometheusSpec) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *PrometheusSpec) GetDeploymentStrategy() string {
	if x != nil {
		return x.DeploymentStrategy
	}
	return ""
}

type SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentConfig *MetricsCapabilityConfig `protobuf:"bytes,1,opt,name=currentConfig,proto3" json:"currentConfig,omitempty"`
}

func (x *SyncRequest) Reset() {
	*x = SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRequest) ProtoMessage() {}

func (x *SyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRequest.ProtoReflect.Descriptor instead.
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDescGZIP(), []int{3}
}

func (x *SyncRequest) GetCurrentConfig() *MetricsCapabilityConfig {
	if x != nil {
		return x.CurrentConfig
	}
	return nil
}

type SyncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigStatus  ConfigStatus             `protobuf:"varint,1,opt,name=configStatus,proto3,enum=node.metrics.ConfigStatus" json:"configStatus,omitempty"`
	UpdatedConfig *MetricsCapabilityConfig `protobuf:"bytes,2,opt,name=updatedConfig,proto3" json:"updatedConfig,omitempty"`
}

func (x *SyncResponse) Reset() {
	*x = SyncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncResponse) ProtoMessage() {}

func (x *SyncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncResponse.ProtoReflect.Descriptor instead.
func (*SyncResponse) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDescGZIP(), []int{4}
}

func (x *SyncResponse) GetConfigStatus() ConfigStatus {
	if x != nil {
		return x.ConfigStatus
	}
	return ConfigStatus_Unknown
}

func (x *SyncResponse) GetUpdatedConfig() *MetricsCapabilityConfig {
	if x != nil {
		return x.UpdatedConfig
	}
	return nil
}

var File_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65,
	0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72,
	0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01,
	0x0a, 0x17, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x86, 0x01, 0x0a,
	0x15, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2f, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x65,
	0x74, 0x68, 0x65, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x65,
	0x74, 0x68, 0x65, 0x75, 0x73, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x65,
	0x74, 0x68, 0x65, 0x75, 0x73, 0x22, 0x56, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68,
	0x65, 0x75, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a,
	0x12, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x22, 0x5a, 0x0a,
	0x0b, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x0d,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x9b, 0x01, 0x0a, 0x0c, 0x53, 0x79,
	0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4b, 0x0a, 0x0d, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2a, 0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x70, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x65, 0x65, 0x64, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x10, 0x02, 0x32, 0x56, 0x0a, 0x15, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x04,
	0x53, 0x79, 0x6e, 0x63, 0x12, 0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x53,
	0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65,
	0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDescData = file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDesc
)

func file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDescData)
	})
	return file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDescData
}

var file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_goTypes = []interface{}{
	(ConfigStatus)(0),               // 0: node.metrics.ConfigStatus
	(*MetricsCapabilityConfig)(nil), // 1: node.metrics.MetricsCapabilityConfig
	(*MetricsCapabilitySpec)(nil),   // 2: node.metrics.MetricsCapabilitySpec
	(*PrometheusSpec)(nil),          // 3: node.metrics.PrometheusSpec
	(*SyncRequest)(nil),             // 4: node.metrics.SyncRequest
	(*SyncResponse)(nil),            // 5: node.metrics.SyncResponse
	(*v1beta1.RulesSpec)(nil),       // 6: config.v1beta1.RulesSpec
}
var file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_depIdxs = []int32{
	2, // 0: node.metrics.MetricsCapabilityConfig.spec:type_name -> node.metrics.MetricsCapabilitySpec
	6, // 1: node.metrics.MetricsCapabilitySpec.rules:type_name -> config.v1beta1.RulesSpec
	3, // 2: node.metrics.MetricsCapabilitySpec.prometheus:type_name -> node.metrics.PrometheusSpec
	1, // 3: node.metrics.SyncRequest.currentConfig:type_name -> node.metrics.MetricsCapabilityConfig
	0, // 4: node.metrics.SyncResponse.configStatus:type_name -> node.metrics.ConfigStatus
	1, // 5: node.metrics.SyncResponse.updatedConfig:type_name -> node.metrics.MetricsCapabilityConfig
	4, // 6: node.metrics.NodeMetricsCapability.Sync:input_type -> node.metrics.SyncRequest
	5, // 7: node.metrics.NodeMetricsCapability.Sync:output_type -> node.metrics.SyncResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_init() }
func file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_init() {
	if File_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsCapabilityConfig); i {
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
		file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsCapabilitySpec); i {
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
		file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrometheusSpec); i {
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
		file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRequest); i {
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
		file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncResponse); i {
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
			RawDescriptor: file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_depIdxs,
		EnumInfos:         file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_enumTypes,
		MessageInfos:      file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto = out.File
	file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_rawDesc = nil
	file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_goTypes = nil
	file_github_com_rancher_opni_plugins_metrics_pkg_apis_node_node_proto_depIdxs = nil
}
