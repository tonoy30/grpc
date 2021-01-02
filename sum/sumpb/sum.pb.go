// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: sum/sumpb/sum.proto

package sumpb

import (
	context "context"
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

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int64 `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 int64 `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sum_sumpb_sum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_sum_sumpb_sum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_sum_sumpb_sum_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetNum1() int64 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *Input) GetNum2() int64 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ans int64 `protobuf:"varint,1,opt,name=ans,proto3" json:"ans,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sum_sumpb_sum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_sum_sumpb_sum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_sum_sumpb_sum_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetAns() int64 {
	if x != nil {
		return x.Ans
	}
	return 0
}

var File_sum_sumpb_sum_proto protoreflect.FileDescriptor

var file_sum_sumpb_sum_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x75, 0x6d, 0x2f, 0x73, 0x75, 0x6d, 0x70, 0x62, 0x2f, 0x73, 0x75, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x75, 0x6d, 0x70, 0x62, 0x22, 0x2f, 0x0a, 0x05,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x32, 0x22, 0x1a, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x6e, 0x73, 0x32, 0x32, 0x0a, 0x0a, 0x53, 0x75, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x0c,
	0x2e, 0x73, 0x75, 0x6d, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0d, 0x2e, 0x73,
	0x75, 0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x0b, 0x5a,
	0x09, 0x73, 0x75, 0x6d, 0x2f, 0x73, 0x75, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sum_sumpb_sum_proto_rawDescOnce sync.Once
	file_sum_sumpb_sum_proto_rawDescData = file_sum_sumpb_sum_proto_rawDesc
)

func file_sum_sumpb_sum_proto_rawDescGZIP() []byte {
	file_sum_sumpb_sum_proto_rawDescOnce.Do(func() {
		file_sum_sumpb_sum_proto_rawDescData = protoimpl.X.CompressGZIP(file_sum_sumpb_sum_proto_rawDescData)
	})
	return file_sum_sumpb_sum_proto_rawDescData
}

var file_sum_sumpb_sum_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sum_sumpb_sum_proto_goTypes = []interface{}{
	(*Input)(nil),  // 0: sumpb.Input
	(*Result)(nil), // 1: sumpb.Result
}
var file_sum_sumpb_sum_proto_depIdxs = []int32{
	0, // 0: sumpb.SumService.Sum:input_type -> sumpb.Input
	1, // 1: sumpb.SumService.Sum:output_type -> sumpb.Result
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sum_sumpb_sum_proto_init() }
func file_sum_sumpb_sum_proto_init() {
	if File_sum_sumpb_sum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sum_sumpb_sum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_sum_sumpb_sum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_sum_sumpb_sum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sum_sumpb_sum_proto_goTypes,
		DependencyIndexes: file_sum_sumpb_sum_proto_depIdxs,
		MessageInfos:      file_sum_sumpb_sum_proto_msgTypes,
	}.Build()
	File_sum_sumpb_sum_proto = out.File
	file_sum_sumpb_sum_proto_rawDesc = nil
	file_sum_sumpb_sum_proto_goTypes = nil
	file_sum_sumpb_sum_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumServiceClient interface {
	// unary grpc
	Sum(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Result, error)
}

type sumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSumServiceClient(cc grpc.ClientConnInterface) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) Sum(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/sumpb.SumService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SumServiceServer is the server API for SumService service.
type SumServiceServer interface {
	// unary grpc
	Sum(context.Context, *Input) (*Result, error)
}

// UnimplementedSumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSumServiceServer struct {
}

func (*UnimplementedSumServiceServer) Sum(context.Context, *Input) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sumpb.SumService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Sum(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sumpb.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _SumService_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sum/sumpb/sum.proto",
}
