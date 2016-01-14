// Code generated by protoc-gen-go.
// source: pb.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb.proto

It has these top-level messages:
	OpeningTx
	UpdateTx
	Fulfillment
	Envelope
	Channel
	Identity
	Peer
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Envelope_Type int32

const (
	Envelope_OpeningTx         Envelope_Type = 0
	Envelope_OpeningTxProposal Envelope_Type = 1
	Envelope_UpdateTx          Envelope_Type = 2
	Envelope_UpdateTxProposal  Envelope_Type = 3
	Envelope_Fulfillment       Envelope_Type = 4
)

var Envelope_Type_name = map[int32]string{
	0: "OpeningTx",
	1: "OpeningTxProposal",
	2: "UpdateTx",
	3: "UpdateTxProposal",
	4: "Fulfillment",
}
var Envelope_Type_value = map[string]int32{
	"OpeningTx":         0,
	"OpeningTxProposal": 1,
	"UpdateTx":          2,
	"UpdateTxProposal":  3,
	"Fulfillment":       4,
}

func (x Envelope_Type) String() string {
	return proto.EnumName(Envelope_Type_name, int32(x))
}
func (Envelope_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type OpeningTx struct {
	ChannelID  string `protobuf:"bytes,1,opt,name=channelID" json:"channelID,omitempty"`
	Pubkey1    []byte `protobuf:"bytes,2,opt,name=pubkey1,proto3" json:"pubkey1,omitempty"`
	Pubkey2    []byte `protobuf:"bytes,3,opt,name=pubkey2,proto3" json:"pubkey2,omitempty"`
	Amount1    uint32 `protobuf:"varint,4,opt,name=amount1" json:"amount1,omitempty"`
	Amount2    uint32 `protobuf:"varint,5,opt,name=amount2" json:"amount2,omitempty"`
	HoldPeriod uint32 `protobuf:"varint,6,opt,name=holdPeriod" json:"holdPeriod,omitempty"`
}

func (m *OpeningTx) Reset()                    { *m = OpeningTx{} }
func (m *OpeningTx) String() string            { return proto.CompactTextString(m) }
func (*OpeningTx) ProtoMessage()               {}
func (*OpeningTx) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type UpdateTx struct {
	ChannelID      string                `protobuf:"bytes,1,opt,name=channelID" json:"channelID,omitempty"`
	NetTransfer    int64                 `protobuf:"zigzag64,2,opt,name=netTransfer" json:"netTransfer,omitempty"`
	SequenceNumber uint32                `protobuf:"varint,3,opt,name=sequenceNumber" json:"sequenceNumber,omitempty"`
	Fast           bool                  `protobuf:"varint,4,opt,name=fast" json:"fast,omitempty"`
	Conditions     []*UpdateTx_Condition `protobuf:"bytes,5,rep,name=conditions" json:"conditions,omitempty"`
}

func (m *UpdateTx) Reset()                    { *m = UpdateTx{} }
func (m *UpdateTx) String() string            { return proto.CompactTextString(m) }
func (*UpdateTx) ProtoMessage()               {}
func (*UpdateTx) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpdateTx) GetConditions() []*UpdateTx_Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

type UpdateTx_Condition struct {
	PresetCondition     uint32 `protobuf:"varint,1,opt,name=presetCondition" json:"presetCondition,omitempty"`
	ConditionalTransfer int64  `protobuf:"varint,2,opt,name=conditionalTransfer" json:"conditionalTransfer,omitempty"`
	Data                string `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *UpdateTx_Condition) Reset()                    { *m = UpdateTx_Condition{} }
func (m *UpdateTx_Condition) String() string            { return proto.CompactTextString(m) }
func (*UpdateTx_Condition) ProtoMessage()               {}
func (*UpdateTx_Condition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Fulfillment struct {
	ChannelID string `protobuf:"bytes,1,opt,name=channelID" json:"channelID,omitempty"`
	Condition uint32 `protobuf:"varint,2,opt,name=condition" json:"condition,omitempty"`
	Data      string `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *Fulfillment) Reset()                    { *m = Fulfillment{} }
func (m *Fulfillment) String() string            { return proto.CompactTextString(m) }
func (*Fulfillment) ProtoMessage()               {}
func (*Fulfillment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Envelope struct {
	Type       Envelope_Type `protobuf:"varint,1,opt,name=type,enum=pb.Envelope_Type" json:"type,omitempty"`
	Payload    []byte        `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature1 []byte        `protobuf:"bytes,3,opt,name=signature1,proto3" json:"signature1,omitempty"`
	Signature2 []byte        `protobuf:"bytes,4,opt,name=signature2,proto3" json:"signature2,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Channel struct {
	ChannelID        string     `protobuf:"bytes,1,opt,name=channelID" json:"channelID,omitempty"`
	OpeningTx        *OpeningTx `protobuf:"bytes,2,opt,name=openingTx" json:"openingTx,omitempty"`
	LastUpdateTx     *UpdateTx  `protobuf:"bytes,3,opt,name=lastUpdateTx" json:"lastUpdateTx,omitempty"`
	LastFullUpdateTx *UpdateTx  `protobuf:"bytes,4,opt,name=lastFullUpdateTx" json:"lastFullUpdateTx,omitempty"`
	Me               uint32     `protobuf:"varint,5,opt,name=me" json:"me,omitempty"`
}

func (m *Channel) Reset()                    { *m = Channel{} }
func (m *Channel) String() string            { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()               {}
func (*Channel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Channel) GetOpeningTx() *OpeningTx {
	if m != nil {
		return m.OpeningTx
	}
	return nil
}

func (m *Channel) GetLastUpdateTx() *UpdateTx {
	if m != nil {
		return m.LastUpdateTx
	}
	return nil
}

func (m *Channel) GetLastFullUpdateTx() *UpdateTx {
	if m != nil {
		return m.LastFullUpdateTx
	}
	return nil
}

type Identity struct {
	Nickname string   `protobuf:"bytes,1,opt,name=nickname" json:"nickname,omitempty"`
	Pubkey   []byte   `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Privkey  []byte   `protobuf:"bytes,3,opt,name=privkey,proto3" json:"privkey,omitempty"`
	Channels []string `protobuf:"bytes,4,rep,name=channels" json:"channels,omitempty"`
}

func (m *Identity) Reset()                    { *m = Identity{} }
func (m *Identity) String() string            { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()               {}
func (*Identity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Peer struct {
	Nickname string   `protobuf:"bytes,1,opt,name=nickname" json:"nickname,omitempty"`
	Pubkey   []byte   `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Channels []string `protobuf:"bytes,3,rep,name=channels" json:"channels,omitempty"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*OpeningTx)(nil), "pb.OpeningTx")
	proto.RegisterType((*UpdateTx)(nil), "pb.UpdateTx")
	proto.RegisterType((*UpdateTx_Condition)(nil), "pb.UpdateTx.Condition")
	proto.RegisterType((*Fulfillment)(nil), "pb.Fulfillment")
	proto.RegisterType((*Envelope)(nil), "pb.Envelope")
	proto.RegisterType((*Channel)(nil), "pb.Channel")
	proto.RegisterType((*Identity)(nil), "pb.Identity")
	proto.RegisterType((*Peer)(nil), "pb.Peer")
	proto.RegisterEnum("pb.Envelope_Type", Envelope_Type_name, Envelope_Type_value)
}

var fileDescriptor0 = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xc7, 0xcd, 0x87, 0x6d, 0x72, 0xf6, 0xa3, 0xe9, 0xa8, 0x25, 0x14, 0x91, 0x12, 0x50, 0x16,
	0x84, 0xc5, 0x8d, 0x20, 0xde, 0x57, 0x85, 0xde, 0xe8, 0x32, 0xac, 0xde, 0x79, 0x31, 0xbb, 0x99,
	0x6d, 0x43, 0x67, 0x67, 0x62, 0x32, 0x29, 0x2e, 0x3e, 0x95, 0xd7, 0xbe, 0x84, 0x4f, 0xe2, 0x33,
	0x38, 0x33, 0x9b, 0xcf, 0xba, 0xb0, 0xe0, 0x5d, 0xfe, 0xe7, 0x77, 0x72, 0xe6, 0x7f, 0xce, 0x9c,
	0x04, 0xbc, 0x6c, 0x39, 0xcd, 0x72, 0x21, 0x05, 0xb2, 0xb3, 0x65, 0xf4, 0xd3, 0x02, 0xff, 0x53,
	0x46, 0x79, 0xca, 0xaf, 0x17, 0xdf, 0xd1, 0x53, 0xf0, 0x57, 0x37, 0x84, 0x73, 0xca, 0xae, 0xde,
	0x85, 0xd6, 0x85, 0x35, 0xf1, 0x71, 0x1b, 0x40, 0x21, 0x1c, 0x67, 0xe5, 0xf2, 0x96, 0x6e, 0x67,
	0xa1, 0xad, 0xd8, 0x10, 0xd7, 0xb2, 0x25, 0x71, 0xe8, 0x74, 0x49, 0xac, 0x09, 0xd9, 0x88, 0x92,
	0xcb, 0x59, 0xe8, 0x2a, 0x32, 0xc2, 0xb5, 0x6c, 0x49, 0x1c, 0x3e, 0xec, 0x92, 0x18, 0x3d, 0x03,
	0xb8, 0x11, 0x2c, 0x99, 0xd3, 0x3c, 0x15, 0x49, 0x78, 0x64, 0x60, 0x27, 0x12, 0xfd, 0xb2, 0xc1,
	0xfb, 0x9c, 0x25, 0x44, 0xd2, 0x83, 0x96, 0x2f, 0x60, 0xc0, 0xa9, 0x5c, 0xe4, 0x84, 0x17, 0x6b,
	0x9a, 0x1b, 0xdb, 0x08, 0x77, 0x43, 0xe8, 0x05, 0x8c, 0x0b, 0xfa, 0xad, 0xa4, 0x7c, 0x45, 0x3f,
	0x96, 0x9b, 0xa5, 0x4a, 0x72, 0xcc, 0x81, 0xf7, 0xa2, 0x08, 0x81, 0xbb, 0x26, 0x85, 0x34, 0x5d,
	0x78, 0xd8, 0x3c, 0xa3, 0x37, 0x00, 0x2b, 0xc1, 0x93, 0x54, 0xa6, 0x82, 0x17, 0xaa, 0x0b, 0x67,
	0x32, 0x88, 0xcf, 0xa6, 0x6a, 0xbe, 0xb5, 0xbb, 0xe9, 0x65, 0x8d, 0x71, 0x27, 0xf3, 0xfc, 0x07,
	0xf8, 0x0d, 0x40, 0x13, 0x38, 0xc9, 0x72, 0x5a, 0x50, 0xd9, 0x84, 0x4c, 0x1b, 0x23, 0x7c, 0x3f,
	0x8c, 0x5e, 0xc1, 0xa3, 0xa6, 0x08, 0x61, 0xbd, 0xa6, 0x1c, 0xbc, 0x0f, 0x69, 0xd3, 0xca, 0x08,
	0x31, 0x2d, 0xf9, 0xd8, 0x3c, 0x47, 0x5f, 0x61, 0xf0, 0xa1, 0x64, 0xeb, 0x94, 0xb1, 0x0d, 0xe5,
	0xf2, 0xc0, 0xfc, 0x34, 0x6d, 0x6c, 0xd9, 0xc6, 0x56, 0x1b, 0xd8, 0x5b, 0xfe, 0x8f, 0x05, 0xde,
	0x7b, 0x7e, 0x47, 0x99, 0xc8, 0x28, 0x7a, 0x0e, 0xae, 0xdc, 0x66, 0xd4, 0xd4, 0x1d, 0xc7, 0xa7,
	0x7a, 0x34, 0x35, 0x9b, 0x2e, 0x14, 0xc0, 0x06, 0x9b, 0xf5, 0x21, 0x5b, 0x26, 0x48, 0xd2, 0x2c,
	0xd6, 0x4e, 0xea, 0x55, 0x28, 0xd2, 0x6b, 0x4e, 0x64, 0x99, 0xd3, 0x59, 0xb5, 0x5b, 0x9d, 0x48,
	0x8f, 0xc7, 0xe6, 0x6e, 0xba, 0x3c, 0x8e, 0x08, 0xb8, 0xfa, 0x1c, 0x34, 0xea, 0x6c, 0x79, 0xf0,
	0x00, 0x3d, 0x81, 0xd3, 0x46, 0xce, 0x73, 0x91, 0x89, 0x82, 0xb0, 0xc0, 0x42, 0xc3, 0x76, 0xaf,
	0x02, 0x1b, 0x3d, 0x86, 0xa0, 0x56, 0x4d, 0x8e, 0x83, 0x4e, 0x7a, 0xe3, 0x0b, 0xdc, 0xe8, 0xb7,
	0x05, 0xc7, 0x97, 0xbb, 0x81, 0x1d, 0x18, 0xe6, 0x4b, 0xf0, 0x45, 0x7d, 0xaa, 0x69, 0x74, 0x10,
	0x8f, 0xf4, 0x48, 0x1a, 0x2b, 0xb8, 0xe5, 0xea, 0xb2, 0x87, 0x4c, 0xed, 0x58, 0xed, 0xc0, 0xf4,
	0x3e, 0x88, 0x87, 0xdd, 0xed, 0xc2, 0xbd, 0x0c, 0xf4, 0x16, 0x02, 0xad, 0x95, 0x3b, 0xd6, 0xbc,
	0xe5, 0xee, 0x79, 0xeb, 0x9f, 0x2c, 0x34, 0x06, 0x7b, 0x43, 0xab, 0xaf, 0x50, 0x3d, 0x45, 0x12,
	0xbc, 0xab, 0x44, 0x75, 0x97, 0xca, 0x2d, 0x3a, 0x07, 0x8f, 0xa7, 0xab, 0x5b, 0x4e, 0x36, 0xb4,
	0xea, 0xa8, 0xd1, 0xe8, 0x0c, 0x8e, 0x76, 0xdf, 0x79, 0x75, 0x6d, 0x95, 0x32, 0xf7, 0x99, 0xa7,
	0x77, 0x1a, 0xd4, 0xbf, 0x83, 0x9d, 0xd4, 0xd5, 0xaa, 0x79, 0x14, 0xca, 0x9b, 0xa3, 0xab, 0xd5,
	0x3a, 0xfa, 0x02, 0xee, 0x9c, 0xaa, 0xa5, 0xfd, 0x9f, 0x13, 0xbb, 0x75, 0x9d, 0x7e, 0xdd, 0xe5,
	0x91, 0xf9, 0xdb, 0xbd, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x11, 0xcd, 0x23, 0x83, 0xf9, 0x04,
	0x00, 0x00,
}