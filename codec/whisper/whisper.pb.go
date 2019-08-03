// Code generated by protoc-gen-go. DO NOT EDIT.
// source: whisper.proto

package whisper

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

// Request definition
//
// In Request, tracing `traceContext` is stored in map `meta`.
type Request struct {
	Seqno                *uint64           `protobuf:"varint,1,opt,name=seqno" json:"seqno,omitempty"`
	Appid                *string           `protobuf:"bytes,2,opt,name=appid" json:"appid,omitempty"`
	Userid               *string           `protobuf:"bytes,3,opt,name=userid" json:"userid,omitempty"`
	Userkey              *string           `protobuf:"bytes,4,opt,name=userkey" json:"userkey,omitempty"`
	Version              *uint32           `protobuf:"varint,5,opt,name=version" json:"version,omitempty"`
	Body                 []byte            `protobuf:"bytes,6,opt,name=body" json:"body,omitempty"`
	Meta                 map[string]string `protobuf:"bytes,7,rep,name=meta" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2186b3ddb7c8c37, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSeqno() uint64 {
	if m != nil && m.Seqno != nil {
		return *m.Seqno
	}
	return 0
}

func (m *Request) GetAppid() string {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return ""
}

func (m *Request) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

func (m *Request) GetUserkey() string {
	if m != nil && m.Userkey != nil {
		return *m.Userkey
	}
	return ""
}

func (m *Request) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Request) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

// Response definition
//
// `err_code` and `err_msg` should indicate errors in framework,
// rather than business logic error or error description.
type Response struct {
	Seqno                *uint64  `protobuf:"varint,1,opt,name=seqno" json:"seqno,omitempty"`
	ErrCode              *uint32  `protobuf:"varint,2,opt,name=err_code,json=errCode" json:"err_code,omitempty"`
	ErrMsg               *string  `protobuf:"bytes,3,opt,name=err_msg,json=errMsg" json:"err_msg,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2186b3ddb7c8c37, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSeqno() uint64 {
	if m != nil && m.Seqno != nil {
		return *m.Seqno
	}
	return 0
}

func (m *Response) GetErrCode() uint32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "whisper.Request")
	proto.RegisterMapType((map[string]string)(nil), "whisper.Request.MetaEntry")
	proto.RegisterType((*Response)(nil), "whisper.Response")
}

func init() { proto.RegisterFile("whisper.proto", fileDescriptor_e2186b3ddb7c8c37) }

var fileDescriptor_e2186b3ddb7c8c37 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x14, 0x84, 0xe5, 0x26, 0x6d, 0xda, 0x07, 0x91, 0x90, 0x85, 0xc0, 0x74, 0x8a, 0x3a, 0x79, 0xca,
	0xc0, 0x02, 0x62, 0x45, 0x8c, 0x5d, 0xfc, 0x07, 0xaa, 0x80, 0x9f, 0xda, 0x08, 0x1a, 0xbb, 0xcf,
	0x4e, 0x51, 0xfe, 0x3c, 0x42, 0x76, 0xe2, 0x32, 0xb1, 0xdd, 0x77, 0xcf, 0x96, 0xee, 0x0e, 0xca,
	0xef, 0x43, 0xeb, 0x2c, 0x52, 0x6d, 0xc9, 0x78, 0xc3, 0x8b, 0x09, 0x37, 0x3f, 0x0c, 0x0a, 0x85,
	0xa7, 0x1e, 0x9d, 0xe7, 0xb7, 0x30, 0x77, 0x78, 0xea, 0x8c, 0x60, 0x15, 0x93, 0xb9, 0x1a, 0x21,
	0xb8, 0x8d, 0xb5, 0xad, 0x16, 0xb3, 0x8a, 0xc9, 0x95, 0x1a, 0x81, 0xdf, 0xc1, 0xa2, 0x77, 0x48,
	0xad, 0x16, 0x59, 0xb4, 0x27, 0xe2, 0x02, 0x8a, 0xa0, 0x3e, 0x71, 0x10, 0x79, 0x3c, 0x24, 0x0c,
	0x97, 0x33, 0x92, 0x6b, 0x4d, 0x27, 0xe6, 0x15, 0x93, 0xa5, 0x4a, 0xc8, 0x39, 0xe4, 0xef, 0x46,
	0x0f, 0x62, 0x51, 0x31, 0x79, 0xad, 0xa2, 0xe6, 0x35, 0xe4, 0x47, 0xf4, 0x8d, 0x28, 0xaa, 0x4c,
	0x5e, 0x3d, 0xae, 0xeb, 0x14, 0x7f, 0xca, 0x5a, 0x6f, 0xd1, 0x37, 0x6f, 0x9d, 0xa7, 0x41, 0xc5,
	0x77, 0xeb, 0x27, 0x58, 0x5d, 0x2c, 0x7e, 0x03, 0x59, 0x08, 0xc0, 0x62, 0x80, 0x20, 0x43, 0x89,
	0x73, 0xf3, 0xd5, 0x63, 0x2a, 0x11, 0xe1, 0x65, 0xf6, 0xcc, 0x36, 0x07, 0x58, 0x2a, 0x74, 0xd6,
	0x74, 0x0e, 0xff, 0x19, 0xe0, 0x01, 0x96, 0x48, 0xb4, 0xfb, 0x30, 0x7a, 0xfc, 0x5e, 0xaa, 0x02,
	0x89, 0x5e, 0x8d, 0x46, 0x7e, 0x0f, 0x41, 0xee, 0x8e, 0x6e, 0x9f, 0x66, 0x40, 0xa2, 0xad, 0xdb,
	0x5f, 0x2a, 0xe5, 0x7f, 0x95, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x37, 0x5c, 0x5b, 0x83,
	0x01, 0x00, 0x00,
}