// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node_attestor.proto

/*
Package nodeattestor is a generated protocol buffer package.

It is generated from these files:
	node_attestor.proto

It has these top-level messages:
	FetchAttestationDataRequest
	FetchAttestationDataResponse
*/
package nodeattestor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/spiffe/sri/pkg/common"
import sriplugin "github.com/spiffe/sri/pkg/common/plugin"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Empty from public import github.com/spiffe/sri/pkg/common/common.proto
type Empty common.Empty

func (m *Empty) Reset()         { (*common.Empty)(m).Reset() }
func (m *Empty) String() string { return (*common.Empty)(m).String() }
func (*Empty) ProtoMessage()    {}

// AttestedData from public import github.com/spiffe/sri/pkg/common/common.proto
type AttestedData common.AttestedData

func (m *AttestedData) Reset()          { (*common.AttestedData)(m).Reset() }
func (m *AttestedData) String() string  { return (*common.AttestedData)(m).String() }
func (*AttestedData) ProtoMessage()     {}
func (m *AttestedData) GetType() string { return (*common.AttestedData)(m).GetType() }
func (m *AttestedData) GetData() []byte { return (*common.AttestedData)(m).GetData() }

// Selector from public import github.com/spiffe/sri/pkg/common/common.proto
type Selector common.Selector

func (m *Selector) Reset()           { (*common.Selector)(m).Reset() }
func (m *Selector) String() string   { return (*common.Selector)(m).String() }
func (*Selector) ProtoMessage()      {}
func (m *Selector) GetType() string  { return (*common.Selector)(m).GetType() }
func (m *Selector) GetValue() string { return (*common.Selector)(m).GetValue() }

// Selectors from public import github.com/spiffe/sri/pkg/common/common.proto
type Selectors common.Selectors

func (m *Selectors) Reset()         { (*common.Selectors)(m).Reset() }
func (m *Selectors) String() string { return (*common.Selectors)(m).String() }
func (*Selectors) ProtoMessage()    {}
func (m *Selectors) GetEntries() []*Selector {
	o := (*common.Selectors)(m).GetEntries()
	if o == nil {
		return nil
	}
	s := make([]*Selector, len(o))
	for i, x := range o {
		s[i] = (*Selector)(x)
	}
	return s
}

// RegistrationEntry from public import github.com/spiffe/sri/pkg/common/common.proto
type RegistrationEntry common.RegistrationEntry

func (m *RegistrationEntry) Reset()         { (*common.RegistrationEntry)(m).Reset() }
func (m *RegistrationEntry) String() string { return (*common.RegistrationEntry)(m).String() }
func (*RegistrationEntry) ProtoMessage()    {}
func (m *RegistrationEntry) GetSelectors() []*Selector {
	o := (*common.RegistrationEntry)(m).GetSelectors()
	if o == nil {
		return nil
	}
	s := make([]*Selector, len(o))
	for i, x := range o {
		s[i] = (*Selector)(x)
	}
	return s
}
func (m *RegistrationEntry) GetParentId() string { return (*common.RegistrationEntry)(m).GetParentId() }
func (m *RegistrationEntry) GetSpiffeId() string { return (*common.RegistrationEntry)(m).GetSpiffeId() }
func (m *RegistrationEntry) GetTtl() int32       { return (*common.RegistrationEntry)(m).GetTtl() }
func (m *RegistrationEntry) GetFbSpiffeIds() []string {
	return (*common.RegistrationEntry)(m).GetFbSpiffeIds()
}

// RegistrationEntries from public import github.com/spiffe/sri/pkg/common/common.proto
type RegistrationEntries common.RegistrationEntries

func (m *RegistrationEntries) Reset()         { (*common.RegistrationEntries)(m).Reset() }
func (m *RegistrationEntries) String() string { return (*common.RegistrationEntries)(m).String() }
func (*RegistrationEntries) ProtoMessage()    {}
func (m *RegistrationEntries) GetEntries() []*RegistrationEntry {
	o := (*common.RegistrationEntries)(m).GetEntries()
	if o == nil {
		return nil
	}
	s := make([]*RegistrationEntry, len(o))
	for i, x := range o {
		s[i] = (*RegistrationEntry)(x)
	}
	return s
}

// ConfigureRequest from public import github.com/spiffe/sri/pkg/common/plugin/plugin.proto
type ConfigureRequest sriplugin.ConfigureRequest

func (m *ConfigureRequest) Reset()         { (*sriplugin.ConfigureRequest)(m).Reset() }
func (m *ConfigureRequest) String() string { return (*sriplugin.ConfigureRequest)(m).String() }
func (*ConfigureRequest) ProtoMessage()    {}
func (m *ConfigureRequest) GetConfiguration() string {
	return (*sriplugin.ConfigureRequest)(m).GetConfiguration()
}

// ConfigureResponse from public import github.com/spiffe/sri/pkg/common/plugin/plugin.proto
type ConfigureResponse sriplugin.ConfigureResponse

func (m *ConfigureResponse) Reset()         { (*sriplugin.ConfigureResponse)(m).Reset() }
func (m *ConfigureResponse) String() string { return (*sriplugin.ConfigureResponse)(m).String() }
func (*ConfigureResponse) ProtoMessage()    {}
func (m *ConfigureResponse) GetErrorList() []string {
	return (*sriplugin.ConfigureResponse)(m).GetErrorList()
}

// GetPluginInfoRequest from public import github.com/spiffe/sri/pkg/common/plugin/plugin.proto
type GetPluginInfoRequest sriplugin.GetPluginInfoRequest

func (m *GetPluginInfoRequest) Reset()         { (*sriplugin.GetPluginInfoRequest)(m).Reset() }
func (m *GetPluginInfoRequest) String() string { return (*sriplugin.GetPluginInfoRequest)(m).String() }
func (*GetPluginInfoRequest) ProtoMessage()    {}

// GetPluginInfoResponse from public import github.com/spiffe/sri/pkg/common/plugin/plugin.proto
type GetPluginInfoResponse sriplugin.GetPluginInfoResponse

func (m *GetPluginInfoResponse) Reset()         { (*sriplugin.GetPluginInfoResponse)(m).Reset() }
func (m *GetPluginInfoResponse) String() string { return (*sriplugin.GetPluginInfoResponse)(m).String() }
func (*GetPluginInfoResponse) ProtoMessage()    {}
func (m *GetPluginInfoResponse) GetName() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetName()
}
func (m *GetPluginInfoResponse) GetCategory() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetCategory()
}
func (m *GetPluginInfoResponse) GetType() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetType()
}
func (m *GetPluginInfoResponse) GetDescription() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetDescription()
}
func (m *GetPluginInfoResponse) GetDateCreated() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetDateCreated()
}
func (m *GetPluginInfoResponse) GetLocation() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetLocation()
}
func (m *GetPluginInfoResponse) GetVersion() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetVersion()
}
func (m *GetPluginInfoResponse) GetAuthor() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetAuthor()
}
func (m *GetPluginInfoResponse) GetCompany() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetCompany()
}

// PluginInfoRequest from public import github.com/spiffe/sri/pkg/common/plugin/plugin.proto
type PluginInfoRequest sriplugin.PluginInfoRequest

func (m *PluginInfoRequest) Reset()         { (*sriplugin.PluginInfoRequest)(m).Reset() }
func (m *PluginInfoRequest) String() string { return (*sriplugin.PluginInfoRequest)(m).String() }
func (*PluginInfoRequest) ProtoMessage()    {}

// PluginInfoReply from public import github.com/spiffe/sri/pkg/common/plugin/plugin.proto
type PluginInfoReply sriplugin.PluginInfoReply

func (m *PluginInfoReply) Reset()         { (*sriplugin.PluginInfoReply)(m).Reset() }
func (m *PluginInfoReply) String() string { return (*sriplugin.PluginInfoReply)(m).String() }
func (*PluginInfoReply) ProtoMessage()    {}
func (m *PluginInfoReply) GetPluginInfo() []*GetPluginInfoResponse {
	o := (*sriplugin.PluginInfoReply)(m).GetPluginInfo()
	if o == nil {
		return nil
	}
	s := make([]*GetPluginInfoResponse, len(o))
	for i, x := range o {
		s[i] = (*GetPluginInfoResponse)(x)
	}
	return s
}

// StopRequest from public import github.com/spiffe/sri/pkg/common/plugin/plugin.proto
type StopRequest sriplugin.StopRequest

func (m *StopRequest) Reset()         { (*sriplugin.StopRequest)(m).Reset() }
func (m *StopRequest) String() string { return (*sriplugin.StopRequest)(m).String() }
func (*StopRequest) ProtoMessage()    {}

// StopReply from public import github.com/spiffe/sri/pkg/common/plugin/plugin.proto
type StopReply sriplugin.StopReply

func (m *StopReply) Reset()         { (*sriplugin.StopReply)(m).Reset() }
func (m *StopReply) String() string { return (*sriplugin.StopReply)(m).String() }
func (*StopReply) ProtoMessage()    {}

// * Represents an empty request.
type FetchAttestationDataRequest struct {
}

func (m *FetchAttestationDataRequest) Reset()                    { *m = FetchAttestationDataRequest{} }
func (m *FetchAttestationDataRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchAttestationDataRequest) ProtoMessage()               {}
func (*FetchAttestationDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// * Represents the attested data and base SPIFFE ID.
type FetchAttestationDataResponse struct {
	AttestedData *common.AttestedData `protobuf:"bytes,1,opt,name=attestedData" json:"attestedData,omitempty"`
	SpiffeId     string               `protobuf:"bytes,2,opt,name=spiffeId" json:"spiffeId,omitempty"`
}

func (m *FetchAttestationDataResponse) Reset()                    { *m = FetchAttestationDataResponse{} }
func (m *FetchAttestationDataResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchAttestationDataResponse) ProtoMessage()               {}
func (*FetchAttestationDataResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FetchAttestationDataResponse) GetAttestedData() *common.AttestedData {
	if m != nil {
		return m.AttestedData
	}
	return nil
}

func (m *FetchAttestationDataResponse) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func init() {
	proto.RegisterType((*FetchAttestationDataRequest)(nil), "nodeattestor.FetchAttestationDataRequest")
	proto.RegisterType((*FetchAttestationDataResponse)(nil), "nodeattestor.FetchAttestationDataResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeAttestor service

type NodeAttestorClient interface {
	// / Returns the node attestation data for specific platform and the generated Base SPIFFE ID for CSR formation.
	FetchAttestationData(ctx context.Context, in *FetchAttestationDataRequest, opts ...grpc.CallOption) (*FetchAttestationDataResponse, error)
	// / Applies the plugin configuration and returns configuration errors.
	Configure(ctx context.Context, in *sriplugin.ConfigureRequest, opts ...grpc.CallOption) (*sriplugin.ConfigureResponse, error)
	// / Returns the version and related metadata of the plugin.
	GetPluginInfo(ctx context.Context, in *sriplugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*sriplugin.GetPluginInfoResponse, error)
}

type nodeAttestorClient struct {
	cc *grpc.ClientConn
}

func NewNodeAttestorClient(cc *grpc.ClientConn) NodeAttestorClient {
	return &nodeAttestorClient{cc}
}

func (c *nodeAttestorClient) FetchAttestationData(ctx context.Context, in *FetchAttestationDataRequest, opts ...grpc.CallOption) (*FetchAttestationDataResponse, error) {
	out := new(FetchAttestationDataResponse)
	err := grpc.Invoke(ctx, "/nodeattestor.NodeAttestor/FetchAttestationData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) Configure(ctx context.Context, in *sriplugin.ConfigureRequest, opts ...grpc.CallOption) (*sriplugin.ConfigureResponse, error) {
	out := new(sriplugin.ConfigureResponse)
	err := grpc.Invoke(ctx, "/nodeattestor.NodeAttestor/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) GetPluginInfo(ctx context.Context, in *sriplugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*sriplugin.GetPluginInfoResponse, error) {
	out := new(sriplugin.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/nodeattestor.NodeAttestor/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeAttestor service

type NodeAttestorServer interface {
	// / Returns the node attestation data for specific platform and the generated Base SPIFFE ID for CSR formation.
	FetchAttestationData(context.Context, *FetchAttestationDataRequest) (*FetchAttestationDataResponse, error)
	// / Applies the plugin configuration and returns configuration errors.
	Configure(context.Context, *sriplugin.ConfigureRequest) (*sriplugin.ConfigureResponse, error)
	// / Returns the version and related metadata of the plugin.
	GetPluginInfo(context.Context, *sriplugin.GetPluginInfoRequest) (*sriplugin.GetPluginInfoResponse, error)
}

func RegisterNodeAttestorServer(s *grpc.Server, srv NodeAttestorServer) {
	s.RegisterService(&_NodeAttestor_serviceDesc, srv)
}

func _NodeAttestor_FetchAttestationData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchAttestationDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).FetchAttestationData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeattestor.NodeAttestor/FetchAttestationData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).FetchAttestationData(ctx, req.(*FetchAttestationDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sriplugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeattestor.NodeAttestor/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).Configure(ctx, req.(*sriplugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sriplugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeattestor.NodeAttestor/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, req.(*sriplugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeAttestor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nodeattestor.NodeAttestor",
	HandlerType: (*NodeAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchAttestationData",
			Handler:    _NodeAttestor_FetchAttestationData_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _NodeAttestor_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _NodeAttestor_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node_attestor.proto",
}

func init() { proto.RegisterFile("node_attestor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x4d, 0x4f, 0x83, 0x40,
	0x10, 0x95, 0x1e, 0x8c, 0x5d, 0xf1, 0xb2, 0xf6, 0xd0, 0xd0, 0x1a, 0x49, 0x4f, 0xd5, 0xc4, 0x25,
	0xa9, 0x1e, 0xbc, 0x12, 0x4d, 0x4d, 0x2f, 0x86, 0xf0, 0x07, 0x0c, 0x85, 0x81, 0x6e, 0x94, 0x1d,
	0x64, 0x87, 0x1f, 0xe3, 0xbf, 0x35, 0x65, 0x97, 0x86, 0x26, 0xf8, 0x71, 0x9a, 0xf0, 0xde, 0x9b,
	0xc7, 0x9b, 0xb7, 0xec, 0x52, 0x61, 0x06, 0x6f, 0x09, 0x11, 0x68, 0xc2, 0x5a, 0x54, 0x35, 0x12,
	0x72, 0x77, 0x0f, 0x76, 0x98, 0x77, 0x57, 0x48, 0xda, 0x35, 0x5b, 0x91, 0x62, 0x19, 0xe8, 0x4a,
	0xe6, 0x39, 0x04, 0xba, 0x96, 0x41, 0xf5, 0x5e, 0x04, 0x29, 0x96, 0x25, 0x2a, 0x3b, 0xcc, 0xb2,
	0xf7, 0xf0, 0xa7, 0xbc, 0xfa, 0x68, 0x0a, 0xd9, 0x0d, 0xb3, 0xb5, 0xb8, 0x62, 0xb3, 0x35, 0x50,
	0xba, 0x0b, 0xdb, 0xbf, 0x26, 0x24, 0x51, 0x3d, 0x27, 0x94, 0xc4, 0xf0, 0xd9, 0x80, 0xa6, 0x05,
	0xb1, 0xf9, 0x30, 0xad, 0x2b, 0x54, 0x1a, 0xf8, 0x23, 0x73, 0x4d, 0x5e, 0xc8, 0xf6, 0xf8, 0xd4,
	0xf1, 0x9d, 0xe5, 0xf9, 0x6a, 0x22, 0x6c, 0xb2, 0xb0, 0xc7, 0xc5, 0x47, 0x4a, 0xee, 0xb1, 0x33,
	0x93, 0x72, 0x93, 0x4d, 0x47, 0xbe, 0xb3, 0x1c, 0xc7, 0x87, 0xef, 0xd5, 0xd7, 0x88, 0xb9, 0xaf,
	0x98, 0x41, 0x68, 0xab, 0xe0, 0x25, 0x9b, 0x0c, 0xc5, 0xe0, 0x37, 0xa2, 0xdf, 0x98, 0xf8, 0xe5,
	0x12, 0xef, 0xf6, 0x3f, 0x52, 0x7b, 0xd5, 0x9a, 0x8d, 0x9f, 0x50, 0xe5, 0xb2, 0x68, 0x6a, 0xe0,
	0x33, 0xa1, 0x6b, 0x69, 0x3b, 0x3b, 0xa0, 0x9d, 0xeb, 0x7c, 0x98, 0xb4, 0x3e, 0x31, 0xbb, 0x78,
	0x01, 0x8a, 0x5a, 0x7a, 0xa3, 0x72, 0xe4, 0xd7, 0x3d, 0xf9, 0x11, 0xd3, 0xf9, 0xf9, 0x3f, 0x0b,
	0x8c, 0x67, 0x74, 0x12, 0x39, 0xdb, 0xd3, 0xf6, 0xed, 0xee, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xd4, 0x8d, 0x16, 0xc2, 0x45, 0x02, 0x00, 0x00,
}
