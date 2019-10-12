// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protob/ecdsa-signing.proto

package signing

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

//
// Represents a P2P message sent to each party during Round 1 of the ECDSA TSS signing protocol.
type SignRound1Message1 struct {
	C                    []byte   `protobuf:"bytes,1,opt,name=c,proto3" json:"c,omitempty"`
	RangeProofAlice      [][]byte `protobuf:"bytes,2,rep,name=range_proof_alice,json=rangeProofAlice,proto3" json:"range_proof_alice,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRound1Message1) Reset()         { *m = SignRound1Message1{} }
func (m *SignRound1Message1) String() string { return proto.CompactTextString(m) }
func (*SignRound1Message1) ProtoMessage()    {}
func (*SignRound1Message1) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f861bfc687bec19, []int{0}
}

func (m *SignRound1Message1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRound1Message1.Unmarshal(m, b)
}
func (m *SignRound1Message1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRound1Message1.Marshal(b, m, deterministic)
}
func (m *SignRound1Message1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRound1Message1.Merge(m, src)
}
func (m *SignRound1Message1) XXX_Size() int {
	return xxx_messageInfo_SignRound1Message1.Size(m)
}
func (m *SignRound1Message1) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRound1Message1.DiscardUnknown(m)
}

var xxx_messageInfo_SignRound1Message1 proto.InternalMessageInfo

func (m *SignRound1Message1) GetC() []byte {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *SignRound1Message1) GetRangeProofAlice() [][]byte {
	if m != nil {
		return m.RangeProofAlice
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 1 of the ECDSA TSS signing protocol.
type SignRound1Message2 struct {
	Commitment           []byte   `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRound1Message2) Reset()         { *m = SignRound1Message2{} }
func (m *SignRound1Message2) String() string { return proto.CompactTextString(m) }
func (*SignRound1Message2) ProtoMessage()    {}
func (*SignRound1Message2) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f861bfc687bec19, []int{1}
}

func (m *SignRound1Message2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRound1Message2.Unmarshal(m, b)
}
func (m *SignRound1Message2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRound1Message2.Marshal(b, m, deterministic)
}
func (m *SignRound1Message2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRound1Message2.Merge(m, src)
}
func (m *SignRound1Message2) XXX_Size() int {
	return xxx_messageInfo_SignRound1Message2.Size(m)
}
func (m *SignRound1Message2) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRound1Message2.DiscardUnknown(m)
}

var xxx_messageInfo_SignRound1Message2 proto.InternalMessageInfo

func (m *SignRound1Message2) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

//
// Represents a P2P message sent to each party during Round 2 of the ECDSA TSS signing protocol.
type SignRound2Message struct {
	C1                   []byte   `protobuf:"bytes,1,opt,name=c1,proto3" json:"c1,omitempty"`
	C2                   []byte   `protobuf:"bytes,2,opt,name=c2,proto3" json:"c2,omitempty"`
	ProofBob             [][]byte `protobuf:"bytes,3,rep,name=proof_bob,json=proofBob,proto3" json:"proof_bob,omitempty"`
	ProofBobWc           [][]byte `protobuf:"bytes,4,rep,name=proof_bob_wc,json=proofBobWc,proto3" json:"proof_bob_wc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRound2Message) Reset()         { *m = SignRound2Message{} }
func (m *SignRound2Message) String() string { return proto.CompactTextString(m) }
func (*SignRound2Message) ProtoMessage()    {}
func (*SignRound2Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f861bfc687bec19, []int{2}
}

func (m *SignRound2Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRound2Message.Unmarshal(m, b)
}
func (m *SignRound2Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRound2Message.Marshal(b, m, deterministic)
}
func (m *SignRound2Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRound2Message.Merge(m, src)
}
func (m *SignRound2Message) XXX_Size() int {
	return xxx_messageInfo_SignRound2Message.Size(m)
}
func (m *SignRound2Message) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRound2Message.DiscardUnknown(m)
}

var xxx_messageInfo_SignRound2Message proto.InternalMessageInfo

func (m *SignRound2Message) GetC1() []byte {
	if m != nil {
		return m.C1
	}
	return nil
}

func (m *SignRound2Message) GetC2() []byte {
	if m != nil {
		return m.C2
	}
	return nil
}

func (m *SignRound2Message) GetProofBob() [][]byte {
	if m != nil {
		return m.ProofBob
	}
	return nil
}

func (m *SignRound2Message) GetProofBobWc() [][]byte {
	if m != nil {
		return m.ProofBobWc
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 3 of the ECDSA TSS signing protocol.
type SignRound3Message struct {
	Theta                []byte   `protobuf:"bytes,1,opt,name=theta,proto3" json:"theta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRound3Message) Reset()         { *m = SignRound3Message{} }
func (m *SignRound3Message) String() string { return proto.CompactTextString(m) }
func (*SignRound3Message) ProtoMessage()    {}
func (*SignRound3Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f861bfc687bec19, []int{3}
}

func (m *SignRound3Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRound3Message.Unmarshal(m, b)
}
func (m *SignRound3Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRound3Message.Marshal(b, m, deterministic)
}
func (m *SignRound3Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRound3Message.Merge(m, src)
}
func (m *SignRound3Message) XXX_Size() int {
	return xxx_messageInfo_SignRound3Message.Size(m)
}
func (m *SignRound3Message) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRound3Message.DiscardUnknown(m)
}

var xxx_messageInfo_SignRound3Message proto.InternalMessageInfo

func (m *SignRound3Message) GetTheta() []byte {
	if m != nil {
		return m.Theta
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 4 of the ECDSA TSS signing protocol.
type SignRound4Message struct {
	DeCommitment         [][]byte `protobuf:"bytes,1,rep,name=de_commitment,json=deCommitment,proto3" json:"de_commitment,omitempty"`
	ProofAlphaX          []byte   `protobuf:"bytes,2,opt,name=proof_alpha_x,json=proofAlphaX,proto3" json:"proof_alpha_x,omitempty"`
	ProofAlphaY          []byte   `protobuf:"bytes,3,opt,name=proof_alpha_y,json=proofAlphaY,proto3" json:"proof_alpha_y,omitempty"`
	ProofT               []byte   `protobuf:"bytes,4,opt,name=proof_t,json=proofT,proto3" json:"proof_t,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRound4Message) Reset()         { *m = SignRound4Message{} }
func (m *SignRound4Message) String() string { return proto.CompactTextString(m) }
func (*SignRound4Message) ProtoMessage()    {}
func (*SignRound4Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f861bfc687bec19, []int{4}
}

func (m *SignRound4Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRound4Message.Unmarshal(m, b)
}
func (m *SignRound4Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRound4Message.Marshal(b, m, deterministic)
}
func (m *SignRound4Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRound4Message.Merge(m, src)
}
func (m *SignRound4Message) XXX_Size() int {
	return xxx_messageInfo_SignRound4Message.Size(m)
}
func (m *SignRound4Message) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRound4Message.DiscardUnknown(m)
}

var xxx_messageInfo_SignRound4Message proto.InternalMessageInfo

func (m *SignRound4Message) GetDeCommitment() [][]byte {
	if m != nil {
		return m.DeCommitment
	}
	return nil
}

func (m *SignRound4Message) GetProofAlphaX() []byte {
	if m != nil {
		return m.ProofAlphaX
	}
	return nil
}

func (m *SignRound4Message) GetProofAlphaY() []byte {
	if m != nil {
		return m.ProofAlphaY
	}
	return nil
}

func (m *SignRound4Message) GetProofT() []byte {
	if m != nil {
		return m.ProofT
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 5 of the ECDSA TSS signing protocol.
type SignRound5Message struct {
	Commitment           []byte   `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRound5Message) Reset()         { *m = SignRound5Message{} }
func (m *SignRound5Message) String() string { return proto.CompactTextString(m) }
func (*SignRound5Message) ProtoMessage()    {}
func (*SignRound5Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f861bfc687bec19, []int{5}
}

func (m *SignRound5Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRound5Message.Unmarshal(m, b)
}
func (m *SignRound5Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRound5Message.Marshal(b, m, deterministic)
}
func (m *SignRound5Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRound5Message.Merge(m, src)
}
func (m *SignRound5Message) XXX_Size() int {
	return xxx_messageInfo_SignRound5Message.Size(m)
}
func (m *SignRound5Message) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRound5Message.DiscardUnknown(m)
}

var xxx_messageInfo_SignRound5Message proto.InternalMessageInfo

func (m *SignRound5Message) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 6 of the ECDSA TSS signing protocol.
type SignRound6Message struct {
	DeCommitment         [][]byte `protobuf:"bytes,1,rep,name=de_commitment,json=deCommitment,proto3" json:"de_commitment,omitempty"`
	ProofAlphaX          []byte   `protobuf:"bytes,2,opt,name=proof_alpha_x,json=proofAlphaX,proto3" json:"proof_alpha_x,omitempty"`
	ProofAlphaY          []byte   `protobuf:"bytes,3,opt,name=proof_alpha_y,json=proofAlphaY,proto3" json:"proof_alpha_y,omitempty"`
	ProofT               []byte   `protobuf:"bytes,4,opt,name=proof_t,json=proofT,proto3" json:"proof_t,omitempty"`
	VProofAlphaX         []byte   `protobuf:"bytes,5,opt,name=v_proof_alpha_x,json=vProofAlphaX,proto3" json:"v_proof_alpha_x,omitempty"`
	VProofAlphaY         []byte   `protobuf:"bytes,6,opt,name=v_proof_alpha_y,json=vProofAlphaY,proto3" json:"v_proof_alpha_y,omitempty"`
	VProofT              []byte   `protobuf:"bytes,7,opt,name=v_proof_t,json=vProofT,proto3" json:"v_proof_t,omitempty"`
	VProofU              []byte   `protobuf:"bytes,8,opt,name=v_proof_u,json=vProofU,proto3" json:"v_proof_u,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRound6Message) Reset()         { *m = SignRound6Message{} }
func (m *SignRound6Message) String() string { return proto.CompactTextString(m) }
func (*SignRound6Message) ProtoMessage()    {}
func (*SignRound6Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f861bfc687bec19, []int{6}
}

func (m *SignRound6Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRound6Message.Unmarshal(m, b)
}
func (m *SignRound6Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRound6Message.Marshal(b, m, deterministic)
}
func (m *SignRound6Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRound6Message.Merge(m, src)
}
func (m *SignRound6Message) XXX_Size() int {
	return xxx_messageInfo_SignRound6Message.Size(m)
}
func (m *SignRound6Message) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRound6Message.DiscardUnknown(m)
}

var xxx_messageInfo_SignRound6Message proto.InternalMessageInfo

func (m *SignRound6Message) GetDeCommitment() [][]byte {
	if m != nil {
		return m.DeCommitment
	}
	return nil
}

func (m *SignRound6Message) GetProofAlphaX() []byte {
	if m != nil {
		return m.ProofAlphaX
	}
	return nil
}

func (m *SignRound6Message) GetProofAlphaY() []byte {
	if m != nil {
		return m.ProofAlphaY
	}
	return nil
}

func (m *SignRound6Message) GetProofT() []byte {
	if m != nil {
		return m.ProofT
	}
	return nil
}

func (m *SignRound6Message) GetVProofAlphaX() []byte {
	if m != nil {
		return m.VProofAlphaX
	}
	return nil
}

func (m *SignRound6Message) GetVProofAlphaY() []byte {
	if m != nil {
		return m.VProofAlphaY
	}
	return nil
}

func (m *SignRound6Message) GetVProofT() []byte {
	if m != nil {
		return m.VProofT
	}
	return nil
}

func (m *SignRound6Message) GetVProofU() []byte {
	if m != nil {
		return m.VProofU
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 7 of the ECDSA TSS signing protocol.
type SignRound7Message struct {
	Commitment           []byte   `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRound7Message) Reset()         { *m = SignRound7Message{} }
func (m *SignRound7Message) String() string { return proto.CompactTextString(m) }
func (*SignRound7Message) ProtoMessage()    {}
func (*SignRound7Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f861bfc687bec19, []int{7}
}

func (m *SignRound7Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRound7Message.Unmarshal(m, b)
}
func (m *SignRound7Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRound7Message.Marshal(b, m, deterministic)
}
func (m *SignRound7Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRound7Message.Merge(m, src)
}
func (m *SignRound7Message) XXX_Size() int {
	return xxx_messageInfo_SignRound7Message.Size(m)
}
func (m *SignRound7Message) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRound7Message.DiscardUnknown(m)
}

var xxx_messageInfo_SignRound7Message proto.InternalMessageInfo

func (m *SignRound7Message) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 8 of the ECDSA TSS signing protocol.
type SignRound8Message struct {
	DeCommitment         [][]byte `protobuf:"bytes,1,rep,name=de_commitment,json=deCommitment,proto3" json:"de_commitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRound8Message) Reset()         { *m = SignRound8Message{} }
func (m *SignRound8Message) String() string { return proto.CompactTextString(m) }
func (*SignRound8Message) ProtoMessage()    {}
func (*SignRound8Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f861bfc687bec19, []int{8}
}

func (m *SignRound8Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRound8Message.Unmarshal(m, b)
}
func (m *SignRound8Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRound8Message.Marshal(b, m, deterministic)
}
func (m *SignRound8Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRound8Message.Merge(m, src)
}
func (m *SignRound8Message) XXX_Size() int {
	return xxx_messageInfo_SignRound8Message.Size(m)
}
func (m *SignRound8Message) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRound8Message.DiscardUnknown(m)
}

var xxx_messageInfo_SignRound8Message proto.InternalMessageInfo

func (m *SignRound8Message) GetDeCommitment() [][]byte {
	if m != nil {
		return m.DeCommitment
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 9 of the ECDSA TSS signing protocol.
type SignRound9Message struct {
	S                    []byte   `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRound9Message) Reset()         { *m = SignRound9Message{} }
func (m *SignRound9Message) String() string { return proto.CompactTextString(m) }
func (*SignRound9Message) ProtoMessage()    {}
func (*SignRound9Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f861bfc687bec19, []int{9}
}

func (m *SignRound9Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRound9Message.Unmarshal(m, b)
}
func (m *SignRound9Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRound9Message.Marshal(b, m, deterministic)
}
func (m *SignRound9Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRound9Message.Merge(m, src)
}
func (m *SignRound9Message) XXX_Size() int {
	return xxx_messageInfo_SignRound9Message.Size(m)
}
func (m *SignRound9Message) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRound9Message.DiscardUnknown(m)
}

var xxx_messageInfo_SignRound9Message proto.InternalMessageInfo

func (m *SignRound9Message) GetS() []byte {
	if m != nil {
		return m.S
	}
	return nil
}

func init() {
	proto.RegisterType((*SignRound1Message1)(nil), "SignRound1Message1")
	proto.RegisterType((*SignRound1Message2)(nil), "SignRound1Message2")
	proto.RegisterType((*SignRound2Message)(nil), "SignRound2Message")
	proto.RegisterType((*SignRound3Message)(nil), "SignRound3Message")
	proto.RegisterType((*SignRound4Message)(nil), "SignRound4Message")
	proto.RegisterType((*SignRound5Message)(nil), "SignRound5Message")
	proto.RegisterType((*SignRound6Message)(nil), "SignRound6Message")
	proto.RegisterType((*SignRound7Message)(nil), "SignRound7Message")
	proto.RegisterType((*SignRound8Message)(nil), "SignRound8Message")
	proto.RegisterType((*SignRound9Message)(nil), "SignRound9Message")
}

func init() { proto.RegisterFile("protob/ecdsa-signing.proto", fileDescriptor_5f861bfc687bec19) }

var fileDescriptor_5f861bfc687bec19 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0x5f, 0x4b, 0xe3, 0x40,
	0x14, 0xc5, 0x49, 0xfa, 0xff, 0x6e, 0xba, 0xa5, 0xc3, 0xc2, 0x0e, 0x5d, 0x58, 0xba, 0xb3, 0x2c,
	0x74, 0x17, 0xd6, 0x92, 0xb4, 0x6a, 0x7d, 0xb4, 0x3e, 0x2b, 0xa5, 0x56, 0xb4, 0xbe, 0x84, 0x64,
	0x32, 0xa6, 0x01, 0x9b, 0x09, 0x49, 0x5a, 0xed, 0x47, 0xf1, 0xc9, 0xaf, 0x2a, 0x9d, 0x64, 0xca,
	0xb4, 0x15, 0xd4, 0x37, 0x1f, 0xef, 0x3d, 0xbf, 0x3b, 0xe7, 0xce, 0x81, 0x0b, 0xad, 0x28, 0xe6,
	0x29, 0x77, 0xbb, 0x8c, 0x7a, 0x89, 0xf3, 0x3f, 0x09, 0xfc, 0x30, 0x08, 0xfd, 0x03, 0xd1, 0x24,
	0x17, 0x80, 0x2e, 0x03, 0x3f, 0x1c, 0xf3, 0x45, 0xe8, 0x99, 0xe7, 0x2c, 0x49, 0x1c, 0x9f, 0x99,
	0xc8, 0x00, 0x8d, 0x62, 0xad, 0xad, 0x75, 0x8c, 0xb1, 0x46, 0xd1, 0x3f, 0x68, 0xc6, 0x4e, 0xe8,
	0x33, 0x3b, 0x8a, 0x39, 0xbf, 0xb3, 0x9d, 0xfb, 0x80, 0x32, 0xac, 0xb7, 0x0b, 0x1d, 0x63, 0xdc,
	0x10, 0xc2, 0x68, 0xdd, 0x3f, 0x5d, 0xb7, 0x49, 0xff, 0x95, 0xf7, 0x2c, 0xf4, 0x13, 0x80, 0xf2,
	0xf9, 0x3c, 0x48, 0xe7, 0x2c, 0x4c, 0xf3, 0x87, 0x95, 0x0e, 0x89, 0xa1, 0xb9, 0x99, 0xb2, 0xf2,
	0x29, 0xf4, 0x15, 0x74, 0x6a, 0xe6, 0xb0, 0x4e, 0x4d, 0x51, 0x5b, 0x58, 0xcf, 0x6b, 0x0b, 0xfd,
	0x80, 0x5a, 0xb6, 0x90, 0xcb, 0x5d, 0x5c, 0x10, 0xeb, 0x54, 0x45, 0x63, 0xc8, 0x5d, 0xd4, 0x06,
	0x63, 0x23, 0xda, 0x0f, 0x14, 0x17, 0x85, 0x0e, 0x52, 0xbf, 0xa6, 0xe4, 0xaf, 0xe2, 0xd9, 0x93,
	0x9e, 0xdf, 0xa0, 0x94, 0xce, 0x58, 0xea, 0xe4, 0xb6, 0x59, 0x41, 0x9e, 0x34, 0x85, 0xed, 0x4b,
	0xf6, 0x37, 0xd4, 0x3d, 0x66, 0x6f, 0xfd, 0x6b, 0xed, 0x61, 0x78, 0xec, 0x6c, 0xd3, 0x43, 0x04,
	0xea, 0x32, 0xb5, 0x68, 0xe6, 0xd8, 0x8f, 0xf9, 0xfe, 0x5f, 0xa2, 0x2c, 0xb2, 0x68, 0xe6, 0xdc,
	0xec, 0x32, 0x2b, 0x5c, 0xd8, 0x65, 0xa6, 0xe8, 0x3b, 0x54, 0x32, 0x26, 0xc5, 0x45, 0xa1, 0x96,
	0x45, 0x39, 0x21, 0x3d, 0x65, 0xb5, 0x43, 0xb9, 0xda, 0x5b, 0x79, 0x3f, 0xeb, 0xca, 0xd4, 0xd1,
	0xa7, 0xfa, 0x10, 0xfa, 0x03, 0x8d, 0xa5, 0xbd, 0x6d, 0x51, 0x12, 0x80, 0xb1, 0x1c, 0x29, 0x1e,
	0x7b, 0xd8, 0x0a, 0x97, 0xf7, 0xb0, 0x29, 0x6a, 0x41, 0x4d, 0x62, 0x29, 0xae, 0x08, 0xa0, 0x92,
	0x01, 0x13, 0x55, 0x5b, 0xe0, 0xaa, 0xaa, 0x5d, 0x6d, 0xc5, 0x7a, 0xfc, 0xde, 0x58, 0x07, 0xca,
	0xd0, 0xe0, 0x23, 0xa9, 0x92, 0x5f, 0xca, 0xe4, 0x89, 0x9c, 0x34, 0x40, 0x4b, 0xe4, 0x15, 0x26,
	0xc3, 0xc6, 0x6d, 0x5d, 0x1c, 0x70, 0x37, 0x3f, 0x60, 0xb7, 0x2c, 0x2e, 0xb8, 0xf7, 0x12, 0x00,
	0x00, 0xff, 0xff, 0xe8, 0xdb, 0x9c, 0xf3, 0xdf, 0x03, 0x00, 0x00,
}