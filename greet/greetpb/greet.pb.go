// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: greet/greetpb/greet.proto

package greetpb

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

type Greet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Greet) Reset() {
	*x = Greet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_greet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Greet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Greet) ProtoMessage() {}

func (x *Greet) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_greet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Greet.ProtoReflect.Descriptor instead.
func (*Greet) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_greet_proto_rawDescGZIP(), []int{0}
}

func (x *Greet) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_greet_greetpb_greet_proto protoreflect.FileDescriptor

var file_greet_greetpb_greet_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2f,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x70, 0x62, 0x22, 0x21, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x3c, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greet_greetpb_greet_proto_rawDescOnce sync.Once
	file_greet_greetpb_greet_proto_rawDescData = file_greet_greetpb_greet_proto_rawDesc
)

func file_greet_greetpb_greet_proto_rawDescGZIP() []byte {
	file_greet_greetpb_greet_proto_rawDescOnce.Do(func() {
		file_greet_greetpb_greet_proto_rawDescData = protoimpl.X.CompressGZIP(file_greet_greetpb_greet_proto_rawDescData)
	})
	return file_greet_greetpb_greet_proto_rawDescData
}

var file_greet_greetpb_greet_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_greet_greetpb_greet_proto_goTypes = []interface{}{
	(*Greet)(nil), // 0: greetpb.Greet
}
var file_greet_greetpb_greet_proto_depIdxs = []int32{
	0, // 0: greetpb.GreetService.Greeting:input_type -> greetpb.Greet
	0, // 1: greetpb.GreetService.Greeting:output_type -> greetpb.Greet
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_greet_greetpb_greet_proto_init() }
func file_greet_greetpb_greet_proto_init() {
	if File_greet_greetpb_greet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greet_greetpb_greet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Greet); i {
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
			RawDescriptor: file_greet_greetpb_greet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greet_greetpb_greet_proto_goTypes,
		DependencyIndexes: file_greet_greetpb_greet_proto_depIdxs,
		MessageInfos:      file_greet_greetpb_greet_proto_msgTypes,
	}.Build()
	File_greet_greetpb_greet_proto = out.File
	file_greet_greetpb_greet_proto_rawDesc = nil
	file_greet_greetpb_greet_proto_goTypes = nil
	file_greet_greetpb_greet_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	Greeting(ctx context.Context, in *Greet, opts ...grpc.CallOption) (*Greet, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greeting(ctx context.Context, in *Greet, opts ...grpc.CallOption) (*Greet, error) {
	out := new(Greet)
	err := c.cc.Invoke(ctx, "/greetpb.GreetService/Greeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	Greeting(context.Context, *Greet) (*Greet, error)
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greeting(context.Context, *Greet) (*Greet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greeting not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Greet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greetpb.GreetService/Greeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greeting(ctx, req.(*Greet))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greetpb.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greeting",
			Handler:    _GreetService_Greeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greet/greetpb/greet.proto",
}
