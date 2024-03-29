// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: auth/v1/auth.proto

package v1

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

type AuthUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserName    string            `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Email       string            `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	AvatarUrl   string            `protobuf:"bytes,5,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	RawUser     map[string]string `protobuf:"bytes,6,rep,name=raw_user,json=rawUser,proto3" json:"raw_user,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AccessToken string            `protobuf:"bytes,7,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *AuthUser) Reset() {
	*x = AuthUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUser) ProtoMessage() {}

func (x *AuthUser) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUser.ProtoReflect.Descriptor instead.
func (*AuthUser) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthUser) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *AuthUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthUser) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *AuthUser) GetRawUser() map[string]string {
	if x != nil {
		return x.RawUser
	}
	return nil
}

func (x *AuthUser) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

var File_auth_v1_auth_proto protoreflect.FileDescriptor

var file_auth_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x61, 0x70, 0x69, 0x69, 0x6f, 0x22, 0x99, 0x02, 0x0a,
	0x08, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12,
	0x38, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x61, 0x70, 0x69, 0x69, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x77, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x72, 0x61, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x3a, 0x0a, 0x0c,
	0x52, 0x61, 0x77, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x92, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d,
	0x2e, 0x68, 0x61, 0x70, 0x69, 0x69, 0x6f, 0x42, 0x09, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6a, 0x6f, 0x68, 0x61, 0x6e, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x68, 0x6f, 0x72, 0x73, 0x74,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x62, 0x6f,
	0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x58, 0x58, 0xaa, 0x02, 0x06,
	0x48, 0x61, 0x70, 0x69, 0x69, 0x6f, 0xca, 0x02, 0x06, 0x48, 0x61, 0x70, 0x69, 0x69, 0x6f, 0xe2,
	0x02, 0x12, 0x48, 0x61, 0x70, 0x69, 0x69, 0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x48, 0x61, 0x70, 0x69, 0x69, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_v1_auth_proto_rawDescOnce sync.Once
	file_auth_v1_auth_proto_rawDescData = file_auth_v1_auth_proto_rawDesc
)

func file_auth_v1_auth_proto_rawDescGZIP() []byte {
	file_auth_v1_auth_proto_rawDescOnce.Do(func() {
		file_auth_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_v1_auth_proto_rawDescData)
	})
	return file_auth_v1_auth_proto_rawDescData
}

var file_auth_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_v1_auth_proto_goTypes = []interface{}{
	(*AuthUser)(nil), // 0: hapiio.AuthUser
	nil,              // 1: hapiio.AuthUser.RawUserEntry
}
var file_auth_v1_auth_proto_depIdxs = []int32{
	1, // 0: hapiio.AuthUser.raw_user:type_name -> hapiio.AuthUser.RawUserEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auth_v1_auth_proto_init() }
func file_auth_v1_auth_proto_init() {
	if File_auth_v1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_v1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthUser); i {
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
			RawDescriptor: file_auth_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_v1_auth_proto_goTypes,
		DependencyIndexes: file_auth_v1_auth_proto_depIdxs,
		MessageInfos:      file_auth_v1_auth_proto_msgTypes,
	}.Build()
	File_auth_v1_auth_proto = out.File
	file_auth_v1_auth_proto_rawDesc = nil
	file_auth_v1_auth_proto_goTypes = nil
	file_auth_v1_auth_proto_depIdxs = nil
}
