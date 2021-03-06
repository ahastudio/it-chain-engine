// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PeerService service

type PeerServiceClient interface {
	GetPeer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Peer, error)
}

type peerServiceClient struct {
	cc *grpc.ClientConn
}

func NewPeerServiceClient(cc *grpc.ClientConn) PeerServiceClient {
	return &peerServiceClient{cc}
}

func (c *peerServiceClient) GetPeer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Peer, error) {
	out := new(Peer)
	err := grpc.Invoke(ctx, "/message.PeerService/GetPeer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PeerService service

type PeerServiceServer interface {
	GetPeer(context.Context, *Empty) (*Peer, error)
}

func RegisterPeerServiceServer(s *grpc.Server, srv PeerServiceServer) {
	s.RegisterService(&_PeerService_serviceDesc, srv)
}

func _PeerService_GetPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServiceServer).GetPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.PeerService/GetPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServiceServer).GetPeer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _PeerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.PeerService",
	HandlerType: (*PeerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPeer",
			Handler:    _PeerService_GetPeer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

// Client API for TransactionService service

type TransactionServiceClient interface {
	PostTransaction(ctx context.Context, in *TxData, opts ...grpc.CallOption) (*Transaction, error)
}

type transactionServiceClient struct {
	cc *grpc.ClientConn
}

func NewTransactionServiceClient(cc *grpc.ClientConn) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) PostTransaction(ctx context.Context, in *TxData, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := grpc.Invoke(ctx, "/message.TransactionService/PostTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TransactionService service

type TransactionServiceServer interface {
	PostTransaction(context.Context, *TxData) (*Transaction, error)
}

func RegisterTransactionServiceServer(s *grpc.Server, srv TransactionServiceServer) {
	s.RegisterService(&_TransactionService_serviceDesc, srv)
}

func _TransactionService_PostTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).PostTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.TransactionService/PostTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).PostTransaction(ctx, req.(*TxData))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransactionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostTransaction",
			Handler:    _TransactionService_PostTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e,
	0x4c, 0x4f, 0x95, 0xe2, 0x29, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x85, 0x08, 0x4b, 0xf1, 0x24, 0xe7,
	0xe7, 0xe6, 0xe6, 0xe7, 0x41, 0x78, 0x46, 0x96, 0x5c, 0xdc, 0x01, 0xa9, 0xa9, 0x45, 0xc1, 0x10,
	0x9d, 0x42, 0x5a, 0x5c, 0xec, 0xee, 0xa9, 0x25, 0x20, 0x11, 0x21, 0x3e, 0x3d, 0xa8, 0x7e, 0x3d,
	0xd7, 0xdc, 0x82, 0x92, 0x4a, 0x29, 0x5e, 0x38, 0x1f, 0x24, 0xad, 0xc4, 0x60, 0x14, 0xc0, 0x25,
	0x14, 0x52, 0x94, 0x98, 0x57, 0x9c, 0x98, 0x5c, 0x92, 0x99, 0x9f, 0x07, 0x33, 0xc1, 0x8a, 0x8b,
	0x3f, 0x20, 0xbf, 0xb8, 0x04, 0x49, 0x46, 0x88, 0x1f, 0xae, 0x33, 0xa4, 0xc2, 0x25, 0xb1, 0x24,
	0x51, 0x4a, 0x04, 0x21, 0x80, 0x50, 0xa6, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0x93, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x5f, 0x44, 0x82, 0x21, 0xc9, 0x00, 0x00, 0x00,
}
