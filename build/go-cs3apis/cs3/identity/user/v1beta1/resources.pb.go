// Copyright 2018-2019 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.0
// source: cs3/identity/user/v1beta1/resources.proto

package userv1beta1

import (
	v1beta1 "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/types/v1beta1"
	proto "github.com/golang/protobuf/proto"
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

// A UserId represents a user.
type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The identity provider for the user.
	Idp string `protobuf:"bytes,1,opt,name=idp,proto3" json:"idp,omitempty"`
	// REQUIRED.
	// the unique identifier for the user in the scope of
	// the identity provider.
	OpaqueId string `protobuf:"bytes,2,opt,name=opaque_id,json=opaqueId,proto3" json:"opaque_id,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_identity_user_v1beta1_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_identity_user_v1beta1_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_cs3_identity_user_v1beta1_resources_proto_rawDescGZIP(), []int{0}
}

func (x *UserId) GetIdp() string {
	if x != nil {
		return x.Idp
	}
	return ""
}

func (x *UserId) GetOpaqueId() string {
	if x != nil {
		return x.OpaqueId
	}
	return ""
}

// Represents a user of the system.
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           *UserId         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username     string          `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Mail         string          `protobuf:"bytes,3,opt,name=mail,proto3" json:"mail,omitempty"`
	MailVerified bool            `protobuf:"varint,4,opt,name=mail_verified,json=mailVerified,proto3" json:"mail_verified,omitempty"`
	DisplayName  string          `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Groups       []string        `protobuf:"bytes,6,rep,name=groups,proto3" json:"groups,omitempty"`
	Opaque       *v1beta1.Opaque `protobuf:"bytes,7,opt,name=opaque,proto3" json:"opaque,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_identity_user_v1beta1_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_identity_user_v1beta1_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_cs3_identity_user_v1beta1_resources_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() *UserId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

func (x *User) GetMailVerified() bool {
	if x != nil {
		return x.MailVerified
	}
	return false
}

func (x *User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *User) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *User) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

var File_cs3_identity_user_v1beta1_resources_proto protoreflect.FileDescriptor

var file_cs3_identity_user_v1beta1_resources_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x73, 0x33, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x73, 0x33,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1d, 0x63, 0x73, 0x33, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x64, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64,
	0x70, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x49, 0x64, 0x22, 0xfc,
	0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61,
	0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x73, 0x33,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f,
	0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x42, 0x7c, 0x0a,
	0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x43, 0x49, 0x55, 0xaa, 0x02, 0x19, 0x43, 0x73, 0x33, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca,
	0x02, 0x19, 0x43, 0x73, 0x33, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x55,
	0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cs3_identity_user_v1beta1_resources_proto_rawDescOnce sync.Once
	file_cs3_identity_user_v1beta1_resources_proto_rawDescData = file_cs3_identity_user_v1beta1_resources_proto_rawDesc
)

func file_cs3_identity_user_v1beta1_resources_proto_rawDescGZIP() []byte {
	file_cs3_identity_user_v1beta1_resources_proto_rawDescOnce.Do(func() {
		file_cs3_identity_user_v1beta1_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_cs3_identity_user_v1beta1_resources_proto_rawDescData)
	})
	return file_cs3_identity_user_v1beta1_resources_proto_rawDescData
}

var file_cs3_identity_user_v1beta1_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cs3_identity_user_v1beta1_resources_proto_goTypes = []interface{}{
	(*UserId)(nil),         // 0: cs3.identity.user.v1beta1.UserId
	(*User)(nil),           // 1: cs3.identity.user.v1beta1.User
	(*v1beta1.Opaque)(nil), // 2: cs3.types.v1beta1.Opaque
}
var file_cs3_identity_user_v1beta1_resources_proto_depIdxs = []int32{
	0, // 0: cs3.identity.user.v1beta1.User.id:type_name -> cs3.identity.user.v1beta1.UserId
	2, // 1: cs3.identity.user.v1beta1.User.opaque:type_name -> cs3.types.v1beta1.Opaque
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cs3_identity_user_v1beta1_resources_proto_init() }
func file_cs3_identity_user_v1beta1_resources_proto_init() {
	if File_cs3_identity_user_v1beta1_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cs3_identity_user_v1beta1_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
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
		file_cs3_identity_user_v1beta1_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_cs3_identity_user_v1beta1_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cs3_identity_user_v1beta1_resources_proto_goTypes,
		DependencyIndexes: file_cs3_identity_user_v1beta1_resources_proto_depIdxs,
		MessageInfos:      file_cs3_identity_user_v1beta1_resources_proto_msgTypes,
	}.Build()
	File_cs3_identity_user_v1beta1_resources_proto = out.File
	file_cs3_identity_user_v1beta1_resources_proto_rawDesc = nil
	file_cs3_identity_user_v1beta1_resources_proto_goTypes = nil
	file_cs3_identity_user_v1beta1_resources_proto_depIdxs = nil
}