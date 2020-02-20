// Code generated by protoc-gen-go. DO NOT EDIT.
// source: price.proto

package pricerpc

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

type GetSubscriptionPricesParams struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductSku           string   `protobuf:"bytes,2,opt,name=product_sku,json=productSku,proto3" json:"product_sku,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSubscriptionPricesParams) Reset()         { *m = GetSubscriptionPricesParams{} }
func (m *GetSubscriptionPricesParams) String() string { return proto.CompactTextString(m) }
func (*GetSubscriptionPricesParams) ProtoMessage()    {}
func (*GetSubscriptionPricesParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_94c214910ad4a3eb, []int{0}
}

func (m *GetSubscriptionPricesParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSubscriptionPricesParams.Unmarshal(m, b)
}
func (m *GetSubscriptionPricesParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSubscriptionPricesParams.Marshal(b, m, deterministic)
}
func (m *GetSubscriptionPricesParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSubscriptionPricesParams.Merge(m, src)
}
func (m *GetSubscriptionPricesParams) XXX_Size() int {
	return xxx_messageInfo_GetSubscriptionPricesParams.Size(m)
}
func (m *GetSubscriptionPricesParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSubscriptionPricesParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetSubscriptionPricesParams proto.InternalMessageInfo

func (m *GetSubscriptionPricesParams) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *GetSubscriptionPricesParams) GetProductSku() string {
	if m != nil {
		return m.ProductSku
	}
	return ""
}

type GetSubscriptionPricesResponse struct {
	Prices               []*Price `protobuf:"bytes,1,rep,name=prices,proto3" json:"prices,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSubscriptionPricesResponse) Reset()         { *m = GetSubscriptionPricesResponse{} }
func (m *GetSubscriptionPricesResponse) String() string { return proto.CompactTextString(m) }
func (*GetSubscriptionPricesResponse) ProtoMessage()    {}
func (*GetSubscriptionPricesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94c214910ad4a3eb, []int{1}
}

func (m *GetSubscriptionPricesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSubscriptionPricesResponse.Unmarshal(m, b)
}
func (m *GetSubscriptionPricesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSubscriptionPricesResponse.Marshal(b, m, deterministic)
}
func (m *GetSubscriptionPricesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSubscriptionPricesResponse.Merge(m, src)
}
func (m *GetSubscriptionPricesResponse) XXX_Size() int {
	return xxx_messageInfo_GetSubscriptionPricesResponse.Size(m)
}
func (m *GetSubscriptionPricesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSubscriptionPricesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSubscriptionPricesResponse proto.InternalMessageInfo

func (m *GetSubscriptionPricesResponse) GetPrices() []*Price {
	if m != nil {
		return m.Prices
	}
	return nil
}

type EndPriceParams struct {
	Params               *GetPriceParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	EndAtDay             int32           `protobuf:"varint,2,opt,name=end_at_day,json=endAtDay,proto3" json:"end_at_day,omitempty"`
	EndAtMonth           int32           `protobuf:"varint,3,opt,name=end_at_month,json=endAtMonth,proto3" json:"end_at_month,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EndPriceParams) Reset()         { *m = EndPriceParams{} }
func (m *EndPriceParams) String() string { return proto.CompactTextString(m) }
func (*EndPriceParams) ProtoMessage()    {}
func (*EndPriceParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_94c214910ad4a3eb, []int{2}
}

func (m *EndPriceParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndPriceParams.Unmarshal(m, b)
}
func (m *EndPriceParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndPriceParams.Marshal(b, m, deterministic)
}
func (m *EndPriceParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndPriceParams.Merge(m, src)
}
func (m *EndPriceParams) XXX_Size() int {
	return xxx_messageInfo_EndPriceParams.Size(m)
}
func (m *EndPriceParams) XXX_DiscardUnknown() {
	xxx_messageInfo_EndPriceParams.DiscardUnknown(m)
}

var xxx_messageInfo_EndPriceParams proto.InternalMessageInfo

func (m *EndPriceParams) GetParams() *GetPriceParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *EndPriceParams) GetEndAtDay() int32 {
	if m != nil {
		return m.EndAtDay
	}
	return 0
}

func (m *EndPriceParams) GetEndAtMonth() int32 {
	if m != nil {
		return m.EndAtMonth
	}
	return 0
}

type Price struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Price                uint32   `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	Discount             uint32   `protobuf:"varint,3,opt,name=discount,proto3" json:"discount,omitempty"`
	StartAt              string   `protobuf:"bytes,4,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt                string   `protobuf:"bytes,5,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	IsSubscription       bool     `protobuf:"varint,6,opt,name=is_subscription,json=isSubscription,proto3" json:"is_subscription,omitempty"`
	ThirdPartyId         int32    `protobuf:"varint,7,opt,name=third_party_id,json=thirdPartyId,proto3" json:"third_party_id,omitempty"`
	PlanId               int64    `protobuf:"varint,8,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	ProductId            int64    `protobuf:"varint,9,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Price) Reset()         { *m = Price{} }
func (m *Price) String() string { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()    {}
func (*Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_94c214910ad4a3eb, []int{3}
}

func (m *Price) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Price.Unmarshal(m, b)
}
func (m *Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Price.Marshal(b, m, deterministic)
}
func (m *Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Price.Merge(m, src)
}
func (m *Price) XXX_Size() int {
	return xxx_messageInfo_Price.Size(m)
}
func (m *Price) XXX_DiscardUnknown() {
	xxx_messageInfo_Price.DiscardUnknown(m)
}

var xxx_messageInfo_Price proto.InternalMessageInfo

func (m *Price) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Price) GetPrice() uint32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Price) GetDiscount() uint32 {
	if m != nil {
		return m.Discount
	}
	return 0
}

func (m *Price) GetStartAt() string {
	if m != nil {
		return m.StartAt
	}
	return ""
}

func (m *Price) GetEndAt() string {
	if m != nil {
		return m.EndAt
	}
	return ""
}

func (m *Price) GetIsSubscription() bool {
	if m != nil {
		return m.IsSubscription
	}
	return false
}

func (m *Price) GetThirdPartyId() int32 {
	if m != nil {
		return m.ThirdPartyId
	}
	return 0
}

func (m *Price) GetPlanId() int64 {
	if m != nil {
		return m.PlanId
	}
	return 0
}

func (m *Price) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

type CreatePriceParams struct {
	StartAtDay           int32    `protobuf:"varint,1,opt,name=start_at_day,json=startAtDay,proto3" json:"start_at_day,omitempty"`
	StartAtMonth         int32    `protobuf:"varint,2,opt,name=start_at_month,json=startAtMonth,proto3" json:"start_at_month,omitempty"`
	EndAtDay             int32    `protobuf:"varint,3,opt,name=end_at_day,json=endAtDay,proto3" json:"end_at_day,omitempty"`
	EndAtMonth           int32    `protobuf:"varint,4,opt,name=end_at_month,json=endAtMonth,proto3" json:"end_at_month,omitempty"`
	Price                uint32   `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	Discount             uint32   `protobuf:"varint,6,opt,name=discount,proto3" json:"discount,omitempty"`
	ProductId            int64    `protobuf:"varint,7,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	PlanId               int64    `protobuf:"varint,8,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	ThirdPartyId         int32    `protobuf:"varint,9,opt,name=third_party_id,json=thirdPartyId,proto3" json:"third_party_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePriceParams) Reset()         { *m = CreatePriceParams{} }
func (m *CreatePriceParams) String() string { return proto.CompactTextString(m) }
func (*CreatePriceParams) ProtoMessage()    {}
func (*CreatePriceParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_94c214910ad4a3eb, []int{4}
}

func (m *CreatePriceParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePriceParams.Unmarshal(m, b)
}
func (m *CreatePriceParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePriceParams.Marshal(b, m, deterministic)
}
func (m *CreatePriceParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePriceParams.Merge(m, src)
}
func (m *CreatePriceParams) XXX_Size() int {
	return xxx_messageInfo_CreatePriceParams.Size(m)
}
func (m *CreatePriceParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePriceParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePriceParams proto.InternalMessageInfo

func (m *CreatePriceParams) GetStartAtDay() int32 {
	if m != nil {
		return m.StartAtDay
	}
	return 0
}

func (m *CreatePriceParams) GetStartAtMonth() int32 {
	if m != nil {
		return m.StartAtMonth
	}
	return 0
}

func (m *CreatePriceParams) GetEndAtDay() int32 {
	if m != nil {
		return m.EndAtDay
	}
	return 0
}

func (m *CreatePriceParams) GetEndAtMonth() int32 {
	if m != nil {
		return m.EndAtMonth
	}
	return 0
}

func (m *CreatePriceParams) GetPrice() uint32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *CreatePriceParams) GetDiscount() uint32 {
	if m != nil {
		return m.Discount
	}
	return 0
}

func (m *CreatePriceParams) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *CreatePriceParams) GetPlanId() int64 {
	if m != nil {
		return m.PlanId
	}
	return 0
}

func (m *CreatePriceParams) GetThirdPartyId() int32 {
	if m != nil {
		return m.ThirdPartyId
	}
	return 0
}

type GetPriceParams struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductSku           string   `protobuf:"bytes,2,opt,name=product_sku,json=productSku,proto3" json:"product_sku,omitempty"`
	PriceId              int64    `protobuf:"varint,3,opt,name=price_id,json=priceId,proto3" json:"price_id,omitempty"`
	PlanId               int64    `protobuf:"varint,4,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPriceParams) Reset()         { *m = GetPriceParams{} }
func (m *GetPriceParams) String() string { return proto.CompactTextString(m) }
func (*GetPriceParams) ProtoMessage()    {}
func (*GetPriceParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_94c214910ad4a3eb, []int{5}
}

func (m *GetPriceParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPriceParams.Unmarshal(m, b)
}
func (m *GetPriceParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPriceParams.Marshal(b, m, deterministic)
}
func (m *GetPriceParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPriceParams.Merge(m, src)
}
func (m *GetPriceParams) XXX_Size() int {
	return xxx_messageInfo_GetPriceParams.Size(m)
}
func (m *GetPriceParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPriceParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetPriceParams proto.InternalMessageInfo

func (m *GetPriceParams) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *GetPriceParams) GetProductSku() string {
	if m != nil {
		return m.ProductSku
	}
	return ""
}

func (m *GetPriceParams) GetPriceId() int64 {
	if m != nil {
		return m.PriceId
	}
	return 0
}

func (m *GetPriceParams) GetPlanId() int64 {
	if m != nil {
		return m.PlanId
	}
	return 0
}

func init() {
	proto.RegisterType((*GetSubscriptionPricesParams)(nil), "pepeunlimited.products.GetSubscriptionPricesParams")
	proto.RegisterType((*GetSubscriptionPricesResponse)(nil), "pepeunlimited.products.GetSubscriptionPricesResponse")
	proto.RegisterType((*EndPriceParams)(nil), "pepeunlimited.products.EndPriceParams")
	proto.RegisterType((*Price)(nil), "pepeunlimited.products.Price")
	proto.RegisterType((*CreatePriceParams)(nil), "pepeunlimited.products.CreatePriceParams")
	proto.RegisterType((*GetPriceParams)(nil), "pepeunlimited.products.GetPriceParams")
}

func init() { proto.RegisterFile("price.proto", fileDescriptor_94c214910ad4a3eb) }

var fileDescriptor_94c214910ad4a3eb = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5f, 0x8b, 0xd3, 0x40,
	0x10, 0x27, 0xcd, 0x25, 0x4d, 0xa7, 0xbd, 0x88, 0x8b, 0xa7, 0xb1, 0x5a, 0x0c, 0xe1, 0xd0, 0xfa,
	0xd2, 0x87, 0x1e, 0xf7, 0x2a, 0x9c, 0x7f, 0x38, 0xfa, 0x20, 0xd4, 0x14, 0x14, 0x04, 0x09, 0xb9,
	0xee, 0xc2, 0x2d, 0x77, 0x4d, 0xc2, 0xee, 0x46, 0xe8, 0xab, 0xfa, 0x01, 0x04, 0x3f, 0x86, 0x5f,
	0x52, 0x76, 0x92, 0xd4, 0x34, 0x6d, 0xee, 0x4e, 0x7c, 0xcb, 0xcc, 0xec, 0x4c, 0x7e, 0xf3, 0xfb,
	0xcd, 0x0c, 0xf4, 0x33, 0xc1, 0x97, 0x6c, 0x92, 0x89, 0x54, 0xa5, 0xe4, 0x61, 0xc6, 0x32, 0x96,
	0x27, 0xd7, 0x7c, 0xc5, 0x15, 0xa3, 0xda, 0x49, 0xf3, 0xa5, 0x92, 0xc1, 0x17, 0x78, 0x72, 0xce,
	0xd4, 0x22, 0xbf, 0x90, 0x4b, 0xc1, 0x33, 0xc5, 0xd3, 0x64, 0xae, 0xb3, 0xe4, 0x3c, 0x16, 0xf1,
	0x4a, 0x92, 0x11, 0x40, 0xf9, 0x34, 0xe2, 0xd4, 0x33, 0x7c, 0x63, 0x6c, 0x86, 0xbd, 0xd2, 0x33,
	0xa3, 0xe4, 0x99, 0xfe, 0x49, 0x11, 0x96, 0x57, 0xb9, 0xd7, 0xf1, 0x8d, 0x71, 0x2f, 0xac, 0x32,
	0x16, 0x57, 0x79, 0xf0, 0x11, 0x46, 0x7b, 0xcb, 0x87, 0x4c, 0x66, 0x69, 0x22, 0x19, 0x39, 0x05,
	0x1b, 0x61, 0x4a, 0xcf, 0xf0, 0xcd, 0x71, 0x7f, 0x3a, 0x9a, 0xec, 0x07, 0x3a, 0xc1, 0xbc, 0xb0,
	0x7c, 0x1c, 0xfc, 0x34, 0xc0, 0x7d, 0x97, 0x50, 0x74, 0x96, 0x50, 0x5f, 0x81, 0x9d, 0xe1, 0x17,
	0xc2, 0xec, 0x4f, 0x9f, 0xb7, 0x55, 0x3a, 0x67, 0xaa, 0x96, 0x17, 0x96, 0x59, 0xe4, 0x29, 0x00,
	0x4b, 0x68, 0x14, 0xab, 0x88, 0xc6, 0x6b, 0x6c, 0xc5, 0x0a, 0x1d, 0x96, 0xd0, 0x33, 0xf5, 0x36,
	0x5e, 0x13, 0x1f, 0x06, 0x65, 0x74, 0x95, 0x26, 0xea, 0xd2, 0x33, 0x31, 0x0e, 0x18, 0x7f, 0xaf,
	0x3d, 0xc1, 0xb7, 0x0e, 0x58, 0x58, 0x97, 0xb8, 0xd0, 0xd9, 0x90, 0xd5, 0xe1, 0x94, 0x3c, 0x00,
	0x0b, 0x61, 0x63, 0xd1, 0xc3, 0xb0, 0x30, 0xc8, 0x10, 0x1c, 0xca, 0xe5, 0x32, 0xcd, 0x13, 0x85,
	0xd5, 0x0e, 0xc3, 0x8d, 0x4d, 0x1e, 0x83, 0x23, 0x55, 0x2c, 0x54, 0x14, 0x2b, 0xef, 0x00, 0x49,
	0xed, 0xa2, 0x7d, 0xa6, 0xc8, 0x11, 0xd8, 0x05, 0x10, 0xcf, 0xc2, 0x80, 0x85, 0x10, 0xc8, 0x0b,
	0xb8, 0xc7, 0x65, 0x24, 0x6b, 0x44, 0x7b, 0xb6, 0x6f, 0x8c, 0x9d, 0xd0, 0xe5, 0xb2, 0x4e, 0x3f,
	0x39, 0x06, 0x57, 0x5d, 0x72, 0x41, 0xa3, 0x2c, 0x16, 0x6a, 0xad, 0x55, 0xed, 0x62, 0x2b, 0x03,
	0xf4, 0xce, 0xb5, 0x73, 0x46, 0xc9, 0x23, 0xe8, 0x66, 0xd7, 0x71, 0xa2, 0xc3, 0x0e, 0xf6, 0x61,
	0x6b, 0x73, 0x46, 0x1b, 0x03, 0xd1, 0x6b, 0x0c, 0x44, 0xf0, 0xbb, 0x03, 0xf7, 0xdf, 0x08, 0x16,
	0x2b, 0x56, 0x97, 0xc6, 0x87, 0x41, 0xd5, 0x0e, 0x92, 0x6b, 0x14, 0xe4, 0x95, 0x2d, 0x69, 0x7a,
	0x8f, 0xc1, 0xdd, 0xbc, 0x28, 0x08, 0x2e, 0x04, 0x18, 0x94, 0x6f, 0x90, 0xe2, 0x86, 0x44, 0xe6,
	0x2d, 0x12, 0x1d, 0x34, 0x25, 0xfa, 0x2b, 0x84, 0xd5, 0x26, 0x84, 0xdd, 0x10, 0x62, 0xbb, 0xdd,
	0x6e, 0x73, 0xfe, 0x5b, 0x69, 0xda, 0x65, 0xb9, 0xb7, 0xcb, 0x72, 0xf0, 0xc3, 0x00, 0x77, 0x7b,
	0x1a, 0xff, 0x77, 0xe1, 0xf4, 0xe4, 0x60, 0x57, 0x3a, 0xdb, 0xc4, 0xec, 0x2e, 0xda, 0xdb, 0x60,
	0x0f, 0xea, 0x60, 0xa7, 0xbf, 0x4c, 0x18, 0x20, 0x86, 0x05, 0x13, 0x5f, 0x35, 0x23, 0x9f, 0xa0,
	0x5f, 0x13, 0x91, 0xbc, 0x6c, 0xdb, 0xa4, 0x1d, 0xa5, 0x87, 0x37, 0xaf, 0x2f, 0xf9, 0x00, 0x4e,
	0xd5, 0x2f, 0xb9, 0xe3, 0x7e, 0xde, 0xa1, 0x64, 0x75, 0x08, 0xda, 0x4b, 0x6e, 0x9f, 0x8a, 0xdb,
	0x4a, 0x7e, 0x37, 0xe0, 0x68, 0xef, 0xd5, 0x22, 0x27, 0x37, 0x60, 0x6e, 0xbb, 0xa1, 0xc3, 0xd3,
	0x7f, 0x4a, 0xaa, 0x2e, 0xe3, 0x6b, 0xf8, 0x5c, 0x28, 0x29, 0xb2, 0xe5, 0x85, 0x8d, 0x47, 0xfc,
	0xe4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0x16, 0x14, 0x98, 0xd3, 0x05, 0x00, 0x00,
}
