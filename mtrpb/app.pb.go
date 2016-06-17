// Code generated by protoc-gen-go.
// source: app.proto
// DO NOT EDIT!

/*
Package mtrpb is a generated protocol buffer package.

It is generated from these files:
	app.proto
	data.proto
	field.proto
	tag.proto

It has these top-level messages:
	AppIDSummary
	AppIDSummaryResult
	DataLatencySummary
	DataLatencySummaryResult
	DataSite
	DataSiteResult
	DataLatencyTag
	DataLatencyTagResult
	DataLatencyThreshold
	DataLatencyThresholdResult
	DataType
	DataTypeResult
	DataCompleteness
	DataCompletenessResult
	DataCompletenessTag
	DataCompletenessTagResult
	FieldMetricSummary
	FieldMetricSummaryResult
	FieldMetricTag
	FieldMetricTagResult
	FieldMetricThreshold
	FieldMetricThresholdResult
	FieldModel
	FieldModelResult
	FieldDevice
	FieldDeviceResult
	FieldType
	FieldTypeResult
	FieldState
	FieldStateResult
	FieldStateTag
	FieldStateTagResult
	Tag
	TagResult
	TagSearchResult
*/
package mtrpb

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

type AppIDSummary struct {
	// The applicationid for the metric e.g., mtr-api
	ApplicationID string `protobuf:"bytes,1,opt,name=application_iD" json:"application_iD,omitempty"`
}

func (m *AppIDSummary) Reset()                    { *m = AppIDSummary{} }
func (m *AppIDSummary) String() string            { return proto.CompactTextString(m) }
func (*AppIDSummary) ProtoMessage()               {}
func (*AppIDSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AppIDSummaryResult struct {
	Result []*AppIDSummary `protobuf:"bytes,1,rep,name=result" json:"result,omitempty"`
}

func (m *AppIDSummaryResult) Reset()                    { *m = AppIDSummaryResult{} }
func (m *AppIDSummaryResult) String() string            { return proto.CompactTextString(m) }
func (*AppIDSummaryResult) ProtoMessage()               {}
func (*AppIDSummaryResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AppIDSummaryResult) GetResult() []*AppIDSummary {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*AppIDSummary)(nil), "mtrpb.AppIDSummary")
	proto.RegisterType((*AppIDSummaryResult)(nil), "mtrpb.AppIDSummaryResult")
}

var fileDescriptor0 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0x28, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0x2d, 0x29, 0x2a, 0x48, 0x52, 0x52, 0xe3, 0xe2,
	0x71, 0x2c, 0x28, 0xf0, 0x74, 0x09, 0x2e, 0xcd, 0xcd, 0x4d, 0x2c, 0xaa, 0x14, 0x12, 0xe3, 0xe2,
	0x03, 0xaa, 0xc9, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x8b, 0xcf, 0x74, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x54, 0xb2, 0xe4, 0x12, 0x42, 0x56, 0x17, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x22,
	0xa4, 0xcc, 0xc5, 0x56, 0x04, 0x66, 0x01, 0x55, 0x31, 0x6b, 0x70, 0x1b, 0x09, 0xeb, 0x81, 0x4d,
	0xd5, 0x43, 0x56, 0xea, 0xc4, 0x1e, 0x05, 0xb1, 0x2b, 0x89, 0x0d, 0x6c, 0xb3, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xdd, 0x72, 0x74, 0x7b, 0x86, 0x00, 0x00, 0x00,
}
