// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FOO int32

const (
	FOO_X FOO = 17
)

var FOO_name = map[int32]string{
	17: "X",
}
var FOO_value = map[string]int32{
	"X": 17,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}
func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}
func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}
func (FOO) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type Test struct {
	Label            *string             `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Type             *int32              `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
	Reps             []int64             `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
	Optionalgroup    *Test_OptionalGroup `protobuf:"group,4,opt,name=OptionalGroup" json:"optionalgroup,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

const Default_Test_Type int32 = 77

func (m *Test) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Test_Type
}

func (m *Test) GetReps() []int64 {
	if m != nil {
		return m.Reps
	}
	return nil
}

func (m *Test) GetOptionalgroup() *Test_OptionalGroup {
	if m != nil {
		return m.Optionalgroup
	}
	return nil
}

type Test_OptionalGroup struct {
	RequiredField    *string `protobuf:"bytes,5,req,name=RequiredField" json:"RequiredField,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Test_OptionalGroup) Reset()                    { *m = Test_OptionalGroup{} }
func (m *Test_OptionalGroup) String() string            { return proto.CompactTextString(m) }
func (*Test_OptionalGroup) ProtoMessage()               {}
func (*Test_OptionalGroup) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *Test_OptionalGroup) GetRequiredField() string {
	if m != nil && m.RequiredField != nil {
		return *m.RequiredField
	}
	return ""
}

func init() {
	proto.RegisterType((*Test)(nil), "protocol.Test")
	proto.RegisterType((*Test_OptionalGroup)(nil), "protocol.Test.OptionalGroup")
	proto.RegisterEnum("protocol.FOO", FOO_name, FOO_value)
}

var fileDescriptor1 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0x93, 0x18,
	0xb9, 0x58, 0x42, 0x80, 0x12, 0x42, 0xbc, 0x5c, 0xac, 0x39, 0x89, 0x49, 0xa9, 0x39, 0x12, 0x8c,
	0x0a, 0x4c, 0x1a, 0x9c, 0x42, 0x02, 0x5c, 0x2c, 0x25, 0x95, 0x05, 0xa9, 0x12, 0x4c, 0x0a, 0x8c,
	0x1a, 0xac, 0x56, 0x4c, 0xe6, 0xe6, 0x42, 0x3c, 0x5c, 0x2c, 0x45, 0xa9, 0x05, 0xc5, 0x12, 0xcc,
	0x0a, 0xcc, 0x1a, 0xcc, 0x42, 0xc6, 0x5c, 0xbc, 0xf9, 0x05, 0x25, 0x99, 0xf9, 0x79, 0x89, 0x39,
	0xe9, 0x45, 0xf9, 0xa5, 0x05, 0x12, 0x2c, 0x40, 0x85, 0x5c, 0x46, 0x32, 0x7a, 0x30, 0x93, 0xf5,
	0x40, 0xa6, 0xea, 0xf9, 0x43, 0xd5, 0xb8, 0x83, 0xd4, 0x48, 0xa9, 0x71, 0xf1, 0xa2, 0x08, 0x08,
	0x89, 0x72, 0xf1, 0x06, 0xa5, 0x16, 0x96, 0x66, 0x16, 0xa5, 0xa6, 0xb8, 0x65, 0xa6, 0xe6, 0xa4,
	0x48, 0xb0, 0x82, 0x2c, 0xd7, 0xe2, 0xe1, 0x62, 0x76, 0xf3, 0xf7, 0x17, 0x62, 0xe5, 0x62, 0x8c,
	0x10, 0x10, 0x04, 0x04, 0x00, 0x00, 0xff, 0xff, 0x88, 0x22, 0xf9, 0x41, 0xb9, 0x00, 0x00, 0x00,
}