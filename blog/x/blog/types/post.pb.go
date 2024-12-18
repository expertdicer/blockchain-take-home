// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blog/blog/post.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Post struct {
	Title         string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body          string     `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Creator       string     `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	Id            uint64     `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt     *time.Time `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at,omitempty"`
	LastUpdatedAt *time.Time `protobuf:"bytes,6,opt,name=last_updated_at,json=lastUpdatedAt,proto3,stdtime" json:"last_updated_at,omitempty"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f060607f92e3b72, []int{0}
}
func (m *Post) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Post.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return m.Size()
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Post) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Post) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Post) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Post) GetLastUpdatedAt() *time.Time {
	if m != nil {
		return m.LastUpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Post)(nil), "blog.blog.Post")
}

func init() { proto.RegisterFile("blog/blog/post.proto", fileDescriptor_8f060607f92e3b72) }

var fileDescriptor_8f060607f92e3b72 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x50, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0xcd, 0xc4, 0xdc, 0x69, 0x56, 0x54, 0x5c, 0x52, 0x2c, 0x29, 0x36, 0xc1, 0x2a, 0x20, 0x24,
	0xa0, 0x5f, 0x70, 0x67, 0x6b, 0x21, 0x41, 0x1b, 0x9b, 0x23, 0x31, 0x6b, 0x08, 0xe4, 0x98, 0x90,
	0xcc, 0x81, 0xf7, 0x17, 0xf7, 0x59, 0x57, 0x5e, 0x69, 0xa5, 0x92, 0xd4, 0xfe, 0x83, 0x64, 0xd7,
	0x7c, 0xc0, 0x35, 0xc3, 0x9b, 0x37, 0xf3, 0x1e, 0x8f, 0xc7, 0xbc, 0xbc, 0xc6, 0x32, 0xd1, 0xa3,
	0xc1, 0x8e, 0xe2, 0xa6, 0x45, 0x42, 0xee, 0x8e, 0x44, 0x3c, 0x0e, 0x3f, 0x28, 0x11, 0xcb, 0x5a,
	0x25, 0xfa, 0x90, 0x6f, 0xde, 0x13, 0xaa, 0xd6, 0xaa, 0xa3, 0x6c, 0xdd, 0x98, 0x5f, 0xdf, 0x2b,
	0xb1, 0x44, 0x0d, 0x93, 0x11, 0x19, 0xf6, 0xe6, 0x17, 0x98, 0xf3, 0x84, 0x1d, 0x71, 0x8f, 0xcd,
	0xa8, 0xa2, 0x5a, 0x09, 0x08, 0x21, 0x72, 0x53, 0xb3, 0x70, 0xce, 0x9c, 0x1c, 0x8b, 0xad, 0xb0,
	0x35, 0xa9, 0x31, 0x17, 0xec, 0xf4, 0xad, 0x55, 0x19, 0x61, 0x2b, 0x4e, 0x34, 0x3d, 0xad, 0xfc,
	0x92, 0xd9, 0x55, 0x21, 0x9c, 0x10, 0x22, 0x27, 0xb5, 0xab, 0x82, 0x3f, 0x30, 0xa6, 0x4f, 0xaa,
	0x58, 0x65, 0x24, 0x66, 0x21, 0x44, 0xe7, 0x77, 0x7e, 0x6c, 0x82, 0xc6, 0x53, 0xd0, 0xf8, 0x79,
	0x0a, 0xba, 0x3c, 0xdb, 0x7f, 0x05, 0xb0, 0xfb, 0x0e, 0x20, 0x75, 0xff, 0x75, 0x0b, 0xe2, 0x8f,
	0xec, 0xaa, 0xce, 0x3a, 0x5a, 0x6d, 0x9a, 0x62, 0x72, 0x9a, 0x1f, 0xe1, 0x74, 0x31, 0x8a, 0x5f,
	0x8c, 0x76, 0x41, 0xcb, 0xdb, 0x7d, 0x2f, 0xe1, 0xd0, 0x4b, 0xf8, 0xe9, 0x25, 0xec, 0x06, 0x69,
	0x1d, 0x06, 0x69, 0x7d, 0x0e, 0xd2, 0x7a, 0xbd, 0xd6, 0xe5, 0x7e, 0x98, 0x8e, 0x69, 0xdb, 0xa8,
	0x2e, 0x9f, 0x6b, 0xe7, 0xfb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xd7, 0x07, 0x63, 0x7d,
	0x01, 0x00, 0x00,
}

func (m *Post) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Post) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Post) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastUpdatedAt != nil {
		n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.LastUpdatedAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.LastUpdatedAt):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintPost(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x32
	}
	if m.CreatedAt != nil {
		n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.CreatedAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.CreatedAt):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintPost(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x2a
	}
	if m.Id != 0 {
		i = encodeVarintPost(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPost(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintPost(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintPost(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPost(dAtA []byte, offset int, v uint64) int {
	offset -= sovPost(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Post) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovPost(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovPost(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPost(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovPost(uint64(m.Id))
	}
	if m.CreatedAt != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.CreatedAt)
		n += 1 + l + sovPost(uint64(l))
	}
	if m.LastUpdatedAt != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.LastUpdatedAt)
		n += 1 + l + sovPost(uint64(l))
	}
	return n
}

func sovPost(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPost(x uint64) (n int) {
	return sovPost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Post) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPost
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
			return fmt.Errorf("proto: Post: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Post: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
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
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
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
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
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
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
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
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedAt == nil {
				m.CreatedAt = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
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
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastUpdatedAt == nil {
				m.LastUpdatedAt = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.LastUpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPost
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
func skipPost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPost
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
					return 0, ErrIntOverflowPost
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
					return 0, ErrIntOverflowPost
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
				return 0, ErrInvalidLengthPost
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPost
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPost
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPost        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPost          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPost = fmt.Errorf("proto: unexpected end of group")
)
