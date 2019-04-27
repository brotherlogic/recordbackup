// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recordbackup.proto

package recordbackup

import (
	fmt "fmt"
	proto1 "github.com/brotherlogic/recordcollection/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	Metadata             []*proto1.ReleaseMetadata `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eeef516d983bdc6, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetMetadata() []*proto1.ReleaseMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "recordbackup.Config")
}

func init() { proto.RegisterFile("recordbackup.proto", fileDescriptor_4eeef516d983bdc6) }

var fileDescriptor_4eeef516d983bdc6 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4a, 0x4d, 0xce,
	0x2f, 0x4a, 0x49, 0x4a, 0x4c, 0xce, 0x2e, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x41, 0x16, 0x93, 0x72, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f,
	0x2a, 0xca, 0x2f, 0xc9, 0x48, 0x2d, 0xca, 0xc9, 0x4f, 0xcf, 0x4c, 0xd6, 0x87, 0xa8, 0x4a, 0xce,
	0xcf, 0xc9, 0x49, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x07, 0xeb, 0xc6, 0x10, 0x86, 0x18, 0xaa,
	0xe4, 0xce, 0xc5, 0xe6, 0x9c, 0x9f, 0x97, 0x96, 0x99, 0x2e, 0x64, 0xcb, 0xc5, 0x91, 0x9b, 0x5a,
	0x92, 0x98, 0x92, 0x58, 0x92, 0x28, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xa8, 0x87, 0xa1,
	0x29, 0x28, 0x35, 0x27, 0x35, 0xb1, 0x38, 0xd5, 0x17, 0xaa, 0x30, 0x08, 0xae, 0x25, 0x89, 0x0d,
	0x6c, 0x9e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x81, 0x62, 0xa7, 0x11, 0xba, 0x00, 0x00, 0x00,
}