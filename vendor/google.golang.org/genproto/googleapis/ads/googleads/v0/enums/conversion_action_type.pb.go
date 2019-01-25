// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/conversion_action_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// Possible types of a conversion action.
type ConversionActionTypeEnum_ConversionActionType int32

const (
	// Not specified.
	ConversionActionTypeEnum_UNSPECIFIED ConversionActionTypeEnum_ConversionActionType = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionActionTypeEnum_UNKNOWN ConversionActionTypeEnum_ConversionActionType = 1
	// Conversions that occur when a user clicks on an ad's call extension.
	ConversionActionTypeEnum_AD_CALL ConversionActionTypeEnum_ConversionActionType = 2
	// Conversions that occur when a user on a mobile device clicks a phone
	// number.
	ConversionActionTypeEnum_CLICK_TO_CALL ConversionActionTypeEnum_ConversionActionType = 3
	// Conversions that occur when a user downloads a mobile app from the Google
	// Play Store.
	ConversionActionTypeEnum_GOOGLE_PLAY_DOWNLOAD ConversionActionTypeEnum_ConversionActionType = 4
	// Conversions that occur when a user makes a purchase in an app through
	// Android billing.
	ConversionActionTypeEnum_GOOGLE_PLAY_IN_APP_PURCHASE ConversionActionTypeEnum_ConversionActionType = 5
	// Call conversions that are tracked by the advertiser and uploaded.
	ConversionActionTypeEnum_UPLOAD_CALLS ConversionActionTypeEnum_ConversionActionType = 6
	// Conversions that are tracked by the advertiser and uploaded with
	// attributed clicks.
	ConversionActionTypeEnum_UPLOAD_CLICKS ConversionActionTypeEnum_ConversionActionType = 7
	// Conversions that occur on a webpage.
	ConversionActionTypeEnum_WEBPAGE ConversionActionTypeEnum_ConversionActionType = 8
	// Conversions that occur when a user calls a dynamically-generated phone
	// number from an advertiser's website.
	ConversionActionTypeEnum_WEBSITE_CALL ConversionActionTypeEnum_ConversionActionType = 9
)

var ConversionActionTypeEnum_ConversionActionType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AD_CALL",
	3: "CLICK_TO_CALL",
	4: "GOOGLE_PLAY_DOWNLOAD",
	5: "GOOGLE_PLAY_IN_APP_PURCHASE",
	6: "UPLOAD_CALLS",
	7: "UPLOAD_CLICKS",
	8: "WEBPAGE",
	9: "WEBSITE_CALL",
}
var ConversionActionTypeEnum_ConversionActionType_value = map[string]int32{
	"UNSPECIFIED":                 0,
	"UNKNOWN":                     1,
	"AD_CALL":                     2,
	"CLICK_TO_CALL":               3,
	"GOOGLE_PLAY_DOWNLOAD":        4,
	"GOOGLE_PLAY_IN_APP_PURCHASE": 5,
	"UPLOAD_CALLS":                6,
	"UPLOAD_CLICKS":               7,
	"WEBPAGE":                     8,
	"WEBSITE_CALL":                9,
}

func (x ConversionActionTypeEnum_ConversionActionType) String() string {
	return proto.EnumName(ConversionActionTypeEnum_ConversionActionType_name, int32(x))
}
func (ConversionActionTypeEnum_ConversionActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_conversion_action_type_b6eff01350215c1e, []int{0, 0}
}

// Container for enum describing possible types of a conversion action.
type ConversionActionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionActionTypeEnum) Reset()         { *m = ConversionActionTypeEnum{} }
func (m *ConversionActionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionActionTypeEnum) ProtoMessage()    {}
func (*ConversionActionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_action_type_b6eff01350215c1e, []int{0}
}
func (m *ConversionActionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionActionTypeEnum.Unmarshal(m, b)
}
func (m *ConversionActionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionActionTypeEnum.Marshal(b, m, deterministic)
}
func (dst *ConversionActionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionActionTypeEnum.Merge(dst, src)
}
func (m *ConversionActionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionActionTypeEnum.Size(m)
}
func (m *ConversionActionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionActionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionActionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConversionActionTypeEnum)(nil), "google.ads.googleads.v0.enums.ConversionActionTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.ConversionActionTypeEnum_ConversionActionType", ConversionActionTypeEnum_ConversionActionType_name, ConversionActionTypeEnum_ConversionActionType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/conversion_action_type.proto", fileDescriptor_conversion_action_type_b6eff01350215c1e)
}

var fileDescriptor_conversion_action_type_b6eff01350215c1e = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xd1, 0x4e, 0xbb, 0x30,
	0x14, 0xc6, 0xff, 0x6c, 0x7f, 0x37, 0xed, 0x34, 0x56, 0xb2, 0x8b, 0x19, 0xb3, 0x18, 0xf7, 0x00,
	0x85, 0xc4, 0x3b, 0xbd, 0x2a, 0x50, 0x91, 0x8c, 0x40, 0x23, 0x63, 0x44, 0x43, 0x42, 0x70, 0x10,
	0xb2, 0x64, 0xa3, 0x64, 0xdd, 0x96, 0xec, 0x75, 0xbc, 0xf4, 0x51, 0x76, 0xeb, 0x8b, 0xf8, 0x08,
	0xa6, 0xc5, 0x2d, 0x5e, 0x4c, 0x6f, 0xe0, 0xeb, 0xf9, 0xfa, 0x3b, 0x5f, 0x7a, 0x0e, 0xb8, 0x2b,
	0x18, 0x2b, 0x66, 0xb9, 0x96, 0x66, 0x5c, 0xab, 0xa5, 0x50, 0x6b, 0x5d, 0xcb, 0xcb, 0xd5, 0x9c,
	0x6b, 0x13, 0x56, 0xae, 0xf3, 0x05, 0x9f, 0xb2, 0x32, 0x49, 0x27, 0x4b, 0xf1, 0x5b, 0x6e, 0xaa,
	0x1c, 0x55, 0x0b, 0xb6, 0x64, 0x6a, 0xbf, 0x06, 0x50, 0x9a, 0x71, 0xb4, 0x67, 0xd1, 0x5a, 0x47,
	0x92, 0x1d, 0x7c, 0x2a, 0xa0, 0x67, 0xee, 0x79, 0x2c, 0xf1, 0xd1, 0xa6, 0xca, 0x49, 0xb9, 0x9a,
	0x0f, 0x3e, 0x14, 0xd0, 0x3d, 0x64, 0xaa, 0xe7, 0xa0, 0x13, 0x7a, 0x01, 0x25, 0xa6, 0xf3, 0xe0,
	0x10, 0x0b, 0xfe, 0x53, 0x3b, 0xa0, 0x1d, 0x7a, 0x43, 0xcf, 0x8f, 0x3c, 0xa8, 0x88, 0x03, 0xb6,
	0x12, 0x13, 0xbb, 0x2e, 0x6c, 0xa8, 0x17, 0xe0, 0xcc, 0x74, 0x1d, 0x73, 0x98, 0x8c, 0xfc, 0xba,
	0xd4, 0x54, 0x7b, 0xa0, 0x6b, 0xfb, 0xbe, 0xed, 0x92, 0x84, 0xba, 0xf8, 0x39, 0xb1, 0xfc, 0xc8,
	0x73, 0x7d, 0x6c, 0xc1, 0xff, 0xea, 0x35, 0xb8, 0xfa, 0xe9, 0x38, 0x5e, 0x82, 0x29, 0x4d, 0x68,
	0xf8, 0x64, 0x3e, 0xe2, 0x80, 0xc0, 0x23, 0x15, 0x82, 0xd3, 0x90, 0x8a, 0xcb, 0xb2, 0x57, 0x00,
	0x5b, 0xa2, 0xff, 0xae, 0x22, 0x62, 0x02, 0xd8, 0x16, 0xf9, 0x11, 0x31, 0x28, 0xb6, 0x09, 0x3c,
	0x16, 0x44, 0x44, 0x8c, 0xc0, 0x19, 0x91, 0x3a, 0xfe, 0xc4, 0xd8, 0x2a, 0xe0, 0x66, 0xc2, 0xe6,
	0xe8, 0xcf, 0xc1, 0x18, 0x97, 0x87, 0x1e, 0x4e, 0xc5, 0x48, 0xa9, 0xf2, 0x62, 0x7c, 0xb3, 0x05,
	0x9b, 0xa5, 0x65, 0x81, 0xd8, 0xa2, 0xd0, 0x8a, 0xbc, 0x94, 0x03, 0xdf, 0x2d, 0xa8, 0x9a, 0xf2,
	0x5f, 0xf6, 0x75, 0x2f, 0xbf, 0x6f, 0x8d, 0xa6, 0x8d, 0xf1, 0x7b, 0xa3, 0x6f, 0xd7, 0xad, 0x70,
	0xc6, 0x51, 0x2d, 0x85, 0x1a, 0xeb, 0x48, 0x6c, 0x80, 0x6f, 0x77, 0x7e, 0x8c, 0x33, 0x1e, 0xef,
	0xfd, 0x78, 0xac, 0xc7, 0xd2, 0x7f, 0x6d, 0xc9, 0xd0, 0xdb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x2c, 0xd4, 0xfc, 0x8a, 0x23, 0x02, 0x00, 0x00,
}