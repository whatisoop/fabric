// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/proposal_response.proto

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A ProposalResponse is returned from an endorser to the proposal submitter.
// The idea is that this message contains the endorser's response to the
// request of a client to perform an action over a chaincode (or more
// generically on the ledger); the response might be success/error (conveyed in
// the Response field) together with a description of the action and a
// signature over it by that endorser.  If a sufficient number of distinct
// endorsers agree on the same action and produce signature to that effect, a
// transaction can be generated and sent for ordering.
type ProposalResponse struct {
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	// Timestamp is the time that the message
	// was created as  defined by the sender
	Timestamp *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	// A response message indicating whether the
	// endorsement of the action was successful
	Response *Response `protobuf:"bytes,4,opt,name=response" json:"response,omitempty"`
	// The payload of response. It is the bytes of ProposalResponsePayload
	Payload []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	// The endorsement of the proposal, basically
	// the endorser's signature over the payload
	Endorsement *Endorsement `protobuf:"bytes,6,opt,name=endorsement" json:"endorsement,omitempty"`
}

func (m *ProposalResponse) Reset()                    { *m = ProposalResponse{} }
func (m *ProposalResponse) String() string            { return proto.CompactTextString(m) }
func (*ProposalResponse) ProtoMessage()               {}
func (*ProposalResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *ProposalResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ProposalResponse) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ProposalResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ProposalResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ProposalResponse) GetEndorsement() *Endorsement {
	if m != nil {
		return m.Endorsement
	}
	return nil
}

// A response with a representation similar to an HTTP response that can
// be used within another message.
type Response struct {
	// A status code that should follow the HTTP status codes.
	Status int32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	// A message associated with the response code.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// A payload that can be used to include metadata with this response.
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// ProposalResponsePayload is the payload of a proposal response.  This message
// is the "bridge" between the client's request and the endorser's action in
// response to that request. Concretely, for chaincodes, it contains a hashed
// representation of the proposal (proposalHash) and a representation of the
// chaincode state changes and events inside the extension field.
type ProposalResponsePayload struct {
	// Hash of the proposal that triggered this response. The hash is used to
	// link a response with its proposal, both for bookeeping purposes on an
	// asynchronous system and for security reasons (accountability,
	// non-repudiation). The hash usually covers the entire Proposal message
	// (byte-by-byte). However this implies that the hash can only be verified
	// if the entire proposal message is available when ProposalResponsePayload is
	// included in a transaction or stored in the ledger. For confidentiality
	// reasons, with chaincodes it might be undesirable to store the proposal
	// payload in the ledger.  If the type is CHAINCODE, this is handled by
	// separating the proposal's header and
	// the payload: the header is always hashed in its entirety whereas the
	// payload can either be hashed fully, or only its hash may be hashed, or
	// nothing from the payload can be hashed. The PayloadVisibility field in the
	// Header's extension controls to which extent the proposal payload is
	// "visible" in the sense that was just explained.
	ProposalHash []byte `protobuf:"bytes,1,opt,name=proposal_hash,json=proposalHash,proto3" json:"proposal_hash,omitempty"`
	// Extension should be unmarshaled to a type-specific message. The type of
	// the extension in any proposal response depends on the type of the proposal
	// that the client selected when the proposal was initially sent out.  In
	// particular, this information is stored in the type field of a Header.  For
	// chaincode, it's a ChaincodeAction message
	Extension []byte `protobuf:"bytes,2,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (m *ProposalResponsePayload) Reset()                    { *m = ProposalResponsePayload{} }
func (m *ProposalResponsePayload) String() string            { return proto.CompactTextString(m) }
func (*ProposalResponsePayload) ProtoMessage()               {}
func (*ProposalResponsePayload) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *ProposalResponsePayload) GetProposalHash() []byte {
	if m != nil {
		return m.ProposalHash
	}
	return nil
}

func (m *ProposalResponsePayload) GetExtension() []byte {
	if m != nil {
		return m.Extension
	}
	return nil
}

// An endorsement is a signature of an endorser over a proposal response.  By
// producing an endorsement message, an endorser implicitly "approves" that
// proposal response and the actions contained therein. When enough
// endorsements have been collected, a transaction can be generated out of a
// set of proposal responses.  Note that this message only contains an identity
// and a signature but no signed payload. This is intentional because
// endorsements are supposed to be collected in a transaction, and they are all
// expected to endorse a single proposal response/action (many endorsements
// over a single proposal response)
type Endorsement struct {
	// Identity of the endorser (e.g. its certificate)
	Endorser []byte `protobuf:"bytes,1,opt,name=endorser,proto3" json:"endorser,omitempty"`
	// Signature of the payload included in ProposalResponse concatenated with
	// the endorser's certificate; ie, sign(ProposalResponse.payload + endorser)
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Endorsement) Reset()                    { *m = Endorsement{} }
func (m *Endorsement) String() string            { return proto.CompactTextString(m) }
func (*Endorsement) ProtoMessage()               {}
func (*Endorsement) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *Endorsement) GetEndorser() []byte {
	if m != nil {
		return m.Endorser
	}
	return nil
}

func (m *Endorsement) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*ProposalResponse)(nil), "protos.ProposalResponse")
	proto.RegisterType((*Response)(nil), "protos.Response")
	proto.RegisterType((*ProposalResponsePayload)(nil), "protos.ProposalResponsePayload")
	proto.RegisterType((*Endorsement)(nil), "protos.Endorsement")
}

func init() { proto.RegisterFile("peer/proposal_response.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0x4b, 0xc3, 0x30,
	0x10, 0xc6, 0xe9, 0x74, 0x73, 0xcb, 0x26, 0x8c, 0x0a, 0x5a, 0xc6, 0xc0, 0x51, 0x1f, 0xdc, 0x83,
	0xa4, 0xa0, 0x08, 0x3e, 0x0f, 0x44, 0x1f, 0x47, 0x10, 0x1f, 0x44, 0x90, 0x74, 0xbb, 0xa5, 0xc5,
	0xb6, 0x09, 0xb9, 0x54, 0xdd, 0x1f, 0xec, 0xff, 0x21, 0x4d, 0x9b, 0x6e, 0x8a, 0x4f, 0xe3, 0xbb,
	0x7d, 0xf9, 0xdd, 0x7d, 0xd7, 0x23, 0x53, 0x05, 0xa0, 0x23, 0xa5, 0xa5, 0x92, 0xc8, 0xb3, 0x37,
	0x0d, 0xa8, 0x64, 0x81, 0x40, 0x95, 0x96, 0x46, 0xfa, 0x3d, 0xfb, 0x83, 0x93, 0x73, 0x21, 0xa5,
	0xc8, 0x20, 0xb2, 0x32, 0x2e, 0x37, 0x91, 0x49, 0x73, 0x40, 0xc3, 0x73, 0x55, 0x1b, 0xc3, 0x6f,
	0x8f, 0x8c, 0x97, 0x0d, 0x84, 0x35, 0x0c, 0x3f, 0x20, 0x47, 0x1f, 0xa0, 0x31, 0x95, 0x45, 0xe0,
	0xcd, 0xbc, 0x79, 0x97, 0x39, 0xe9, 0xdf, 0x91, 0x41, 0x4b, 0x08, 0x3a, 0x33, 0x6f, 0x3e, 0xbc,
	0x9e, 0xd0, 0xba, 0x07, 0x75, 0x3d, 0xe8, 0x93, 0x73, 0xb0, 0x9d, 0xd9, 0xbf, 0x22, 0x7d, 0x37,
	0x63, 0x70, 0x68, 0x1f, 0x8e, 0xeb, 0x17, 0x48, 0x5d, 0x5f, 0xd6, 0x3a, 0xaa, 0x09, 0x14, 0xdf,
	0x66, 0x92, 0xaf, 0x83, 0xee, 0xcc, 0x9b, 0x8f, 0x98, 0x93, 0xfe, 0x2d, 0x19, 0x42, 0xb1, 0x96,
	0x1a, 0x21, 0x87, 0xc2, 0x04, 0x3d, 0x8b, 0x3a, 0x71, 0xa8, 0xfb, 0xdd, 0x5f, 0x6c, 0xdf, 0x17,
	0x3e, 0x93, 0x7e, 0x1b, 0xef, 0x94, 0xf4, 0xd0, 0x70, 0x53, 0x62, 0x93, 0xae, 0x51, 0x55, 0xd3,
	0x1c, 0x10, 0xb9, 0x00, 0x1b, 0x6d, 0xc0, 0x9c, 0xdc, 0x1f, 0xe7, 0xe0, 0xd7, 0x38, 0xe1, 0x2b,
	0x39, 0xfb, 0xbb, 0xbe, 0x65, 0x33, 0xe9, 0x05, 0x39, 0x6e, 0x3f, 0x4f, 0xc2, 0x31, 0xb1, 0xdd,
	0x46, 0x6c, 0xe4, 0x8a, 0x8f, 0x1c, 0x13, 0x7f, 0x4a, 0x06, 0xf0, 0x65, 0xa0, 0xb0, 0xcb, 0xee,
	0x58, 0xc3, 0xae, 0x10, 0x3e, 0x90, 0xe1, 0x5e, 0x22, 0x7f, 0x42, 0xfa, 0x4d, 0x26, 0xdd, 0xc0,
	0x5a, 0x5d, 0x81, 0x30, 0x15, 0x05, 0x37, 0xa5, 0x06, 0x07, 0x6a, 0x0b, 0x8b, 0x0d, 0x09, 0xa5,
	0x16, 0x34, 0xd9, 0x2a, 0xd0, 0x19, 0xac, 0x05, 0x68, 0xba, 0xe1, 0xb1, 0x4e, 0x57, 0x6e, 0x71,
	0xd5, 0x35, 0x2d, 0xfe, 0x89, 0xb2, 0x7a, 0xe7, 0x02, 0x5e, 0x2e, 0x45, 0x6a, 0x92, 0x32, 0xa6,
	0x2b, 0x99, 0x47, 0x9f, 0x09, 0x37, 0x29, 0x4a, 0xa9, 0xa2, 0x9a, 0x50, 0xdf, 0x16, 0x46, 0x15,
	0x21, 0xae, 0xef, 0xee, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x70, 0x20, 0xb8, 0x9e, 0x02,
	0x00, 0x00,
}
