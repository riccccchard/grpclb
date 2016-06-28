// Code generated by protoc-gen-go.
// source: grpclb_balancer_v1/balancer.proto
// DO NOT EDIT!

/*
Package grpclb_balancer_v1 is a generated protocol buffer package.

It is generated from these files:
	grpclb_balancer_v1/balancer.proto

It has these top-level messages:
	Server
	ServersRequest
	ServersResponse
*/
package grpclb_balancer_v1

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Server struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Score   int64  `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
}

func (m *Server) Reset()                    { *m = Server{} }
func (m *Server) String() string            { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()               {}
func (*Server) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServersRequest struct {
	Target string `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
}

func (m *ServersRequest) Reset()                    { *m = ServersRequest{} }
func (m *ServersRequest) String() string            { return proto.CompactTextString(m) }
func (*ServersRequest) ProtoMessage()               {}
func (*ServersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ServersResponse struct {
	Servers []*Server `protobuf:"bytes,1,rep,name=servers" json:"servers,omitempty"`
}

func (m *ServersResponse) Reset()                    { *m = ServersResponse{} }
func (m *ServersResponse) String() string            { return proto.CompactTextString(m) }
func (*ServersResponse) ProtoMessage()               {}
func (*ServersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServersResponse) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

func init() {
	proto.RegisterType((*Server)(nil), "grpclb.balancer.v1.Server")
	proto.RegisterType((*ServersRequest)(nil), "grpclb.balancer.v1.ServersRequest")
	proto.RegisterType((*ServersResponse)(nil), "grpclb.balancer.v1.ServersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for LoadBalancer service

type LoadBalancerClient interface {
	Servers(ctx context.Context, in *ServersRequest, opts ...grpc.CallOption) (*ServersResponse, error)
}

type loadBalancerClient struct {
	cc *grpc.ClientConn
}

func NewLoadBalancerClient(cc *grpc.ClientConn) LoadBalancerClient {
	return &loadBalancerClient{cc}
}

func (c *loadBalancerClient) Servers(ctx context.Context, in *ServersRequest, opts ...grpc.CallOption) (*ServersResponse, error) {
	out := new(ServersResponse)
	err := grpc.Invoke(ctx, "/grpclb.balancer.v1.LoadBalancer/Servers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoadBalancer service

type LoadBalancerServer interface {
	Servers(context.Context, *ServersRequest) (*ServersResponse, error)
}

func RegisterLoadBalancerServer(s *grpc.Server, srv LoadBalancerServer) {
	s.RegisterService(&_LoadBalancer_serviceDesc, srv)
}

func _LoadBalancer_Servers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServer).Servers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpclb.balancer.v1.LoadBalancer/Servers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServer).Servers(ctx, req.(*ServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadBalancer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpclb.balancer.v1.LoadBalancer",
	HandlerType: (*LoadBalancerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Servers",
			Handler:    _LoadBalancer_Servers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("grpclb_balancer_v1/balancer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x39, 0x0f, 0x13, 0x1c, 0x45, 0x61, 0x10, 0x09, 0x57, 0xe9, 0xda, 0xa4, 0x5a, 0xb9,
	0x68, 0x61, 0x6d, 0x63, 0x63, 0xb5, 0xfe, 0x80, 0xb0, 0x9b, 0x0c, 0x69, 0x42, 0x36, 0xce, 0xac,
	0xf9, 0xfd, 0x06, 0x77, 0x37, 0x8d, 0x78, 0xdd, 0xbc, 0xc7, 0xf7, 0x1e, 0x8f, 0x81, 0x87, 0x81,
	0xe7, 0x6e, 0x74, 0xad, 0xb3, 0xa3, 0x9d, 0x3a, 0xe2, 0x76, 0x39, 0x3e, 0xe5, 0x5b, 0xcf, 0xec,
	0x83, 0x47, 0x8c, 0x88, 0xde, 0xec, 0xe5, 0xa8, 0x5e, 0xa1, 0xf8, 0x24, 0x5e, 0x88, 0xb1, 0x82,
	0xd2, 0xf6, 0x3d, 0x93, 0x48, 0xb5, 0xbb, 0xdf, 0xd5, 0x17, 0x26, 0x4b, 0xbc, 0x85, 0x73, 0xe9,
	0x3c, 0x53, 0x75, 0xb6, 0xfa, 0x7b, 0x13, 0x85, 0xaa, 0xe1, 0x3a, 0x26, 0xc5, 0xd0, 0xd7, 0x37,
	0x49, 0xc0, 0x3b, 0x28, 0x82, 0xe5, 0x81, 0x42, 0x2a, 0x48, 0x4a, 0xbd, 0xc3, 0xcd, 0x46, 0xca,
	0xec, 0x27, 0x21, 0x7c, 0x81, 0x52, 0xa2, 0xb5, 0xb2, 0xfb, 0xfa, 0xb2, 0x39, 0xe8, 0xbf, 0xe3,
	0x74, 0x4c, 0x99, 0x8c, 0x36, 0x0e, 0xae, 0x3e, 0xbc, 0xed, 0xdf, 0x12, 0x82, 0x06, 0xca, 0x54,
	0x8c, 0xea, 0xff, 0x7c, 0xde, 0x77, 0x78, 0x3c, 0xc9, 0xc4, 0x65, 0xae, 0xf8, 0xfd, 0xd5, 0xf3,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x35, 0x68, 0x36, 0x50, 0x01, 0x00, 0x00,
}
