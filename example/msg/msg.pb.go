// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg/msg.proto

/*
Package cmsg is a generated protocol buffer package.

It is generated from these files:
	msg/msg.proto

It has these top-level messages:
	ReqHello
	RespHello
	ReqSession2Server
	RespSession2Server
	ReqServer2Server
	RespServer2Server
	ReqRequest
	RespRequest
*/
package cmsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ReqHello struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ReqHello) Reset()                    { *m = ReqHello{} }
func (m *ReqHello) String() string            { return proto.CompactTextString(m) }
func (*ReqHello) ProtoMessage()               {}
func (*ReqHello) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReqHello) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RespHello struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *RespHello) Reset()                    { *m = RespHello{} }
func (m *RespHello) String() string            { return proto.CompactTextString(m) }
func (*RespHello) ProtoMessage()               {}
func (*RespHello) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RespHello) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ReqSession2Server struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ReqSession2Server) Reset()                    { *m = ReqSession2Server{} }
func (m *ReqSession2Server) String() string            { return proto.CompactTextString(m) }
func (*ReqSession2Server) ProtoMessage()               {}
func (*ReqSession2Server) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReqSession2Server) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RespSession2Server struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *RespSession2Server) Reset()                    { *m = RespSession2Server{} }
func (m *RespSession2Server) String() string            { return proto.CompactTextString(m) }
func (*RespSession2Server) ProtoMessage()               {}
func (*RespSession2Server) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RespSession2Server) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ReqServer2Server struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ReqServer2Server) Reset()                    { *m = ReqServer2Server{} }
func (m *ReqServer2Server) String() string            { return proto.CompactTextString(m) }
func (*ReqServer2Server) ProtoMessage()               {}
func (*ReqServer2Server) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReqServer2Server) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RespServer2Server struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *RespServer2Server) Reset()                    { *m = RespServer2Server{} }
func (m *RespServer2Server) String() string            { return proto.CompactTextString(m) }
func (*RespServer2Server) ProtoMessage()               {}
func (*RespServer2Server) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RespServer2Server) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ReqRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ReqRequest) Reset()                    { *m = ReqRequest{} }
func (m *ReqRequest) String() string            { return proto.CompactTextString(m) }
func (*ReqRequest) ProtoMessage()               {}
func (*ReqRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReqRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RespRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *RespRequest) Reset()                    { *m = RespRequest{} }
func (m *RespRequest) String() string            { return proto.CompactTextString(m) }
func (*RespRequest) ProtoMessage()               {}
func (*RespRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RespRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqHello)(nil), "cmsg.ReqHello")
	proto.RegisterType((*RespHello)(nil), "cmsg.RespHello")
	proto.RegisterType((*ReqSession2Server)(nil), "cmsg.ReqSession2Server")
	proto.RegisterType((*RespSession2Server)(nil), "cmsg.RespSession2Server")
	proto.RegisterType((*ReqServer2Server)(nil), "cmsg.ReqServer2Server")
	proto.RegisterType((*RespServer2Server)(nil), "cmsg.RespServer2Server")
	proto.RegisterType((*ReqRequest)(nil), "cmsg.ReqRequest")
	proto.RegisterType((*RespRequest)(nil), "cmsg.RespRequest")
}

func init() { proto.RegisterFile("msg/msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x2d, 0x4e, 0xd7,
	0xcf, 0x2d, 0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0xce, 0x2d, 0x4e, 0x57,
	0x92, 0xe3, 0xe2, 0x08, 0x4a, 0x2d, 0xf4, 0x48, 0xcd, 0xc9, 0xc9, 0x17, 0x12, 0xe2, 0x62, 0xc9,
	0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xe4, 0xb9, 0x38,
	0x83, 0x52, 0x8b, 0x0b, 0x70, 0x2b, 0x50, 0xe7, 0x12, 0x0c, 0x4a, 0x2d, 0x0c, 0x4e, 0x2d, 0x2e,
	0xce, 0xcc, 0xcf, 0x33, 0x0a, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xc2, 0xaa, 0x50, 0x83, 0x4b, 0x08,
	0x64, 0x12, 0x11, 0x2a, 0xd5, 0xb8, 0x04, 0xc0, 0x46, 0x82, 0x14, 0xe0, 0x53, 0x07, 0xb6, 0x1a,
	0x64, 0x22, 0x21, 0x85, 0x0a, 0x5c, 0x5c, 0x41, 0xa9, 0x85, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x58, 0x55, 0x28, 0x72, 0x71, 0x83, 0x8c, 0xc2, 0xa3, 0x24, 0x89, 0x0d, 0x1c, 0x6c, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x1c, 0xf6, 0xf1, 0x47, 0x01, 0x00, 0x00,
}
