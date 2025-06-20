// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: user/user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "protobuf/pb/common"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName      string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Age           int64                  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Balance       float64                `protobuf:"fixed64,4,opt,name=balance,proto3" json:"balance,omitempty"`
	IsActive      bool                   `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Address       *Address               `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Educations    []string               `protobuf:"bytes,7,rep,name=educations,proto3" json:"educations,omitempty"`
	SpouseName    string                 `protobuf:"bytes,8,opt,name=spouse_name,json=spouseName,proto3" json:"spouse_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_user_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[0]
	if x != nil {
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
	return file_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *User) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *User) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *User) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *User) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *User) GetEducations() []string {
	if x != nil {
		return x.Educations
	}
	return nil
}

func (x *User) GetSpouseName() string {
	if x != nil {
		return x.SpouseName
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FullAddress   string                 `protobuf:"bytes,2,opt,name=full_address,json=fullAddress,proto3" json:"full_address,omitempty"`
	Province      string                 `protobuf:"bytes,3,opt,name=province,proto3" json:"province,omitempty"`
	City          string                 `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_user_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Address) GetFullAddress() string {
	if x != nil {
		return x.FullAddress
	}
	return ""
}

func (x *Address) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BaseResponse  *common.BaseResponse   `protobuf:"bytes,1,opt,name=base_response,json=baseResponse,proto3" json:"base_response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	mi := &file_user_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetBaseResponse() *common.BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

var File_user_user_proto protoreflect.FileDescriptor

const file_user_user_proto_rawDesc = "" +
	"\n" +
	"\x0fuser/user.proto\x12\x04user\x1a\x1acommon/base_response.proto\"\xe6\x01\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x1b\n" +
	"\tfull_name\x18\x02 \x01(\tR\bfullName\x12\x10\n" +
	"\x03age\x18\x03 \x01(\x03R\x03age\x12\x18\n" +
	"\abalance\x18\x04 \x01(\x01R\abalance\x12\x1b\n" +
	"\tis_active\x18\x05 \x01(\bR\bisActive\x12'\n" +
	"\aaddress\x18\x06 \x01(\v2\r.user.AddressR\aaddress\x12\x1e\n" +
	"\n" +
	"educations\x18\a \x03(\tR\n" +
	"educations\x12\x1f\n" +
	"\vspouse_name\x18\b \x01(\tR\n" +
	"spouseName\"l\n" +
	"\aAddress\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12!\n" +
	"\ffull_address\x18\x02 \x01(\tR\vfullAddress\x12\x1a\n" +
	"\bprovince\x18\x03 \x01(\tR\bprovince\x12\x12\n" +
	"\x04city\x18\x04 \x01(\tR\x04city\"K\n" +
	"\x0eCreateResponse\x129\n" +
	"\rbase_response\x18\x01 \x01(\v2\x14.common.BaseResponseR\fbaseResponse2=\n" +
	"\vUserService\x12.\n" +
	"\n" +
	"CreateUser\x12\n" +
	".user.User\x1a\x14.user.CreateResponseB\x12Z\x10protobuf/pb;userb\x06proto3"

var (
	file_user_user_proto_rawDescOnce sync.Once
	file_user_user_proto_rawDescData []byte
)

func file_user_user_proto_rawDescGZIP() []byte {
	file_user_user_proto_rawDescOnce.Do(func() {
		file_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_user_proto_rawDesc), len(file_user_user_proto_rawDesc)))
	})
	return file_user_user_proto_rawDescData
}

var file_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_user_user_proto_goTypes = []any{
	(*User)(nil),                // 0: user.User
	(*Address)(nil),             // 1: user.Address
	(*CreateResponse)(nil),      // 2: user.CreateResponse
	(*common.BaseResponse)(nil), // 3: common.BaseResponse
}
var file_user_user_proto_depIdxs = []int32{
	1, // 0: user.User.address:type_name -> user.Address
	3, // 1: user.CreateResponse.base_response:type_name -> common.BaseResponse
	0, // 2: user.UserService.CreateUser:input_type -> user.User
	2, // 3: user.UserService.CreateUser:output_type -> user.CreateResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_user_proto_init() }
func file_user_user_proto_init() {
	if File_user_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_user_proto_rawDesc), len(file_user_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_user_proto_goTypes,
		DependencyIndexes: file_user_user_proto_depIdxs,
		MessageInfos:      file_user_user_proto_msgTypes,
	}.Build()
	File_user_user_proto = out.File
	file_user_user_proto_goTypes = nil
	file_user_user_proto_depIdxs = nil
}
