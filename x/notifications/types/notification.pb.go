// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canine_chain/notifications/notification.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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

type Notification struct {
	To              string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	From            string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Time            int64  `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	Contents        string `protobuf:"bytes,4,opt,name=contents,proto3" json:"contents,omitempty"`
	PrivateContents []byte `protobuf:"bytes,5,opt,name=private_contents,json=privateContents,proto3" json:"private_contents,omitempty"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbec842809b00e30, []int{0}
}
func (m *Notification) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return m.Size()
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Notification) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Notification) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Notification) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

func (m *Notification) GetPrivateContents() []byte {
	if m != nil {
		return m.PrivateContents
	}
	return nil
}

type Block struct {
	Address        string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	BlockedAddress string `protobuf:"bytes,2,opt,name=blocked_address,json=blockedAddress,proto3" json:"blocked_address,omitempty"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbec842809b00e30, []int{1}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Block.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return m.Size()
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Block) GetBlockedAddress() string {
	if m != nil {
		return m.BlockedAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Notification)(nil), "canine_chain.notifications.Notification")
	proto.RegisterType((*Block)(nil), "canine_chain.notifications.Block")
}

func init() {
	proto.RegisterFile("canine_chain/notifications/notification.proto", fileDescriptor_cbec842809b00e30)
}

var fileDescriptor_cbec842809b00e30 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xbf, 0x4f, 0x02, 0x31,
	0x14, 0xc7, 0x29, 0x3f, 0xfc, 0xd1, 0x10, 0x30, 0x8d, 0xc3, 0xe5, 0x86, 0x93, 0xb0, 0x88, 0x03,
	0x74, 0x70, 0x71, 0x15, 0x37, 0x63, 0x1c, 0x6e, 0x74, 0x21, 0xbd, 0x52, 0x8e, 0x0a, 0xd7, 0x77,
	0xb9, 0x3e, 0x8c, 0xfe, 0x11, 0x26, 0xfe, 0x59, 0x8e, 0x8c, 0x8e, 0x06, 0xfe, 0x11, 0x73, 0xed,
	0x41, 0xc0, 0xed, 0xfb, 0xfd, 0xe4, 0xd3, 0xe6, 0xe5, 0x3d, 0x3a, 0x94, 0xc2, 0x68, 0xa3, 0x26,
	0x72, 0x2e, 0xb4, 0xe1, 0x06, 0x50, 0xcf, 0xb4, 0x14, 0xa8, 0xc1, 0xd8, 0xa3, 0x36, 0xca, 0x0b,
	0x40, 0x60, 0xe1, 0xa1, 0x3e, 0x3a, 0xd2, 0xc3, 0xcb, 0x14, 0x52, 0x70, 0x1a, 0x2f, 0x93, 0x7f,
	0x11, 0x5e, 0xa5, 0x00, 0xe9, 0x52, 0x71, 0xd7, 0x92, 0xd5, 0x8c, 0xa3, 0xce, 0x94, 0x45, 0x91,
	0xe5, 0x5e, 0xe8, 0x7f, 0x12, 0xda, 0x7e, 0x3e, 0xf8, 0x88, 0x75, 0x68, 0x1d, 0x21, 0x20, 0x3d,
	0x32, 0x38, 0x8f, 0xeb, 0x08, 0x8c, 0xd1, 0xe6, 0xac, 0x80, 0x2c, 0xa8, 0x3b, 0xe2, 0x72, 0xc9,
	0xca, 0x7f, 0x82, 0x46, 0x8f, 0x0c, 0x1a, 0xb1, 0xcb, 0x2c, 0xa4, 0x67, 0x12, 0x0c, 0x2a, 0x83,
	0x36, 0x68, 0x3a, 0x77, 0xdf, 0xd9, 0x0d, 0xbd, 0xc8, 0x0b, 0xfd, 0x26, 0x50, 0x4d, 0xf6, 0x4e,
	0xab, 0x47, 0x06, 0xed, 0xb8, 0x5b, 0xf1, 0x87, 0x0a, 0xf7, 0x1f, 0x69, 0x6b, 0xbc, 0x04, 0xb9,
	0x60, 0x01, 0x3d, 0x15, 0xd3, 0x69, 0xa1, 0xac, 0xad, 0x86, 0xd9, 0x55, 0x76, 0x4d, 0xbb, 0x49,
	0xa9, 0xa8, 0xe9, 0x64, 0x67, 0xf8, 0xe1, 0x3a, 0x15, 0xbe, 0xf7, 0x74, 0x1c, 0x7f, 0x6f, 0x22,
	0xb2, 0xde, 0x44, 0xe4, 0x77, 0x13, 0x91, 0xaf, 0x6d, 0x54, 0x5b, 0x6f, 0xa3, 0xda, 0xcf, 0x36,
	0xaa, 0xbd, 0xdc, 0xa5, 0x1a, 0xe7, 0xab, 0x64, 0x24, 0x21, 0xe3, 0xaf, 0x42, 0x2e, 0xc4, 0xf2,
	0x49, 0x24, 0x96, 0xfb, 0xf5, 0x0e, 0xfd, 0x35, 0xde, 0xff, 0xdd, 0x03, 0x3f, 0x72, 0x65, 0x93,
	0x13, 0xb7, 0xb6, 0xdb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0x11, 0xde, 0x53, 0xba, 0x01,
	0x00, 0x00,
}

func (m *Notification) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Notification) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Notification) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PrivateContents) > 0 {
		i -= len(m.PrivateContents)
		copy(dAtA[i:], m.PrivateContents)
		i = encodeVarintNotification(dAtA, i, uint64(len(m.PrivateContents)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Contents) > 0 {
		i -= len(m.Contents)
		copy(dAtA[i:], m.Contents)
		i = encodeVarintNotification(dAtA, i, uint64(len(m.Contents)))
		i--
		dAtA[i] = 0x22
	}
	if m.Time != 0 {
		i = encodeVarintNotification(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x18
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintNotification(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintNotification(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Block) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Block) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Block) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BlockedAddress) > 0 {
		i -= len(m.BlockedAddress)
		copy(dAtA[i:], m.BlockedAddress)
		i = encodeVarintNotification(dAtA, i, uint64(len(m.BlockedAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintNotification(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNotification(dAtA []byte, offset int, v uint64) int {
	offset -= sovNotification(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Notification) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovNotification(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovNotification(uint64(l))
	}
	if m.Time != 0 {
		n += 1 + sovNotification(uint64(m.Time))
	}
	l = len(m.Contents)
	if l > 0 {
		n += 1 + l + sovNotification(uint64(l))
	}
	l = len(m.PrivateContents)
	if l > 0 {
		n += 1 + l + sovNotification(uint64(l))
	}
	return n
}

func (m *Block) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovNotification(uint64(l))
	}
	l = len(m.BlockedAddress)
	if l > 0 {
		n += 1 + l + sovNotification(uint64(l))
	}
	return n
}

func sovNotification(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNotification(x uint64) (n int) {
	return sovNotification(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Notification) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotification
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
			return fmt.Errorf("proto: Notification: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Notification: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotification
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
				return ErrInvalidLengthNotification
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotification
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotification
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
				return ErrInvalidLengthNotification
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotification
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotification
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contents", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotification
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
				return ErrInvalidLengthNotification
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotification
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contents = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrivateContents", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotification
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNotification
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNotification
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrivateContents = append(m.PrivateContents[:0], dAtA[iNdEx:postIndex]...)
			if m.PrivateContents == nil {
				m.PrivateContents = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNotification(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNotification
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
func (m *Block) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotification
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
			return fmt.Errorf("proto: Block: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Block: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotification
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
				return ErrInvalidLengthNotification
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotification
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockedAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotification
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
				return ErrInvalidLengthNotification
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotification
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockedAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNotification(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNotification
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
func skipNotification(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNotification
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
					return 0, ErrIntOverflowNotification
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
					return 0, ErrIntOverflowNotification
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
				return 0, ErrInvalidLengthNotification
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNotification
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNotification
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNotification        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNotification          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNotification = fmt.Errorf("proto: unexpected end of group")
)
