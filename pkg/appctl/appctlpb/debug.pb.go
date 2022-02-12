// Copyright (C) 2021  mieru authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: debug.proto

package appctlpb

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

type ThreadDump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThreadDump string `protobuf:"bytes,1,opt,name=threadDump,proto3" json:"threadDump,omitempty"`
}

func (x *ThreadDump) Reset() {
	*x = ThreadDump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreadDump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadDump) ProtoMessage() {}

func (x *ThreadDump) ProtoReflect() protoreflect.Message {
	mi := &file_debug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadDump.ProtoReflect.Descriptor instead.
func (*ThreadDump) Descriptor() ([]byte, []int) {
	return file_debug_proto_rawDescGZIP(), []int{0}
}

func (x *ThreadDump) GetThreadDump() string {
	if x != nil {
		return x.ThreadDump
	}
	return ""
}

type ProfileSavePath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string `protobuf:"bytes,1,opt,name=filePath,proto3" json:"filePath,omitempty"`
}

func (x *ProfileSavePath) Reset() {
	*x = ProfileSavePath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileSavePath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileSavePath) ProtoMessage() {}

func (x *ProfileSavePath) ProtoReflect() protoreflect.Message {
	mi := &file_debug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileSavePath.ProtoReflect.Descriptor instead.
func (*ProfileSavePath) Descriptor() ([]byte, []int) {
	return file_debug_proto_rawDescGZIP(), []int{1}
}

func (x *ProfileSavePath) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

var File_debug_proto protoreflect.FileDescriptor

var file_debug_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61,
	0x70, 0x70, 0x63, 0x74, 0x6c, 0x22, 0x2c, 0x0a, 0x0a, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44,
	0x75, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44,
	0x75, 0x6d, 0x70, 0x22, 0x2d, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x61,
	0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6e, 0x66, 0x65, 0x69, 0x6e, 0x2f, 0x6d, 0x69, 0x65, 0x72, 0x75, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_debug_proto_rawDescOnce sync.Once
	file_debug_proto_rawDescData = file_debug_proto_rawDesc
)

func file_debug_proto_rawDescGZIP() []byte {
	file_debug_proto_rawDescOnce.Do(func() {
		file_debug_proto_rawDescData = protoimpl.X.CompressGZIP(file_debug_proto_rawDescData)
	})
	return file_debug_proto_rawDescData
}

var file_debug_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_debug_proto_goTypes = []interface{}{
	(*ThreadDump)(nil),      // 0: appctl.ThreadDump
	(*ProfileSavePath)(nil), // 1: appctl.ProfileSavePath
}
var file_debug_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_debug_proto_init() }
func file_debug_proto_init() {
	if File_debug_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_debug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThreadDump); i {
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
		file_debug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileSavePath); i {
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
			RawDescriptor: file_debug_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_debug_proto_goTypes,
		DependencyIndexes: file_debug_proto_depIdxs,
		MessageInfos:      file_debug_proto_msgTypes,
	}.Build()
	File_debug_proto = out.File
	file_debug_proto_rawDesc = nil
	file_debug_proto_goTypes = nil
	file_debug_proto_depIdxs = nil
}
