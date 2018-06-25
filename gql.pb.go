// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gql.proto

/*
Package gql is a generated protocol buffer package.

It is generated from these files:
	gql.proto

It has these top-level messages:
*/
package gql

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Type int32

const (
	Type_MUTATION     Type = 0
	Type_QUERY        Type = 1
	Type_SUBSCRIPTION Type = 2
)

var Type_name = map[int32]string{
	0: "MUTATION",
	1: "QUERY",
	2: "SUBSCRIPTION",
}
var Type_value = map[string]int32{
	"MUTATION":     0,
	"QUERY":        1,
	"SUBSCRIPTION": 2,
}

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}
func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (x *Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Type_value, data, "Type")
	if err != nil {
		return err
	}
	*x = Type(value)
	return nil
}
func (Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorGql, []int{0} }

var E_Type = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*Type)(nil),
	Field:         65080,
	Name:          "gql.type",
	Tag:           "varint,65080,opt,name=type,enum=gql.Type",
	Filename:      "gql.proto",
}

func init() {
	proto.RegisterEnum("gql.Type", Type_name, Type_value)
	proto.RegisterExtension(E_Type)
}

func init() { proto.RegisterFile("gql.proto", fileDescriptorGql) }

var fileDescriptorGql = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2f, 0xcc, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2f, 0xcc, 0x91, 0x52, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x25, 0x95, 0xa6, 0xe9, 0xa7, 0xa4, 0x16, 0x27, 0x17, 0x65, 0x16,
	0x94, 0xe4, 0x17, 0x41, 0x94, 0x69, 0x19, 0x72, 0xb1, 0x84, 0x54, 0x16, 0xa4, 0x0a, 0xf1, 0x70,
	0x71, 0xf8, 0x86, 0x86, 0x38, 0x86, 0x78, 0xfa, 0xfb, 0x09, 0x30, 0x08, 0x71, 0x72, 0xb1, 0x06,
	0x86, 0xba, 0x06, 0x45, 0x0a, 0x30, 0x0a, 0x09, 0x70, 0xf1, 0x04, 0x87, 0x3a, 0x05, 0x3b, 0x07,
	0x79, 0x06, 0x80, 0x25, 0x99, 0xac, 0xec, 0xb9, 0x58, 0x4a, 0x40, 0x5a, 0xe4, 0xf4, 0x20, 0xa6,
	0xeb, 0xc1, 0x4c, 0xd7, 0xf3, 0x4d, 0x2d, 0xc9, 0xc8, 0x4f, 0xf1, 0x2f, 0x28, 0xc9, 0xcc, 0xcf,
	0x2b, 0x96, 0xd8, 0xf1, 0x87, 0x59, 0x81, 0x51, 0x83, 0xcf, 0x88, 0x53, 0x0f, 0xe4, 0x2a, 0x90,
	0x2d, 0x41, 0x60, 0x8d, 0x4e, 0xac, 0x51, 0x20, 0xc7, 0x01, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0xb3, 0xf8, 0x5c, 0xad, 0x00, 0x00, 0x00,
}
