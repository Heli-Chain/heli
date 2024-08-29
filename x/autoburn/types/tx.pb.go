// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: helichain/autoburn/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgBurnCoins struct {
	Sender string       `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Coins  []types.Coin `protobuf:"bytes,2,rep,name=coins,proto3" json:"coins"`
}

func (m *MsgBurnCoins) Reset()         { *m = MsgBurnCoins{} }
func (m *MsgBurnCoins) String() string { return proto.CompactTextString(m) }
func (*MsgBurnCoins) ProtoMessage()    {}
func (*MsgBurnCoins) Descriptor() ([]byte, []int) {
	return fileDescriptor_041bdc935fcea6b3, []int{0}
}
func (m *MsgBurnCoins) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBurnCoins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBurnCoins.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBurnCoins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurnCoins.Merge(m, src)
}
func (m *MsgBurnCoins) XXX_Size() int {
	return m.Size()
}
func (m *MsgBurnCoins) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurnCoins.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurnCoins proto.InternalMessageInfo

func (m *MsgBurnCoins) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgBurnCoins) GetCoins() []types.Coin {
	if m != nil {
		return m.Coins
	}
	return nil
}

type MsgBurnCoinsResponse struct {
}

func (m *MsgBurnCoinsResponse) Reset()         { *m = MsgBurnCoinsResponse{} }
func (m *MsgBurnCoinsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBurnCoinsResponse) ProtoMessage()    {}
func (*MsgBurnCoinsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_041bdc935fcea6b3, []int{1}
}
func (m *MsgBurnCoinsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBurnCoinsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBurnCoinsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBurnCoinsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurnCoinsResponse.Merge(m, src)
}
func (m *MsgBurnCoinsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBurnCoinsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurnCoinsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurnCoinsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgBurnCoins)(nil), "helichain.autoburn.MsgBurnCoins")
	proto.RegisterType((*MsgBurnCoinsResponse)(nil), "helichain.autoburn.MsgBurnCoinsResponse")
}

func init() { proto.RegisterFile("helichain/autoburn/tx.proto", fileDescriptor_041bdc935fcea6b3) }

var fileDescriptor_041bdc935fcea6b3 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x48, 0xcd, 0xc9,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x2c, 0x2d, 0xc9, 0x4f, 0x2a, 0x2d, 0xca, 0xd3, 0x2f,
	0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x82, 0x4b, 0xea, 0xc1, 0x24, 0xa5, 0xc4,
	0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xf5, 0x73, 0x8b, 0xd3, 0xf5, 0xcb, 0x0c, 0x41, 0x14, 0x44,
	0xb1, 0x94, 0x24, 0x44, 0x22, 0x1e, 0xcc, 0xd3, 0x87, 0x70, 0xa0, 0x52, 0x72, 0x50, 0x3d, 0x49,
	0x89, 0xc5, 0xa9, 0xfa, 0x65, 0x86, 0x49, 0xa9, 0x25, 0x89, 0x86, 0xfa, 0xc9, 0xf9, 0x99, 0x79,
	0x50, 0x79, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0x88, 0x3e, 0x10, 0x0b, 0x22, 0xaa, 0xd4, 0xc2, 0xc8,
	0xc5, 0xe3, 0x5b, 0x9c, 0xee, 0x54, 0x5a, 0x94, 0xe7, 0x9c, 0x9f, 0x99, 0x57, 0x2c, 0x64, 0xc0,
	0xc5, 0x56, 0x9c, 0x9a, 0x97, 0x92, 0x5a, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0x24, 0x71,
	0x69, 0x8b, 0xae, 0x08, 0xd4, 0x22, 0xc7, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0xe2, 0xe0, 0x92, 0xa2,
	0xcc, 0xbc, 0xf4, 0x20, 0xa8, 0x3a, 0x21, 0x53, 0x2e, 0x56, 0x90, 0x35, 0xc5, 0x12, 0x4c, 0x0a,
	0xcc, 0x1a, 0xdc, 0x46, 0x92, 0x7a, 0x50, 0xd5, 0x20, 0x87, 0xe8, 0x41, 0x1d, 0xa2, 0x07, 0x32,
	0xdc, 0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0x88, 0x6a, 0x2b, 0xee, 0xa6, 0xe7, 0x1b, 0xb4,
	0xa0, 0x66, 0x28, 0x89, 0x71, 0x89, 0x20, 0xbb, 0x22, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38,
	0xd5, 0x28, 0x8e, 0x8b, 0xd9, 0xb7, 0x38, 0x5d, 0x28, 0x9c, 0x8b, 0x13, 0xe1, 0x42, 0x05, 0x3d,
	0xcc, 0x10, 0xd3, 0x43, 0xd6, 0x2d, 0xa5, 0x41, 0x48, 0x05, 0xcc, 0x7c, 0x27, 0x93, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b,
	0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x92, 0x42, 0xc4, 0x59, 0x05, 0x52, 0xac, 0x55,
	0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xc3, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x22,
	0xc1, 0x79, 0xd8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	BurnCoins(ctx context.Context, in *MsgBurnCoins, opts ...grpc.CallOption) (*MsgBurnCoinsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) BurnCoins(ctx context.Context, in *MsgBurnCoins, opts ...grpc.CallOption) (*MsgBurnCoinsResponse, error) {
	out := new(MsgBurnCoinsResponse)
	err := c.cc.Invoke(ctx, "/helichain.autoburn.Msg/BurnCoins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	BurnCoins(context.Context, *MsgBurnCoins) (*MsgBurnCoinsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) BurnCoins(ctx context.Context, req *MsgBurnCoins) (*MsgBurnCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BurnCoins not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_BurnCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurnCoins)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BurnCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helichain.autoburn.Msg/BurnCoins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BurnCoins(ctx, req.(*MsgBurnCoins))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helichain.autoburn.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BurnCoins",
			Handler:    _Msg_BurnCoins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helichain/autoburn/tx.proto",
}

func (m *MsgBurnCoins) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBurnCoins) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBurnCoins) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBurnCoinsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBurnCoinsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBurnCoinsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgBurnCoins) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgBurnCoinsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgBurnCoins) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBurnCoins: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBurnCoins: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBurnCoinsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBurnCoinsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBurnCoinsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
