// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/bstream/v1/bstream.proto

package pbbstream

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ForkStep int32

const (
	ForkStep_STEP_UNKNOWN      ForkStep = 0
	ForkStep_STEP_NEW          ForkStep = 1
	ForkStep_STEP_UNDO         ForkStep = 2
	ForkStep_STEP_REDO         ForkStep = 3
	ForkStep_STEP_IRREVERSIBLE ForkStep = 4
	ForkStep_STEP_STALLED      ForkStep = 5
)

var ForkStep_name = map[int32]string{
	0: "STEP_UNKNOWN",
	1: "STEP_NEW",
	2: "STEP_UNDO",
	3: "STEP_REDO",
	4: "STEP_IRREVERSIBLE",
	5: "STEP_STALLED",
}

var ForkStep_value = map[string]int32{
	"STEP_UNKNOWN":      0,
	"STEP_NEW":          1,
	"STEP_UNDO":         2,
	"STEP_REDO":         3,
	"STEP_IRREVERSIBLE": 4,
	"STEP_STALLED":      5,
}

func (x ForkStep) String() string {
	return proto.EnumName(ForkStep_name, int32(x))
}

func (ForkStep) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9cfb17be2e7d1b7d, []int{0}
}

type Protocol int32

const (
	Protocol_UNKNOWN Protocol = 0
	Protocol_EOS     Protocol = 1
	Protocol_ETH     Protocol = 2
)

var Protocol_name = map[int32]string{
	0: "UNKNOWN",
	1: "EOS",
	2: "ETH",
}

var Protocol_value = map[string]int32{
	"UNKNOWN": 0,
	"EOS":     1,
	"ETH":     2,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}

func (Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9cfb17be2e7d1b7d, []int{1}
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

var BlockRequest_Order_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "ORDERED",
	2: "UNORDERED",
}

var BlockRequest_Order_value = map[string]int32{
	"UNSPECIFIED": 0,
	"ORDERED":     1,
	"UNORDERED":   2,
}

func (x BlockRequest_Order) String() string {
	return proto.EnumName(BlockRequest_Order_name, int32(x))
}

func (BlockRequest_Order) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9cfb17be2e7d1b7d, []int{0, 0}
}

type BlockRequest struct {
	// Number of blocks we want to get in burst upon connection, on a best effort basis.
	Burst int64 `protobuf:"varint,1,opt,name=burst,proto3" json:"burst,omitempty"`
	// Type of blocks we're after here, is it 'ethereum' data, 'eos', etc.. The server can fail early
	// if he doesn't match the data he serves (services mismatch, etc..)
	ContentType          string             `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Order                BlockRequest_Order `protobuf:"varint,3,opt,name=order,proto3,enum=dfuse.bstream.v1.BlockRequest_Order" json:"order,omitempty"`
	Requester            string             `protobuf:"bytes,4,opt,name=requester,proto3" json:"requester,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BlockRequest) Reset()         { *m = BlockRequest{} }
func (m *BlockRequest) String() string { return proto.CompactTextString(m) }
func (*BlockRequest) ProtoMessage()    {}
func (*BlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cfb17be2e7d1b7d, []int{0}
}

func (m *BlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockRequest.Unmarshal(m, b)
}
func (m *BlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockRequest.Marshal(b, m, deterministic)
}
func (m *BlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockRequest.Merge(m, src)
}
func (m *BlockRequest) XXX_Size() int {
	return xxx_messageInfo_BlockRequest.Size(m)
}
func (m *BlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlockRequest proto.InternalMessageInfo

func (m *BlockRequest) GetBurst() int64 {
	if m != nil {
		return m.Burst
	}
	return 0
}

func (m *BlockRequest) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *BlockRequest) GetOrder() BlockRequest_Order {
	if m != nil {
		return m.Order
	}
	return BlockRequest_UNSPECIFIED
}

func (m *BlockRequest) GetRequester() string {
	if m != nil {
		return m.Requester
	}
	return ""
}

type BlocksRequestV2 struct {
	StartBlockNum       int64      `protobuf:"varint,1,opt,name=start_block_num,json=startBlockNum,proto3" json:"start_block_num,omitempty"`
	StartBlockId        uint64     `protobuf:"varint,2,opt,name=start_block_id,json=startBlockId,proto3" json:"start_block_id,omitempty"`
	ExcludeStartBlock   bool       `protobuf:"varint,3,opt,name=exclude_start_block,json=excludeStartBlock,proto3" json:"exclude_start_block,omitempty"`
	StreamBlocksFromLib bool       `protobuf:"varint,4,opt,name=stream_blocks_from_lib,json=streamBlocksFromLib,proto3" json:"stream_blocks_from_lib,omitempty"`
	StopBlockNum        uint64     `protobuf:"varint,5,opt,name=stop_block_num,json=stopBlockNum,proto3" json:"stop_block_num,omitempty"`
	ExcludeStopBlock    bool       `protobuf:"varint,6,opt,name=exclude_stop_block,json=excludeStopBlock,proto3" json:"exclude_stop_block,omitempty"`
	HandleForks         bool       `protobuf:"varint,7,opt,name=handleForks,proto3" json:"handleForks,omitempty"`
	HandleForksSteps    []ForkStep `protobuf:"varint,8,rep,packed,name=handle_forks_steps,json=handleForksSteps,proto3,enum=dfuse.bstream.v1.ForkStep" json:"handle_forks_steps,omitempty"`
	Decoded             bool       `protobuf:"varint,9,opt,name=decoded,proto3" json:"decoded,omitempty"`
	// The CEL filter expression used to include transactions, specific to the target protocol, works
	// in combination with `exclude_filter_expr` value.
	IncludeFilterExpr string `protobuf:"bytes,10,opt,name=include_filter_expr,json=includeFilterExpr,proto3" json:"include_filter_expr,omitempty"`
	// The CEL filter expression used to exclude transactions, specific to the target protocol, works
	// in combination with `include_filter_expr` value.
	ExcludeFilterExpr    string   `protobuf:"bytes,11,opt,name=exclude_filter_expr,json=excludeFilterExpr,proto3" json:"exclude_filter_expr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlocksRequestV2) Reset()         { *m = BlocksRequestV2{} }
func (m *BlocksRequestV2) String() string { return proto.CompactTextString(m) }
func (*BlocksRequestV2) ProtoMessage()    {}
func (*BlocksRequestV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cfb17be2e7d1b7d, []int{1}
}

func (m *BlocksRequestV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlocksRequestV2.Unmarshal(m, b)
}
func (m *BlocksRequestV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlocksRequestV2.Marshal(b, m, deterministic)
}
func (m *BlocksRequestV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlocksRequestV2.Merge(m, src)
}
func (m *BlocksRequestV2) XXX_Size() int {
	return xxx_messageInfo_BlocksRequestV2.Size(m)
}
func (m *BlocksRequestV2) XXX_DiscardUnknown() {
	xxx_messageInfo_BlocksRequestV2.DiscardUnknown(m)
}

var xxx_messageInfo_BlocksRequestV2 proto.InternalMessageInfo

func (m *BlocksRequestV2) GetStartBlockNum() int64 {
	if m != nil {
		return m.StartBlockNum
	}
	return 0
}

func (m *BlocksRequestV2) GetStartBlockId() uint64 {
	if m != nil {
		return m.StartBlockId
	}
	return 0
}

func (m *BlocksRequestV2) GetExcludeStartBlock() bool {
	if m != nil {
		return m.ExcludeStartBlock
	}
	return false
}

func (m *BlocksRequestV2) GetStreamBlocksFromLib() bool {
	if m != nil {
		return m.StreamBlocksFromLib
	}
	return false
}

func (m *BlocksRequestV2) GetStopBlockNum() uint64 {
	if m != nil {
		return m.StopBlockNum
	}
	return 0
}

func (m *BlocksRequestV2) GetExcludeStopBlock() bool {
	if m != nil {
		return m.ExcludeStopBlock
	}
	return false
}

func (m *BlocksRequestV2) GetHandleForks() bool {
	if m != nil {
		return m.HandleForks
	}
	return false
}

func (m *BlocksRequestV2) GetHandleForksSteps() []ForkStep {
	if m != nil {
		return m.HandleForksSteps
	}
	return nil
}

func (m *BlocksRequestV2) GetDecoded() bool {
	if m != nil {
		return m.Decoded
	}
	return false
}

func (m *BlocksRequestV2) GetIncludeFilterExpr() string {
	if m != nil {
		return m.IncludeFilterExpr
	}
	return ""
}

func (m *BlocksRequestV2) GetExcludeFilterExpr() string {
	if m != nil {
		return m.ExcludeFilterExpr
	}
	return ""
}

type BlockResponseV2 struct {
	// Chain specific blocks:
	// * dfuse.ethereum.codec.v1.Block
	// * dfuse.eosio.codec.v1.Block
	// Chain agnostic block, when request asks for packed blocks:
	// * dfuse.bstream.v1.Block
	Block                 *any.Any `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Undo                  bool     `protobuf:"varint,5,opt,name=undo,proto3" json:"undo,omitempty"`
	Step                  ForkStep `protobuf:"varint,6,opt,name=step,proto3,enum=dfuse.bstream.v1.ForkStep" json:"step,omitempty"`
	StepCount             uint64   `protobuf:"varint,7,opt,name=step_count,json=stepCount,proto3" json:"step_count,omitempty"`
	StepIndex             uint64   `protobuf:"varint,8,opt,name=step_index,json=stepIndex,proto3" json:"step_index,omitempty"`
	KnownToBeIrreversible bool     `protobuf:"varint,9,opt,name=known_to_be_irreversible,json=knownToBeIrreversible,proto3" json:"known_to_be_irreversible,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *BlockResponseV2) Reset()         { *m = BlockResponseV2{} }
func (m *BlockResponseV2) String() string { return proto.CompactTextString(m) }
func (*BlockResponseV2) ProtoMessage()    {}
func (*BlockResponseV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cfb17be2e7d1b7d, []int{2}
}

func (m *BlockResponseV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockResponseV2.Unmarshal(m, b)
}
func (m *BlockResponseV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockResponseV2.Marshal(b, m, deterministic)
}
func (m *BlockResponseV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockResponseV2.Merge(m, src)
}
func (m *BlockResponseV2) XXX_Size() int {
	return xxx_messageInfo_BlockResponseV2.Size(m)
}
func (m *BlockResponseV2) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockResponseV2.DiscardUnknown(m)
}

var xxx_messageInfo_BlockResponseV2 proto.InternalMessageInfo

func (m *BlockResponseV2) GetBlock() *any.Any {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *BlockResponseV2) GetUndo() bool {
	if m != nil {
		return m.Undo
	}
	return false
}

func (m *BlockResponseV2) GetStep() ForkStep {
	if m != nil {
		return m.Step
	}
	return ForkStep_STEP_UNKNOWN
}

func (m *BlockResponseV2) GetStepCount() uint64 {
	if m != nil {
		return m.StepCount
	}
	return 0
}

func (m *BlockResponseV2) GetStepIndex() uint64 {
	if m != nil {
		return m.StepIndex
	}
	return 0
}

func (m *BlockResponseV2) GetKnownToBeIrreversible() bool {
	if m != nil {
		return m.KnownToBeIrreversible
	}
	return false
}

type Block struct {
	Number               uint64               `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Id                   string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	PreviousId           string               `protobuf:"bytes,3,opt,name=previous_id,json=previousId,proto3" json:"previous_id,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	LibNum               uint64               `protobuf:"varint,5,opt,name=lib_num,json=libNum,proto3" json:"lib_num,omitempty"`
	PayloadKind          Protocol             `protobuf:"varint,6,opt,name=payload_kind,json=payloadKind,proto3,enum=dfuse.bstream.v1.Protocol" json:"payload_kind,omitempty"`
	PayloadVersion       int32                `protobuf:"varint,7,opt,name=payload_version,json=payloadVersion,proto3" json:"payload_version,omitempty"`
	PayloadBuffer        []byte               `protobuf:"bytes,8,opt,name=payload_buffer,json=payloadBuffer,proto3" json:"payload_buffer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cfb17be2e7d1b7d, []int{3}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Block) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Block) GetPreviousId() string {
	if m != nil {
		return m.PreviousId
	}
	return ""
}

func (m *Block) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Block) GetLibNum() uint64 {
	if m != nil {
		return m.LibNum
	}
	return 0
}

func (m *Block) GetPayloadKind() Protocol {
	if m != nil {
		return m.PayloadKind
	}
	return Protocol_UNKNOWN
}

func (m *Block) GetPayloadVersion() int32 {
	if m != nil {
		return m.PayloadVersion
	}
	return 0
}

func (m *Block) GetPayloadBuffer() []byte {
	if m != nil {
		return m.PayloadBuffer
	}
	return nil
}

type BlockRef struct {
	Num                  uint64   `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockRef) Reset()         { *m = BlockRef{} }
func (m *BlockRef) String() string { return proto.CompactTextString(m) }
func (*BlockRef) ProtoMessage()    {}
func (*BlockRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cfb17be2e7d1b7d, []int{4}
}

func (m *BlockRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockRef.Unmarshal(m, b)
}
func (m *BlockRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockRef.Marshal(b, m, deterministic)
}
func (m *BlockRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockRef.Merge(m, src)
}
func (m *BlockRef) XXX_Size() int {
	return xxx_messageInfo_BlockRef.Size(m)
}
func (m *BlockRef) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockRef.DiscardUnknown(m)
}

var xxx_messageInfo_BlockRef proto.InternalMessageInfo

func (m *BlockRef) GetNum() uint64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *BlockRef) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterEnum("dfuse.bstream.v1.ForkStep", ForkStep_name, ForkStep_value)
	proto.RegisterEnum("dfuse.bstream.v1.Protocol", Protocol_name, Protocol_value)
	proto.RegisterEnum("dfuse.bstream.v1.BlockRequest_Order", BlockRequest_Order_name, BlockRequest_Order_value)
	proto.RegisterType((*BlockRequest)(nil), "dfuse.bstream.v1.BlockRequest")
	proto.RegisterType((*BlocksRequestV2)(nil), "dfuse.bstream.v1.BlocksRequestV2")
	proto.RegisterType((*BlockResponseV2)(nil), "dfuse.bstream.v1.BlockResponseV2")
	proto.RegisterType((*Block)(nil), "dfuse.bstream.v1.Block")
	proto.RegisterType((*BlockRef)(nil), "dfuse.bstream.v1.BlockRef")
}

func init() { proto.RegisterFile("dfuse/bstream/v1/bstream.proto", fileDescriptor_9cfb17be2e7d1b7d) }

var fileDescriptor_9cfb17be2e7d1b7d = []byte{
	// 928 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xdb, 0x6e, 0xdb, 0x46,
	0x10, 0x0d, 0x75, 0xd7, 0x50, 0x96, 0xe9, 0xc9, 0x8d, 0x15, 0xda, 0x44, 0x16, 0xd2, 0xd6, 0x35,
	0x52, 0x2a, 0x91, 0x5b, 0xb4, 0x68, 0xd1, 0x87, 0xc8, 0xa2, 0x11, 0x21, 0x86, 0x64, 0xac, 0x64,
	0x05, 0xe8, 0x0b, 0x21, 0x8a, 0x2b, 0x87, 0x30, 0xc9, 0x65, 0x96, 0xa4, 0x6b, 0xfd, 0x56, 0x7f,
	0xa6, 0xbf, 0xd2, 0xa7, 0xa2, 0xe0, 0x2e, 0x29, 0x11, 0x0e, 0xec, 0x27, 0xed, 0x9c, 0x73, 0x66,
	0x77, 0xf6, 0xec, 0x70, 0x04, 0x2f, 0x9c, 0x75, 0x12, 0xd1, 0xbe, 0x1d, 0xc5, 0x9c, 0x2e, 0xfd,
	0xfe, 0xcd, 0xdb, 0x7c, 0x69, 0x84, 0x9c, 0xc5, 0x0c, 0x35, 0xc1, 0x1b, 0x39, 0x78, 0xf3, 0xb6,
	0xf3, 0xf2, 0x8a, 0xb1, 0x2b, 0x8f, 0xf6, 0x05, 0x6f, 0x27, 0xeb, 0x7e, 0xec, 0xfa, 0x34, 0x8a,
	0x97, 0x7e, 0x28, 0x53, 0x3a, 0x5f, 0xdd, 0x15, 0x2c, 0x83, 0x8d, 0xa4, 0x7a, 0xff, 0x28, 0xd0,
	0x1a, 0x7a, 0x6c, 0x75, 0x4d, 0xe8, 0xe7, 0x84, 0x46, 0x31, 0x3e, 0x81, 0xaa, 0x9d, 0xf0, 0x28,
	0xd6, 0x95, 0xae, 0x72, 0x54, 0x26, 0x32, 0xc0, 0x43, 0x68, 0xad, 0x58, 0x10, 0xd3, 0x20, 0xb6,
	0xe2, 0x4d, 0x48, 0xf5, 0x52, 0x57, 0x39, 0x6a, 0x12, 0x35, 0xc3, 0xe6, 0x9b, 0x90, 0xe2, 0x6f,
	0x50, 0x65, 0xdc, 0xa1, 0x5c, 0x2f, 0x77, 0x95, 0xa3, 0xf6, 0xe0, 0x95, 0x71, 0xb7, 0x4e, 0xa3,
	0x78, 0x8e, 0x31, 0x4d, 0xb5, 0x44, 0xa6, 0xe0, 0xd7, 0xd0, 0xe4, 0x12, 0xa7, 0x5c, 0xaf, 0x88,
	0xbd, 0x77, 0x40, 0xef, 0x27, 0xa8, 0x0a, 0x35, 0xee, 0x83, 0x7a, 0x39, 0x99, 0x5d, 0x98, 0xa7,
	0xe3, 0xb3, 0xb1, 0x39, 0xd2, 0x1e, 0xa1, 0x0a, 0xf5, 0x29, 0x19, 0x99, 0xc4, 0x1c, 0x69, 0x0a,
	0xee, 0x41, 0xf3, 0x72, 0x92, 0x87, 0xa5, 0xde, 0x7f, 0x65, 0xd8, 0x17, 0x27, 0x46, 0xd9, 0x91,
	0x8b, 0x01, 0x7e, 0x07, 0xfb, 0x51, 0xbc, 0xe4, 0xb1, 0x65, 0xa7, 0x84, 0x15, 0x24, 0x7e, 0x76,
	0xcd, 0x3d, 0x01, 0x0b, 0xf9, 0x24, 0xf1, 0xf1, 0x15, 0xb4, 0x8b, 0x3a, 0xd7, 0x11, 0x17, 0xae,
	0x90, 0xd6, 0x4e, 0x36, 0x76, 0xd0, 0x80, 0xc7, 0xf4, 0x76, 0xe5, 0x25, 0x0e, 0xb5, 0x0a, 0x6a,
	0x71, 0xff, 0x06, 0x39, 0xc8, 0xa8, 0xd9, 0x36, 0x03, 0x4f, 0xe0, 0x99, 0x34, 0x43, 0x0a, 0x23,
	0x6b, 0xcd, 0x99, 0x6f, 0x79, 0xae, 0x2d, 0xae, 0xdc, 0x20, 0x8f, 0x25, 0x2b, 0x8b, 0x3e, 0xe3,
	0xcc, 0x3f, 0x77, 0x6d, 0x59, 0x0a, 0x0b, 0x0b, 0x15, 0x57, 0xf3, 0x52, 0x58, 0xb8, 0x2d, 0xf8,
	0x35, 0xe0, 0xae, 0x94, 0x5c, 0xad, 0xd7, 0xc4, 0xb6, 0xda, 0xb6, 0x92, 0x2c, 0x01, 0xbb, 0xa0,
	0x7e, 0x5a, 0x06, 0x8e, 0x47, 0xcf, 0x18, 0xbf, 0x8e, 0xf4, 0xba, 0x90, 0x15, 0x21, 0x7c, 0x0f,
	0x28, 0x43, 0x6b, 0x9d, 0xc6, 0x56, 0x14, 0xd3, 0x30, 0xd2, 0x1b, 0xdd, 0xf2, 0x51, 0x7b, 0xd0,
	0xf9, 0xf2, 0x65, 0xd3, 0xa4, 0x59, 0x4c, 0x43, 0xa2, 0x15, 0x36, 0x49, 0x81, 0x08, 0x75, 0xa8,
	0x3b, 0x74, 0xc5, 0x1c, 0xea, 0xe8, 0x4d, 0x71, 0x4e, 0x1e, 0xa6, 0xf6, 0xb9, 0x81, 0xac, 0x79,
	0xed, 0x7a, 0x31, 0xe5, 0x16, 0xbd, 0x0d, 0xb9, 0x0e, 0xe2, 0xf9, 0x0f, 0x32, 0xea, 0x4c, 0x30,
	0xe6, 0x6d, 0xc8, 0x8b, 0x76, 0x17, 0xf5, 0xaa, 0xd4, 0x67, 0xd4, 0x4e, 0xdf, 0xfb, 0x57, 0xc9,
	0x1a, 0x80, 0xd0, 0x28, 0x64, 0x41, 0x44, 0x17, 0x03, 0x3c, 0x86, 0xaa, 0xb4, 0x26, 0x7d, 0x76,
	0x75, 0xf0, 0xc4, 0x90, 0x5f, 0x86, 0x91, 0x7f, 0x19, 0xc6, 0xbb, 0x60, 0x43, 0xa4, 0x04, 0x11,
	0x2a, 0x49, 0xe0, 0x30, 0xe1, 0x77, 0x83, 0x88, 0x35, 0x1a, 0x50, 0x49, 0xad, 0x10, 0xce, 0x3e,
	0xec, 0x84, 0xd0, 0xe1, 0x37, 0x00, 0xe9, 0xaf, 0xb5, 0x62, 0x49, 0x10, 0x0b, 0xa3, 0x2b, 0xa4,
	0x99, 0x22, 0xa7, 0x29, 0xb0, 0xa5, 0xdd, 0xc0, 0xa1, 0xb7, 0x7a, 0x63, 0x47, 0x8f, 0x53, 0x00,
	0x7f, 0x01, 0xfd, 0x3a, 0x60, 0x7f, 0x05, 0x56, 0xcc, 0x2c, 0x9b, 0x5a, 0x2e, 0xe7, 0xf4, 0x86,
	0xf2, 0xc8, 0xb5, 0x3d, 0x9a, 0x99, 0xf9, 0x54, 0xf0, 0x73, 0x36, 0xa4, 0xe3, 0x02, 0xd9, 0xfb,
	0xbb, 0x04, 0x55, 0xf9, 0xd4, 0xcf, 0xa0, 0x16, 0x24, 0xbe, 0x4d, 0xb9, 0xb8, 0x71, 0x85, 0x64,
	0x11, 0xb6, 0xa1, 0x94, 0x75, 0x75, 0x93, 0x94, 0x5c, 0x07, 0x5f, 0x82, 0x1a, 0x72, 0x7a, 0xe3,
	0xb2, 0x24, 0x4a, 0xdb, 0xbd, 0x2c, 0x08, 0xc8, 0xa1, 0xb1, 0x83, 0xbf, 0x42, 0x73, 0x3b, 0x56,
	0x44, 0xbf, 0xaa, 0x83, 0xce, 0x17, 0xee, 0xcd, 0x73, 0x05, 0xd9, 0x89, 0xf1, 0x39, 0xd4, 0x3d,
	0xd7, 0x2e, 0xb4, 0x6e, 0xcd, 0x73, 0xed, 0xb4, 0x69, 0xff, 0x80, 0x56, 0xb8, 0xdc, 0x78, 0x6c,
	0xe9, 0x58, 0xd7, 0x6e, 0xe0, 0xdc, 0x6f, 0xea, 0x45, 0xba, 0xff, 0x8a, 0x79, 0x44, 0xcd, 0xf4,
	0x1f, 0xdc, 0xc0, 0xc1, 0xef, 0x61, 0x3f, 0x4f, 0x17, 0x17, 0x67, 0x81, 0x30, 0xb8, 0x4a, 0xda,
	0x19, 0xbc, 0x90, 0x28, 0x7e, 0x0b, 0x39, 0x62, 0xd9, 0xc9, 0x7a, 0x4d, 0xb9, 0x70, 0xba, 0x45,
	0xf6, 0x32, 0x74, 0x28, 0xc0, 0xde, 0x6b, 0x68, 0x64, 0xed, 0xb2, 0x46, 0x0d, 0xca, 0xf9, 0x70,
	0xa8, 0x90, 0x74, 0x79, 0xd7, 0xb0, 0xe3, 0xcf, 0xd0, 0xc8, 0xdf, 0x1a, 0x35, 0x68, 0xcd, 0xe6,
	0xe6, 0x85, 0x75, 0x39, 0xf9, 0x30, 0x99, 0x7e, 0x9c, 0x68, 0x8f, 0xb0, 0x05, 0x0d, 0x81, 0x4c,
	0xcc, 0x8f, 0x72, 0x32, 0x65, 0xfc, 0x68, 0xaa, 0x95, 0xb6, 0x21, 0x31, 0x47, 0x53, 0xad, 0x8c,
	0x4f, 0xe1, 0x40, 0x84, 0x63, 0x42, 0xcc, 0x85, 0x49, 0x66, 0xe3, 0xe1, 0xb9, 0xa9, 0x55, 0xb6,
	0x9b, 0xce, 0xe6, 0xef, 0xce, 0xcf, 0xcd, 0x91, 0x56, 0x3d, 0xfe, 0x01, 0x1a, 0xb9, 0x13, 0xe9,
	0xe4, 0xdb, 0x9d, 0x56, 0x87, 0xb2, 0x39, 0x9d, 0x69, 0x8a, 0x58, 0xcc, 0xdf, 0x6b, 0xa5, 0x01,
	0x01, 0x55, 0xdc, 0x65, 0x26, 0x3c, 0xc4, 0x53, 0xa8, 0xc9, 0xa9, 0x82, 0x2f, 0x1e, 0x1e, 0xcb,
	0x9d, 0xe7, 0xf7, 0xf0, 0x6f, 0x94, 0xc1, 0x12, 0xf6, 0x0a, 0x7b, 0x2e, 0x06, 0x78, 0xb1, 0xdd,
	0xf5, 0xf0, 0x9e, 0xac, 0xdd, 0xe8, 0xed, 0x1c, 0xde, 0x7b, 0x70, 0xfe, 0x71, 0xbe, 0x51, 0x86,
	0x3f, 0xff, 0x79, 0x72, 0xe5, 0xc6, 0x9f, 0x12, 0xdb, 0x58, 0x31, 0xbf, 0x2f, 0x12, 0x7e, 0x74,
	0x59, 0x3f, 0xb4, 0xaf, 0x58, 0xff, 0xee, 0xdf, 0xe2, 0xef, 0xa1, 0x9d, 0x05, 0x76, 0x4d, 0x34,
	0xe0, 0xc9, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xc4, 0x2f, 0xd8, 0x3b, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlockStreamClient is the client API for BlockStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockStreamClient interface {
	Blocks(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (BlockStream_BlocksClient, error)
}

type blockStreamClient struct {
	cc *grpc.ClientConn
}

func NewBlockStreamClient(cc *grpc.ClientConn) BlockStreamClient {
	return &blockStreamClient{cc}
}

func (c *blockStreamClient) Blocks(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (BlockStream_BlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlockStream_serviceDesc.Streams[0], "/dfuse.bstream.v1.BlockStream/Blocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockStreamBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockStream_BlocksClient interface {
	Recv() (*Block, error)
	grpc.ClientStream
}

type blockStreamBlocksClient struct {
	grpc.ClientStream
}

func (x *blockStreamBlocksClient) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlockStreamServer is the server API for BlockStream service.
type BlockStreamServer interface {
	Blocks(*BlockRequest, BlockStream_BlocksServer) error
}

// UnimplementedBlockStreamServer can be embedded to have forward compatible implementations.
type UnimplementedBlockStreamServer struct {
}

func (*UnimplementedBlockStreamServer) Blocks(req *BlockRequest, srv BlockStream_BlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method Blocks not implemented")
}

func RegisterBlockStreamServer(s *grpc.Server, srv BlockStreamServer) {
	s.RegisterService(&_BlockStream_serviceDesc, srv)
}

func _BlockStream_Blocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockStreamServer).Blocks(m, &blockStreamBlocksServer{stream})
}

type BlockStream_BlocksServer interface {
	Send(*Block) error
	grpc.ServerStream
}

type blockStreamBlocksServer struct {
	grpc.ServerStream
}

func (x *blockStreamBlocksServer) Send(m *Block) error {
	return x.ServerStream.SendMsg(m)
}

var _BlockStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.bstream.v1.BlockStream",
	HandlerType: (*BlockStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Blocks",
			Handler:       _BlockStream_Blocks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dfuse/bstream/v1/bstream.proto",
}

// BlockStreamV2Client is the client API for BlockStreamV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockStreamV2Client interface {
	Blocks(ctx context.Context, in *BlocksRequestV2, opts ...grpc.CallOption) (BlockStreamV2_BlocksClient, error)
}

type blockStreamV2Client struct {
	cc *grpc.ClientConn
}

func NewBlockStreamV2Client(cc *grpc.ClientConn) BlockStreamV2Client {
	return &blockStreamV2Client{cc}
}

func (c *blockStreamV2Client) Blocks(ctx context.Context, in *BlocksRequestV2, opts ...grpc.CallOption) (BlockStreamV2_BlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlockStreamV2_serviceDesc.Streams[0], "/dfuse.bstream.v1.BlockStreamV2/Blocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockStreamV2BlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockStreamV2_BlocksClient interface {
	Recv() (*BlockResponseV2, error)
	grpc.ClientStream
}

type blockStreamV2BlocksClient struct {
	grpc.ClientStream
}

func (x *blockStreamV2BlocksClient) Recv() (*BlockResponseV2, error) {
	m := new(BlockResponseV2)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlockStreamV2Server is the server API for BlockStreamV2 service.
type BlockStreamV2Server interface {
	Blocks(*BlocksRequestV2, BlockStreamV2_BlocksServer) error
}

// UnimplementedBlockStreamV2Server can be embedded to have forward compatible implementations.
type UnimplementedBlockStreamV2Server struct {
}

func (*UnimplementedBlockStreamV2Server) Blocks(req *BlocksRequestV2, srv BlockStreamV2_BlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method Blocks not implemented")
}

func RegisterBlockStreamV2Server(s *grpc.Server, srv BlockStreamV2Server) {
	s.RegisterService(&_BlockStreamV2_serviceDesc, srv)
}

func _BlockStreamV2_Blocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlocksRequestV2)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockStreamV2Server).Blocks(m, &blockStreamV2BlocksServer{stream})
}

type BlockStreamV2_BlocksServer interface {
	Send(*BlockResponseV2) error
	grpc.ServerStream
}

type blockStreamV2BlocksServer struct {
	grpc.ServerStream
}

func (x *blockStreamV2BlocksServer) Send(m *BlockResponseV2) error {
	return x.ServerStream.SendMsg(m)
}

var _BlockStreamV2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.bstream.v1.BlockStreamV2",
	HandlerType: (*BlockStreamV2Server)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Blocks",
			Handler:       _BlockStreamV2_Blocks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dfuse/bstream/v1/bstream.proto",
}
