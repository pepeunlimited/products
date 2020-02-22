// Code generated by protoc-gen-go. DO NOT EDIT.
// source: third_party.proto

package thirdpartypricerpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type EndThirdPartyParams struct {
	Params               *GetThirdPartyParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	EndAtMonth           int32                `protobuf:"varint,2,opt,name=end_at_month,json=endAtMonth,proto3" json:"end_at_month,omitempty"`
	EndAtDay             int32                `protobuf:"varint,3,opt,name=end_at_day,json=endAtDay,proto3" json:"end_at_day,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EndThirdPartyParams) Reset()         { *m = EndThirdPartyParams{} }
func (m *EndThirdPartyParams) String() string { return proto.CompactTextString(m) }
func (*EndThirdPartyParams) ProtoMessage()    {}
func (*EndThirdPartyParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_66e319316a0c7196, []int{0}
}

func (m *EndThirdPartyParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndThirdPartyParams.Unmarshal(m, b)
}
func (m *EndThirdPartyParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndThirdPartyParams.Marshal(b, m, deterministic)
}
func (m *EndThirdPartyParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndThirdPartyParams.Merge(m, src)
}
func (m *EndThirdPartyParams) XXX_Size() int {
	return xxx_messageInfo_EndThirdPartyParams.Size(m)
}
func (m *EndThirdPartyParams) XXX_DiscardUnknown() {
	xxx_messageInfo_EndThirdPartyParams.DiscardUnknown(m)
}

var xxx_messageInfo_EndThirdPartyParams proto.InternalMessageInfo

func (m *EndThirdPartyParams) GetParams() *GetThirdPartyParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *EndThirdPartyParams) GetEndAtMonth() int32 {
	if m != nil {
		return m.EndAtMonth
	}
	return 0
}

func (m *EndThirdPartyParams) GetEndAtDay() int32 {
	if m != nil {
		return m.EndAtDay
	}
	return 0
}

type GetThirdPartyParams struct {
	ThirdPartyId            int32    `protobuf:"varint,1,opt,name=third_party_id,json=thirdPartyId,proto3" json:"third_party_id,omitempty"`
	InAppPurchaseSku        string   `protobuf:"bytes,2,opt,name=in_app_purchase_sku,json=inAppPurchaseSku,proto3" json:"in_app_purchase_sku,omitempty"`
	GoogleBillingServiceSku string   `protobuf:"bytes,3,opt,name=google_billing_service_sku,json=googleBillingServiceSku,proto3" json:"google_billing_service_sku,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *GetThirdPartyParams) Reset()         { *m = GetThirdPartyParams{} }
func (m *GetThirdPartyParams) String() string { return proto.CompactTextString(m) }
func (*GetThirdPartyParams) ProtoMessage()    {}
func (*GetThirdPartyParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_66e319316a0c7196, []int{1}
}

func (m *GetThirdPartyParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetThirdPartyParams.Unmarshal(m, b)
}
func (m *GetThirdPartyParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetThirdPartyParams.Marshal(b, m, deterministic)
}
func (m *GetThirdPartyParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetThirdPartyParams.Merge(m, src)
}
func (m *GetThirdPartyParams) XXX_Size() int {
	return xxx_messageInfo_GetThirdPartyParams.Size(m)
}
func (m *GetThirdPartyParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetThirdPartyParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetThirdPartyParams proto.InternalMessageInfo

func (m *GetThirdPartyParams) GetThirdPartyId() int32 {
	if m != nil {
		return m.ThirdPartyId
	}
	return 0
}

func (m *GetThirdPartyParams) GetInAppPurchaseSku() string {
	if m != nil {
		return m.InAppPurchaseSku
	}
	return ""
}

func (m *GetThirdPartyParams) GetGoogleBillingServiceSku() string {
	if m != nil {
		return m.GoogleBillingServiceSku
	}
	return ""
}

type CreateThirdPartyParams struct {
	InAppPurchaseSku        string   `protobuf:"bytes,1,opt,name=in_app_purchase_sku,json=inAppPurchaseSku,proto3" json:"in_app_purchase_sku,omitempty"`
	GoogleBillingServiceSku string   `protobuf:"bytes,2,opt,name=google_billing_service_sku,json=googleBillingServiceSku,proto3" json:"google_billing_service_sku,omitempty"`
	StartAtMonth            int32    `protobuf:"varint,3,opt,name=start_at_month,json=startAtMonth,proto3" json:"start_at_month,omitempty"`
	StartAtDay              int32    `protobuf:"varint,4,opt,name=start_at_day,json=startAtDay,proto3" json:"start_at_day,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *CreateThirdPartyParams) Reset()         { *m = CreateThirdPartyParams{} }
func (m *CreateThirdPartyParams) String() string { return proto.CompactTextString(m) }
func (*CreateThirdPartyParams) ProtoMessage()    {}
func (*CreateThirdPartyParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_66e319316a0c7196, []int{2}
}

func (m *CreateThirdPartyParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateThirdPartyParams.Unmarshal(m, b)
}
func (m *CreateThirdPartyParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateThirdPartyParams.Marshal(b, m, deterministic)
}
func (m *CreateThirdPartyParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateThirdPartyParams.Merge(m, src)
}
func (m *CreateThirdPartyParams) XXX_Size() int {
	return xxx_messageInfo_CreateThirdPartyParams.Size(m)
}
func (m *CreateThirdPartyParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateThirdPartyParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreateThirdPartyParams proto.InternalMessageInfo

func (m *CreateThirdPartyParams) GetInAppPurchaseSku() string {
	if m != nil {
		return m.InAppPurchaseSku
	}
	return ""
}

func (m *CreateThirdPartyParams) GetGoogleBillingServiceSku() string {
	if m != nil {
		return m.GoogleBillingServiceSku
	}
	return ""
}

func (m *CreateThirdPartyParams) GetStartAtMonth() int32 {
	if m != nil {
		return m.StartAtMonth
	}
	return 0
}

func (m *CreateThirdPartyParams) GetStartAtDay() int32 {
	if m != nil {
		return m.StartAtDay
	}
	return 0
}

type GetThirdPartiesParams struct {
	Show                 bool     `protobuf:"varint,1,opt,name=show,proto3" json:"show,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetThirdPartiesParams) Reset()         { *m = GetThirdPartiesParams{} }
func (m *GetThirdPartiesParams) String() string { return proto.CompactTextString(m) }
func (*GetThirdPartiesParams) ProtoMessage()    {}
func (*GetThirdPartiesParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_66e319316a0c7196, []int{3}
}

func (m *GetThirdPartiesParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetThirdPartiesParams.Unmarshal(m, b)
}
func (m *GetThirdPartiesParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetThirdPartiesParams.Marshal(b, m, deterministic)
}
func (m *GetThirdPartiesParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetThirdPartiesParams.Merge(m, src)
}
func (m *GetThirdPartiesParams) XXX_Size() int {
	return xxx_messageInfo_GetThirdPartiesParams.Size(m)
}
func (m *GetThirdPartiesParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetThirdPartiesParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetThirdPartiesParams proto.InternalMessageInfo

func (m *GetThirdPartiesParams) GetShow() bool {
	if m != nil {
		return m.Show
	}
	return false
}

type ThirdParty struct {
	Id                      int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	InAppPurchaseSku        string   `protobuf:"bytes,2,opt,name=in_app_purchase_sku,json=inAppPurchaseSku,proto3" json:"in_app_purchase_sku,omitempty"`
	GoogleBillingServiceSku string   `protobuf:"bytes,3,opt,name=google_billing_service_sku,json=googleBillingServiceSku,proto3" json:"google_billing_service_sku,omitempty"`
	StartAt                 string   `protobuf:"bytes,4,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt                   string   `protobuf:"bytes,5,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ThirdParty) Reset()         { *m = ThirdParty{} }
func (m *ThirdParty) String() string { return proto.CompactTextString(m) }
func (*ThirdParty) ProtoMessage()    {}
func (*ThirdParty) Descriptor() ([]byte, []int) {
	return fileDescriptor_66e319316a0c7196, []int{4}
}

func (m *ThirdParty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThirdParty.Unmarshal(m, b)
}
func (m *ThirdParty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThirdParty.Marshal(b, m, deterministic)
}
func (m *ThirdParty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThirdParty.Merge(m, src)
}
func (m *ThirdParty) XXX_Size() int {
	return xxx_messageInfo_ThirdParty.Size(m)
}
func (m *ThirdParty) XXX_DiscardUnknown() {
	xxx_messageInfo_ThirdParty.DiscardUnknown(m)
}

var xxx_messageInfo_ThirdParty proto.InternalMessageInfo

func (m *ThirdParty) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ThirdParty) GetInAppPurchaseSku() string {
	if m != nil {
		return m.InAppPurchaseSku
	}
	return ""
}

func (m *ThirdParty) GetGoogleBillingServiceSku() string {
	if m != nil {
		return m.GoogleBillingServiceSku
	}
	return ""
}

func (m *ThirdParty) GetStartAt() string {
	if m != nil {
		return m.StartAt
	}
	return ""
}

func (m *ThirdParty) GetEndAt() string {
	if m != nil {
		return m.EndAt
	}
	return ""
}

type GetThirdPartiesResponse struct {
	ThirdParties         []*ThirdParty `protobuf:"bytes,1,rep,name=third_parties,json=thirdParties,proto3" json:"third_parties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetThirdPartiesResponse) Reset()         { *m = GetThirdPartiesResponse{} }
func (m *GetThirdPartiesResponse) String() string { return proto.CompactTextString(m) }
func (*GetThirdPartiesResponse) ProtoMessage()    {}
func (*GetThirdPartiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66e319316a0c7196, []int{5}
}

func (m *GetThirdPartiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetThirdPartiesResponse.Unmarshal(m, b)
}
func (m *GetThirdPartiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetThirdPartiesResponse.Marshal(b, m, deterministic)
}
func (m *GetThirdPartiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetThirdPartiesResponse.Merge(m, src)
}
func (m *GetThirdPartiesResponse) XXX_Size() int {
	return xxx_messageInfo_GetThirdPartiesResponse.Size(m)
}
func (m *GetThirdPartiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetThirdPartiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetThirdPartiesResponse proto.InternalMessageInfo

func (m *GetThirdPartiesResponse) GetThirdParties() []*ThirdParty {
	if m != nil {
		return m.ThirdParties
	}
	return nil
}

func init() {
	proto.RegisterType((*EndThirdPartyParams)(nil), "pepeunlimited.products.EndThirdPartyParams")
	proto.RegisterType((*GetThirdPartyParams)(nil), "pepeunlimited.products.GetThirdPartyParams")
	proto.RegisterType((*CreateThirdPartyParams)(nil), "pepeunlimited.products.CreateThirdPartyParams")
	proto.RegisterType((*GetThirdPartiesParams)(nil), "pepeunlimited.products.GetThirdPartiesParams")
	proto.RegisterType((*ThirdParty)(nil), "pepeunlimited.products.ThirdParty")
	proto.RegisterType((*GetThirdPartiesResponse)(nil), "pepeunlimited.products.GetThirdPartiesResponse")
}

func init() { proto.RegisterFile("third_party.proto", fileDescriptor_66e319316a0c7196) }

var fileDescriptor_66e319316a0c7196 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcd, 0x6e, 0xda, 0x40,
	0x10, 0x96, 0x21, 0x50, 0x32, 0x81, 0xfc, 0x2c, 0x4a, 0x42, 0x69, 0x55, 0x21, 0xab, 0x07, 0xa4,
	0x28, 0x20, 0xd1, 0x63, 0x4f, 0x49, 0x5a, 0x45, 0x3d, 0x54, 0x42, 0xa6, 0xa7, 0x5e, 0x9c, 0x05,
	0x4f, 0x60, 0x05, 0xac, 0xb7, 0xbb, 0xeb, 0x46, 0xbc, 0x4a, 0x5f, 0xa0, 0xf7, 0x1e, 0xfb, 0x0c,
	0x7d, 0xa8, 0xca, 0xbb, 0xc6, 0x26, 0x2e, 0x54, 0x6e, 0x0f, 0xb9, 0x59, 0x3b, 0xdf, 0x78, 0xbe,
	0xef, 0x9b, 0x1f, 0x38, 0xd1, 0x33, 0x26, 0x03, 0x5f, 0x50, 0xa9, 0x57, 0x3d, 0x21, 0x43, 0x1d,
	0x92, 0x33, 0x81, 0x02, 0x23, 0xbe, 0x60, 0x4b, 0xa6, 0x31, 0x88, 0x1f, 0x83, 0x68, 0xa2, 0x55,
	0xfb, 0xc5, 0x34, 0x0c, 0xa7, 0x0b, 0xec, 0x1b, 0xd4, 0x38, 0xba, 0xef, 0xe3, 0x52, 0xac, 0x93,
	0xda, 0xaf, 0xf2, 0xc1, 0x07, 0x49, 0x85, 0x40, 0xa9, 0x6c, 0xdc, 0xfd, 0xe6, 0x40, 0xf3, 0x3d,
	0x0f, 0x3e, 0xc5, 0xd5, 0x86, 0x71, 0xb1, 0x21, 0x95, 0x74, 0xa9, 0xc8, 0x0d, 0x54, 0x85, 0xf9,
	0x6a, 0x39, 0x1d, 0xa7, 0x7b, 0x30, 0xb8, 0xe8, 0x6d, 0xaf, 0xde, 0xbb, 0x45, 0x9d, 0x4f, 0xf6,
	0x92, 0x54, 0xd2, 0x81, 0x3a, 0xf2, 0xc0, 0xa7, 0xda, 0x5f, 0x86, 0x5c, 0xcf, 0x5a, 0xa5, 0x8e,
	0xd3, 0xad, 0x78, 0x80, 0x3c, 0xb8, 0xd2, 0x1f, 0xe3, 0x17, 0xf2, 0x12, 0x20, 0x41, 0x04, 0x74,
	0xd5, 0x2a, 0x9b, 0x78, 0xcd, 0xc4, 0xdf, 0xd1, 0x95, 0xfb, 0xdd, 0x81, 0xe6, 0x96, 0xff, 0x93,
	0xd7, 0x70, 0xb8, 0x61, 0x8f, 0xcf, 0x02, 0x43, 0xb2, 0xe2, 0xd5, 0x75, 0x8a, 0xfc, 0x10, 0x90,
	0x4b, 0x68, 0x32, 0xee, 0x53, 0x21, 0x7c, 0x11, 0xc9, 0xc9, 0x8c, 0x2a, 0xf4, 0xd5, 0x3c, 0x32,
	0x24, 0xf6, 0xbd, 0x63, 0xc6, 0xaf, 0x84, 0x18, 0x26, 0x81, 0xd1, 0x3c, 0x22, 0x6f, 0xa1, 0x6d,
	0xbd, 0xf2, 0xc7, 0x6c, 0xb1, 0x60, 0x7c, 0xea, 0x2b, 0x94, 0x5f, 0xd9, 0xc4, 0x66, 0x95, 0x4d,
	0xd6, 0xb9, 0x45, 0x5c, 0x5b, 0xc0, 0xc8, 0xc6, 0x47, 0xf3, 0xc8, 0xfd, 0xe5, 0xc0, 0xd9, 0x8d,
	0x44, 0xaa, 0xf1, 0x0f, 0xb2, 0x3b, 0x68, 0x38, 0xff, 0x45, 0xa3, 0xf4, 0x57, 0x1a, 0xb1, 0x31,
	0x4a, 0x53, 0xa9, 0x33, 0xcb, 0xad, 0xa5, 0x75, 0xf3, 0xba, 0x36, 0xbd, 0x03, 0xf5, 0x14, 0x15,
	0xdb, 0xbe, 0x67, 0xdb, 0x92, 0x60, 0x62, 0xe3, 0x2f, 0xe0, 0x74, 0xd3, 0x77, 0x86, 0x2a, 0x11,
	0x43, 0x60, 0x4f, 0xcd, 0xc2, 0x07, 0xc3, 0xbe, 0xe6, 0x99, 0x6f, 0xf7, 0xa7, 0x03, 0x90, 0xa9,
	0x26, 0x87, 0x50, 0x4a, 0x1b, 0x52, 0x62, 0x4f, 0xda, 0x06, 0xf2, 0x1c, 0x6a, 0x6b, 0x65, 0x46,
	0xd5, 0xbe, 0xf7, 0x2c, 0x51, 0x45, 0x4e, 0xa1, 0x6a, 0x27, 0xad, 0x55, 0x31, 0x81, 0x8a, 0x99,
	0x32, 0x77, 0x0c, 0xe7, 0x39, 0xa5, 0x1e, 0x2a, 0x11, 0x72, 0x85, 0xe4, 0x16, 0x1a, 0xd9, 0x94,
	0x31, 0x8c, 0x37, 0xa1, 0xdc, 0x3d, 0x18, 0xb8, 0xbb, 0x36, 0x21, 0xf3, 0x60, 0x63, 0x10, 0x19,
	0xaa, 0xc1, 0x8f, 0x32, 0x9c, 0x64, 0xc1, 0x84, 0x2e, 0xb9, 0x87, 0xe3, 0xfc, 0xc4, 0x90, 0xde,
	0xae, 0x7f, 0x6f, 0x9f, 0xad, 0x76, 0x01, 0x2e, 0xe4, 0x0b, 0x1c, 0xe5, 0x14, 0x92, 0xcb, 0x22,
	0xcb, 0x9c, 0x36, 0xbd, 0xdd, 0x2f, 0x08, 0x4f, 0x9d, 0xbb, 0x83, 0xc6, 0xa3, 0xb5, 0x25, 0xff,
	0x72, 0x3d, 0x0a, 0x89, 0xba, 0x83, 0xc6, 0xa3, 0xab, 0xb5, 0xbb, 0xc2, 0x96, 0xe3, 0x56, 0xa4,
	0xc2, 0xf5, 0xd1, 0x67, 0xdb, 0x7d, 0x73, 0x62, 0xa4, 0x98, 0x8c, 0xab, 0xe6, 0x60, 0xbe, 0xf9,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xd5, 0x9e, 0xf2, 0x9a, 0x05, 0x00, 0x00,
}