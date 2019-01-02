// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/predict.proto

package containers_ai_alameda_v1alpha1_datahub

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RecommendationPolicy int32

const (
	RecommendationPolicy_RECOMMENDATIONPOLICY_UNDEFINED RecommendationPolicy = 0
	RecommendationPolicy_STABLE                         RecommendationPolicy = 1
	RecommendationPolicy_COMPACT                        RecommendationPolicy = 2
)

var RecommendationPolicy_name = map[int32]string{
	0: "RECOMMENDATIONPOLICY_UNDEFINED",
	1: "STABLE",
	2: "COMPACT",
}
var RecommendationPolicy_value = map[string]int32{
	"RECOMMENDATIONPOLICY_UNDEFINED": 0,
	"STABLE":  1,
	"COMPACT": 2,
}

func (x RecommendationPolicy) String() string {
	return proto.EnumName(RecommendationPolicy_name, int32(x))
}
func (RecommendationPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_predict_5ee7a71176176971, []int{0}
}

type ContainerPrediction struct {
	Name                            string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PredictedRawData                []*MetricData `protobuf:"bytes,2,rep,name=predicted_raw_data,json=predictedRawData,proto3" json:"predicted_raw_data,omitempty"`
	PredictedLimitData              []*MetricData `protobuf:"bytes,3,rep,name=predicted_limit_data,json=predictedLimitData,proto3" json:"predicted_limit_data,omitempty"`
	PredictedRequestData            []*MetricData `protobuf:"bytes,4,rep,name=predicted_request_data,json=predictedRequestData,proto3" json:"predicted_request_data,omitempty"`
	PredictedInitialLimitResource   []*MetricData `protobuf:"bytes,5,rep,name=predicted_initial_limit_resource,json=predictedInitialLimitResource,proto3" json:"predicted_initial_limit_resource,omitempty"`
	PredictedInitialRequestResource []*MetricData `protobuf:"bytes,6,rep,name=predicted_initial_request_resource,json=predictedInitialRequestResource,proto3" json:"predicted_initial_request_resource,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}      `json:"-"`
	XXX_unrecognized                []byte        `json:"-"`
	XXX_sizecache                   int32         `json:"-"`
}

func (m *ContainerPrediction) Reset()         { *m = ContainerPrediction{} }
func (m *ContainerPrediction) String() string { return proto.CompactTextString(m) }
func (*ContainerPrediction) ProtoMessage()    {}
func (*ContainerPrediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_predict_5ee7a71176176971, []int{0}
}
func (m *ContainerPrediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerPrediction.Unmarshal(m, b)
}
func (m *ContainerPrediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerPrediction.Marshal(b, m, deterministic)
}
func (dst *ContainerPrediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerPrediction.Merge(dst, src)
}
func (m *ContainerPrediction) XXX_Size() int {
	return xxx_messageInfo_ContainerPrediction.Size(m)
}
func (m *ContainerPrediction) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerPrediction.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerPrediction proto.InternalMessageInfo

func (m *ContainerPrediction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerPrediction) GetPredictedRawData() []*MetricData {
	if m != nil {
		return m.PredictedRawData
	}
	return nil
}

func (m *ContainerPrediction) GetPredictedLimitData() []*MetricData {
	if m != nil {
		return m.PredictedLimitData
	}
	return nil
}

func (m *ContainerPrediction) GetPredictedRequestData() []*MetricData {
	if m != nil {
		return m.PredictedRequestData
	}
	return nil
}

func (m *ContainerPrediction) GetPredictedInitialLimitResource() []*MetricData {
	if m != nil {
		return m.PredictedInitialLimitResource
	}
	return nil
}

func (m *ContainerPrediction) GetPredictedInitialRequestResource() []*MetricData {
	if m != nil {
		return m.PredictedInitialRequestResource
	}
	return nil
}

type PodPrediction struct {
	NamespacedName       *NamespacedName        `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	ContainerPredictions []*ContainerPrediction `protobuf:"bytes,2,rep,name=container_predictions,json=containerPredictions,proto3" json:"container_predictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *PodPrediction) Reset()         { *m = PodPrediction{} }
func (m *PodPrediction) String() string { return proto.CompactTextString(m) }
func (*PodPrediction) ProtoMessage()    {}
func (*PodPrediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_predict_5ee7a71176176971, []int{1}
}
func (m *PodPrediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodPrediction.Unmarshal(m, b)
}
func (m *PodPrediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodPrediction.Marshal(b, m, deterministic)
}
func (dst *PodPrediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodPrediction.Merge(dst, src)
}
func (m *PodPrediction) XXX_Size() int {
	return xxx_messageInfo_PodPrediction.Size(m)
}
func (m *PodPrediction) XXX_DiscardUnknown() {
	xxx_messageInfo_PodPrediction.DiscardUnknown(m)
}

var xxx_messageInfo_PodPrediction proto.InternalMessageInfo

func (m *PodPrediction) GetNamespacedName() *NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *PodPrediction) GetContainerPredictions() []*ContainerPrediction {
	if m != nil {
		return m.ContainerPredictions
	}
	return nil
}

type NodePrediction struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PredictedRawData     []*MetricData `protobuf:"bytes,2,rep,name=predicted_raw_data,json=predictedRawData,proto3" json:"predicted_raw_data,omitempty"`
	IsScheduled          bool          `protobuf:"varint,3,opt,name=is_scheduled,json=isScheduled,proto3" json:"is_scheduled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NodePrediction) Reset()         { *m = NodePrediction{} }
func (m *NodePrediction) String() string { return proto.CompactTextString(m) }
func (*NodePrediction) ProtoMessage()    {}
func (*NodePrediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_predict_5ee7a71176176971, []int{2}
}
func (m *NodePrediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodePrediction.Unmarshal(m, b)
}
func (m *NodePrediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodePrediction.Marshal(b, m, deterministic)
}
func (dst *NodePrediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodePrediction.Merge(dst, src)
}
func (m *NodePrediction) XXX_Size() int {
	return xxx_messageInfo_NodePrediction.Size(m)
}
func (m *NodePrediction) XXX_DiscardUnknown() {
	xxx_messageInfo_NodePrediction.DiscardUnknown(m)
}

var xxx_messageInfo_NodePrediction proto.InternalMessageInfo

func (m *NodePrediction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodePrediction) GetPredictedRawData() []*MetricData {
	if m != nil {
		return m.PredictedRawData
	}
	return nil
}

func (m *NodePrediction) GetIsScheduled() bool {
	if m != nil {
		return m.IsScheduled
	}
	return false
}

func init() {
	proto.RegisterType((*ContainerPrediction)(nil), "containers_ai.alameda.v1alpha1.datahub.ContainerPrediction")
	proto.RegisterType((*PodPrediction)(nil), "containers_ai.alameda.v1alpha1.datahub.PodPrediction")
	proto.RegisterType((*NodePrediction)(nil), "containers_ai.alameda.v1alpha1.datahub.NodePrediction")
	proto.RegisterEnum("containers_ai.alameda.v1alpha1.datahub.RecommendationPolicy", RecommendationPolicy_name, RecommendationPolicy_value)
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/predict.proto", fileDescriptor_predict_5ee7a71176176971)
}

var fileDescriptor_predict_5ee7a71176176971 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xcf, 0x8a, 0x13, 0x41,
	0x10, 0xc6, 0x9d, 0x4d, 0x8c, 0x5a, 0xd1, 0x35, 0xb4, 0x51, 0xc2, 0x82, 0x1a, 0x73, 0x90, 0xb8,
	0xc2, 0x84, 0x8d, 0xe0, 0xc5, 0x53, 0x9c, 0x8c, 0x10, 0x48, 0x26, 0x43, 0x27, 0x1e, 0x3c, 0xb5,
	0xb5, 0xdd, 0x0d, 0x69, 0x98, 0x99, 0x1e, 0xa7, 0x27, 0x2e, 0xe2, 0xc1, 0xe7, 0xf1, 0x91, 0xbc,
	0xfa, 0x24, 0x32, 0x7f, 0xb3, 0xba, 0xa2, 0x21, 0x97, 0xbd, 0xd5, 0x14, 0x5d, 0xdf, 0xef, 0xfb,
	0xa0, 0x6a, 0xe0, 0x14, 0x03, 0x0c, 0xa5, 0x40, 0x86, 0xb1, 0x1a, 0x7d, 0x3e, 0xc3, 0x20, 0xde,
	0xe0, 0xd9, 0x48, 0x60, 0x8a, 0x9b, 0xed, 0xf9, 0x28, 0x4e, 0xa4, 0x50, 0x3c, 0xb5, 0xe3, 0x44,
	0xa7, 0x9a, 0x3c, 0xe7, 0x3a, 0x4a, 0x51, 0x45, 0x32, 0x31, 0x0c, 0x95, 0x5d, 0x4e, 0xda, 0xd5,
	0x94, 0x5d, 0x4e, 0x9d, 0xbc, 0xfc, 0xa7, 0x66, 0x28, 0x53, 0xcc, 0xea, 0x42, 0xf4, 0xe4, 0xc5,
	0xff, 0x1e, 0x27, 0x8a, 0x17, 0x4f, 0x07, 0x3f, 0x9a, 0xf0, 0xc0, 0xa9, 0x2c, 0xf8, 0x85, 0x35,
	0xa5, 0x23, 0x42, 0xa0, 0x19, 0x61, 0x28, 0x7b, 0x56, 0xdf, 0x1a, 0xde, 0xa1, 0x79, 0x4d, 0x3e,
	0x02, 0x29, 0xcd, 0x4b, 0xc1, 0x12, 0xbc, 0x60, 0x99, 0x62, 0xef, 0xa8, 0xdf, 0x18, 0xb6, 0xc7,
	0x63, 0x7b, 0xbf, 0x20, 0xf6, 0x22, 0xa7, 0x4f, 0x31, 0x45, 0xda, 0xa9, 0xd5, 0x28, 0x5e, 0x64,
	0x1d, 0x22, 0xa0, 0xbb, 0x23, 0x04, 0x2a, 0x54, 0x69, 0xc1, 0x68, 0x1c, 0xcc, 0xd8, 0x39, 0x9e,
	0x67, 0x72, 0x39, 0x65, 0x03, 0x8f, 0x2e, 0xe5, 0x90, 0x9f, 0xb6, 0xd2, 0x94, 0x9c, 0xe6, 0xc1,
	0x9c, 0x9d, 0x6f, 0x5a, 0x08, 0xe6, 0xa4, 0xaf, 0xd0, 0xdf, 0x91, 0x54, 0xa4, 0x52, 0x85, 0x41,
	0x99, 0x2b, 0x91, 0x46, 0x6f, 0x13, 0x2e, 0x7b, 0x37, 0x0f, 0x66, 0x3e, 0xae, 0xb5, 0x67, 0x85,
	0x74, 0x1e, 0x91, 0x96, 0xc2, 0xe4, 0x1b, 0x0c, 0xae, 0xc2, 0xab, 0xb8, 0x35, 0xbe, 0x75, 0x30,
	0xfe, 0xe9, 0x9f, 0xf8, 0x32, 0x79, 0x65, 0x60, 0xf0, 0xd3, 0x82, 0x7b, 0xbe, 0x16, 0x97, 0xb6,
	0x8a, 0xc1, 0xfd, 0x6c, 0x93, 0x4c, 0x8c, 0x5c, 0x0a, 0x56, 0x2f, 0x58, 0x7b, 0xfc, 0x7a, 0x5f,
	0xbe, 0x57, 0x8f, 0x67, 0x15, 0x3d, 0x8e, 0x7e, 0xfb, 0x26, 0x31, 0x3c, 0xac, 0x85, 0x58, 0x5c,
	0x83, 0x4d, 0xb9, 0xa5, 0x6f, 0xf6, 0xc5, 0xfc, 0xe5, 0x24, 0x68, 0x97, 0x5f, 0x6d, 0x9a, 0xc1,
	0x77, 0x0b, 0x8e, 0x3d, 0x2d, 0xe4, 0xb5, 0xdf, 0xce, 0x33, 0xb8, 0xab, 0x0c, 0x33, 0x7c, 0x23,
	0xc5, 0x36, 0x90, 0xa2, 0xd7, 0xe8, 0x5b, 0xc3, 0xdb, 0xb4, 0xad, 0xcc, 0xaa, 0x6a, 0x9d, 0xae,
	0xa0, 0x4b, 0x25, 0xd7, 0x61, 0x28, 0x23, 0x81, 0x99, 0x55, 0x5f, 0x07, 0x8a, 0x7f, 0x21, 0x03,
	0x78, 0x42, 0x5d, 0x67, 0xb9, 0x58, 0xb8, 0xde, 0x74, 0xb2, 0x9e, 0x2d, 0x3d, 0x7f, 0x39, 0x9f,
	0x39, 0x1f, 0xd8, 0x7b, 0x6f, 0xea, 0xbe, 0x9b, 0x79, 0xee, 0xb4, 0x73, 0x83, 0x00, 0xb4, 0x56,
	0xeb, 0xc9, 0xdb, 0xb9, 0xdb, 0xb1, 0x48, 0x1b, 0x6e, 0x39, 0xcb, 0x85, 0x3f, 0x71, 0xd6, 0x9d,
	0xa3, 0xf3, 0x56, 0xfe, 0x23, 0x79, 0xf5, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x5a, 0x4c, 0xd7,
	0xf6, 0x04, 0x00, 0x00,
}
