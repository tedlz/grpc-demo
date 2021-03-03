// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: client_stream.proto

package proto

import (
	"context"
	"reflect"
	"sync"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
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

// 定义请求信息
type SimpleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SimpleRequest) Reset() {
	*x = SimpleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleRequest) ProtoMessage() {}

func (x *SimpleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleRequest.ProtoReflect.Descriptor instead.
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return file_client_stream_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// 定义流式请求信息
type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamData string `protobuf:"bytes,1,opt,name=stream_data,json=streamData,proto3" json:"stream_data,omitempty"`
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_client_stream_proto_rawDescGZIP(), []int{1}
}

func (x *StreamRequest) GetStreamData() string {
	if x != nil {
		return x.StreamData
	}
	return ""
}

// 定义响应信息
type SimpleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SimpleResponse) Reset() {
	*x = SimpleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleResponse) ProtoMessage() {}

func (x *SimpleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleResponse.ProtoReflect.Descriptor instead.
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return file_client_stream_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SimpleResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_client_stream_proto protoreflect.FileDescriptor

var file_client_stream_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0d,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x30, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x0e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32,
	0x84, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x36, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_stream_proto_rawDescOnce sync.Once
	file_client_stream_proto_rawDescData = file_client_stream_proto_rawDesc
)

func file_client_stream_proto_rawDescGZIP() []byte {
	file_client_stream_proto_rawDescOnce.Do(func() {
		file_client_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_stream_proto_rawDescData)
	})
	return file_client_stream_proto_rawDescData
}

var file_client_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_client_stream_proto_goTypes = []interface{}{
	(*SimpleRequest)(nil),  // 0: proto.SimpleRequest
	(*StreamRequest)(nil),  // 1: proto.StreamRequest
	(*SimpleResponse)(nil), // 2: proto.SimpleResponse
}
var file_client_stream_proto_depIdxs = []int32{
	0, // 0: proto.StreamClient.Route:input_type -> proto.SimpleRequest
	1, // 1: proto.StreamClient.RouteList:input_type -> proto.StreamRequest
	2, // 2: proto.StreamClient.Route:output_type -> proto.SimpleResponse
	2, // 3: proto.StreamClient.RouteList:output_type -> proto.SimpleResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_client_stream_proto_init() }
func file_client_stream_proto_init() {
	if File_client_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleRequest); i {
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
		file_client_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRequest); i {
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
		file_client_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleResponse); i {
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
			RawDescriptor: file_client_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_stream_proto_goTypes,
		DependencyIndexes: file_client_stream_proto_depIdxs,
		MessageInfos:      file_client_stream_proto_msgTypes,
	}.Build()
	File_client_stream_proto = out.File
	file_client_stream_proto_rawDesc = nil
	file_client_stream_proto_goTypes = nil
	file_client_stream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamClientClient is the client API for StreamClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamClientClient interface {
	Route(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// 客户端流式 RPC
	RouteList(ctx context.Context, opts ...grpc.CallOption) (StreamClient_RouteListClient, error)
}

type streamClientClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamClientClient(cc grpc.ClientConnInterface) StreamClientClient {
	return &streamClientClient{cc}
}

func (c *streamClientClient) Route(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/proto.StreamClient/Route", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamClientClient) RouteList(ctx context.Context, opts ...grpc.CallOption) (StreamClient_RouteListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamClient_serviceDesc.Streams[0], "/proto.StreamClient/RouteList", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamClientRouteListClient{stream}
	return x, nil
}

type StreamClient_RouteListClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*SimpleResponse, error)
	grpc.ClientStream
}

type streamClientRouteListClient struct {
	grpc.ClientStream
}

func (x *streamClientRouteListClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamClientRouteListClient) CloseAndRecv() (*SimpleResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamClientServer is the server API for StreamClient service.
type StreamClientServer interface {
	Route(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// 客户端流式 RPC
	RouteList(StreamClient_RouteListServer) error
}

// UnimplementedStreamClientServer can be embedded to have forward compatible implementations.
type UnimplementedStreamClientServer struct {
}

func (*UnimplementedStreamClientServer) Route(context.Context, *SimpleRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Route not implemented")
}
func (*UnimplementedStreamClientServer) RouteList(StreamClient_RouteListServer) error {
	return status.Errorf(codes.Unimplemented, "method RouteList not implemented")
}

func RegisterStreamClientServer(s *grpc.Server, srv StreamClientServer) {
	s.RegisterService(&_StreamClient_serviceDesc, srv)
}

func _StreamClient_Route_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamClientServer).Route(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StreamClient/Route",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamClientServer).Route(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamClient_RouteList_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamClientServer).RouteList(&streamClientRouteListServer{stream})
}

type StreamClient_RouteListServer interface {
	SendAndClose(*SimpleResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamClientRouteListServer struct {
	grpc.ServerStream
}

func (x *streamClientRouteListServer) SendAndClose(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamClientRouteListServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StreamClient",
	HandlerType: (*StreamClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Route",
			Handler:    _StreamClient_Route_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RouteList",
			Handler:       _StreamClient_RouteList_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "client_stream.proto",
}
