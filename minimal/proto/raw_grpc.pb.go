// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RawClient is the client API for Raw service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RawClient interface {
	RawMsg(ctx context.Context, in *RawMsgReq, opts ...grpc.CallOption) (*empty.Empty, error)
}

type rawClient struct {
	cc grpc.ClientConnInterface
}

func NewRawClient(cc grpc.ClientConnInterface) RawClient {
	return &rawClient{cc}
}

func (c *rawClient) RawMsg(ctx context.Context, in *RawMsgReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.Raw/RawMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RawServer is the server API for Raw service.
// All implementations must embed UnimplementedRawServer
// for forward compatibility
type RawServer interface {
	RawMsg(context.Context, *RawMsgReq) (*empty.Empty, error)
	mustEmbedUnimplementedRawServer()
}

// UnimplementedRawServer must be embedded to have forward compatible implementations.
type UnimplementedRawServer struct {
}

func (UnimplementedRawServer) RawMsg(context.Context, *RawMsgReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RawMsg not implemented")
}
func (UnimplementedRawServer) mustEmbedUnimplementedRawServer() {}

// UnsafeRawServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RawServer will
// result in compilation errors.
type UnsafeRawServer interface {
	mustEmbedUnimplementedRawServer()
}

func RegisterRawServer(s grpc.ServiceRegistrar, srv RawServer) {
	s.RegisterService(&Raw_ServiceDesc, srv)
}

func _Raw_RawMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RawServer).RawMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Raw/RawMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RawServer).RawMsg(ctx, req.(*RawMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Raw_ServiceDesc is the grpc.ServiceDesc for Raw service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Raw_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Raw",
	HandlerType: (*RawServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RawMsg",
			Handler:    _Raw_RawMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "minimal/proto/raw.proto",
}
