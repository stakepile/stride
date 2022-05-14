// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stakeibc/validator.proto

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

type Validator struct {
	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address        string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Status         string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	CommissionRate int32  `protobuf:"varint,4,opt,name=commissionRate,proto3" json:"commissionRate,omitempty"`
	DelegationAmt  int32  `protobuf:"varint,5,opt,name=delegationAmt,proto3" json:"delegationAmt,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_135ed83653830bac, []int{0}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Validator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Validator) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Validator) GetCommissionRate() int32 {
	if m != nil {
		return m.CommissionRate
	}
	return 0
}

func (m *Validator) GetDelegationAmt() int32 {
	if m != nil {
		return m.DelegationAmt
	}
	return 0
}

func init() {
	proto.RegisterType((*Validator)(nil), "Stridelabs.stride.stakeibc.Validator")
}

func init() { proto.RegisterFile("stakeibc/validator.proto", fileDescriptor_135ed83653830bac) }

var fileDescriptor_135ed83653830bac = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4a, 0x03, 0x31,
	0x14, 0x86, 0x1b, 0x6d, 0x2b, 0x0d, 0xe8, 0x22, 0x88, 0xc4, 0x2e, 0x42, 0x11, 0x91, 0x6e, 0x3a,
	0x01, 0x3d, 0x41, 0xbb, 0x72, 0x3d, 0x82, 0x0b, 0x37, 0x92, 0x99, 0x84, 0x31, 0x38, 0x99, 0x94,
	0x79, 0xaf, 0xa2, 0xb7, 0xf0, 0x30, 0x6e, 0xbc, 0x81, 0xcb, 0xe2, 0xca, 0xa5, 0xcc, 0x5c, 0x44,
	0x26, 0xe9, 0x08, 0xba, 0x7b, 0xef, 0xcf, 0x97, 0xff, 0xc1, 0x47, 0x39, 0xa0, 0x7a, 0x34, 0x36,
	0xcb, 0xe5, 0x93, 0x2a, 0xad, 0x56, 0xe8, 0xeb, 0x64, 0x5d, 0x7b, 0xf4, 0x6c, 0x7a, 0x83, 0xb5,
	0xd5, 0xa6, 0x54, 0x19, 0x24, 0x10, 0xc6, 0xa4, 0x67, 0xa7, 0xa7, 0xb9, 0x07, 0xe7, 0xe1, 0x3e,
	0x90, 0x32, 0x2e, 0xf1, 0xdb, 0xd9, 0x3b, 0xa1, 0x93, 0xdb, 0xbe, 0x8a, 0x31, 0x3a, 0xac, 0x94,
	0x33, 0x9c, 0xcc, 0xc8, 0x7c, 0x92, 0x86, 0x99, 0x5d, 0xd2, 0x03, 0xa5, 0x75, 0x6d, 0x00, 0xf8,
	0x5e, 0x17, 0xaf, 0xf8, 0xe7, 0xdb, 0xe2, 0x78, 0x57, 0xb2, 0x8c, 0x2f, 0xdd, 0xed, 0xaa, 0x48,
	0x7b, 0x90, 0x9d, 0xd0, 0x31, 0xa0, 0xc2, 0x0d, 0xf0, 0xfd, 0xd0, 0xb4, 0xdb, 0xd8, 0x05, 0x3d,
	0xca, 0xbd, 0x73, 0x16, 0xc0, 0xfa, 0x2a, 0x55, 0x68, 0xf8, 0x70, 0x46, 0xe6, 0xa3, 0xf4, 0x5f,
	0xca, 0xce, 0xe9, 0xa1, 0x36, 0xa5, 0x29, 0x14, 0x5a, 0x5f, 0x2d, 0x1d, 0xf2, 0x51, 0xc0, 0xfe,
	0x86, 0xab, 0xeb, 0x8f, 0x46, 0x90, 0x6d, 0x23, 0xc8, 0x77, 0x23, 0xc8, 0x6b, 0x2b, 0x06, 0xdb,
	0x56, 0x0c, 0xbe, 0x5a, 0x31, 0xb8, 0x4b, 0x0a, 0x8b, 0x0f, 0x9b, 0x2c, 0xc9, 0xbd, 0x93, 0xd1,
	0xcb, 0xa2, 0x13, 0x23, 0xa3, 0x18, 0xf9, 0x2c, 0x7f, 0x35, 0xe2, 0xcb, 0xda, 0x40, 0x36, 0x0e,
	0x32, 0xae, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x86, 0xc1, 0xf6, 0x94, 0x5f, 0x01, 0x00, 0x00,
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DelegationAmt != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.DelegationAmt))
		i--
		dAtA[i] = 0x28
	}
	if m.CommissionRate != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.CommissionRate))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.CommissionRate != 0 {
		n += 1 + sovValidator(uint64(m.CommissionRate))
	}
	if m.DelegationAmt != 0 {
		n += 1 + sovValidator(uint64(m.DelegationAmt))
	}
	return n
}

func sovValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidator(x uint64) (n int) {
	return sovValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionRate", wireType)
			}
			m.CommissionRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommissionRate |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationAmt", wireType)
			}
			m.DelegationAmt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelegationAmt |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func skipValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
				return 0, ErrInvalidLengthValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidator = fmt.Errorf("proto: unexpected end of group")
)
