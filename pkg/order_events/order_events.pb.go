// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: order_events.proto

package order_events

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderType int32

const (
	OrderType_PLACEHOLDER_ORDER_TYPE OrderType = 0
	OrderType_WITHDRAW               OrderType = 1
	OrderType_DEPOSIT                OrderType = 2
)

// Enum value maps for OrderType.
var (
	OrderType_name = map[int32]string{
		0: "PLACEHOLDER_ORDER_TYPE",
		1: "WITHDRAW",
		2: "DEPOSIT",
	}
	OrderType_value = map[string]int32{
		"PLACEHOLDER_ORDER_TYPE": 0,
		"WITHDRAW":               1,
		"DEPOSIT":                2,
	}
)

func (x OrderType) Enum() *OrderType {
	p := new(OrderType)
	*p = x
	return p
}

func (x OrderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderType) Descriptor() protoreflect.EnumDescriptor {
	return file_order_events_proto_enumTypes[0].Descriptor()
}

func (OrderType) Type() protoreflect.EnumType {
	return &file_order_events_proto_enumTypes[0]
}

func (x OrderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderType.Descriptor instead.
func (OrderType) EnumDescriptor() ([]byte, []int) {
	return file_order_events_proto_rawDescGZIP(), []int{0}
}

type OrderStatus int32

const (
	OrderStatus_PLACEHOLDER_ORDER_STATUS OrderStatus = 0
	OrderStatus_ACCEPTED                 OrderStatus = 1
	OrderStatus_CREATED                  OrderStatus = 2
	OrderStatus_PROCESSING               OrderStatus = 3
	OrderStatus_CONFIRMED                OrderStatus = 4
	OrderStatus_FAILED                   OrderStatus = 5
	OrderStatus_SUCCESS                  OrderStatus = 6
	OrderStatus_CANCELED                 OrderStatus = 7
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "PLACEHOLDER_ORDER_STATUS",
		1: "ACCEPTED",
		2: "CREATED",
		3: "PROCESSING",
		4: "CONFIRMED",
		5: "FAILED",
		6: "SUCCESS",
		7: "CANCELED",
	}
	OrderStatus_value = map[string]int32{
		"PLACEHOLDER_ORDER_STATUS": 0,
		"ACCEPTED":                 1,
		"CREATED":                  2,
		"PROCESSING":               3,
		"CONFIRMED":                4,
		"FAILED":                   5,
		"SUCCESS":                  6,
		"CANCELED":                 7,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_order_events_proto_enumTypes[1].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_order_events_proto_enumTypes[1]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_order_events_proto_rawDescGZIP(), []int{1}
}

type BcTxStatus int32

const (
	BcTxStatus_PLACEHOLDER_BC_TX_STATUS BcTxStatus = 0
	BcTxStatus_BC_TX_IN_MINING          BcTxStatus = 1
	BcTxStatus_BC_TX_SUCCESS            BcTxStatus = 2
	BcTxStatus_BC_TX_EXECUTION_FAILED   BcTxStatus = 3
)

// Enum value maps for BcTxStatus.
var (
	BcTxStatus_name = map[int32]string{
		0: "PLACEHOLDER_BC_TX_STATUS",
		1: "BC_TX_IN_MINING",
		2: "BC_TX_SUCCESS",
		3: "BC_TX_EXECUTION_FAILED",
	}
	BcTxStatus_value = map[string]int32{
		"PLACEHOLDER_BC_TX_STATUS": 0,
		"BC_TX_IN_MINING":          1,
		"BC_TX_SUCCESS":            2,
		"BC_TX_EXECUTION_FAILED":   3,
	}
)

func (x BcTxStatus) Enum() *BcTxStatus {
	p := new(BcTxStatus)
	*p = x
	return p
}

func (x BcTxStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BcTxStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_order_events_proto_enumTypes[2].Descriptor()
}

func (BcTxStatus) Type() protoreflect.EnumType {
	return &file_order_events_proto_enumTypes[2]
}

func (x BcTxStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BcTxStatus.Descriptor instead.
func (BcTxStatus) EnumDescriptor() ([]byte, []int) {
	return file_order_events_proto_rawDescGZIP(), []int{2}
}

type OrderIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID              string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	ExternalOrderUUID string `protobuf:"bytes,2,opt,name=ExternalOrderUUID,proto3" json:"ExternalOrderUUID,omitempty"`
	ProviderOrderUUID string `protobuf:"bytes,3,opt,name=ProviderOrderUUID,proto3" json:"ProviderOrderUUID,omitempty"`
}

func (x *OrderIdentity) Reset() {
	*x = OrderIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderIdentity) ProtoMessage() {}

func (x *OrderIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_order_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderIdentity.ProtoReflect.Descriptor instead.
func (*OrderIdentity) Descriptor() ([]byte, []int) {
	return file_order_events_proto_rawDescGZIP(), []int{0}
}

func (x *OrderIdentity) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *OrderIdentity) GetExternalOrderUUID() string {
	if x != nil {
		return x.ExternalOrderUUID
	}
	return ""
}

func (x *OrderIdentity) GetProviderOrderUUID() string {
	if x != nil {
		return x.ProviderOrderUUID
	}
	return ""
}

type BcTxIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxID string `protobuf:"bytes,1,opt,name=TxID,proto3" json:"TxID,omitempty"`
	Meta []byte `protobuf:"bytes,2,opt,name=Meta,proto3" json:"Meta,omitempty"`
}

func (x *BcTxIdentity) Reset() {
	*x = BcTxIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BcTxIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BcTxIdentity) ProtoMessage() {}

func (x *BcTxIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_order_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BcTxIdentity.ProtoReflect.Descriptor instead.
func (*BcTxIdentity) Descriptor() ([]byte, []int) {
	return file_order_events_proto_rawDescGZIP(), []int{1}
}

func (x *BcTxIdentity) GetTxID() string {
	if x != nil {
		return x.TxID
	}
	return ""
}

func (x *BcTxIdentity) GetMeta() []byte {
	if x != nil {
		return x.Meta
	}
	return nil
}

type BcTxOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index   uint32 `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	Amount  uint64 `protobuf:"varint,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *BcTxOperation) Reset() {
	*x = BcTxOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BcTxOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BcTxOperation) ProtoMessage() {}

func (x *BcTxOperation) ProtoReflect() protoreflect.Message {
	mi := &file_order_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BcTxOperation.ProtoReflect.Descriptor instead.
func (*BcTxOperation) Descriptor() ([]byte, []int) {
	return file_order_events_proto_rawDescGZIP(), []int{2}
}

func (x *BcTxOperation) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *BcTxOperation) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BcTxOperation) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BcTxInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkIdentifier uint32           `protobuf:"varint,1,opt,name=NetworkIdentifier,proto3" json:"NetworkIdentifier,omitempty"`
	BcTxIdentifier    *BcTxIdentity    `protobuf:"bytes,2,opt,name=BcTxIdentifier,proto3" json:"BcTxIdentifier,omitempty"`
	Confirmations     uint64           `protobuf:"varint,3,opt,name=Confirmations,proto3" json:"Confirmations,omitempty"`
	BlockNumber       uint64           `protobuf:"varint,4,opt,name=BlockNumber,proto3" json:"BlockNumber,omitempty"`
	Operations        []*BcTxOperation `protobuf:"bytes,5,rep,name=Operations,proto3" json:"Operations,omitempty"`
	Status            BcTxStatus       `protobuf:"varint,6,opt,name=Status,proto3,enum=order_events.BcTxStatus" json:"Status,omitempty"`
	RealFee           uint64           `protobuf:"varint,7,opt,name=RealFee,proto3" json:"RealFee,omitempty"`
	ExplorerURL       string           `protobuf:"bytes,8,opt,name=ExplorerURL,proto3" json:"ExplorerURL,omitempty"`
}

func (x *BcTxInfo) Reset() {
	*x = BcTxInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BcTxInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BcTxInfo) ProtoMessage() {}

func (x *BcTxInfo) ProtoReflect() protoreflect.Message {
	mi := &file_order_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BcTxInfo.ProtoReflect.Descriptor instead.
func (*BcTxInfo) Descriptor() ([]byte, []int) {
	return file_order_events_proto_rawDescGZIP(), []int{3}
}

func (x *BcTxInfo) GetNetworkIdentifier() uint32 {
	if x != nil {
		return x.NetworkIdentifier
	}
	return 0
}

func (x *BcTxInfo) GetBcTxIdentifier() *BcTxIdentity {
	if x != nil {
		return x.BcTxIdentifier
	}
	return nil
}

func (x *BcTxInfo) GetConfirmations() uint64 {
	if x != nil {
		return x.Confirmations
	}
	return 0
}

func (x *BcTxInfo) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *BcTxInfo) GetOperations() []*BcTxOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *BcTxInfo) GetStatus() BcTxStatus {
	if x != nil {
		return x.Status
	}
	return BcTxStatus_PLACEHOLDER_BC_TX_STATUS
}

func (x *BcTxInfo) GetRealFee() uint64 {
	if x != nil {
		return x.RealFee
	}
	return 0
}

func (x *BcTxInfo) GetExplorerURL() string {
	if x != nil {
		return x.ExplorerURL
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderIdentifier         *OrderIdentity `protobuf:"bytes,1,opt,name=OrderIdentifier,proto3" json:"OrderIdentifier,omitempty"`
	Type                    OrderType      `protobuf:"varint,2,opt,name=Type,proto3,enum=order_events.OrderType" json:"Type,omitempty"`
	Status                  OrderStatus    `protobuf:"varint,3,opt,name=Status,proto3,enum=order_events.OrderStatus" json:"Status,omitempty"`
	NetworkIdentifier       uint32         `protobuf:"varint,4,opt,name=NetworkIdentifier,proto3" json:"NetworkIdentifier,omitempty"`
	ProviderIdentifier      uint32         `protobuf:"varint,5,opt,name=ProviderIdentifier,proto3" json:"ProviderIdentifier,omitempty"`
	MerchantIdentifier      string         `protobuf:"bytes,6,opt,name=MerchantIdentifier,proto3" json:"MerchantIdentifier,omitempty"`
	Amount                  uint64         `protobuf:"varint,7,opt,name=Amount,proto3" json:"Amount,omitempty"`
	RealFee                 uint64         `protobuf:"varint,8,opt,name=RealFee,proto3" json:"RealFee,omitempty"`
	EstimationFee           uint64         `protobuf:"varint,9,opt,name=EstimationFee,proto3" json:"EstimationFee,omitempty"`
	EstimationFeeIdentifier string         `protobuf:"bytes,10,opt,name=EstimationFeeIdentifier,proto3" json:"EstimationFeeIdentifier,omitempty"`
	WalletIdentifier        string         `protobuf:"bytes,11,opt,name=WalletIdentifier,proto3" json:"WalletIdentifier,omitempty"`
	AddressIdentifier       string         `protobuf:"bytes,12,opt,name=AddressIdentifier,proto3" json:"AddressIdentifier,omitempty"`
	Currency                string         `protobuf:"bytes,13,opt,name=Currency,proto3" json:"Currency,omitempty"`
	BcTx                    []*BcTxInfo    `protobuf:"bytes,14,rep,name=BcTx,proto3" json:"BcTx,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_events_proto_rawDescGZIP(), []int{4}
}

func (x *Order) GetOrderIdentifier() *OrderIdentity {
	if x != nil {
		return x.OrderIdentifier
	}
	return nil
}

func (x *Order) GetType() OrderType {
	if x != nil {
		return x.Type
	}
	return OrderType_PLACEHOLDER_ORDER_TYPE
}

func (x *Order) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_PLACEHOLDER_ORDER_STATUS
}

func (x *Order) GetNetworkIdentifier() uint32 {
	if x != nil {
		return x.NetworkIdentifier
	}
	return 0
}

func (x *Order) GetProviderIdentifier() uint32 {
	if x != nil {
		return x.ProviderIdentifier
	}
	return 0
}

func (x *Order) GetMerchantIdentifier() string {
	if x != nil {
		return x.MerchantIdentifier
	}
	return ""
}

func (x *Order) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Order) GetRealFee() uint64 {
	if x != nil {
		return x.RealFee
	}
	return 0
}

func (x *Order) GetEstimationFee() uint64 {
	if x != nil {
		return x.EstimationFee
	}
	return 0
}

func (x *Order) GetEstimationFeeIdentifier() string {
	if x != nil {
		return x.EstimationFeeIdentifier
	}
	return ""
}

func (x *Order) GetWalletIdentifier() string {
	if x != nil {
		return x.WalletIdentifier
	}
	return ""
}

func (x *Order) GetAddressIdentifier() string {
	if x != nil {
		return x.AddressIdentifier
	}
	return ""
}

func (x *Order) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Order) GetBcTx() []*BcTxInfo {
	if x != nil {
		return x.BcTx
	}
	return nil
}

type OrderEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventIdentity      string `protobuf:"bytes,1,opt,name=EventIdentity,proto3" json:"EventIdentity,omitempty"`
	NetworkIdentifier  uint32 `protobuf:"varint,2,opt,name=NetworkIdentifier,proto3" json:"NetworkIdentifier,omitempty"`
	ProviderIdentifier uint32 `protobuf:"varint,3,opt,name=ProviderIdentifier,proto3" json:"ProviderIdentifier,omitempty"`
	MerchantIdentifier string `protobuf:"bytes,4,opt,name=MerchantIdentifier,proto3" json:"MerchantIdentifier,omitempty"`
	OrderIdentifier    string `protobuf:"bytes,5,opt,name=OrderIdentifier,proto3" json:"OrderIdentifier,omitempty"`
	Type               string `protobuf:"bytes,6,opt,name=Type,proto3" json:"Type,omitempty"`
	Producer           string `protobuf:"bytes,7,opt,name=Producer,proto3" json:"Producer,omitempty"`
	Message            string `protobuf:"bytes,8,opt,name=Message,proto3" json:"Message,omitempty"`
	OrderInfo          *Order `protobuf:"bytes,9,opt,name=OrderInfo,proto3" json:"OrderInfo,omitempty"`
}

func (x *OrderEvent) Reset() {
	*x = OrderEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderEvent) ProtoMessage() {}

func (x *OrderEvent) ProtoReflect() protoreflect.Message {
	mi := &file_order_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderEvent.ProtoReflect.Descriptor instead.
func (*OrderEvent) Descriptor() ([]byte, []int) {
	return file_order_events_proto_rawDescGZIP(), []int{5}
}

func (x *OrderEvent) GetEventIdentity() string {
	if x != nil {
		return x.EventIdentity
	}
	return ""
}

func (x *OrderEvent) GetNetworkIdentifier() uint32 {
	if x != nil {
		return x.NetworkIdentifier
	}
	return 0
}

func (x *OrderEvent) GetProviderIdentifier() uint32 {
	if x != nil {
		return x.ProviderIdentifier
	}
	return 0
}

func (x *OrderEvent) GetMerchantIdentifier() string {
	if x != nil {
		return x.MerchantIdentifier
	}
	return ""
}

func (x *OrderEvent) GetOrderIdentifier() string {
	if x != nil {
		return x.OrderIdentifier
	}
	return ""
}

func (x *OrderEvent) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OrderEvent) GetProducer() string {
	if x != nil {
		return x.Producer
	}
	return ""
}

func (x *OrderEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OrderEvent) GetOrderInfo() *Order {
	if x != nil {
		return x.OrderInfo
	}
	return nil
}

var File_order_events_proto protoreflect.FileDescriptor

var file_order_events_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x7f, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x55, 0x55, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55,
	0x55, 0x49, 0x44, 0x22, 0x36, 0x0a, 0x0c, 0x42, 0x63, 0x54, 0x78, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x78, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x54, 0x78, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x57, 0x0a, 0x0d, 0x42,
	0x63, 0x54, 0x78, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xef, 0x02, 0x0a, 0x08, 0x42, 0x63, 0x54, 0x78, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2c, 0x0a, 0x11, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x42, 0x0a, 0x0e, 0x42, 0x63, 0x54, 0x78, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x42, 0x63, 0x54, 0x78, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x0e, 0x42, 0x63, 0x54, 0x78, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0a, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x42,
	0x63, 0x54, 0x78, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x42, 0x63, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65,
	0x61, 0x6c, 0x46, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x52, 0x65, 0x61,
	0x6c, 0x46, 0x65, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72,
	0x55, 0x52, 0x4c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x45, 0x78, 0x70, 0x6c, 0x6f,
	0x72, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x22, 0xf0, 0x04, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x45, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x11, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x12, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x52, 0x65, 0x61, 0x6c, 0x46, 0x65, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x52, 0x65, 0x61, 0x6c, 0x46, 0x65, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x45, 0x73, 0x74, 0x69, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d,
	0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x12, 0x38, 0x0a,
	0x17, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17,
	0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2a, 0x0a,
	0x04, 0x42, 0x63, 0x54, 0x78, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x42, 0x63, 0x54, 0x78, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x42, 0x63, 0x54, 0x78, 0x22, 0xe7, 0x02, 0x0a, 0x0a, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2c,
	0x0a, 0x11, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x31, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2a, 0x42, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x16, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x5f,
	0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45,
	0x50, 0x4f, 0x53, 0x49, 0x54, 0x10, 0x02, 0x2a, 0x8c, 0x01, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x4c, 0x41, 0x43, 0x45,
	0x48, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x03,
	0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x04, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43,
	0x45, 0x4c, 0x45, 0x44, 0x10, 0x07, 0x2a, 0x6e, 0x0a, 0x0a, 0x42, 0x63, 0x54, 0x78, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c,
	0x44, 0x45, 0x52, 0x5f, 0x42, 0x43, 0x5f, 0x54, 0x58, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x43, 0x5f, 0x54, 0x58, 0x5f, 0x49, 0x4e, 0x5f, 0x4d,
	0x49, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x43, 0x5f, 0x54, 0x58,
	0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x42, 0x43,
	0x5f, 0x54, 0x58, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x77, 0x69, 0x7a, 0x65, 0x2d,
	0x74, 0x65, 0x63, 0x68, 0x2f, 0x62, 0x63, 0x2d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_events_proto_rawDescOnce sync.Once
	file_order_events_proto_rawDescData = file_order_events_proto_rawDesc
)

func file_order_events_proto_rawDescGZIP() []byte {
	file_order_events_proto_rawDescOnce.Do(func() {
		file_order_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_events_proto_rawDescData)
	})
	return file_order_events_proto_rawDescData
}

var file_order_events_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_order_events_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_order_events_proto_goTypes = []interface{}{
	(OrderType)(0),        // 0: order_events.OrderType
	(OrderStatus)(0),      // 1: order_events.OrderStatus
	(BcTxStatus)(0),       // 2: order_events.BcTxStatus
	(*OrderIdentity)(nil), // 3: order_events.OrderIdentity
	(*BcTxIdentity)(nil),  // 4: order_events.BcTxIdentity
	(*BcTxOperation)(nil), // 5: order_events.BcTxOperation
	(*BcTxInfo)(nil),      // 6: order_events.BcTxInfo
	(*Order)(nil),         // 7: order_events.Order
	(*OrderEvent)(nil),    // 8: order_events.OrderEvent
}
var file_order_events_proto_depIdxs = []int32{
	4, // 0: order_events.BcTxInfo.BcTxIdentifier:type_name -> order_events.BcTxIdentity
	5, // 1: order_events.BcTxInfo.Operations:type_name -> order_events.BcTxOperation
	2, // 2: order_events.BcTxInfo.Status:type_name -> order_events.BcTxStatus
	3, // 3: order_events.Order.OrderIdentifier:type_name -> order_events.OrderIdentity
	0, // 4: order_events.Order.Type:type_name -> order_events.OrderType
	1, // 5: order_events.Order.Status:type_name -> order_events.OrderStatus
	6, // 6: order_events.Order.BcTx:type_name -> order_events.BcTxInfo
	7, // 7: order_events.OrderEvent.OrderInfo:type_name -> order_events.Order
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_order_events_proto_init() }
func file_order_events_proto_init() {
	if File_order_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderIdentity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BcTxIdentity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BcTxOperation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BcTxInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_events_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_order_events_proto_goTypes,
		DependencyIndexes: file_order_events_proto_depIdxs,
		EnumInfos:         file_order_events_proto_enumTypes,
		MessageInfos:      file_order_events_proto_msgTypes,
	}.Build()
	File_order_events_proto = out.File
	file_order_events_proto_rawDesc = nil
	file_order_events_proto_goTypes = nil
	file_order_events_proto_depIdxs = nil
}
