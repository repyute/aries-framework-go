// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/ecdh_aead.proto

package ecdh_aead_go_proto

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	common_go_proto "github.com/google/tink/go/proto/common_go_proto"
	tink_go_proto "github.com/google/tink/go/proto/tink_go_proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type KeyType int32

const (
	KeyType_UNKNOWN_KEY_TYPE KeyType = 0
	KeyType_EC               KeyType = 1
	KeyType_OKP              KeyType = 2
)

// Enum value maps for KeyType.
var (
	KeyType_name = map[int32]string{
		0: "UNKNOWN_KEY_TYPE",
		1: "EC",
		2: "OKP",
	}
	KeyType_value = map[string]int32{
		"UNKNOWN_KEY_TYPE": 0,
		"EC":               1,
		"OKP":              2,
	}
)

func (x KeyType) Enum() *KeyType {
	p := new(KeyType)
	*p = x
	return p
}

func (x KeyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_ecdh_aead_proto_enumTypes[0].Descriptor()
}

func (KeyType) Type() protoreflect.EnumType {
	return &file_proto_ecdh_aead_proto_enumTypes[0]
}

func (x KeyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyType.Descriptor instead.
func (KeyType) EnumDescriptor() ([]byte, []int) {
	return file_proto_ecdh_aead_proto_rawDescGZIP(), []int{0}
}

type EcdhKwParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurveType common_go_proto.EllipticCurveType `protobuf:"varint,1,opt,name=curve_type,json=curveType,proto3,enum=google.crypto.tink.EllipticCurveType" json:"curve_type,omitempty"`
	KeyType   KeyType                           `protobuf:"varint,2,opt,name=key_type,json=keyType,proto3,enum=google.crypto.tink.KeyType" json:"key_type,omitempty"`
}

func (x *EcdhKwParams) Reset() {
	*x = EcdhKwParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecdh_aead_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcdhKwParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcdhKwParams) ProtoMessage() {}

func (x *EcdhKwParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecdh_aead_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcdhKwParams.ProtoReflect.Descriptor instead.
func (*EcdhKwParams) Descriptor() ([]byte, []int) {
	return file_proto_ecdh_aead_proto_rawDescGZIP(), []int{0}
}

func (x *EcdhKwParams) GetCurveType() common_go_proto.EllipticCurveType {
	if x != nil {
		return x.CurveType
	}
	return common_go_proto.EllipticCurveType_UNKNOWN_CURVE
}

func (x *EcdhKwParams) GetKeyType() KeyType {
	if x != nil {
		return x.KeyType
	}
	return KeyType_UNKNOWN_KEY_TYPE
}

type EcdhAeadEncParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AeadEnc *tink_go_proto.KeyTemplate `protobuf:"bytes,1,opt,name=aead_enc,json=aeadEnc,proto3" json:"aead_enc,omitempty"`
	CEK     []byte                     `protobuf:"bytes,2,opt,name=CEK,proto3" json:"CEK,omitempty"`
}

func (x *EcdhAeadEncParams) Reset() {
	*x = EcdhAeadEncParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecdh_aead_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcdhAeadEncParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcdhAeadEncParams) ProtoMessage() {}

func (x *EcdhAeadEncParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecdh_aead_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcdhAeadEncParams.ProtoReflect.Descriptor instead.
func (*EcdhAeadEncParams) Descriptor() ([]byte, []int) {
	return file_proto_ecdh_aead_proto_rawDescGZIP(), []int{1}
}

func (x *EcdhAeadEncParams) GetAeadEnc() *tink_go_proto.KeyTemplate {
	if x != nil {
		return x.AeadEnc
	}
	return nil
}

func (x *EcdhAeadEncParams) GetCEK() []byte {
	if x != nil {
		return x.CEK
	}
	return nil
}

type EcdhAeadParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KwParams      *EcdhKwParams                 `protobuf:"bytes,1,opt,name=kw_params,json=kwParams,proto3" json:"kw_params,omitempty"`
	EncParams     *EcdhAeadEncParams            `protobuf:"bytes,2,opt,name=enc_params,json=encParams,proto3" json:"enc_params,omitempty"`
	EcPointFormat common_go_proto.EcPointFormat `protobuf:"varint,3,opt,name=ec_point_format,json=ecPointFormat,proto3,enum=google.crypto.tink.EcPointFormat" json:"ec_point_format,omitempty"`
}

func (x *EcdhAeadParams) Reset() {
	*x = EcdhAeadParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecdh_aead_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcdhAeadParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcdhAeadParams) ProtoMessage() {}

func (x *EcdhAeadParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecdh_aead_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcdhAeadParams.ProtoReflect.Descriptor instead.
func (*EcdhAeadParams) Descriptor() ([]byte, []int) {
	return file_proto_ecdh_aead_proto_rawDescGZIP(), []int{2}
}

func (x *EcdhAeadParams) GetKwParams() *EcdhKwParams {
	if x != nil {
		return x.KwParams
	}
	return nil
}

func (x *EcdhAeadParams) GetEncParams() *EcdhAeadEncParams {
	if x != nil {
		return x.EncParams
	}
	return nil
}

func (x *EcdhAeadParams) GetEcPointFormat() common_go_proto.EcPointFormat {
	if x != nil {
		return x.EcPointFormat
	}
	return common_go_proto.EcPointFormat_UNKNOWN_FORMAT
}

type EcdhAeadPublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint32          `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Params  *EcdhAeadParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	KID     string          `protobuf:"bytes,3,opt,name=KID,proto3" json:"KID,omitempty"`
	X       []byte          `protobuf:"bytes,4,opt,name=x,proto3" json:"x,omitempty"`
	Y       []byte          `protobuf:"bytes,5,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *EcdhAeadPublicKey) Reset() {
	*x = EcdhAeadPublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecdh_aead_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcdhAeadPublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcdhAeadPublicKey) ProtoMessage() {}

func (x *EcdhAeadPublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecdh_aead_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcdhAeadPublicKey.ProtoReflect.Descriptor instead.
func (*EcdhAeadPublicKey) Descriptor() ([]byte, []int) {
	return file_proto_ecdh_aead_proto_rawDescGZIP(), []int{3}
}

func (x *EcdhAeadPublicKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *EcdhAeadPublicKey) GetParams() *EcdhAeadParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *EcdhAeadPublicKey) GetKID() string {
	if x != nil {
		return x.KID
	}
	return ""
}

func (x *EcdhAeadPublicKey) GetX() []byte {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *EcdhAeadPublicKey) GetY() []byte {
	if x != nil {
		return x.Y
	}
	return nil
}

type EcdhAeadPrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint32             `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	PublicKey *EcdhAeadPublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	KeyValue  []byte             `protobuf:"bytes,3,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
}

func (x *EcdhAeadPrivateKey) Reset() {
	*x = EcdhAeadPrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecdh_aead_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcdhAeadPrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcdhAeadPrivateKey) ProtoMessage() {}

func (x *EcdhAeadPrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecdh_aead_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcdhAeadPrivateKey.ProtoReflect.Descriptor instead.
func (*EcdhAeadPrivateKey) Descriptor() ([]byte, []int) {
	return file_proto_ecdh_aead_proto_rawDescGZIP(), []int{4}
}

func (x *EcdhAeadPrivateKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *EcdhAeadPrivateKey) GetPublicKey() *EcdhAeadPublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *EcdhAeadPrivateKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

type EcdhAeadKeyFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *EcdhAeadParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *EcdhAeadKeyFormat) Reset() {
	*x = EcdhAeadKeyFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecdh_aead_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcdhAeadKeyFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcdhAeadKeyFormat) ProtoMessage() {}

func (x *EcdhAeadKeyFormat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecdh_aead_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcdhAeadKeyFormat.ProtoReflect.Descriptor instead.
func (*EcdhAeadKeyFormat) Descriptor() ([]byte, []int) {
	return file_proto_ecdh_aead_proto_rawDescGZIP(), []int{5}
}

func (x *EcdhAeadKeyFormat) GetParams() *EcdhAeadParams {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_proto_ecdh_aead_proto protoreflect.FileDescriptor

var file_proto_ecdh_aead_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x63, 0x64, 0x68, 0x5f, 0x61, 0x65, 0x61,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x1a, 0x12, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x0c, 0x45, 0x63, 0x64, 0x68, 0x4b, 0x77, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x44, 0x0a, 0x0a, 0x63, 0x75, 0x72, 0x76, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x45, 0x6c, 0x6c, 0x69,
	0x70, 0x74, 0x69, 0x63, 0x43, 0x75, 0x72, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x63,
	0x75, 0x72, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e,
	0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x61, 0x0a, 0x11, 0x45, 0x63, 0x64, 0x68, 0x41, 0x65, 0x61, 0x64, 0x45, 0x6e, 0x63, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x3a, 0x0a, 0x08, 0x61, 0x65, 0x61, 0x64, 0x5f, 0x65, 0x6e,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x4b, 0x65, 0x79,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x07, 0x61, 0x65, 0x61, 0x64, 0x45, 0x6e,
	0x63, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x45, 0x4b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x43, 0x45, 0x4b, 0x22, 0xe0, 0x01, 0x0a, 0x0e, 0x45, 0x63, 0x64, 0x68, 0x41, 0x65, 0x61, 0x64,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x3d, 0x0a, 0x09, 0x6b, 0x77, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x45,
	0x63, 0x64, 0x68, 0x4b, 0x77, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x08, 0x6b, 0x77, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x44, 0x0a, 0x0a, 0x65, 0x6e, 0x63, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x45,
	0x63, 0x64, 0x68, 0x41, 0x65, 0x61, 0x64, 0x45, 0x6e, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x52, 0x09, 0x65, 0x6e, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x49, 0x0a, 0x0f, 0x65,
	0x63, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x45, 0x63, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x0d, 0x65, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x97, 0x01, 0x0a, 0x11, 0x45, 0x63, 0x64, 0x68, 0x41,
	0x65, 0x61, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x45, 0x63, 0x64, 0x68,
	0x41, 0x65, 0x61, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4b, 0x49, 0x44, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x79,
	0x22, 0x91, 0x01, 0x0a, 0x12, 0x45, 0x63, 0x64, 0x68, 0x41, 0x65, 0x61, 0x64, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x44, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x45, 0x63, 0x64, 0x68, 0x41,
	0x65, 0x61, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x4f, 0x0a, 0x11, 0x45, 0x63, 0x64, 0x68, 0x41, 0x65, 0x61, 0x64,
	0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x45,
	0x63, 0x64, 0x68, 0x41, 0x65, 0x61, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x2a, 0x30, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4b, 0x45, 0x59, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x45, 0x43, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x4f, 0x4b, 0x50, 0x10, 0x02, 0x42, 0x8d, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69,
	0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x62, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2f, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2d, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x6d,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x63, 0x64, 0x68,
	0x5f, 0x61, 0x65, 0x61, 0x64, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02,
	0x06, 0x54, 0x49, 0x4e, 0x4b, 0x50, 0x42, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ecdh_aead_proto_rawDescOnce sync.Once
	file_proto_ecdh_aead_proto_rawDescData = file_proto_ecdh_aead_proto_rawDesc
)

func file_proto_ecdh_aead_proto_rawDescGZIP() []byte {
	file_proto_ecdh_aead_proto_rawDescOnce.Do(func() {
		file_proto_ecdh_aead_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ecdh_aead_proto_rawDescData)
	})
	return file_proto_ecdh_aead_proto_rawDescData
}

var file_proto_ecdh_aead_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_ecdh_aead_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_ecdh_aead_proto_goTypes = []interface{}{
	(KeyType)(0),                           // 0: google.crypto.tink.KeyType
	(*EcdhKwParams)(nil),                   // 1: google.crypto.tink.EcdhKwParams
	(*EcdhAeadEncParams)(nil),              // 2: google.crypto.tink.EcdhAeadEncParams
	(*EcdhAeadParams)(nil),                 // 3: google.crypto.tink.EcdhAeadParams
	(*EcdhAeadPublicKey)(nil),              // 4: google.crypto.tink.EcdhAeadPublicKey
	(*EcdhAeadPrivateKey)(nil),             // 5: google.crypto.tink.EcdhAeadPrivateKey
	(*EcdhAeadKeyFormat)(nil),              // 6: google.crypto.tink.EcdhAeadKeyFormat
	(common_go_proto.EllipticCurveType)(0), // 7: google.crypto.tink.EllipticCurveType
	(*tink_go_proto.KeyTemplate)(nil),      // 8: google.crypto.tink.KeyTemplate
	(common_go_proto.EcPointFormat)(0),     // 9: google.crypto.tink.EcPointFormat
}
var file_proto_ecdh_aead_proto_depIdxs = []int32{
	7, // 0: google.crypto.tink.EcdhKwParams.curve_type:type_name -> google.crypto.tink.EllipticCurveType
	0, // 1: google.crypto.tink.EcdhKwParams.key_type:type_name -> google.crypto.tink.KeyType
	8, // 2: google.crypto.tink.EcdhAeadEncParams.aead_enc:type_name -> google.crypto.tink.KeyTemplate
	1, // 3: google.crypto.tink.EcdhAeadParams.kw_params:type_name -> google.crypto.tink.EcdhKwParams
	2, // 4: google.crypto.tink.EcdhAeadParams.enc_params:type_name -> google.crypto.tink.EcdhAeadEncParams
	9, // 5: google.crypto.tink.EcdhAeadParams.ec_point_format:type_name -> google.crypto.tink.EcPointFormat
	3, // 6: google.crypto.tink.EcdhAeadPublicKey.params:type_name -> google.crypto.tink.EcdhAeadParams
	4, // 7: google.crypto.tink.EcdhAeadPrivateKey.public_key:type_name -> google.crypto.tink.EcdhAeadPublicKey
	3, // 8: google.crypto.tink.EcdhAeadKeyFormat.params:type_name -> google.crypto.tink.EcdhAeadParams
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_proto_ecdh_aead_proto_init() }
func file_proto_ecdh_aead_proto_init() {
	if File_proto_ecdh_aead_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ecdh_aead_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcdhKwParams); i {
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
		file_proto_ecdh_aead_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcdhAeadEncParams); i {
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
		file_proto_ecdh_aead_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcdhAeadParams); i {
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
		file_proto_ecdh_aead_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcdhAeadPublicKey); i {
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
		file_proto_ecdh_aead_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcdhAeadPrivateKey); i {
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
		file_proto_ecdh_aead_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcdhAeadKeyFormat); i {
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
			RawDescriptor: file_proto_ecdh_aead_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ecdh_aead_proto_goTypes,
		DependencyIndexes: file_proto_ecdh_aead_proto_depIdxs,
		EnumInfos:         file_proto_ecdh_aead_proto_enumTypes,
		MessageInfos:      file_proto_ecdh_aead_proto_msgTypes,
	}.Build()
	File_proto_ecdh_aead_proto = out.File
	file_proto_ecdh_aead_proto_rawDesc = nil
	file_proto_ecdh_aead_proto_goTypes = nil
	file_proto_ecdh_aead_proto_depIdxs = nil
}