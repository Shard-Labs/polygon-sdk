// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: network2/grpc/test/test.proto

package test

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AReq) Reset() {
	*x = AReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network2_grpc_test_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AReq) ProtoMessage() {}

func (x *AReq) ProtoReflect() protoreflect.Message {
	mi := &file_network2_grpc_test_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AReq.ProtoReflect.Descriptor instead.
func (*AReq) Descriptor() ([]byte, []int) {
	return file_network2_grpc_test_test_proto_rawDescGZIP(), []int{0}
}

type AResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AResp) Reset() {
	*x = AResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network2_grpc_test_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AResp) ProtoMessage() {}

func (x *AResp) ProtoReflect() protoreflect.Message {
	mi := &file_network2_grpc_test_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AResp.ProtoReflect.Descriptor instead.
func (*AResp) Descriptor() ([]byte, []int) {
	return file_network2_grpc_test_test_proto_rawDescGZIP(), []int{1}
}

var File_network2_grpc_test_test_proto protoreflect.FileDescriptor

var file_network2_grpc_test_test_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x32, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x22, 0x06, 0x0a, 0x04, 0x41, 0x52, 0x65, 0x71, 0x22, 0x07, 0x0a, 0x05, 0x41,
	0x52, 0x65, 0x73, 0x70, 0x32, 0x20, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x01,
	0x41, 0x12, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x52, 0x65, 0x73, 0x70, 0x42, 0x15, 0x5a, 0x13, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x32, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_network2_grpc_test_test_proto_rawDescOnce sync.Once
	file_network2_grpc_test_test_proto_rawDescData = file_network2_grpc_test_test_proto_rawDesc
)

func file_network2_grpc_test_test_proto_rawDescGZIP() []byte {
	file_network2_grpc_test_test_proto_rawDescOnce.Do(func() {
		file_network2_grpc_test_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_network2_grpc_test_test_proto_rawDescData)
	})
	return file_network2_grpc_test_test_proto_rawDescData
}

var file_network2_grpc_test_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_network2_grpc_test_test_proto_goTypes = []interface{}{
	(*AReq)(nil),  // 0: v1.AReq
	(*AResp)(nil), // 1: v1.AResp
}
var file_network2_grpc_test_test_proto_depIdxs = []int32{
	0, // 0: v1.Test.A:input_type -> v1.AReq
	1, // 1: v1.Test.A:output_type -> v1.AResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_network2_grpc_test_test_proto_init() }
func file_network2_grpc_test_test_proto_init() {
	if File_network2_grpc_test_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_network2_grpc_test_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AReq); i {
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
		file_network2_grpc_test_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AResp); i {
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
			RawDescriptor: file_network2_grpc_test_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_network2_grpc_test_test_proto_goTypes,
		DependencyIndexes: file_network2_grpc_test_test_proto_depIdxs,
		MessageInfos:      file_network2_grpc_test_test_proto_msgTypes,
	}.Build()
	File_network2_grpc_test_test_proto = out.File
	file_network2_grpc_test_test_proto_rawDesc = nil
	file_network2_grpc_test_test_proto_goTypes = nil
	file_network2_grpc_test_test_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	A(ctx context.Context, in *AReq, opts ...grpc.CallOption) (*AResp, error)
}

type testClient struct {
	cc grpc.ClientConnInterface
}

func NewTestClient(cc grpc.ClientConnInterface) TestClient {
	return &testClient{cc}
}

func (c *testClient) A(ctx context.Context, in *AReq, opts ...grpc.CallOption) (*AResp, error) {
	out := new(AResp)
	err := c.cc.Invoke(ctx, "/v1.Test/A", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	A(context.Context, *AReq) (*AResp, error)
}

// UnimplementedTestServer can be embedded to have forward compatible implementations.
type UnimplementedTestServer struct {
}

func (*UnimplementedTestServer) A(context.Context, *AReq) (*AResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method A not implemented")
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_A_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).A(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Test/A",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).A(ctx, req.(*AReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "A",
			Handler:    _Test_A_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "network2/grpc/test/test.proto",
}
