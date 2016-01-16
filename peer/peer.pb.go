// Code generated by protoc-gen-go.
// source: github.com/jtremback/upc/peer/peer.proto
// DO NOT EDIT!

/*
Package peer is a generated protocol buffer package.

It is generated from these files:
	github.com/jtremback/upc/peer/peer.proto

It has these top-level messages:
	Channel
	Identity
	Peer
*/
package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wire "github.com/jtremback/upc/wire"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Channel_State int32

const (
	Channel_PendingOpen   Channel_State = 0
	Channel_Open          Channel_State = 1
	Channel_PendingClosed Channel_State = 2
	Channel_Closed        Channel_State = 3
)

var Channel_State_name = map[int32]string{
	0: "PendingOpen",
	1: "Open",
	2: "PendingClosed",
	3: "Closed",
}
var Channel_State_value = map[string]int32{
	"PendingOpen":   0,
	"Open":          1,
	"PendingClosed": 2,
	"Closed":        3,
}

func (x Channel_State) String() string {
	return proto.EnumName(Channel_State_name, int32(x))
}
func (Channel_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Channel struct {
	ChannelID                []byte          `protobuf:"bytes,1,opt,name=channelID,proto3" json:"channelID,omitempty"`
	OpeningTx                *wire.OpeningTx `protobuf:"bytes,2,opt,name=openingTx" json:"openingTx,omitempty"`
	OpeningTxEnvelope        *wire.Envelope  `protobuf:"bytes,3,opt,name=openingTxEnvelope" json:"openingTxEnvelope,omitempty"`
	LastUpdateTx             *wire.UpdateTx  `protobuf:"bytes,4,opt,name=lastUpdateTx" json:"lastUpdateTx,omitempty"`
	LastUpdateTxEnvelope     *wire.Envelope  `protobuf:"bytes,5,opt,name=lastUpdateTxEnvelope" json:"lastUpdateTxEnvelope,omitempty"`
	LastFullUpdateTx         *wire.UpdateTx  `protobuf:"bytes,6,opt,name=lastFullUpdateTx" json:"lastFullUpdateTx,omitempty"`
	LastFullUpdateTxEnvelope *wire.Envelope  `protobuf:"bytes,7,opt,name=lastFullUpdateTxEnvelope" json:"lastFullUpdateTxEnvelope,omitempty"`
	Me                       uint32          `protobuf:"varint,8,opt,name=me" json:"me,omitempty"`
	State                    Channel_State   `protobuf:"varint,9,opt,name=state,enum=peer.Channel_State" json:"state,omitempty"`
	Identity                 *Identity       `protobuf:"bytes,10,opt,name=identity" json:"identity,omitempty"`
	Peer                     *Peer           `protobuf:"bytes,11,opt,name=peer" json:"peer,omitempty"`
}

func (m *Channel) Reset()                    { *m = Channel{} }
func (m *Channel) String() string            { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()               {}
func (*Channel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Channel) GetOpeningTx() *wire.OpeningTx {
	if m != nil {
		return m.OpeningTx
	}
	return nil
}

func (m *Channel) GetOpeningTxEnvelope() *wire.Envelope {
	if m != nil {
		return m.OpeningTxEnvelope
	}
	return nil
}

func (m *Channel) GetLastUpdateTx() *wire.UpdateTx {
	if m != nil {
		return m.LastUpdateTx
	}
	return nil
}

func (m *Channel) GetLastUpdateTxEnvelope() *wire.Envelope {
	if m != nil {
		return m.LastUpdateTxEnvelope
	}
	return nil
}

func (m *Channel) GetLastFullUpdateTx() *wire.UpdateTx {
	if m != nil {
		return m.LastFullUpdateTx
	}
	return nil
}

func (m *Channel) GetLastFullUpdateTxEnvelope() *wire.Envelope {
	if m != nil {
		return m.LastFullUpdateTxEnvelope
	}
	return nil
}

func (m *Channel) GetIdentity() *Identity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Channel) GetPeer() *Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

type Identity struct {
	Nickname string   `protobuf:"bytes,1,opt,name=nickname" json:"nickname,omitempty"`
	Pubkey   []byte   `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Privkey  []byte   `protobuf:"bytes,3,opt,name=privkey,proto3" json:"privkey,omitempty"`
	Channels [][]byte `protobuf:"bytes,4,rep,name=channels,proto3" json:"channels,omitempty"`
}

func (m *Identity) Reset()                    { *m = Identity{} }
func (m *Identity) String() string            { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()               {}
func (*Identity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Peer struct {
	Nickname string   `protobuf:"bytes,1,opt,name=nickname" json:"nickname,omitempty"`
	Pubkey   []byte   `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Channels [][]byte `protobuf:"bytes,3,rep,name=channels,proto3" json:"channels,omitempty"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Channel)(nil), "peer.Channel")
	proto.RegisterType((*Identity)(nil), "peer.Identity")
	proto.RegisterType((*Peer)(nil), "peer.Peer")
	proto.RegisterEnum("peer.Channel_State", Channel_State_name, Channel_State_value)
}

var fileDescriptor0 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0x4b, 0x8f, 0xd3, 0x30,
	0x10, 0xa6, 0x4d, 0x9a, 0x26, 0xd3, 0x6e, 0xb7, 0x3b, 0x20, 0x64, 0xad, 0x10, 0x42, 0x39, 0x2d,
	0x48, 0xa4, 0x52, 0xb9, 0x21, 0x2e, 0xb0, 0x80, 0x54, 0x2e, 0x54, 0xe6, 0x71, 0xcf, 0xc3, 0x6a,
	0x43, 0x53, 0x27, 0x4a, 0x9c, 0x42, 0xff, 0x2c, 0xbf, 0x05, 0xdb, 0x79, 0xb4, 0xa5, 0xe9, 0x65,
	0x2f, 0xd6, 0x7c, 0x0f, 0xcf, 0x17, 0x6b, 0x26, 0x70, 0xb7, 0x8a, 0xc5, 0xba, 0x0c, 0xbc, 0x30,
	0xdd, 0xce, 0x7e, 0x89, 0x9c, 0x6d, 0x03, 0x3f, 0xdc, 0xcc, 0xca, 0x2c, 0x9c, 0x65, 0x8c, 0xe5,
	0xfa, 0xf0, 0xb2, 0x3c, 0x15, 0x29, 0x9a, 0xaa, 0xbe, 0xbd, 0xec, 0xff, 0x1d, 0xe7, 0x4c, 0x1f,
	0x95, 0xdf, 0xfd, 0x6b, 0xc2, 0xf0, 0x7e, 0xed, 0x73, 0xce, 0x12, 0x7c, 0x06, 0x4e, 0x58, 0x95,
	0x8b, 0x8f, 0xa4, 0xf7, 0xa2, 0x77, 0x37, 0xa6, 0x07, 0x02, 0x5f, 0x83, 0x93, 0x66, 0x8c, 0xc7,
	0x7c, 0xf5, 0xfd, 0x0f, 0xe9, 0x4b, 0x75, 0x34, 0xbf, 0xf6, 0x74, 0xa7, 0xaf, 0x0d, 0x4d, 0x0f,
	0x0e, 0x7c, 0x07, 0x37, 0x2d, 0xf8, 0xc4, 0x77, 0x2c, 0x91, 0x88, 0x18, 0xfa, 0xda, 0xa4, 0xba,
	0xd6, 0xb0, 0xf4, 0xdc, 0x88, 0x73, 0x18, 0x27, 0x7e, 0x21, 0x7e, 0x64, 0x91, 0x2f, 0x98, 0xcc,
	0x33, 0x8f, 0x2f, 0x36, 0x2c, 0x3d, 0xf1, 0xe0, 0x07, 0x78, 0x72, 0x8c, 0xdb, 0xd0, 0x41, 0x67,
	0x68, 0xa7, 0x17, 0xdf, 0xc2, 0x54, 0xf1, 0x9f, 0xcb, 0x24, 0x69, 0xb3, 0xad, 0xce, 0xec, 0x33,
	0x1f, 0x7e, 0x01, 0xf2, 0x3f, 0xd7, 0x7e, 0xc3, 0xb0, 0xf3, 0x1b, 0x2e, 0xfa, 0x71, 0x02, 0xfd,
	0x2d, 0x23, 0xb6, 0xbc, 0x75, 0x45, 0x65, 0x85, 0x2f, 0x61, 0x50, 0x08, 0x69, 0x21, 0x8e, 0xa4,
	0x26, 0xf3, 0xc7, 0x9e, 0x1e, 0x79, 0x3d, 0x38, 0xef, 0x9b, 0x92, 0x68, 0xe5, 0xc0, 0x57, 0x60,
	0xc7, 0x11, 0xe3, 0x22, 0x16, 0x7b, 0x02, 0x75, 0xac, 0x76, 0x2f, 0x6a, 0x96, 0xb6, 0x3a, 0x3e,
	0x07, 0xbd, 0x2f, 0x64, 0xa4, 0x7d, 0x50, 0xf9, 0x96, 0xf2, 0xa0, 0x9a, 0x77, 0xdf, 0xc3, 0x40,
	0xf7, 0xc6, 0x6b, 0x18, 0x2d, 0x19, 0x8f, 0xe4, 0x90, 0xd4, 0xb0, 0xa7, 0x8f, 0xd0, 0x06, 0x53,
	0x57, 0x3d, 0xbc, 0x81, 0xab, 0x5a, 0xba, 0x4f, 0xd2, 0x82, 0x45, 0xd3, 0x3e, 0x02, 0x58, 0x75,
	0x6d, 0xb8, 0x02, 0xec, 0x26, 0x18, 0x6f, 0xc1, 0xe6, 0x71, 0xb8, 0xe1, 0xbe, 0x7c, 0x9b, 0xda,
	0x2f, 0x87, 0xb6, 0x18, 0x9f, 0x82, 0x95, 0x95, 0xc1, 0x86, 0xed, 0xf5, 0x6e, 0x8d, 0x69, 0x8d,
	0x90, 0xc0, 0x30, 0xcb, 0xe3, 0x9d, 0x12, 0x0c, 0x2d, 0x34, 0x50, 0x75, 0xab, 0xb7, 0xb3, 0x90,
	0xfb, 0x61, 0x48, 0xa9, 0xc5, 0xee, 0x4f, 0x30, 0xd5, 0x33, 0x1e, 0x94, 0x78, 0xdc, 0xd7, 0x38,
	0xed, 0x1b, 0x58, 0xfa, 0xaf, 0x79, 0xf3, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x62, 0xa3, 0x2b, 0x9a,
	0x91, 0x03, 0x00, 0x00,
}