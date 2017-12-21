// Code generated by protoc-gen-go.
// source: status.ext.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Status struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Code   int32  `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	Detail string `protobuf:"bytes,3,opt,name=detail" json:"detail,omitempty"`
	Status string `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Status) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *Status) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Status)(nil), "pb.Status")
}

func init() { proto.RegisterFile("status.ext.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2e, 0x49, 0x2c,
	0x29, 0x2d, 0xd6, 0x4b, 0xad, 0x28, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x52, 0x8a, 0xe1, 0x62, 0x0b, 0x06, 0x8b, 0x0b, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1, 0x24, 0xe7, 0xa7, 0xa4, 0x4a,
	0x30, 0x29, 0x30, 0x6a, 0xb0, 0x06, 0x81, 0xd9, 0x42, 0x62, 0x5c, 0x6c, 0x29, 0xa9, 0x25, 0x89,
	0x99, 0x39, 0x12, 0xcc, 0x60, 0x75, 0x50, 0x1e, 0x48, 0x1c, 0x62, 0xba, 0x04, 0x0b, 0x44, 0x1c,
	0xc2, 0x4b, 0x62, 0x03, 0x5b, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x55, 0x3a, 0xd7,
	0x7c, 0x00, 0x00, 0x00,
}
