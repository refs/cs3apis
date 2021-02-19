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
// source: cs3/app/provider/v1beta1/provider_api.proto

package providerv1beta1

import (
	context "context"
	v1beta12 "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/rpc/v1beta1"
	v1beta11 "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/storage/provider/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/types/v1beta1"
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

// REQUIRED.
// View mode.
type OpenFileInAppProviderRequest_ViewMode int32

const (
	OpenFileInAppProviderRequest_VIEW_MODE_INVALID OpenFileInAppProviderRequest_ViewMode = 0
	// The file can be opened but not downloaded.
	OpenFileInAppProviderRequest_VIEW_MODE_VIEW_ONLY OpenFileInAppProviderRequest_ViewMode = 1
	// The file can be downloaded.
	OpenFileInAppProviderRequest_VIEW_MODE_READ_ONLY OpenFileInAppProviderRequest_ViewMode = 2
	// The file can be downloaded and updated.
	OpenFileInAppProviderRequest_VIEW_MODE_READ_WRITE OpenFileInAppProviderRequest_ViewMode = 3
)

// Enum value maps for OpenFileInAppProviderRequest_ViewMode.
var (
	OpenFileInAppProviderRequest_ViewMode_name = map[int32]string{
		0: "VIEW_MODE_INVALID",
		1: "VIEW_MODE_VIEW_ONLY",
		2: "VIEW_MODE_READ_ONLY",
		3: "VIEW_MODE_READ_WRITE",
	}
	OpenFileInAppProviderRequest_ViewMode_value = map[string]int32{
		"VIEW_MODE_INVALID":    0,
		"VIEW_MODE_VIEW_ONLY":  1,
		"VIEW_MODE_READ_ONLY":  2,
		"VIEW_MODE_READ_WRITE": 3,
	}
)

func (x OpenFileInAppProviderRequest_ViewMode) Enum() *OpenFileInAppProviderRequest_ViewMode {
	p := new(OpenFileInAppProviderRequest_ViewMode)
	*p = x
	return p
}

func (x OpenFileInAppProviderRequest_ViewMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpenFileInAppProviderRequest_ViewMode) Descriptor() protoreflect.EnumDescriptor {
	return file_cs3_app_provider_v1beta1_provider_api_proto_enumTypes[0].Descriptor()
}

func (OpenFileInAppProviderRequest_ViewMode) Type() protoreflect.EnumType {
	return &file_cs3_app_provider_v1beta1_provider_api_proto_enumTypes[0]
}

func (x OpenFileInAppProviderRequest_ViewMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpenFileInAppProviderRequest_ViewMode.Descriptor instead.
func (OpenFileInAppProviderRequest_ViewMode) EnumDescriptor() ([]byte, []int) {
	return file_cs3_app_provider_v1beta1_provider_api_proto_rawDescGZIP(), []int{0, 0}
}

type OpenFileInAppProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The resourceInfo to be opened. The gateway grpc message has a ref instead.
	ResourceInfo *v1beta11.ResourceInfo                `protobuf:"bytes,2,opt,name=resource_info,json=resourceInfo,proto3" json:"resource_info,omitempty"`
	ViewMode     OpenFileInAppProviderRequest_ViewMode `protobuf:"varint,3,opt,name=view_mode,json=viewMode,proto3,enum=cs3.app.provider.v1beta1.OpenFileInAppProviderRequest_ViewMode" json:"view_mode,omitempty"`
	// REQUIRED.
	// The access token this application provider will use when contacting
	// the storage provider to read and write.
	// Service implementors MUST make sure that the access token only grants
	// access to the requested resource.
	// Service implementors should use a ResourceId rather than a filename to grant access, as
	// ResourceIds MUST NOT change when a resource is renamed.
	// The access token MUST be short-lived.
	// TODO(labkode): investigate token derivation techniques.
	AccessToken string `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *OpenFileInAppProviderRequest) Reset() {
	*x = OpenFileInAppProviderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_app_provider_v1beta1_provider_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenFileInAppProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenFileInAppProviderRequest) ProtoMessage() {}

func (x *OpenFileInAppProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_app_provider_v1beta1_provider_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenFileInAppProviderRequest.ProtoReflect.Descriptor instead.
func (*OpenFileInAppProviderRequest) Descriptor() ([]byte, []int) {
	return file_cs3_app_provider_v1beta1_provider_api_proto_rawDescGZIP(), []int{0}
}

func (x *OpenFileInAppProviderRequest) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *OpenFileInAppProviderRequest) GetResourceInfo() *v1beta11.ResourceInfo {
	if x != nil {
		return x.ResourceInfo
	}
	return nil
}

func (x *OpenFileInAppProviderRequest) GetViewMode() OpenFileInAppProviderRequest_ViewMode {
	if x != nil {
		return x.ViewMode
	}
	return OpenFileInAppProviderRequest_VIEW_MODE_INVALID
}

func (x *OpenFileInAppProviderRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type OpenFileInAppProviderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The url that user agents will render to clients.
	// Usually the rendering happens by using HTML iframes or in separate browser tabs.
	AppProviderUrl string `protobuf:"bytes,3,opt,name=app_provider_url,json=appProviderUrl,proto3" json:"app_provider_url,omitempty"`
}

func (x *OpenFileInAppProviderResponse) Reset() {
	*x = OpenFileInAppProviderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_app_provider_v1beta1_provider_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenFileInAppProviderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenFileInAppProviderResponse) ProtoMessage() {}

func (x *OpenFileInAppProviderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_app_provider_v1beta1_provider_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenFileInAppProviderResponse.ProtoReflect.Descriptor instead.
func (*OpenFileInAppProviderResponse) Descriptor() ([]byte, []int) {
	return file_cs3_app_provider_v1beta1_provider_api_proto_rawDescGZIP(), []int{1}
}

func (x *OpenFileInAppProviderResponse) GetStatus() *v1beta12.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *OpenFileInAppProviderResponse) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *OpenFileInAppProviderResponse) GetAppProviderUrl() string {
	if x != nil {
		return x.AppProviderUrl
	}
	return ""
}

var File_cs3_app_provider_v1beta1_provider_api_proto protoreflect.FileDescriptor

var file_cs3_app_provider_v1beta1_provider_api_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x73, 0x33, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63,
	0x73, 0x33, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x63, 0x73, 0x33, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x63, 0x73, 0x33, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x73, 0x33, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x92, 0x03, 0x0a, 0x1c, 0x4f, 0x70, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x41, 0x70, 0x70, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06,
	0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x4f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x63, 0x73, 0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5c, 0x0a, 0x09, 0x76, 0x69, 0x65, 0x77, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x63, 0x73, 0x33,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e,
	0x41, 0x70, 0x70, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x76, 0x69, 0x65,
	0x77, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6d, 0x0a, 0x08, 0x56, 0x69, 0x65, 0x77,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x4d, 0x4f, 0x44,
	0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x56,
	0x49, 0x45, 0x57, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x4f, 0x4e,
	0x4c, 0x59, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x4d, 0x4f, 0x44,
	0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x18, 0x0a,
	0x14, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x5f,
	0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x03, 0x22, 0xad, 0x01, 0x0a, 0x1d, 0x4f, 0x70, 0x65, 0x6e,
	0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x41, 0x70, 0x70, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x73, 0x33, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x73, 0x33,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f,
	0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x61, 0x70, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x32, 0x98, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x41, 0x50, 0x49, 0x12, 0x88, 0x01, 0x0a, 0x15, 0x4f, 0x70, 0x65, 0x6e,
	0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x41, 0x70, 0x70, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x36, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65,
	0x6e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x41, 0x70, 0x70, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x73, 0x33, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x41,
	0x70, 0x70, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x7f, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x10, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x70, 0x69, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x50, 0xaa, 0x02, 0x18,
	0x43, 0x73, 0x33, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x18, 0x43, 0x73, 0x33, 0x5c, 0x41,
	0x70, 0x70, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x42, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cs3_app_provider_v1beta1_provider_api_proto_rawDescOnce sync.Once
	file_cs3_app_provider_v1beta1_provider_api_proto_rawDescData = file_cs3_app_provider_v1beta1_provider_api_proto_rawDesc
)

func file_cs3_app_provider_v1beta1_provider_api_proto_rawDescGZIP() []byte {
	file_cs3_app_provider_v1beta1_provider_api_proto_rawDescOnce.Do(func() {
		file_cs3_app_provider_v1beta1_provider_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_cs3_app_provider_v1beta1_provider_api_proto_rawDescData)
	})
	return file_cs3_app_provider_v1beta1_provider_api_proto_rawDescData
}

var file_cs3_app_provider_v1beta1_provider_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cs3_app_provider_v1beta1_provider_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cs3_app_provider_v1beta1_provider_api_proto_goTypes = []interface{}{
	(OpenFileInAppProviderRequest_ViewMode)(0), // 0: cs3.app.provider.v1beta1.OpenFileInAppProviderRequest.ViewMode
	(*OpenFileInAppProviderRequest)(nil),       // 1: cs3.app.provider.v1beta1.OpenFileInAppProviderRequest
	(*OpenFileInAppProviderResponse)(nil),      // 2: cs3.app.provider.v1beta1.OpenFileInAppProviderResponse
	(*v1beta1.Opaque)(nil),                     // 3: cs3.types.v1beta1.Opaque
	(*v1beta11.ResourceInfo)(nil),              // 4: cs3.storage.provider.v1beta1.ResourceInfo
	(*v1beta12.Status)(nil),                    // 5: cs3.rpc.v1beta1.Status
}
var file_cs3_app_provider_v1beta1_provider_api_proto_depIdxs = []int32{
	3, // 0: cs3.app.provider.v1beta1.OpenFileInAppProviderRequest.opaque:type_name -> cs3.types.v1beta1.Opaque
	4, // 1: cs3.app.provider.v1beta1.OpenFileInAppProviderRequest.resource_info:type_name -> cs3.storage.provider.v1beta1.ResourceInfo
	0, // 2: cs3.app.provider.v1beta1.OpenFileInAppProviderRequest.view_mode:type_name -> cs3.app.provider.v1beta1.OpenFileInAppProviderRequest.ViewMode
	5, // 3: cs3.app.provider.v1beta1.OpenFileInAppProviderResponse.status:type_name -> cs3.rpc.v1beta1.Status
	3, // 4: cs3.app.provider.v1beta1.OpenFileInAppProviderResponse.opaque:type_name -> cs3.types.v1beta1.Opaque
	1, // 5: cs3.app.provider.v1beta1.ProviderAPI.OpenFileInAppProvider:input_type -> cs3.app.provider.v1beta1.OpenFileInAppProviderRequest
	2, // 6: cs3.app.provider.v1beta1.ProviderAPI.OpenFileInAppProvider:output_type -> cs3.app.provider.v1beta1.OpenFileInAppProviderResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cs3_app_provider_v1beta1_provider_api_proto_init() }
func file_cs3_app_provider_v1beta1_provider_api_proto_init() {
	if File_cs3_app_provider_v1beta1_provider_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cs3_app_provider_v1beta1_provider_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenFileInAppProviderRequest); i {
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
		file_cs3_app_provider_v1beta1_provider_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenFileInAppProviderResponse); i {
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
			RawDescriptor: file_cs3_app_provider_v1beta1_provider_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cs3_app_provider_v1beta1_provider_api_proto_goTypes,
		DependencyIndexes: file_cs3_app_provider_v1beta1_provider_api_proto_depIdxs,
		EnumInfos:         file_cs3_app_provider_v1beta1_provider_api_proto_enumTypes,
		MessageInfos:      file_cs3_app_provider_v1beta1_provider_api_proto_msgTypes,
	}.Build()
	File_cs3_app_provider_v1beta1_provider_api_proto = out.File
	file_cs3_app_provider_v1beta1_provider_api_proto_rawDesc = nil
	file_cs3_app_provider_v1beta1_provider_api_proto_goTypes = nil
	file_cs3_app_provider_v1beta1_provider_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProviderAPIClient is the client API for ProviderAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProviderAPIClient interface {
	// Returns the App provider URL
	// MUST return CODE_NOT_FOUND if the resource does not exist.
	OpenFileInAppProvider(ctx context.Context, in *OpenFileInAppProviderRequest, opts ...grpc.CallOption) (*OpenFileInAppProviderResponse, error)
}

type providerAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderAPIClient(cc grpc.ClientConnInterface) ProviderAPIClient {
	return &providerAPIClient{cc}
}

func (c *providerAPIClient) OpenFileInAppProvider(ctx context.Context, in *OpenFileInAppProviderRequest, opts ...grpc.CallOption) (*OpenFileInAppProviderResponse, error) {
	out := new(OpenFileInAppProviderResponse)
	err := c.cc.Invoke(ctx, "/cs3.app.provider.v1beta1.ProviderAPI/OpenFileInAppProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderAPIServer is the server API for ProviderAPI service.
type ProviderAPIServer interface {
	// Returns the App provider URL
	// MUST return CODE_NOT_FOUND if the resource does not exist.
	OpenFileInAppProvider(context.Context, *OpenFileInAppProviderRequest) (*OpenFileInAppProviderResponse, error)
}

// UnimplementedProviderAPIServer can be embedded to have forward compatible implementations.
type UnimplementedProviderAPIServer struct {
}

func (*UnimplementedProviderAPIServer) OpenFileInAppProvider(context.Context, *OpenFileInAppProviderRequest) (*OpenFileInAppProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenFileInAppProvider not implemented")
}

func RegisterProviderAPIServer(s *grpc.Server, srv ProviderAPIServer) {
	s.RegisterService(&_ProviderAPI_serviceDesc, srv)
}

func _ProviderAPI_OpenFileInAppProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenFileInAppProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderAPIServer).OpenFileInAppProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.app.provider.v1beta1.ProviderAPI/OpenFileInAppProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderAPIServer).OpenFileInAppProvider(ctx, req.(*OpenFileInAppProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProviderAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.app.provider.v1beta1.ProviderAPI",
	HandlerType: (*ProviderAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenFileInAppProvider",
			Handler:    _ProviderAPI_OpenFileInAppProvider_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/app/provider/v1beta1/provider_api.proto",
}
