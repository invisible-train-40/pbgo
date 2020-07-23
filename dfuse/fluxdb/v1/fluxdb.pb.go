// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/fluxdb/v1/fluxdb.proto

package pbfluxdb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Index struct {
	SquelchedCount       uint64            `protobuf:"varint,1,opt,name=squelched_count,json=squelchedCount,proto3" json:"squelched_count,omitempty"`
	PrimaryKeyToHeight   map[string]uint64 `protobuf:"bytes,2,rep,name=primary_key_to_height,json=primaryKeyToHeight,proto3" json:"primary_key_to_height,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_74afbad2e99b142c, []int{0}
}

func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetSquelchedCount() uint64 {
	if m != nil {
		return m.SquelchedCount
	}
	return 0
}

func (m *Index) GetPrimaryKeyToHeight() map[string]uint64 {
	if m != nil {
		return m.PrimaryKeyToHeight
	}
	return nil
}

// Row represents a FluxDB storage row as seen by the underlying storage
// engine which is a Key-Value store.
//
// A FluxDB row is always of the form:
// ```
// <collection>/<tablet>/<height>/<primaryKey> => <payload>
// ```
type Row struct {
	Collection           string   `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	TabletKey            string   `protobuf:"bytes,2,opt,name=tablet_key,json=tabletKey,proto3" json:"tablet_key,omitempty"`
	HeightKey            string   `protobuf:"bytes,3,opt,name=height_key,json=heightKey,proto3" json:"height_key,omitempty"`
	PrimKey              string   `protobuf:"bytes,4,opt,name=prim_key,json=primKey,proto3" json:"prim_key,omitempty"`
	Payload              []byte   `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_74afbad2e99b142c, []int{1}
}

func (m *Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Row.Unmarshal(m, b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Row.Marshal(b, m, deterministic)
}
func (m *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(m, src)
}
func (m *Row) XXX_Size() int {
	return xxx_messageInfo_Row.Size(m)
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *Row) GetTabletKey() string {
	if m != nil {
		return m.TabletKey
	}
	return ""
}

func (m *Row) GetHeightKey() string {
	if m != nil {
		return m.HeightKey
	}
	return ""
}

func (m *Row) GetPrimKey() string {
	if m != nil {
		return m.PrimKey
	}
	return ""
}

func (m *Row) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Index)(nil), "dfuse.fluxdb.v1.Index")
	proto.RegisterMapType((map[string]uint64)(nil), "dfuse.fluxdb.v1.Index.PrimaryKeyToHeightEntry")
	proto.RegisterType((*Row)(nil), "dfuse.fluxdb.v1.Row")
}

func init() { proto.RegisterFile("dfuse/fluxdb/v1/fluxdb.proto", fileDescriptor_74afbad2e99b142c) }

var fileDescriptor_74afbad2e99b142c = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4e, 0xf2, 0x40,
	0x10, 0xc0, 0x53, 0x0a, 0x1f, 0x1f, 0xa3, 0x11, 0xb3, 0xd1, 0x58, 0x8d, 0x1a, 0xc2, 0x45, 0x2e,
	0x6e, 0x83, 0x5c, 0x8c, 0xde, 0x34, 0x24, 0x1a, 0x2e, 0xa6, 0xf1, 0xe4, 0xa5, 0xe9, 0x9f, 0x81,
	0x36, 0x2c, 0x9d, 0x5a, 0xb6, 0xc8, 0xbe, 0x89, 0xcf, 0xe6, 0xd3, 0x98, 0xdd, 0x45, 0x63, 0x30,
	0xde, 0x66, 0x7e, 0xbf, 0xce, 0x74, 0x66, 0x07, 0x4e, 0xd3, 0x69, 0xbd, 0x44, 0x7f, 0x2a, 0xea,
	0x75, 0x1a, 0xfb, 0xab, 0xe1, 0x26, 0xe2, 0x65, 0x45, 0x92, 0x58, 0xd7, 0x58, 0xbe, 0x61, 0xab,
	0x61, 0xff, 0xc3, 0x81, 0xd6, 0x63, 0x91, 0xe2, 0x9a, 0x5d, 0x40, 0x77, 0xf9, 0x5a, 0xa3, 0x48,
	0x32, 0x4c, 0xc3, 0x84, 0xea, 0x42, 0x7a, 0x4e, 0xcf, 0x19, 0x34, 0x83, 0xbd, 0x6f, 0x7c, 0xaf,
	0x29, 0x8b, 0xe0, 0xb0, 0xac, 0xf2, 0x45, 0x54, 0xa9, 0x70, 0x8e, 0x2a, 0x94, 0x14, 0x66, 0x98,
	0xcf, 0x32, 0xe9, 0x35, 0x7a, 0xee, 0x60, 0xe7, 0x8a, 0xf3, 0xad, 0x7f, 0x70, 0xd3, 0x9f, 0x3f,
	0xd9, 0x9a, 0x09, 0xaa, 0x67, 0x7a, 0x30, 0x05, 0xe3, 0x42, 0x56, 0x2a, 0x60, 0xe5, 0x2f, 0x71,
	0x32, 0x86, 0xa3, 0x3f, 0x3e, 0x67, 0xfb, 0xe0, 0xce, 0x51, 0x99, 0xd1, 0x3a, 0x81, 0x0e, 0xd9,
	0x01, 0xb4, 0x56, 0x91, 0xa8, 0xd1, 0x6b, 0x98, 0x71, 0x6d, 0x72, 0xd3, 0xb8, 0x76, 0xfa, 0xef,
	0x0e, 0xb8, 0x01, 0xbd, 0xb1, 0x73, 0x80, 0x84, 0x84, 0xc0, 0x44, 0xe6, 0x54, 0x6c, 0x4a, 0x7f,
	0x10, 0x76, 0x06, 0x20, 0xa3, 0x58, 0xa0, 0xd4, 0x0b, 0x99, 0x36, 0x9d, 0xa0, 0x63, 0xc9, 0x04,
	0x95, 0xd6, 0x76, 0x43, 0xa3, 0x5d, 0xab, 0x2d, 0xd1, 0xfa, 0x18, 0xfe, 0xeb, 0x15, 0x8c, 0x6c,
	0x1a, 0xd9, 0xd6, 0xb9, 0x56, 0x1e, 0xb4, 0xcb, 0x48, 0x09, 0x8a, 0x52, 0xaf, 0xd5, 0x73, 0x06,
	0xbb, 0xc1, 0x57, 0x7a, 0x37, 0x7a, 0x19, 0xce, 0x72, 0x99, 0xd5, 0x31, 0x4f, 0x68, 0xe1, 0x9b,
	0x17, 0xbb, 0xcc, 0xc9, 0x2f, 0xe3, 0x19, 0xf9, 0x5b, 0x17, 0xbc, 0x2d, 0x63, 0x1b, 0xc7, 0xff,
	0xcc, 0x11, 0x47, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x35, 0xa5, 0xd6, 0x00, 0xe4, 0x01, 0x00,
	0x00,
}
