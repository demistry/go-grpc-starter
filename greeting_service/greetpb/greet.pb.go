// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: greeting_service/greetpb/greet.proto

package greetpb

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

type Greeting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Age       int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Greeting) Reset() {
	*x = Greeting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greeting_service_greetpb_greet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Greeting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Greeting) ProtoMessage() {}

func (x *Greeting) ProtoReflect() protoreflect.Message {
	mi := &file_greeting_service_greetpb_greet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Greeting.ProtoReflect.Descriptor instead.
func (*Greeting) Descriptor() ([]byte, []int) {
	return file_greeting_service_greetpb_greet_proto_rawDescGZIP(), []int{0}
}

func (x *Greeting) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Greeting) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Greeting) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type GreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetRequest) Reset() {
	*x = GreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greeting_service_greetpb_greet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequest) ProtoMessage() {}

func (x *GreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greeting_service_greetpb_greet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequest.ProtoReflect.Descriptor instead.
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return file_greeting_service_greetpb_greet_proto_rawDescGZIP(), []int{1}
}

func (x *GreetRequest) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

type GreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Greeting string `protobuf:"bytes,3,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetResponse) Reset() {
	*x = GreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greeting_service_greetpb_greet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponse) ProtoMessage() {}

func (x *GreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greeting_service_greetpb_greet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponse.ProtoReflect.Descriptor instead.
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return file_greeting_service_greetpb_greet_proto_rawDescGZIP(), []int{2}
}

func (x *GreetResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GreetResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GreetResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

type GreetManyTimesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Greeting string `protobuf:"bytes,3,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetManyTimesResponse) Reset() {
	*x = GreetManyTimesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greeting_service_greetpb_greet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetManyTimesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetManyTimesResponse) ProtoMessage() {}

func (x *GreetManyTimesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greeting_service_greetpb_greet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetManyTimesResponse.ProtoReflect.Descriptor instead.
func (*GreetManyTimesResponse) Descriptor() ([]byte, []int) {
	return file_greeting_service_greetpb_greet_proto_rawDescGZIP(), []int{3}
}

func (x *GreetManyTimesResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GreetManyTimesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GreetManyTimesResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

var File_greeting_service_greetpb_greet_proto protoreflect.FileDescriptor

var file_greeting_service_greetpb_greet_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x22, 0x58, 0x0a,
	0x08, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x3b, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x22, 0x5d, 0x0a, 0x0d, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x22, 0x66, 0x0a, 0x16, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x32, 0xa0, 0x02, 0x0a, 0x0c,
	0x44, 0x75, 0x6d, 0x6d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x05,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x50, 0x0a, 0x18, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d,
	0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x12, 0x42, 0x0a, 0x13, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65, 0x74, 0x46,
	0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x46, 0x0a, 0x15, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x13, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x1a,
	0x5a, 0x18, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_greeting_service_greetpb_greet_proto_rawDescOnce sync.Once
	file_greeting_service_greetpb_greet_proto_rawDescData = file_greeting_service_greetpb_greet_proto_rawDesc
)

func file_greeting_service_greetpb_greet_proto_rawDescGZIP() []byte {
	file_greeting_service_greetpb_greet_proto_rawDescOnce.Do(func() {
		file_greeting_service_greetpb_greet_proto_rawDescData = protoimpl.X.CompressGZIP(file_greeting_service_greetpb_greet_proto_rawDescData)
	})
	return file_greeting_service_greetpb_greet_proto_rawDescData
}

var file_greeting_service_greetpb_greet_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_greeting_service_greetpb_greet_proto_goTypes = []interface{}{
	(*Greeting)(nil),               // 0: greet.Greeting
	(*GreetRequest)(nil),           // 1: greet.GreetRequest
	(*GreetResponse)(nil),          // 2: greet.GreetResponse
	(*GreetManyTimesResponse)(nil), // 3: greet.GreetManyTimesResponse
}
var file_greeting_service_greetpb_greet_proto_depIdxs = []int32{
	0, // 0: greet.GreetRequest.greeting:type_name -> greet.Greeting
	1, // 1: greet.DummyService.Greet:input_type -> greet.GreetRequest
	1, // 2: greet.DummyService.GreetManyTimesFromServer:input_type -> greet.GreetRequest
	1, // 3: greet.DummyService.LongGreetFromClient:input_type -> greet.GreetRequest
	1, // 4: greet.DummyService.BidirectionalGreeting:input_type -> greet.GreetRequest
	2, // 5: greet.DummyService.Greet:output_type -> greet.GreetResponse
	3, // 6: greet.DummyService.GreetManyTimesFromServer:output_type -> greet.GreetManyTimesResponse
	2, // 7: greet.DummyService.LongGreetFromClient:output_type -> greet.GreetResponse
	2, // 8: greet.DummyService.BidirectionalGreeting:output_type -> greet.GreetResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_greeting_service_greetpb_greet_proto_init() }
func file_greeting_service_greetpb_greet_proto_init() {
	if File_greeting_service_greetpb_greet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greeting_service_greetpb_greet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Greeting); i {
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
		file_greeting_service_greetpb_greet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRequest); i {
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
		file_greeting_service_greetpb_greet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetResponse); i {
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
		file_greeting_service_greetpb_greet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetManyTimesResponse); i {
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
			RawDescriptor: file_greeting_service_greetpb_greet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greeting_service_greetpb_greet_proto_goTypes,
		DependencyIndexes: file_greeting_service_greetpb_greet_proto_depIdxs,
		MessageInfos:      file_greeting_service_greetpb_greet_proto_msgTypes,
	}.Build()
	File_greeting_service_greetpb_greet_proto = out.File
	file_greeting_service_greetpb_greet_proto_rawDesc = nil
	file_greeting_service_greetpb_greet_proto_goTypes = nil
	file_greeting_service_greetpb_greet_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DummyServiceClient is the client API for DummyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DummyServiceClient interface {
	// Unary
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	//Server Streaming
	GreetManyTimesFromServer(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (DummyService_GreetManyTimesFromServerClient, error)
	//Client Streaming
	LongGreetFromClient(ctx context.Context, opts ...grpc.CallOption) (DummyService_LongGreetFromClientClient, error)
	//Bidirectional Streaming
	BidirectionalGreeting(ctx context.Context, opts ...grpc.CallOption) (DummyService_BidirectionalGreetingClient, error)
}

type dummyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDummyServiceClient(cc grpc.ClientConnInterface) DummyServiceClient {
	return &dummyServiceClient{cc}
}

func (c *dummyServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.DummyService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dummyServiceClient) GreetManyTimesFromServer(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (DummyService_GreetManyTimesFromServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DummyService_serviceDesc.Streams[0], "/greet.DummyService/GreetManyTimesFromServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &dummyServiceGreetManyTimesFromServerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DummyService_GreetManyTimesFromServerClient interface {
	Recv() (*GreetManyTimesResponse, error)
	grpc.ClientStream
}

type dummyServiceGreetManyTimesFromServerClient struct {
	grpc.ClientStream
}

func (x *dummyServiceGreetManyTimesFromServerClient) Recv() (*GreetManyTimesResponse, error) {
	m := new(GreetManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dummyServiceClient) LongGreetFromClient(ctx context.Context, opts ...grpc.CallOption) (DummyService_LongGreetFromClientClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DummyService_serviceDesc.Streams[1], "/greet.DummyService/LongGreetFromClient", opts...)
	if err != nil {
		return nil, err
	}
	x := &dummyServiceLongGreetFromClientClient{stream}
	return x, nil
}

type DummyService_LongGreetFromClientClient interface {
	Send(*GreetRequest) error
	CloseAndRecv() (*GreetResponse, error)
	grpc.ClientStream
}

type dummyServiceLongGreetFromClientClient struct {
	grpc.ClientStream
}

func (x *dummyServiceLongGreetFromClientClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dummyServiceLongGreetFromClientClient) CloseAndRecv() (*GreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dummyServiceClient) BidirectionalGreeting(ctx context.Context, opts ...grpc.CallOption) (DummyService_BidirectionalGreetingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DummyService_serviceDesc.Streams[2], "/greet.DummyService/BidirectionalGreeting", opts...)
	if err != nil {
		return nil, err
	}
	x := &dummyServiceBidirectionalGreetingClient{stream}
	return x, nil
}

type DummyService_BidirectionalGreetingClient interface {
	Send(*GreetRequest) error
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type dummyServiceBidirectionalGreetingClient struct {
	grpc.ClientStream
}

func (x *dummyServiceBidirectionalGreetingClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dummyServiceBidirectionalGreetingClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DummyServiceServer is the server API for DummyService service.
type DummyServiceServer interface {
	// Unary
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	//Server Streaming
	GreetManyTimesFromServer(*GreetRequest, DummyService_GreetManyTimesFromServerServer) error
	//Client Streaming
	LongGreetFromClient(DummyService_LongGreetFromClientServer) error
	//Bidirectional Streaming
	BidirectionalGreeting(DummyService_BidirectionalGreetingServer) error
}

// UnimplementedDummyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDummyServiceServer struct {
}

func (*UnimplementedDummyServiceServer) Greet(context.Context, *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (*UnimplementedDummyServiceServer) GreetManyTimesFromServer(*GreetRequest, DummyService_GreetManyTimesFromServerServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyTimesFromServer not implemented")
}
func (*UnimplementedDummyServiceServer) LongGreetFromClient(DummyService_LongGreetFromClientServer) error {
	return status.Errorf(codes.Unimplemented, "method LongGreetFromClient not implemented")
}
func (*UnimplementedDummyServiceServer) BidirectionalGreeting(DummyService_BidirectionalGreetingServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalGreeting not implemented")
}

func RegisterDummyServiceServer(s *grpc.Server, srv DummyServiceServer) {
	s.RegisterService(&_DummyService_serviceDesc, srv)
}

func _DummyService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DummyServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.DummyService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DummyServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DummyService_GreetManyTimesFromServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DummyServiceServer).GreetManyTimesFromServer(m, &dummyServiceGreetManyTimesFromServerServer{stream})
}

type DummyService_GreetManyTimesFromServerServer interface {
	Send(*GreetManyTimesResponse) error
	grpc.ServerStream
}

type dummyServiceGreetManyTimesFromServerServer struct {
	grpc.ServerStream
}

func (x *dummyServiceGreetManyTimesFromServerServer) Send(m *GreetManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DummyService_LongGreetFromClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DummyServiceServer).LongGreetFromClient(&dummyServiceLongGreetFromClientServer{stream})
}

type DummyService_LongGreetFromClientServer interface {
	SendAndClose(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type dummyServiceLongGreetFromClientServer struct {
	grpc.ServerStream
}

func (x *dummyServiceLongGreetFromClientServer) SendAndClose(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dummyServiceLongGreetFromClientServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DummyService_BidirectionalGreeting_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DummyServiceServer).BidirectionalGreeting(&dummyServiceBidirectionalGreetingServer{stream})
}

type DummyService_BidirectionalGreetingServer interface {
	Send(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type dummyServiceBidirectionalGreetingServer struct {
	grpc.ServerStream
}

func (x *dummyServiceBidirectionalGreetingServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dummyServiceBidirectionalGreetingServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DummyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.DummyService",
	HandlerType: (*DummyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _DummyService_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetManyTimesFromServer",
			Handler:       _DummyService_GreetManyTimesFromServer_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LongGreetFromClient",
			Handler:       _DummyService_LongGreetFromClient_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalGreeting",
			Handler:       _DummyService_BidirectionalGreeting_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greeting_service/greetpb/greet.proto",
}
