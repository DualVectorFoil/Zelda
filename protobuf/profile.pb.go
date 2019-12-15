// Code generated by protoc-gen-go. DO NOT EDIT.
// source: profile.proto

package protobuf

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

type ProfileInfo struct {
	PhoneNum             *string  `protobuf:"bytes,1,req,name=phoneNum" json:"phoneNum,omitempty"`
	AvatarUrl            *string  `protobuf:"bytes,2,req,name=avatarUrl" json:"avatarUrl,omitempty"`
	UserName             *string  `protobuf:"bytes,3,req,name=userName" json:"userName,omitempty"`
	Locale               *string  `protobuf:"bytes,4,req,name=locale" json:"locale,omitempty"`
	Bio                  *string  `protobuf:"bytes,5,req,name=bio" json:"bio,omitempty"`
	Followers            *int32   `protobuf:"varint,6,req,name=followers" json:"followers,omitempty"`
	Following            *int32   `protobuf:"varint,7,req,name=following" json:"following,omitempty"`
	ArtworkCount         *int32   `protobuf:"varint,8,req,name=artworkCount" json:"artworkCount,omitempty"`
	Pwd                  *string  `protobuf:"bytes,9,req,name=pwd" json:"pwd,omitempty"`
	RegisteredAt         *int64   `protobuf:"varint,10,req,name=registeredAt" json:"registeredAt,omitempty"`
	LastLoginAt          *int64   `protobuf:"varint,11,req,name=lastLoginAt" json:"lastLoginAt,omitempty"`
	Token                *string  `protobuf:"bytes,12,req,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileInfo) Reset()         { *m = ProfileInfo{} }
func (m *ProfileInfo) String() string { return proto.CompactTextString(m) }
func (*ProfileInfo) ProtoMessage()    {}
func (*ProfileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{0}
}

func (m *ProfileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileInfo.Unmarshal(m, b)
}
func (m *ProfileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileInfo.Marshal(b, m, deterministic)
}
func (m *ProfileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileInfo.Merge(m, src)
}
func (m *ProfileInfo) XXX_Size() int {
	return xxx_messageInfo_ProfileInfo.Size(m)
}
func (m *ProfileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileInfo proto.InternalMessageInfo

func (m *ProfileInfo) GetPhoneNum() string {
	if m != nil && m.PhoneNum != nil {
		return *m.PhoneNum
	}
	return ""
}

func (m *ProfileInfo) GetAvatarUrl() string {
	if m != nil && m.AvatarUrl != nil {
		return *m.AvatarUrl
	}
	return ""
}

func (m *ProfileInfo) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *ProfileInfo) GetLocale() string {
	if m != nil && m.Locale != nil {
		return *m.Locale
	}
	return ""
}

func (m *ProfileInfo) GetBio() string {
	if m != nil && m.Bio != nil {
		return *m.Bio
	}
	return ""
}

func (m *ProfileInfo) GetFollowers() int32 {
	if m != nil && m.Followers != nil {
		return *m.Followers
	}
	return 0
}

func (m *ProfileInfo) GetFollowing() int32 {
	if m != nil && m.Following != nil {
		return *m.Following
	}
	return 0
}

func (m *ProfileInfo) GetArtworkCount() int32 {
	if m != nil && m.ArtworkCount != nil {
		return *m.ArtworkCount
	}
	return 0
}

func (m *ProfileInfo) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *ProfileInfo) GetRegisteredAt() int64 {
	if m != nil && m.RegisteredAt != nil {
		return *m.RegisteredAt
	}
	return 0
}

func (m *ProfileInfo) GetLastLoginAt() int64 {
	if m != nil && m.LastLoginAt != nil {
		return *m.LastLoginAt
	}
	return 0
}

func (m *ProfileInfo) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*ProfileInfo)(nil), "protobuf.ProfileInfo")
}

func init() { proto.RegisterFile("profile.proto", fileDescriptor_744bf7a47b381504) }

var fileDescriptor_744bf7a47b381504 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8d, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0x23, 0x48, 0x85, 0xa1, 0x26, 0x66, 0x63, 0xcc, 0xc4, 0x78, 0x20, 0x3d, 0x71, 0xf2,
	0x1d, 0x1a, 0x4f, 0x26, 0xa6, 0x31, 0x24, 0x3e, 0xc0, 0x56, 0x06, 0xdc, 0x74, 0xbb, 0x43, 0x86,
	0x45, 0x1e, 0xd6, 0x97, 0x31, 0x2c, 0x5a, 0xda, 0x13, 0xfc, 0xdf, 0x37, 0xff, 0xfe, 0x70, 0xdb,
	0x09, 0x37, 0xc6, 0xd2, 0x73, 0x27, 0xec, 0x59, 0xa5, 0xe1, 0xb3, 0x1f, 0x9a, 0xcd, 0x4f, 0x04,
	0xf9, 0xfb, 0xec, 0x5e, 0x5d, 0xc3, 0xea, 0x11, 0xd2, 0xee, 0x8b, 0x1d, 0xed, 0x86, 0x23, 0x5e,
	0x15, 0x51, 0x99, 0x55, 0xa7, 0xac, 0x9e, 0x20, 0xd3, 0xdf, 0xda, 0x6b, 0xf9, 0x10, 0x8b, 0x51,
	0x90, 0x0b, 0x98, 0x9a, 0x43, 0x4f, 0xb2, 0xd3, 0x47, 0xc2, 0x78, 0x6e, 0xfe, 0x67, 0xf5, 0x00,
	0x2b, 0xcb, 0x9f, 0xda, 0x12, 0x5e, 0x07, 0xf3, 0x97, 0xd4, 0x1d, 0xc4, 0x7b, 0xc3, 0x98, 0x04,
	0x38, 0xfd, 0x4e, 0x1b, 0x0d, 0x5b, 0xcb, 0x23, 0x49, 0x8f, 0xab, 0x22, 0x2a, 0x93, 0x6a, 0x01,
	0x8b, 0x35, 0xae, 0xc5, 0x9b, 0x73, 0x6b, 0x5c, 0xab, 0x36, 0xb0, 0xd6, 0xe2, 0x47, 0x96, 0xc3,
	0x0b, 0x0f, 0xce, 0x63, 0x1a, 0x0e, 0x2e, 0xd8, 0xb4, 0xd8, 0x8d, 0x35, 0x66, 0xf3, 0x62, 0x37,
	0xd6, 0x53, 0x4b, 0xa8, 0x35, 0xbd, 0x27, 0xa1, 0x7a, 0xeb, 0x11, 0x8a, 0xa8, 0x8c, 0xab, 0x0b,
	0xa6, 0x0a, 0xc8, 0xad, 0xee, 0xfd, 0x1b, 0xb7, 0xc6, 0x6d, 0x3d, 0xe6, 0xe1, 0xe4, 0x1c, 0xa9,
	0x7b, 0x48, 0x3c, 0x1f, 0xc8, 0xe1, 0x3a, 0xbc, 0x3c, 0x87, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x35, 0x70, 0x77, 0x40, 0x77, 0x01, 0x00, 0x00,
}