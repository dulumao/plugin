// Code generated by protoc-gen-go. DO NOT EDIT.
// source: unfreeze.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	unfreeze.proto

It has these top-level messages:
	Unfreeze
	UnfreezeAction
	UnfreezeCreate
	UnfreezeWithdraw
	UnfreezeTerminate
	ReceiptUnfreeze
	QueryWithdraw
*/
package types

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

type Unfreeze struct {
	// 解冻交易ID（唯一识别码）
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
	// 开始时间
	StartTime int64 `protobuf:"varint,2,opt,name=startTime" json:"startTime,omitempty"`
	// 币种
	TokenName string `protobuf:"bytes,3,opt,name=tokenName" json:"tokenName,omitempty"`
	// 冻结总额
	TotalCount int64 `protobuf:"varint,4,opt,name=totalCount" json:"totalCount,omitempty"`
	// 发币人地址
	Initiator string `protobuf:"bytes,5,opt,name=initiator" json:"initiator,omitempty"`
	// 收币人地址
	Beneficiary string `protobuf:"bytes,6,opt,name=beneficiary" json:"beneficiary,omitempty"`
	// 解冻间隔
	Period int64 `protobuf:"varint,7,opt,name=period" json:"period,omitempty"`
	// 解冻方式（百分比；固额） 1 百分比 -> 2 固额
	Means int32 `protobuf:"varint,8,opt,name=means" json:"means,omitempty"`
	// 解冻数量：若为百分比解冻方式该字段值为百分比乘以100，若为固额该字段值为币数量
	Amount int64 `protobuf:"varint,9,opt,name=amount" json:"amount,omitempty"`
	// 已解冻次数
	WithdrawTimes int32 `protobuf:"varint,10,opt,name=withdrawTimes" json:"withdrawTimes,omitempty"`
}

func (m *Unfreeze) Reset()                    { *m = Unfreeze{} }
func (m *Unfreeze) String() string            { return proto.CompactTextString(m) }
func (*Unfreeze) ProtoMessage()               {}
func (*Unfreeze) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Unfreeze) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

func (m *Unfreeze) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Unfreeze) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *Unfreeze) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *Unfreeze) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *Unfreeze) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

func (m *Unfreeze) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *Unfreeze) GetMeans() int32 {
	if m != nil {
		return m.Means
	}
	return 0
}

func (m *Unfreeze) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Unfreeze) GetWithdrawTimes() int32 {
	if m != nil {
		return m.WithdrawTimes
	}
	return 0
}

// message for execs.unfreeze
type UnfreezeAction struct {
	// Types that are valid to be assigned to Value:
	//	*UnfreezeAction_Create
	//	*UnfreezeAction_Withdraw
	//	*UnfreezeAction_Terminate
	Value isUnfreezeAction_Value `protobuf_oneof:"value"`
	Ty    int32                  `protobuf:"varint,10,opt,name=ty" json:"ty,omitempty"`
}

func (m *UnfreezeAction) Reset()                    { *m = UnfreezeAction{} }
func (m *UnfreezeAction) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeAction) ProtoMessage()               {}
func (*UnfreezeAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isUnfreezeAction_Value interface {
	isUnfreezeAction_Value()
}

type UnfreezeAction_Create struct {
	Create *UnfreezeCreate `protobuf:"bytes,1,opt,name=create,oneof"`
}
type UnfreezeAction_Withdraw struct {
	Withdraw *UnfreezeWithdraw `protobuf:"bytes,2,opt,name=withdraw,oneof"`
}
type UnfreezeAction_Terminate struct {
	Terminate *UnfreezeTerminate `protobuf:"bytes,3,opt,name=terminate,oneof"`
}

func (*UnfreezeAction_Create) isUnfreezeAction_Value()    {}
func (*UnfreezeAction_Withdraw) isUnfreezeAction_Value()  {}
func (*UnfreezeAction_Terminate) isUnfreezeAction_Value() {}

func (m *UnfreezeAction) GetValue() isUnfreezeAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *UnfreezeAction) GetCreate() *UnfreezeCreate {
	if x, ok := m.GetValue().(*UnfreezeAction_Create); ok {
		return x.Create
	}
	return nil
}

func (m *UnfreezeAction) GetWithdraw() *UnfreezeWithdraw {
	if x, ok := m.GetValue().(*UnfreezeAction_Withdraw); ok {
		return x.Withdraw
	}
	return nil
}

func (m *UnfreezeAction) GetTerminate() *UnfreezeTerminate {
	if x, ok := m.GetValue().(*UnfreezeAction_Terminate); ok {
		return x.Terminate
	}
	return nil
}

func (m *UnfreezeAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UnfreezeAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UnfreezeAction_OneofMarshaler, _UnfreezeAction_OneofUnmarshaler, _UnfreezeAction_OneofSizer, []interface{}{
		(*UnfreezeAction_Create)(nil),
		(*UnfreezeAction_Withdraw)(nil),
		(*UnfreezeAction_Terminate)(nil),
	}
}

func _UnfreezeAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UnfreezeAction)
	// value
	switch x := m.Value.(type) {
	case *UnfreezeAction_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *UnfreezeAction_Withdraw:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Withdraw); err != nil {
			return err
		}
	case *UnfreezeAction_Terminate:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Terminate); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UnfreezeAction.Value has unexpected type %T", x)
	}
	return nil
}

func _UnfreezeAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UnfreezeAction)
	switch tag {
	case 1: // value.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UnfreezeCreate)
		err := b.DecodeMessage(msg)
		m.Value = &UnfreezeAction_Create{msg}
		return true, err
	case 2: // value.withdraw
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UnfreezeWithdraw)
		err := b.DecodeMessage(msg)
		m.Value = &UnfreezeAction_Withdraw{msg}
		return true, err
	case 3: // value.terminate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UnfreezeTerminate)
		err := b.DecodeMessage(msg)
		m.Value = &UnfreezeAction_Terminate{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UnfreezeAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UnfreezeAction)
	// value
	switch x := m.Value.(type) {
	case *UnfreezeAction_Create:
		s := proto.Size(x.Create)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UnfreezeAction_Withdraw:
		s := proto.Size(x.Withdraw)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UnfreezeAction_Terminate:
		s := proto.Size(x.Terminate)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type UnfreezeCreate struct {
	StartTime   int64  `protobuf:"varint,1,opt,name=startTime" json:"startTime,omitempty"`
	TokenName   string `protobuf:"bytes,2,opt,name=tokenName" json:"tokenName,omitempty"`
	TotalCount  int64  `protobuf:"varint,3,opt,name=totalCount" json:"totalCount,omitempty"`
	Initiator   string `protobuf:"bytes,4,opt,name=initiator" json:"initiator,omitempty"`
	Beneficiary string `protobuf:"bytes,5,opt,name=beneficiary" json:"beneficiary,omitempty"`
	Period      int64  `protobuf:"varint,6,opt,name=period" json:"period,omitempty"`
	Means       int32  `protobuf:"varint,7,opt,name=means" json:"means,omitempty"`
	Amount      int64  `protobuf:"varint,8,opt,name=amount" json:"amount,omitempty"`
}

func (m *UnfreezeCreate) Reset()                    { *m = UnfreezeCreate{} }
func (m *UnfreezeCreate) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeCreate) ProtoMessage()               {}
func (*UnfreezeCreate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UnfreezeCreate) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *UnfreezeCreate) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *UnfreezeCreate) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *UnfreezeCreate) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *UnfreezeCreate) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

func (m *UnfreezeCreate) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *UnfreezeCreate) GetMeans() int32 {
	if m != nil {
		return m.Means
	}
	return 0
}

func (m *UnfreezeCreate) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type UnfreezeWithdraw struct {
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
}

func (m *UnfreezeWithdraw) Reset()                    { *m = UnfreezeWithdraw{} }
func (m *UnfreezeWithdraw) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeWithdraw) ProtoMessage()               {}
func (*UnfreezeWithdraw) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UnfreezeWithdraw) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

type UnfreezeTerminate struct {
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
}

func (m *UnfreezeTerminate) Reset()                    { *m = UnfreezeTerminate{} }
func (m *UnfreezeTerminate) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeTerminate) ProtoMessage()               {}
func (*UnfreezeTerminate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UnfreezeTerminate) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

// TODO: 更多receipt类型
type ReceiptUnfreeze struct {
	TokenName   string `protobuf:"bytes,1,opt,name=tokenName" json:"tokenName,omitempty"`
	CreateAddr  string `protobuf:"bytes,2,opt,name=createAddr" json:"createAddr,omitempty"`
	ReceiveAddr string `protobuf:"bytes,3,opt,name=receiveAddr" json:"receiveAddr,omitempty"`
}

func (m *ReceiptUnfreeze) Reset()                    { *m = ReceiptUnfreeze{} }
func (m *ReceiptUnfreeze) String() string            { return proto.CompactTextString(m) }
func (*ReceiptUnfreeze) ProtoMessage()               {}
func (*ReceiptUnfreeze) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReceiptUnfreeze) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *ReceiptUnfreeze) GetCreateAddr() string {
	if m != nil {
		return m.CreateAddr
	}
	return ""
}

func (m *ReceiptUnfreeze) GetReceiveAddr() string {
	if m != nil {
		return m.ReceiveAddr
	}
	return ""
}

type QueryWithdraw struct {
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
}

func (m *QueryWithdraw) Reset()                    { *m = QueryWithdraw{} }
func (m *QueryWithdraw) String() string            { return proto.CompactTextString(m) }
func (*QueryWithdraw) ProtoMessage()               {}
func (*QueryWithdraw) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *QueryWithdraw) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

func init() {
	proto.RegisterType((*Unfreeze)(nil), "types.Unfreeze")
	proto.RegisterType((*UnfreezeAction)(nil), "types.UnfreezeAction")
	proto.RegisterType((*UnfreezeCreate)(nil), "types.UnfreezeCreate")
	proto.RegisterType((*UnfreezeWithdraw)(nil), "types.UnfreezeWithdraw")
	proto.RegisterType((*UnfreezeTerminate)(nil), "types.UnfreezeTerminate")
	proto.RegisterType((*ReceiptUnfreeze)(nil), "types.ReceiptUnfreeze")
	proto.RegisterType((*QueryWithdraw)(nil), "types.QueryWithdraw")
}

func init() { proto.RegisterFile("unfreeze.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x6d, 0xd2, 0x4d, 0xda, 0x4e, 0xb5, 0x05, 0x2c, 0x3e, 0x7c, 0x40, 0xa8, 0x8a, 0x38, 0xf4,
	0xd4, 0x95, 0xba, 0x42, 0xe2, 0xba, 0x2c, 0x87, 0x72, 0x41, 0x22, 0x5a, 0xc4, 0xd9, 0xdb, 0xce,
	0x0a, 0x8b, 0x8d, 0x1d, 0x9c, 0xc9, 0x56, 0xe1, 0x67, 0xf1, 0x37, 0xf8, 0x37, 0xfc, 0x02, 0x64,
	0x3b, 0x69, 0xd3, 0x52, 0xa5, 0x7b, 0x9c, 0x37, 0xef, 0x4d, 0x3c, 0xef, 0xd9, 0x81, 0x49, 0xa9,
	0xee, 0x0c, 0xe2, 0x2f, 0x9c, 0xe7, 0x46, 0x93, 0x66, 0x11, 0x55, 0x39, 0x16, 0xc9, 0xef, 0x10,
	0x86, 0x5f, 0xeb, 0x0e, 0x7b, 0x03, 0xd0, 0xb0, 0x3e, 0x7d, 0xe4, 0xc1, 0x34, 0x98, 0x8d, 0xd2,
	0x16, 0xc2, 0x5e, 0xc3, 0xa8, 0x20, 0x61, 0xe8, 0x46, 0x66, 0xc8, 0xc3, 0x69, 0x30, 0xeb, 0xa7,
	0x3b, 0xc0, 0x76, 0x49, 0xff, 0x40, 0xf5, 0x59, 0x64, 0xc8, 0xfb, 0x4e, 0xbc, 0x03, 0xec, 0x6c,
	0xd2, 0x24, 0xee, 0xaf, 0x75, 0xa9, 0x88, 0x9f, 0x39, 0x71, 0x0b, 0xb1, 0x6a, 0xa9, 0x24, 0x49,
	0x41, 0xda, 0xf0, 0xc8, 0xab, 0xb7, 0x00, 0x9b, 0xc2, 0xf8, 0x16, 0x15, 0xde, 0xc9, 0x95, 0x14,
	0xa6, 0xe2, 0xb1, 0xeb, 0xb7, 0x21, 0xf6, 0x12, 0xe2, 0x1c, 0x8d, 0xd4, 0x6b, 0x3e, 0x70, 0xb3,
	0xeb, 0x8a, 0x3d, 0x87, 0x28, 0x43, 0xa1, 0x0a, 0x3e, 0x9c, 0x06, 0xb3, 0x28, 0xf5, 0x85, 0x65,
	0x8b, 0xcc, 0x9d, 0x64, 0xe4, 0xd9, 0xbe, 0x62, 0x6f, 0xe1, 0x7c, 0x23, 0xe9, 0xfb, 0xda, 0x88,
	0x8d, 0xdd, 0xa9, 0xe0, 0xe0, 0x54, 0xfb, 0x60, 0xf2, 0x27, 0x80, 0x49, 0x63, 0xda, 0xd5, 0x8a,
	0xa4, 0x56, 0xec, 0x02, 0xe2, 0x95, 0x41, 0x41, 0xe8, 0x6c, 0x1b, 0x2f, 0x5e, 0xcc, 0x9d, 0xbf,
	0xf3, 0x86, 0x76, 0xed, 0x9a, 0xcb, 0x5e, 0x5a, 0xd3, 0xd8, 0x3b, 0x18, 0x36, 0x43, 0x9d, 0x95,
	0xe3, 0xc5, 0xab, 0x03, 0xc9, 0xb7, 0xba, 0xbd, 0xec, 0xa5, 0x5b, 0x2a, 0x7b, 0x0f, 0x23, 0x42,
	0x93, 0x49, 0x65, 0x3f, 0xd5, 0x77, 0x3a, 0x7e, 0xa0, 0xbb, 0x69, 0xfa, 0xcb, 0x5e, 0xba, 0x23,
	0xb3, 0x09, 0x84, 0x54, 0xd5, 0xfb, 0x84, 0x54, 0x7d, 0x18, 0x40, 0xf4, 0x20, 0xee, 0x4b, 0x4c,
	0xfe, 0xb6, 0xb6, 0xf1, 0xc7, 0xdc, 0x0f, 0x3a, 0xe8, 0x0c, 0x3a, 0xec, 0x0e, 0xba, 0xdf, 0x1d,
	0xf4, 0xd9, 0x89, 0xa0, 0xa3, 0xae, 0xa0, 0xe3, 0xe3, 0x41, 0x0f, 0x8e, 0x07, 0x3d, 0x6c, 0x07,
	0x9d, 0x2c, 0xe0, 0xe9, 0xa1, 0xcf, 0xa7, 0xae, 0x7f, 0x72, 0x09, 0xcf, 0xfe, 0xf3, 0xf8, 0xa4,
	0xe8, 0x27, 0x3c, 0x49, 0x71, 0x85, 0x32, 0xa7, 0xed, 0x33, 0xdb, 0xf3, 0x2f, 0x38, 0xe2, 0x9f,
	0xbf, 0x22, 0x57, 0xeb, 0xb5, 0xa9, 0xed, 0x6d, 0x21, 0xd6, 0x21, 0x63, 0x07, 0x3e, 0x78, 0x82,
	0x7f, 0x68, 0x6d, 0x28, 0xb9, 0x80, 0xf3, 0x2f, 0x25, 0x9a, 0xea, 0xb1, 0x8b, 0xdd, 0xc6, 0xee,
	0x97, 0x70, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x73, 0x88, 0x4e, 0x63, 0x24, 0x04, 0x00, 0x00,
}
