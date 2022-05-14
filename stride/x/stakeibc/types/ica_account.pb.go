// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stakeibc/ica_account.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type ICAAccount struct {
	Address          string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance          int32         `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	DelegatedBalance int32         `protobuf:"varint,3,opt,name=delegatedBalance,proto3" json:"delegatedBalance,omitempty"`
	Delegations      []*Delegation `protobuf:"bytes,4,rep,name=delegations,proto3" json:"delegations,omitempty"`
}

func (m *ICAAccount) Reset()         { *m = ICAAccount{} }
func (m *ICAAccount) String() string { return proto.CompactTextString(m) }
func (*ICAAccount) ProtoMessage()    {}
func (*ICAAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7243c23ee376c2f, []int{0}
}
func (m *ICAAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ICAAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ICAAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ICAAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ICAAccount.Merge(m, src)
}
func (m *ICAAccount) XXX_Size() int {
	return m.Size()
}
func (m *ICAAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ICAAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ICAAccount proto.InternalMessageInfo

func (m *ICAAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ICAAccount) GetBalance() int32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *ICAAccount) GetDelegatedBalance() int32 {
	if m != nil {
		return m.DelegatedBalance
	}
	return 0
}

func (m *ICAAccount) GetDelegations() []*Delegation {
	if m != nil {
		return m.Delegations
	}
	return nil
}

func init() {
	proto.RegisterType((*ICAAccount)(nil), "Stridelabs.stride.stakeibc.ICAAccount")
}

func init() { proto.RegisterFile("stakeibc/ica_account.proto", fileDescriptor_f7243c23ee376c2f) }

var fileDescriptor_f7243c23ee376c2f = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x2e, 0x49, 0xcc,
	0x4e, 0xcd, 0x4c, 0x4a, 0xd6, 0xcf, 0x4c, 0x4e, 0x8c, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x0a, 0x2e, 0x29, 0xca, 0x4c, 0x49, 0xcd, 0x49,
	0x4c, 0x2a, 0xd6, 0x2b, 0x06, 0x33, 0xf5, 0x60, 0xaa, 0xa5, 0x24, 0xe1, 0xfa, 0x52, 0x52, 0x73,
	0x52, 0xd3, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0x20, 0xda, 0xa4, 0x24, 0x93, 0xf3, 0x8b, 0x73, 0xf3,
	0x8b, 0xe3, 0xc1, 0x3c, 0x7d, 0x08, 0x07, 0x22, 0xa5, 0x74, 0x81, 0x91, 0x8b, 0xcb, 0xd3, 0xd9,
	0xd1, 0x11, 0x62, 0x8d, 0x90, 0x11, 0x17, 0x7b, 0x62, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0xa7, 0x93, 0xc4, 0xa5, 0x2d, 0xba, 0x22, 0x50, 0x1d, 0x8e, 0x10, 0x19,
	0x90, 0x1b, 0xf2, 0xd2, 0x83, 0x60, 0x0a, 0x85, 0x24, 0xb8, 0xd8, 0x93, 0x12, 0x73, 0x12, 0xf3,
	0x92, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0x60, 0x5c, 0x21, 0x2d, 0x2e, 0x01, 0xa8,
	0x5b, 0x52, 0x53, 0x9c, 0xa0, 0x4a, 0x98, 0xc1, 0x4a, 0x30, 0xc4, 0x85, 0x3c, 0xb8, 0xb8, 0x11,
	0xee, 0x2e, 0x96, 0x60, 0x51, 0x60, 0xd6, 0xe0, 0x36, 0x52, 0xd3, 0xc3, 0xed, 0x61, 0x3d, 0x17,
	0xb8, 0xf2, 0x20, 0x64, 0xad, 0x4e, 0x1e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8,
	0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7,
	0x10, 0xa5, 0x97, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x31, 0x58,
	0x17, 0x64, 0xb2, 0x3e, 0xc4, 0x64, 0xfd, 0x0a, 0x7d, 0x78, 0x10, 0x96, 0x54, 0x16, 0xa4, 0x16,
	0x27, 0xb1, 0x81, 0xc3, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x79, 0x0c, 0x67, 0x4d, 0x93,
	0x01, 0x00, 0x00,
}

func (m *ICAAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ICAAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ICAAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Delegations) > 0 {
		for iNdEx := len(m.Delegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Delegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIcaAccount(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.DelegatedBalance != 0 {
		i = encodeVarintIcaAccount(dAtA, i, uint64(m.DelegatedBalance))
		i--
		dAtA[i] = 0x18
	}
	if m.Balance != 0 {
		i = encodeVarintIcaAccount(dAtA, i, uint64(m.Balance))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintIcaAccount(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIcaAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovIcaAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ICAAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovIcaAccount(uint64(l))
	}
	if m.Balance != 0 {
		n += 1 + sovIcaAccount(uint64(m.Balance))
	}
	if m.DelegatedBalance != 0 {
		n += 1 + sovIcaAccount(uint64(m.DelegatedBalance))
	}
	if len(m.Delegations) > 0 {
		for _, e := range m.Delegations {
			l = e.Size()
			n += 1 + l + sovIcaAccount(uint64(l))
		}
	}
	return n
}

func sovIcaAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIcaAccount(x uint64) (n int) {
	return sovIcaAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ICAAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIcaAccount
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
			return fmt.Errorf("proto: ICAAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ICAAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaAccount
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
				return ErrInvalidLengthIcaAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			m.Balance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Balance |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatedBalance", wireType)
			}
			m.DelegatedBalance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelegatedBalance |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIcaAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIcaAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegations = append(m.Delegations, &Delegation{})
			if err := m.Delegations[len(m.Delegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIcaAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIcaAccount
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
func skipIcaAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIcaAccount
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
					return 0, ErrIntOverflowIcaAccount
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
					return 0, ErrIntOverflowIcaAccount
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
				return 0, ErrInvalidLengthIcaAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIcaAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIcaAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIcaAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIcaAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIcaAccount = fmt.Errorf("proto: unexpected end of group")
)
