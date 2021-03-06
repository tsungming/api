// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/pod_assignment.proto

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

type NodePriority struct {
	Nodes                []string `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodePriority) Reset()         { *m = NodePriority{} }
func (m *NodePriority) String() string { return proto.CompactTextString(m) }
func (*NodePriority) ProtoMessage()    {}
func (*NodePriority) Descriptor() ([]byte, []int) {
	return fileDescriptor_pod_assignment_2569972705d8d0c8, []int{0}
}
func (m *NodePriority) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodePriority.Unmarshal(m, b)
}
func (m *NodePriority) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodePriority.Marshal(b, m, deterministic)
}
func (dst *NodePriority) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodePriority.Merge(dst, src)
}
func (m *NodePriority) XXX_Size() int {
	return xxx_messageInfo_NodePriority.Size(m)
}
func (m *NodePriority) XXX_DiscardUnknown() {
	xxx_messageInfo_NodePriority.DiscardUnknown(m)
}

var xxx_messageInfo_NodePriority proto.InternalMessageInfo

func (m *NodePriority) GetNodes() []string {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Selector struct {
	Selector             map[string]string `protobuf:"bytes,1,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Selector) Reset()         { *m = Selector{} }
func (m *Selector) String() string { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()    {}
func (*Selector) Descriptor() ([]byte, []int) {
	return fileDescriptor_pod_assignment_2569972705d8d0c8, []int{1}
}
func (m *Selector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selector.Unmarshal(m, b)
}
func (m *Selector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selector.Marshal(b, m, deterministic)
}
func (dst *Selector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selector.Merge(dst, src)
}
func (m *Selector) XXX_Size() int {
	return xxx_messageInfo_Selector.Size(m)
}
func (m *Selector) XXX_DiscardUnknown() {
	xxx_messageInfo_Selector.DiscardUnknown(m)
}

var xxx_messageInfo_Selector proto.InternalMessageInfo

func (m *Selector) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

func init() {
	proto.RegisterType((*NodePriority)(nil), "containers_ai.alameda.v1alpha1.datahub.NodePriority")
	proto.RegisterType((*Selector)(nil), "containers_ai.alameda.v1alpha1.datahub.Selector")
	proto.RegisterMapType((map[string]string)(nil), "containers_ai.alameda.v1alpha1.datahub.Selector.SelectorEntry")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/pod_assignment.proto", fileDescriptor_pod_assignment_2569972705d8d0c8)
}

var fileDescriptor_pod_assignment_2569972705d8d0c8 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcc, 0x49, 0xcc,
	0x4d, 0x4d, 0x49, 0x8c, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34,
	0xd4, 0x4f, 0x49, 0x2c, 0x49, 0xcc, 0x28, 0x4d, 0xd2, 0x2f, 0xc8, 0x4f, 0x89, 0x4f, 0x2c, 0x2e,
	0xce, 0x4c, 0xcf, 0xcb, 0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x4b,
	0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d, 0x2a, 0x8e, 0x4f, 0xcc, 0xd4, 0x83, 0x1a, 0xa0,
	0x07, 0xd3, 0xac, 0x07, 0xd5, 0xac, 0xa4, 0xc2, 0xc5, 0xe3, 0x97, 0x9f, 0x92, 0x1a, 0x50, 0x94,
	0x99, 0x5f, 0x94, 0x59, 0x52, 0x29, 0x24, 0xc2, 0xc5, 0x9a, 0x97, 0x9f, 0x92, 0x5a, 0x2c, 0xc1,
	0xa8, 0xc0, 0xac, 0xc1, 0x19, 0x04, 0xe1, 0x28, 0x2d, 0x66, 0xe4, 0xe2, 0x08, 0x4e, 0xcd, 0x49,
	0x4d, 0x2e, 0xc9, 0x2f, 0x12, 0x8a, 0xe2, 0xe2, 0x28, 0x86, 0xb2, 0xc1, 0xaa, 0xb8, 0x8d, 0xec,
	0xf4, 0x88, 0xb3, 0x4d, 0x0f, 0x66, 0x06, 0x9c, 0xe1, 0x9a, 0x57, 0x52, 0x54, 0x19, 0x04, 0x37,
	0x4f, 0xca, 0x9a, 0x8b, 0x17, 0x45, 0x4a, 0x48, 0x80, 0x8b, 0x39, 0x3b, 0xb5, 0x52, 0x82, 0x51,
	0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x04, 0xb9, 0xb0, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x09,
	0x2c, 0x06, 0xe1, 0x58, 0x31, 0x59, 0x30, 0x26, 0xb1, 0x81, 0xbd, 0x6e, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x38, 0x31, 0x4f, 0xa3, 0x2f, 0x01, 0x00, 0x00,
}
