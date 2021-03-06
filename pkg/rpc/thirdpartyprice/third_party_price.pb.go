// Code generated by protoc-gen-go. DO NOT EDIT.
// source: third_party_price.proto

package thirdpartyprice

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

type EndThirdPartyPriceParams struct {
	Params               *GetThirdPartyPriceParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	EndAtMonth           int32                     `protobuf:"varint,2,opt,name=end_at_month,json=endAtMonth,proto3" json:"end_at_month,omitempty"`
	EndAtDay             int32                     `protobuf:"varint,3,opt,name=end_at_day,json=endAtDay,proto3" json:"end_at_day,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *EndThirdPartyPriceParams) Reset()         { *m = EndThirdPartyPriceParams{} }
func (m *EndThirdPartyPriceParams) String() string { return proto.CompactTextString(m) }
func (*EndThirdPartyPriceParams) ProtoMessage()    {}
func (*EndThirdPartyPriceParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b2355bf6f7ba95, []int{0}
}

func (m *EndThirdPartyPriceParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndThirdPartyPriceParams.Unmarshal(m, b)
}
func (m *EndThirdPartyPriceParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndThirdPartyPriceParams.Marshal(b, m, deterministic)
}
func (m *EndThirdPartyPriceParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndThirdPartyPriceParams.Merge(m, src)
}
func (m *EndThirdPartyPriceParams) XXX_Size() int {
	return xxx_messageInfo_EndThirdPartyPriceParams.Size(m)
}
func (m *EndThirdPartyPriceParams) XXX_DiscardUnknown() {
	xxx_messageInfo_EndThirdPartyPriceParams.DiscardUnknown(m)
}

var xxx_messageInfo_EndThirdPartyPriceParams proto.InternalMessageInfo

func (m *EndThirdPartyPriceParams) GetParams() *GetThirdPartyPriceParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *EndThirdPartyPriceParams) GetEndAtMonth() int32 {
	if m != nil {
		return m.EndAtMonth
	}
	return 0
}

func (m *EndThirdPartyPriceParams) GetEndAtDay() int32 {
	if m != nil {
		return m.EndAtDay
	}
	return 0
}

type GetThirdPartyPriceParams struct {
	ThirdPartyPriceId       int64    `protobuf:"varint,1,opt,name=third_party_price_id,json=thirdPartyPriceId,proto3" json:"third_party_price_id,omitempty"`
	InAppPurchaseSku        string   `protobuf:"bytes,2,opt,name=in_app_purchase_sku,json=inAppPurchaseSku,proto3" json:"in_app_purchase_sku,omitempty"`
	GoogleBillingServiceSku string   `protobuf:"bytes,3,opt,name=google_billing_service_sku,json=googleBillingServiceSku,proto3" json:"google_billing_service_sku,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *GetThirdPartyPriceParams) Reset()         { *m = GetThirdPartyPriceParams{} }
func (m *GetThirdPartyPriceParams) String() string { return proto.CompactTextString(m) }
func (*GetThirdPartyPriceParams) ProtoMessage()    {}
func (*GetThirdPartyPriceParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b2355bf6f7ba95, []int{1}
}

func (m *GetThirdPartyPriceParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetThirdPartyPriceParams.Unmarshal(m, b)
}
func (m *GetThirdPartyPriceParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetThirdPartyPriceParams.Marshal(b, m, deterministic)
}
func (m *GetThirdPartyPriceParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetThirdPartyPriceParams.Merge(m, src)
}
func (m *GetThirdPartyPriceParams) XXX_Size() int {
	return xxx_messageInfo_GetThirdPartyPriceParams.Size(m)
}
func (m *GetThirdPartyPriceParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetThirdPartyPriceParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetThirdPartyPriceParams proto.InternalMessageInfo

func (m *GetThirdPartyPriceParams) GetThirdPartyPriceId() int64 {
	if m != nil {
		return m.ThirdPartyPriceId
	}
	return 0
}

func (m *GetThirdPartyPriceParams) GetInAppPurchaseSku() string {
	if m != nil {
		return m.InAppPurchaseSku
	}
	return ""
}

func (m *GetThirdPartyPriceParams) GetGoogleBillingServiceSku() string {
	if m != nil {
		return m.GoogleBillingServiceSku
	}
	return ""
}

type CreateThirdPartyPriceParams struct {
	InAppPurchaseSku        string   `protobuf:"bytes,1,opt,name=in_app_purchase_sku,json=inAppPurchaseSku,proto3" json:"in_app_purchase_sku,omitempty"`
	GoogleBillingServiceSku string   `protobuf:"bytes,2,opt,name=google_billing_service_sku,json=googleBillingServiceSku,proto3" json:"google_billing_service_sku,omitempty"`
	StartAtMonth            int32    `protobuf:"varint,3,opt,name=start_at_month,json=startAtMonth,proto3" json:"start_at_month,omitempty"`
	StartAtDay              int32    `protobuf:"varint,4,opt,name=start_at_day,json=startAtDay,proto3" json:"start_at_day,omitempty"`
	Type                    string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *CreateThirdPartyPriceParams) Reset()         { *m = CreateThirdPartyPriceParams{} }
func (m *CreateThirdPartyPriceParams) String() string { return proto.CompactTextString(m) }
func (*CreateThirdPartyPriceParams) ProtoMessage()    {}
func (*CreateThirdPartyPriceParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b2355bf6f7ba95, []int{2}
}

func (m *CreateThirdPartyPriceParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateThirdPartyPriceParams.Unmarshal(m, b)
}
func (m *CreateThirdPartyPriceParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateThirdPartyPriceParams.Marshal(b, m, deterministic)
}
func (m *CreateThirdPartyPriceParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateThirdPartyPriceParams.Merge(m, src)
}
func (m *CreateThirdPartyPriceParams) XXX_Size() int {
	return xxx_messageInfo_CreateThirdPartyPriceParams.Size(m)
}
func (m *CreateThirdPartyPriceParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateThirdPartyPriceParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreateThirdPartyPriceParams proto.InternalMessageInfo

func (m *CreateThirdPartyPriceParams) GetInAppPurchaseSku() string {
	if m != nil {
		return m.InAppPurchaseSku
	}
	return ""
}

func (m *CreateThirdPartyPriceParams) GetGoogleBillingServiceSku() string {
	if m != nil {
		return m.GoogleBillingServiceSku
	}
	return ""
}

func (m *CreateThirdPartyPriceParams) GetStartAtMonth() int32 {
	if m != nil {
		return m.StartAtMonth
	}
	return 0
}

func (m *CreateThirdPartyPriceParams) GetStartAtDay() int32 {
	if m != nil {
		return m.StartAtDay
	}
	return 0
}

func (m *CreateThirdPartyPriceParams) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type GetThirdPartyPricesParams struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetThirdPartyPricesParams) Reset()         { *m = GetThirdPartyPricesParams{} }
func (m *GetThirdPartyPricesParams) String() string { return proto.CompactTextString(m) }
func (*GetThirdPartyPricesParams) ProtoMessage()    {}
func (*GetThirdPartyPricesParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b2355bf6f7ba95, []int{3}
}

func (m *GetThirdPartyPricesParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetThirdPartyPricesParams.Unmarshal(m, b)
}
func (m *GetThirdPartyPricesParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetThirdPartyPricesParams.Marshal(b, m, deterministic)
}
func (m *GetThirdPartyPricesParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetThirdPartyPricesParams.Merge(m, src)
}
func (m *GetThirdPartyPricesParams) XXX_Size() int {
	return xxx_messageInfo_GetThirdPartyPricesParams.Size(m)
}
func (m *GetThirdPartyPricesParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetThirdPartyPricesParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetThirdPartyPricesParams proto.InternalMessageInfo

func (m *GetThirdPartyPricesParams) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type ThirdPartyPrice struct {
	Id                      int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	InAppPurchaseSku        string   `protobuf:"bytes,2,opt,name=in_app_purchase_sku,json=inAppPurchaseSku,proto3" json:"in_app_purchase_sku,omitempty"`
	GoogleBillingServiceSku string   `protobuf:"bytes,3,opt,name=google_billing_service_sku,json=googleBillingServiceSku,proto3" json:"google_billing_service_sku,omitempty"`
	StartAt                 string   `protobuf:"bytes,4,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt                   string   `protobuf:"bytes,5,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	Type                    string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ThirdPartyPrice) Reset()         { *m = ThirdPartyPrice{} }
func (m *ThirdPartyPrice) String() string { return proto.CompactTextString(m) }
func (*ThirdPartyPrice) ProtoMessage()    {}
func (*ThirdPartyPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b2355bf6f7ba95, []int{4}
}

func (m *ThirdPartyPrice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThirdPartyPrice.Unmarshal(m, b)
}
func (m *ThirdPartyPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThirdPartyPrice.Marshal(b, m, deterministic)
}
func (m *ThirdPartyPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThirdPartyPrice.Merge(m, src)
}
func (m *ThirdPartyPrice) XXX_Size() int {
	return xxx_messageInfo_ThirdPartyPrice.Size(m)
}
func (m *ThirdPartyPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_ThirdPartyPrice.DiscardUnknown(m)
}

var xxx_messageInfo_ThirdPartyPrice proto.InternalMessageInfo

func (m *ThirdPartyPrice) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ThirdPartyPrice) GetInAppPurchaseSku() string {
	if m != nil {
		return m.InAppPurchaseSku
	}
	return ""
}

func (m *ThirdPartyPrice) GetGoogleBillingServiceSku() string {
	if m != nil {
		return m.GoogleBillingServiceSku
	}
	return ""
}

func (m *ThirdPartyPrice) GetStartAt() string {
	if m != nil {
		return m.StartAt
	}
	return ""
}

func (m *ThirdPartyPrice) GetEndAt() string {
	if m != nil {
		return m.EndAt
	}
	return ""
}

func (m *ThirdPartyPrice) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type GetThirdPartyPricesResponse struct {
	ThirdPartyPrices     []*ThirdPartyPrice `protobuf:"bytes,1,rep,name=third_party_prices,json=thirdPartyPrices,proto3" json:"third_party_prices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetThirdPartyPricesResponse) Reset()         { *m = GetThirdPartyPricesResponse{} }
func (m *GetThirdPartyPricesResponse) String() string { return proto.CompactTextString(m) }
func (*GetThirdPartyPricesResponse) ProtoMessage()    {}
func (*GetThirdPartyPricesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b2355bf6f7ba95, []int{5}
}

func (m *GetThirdPartyPricesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetThirdPartyPricesResponse.Unmarshal(m, b)
}
func (m *GetThirdPartyPricesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetThirdPartyPricesResponse.Marshal(b, m, deterministic)
}
func (m *GetThirdPartyPricesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetThirdPartyPricesResponse.Merge(m, src)
}
func (m *GetThirdPartyPricesResponse) XXX_Size() int {
	return xxx_messageInfo_GetThirdPartyPricesResponse.Size(m)
}
func (m *GetThirdPartyPricesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetThirdPartyPricesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetThirdPartyPricesResponse proto.InternalMessageInfo

func (m *GetThirdPartyPricesResponse) GetThirdPartyPrices() []*ThirdPartyPrice {
	if m != nil {
		return m.ThirdPartyPrices
	}
	return nil
}

func init() {
	proto.RegisterType((*EndThirdPartyPriceParams)(nil), "pepeunlimited.products.EndThirdPartyPriceParams")
	proto.RegisterType((*GetThirdPartyPriceParams)(nil), "pepeunlimited.products.GetThirdPartyPriceParams")
	proto.RegisterType((*CreateThirdPartyPriceParams)(nil), "pepeunlimited.products.CreateThirdPartyPriceParams")
	proto.RegisterType((*GetThirdPartyPricesParams)(nil), "pepeunlimited.products.GetThirdPartyPricesParams")
	proto.RegisterType((*ThirdPartyPrice)(nil), "pepeunlimited.products.ThirdPartyPrice")
	proto.RegisterType((*GetThirdPartyPricesResponse)(nil), "pepeunlimited.products.GetThirdPartyPricesResponse")
}

func init() { proto.RegisterFile("third_party_price.proto", fileDescriptor_31b2355bf6f7ba95) }

var fileDescriptor_31b2355bf6f7ba95 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xd6, 0x66, 0x9b, 0xd0, 0x4e, 0xa3, 0xfe, 0x4c, 0x69, 0xbb, 0x4d, 0x39, 0x44, 0x11, 0x12,
	0xb9, 0x90, 0x42, 0x72, 0xe4, 0xd4, 0x52, 0x04, 0x1c, 0x90, 0xa2, 0x0d, 0x5c, 0xb8, 0x58, 0x6e,
	0xd6, 0x4a, 0xac, 0x24, 0xbb, 0x96, 0xed, 0x45, 0xda, 0x03, 0x6f, 0xc3, 0x3b, 0xf0, 0x20, 0x3c,
	0x03, 0x77, 0x1e, 0x01, 0xd9, 0xeb, 0x84, 0x92, 0xdd, 0xad, 0x12, 0x0e, 0xdc, 0x2c, 0xcf, 0x7c,
	0x9e, 0xf9, 0xe6, 0xfb, 0x3c, 0x70, 0xae, 0xa7, 0x5c, 0x46, 0x44, 0x50, 0xa9, 0x33, 0x22, 0x24,
	0x1f, 0xb3, 0x9e, 0x90, 0x89, 0x4e, 0xf0, 0x4c, 0x30, 0xc1, 0xd2, 0x78, 0xce, 0x17, 0x5c, 0xb3,
	0xc8, 0x5c, 0x46, 0xe9, 0x58, 0xab, 0xce, 0x37, 0x0f, 0x82, 0x37, 0x71, 0xf4, 0xd1, 0xc0, 0x86,
	0x06, 0x35, 0x34, 0xa0, 0x21, 0x95, 0x74, 0xa1, 0xf0, 0x1d, 0x34, 0x84, 0x3d, 0x05, 0x5e, 0xdb,
	0xeb, 0xee, 0xf7, 0x5f, 0xf4, 0xca, 0x5f, 0xe9, 0xbd, 0x65, 0xba, 0xf4, 0x85, 0xd0, 0xe1, 0xb1,
	0x0d, 0x4d, 0x16, 0x47, 0x84, 0x6a, 0xb2, 0x48, 0x62, 0x3d, 0x0d, 0x6a, 0x6d, 0xaf, 0x5b, 0x0f,
	0x81, 0xc5, 0xd1, 0xb5, 0xfe, 0x60, 0x6e, 0xf0, 0x09, 0x80, 0xcb, 0x88, 0x68, 0x16, 0xf8, 0x36,
	0xbe, 0x6b, 0xe3, 0xb7, 0x34, 0xeb, 0x7c, 0xf7, 0x20, 0xa8, 0x2a, 0x82, 0x57, 0xf0, 0xb8, 0x40,
	0x9b, 0xf0, 0xc8, 0x36, 0xed, 0x87, 0xc7, 0xfa, 0x6f, 0xd0, 0xfb, 0x08, 0x9f, 0xc3, 0x09, 0x8f,
	0x09, 0x15, 0x82, 0x88, 0x54, 0x8e, 0xa7, 0x54, 0x31, 0xa2, 0x66, 0xa9, 0x6d, 0x6a, 0x2f, 0x3c,
	0xe2, 0xf1, 0xb5, 0x10, 0x43, 0x17, 0x18, 0xcd, 0x52, 0x7c, 0x05, 0xad, 0x49, 0x92, 0x4c, 0xe6,
	0x8c, 0xdc, 0xf1, 0xf9, 0x9c, 0xc7, 0x13, 0xa2, 0x98, 0xfc, 0x62, 0x8a, 0x18, 0x94, 0x6f, 0x51,
	0xe7, 0x79, 0xc6, 0x4d, 0x9e, 0x30, 0xca, 0xe3, 0xa3, 0x59, 0xda, 0xf9, 0xe9, 0xc1, 0xe5, 0x6b,
	0xc9, 0xa8, 0x66, 0xe5, 0xcd, 0x57, 0xf4, 0xe2, 0xfd, 0x53, 0x2f, 0xb5, 0x07, 0x7b, 0xc1, 0xa7,
	0x70, 0xa0, 0x34, 0x95, 0xfa, 0x8f, 0x0e, 0xf9, 0x9c, 0x9b, 0xf6, 0x76, 0xa9, 0x44, 0x1b, 0x9a,
	0xab, 0x2c, 0xa3, 0xc5, 0x4e, 0xae, 0x95, 0xcb, 0xb9, 0xa5, 0x19, 0x22, 0xec, 0xe8, 0x4c, 0xb0,
	0xa0, 0x6e, 0xcb, 0xd9, 0x73, 0xe7, 0x0a, 0x2e, 0x8a, 0x02, 0x29, 0x47, 0x72, 0x09, 0xf0, 0xee,
	0x01, 0x7e, 0x78, 0x70, 0xb8, 0x96, 0x8e, 0x07, 0x50, 0x5b, 0xe9, 0x56, 0xe3, 0xff, 0x55, 0x28,
	0xbc, 0x80, 0xdd, 0x25, 0x6d, 0x4b, 0x79, 0x2f, 0x7c, 0xe4, 0x28, 0xe3, 0x29, 0x34, 0x72, 0x6f,
	0x3a, 0xc6, 0x75, 0xeb, 0xcb, 0x15, 0xab, 0xc6, 0x3d, 0x56, 0x1a, 0x2e, 0x4b, 0xc6, 0x10, 0x32,
	0x25, 0x92, 0x58, 0x31, 0xfc, 0x04, 0x58, 0xb0, 0xaa, 0xf9, 0x5d, 0x7e, 0x77, 0xbf, 0xff, 0xac,
	0xea, 0x77, 0xad, 0xbd, 0x16, 0x1e, 0xad, 0x39, 0x5a, 0xf5, 0x7f, 0xf9, 0x70, 0xb6, 0x96, 0xe5,
	0x98, 0x61, 0x0a, 0xa7, 0xa5, 0xf6, 0xc3, 0x41, 0x55, 0xb9, 0x07, 0xdc, 0xda, 0xda, 0xb4, 0x47,
	0xfc, 0x0a, 0x27, 0x25, 0x73, 0xc0, 0x97, 0x9b, 0x6f, 0x10, 0xe7, 0x9d, 0xd6, 0x60, 0x0b, 0xc8,
	0x6a, 0xce, 0x09, 0x60, 0x31, 0x8c, 0x5b, 0xef, 0xaf, 0xcd, 0xf9, 0x26, 0x80, 0xc5, 0x35, 0x5a,
	0x5d, 0xb0, 0x6a, 0xe5, 0x6e, 0x5c, 0xf0, 0xe6, 0xf8, 0xf3, 0xa1, 0xb5, 0x81, 0x35, 0x92, 0xf5,
	0xd1, 0x5d, 0xc3, 0xae, 0xfa, 0xc1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xc4, 0xd4, 0xcd,
	0x05, 0x06, 0x00, 0x00,
}
