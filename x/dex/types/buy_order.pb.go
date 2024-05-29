// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bridge/dex/buy_order.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type BuyOrder struct {
	Index       string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	AmountDenom string `protobuf:"bytes,2,opt,name=amountDenom,proto3" json:"amountDenom,omitempty"`
	PriceDenom  string `protobuf:"bytes,3,opt,name=priceDenom,proto3" json:"priceDenom,omitempty"`
}

func (m *BuyOrder) Reset()         { *m = BuyOrder{} }
func (m *BuyOrder) String() string { return proto.CompactTextString(m) }
func (*BuyOrder) ProtoMessage()    {}
func (*BuyOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_23350f7decb5404f, []int{0}
}
func (m *BuyOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BuyOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BuyOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BuyOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyOrder.Merge(m, src)
}
func (m *BuyOrder) XXX_Size() int {
	return m.Size()
}
func (m *BuyOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyOrder.DiscardUnknown(m)
}

var xxx_messageInfo_BuyOrder proto.InternalMessageInfo

func (m *BuyOrder) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *BuyOrder) GetAmountDenom() string {
	if m != nil {
		return m.AmountDenom
	}
	return ""
}

func (m *BuyOrder) GetPriceDenom() string {
	if m != nil {
		return m.PriceDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*BuyOrder)(nil), "bridge.dex.BuyOrder")
}

func init() { proto.RegisterFile("bridge/dex/buy_order.proto", fileDescriptor_23350f7decb5404f) }

var fileDescriptor_23350f7decb5404f = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x2a, 0xca, 0x4c,
	0x49, 0x4f, 0xd5, 0x4f, 0x49, 0xad, 0xd0, 0x4f, 0x2a, 0xad, 0x8c, 0xcf, 0x2f, 0x4a, 0x49, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0xc8, 0xe9, 0xa5, 0xa4, 0x56, 0x28, 0x25,
	0x71, 0x71, 0x38, 0x95, 0x56, 0xfa, 0x83, 0x64, 0x85, 0x44, 0xb8, 0x58, 0x33, 0xf3, 0x52, 0x52,
	0x2b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x05, 0x2e, 0xee, 0xc4, 0xdc,
	0xfc, 0xd2, 0xbc, 0x12, 0x97, 0xd4, 0xbc, 0xfc, 0x5c, 0x09, 0x26, 0xb0, 0x1c, 0xb2, 0x90, 0x90,
	0x1c, 0x17, 0x57, 0x41, 0x51, 0x66, 0x72, 0x2a, 0x44, 0x01, 0x33, 0x58, 0x01, 0x92, 0x88, 0x93,
	0xce, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1,
	0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x09, 0x41, 0x5d, 0x59, 0x01,
	0x76, 0x67, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x91, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x90, 0xc5, 0x3a, 0xd0, 0xc2, 0x00, 0x00, 0x00,
}

func (m *BuyOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BuyOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BuyOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PriceDenom) > 0 {
		i -= len(m.PriceDenom)
		copy(dAtA[i:], m.PriceDenom)
		i = encodeVarintBuyOrder(dAtA, i, uint64(len(m.PriceDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AmountDenom) > 0 {
		i -= len(m.AmountDenom)
		copy(dAtA[i:], m.AmountDenom)
		i = encodeVarintBuyOrder(dAtA, i, uint64(len(m.AmountDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintBuyOrder(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBuyOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovBuyOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BuyOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovBuyOrder(uint64(l))
	}
	l = len(m.AmountDenom)
	if l > 0 {
		n += 1 + l + sovBuyOrder(uint64(l))
	}
	l = len(m.PriceDenom)
	if l > 0 {
		n += 1 + l + sovBuyOrder(uint64(l))
	}
	return n
}

func sovBuyOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBuyOrder(x uint64) (n int) {
	return sovBuyOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BuyOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBuyOrder
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BuyOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BuyOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuyOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBuyOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuyOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBuyOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuyOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBuyOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBuyOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBuyOrder
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBuyOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBuyOrder
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBuyOrder
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBuyOrder
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthBuyOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBuyOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBuyOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBuyOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBuyOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBuyOrder = fmt.Errorf("proto: unexpected end of group")
)
