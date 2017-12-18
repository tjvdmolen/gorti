// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HLAHandle struct {
	Handle uint64 `protobuf:"varint,1,opt,name=handle" json:"handle,omitempty"`
}

func (m *HLAHandle) Reset()                    { *m = HLAHandle{} }
func (m *HLAHandle) String() string            { return proto.CompactTextString(m) }
func (*HLAHandle) ProtoMessage()               {}
func (*HLAHandle) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *HLAHandle) GetHandle() uint64 {
	if m != nil {
		return m.Handle
	}
	return 0
}

func init() {
	proto.RegisterType((*HLAHandle)(nil), "api.HLAHandle")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 76 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x52, 0xe6, 0xe2,
	0xf4, 0xf0, 0x71, 0xf4, 0x48, 0xcc, 0x4b, 0xc9, 0x49, 0x15, 0x12, 0xe3, 0x62, 0xcb, 0x00, 0xb3,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0xa0, 0xbc, 0x24, 0x36, 0xb0, 0x06, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xfb, 0x55, 0x85, 0xd9, 0x3f, 0x00, 0x00, 0x00,
}
