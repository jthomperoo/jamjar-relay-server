//
//Copyright 2021 The JamJar Relay Server Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: v1/relay/relay.proto

package relay

import (
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

type Relay_RelayType int32

const (
	Relay_BROADCAST Relay_RelayType = 0
	Relay_TARGET    Relay_RelayType = 1
	Relay_HOST      Relay_RelayType = 2
)

// Enum value maps for Relay_RelayType.
var (
	Relay_RelayType_name = map[int32]string{
		0: "BROADCAST",
		1: "TARGET",
		2: "HOST",
	}
	Relay_RelayType_value = map[string]int32{
		"BROADCAST": 0,
		"TARGET":    1,
		"HOST":      2,
	}
)

func (x Relay_RelayType) Enum() *Relay_RelayType {
	p := new(Relay_RelayType)
	*p = x
	return p
}

func (x Relay_RelayType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Relay_RelayType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_relay_relay_proto_enumTypes[0].Descriptor()
}

func (Relay_RelayType) Type() protoreflect.EnumType {
	return &file_v1_relay_relay_proto_enumTypes[0]
}

func (x Relay_RelayType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Relay_RelayType.Descriptor instead.
func (Relay_RelayType) EnumDescriptor() ([]byte, []int) {
	return file_v1_relay_relay_proto_rawDescGZIP(), []int{0, 0}
}

type Relay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   Relay_RelayType `protobuf:"varint,1,opt,name=Type,proto3,enum=v1_relay.Relay_RelayType" json:"Type,omitempty"`
	Target *int32          `protobuf:"varint,2,opt,name=Target,proto3,oneof" json:"Target,omitempty"`
	Data   []byte          `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Relay) Reset() {
	*x = Relay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_relay_relay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relay) ProtoMessage() {}

func (x *Relay) ProtoReflect() protoreflect.Message {
	mi := &file_v1_relay_relay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relay.ProtoReflect.Descriptor instead.
func (*Relay) Descriptor() ([]byte, []int) {
	return file_v1_relay_relay_proto_rawDescGZIP(), []int{0}
}

func (x *Relay) GetType() Relay_RelayType {
	if x != nil {
		return x.Type
	}
	return Relay_BROADCAST
}

func (x *Relay) GetTarget() int32 {
	if x != nil && x.Target != nil {
		return *x.Target
	}
	return 0
}

func (x *Relay) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_v1_relay_relay_proto protoreflect.FileDescriptor

var file_v1_relay_relay_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x31, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x22, 0xa4, 0x01, 0x0a, 0x05, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x2d, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x76, 0x31, 0x5f, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x30, 0x0a, 0x09, 0x52, 0x65,
	0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x52, 0x4f, 0x41, 0x44,
	0x43, 0x41, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x02, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x6d, 0x6a, 0x61, 0x72, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x6a, 0x61, 0x6d, 0x6a, 0x61, 0x72, 0x2d, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_relay_relay_proto_rawDescOnce sync.Once
	file_v1_relay_relay_proto_rawDescData = file_v1_relay_relay_proto_rawDesc
)

func file_v1_relay_relay_proto_rawDescGZIP() []byte {
	file_v1_relay_relay_proto_rawDescOnce.Do(func() {
		file_v1_relay_relay_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_relay_relay_proto_rawDescData)
	})
	return file_v1_relay_relay_proto_rawDescData
}

var file_v1_relay_relay_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_relay_relay_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v1_relay_relay_proto_goTypes = []interface{}{
	(Relay_RelayType)(0), // 0: v1_relay.Relay.RelayType
	(*Relay)(nil),        // 1: v1_relay.Relay
}
var file_v1_relay_relay_proto_depIdxs = []int32{
	0, // 0: v1_relay.Relay.Type:type_name -> v1_relay.Relay.RelayType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_relay_relay_proto_init() }
func file_v1_relay_relay_proto_init() {
	if File_v1_relay_relay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_relay_relay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Relay); i {
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
	file_v1_relay_relay_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_relay_relay_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_relay_relay_proto_goTypes,
		DependencyIndexes: file_v1_relay_relay_proto_depIdxs,
		EnumInfos:         file_v1_relay_relay_proto_enumTypes,
		MessageInfos:      file_v1_relay_relay_proto_msgTypes,
	}.Build()
	File_v1_relay_relay_proto = out.File
	file_v1_relay_relay_proto_rawDesc = nil
	file_v1_relay_relay_proto_goTypes = nil
	file_v1_relay_relay_proto_depIdxs = nil
}
