// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/chaincode_event.proto

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ChaincodeEvent is used for events and registrations that are specific to chaincode
// string type - "chaincode"
type ChaincodeEvent struct {
	ChaincodeId string `protobuf:"bytes,1,opt,name=chaincode_id,json=chaincodeId" json:"chaincode_id,omitempty"`
	TxId        string `protobuf:"bytes,2,opt,name=tx_id,json=txId" json:"tx_id,omitempty"`
	EventName   string `protobuf:"bytes,3,opt,name=event_name,json=eventName" json:"event_name,omitempty"`
	Payload     []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *ChaincodeEvent) Reset()                    { *m = ChaincodeEvent{} }
func (m *ChaincodeEvent) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeEvent) ProtoMessage()               {}
func (*ChaincodeEvent) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ChaincodeEvent) GetChaincodeId() string {
	if m != nil {
		return m.ChaincodeId
	}
	return ""
}

func (m *ChaincodeEvent) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *ChaincodeEvent) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *ChaincodeEvent) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*ChaincodeEvent)(nil), "protos.ChaincodeEvent")
}

func init() { proto.RegisterFile("peer/chaincode_event.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x59, 0xad, 0x4a, 0xc7, 0xe2, 0x21, 0x22, 0x2c, 0x82, 0x50, 0x7b, 0xb1, 0xa7, 0xe4,
	0xe0, 0x3f, 0xa8, 0x78, 0xe8, 0x45, 0xa4, 0x47, 0x2f, 0x65, 0x36, 0x19, 0x37, 0xc1, 0xee, 0x4e,
	0xc8, 0x46, 0x6d, 0x8f, 0xfe, 0x73, 0x49, 0xc2, 0x2a, 0x3d, 0x85, 0xbc, 0xf7, 0x3e, 0xe6, 0x3d,
	0xb8, 0xf5, 0x44, 0x41, 0x69, 0x8b, 0xae, 0xd7, 0x6c, 0x68, 0x4b, 0x5f, 0xd4, 0x47, 0xe9, 0x03,
	0x47, 0x16, 0xe7, 0xf9, 0x19, 0x16, 0x3f, 0x15, 0x5c, 0x3d, 0x8d, 0x89, 0xe7, 0x14, 0x10, 0xf7,
	0x30, 0xfb, 0x67, 0x9c, 0xa9, 0xab, 0x79, 0xb5, 0x9c, 0x6e, 0x2e, 0xff, 0xb4, 0xb5, 0x11, 0xd7,
	0x70, 0x16, 0xf7, 0xc9, 0x3b, 0xc9, 0xde, 0x24, 0xee, 0xd7, 0x46, 0xdc, 0x01, 0xe4, 0x0b, 0xdb,
	0x1e, 0x3b, 0xaa, 0x4f, 0xb3, 0x33, 0xcd, 0xca, 0x0b, 0x76, 0x24, 0x6a, 0xb8, 0xf0, 0x78, 0xd8,
	0x31, 0x9a, 0x7a, 0x32, 0xaf, 0x96, 0xb3, 0xcd, 0xf8, 0x5d, 0x19, 0x58, 0x70, 0x68, 0xa5, 0x3d,
	0x78, 0x0a, 0x3b, 0x32, 0x2d, 0x05, 0xf9, 0x8e, 0x4d, 0x70, 0xba, 0x74, 0x1d, 0x64, 0xda, 0xb1,
	0xba, 0x39, 0xae, 0xf9, 0x8a, 0xfa, 0x03, 0x5b, 0x7a, 0x7b, 0x68, 0x5d, 0xb4, 0x9f, 0x8d, 0xd4,
	0xdc, 0xa9, 0x6f, 0x8b, 0xd1, 0x0d, 0xcc, 0x5e, 0x15, 0x5e, 0x15, 0x5e, 0x25, 0xbe, 0x29, 0x8b,
	0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x58, 0x23, 0x73, 0x4e, 0x16, 0x01, 0x00, 0x00,
}
