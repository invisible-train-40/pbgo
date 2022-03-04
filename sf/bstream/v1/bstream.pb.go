// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: sf/bstream/v1/bstream.proto

package pbbstream

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ForkStep int32

const (
	ForkStep_STEP_UNKNOWN ForkStep = 0
	// Block is new head block of the chain, that is linear with the previous block
	ForkStep_STEP_NEW ForkStep = 1
	// Block is now forked and should be undone, it's not the head block of the chain anymore
	ForkStep_STEP_UNDO ForkStep = 2
	// Block is now irreversible and can be committed to (finality is chain specific, see chain documentation for more details)
	ForkStep_STEP_IRREVERSIBLE ForkStep = 4
)

// Enum value maps for ForkStep.
var (
	ForkStep_name = map[int32]string{
		0: "STEP_UNKNOWN",
		1: "STEP_NEW",
		2: "STEP_UNDO",
		4: "STEP_IRREVERSIBLE",
	}
	ForkStep_value = map[string]int32{
		"STEP_UNKNOWN":      0,
		"STEP_NEW":          1,
		"STEP_UNDO":         2,
		"STEP_IRREVERSIBLE": 4,
	}
)

func (x ForkStep) Enum() *ForkStep {
	p := new(ForkStep)
	*p = x
	return p
}

func (x ForkStep) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ForkStep) Descriptor() protoreflect.EnumDescriptor {
	return file_sf_bstream_v1_bstream_proto_enumTypes[0].Descriptor()
}

func (ForkStep) Type() protoreflect.EnumType {
	return &file_sf_bstream_v1_bstream_proto_enumTypes[0]
}

func (x ForkStep) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ForkStep.Descriptor instead.
func (ForkStep) EnumDescriptor() ([]byte, []int) {
	return file_sf_bstream_v1_bstream_proto_rawDescGZIP(), []int{0}
}

type Protocol int32

const (
	Protocol_UNKNOWN    Protocol = 0
	Protocol_EOS        Protocol = 1
	Protocol_ETH        Protocol = 2
	Protocol_SOLANA     Protocol = 3
	Protocol_NEAR       Protocol = 4
	Protocol_TENDERMINT Protocol = 5
)

// Enum value maps for Protocol.
var (
	Protocol_name = map[int32]string{
		0: "UNKNOWN",
		1: "EOS",
		2: "ETH",
		3: "SOLANA",
		4: "NEAR",
		5: "TENDERMINT",
	}
	Protocol_value = map[string]int32{
		"UNKNOWN":    0,
		"EOS":        1,
		"ETH":        2,
		"SOLANA":     3,
		"NEAR":       4,
		"TENDERMINT": 5,
	}
)

func (x Protocol) Enum() *Protocol {
	p := new(Protocol)
	*p = x
	return p
}

func (x Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_sf_bstream_v1_bstream_proto_enumTypes[1].Descriptor()
}

func (Protocol) Type() protoreflect.EnumType {
	return &file_sf_bstream_v1_bstream_proto_enumTypes[1]
}

func (x Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Protocol.Descriptor instead.
func (Protocol) EnumDescriptor() ([]byte, []int) {
	return file_sf_bstream_v1_bstream_proto_rawDescGZIP(), []int{1}
}

// Whether we can assume the data will come ordered, unless there is a chain reorganization.
// mindreaders output ordered data, whereas relayers can output unordered data.
// The server can fail early if the assumption of the caller cannot be fulfilled.
type BlockRequest_Order int32

const (
	BlockRequest_UNSPECIFIED BlockRequest_Order = 0
	BlockRequest_ORDERED     BlockRequest_Order = 1
	BlockRequest_UNORDERED   BlockRequest_Order = 2
)

// Enum value maps for BlockRequest_Order.
var (
	BlockRequest_Order_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "ORDERED",
		2: "UNORDERED",
	}
	BlockRequest_Order_value = map[string]int32{
		"UNSPECIFIED": 0,
		"ORDERED":     1,
		"UNORDERED":   2,
	}
)

func (x BlockRequest_Order) Enum() *BlockRequest_Order {
	p := new(BlockRequest_Order)
	*p = x
	return p
}

func (x BlockRequest_Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockRequest_Order) Descriptor() protoreflect.EnumDescriptor {
	return file_sf_bstream_v1_bstream_proto_enumTypes[2].Descriptor()
}

func (BlockRequest_Order) Type() protoreflect.EnumType {
	return &file_sf_bstream_v1_bstream_proto_enumTypes[2]
}

func (x BlockRequest_Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockRequest_Order.Descriptor instead.
func (BlockRequest_Order) EnumDescriptor() ([]byte, []int) {
	return file_sf_bstream_v1_bstream_proto_rawDescGZIP(), []int{0, 0}
}

type BlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of blocks we want to get in burst upon connection, on a best effort basis.
	Burst int64 `protobuf:"varint,1,opt,name=burst,proto3" json:"burst,omitempty"`
	// Type of blocks we're after here, is it 'ethereum' data, 'eos', etc.. The server can fail early
	// if he doesn't match the data he serves (services mismatch, etc..)
	ContentType string             `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Order       BlockRequest_Order `protobuf:"varint,3,opt,name=order,proto3,enum=sf.bstream.v1.BlockRequest_Order" json:"order,omitempty"`
	Requester   string             `protobuf:"bytes,4,opt,name=requester,proto3" json:"requester,omitempty"`
}

func (x *BlockRequest) Reset() {
	*x = BlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_bstream_v1_bstream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRequest) ProtoMessage() {}

func (x *BlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sf_bstream_v1_bstream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRequest.ProtoReflect.Descriptor instead.
func (*BlockRequest) Descriptor() ([]byte, []int) {
	return file_sf_bstream_v1_bstream_proto_rawDescGZIP(), []int{0}
}

func (x *BlockRequest) GetBurst() int64 {
	if x != nil {
		return x.Burst
	}
	return 0
}

func (x *BlockRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *BlockRequest) GetOrder() BlockRequest_Order {
	if x != nil {
		return x.Order
	}
	return BlockRequest_UNSPECIFIED
}

func (x *BlockRequest) GetRequester() string {
	if x != nil {
		return x.Requester
	}
	return ""
}

// Cursor is used to generate a string cursor, currently being utilized in forkable
type Cursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block     *BlockRef `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	HeadBlock *BlockRef `protobuf:"bytes,2,opt,name=head_block,json=headBlock,proto3" json:"head_block,omitempty"`
	Lib       *BlockRef `protobuf:"bytes,3,opt,name=lib,proto3" json:"lib,omitempty"`
	Step      ForkStep  `protobuf:"varint,4,opt,name=step,proto3,enum=sf.bstream.v1.ForkStep" json:"step,omitempty"`
}

func (x *Cursor) Reset() {
	*x = Cursor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_bstream_v1_bstream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cursor) ProtoMessage() {}

func (x *Cursor) ProtoReflect() protoreflect.Message {
	mi := &file_sf_bstream_v1_bstream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cursor.ProtoReflect.Descriptor instead.
func (*Cursor) Descriptor() ([]byte, []int) {
	return file_sf_bstream_v1_bstream_proto_rawDescGZIP(), []int{1}
}

func (x *Cursor) GetBlock() *BlockRef {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *Cursor) GetHeadBlock() *BlockRef {
	if x != nil {
		return x.HeadBlock
	}
	return nil
}

func (x *Cursor) GetLib() *BlockRef {
	if x != nil {
		return x.Lib
	}
	return nil
}

func (x *Cursor) GetStep() ForkStep {
	if x != nil {
		return x.Step
	}
	return ForkStep_STEP_UNKNOWN
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number         uint64                 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Id             string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	PreviousId     string                 `protobuf:"bytes,3,opt,name=previous_id,json=previousId,proto3" json:"previous_id,omitempty"`
	Timestamp      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	LibNum         uint64                 `protobuf:"varint,5,opt,name=lib_num,json=libNum,proto3" json:"lib_num,omitempty"`
	PayloadKind    Protocol               `protobuf:"varint,6,opt,name=payload_kind,json=payloadKind,proto3,enum=sf.bstream.v1.Protocol" json:"payload_kind,omitempty"`
	PayloadVersion int32                  `protobuf:"varint,7,opt,name=payload_version,json=payloadVersion,proto3" json:"payload_version,omitempty"`
	PayloadBuffer  []byte                 `protobuf:"bytes,8,opt,name=payload_buffer,json=payloadBuffer,proto3" json:"payload_buffer,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_bstream_v1_bstream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_sf_bstream_v1_bstream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_sf_bstream_v1_bstream_proto_rawDescGZIP(), []int{2}
}

func (x *Block) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Block) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Block) GetPreviousId() string {
	if x != nil {
		return x.PreviousId
	}
	return ""
}

func (x *Block) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Block) GetLibNum() uint64 {
	if x != nil {
		return x.LibNum
	}
	return 0
}

func (x *Block) GetPayloadKind() Protocol {
	if x != nil {
		return x.PayloadKind
	}
	return Protocol_UNKNOWN
}

func (x *Block) GetPayloadVersion() int32 {
	if x != nil {
		return x.PayloadVersion
	}
	return 0
}

func (x *Block) GetPayloadBuffer() []byte {
	if x != nil {
		return x.PayloadBuffer
	}
	return nil
}

type BlockRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num uint64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	Id  string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BlockRef) Reset() {
	*x = BlockRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_bstream_v1_bstream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRef) ProtoMessage() {}

func (x *BlockRef) ProtoReflect() protoreflect.Message {
	mi := &file_sf_bstream_v1_bstream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRef.ProtoReflect.Descriptor instead.
func (*BlockRef) Descriptor() ([]byte, []int) {
	return file_sf_bstream_v1_bstream_proto_rawDescGZIP(), []int{3}
}

func (x *BlockRef) GetNum() uint64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *BlockRef) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GenericBlockIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kv []*KeyToBitmap `protobuf:"bytes,4,rep,name=kv,proto3" json:"kv,omitempty"`
}

func (x *GenericBlockIndex) Reset() {
	*x = GenericBlockIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_bstream_v1_bstream_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericBlockIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericBlockIndex) ProtoMessage() {}

func (x *GenericBlockIndex) ProtoReflect() protoreflect.Message {
	mi := &file_sf_bstream_v1_bstream_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericBlockIndex.ProtoReflect.Descriptor instead.
func (*GenericBlockIndex) Descriptor() ([]byte, []int) {
	return file_sf_bstream_v1_bstream_proto_rawDescGZIP(), []int{4}
}

func (x *GenericBlockIndex) GetKv() []*KeyToBitmap {
	if x != nil {
		return x.Kv
	}
	return nil
}

type KeyToBitmap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Bitmap []byte `protobuf:"bytes,2,opt,name=bitmap,proto3" json:"bitmap,omitempty"`
}

func (x *KeyToBitmap) Reset() {
	*x = KeyToBitmap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_bstream_v1_bstream_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyToBitmap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyToBitmap) ProtoMessage() {}

func (x *KeyToBitmap) ProtoReflect() protoreflect.Message {
	mi := &file_sf_bstream_v1_bstream_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyToBitmap.ProtoReflect.Descriptor instead.
func (*KeyToBitmap) Descriptor() ([]byte, []int) {
	return file_sf_bstream_v1_bstream_proto_rawDescGZIP(), []int{5}
}

func (x *KeyToBitmap) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *KeyToBitmap) GetBitmap() []byte {
	if x != nil {
		return x.Bitmap
	}
	return nil
}

var File_sf_bstream_v1_bstream_proto protoreflect.FileDescriptor

var file_sf_bstream_v1_bstream_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x66, 0x2f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73,
	0x66, 0x2e, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01,
	0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x75, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62,
	0x75, 0x72, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x22, 0x34,
	0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x4f, 0x52, 0x44, 0x45, 0x52,
	0x45, 0x44, 0x10, 0x02, 0x22, 0xc7, 0x01, 0x0a, 0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12,
	0x2d, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x73, 0x66, 0x2e, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x36,
	0x0a, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66, 0x52, 0x09, 0x68, 0x65, 0x61,
	0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x29, 0x0a, 0x03, 0x6c, 0x69, 0x62, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66, 0x52, 0x03, 0x6c, 0x69,
	0x62, 0x12, 0x2b, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x17, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x65, 0x70, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0xaf,
	0x02, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x49,
	0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x6c,
	0x69, 0x62, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x69,
	0x62, 0x4e, 0x75, 0x6d, 0x12, 0x3a, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x66, 0x2e,
	0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4b, 0x69, 0x6e, 0x64,
	0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x22, 0x2c, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x51,
	0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x2a, 0x0a, 0x02, 0x6b, 0x76, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x4b, 0x65, 0x79, 0x54, 0x6f, 0x42, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x52, 0x02, 0x6b, 0x76, 0x4a,
	0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10,
	0x04, 0x22, 0x37, 0x0a, 0x0b, 0x4b, 0x65, 0x79, 0x54, 0x6f, 0x42, 0x69, 0x74, 0x6d, 0x61, 0x70,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x62, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x2a, 0x5c, 0x0a, 0x08, 0x46, 0x6f,
	0x72, 0x6b, 0x53, 0x74, 0x65, 0x70, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x45, 0x50,
	0x5f, 0x4e, 0x45, 0x57, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x55,
	0x4e, 0x44, 0x4f, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x49, 0x52,
	0x52, 0x45, 0x56, 0x45, 0x52, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x22, 0x04, 0x08, 0x03,
	0x10, 0x03, 0x22, 0x04, 0x08, 0x05, 0x10, 0x05, 0x2a, 0x4f, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x54,
	0x48, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x4f, 0x4c, 0x41, 0x4e, 0x41, 0x10, 0x03, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x45, 0x41, 0x52, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x45, 0x4e,
	0x44, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x54, 0x10, 0x05, 0x32, 0x4c, 0x0a, 0x0b, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x3d, 0x0a, 0x06, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x30, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x66,
	0x61, 0x73, 0x74, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x2f, 0x73, 0x66, 0x2f, 0x62, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_bstream_v1_bstream_proto_rawDescOnce sync.Once
	file_sf_bstream_v1_bstream_proto_rawDescData = file_sf_bstream_v1_bstream_proto_rawDesc
)

func file_sf_bstream_v1_bstream_proto_rawDescGZIP() []byte {
	file_sf_bstream_v1_bstream_proto_rawDescOnce.Do(func() {
		file_sf_bstream_v1_bstream_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_bstream_v1_bstream_proto_rawDescData)
	})
	return file_sf_bstream_v1_bstream_proto_rawDescData
}

var file_sf_bstream_v1_bstream_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_sf_bstream_v1_bstream_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sf_bstream_v1_bstream_proto_goTypes = []interface{}{
	(ForkStep)(0),                 // 0: sf.bstream.v1.ForkStep
	(Protocol)(0),                 // 1: sf.bstream.v1.Protocol
	(BlockRequest_Order)(0),       // 2: sf.bstream.v1.BlockRequest.Order
	(*BlockRequest)(nil),          // 3: sf.bstream.v1.BlockRequest
	(*Cursor)(nil),                // 4: sf.bstream.v1.Cursor
	(*Block)(nil),                 // 5: sf.bstream.v1.Block
	(*BlockRef)(nil),              // 6: sf.bstream.v1.BlockRef
	(*GenericBlockIndex)(nil),     // 7: sf.bstream.v1.GenericBlockIndex
	(*KeyToBitmap)(nil),           // 8: sf.bstream.v1.KeyToBitmap
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_sf_bstream_v1_bstream_proto_depIdxs = []int32{
	2, // 0: sf.bstream.v1.BlockRequest.order:type_name -> sf.bstream.v1.BlockRequest.Order
	6, // 1: sf.bstream.v1.Cursor.block:type_name -> sf.bstream.v1.BlockRef
	6, // 2: sf.bstream.v1.Cursor.head_block:type_name -> sf.bstream.v1.BlockRef
	6, // 3: sf.bstream.v1.Cursor.lib:type_name -> sf.bstream.v1.BlockRef
	0, // 4: sf.bstream.v1.Cursor.step:type_name -> sf.bstream.v1.ForkStep
	9, // 5: sf.bstream.v1.Block.timestamp:type_name -> google.protobuf.Timestamp
	1, // 6: sf.bstream.v1.Block.payload_kind:type_name -> sf.bstream.v1.Protocol
	8, // 7: sf.bstream.v1.GenericBlockIndex.kv:type_name -> sf.bstream.v1.KeyToBitmap
	3, // 8: sf.bstream.v1.BlockStream.Blocks:input_type -> sf.bstream.v1.BlockRequest
	5, // 9: sf.bstream.v1.BlockStream.Blocks:output_type -> sf.bstream.v1.Block
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_sf_bstream_v1_bstream_proto_init() }
func file_sf_bstream_v1_bstream_proto_init() {
	if File_sf_bstream_v1_bstream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_bstream_v1_bstream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRequest); i {
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
		file_sf_bstream_v1_bstream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cursor); i {
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
		file_sf_bstream_v1_bstream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_sf_bstream_v1_bstream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRef); i {
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
		file_sf_bstream_v1_bstream_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericBlockIndex); i {
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
		file_sf_bstream_v1_bstream_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyToBitmap); i {
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
			RawDescriptor: file_sf_bstream_v1_bstream_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sf_bstream_v1_bstream_proto_goTypes,
		DependencyIndexes: file_sf_bstream_v1_bstream_proto_depIdxs,
		EnumInfos:         file_sf_bstream_v1_bstream_proto_enumTypes,
		MessageInfos:      file_sf_bstream_v1_bstream_proto_msgTypes,
	}.Build()
	File_sf_bstream_v1_bstream_proto = out.File
	file_sf_bstream_v1_bstream_proto_rawDesc = nil
	file_sf_bstream_v1_bstream_proto_goTypes = nil
	file_sf_bstream_v1_bstream_proto_depIdxs = nil
}
