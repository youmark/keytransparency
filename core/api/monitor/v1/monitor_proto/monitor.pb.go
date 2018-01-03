// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitor/v1/monitor_proto/monitor.proto

/*
Package monitor_proto is a generated protocol buffer package.

Monitor Service

The Key Transparency monitor server service consists of APIs to fetch
monitor results queried using the mutations API.

It is generated from these files:
	monitor/v1/monitor_proto/monitor.proto

It has these top-level messages:
	GetStateRequest
	State
*/
package monitor_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"
import trillian "github.com/google/trillian"

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

// GetStateRequest requests the verification state of a keytransparency domain
// for a particular point in time.
type GetStateRequest struct {
	// kt_url is the URL of the keytransparency server for which the monitoring
	// result will be returned.
	KtUrl string `protobuf:"bytes,2,opt,name=kt_url,json=ktUrl" json:"kt_url,omitempty"`
	// domain_id identifies the merkle tree being monitored.
	DomainId string `protobuf:"bytes,3,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	// epoch specifies the revision for which the monitoring results will
	// be returned (epochs start at 0).
	Epoch int64 `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
}

func (m *GetStateRequest) Reset()                    { *m = GetStateRequest{} }
func (m *GetStateRequest) String() string            { return proto.CompactTextString(m) }
func (*GetStateRequest) ProtoMessage()               {}
func (*GetStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetStateRequest) GetKtUrl() string {
	if m != nil {
		return m.KtUrl
	}
	return ""
}

func (m *GetStateRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *GetStateRequest) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

// State represents the monitor's evaluation of a Key Transparency domain
// at a particular epoch.
type State struct {
	// smr contains the map root for the sparse Merkle Tree signed with the
	// monitor's key on success. If the checks were not successful the
	// smr will be empty. The epochs are encoded into the smr map_revision.
	Smr *trillian.SignedMapRoot `protobuf:"bytes,1,opt,name=smr" json:"smr,omitempty"`
	// seen_time contains the time when this particular signed map root was
	// retrieved and processed.
	SeenTime *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=seen_time,json=seenTime" json:"seen_time,omitempty"`
	// errors contains a list of errors representing the verification checks
	// that failed while monitoring the key-transparency server.
	Errors []*google_rpc.Status `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
}

func (m *State) Reset()                    { *m = State{} }
func (m *State) String() string            { return proto.CompactTextString(m) }
func (*State) ProtoMessage()               {}
func (*State) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *State) GetSmr() *trillian.SignedMapRoot {
	if m != nil {
		return m.Smr
	}
	return nil
}

func (m *State) GetSeenTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.SeenTime
	}
	return nil
}

func (m *State) GetErrors() []*google_rpc.Status {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*GetStateRequest)(nil), "google.keytransparency.monitor.v1.GetStateRequest")
	proto.RegisterType((*State)(nil), "google.keytransparency.monitor.v1.State")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Monitor service

type MonitorClient interface {
	// GetSignedMapRoot returns the latest valid signed map root the monitor
	// observed. Additionally, the response contains extra data necessary to
	// reproduce errors on failure.
	//
	// Returns the signed map root for the latest epoch the monitor observed. If
	// the monitor could not reconstruct the map root given the set of mutations
	// from the previous to the current epoch it won't sign the map root and
	// additional data will be provided to reproduce the failure.
	GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*State, error)
	// GetSignedMapRootByRevision returns the monitor's result for a specific map
	// revision.
	//
	// Returns the signed map root for the specified epoch the monitor observed.
	// If the monitor could not reconstruct the map root given the set of
	// mutations from the previous to the current epoch it won't sign the map root
	// and additional data will be provided to reproduce the failure.
	GetStateByRevision(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*State, error)
}

type monitorClient struct {
	cc *grpc.ClientConn
}

func NewMonitorClient(cc *grpc.ClientConn) MonitorClient {
	return &monitorClient{cc}
}

func (c *monitorClient) GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := grpc.Invoke(ctx, "/google.keytransparency.monitor.v1.Monitor/GetState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorClient) GetStateByRevision(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := grpc.Invoke(ctx, "/google.keytransparency.monitor.v1.Monitor/GetStateByRevision", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Monitor service

type MonitorServer interface {
	// GetSignedMapRoot returns the latest valid signed map root the monitor
	// observed. Additionally, the response contains extra data necessary to
	// reproduce errors on failure.
	//
	// Returns the signed map root for the latest epoch the monitor observed. If
	// the monitor could not reconstruct the map root given the set of mutations
	// from the previous to the current epoch it won't sign the map root and
	// additional data will be provided to reproduce the failure.
	GetState(context.Context, *GetStateRequest) (*State, error)
	// GetSignedMapRootByRevision returns the monitor's result for a specific map
	// revision.
	//
	// Returns the signed map root for the specified epoch the monitor observed.
	// If the monitor could not reconstruct the map root given the set of
	// mutations from the previous to the current epoch it won't sign the map root
	// and additional data will be provided to reproduce the failure.
	GetStateByRevision(context.Context, *GetStateRequest) (*State, error)
}

func RegisterMonitorServer(s *grpc.Server, srv MonitorServer) {
	s.RegisterService(&_Monitor_serviceDesc, srv)
}

func _Monitor_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.monitor.v1.Monitor/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).GetState(ctx, req.(*GetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitor_GetStateByRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).GetStateByRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.monitor.v1.Monitor/GetStateByRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).GetStateByRevision(ctx, req.(*GetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Monitor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.keytransparency.monitor.v1.Monitor",
	HandlerType: (*MonitorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetState",
			Handler:    _Monitor_GetState_Handler,
		},
		{
			MethodName: "GetStateByRevision",
			Handler:    _Monitor_GetStateByRevision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitor/v1/monitor_proto/monitor.proto",
}

func init() { proto.RegisterFile("monitor/v1/monitor_proto/monitor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0x26, 0x0d, 0xad, 0xed, 0x14, 0x14, 0x06, 0xe5, 0x96, 0x2a, 0x58, 0xbb, 0x90, 0xea, 0x62,
	0x86, 0x1b, 0x17, 0x82, 0x0b, 0x7f, 0xae, 0xe8, 0xc5, 0xc5, 0xdd, 0xa4, 0xba, 0xd1, 0x45, 0x99,
	0xa6, 0xc7, 0xde, 0xa1, 0xc9, 0x4c, 0x3c, 0x73, 0x52, 0x28, 0xa5, 0x1b, 0x5f, 0xc1, 0x85, 0xaf,
	0xe1, 0xde, 0xc7, 0xf0, 0x15, 0x7c, 0x0a, 0x57, 0x92, 0x99, 0xe4, 0x22, 0x17, 0x44, 0x11, 0xdc,
	0x24, 0xf9, 0x72, 0xbe, 0x33, 0xf9, 0xce, 0x77, 0xbe, 0xb0, 0xbb, 0x85, 0x35, 0x9a, 0x2c, 0xca,
	0xed, 0xb1, 0x6c, 0x1e, 0x17, 0x25, 0x5a, 0xb2, 0x2d, 0x12, 0x1e, 0xf1, 0x3b, 0x6b, 0x6b, 0xd7,
	0x39, 0x88, 0x0d, 0xec, 0x08, 0x95, 0x71, 0xa5, 0x42, 0x30, 0xd9, 0x4e, 0xb4, 0xac, 0xed, 0xf1,
	0xf8, 0x56, 0xa0, 0x48, 0x55, 0x6a, 0xa9, 0x8c, 0xb1, 0xa4, 0x48, 0x5b, 0xe3, 0xc2, 0x01, 0xe3,
	0xdb, 0x4d, 0xd5, 0xa3, 0x65, 0xf5, 0x5e, 0x92, 0x2e, 0xc0, 0x91, 0x2a, 0xca, 0x86, 0x70, 0xd4,
	0x10, 0xb0, 0xcc, 0xa4, 0x23, 0x45, 0x55, 0xdb, 0x79, 0x95, 0x50, 0xe7, 0xb9, 0x56, 0x26, 0xe0,
	0xe9, 0x3b, 0x76, 0xed, 0x14, 0x68, 0x4e, 0x8a, 0x20, 0x85, 0x0f, 0x15, 0x38, 0xe2, 0x37, 0x58,
	0x6f, 0x43, 0x8b, 0x0a, 0xf3, 0x51, 0x67, 0x12, 0xcd, 0x06, 0x69, 0x77, 0x43, 0x6f, 0x30, 0xe7,
	0x37, 0xd9, 0x60, 0x65, 0x0b, 0xa5, 0xcd, 0x42, 0xaf, 0x46, 0xb1, 0xaf, 0xf4, 0xc3, 0x8b, 0x57,
	0x2b, 0x7e, 0x9d, 0x75, 0xa1, 0xb4, 0xd9, 0xf9, 0x28, 0x9a, 0x44, 0xb3, 0x38, 0x0d, 0x60, 0xfa,
	0x39, 0x62, 0x5d, 0x7f, 0x34, 0xbf, 0xc7, 0x62, 0x57, 0xa0, 0xaf, 0x0e, 0x93, 0x23, 0x71, 0x21,
	0x62, 0xae, 0xd7, 0x06, 0x56, 0x67, 0xaa, 0x4c, 0xad, 0xa5, 0xb4, 0xe6, 0xf0, 0x87, 0x6c, 0xe0,
	0x00, 0xcc, 0xa2, 0x1e, 0xc9, 0x2b, 0x18, 0x26, 0x63, 0xd1, 0x18, 0xd6, 0xce, 0x2b, 0x5e, 0xb7,
	0xf3, 0xa6, 0xfd, 0x9a, 0x5c, 0x43, 0x7e, 0x9f, 0xf5, 0x00, 0xd1, 0xa2, 0x1b, 0xc5, 0x93, 0x78,
	0x36, 0x4c, 0x78, 0xdb, 0x85, 0x65, 0x26, 0xe6, 0xde, 0x84, 0xb4, 0x61, 0x24, 0x3f, 0x3a, 0xec,
	0xca, 0x59, 0x70, 0x9b, 0x7f, 0x89, 0x58, 0xbf, 0xf5, 0x80, 0x27, 0xe2, 0x8f, 0xbb, 0x11, 0x97,
	0x0c, 0x1b, 0xcf, 0xfe, 0xa2, 0xc7, 0x37, 0x4c, 0x5f, 0x7e, 0xfc, 0xf6, 0xfd, 0x53, 0xe7, 0x29,
	0x7f, 0x2c, 0x7f, 0x49, 0x8a, 0x03, 0xdc, 0x02, 0x3a, 0xb9, 0x0f, 0xae, 0x1f, 0x64, 0x70, 0xd5,
	0xc9, 0xfd, 0x85, 0xdf, 0x07, 0xbf, 0x44, 0x70, 0x8f, 0xf2, 0xfa, 0x4a, 0xfc, 0x6b, 0xc4, 0x78,
	0xab, 0xe2, 0x64, 0x97, 0xc2, 0x56, 0x3b, 0x6d, 0xcd, 0x7f, 0x16, 0x7f, 0xea, 0xc5, 0x3f, 0xe3,
	0x4f, 0xfe, 0x51, 0xbc, 0xdc, 0xfb, 0x54, 0x1c, 0x4e, 0x5e, 0xbc, 0x7d, 0xbe, 0xd6, 0x74, 0x5e,
	0x2d, 0x45, 0x66, 0x0b, 0xd9, 0x24, 0xf5, 0xd2, 0xe7, 0x65, 0x66, 0x31, 0xa4, 0xff, 0x77, 0xff,
	0xd4, 0xb2, 0xe7, 0x6f, 0x0f, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xa9, 0x74, 0x52, 0x76,
	0x03, 0x00, 0x00,
}