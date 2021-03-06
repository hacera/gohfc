// Code generated by protoc-gen-go.
// source: peer/peer.proto
// DO NOT EDIT!

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PeerID struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PeerID) Reset()                    { *m = PeerID{} }
func (m *PeerID) String() string            { return proto.CompactTextString(m) }
func (*PeerID) ProtoMessage()               {}
func (*PeerID) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

type PeerEndpoint struct {
	Id      *PeerID `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Address string  `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
}

func (m *PeerEndpoint) Reset()                    { *m = PeerEndpoint{} }
func (m *PeerEndpoint) String() string            { return proto.CompactTextString(m) }
func (*PeerEndpoint) ProtoMessage()               {}
func (*PeerEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *PeerEndpoint) GetId() *PeerID {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*PeerID)(nil), "protos.PeerID")
	proto.RegisterType((*PeerEndpoint)(nil), "protos.PeerEndpoint")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Endorser service

type EndorserClient interface {
	ProcessProposal(ctx context.Context, in *SignedProposal, opts ...grpc.CallOption) (*ProposalResponse, error)
}

type endorserClient struct {
	cc *grpc.ClientConn
}

func NewEndorserClient(cc *grpc.ClientConn) EndorserClient {
	return &endorserClient{cc}
}

func (c *endorserClient) ProcessProposal(ctx context.Context, in *SignedProposal, opts ...grpc.CallOption) (*ProposalResponse, error) {
	out := new(ProposalResponse)
	err := grpc.Invoke(ctx, "/protos.Endorser/ProcessProposal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Endorser service

type EndorserServer interface {
	ProcessProposal(context.Context, *SignedProposal) (*ProposalResponse, error)
}

func RegisterEndorserServer(s *grpc.Server, srv EndorserServer) {
	s.RegisterService(&_Endorser_serviceDesc, srv)
}

func _Endorser_ProcessProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndorserServer).ProcessProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Endorser/ProcessProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndorserServer).ProcessProposal(ctx, req.(*SignedProposal))
	}
	return interceptor(ctx, in, info, handler)
}

var _Endorser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Endorser",
	HandlerType: (*EndorserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessProposal",
			Handler:    _Endorser_ProcessProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor6,
}

func init() { proto.RegisterFile("peer/peer.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x6d, 0x90, 0xaa, 0xa3, 0x58, 0x58, 0x41, 0x42, 0x28, 0x22, 0x39, 0x29, 0x42, 0x02,
	0xf5, 0x1b, 0x88, 0x01, 0xbd, 0xc5, 0x78, 0xf3, 0x22, 0x49, 0x66, 0x4c, 0x17, 0xda, 0x9d, 0x65,
	0x26, 0x1e, 0xfc, 0xf6, 0xd2, 0xdd, 0x44, 0xf0, 0xb2, 0x7f, 0xde, 0x6f, 0xe7, 0xed, 0xe3, 0xc1,
	0xca, 0x13, 0x49, 0x79, 0x58, 0x0a, 0x2f, 0x3c, 0xb2, 0x59, 0x86, 0x4d, 0xb3, 0xab, 0x08, 0x84,
	0x3d, 0x6b, 0xbb, 0x8b, 0x30, 0x5b, 0xff, 0x13, 0x3f, 0x85, 0xd4, 0xb3, 0x53, 0x8a, 0x34, 0x5f,
	0xc3, 0xb2, 0x26, 0x92, 0xd7, 0x67, 0x63, 0xe0, 0xd8, 0xb5, 0x7b, 0x4a, 0x17, 0xb7, 0x8b, 0xbb,
	0xb3, 0x26, 0x9c, 0xf3, 0x17, 0xb8, 0x38, 0xd0, 0xca, 0xa1, 0x67, 0xeb, 0x46, 0x73, 0x03, 0x89,
	0xc5, 0xf0, 0xe2, 0x7c, 0x73, 0x19, 0x1d, 0xb4, 0x88, 0xf3, 0x4d, 0x62, 0xd1, 0xa4, 0x70, 0xd2,
	0x22, 0x0a, 0xa9, 0xa6, 0x49, 0xb0, 0x99, 0xaf, 0x9b, 0x37, 0x38, 0xad, 0x1c, 0xb2, 0x28, 0x89,
	0xa9, 0x60, 0x55, 0x0b, 0xf7, 0xa4, 0x5a, 0x4f, 0xa9, 0xcc, 0xf5, 0x6c, 0xf6, 0x6e, 0x07, 0x47,
	0x38, 0xeb, 0x59, 0xfa, 0xf7, 0xc9, 0xa4, 0x34, 0x53, 0xfc, 0xfc, 0xe8, 0xe9, 0xe1, 0xe3, 0x7e,
	0xb0, 0xe3, 0xf6, 0xbb, 0x2b, 0x7a, 0xde, 0x97, 0xdb, 0x1f, 0x4f, 0xb2, 0x23, 0x1c, 0x48, 0xca,
	0xaf, 0xb6, 0x13, 0xdb, 0x97, 0x71, 0x34, 0x14, 0xd5, 0xc5, 0x8a, 0x1e, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x38, 0x9b, 0xb5, 0xcd, 0x3c, 0x01, 0x00, 0x00,
}
