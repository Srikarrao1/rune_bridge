// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: bridge/dex/tx.proto

package dex

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_UpdateParams_FullMethodName   = "/bridge.dex.Msg/UpdateParams"
	Msg_SendCreatePair_FullMethodName = "/bridge.dex.Msg/SendCreatePair"
	Msg_SendSell_FullMethodName       = "/bridge.dex.Msg/SendSell"
	Msg_SendBuy_FullMethodName        = "/bridge.dex.Msg/SendBuy"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	SendCreatePair(ctx context.Context, in *MsgSendCreatePair, opts ...grpc.CallOption) (*MsgSendCreatePairResponse, error)
	SendSell(ctx context.Context, in *MsgSendSell, opts ...grpc.CallOption) (*MsgSendSellResponse, error)
	SendBuy(ctx context.Context, in *MsgSendBuy, opts ...grpc.CallOption) (*MsgSendBuyResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendCreatePair(ctx context.Context, in *MsgSendCreatePair, opts ...grpc.CallOption) (*MsgSendCreatePairResponse, error) {
	out := new(MsgSendCreatePairResponse)
	err := c.cc.Invoke(ctx, Msg_SendCreatePair_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendSell(ctx context.Context, in *MsgSendSell, opts ...grpc.CallOption) (*MsgSendSellResponse, error) {
	out := new(MsgSendSellResponse)
	err := c.cc.Invoke(ctx, Msg_SendSell_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendBuy(ctx context.Context, in *MsgSendBuy, opts ...grpc.CallOption) (*MsgSendBuyResponse, error) {
	out := new(MsgSendBuyResponse)
	err := c.cc.Invoke(ctx, Msg_SendBuy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	SendCreatePair(context.Context, *MsgSendCreatePair) (*MsgSendCreatePairResponse, error)
	SendSell(context.Context, *MsgSendSell) (*MsgSendSellResponse, error)
	SendBuy(context.Context, *MsgSendBuy) (*MsgSendBuyResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) SendCreatePair(context.Context, *MsgSendCreatePair) (*MsgSendCreatePairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCreatePair not implemented")
}
func (UnimplementedMsgServer) SendSell(context.Context, *MsgSendSell) (*MsgSendSellResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSell not implemented")
}
func (UnimplementedMsgServer) SendBuy(context.Context, *MsgSendBuy) (*MsgSendBuyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBuy not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendCreatePair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendCreatePair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendCreatePair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SendCreatePair_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendCreatePair(ctx, req.(*MsgSendCreatePair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendSell)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SendSell_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendSell(ctx, req.(*MsgSendSell))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendBuy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendBuy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendBuy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SendBuy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendBuy(ctx, req.(*MsgSendBuy))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bridge.dex.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "SendCreatePair",
			Handler:    _Msg_SendCreatePair_Handler,
		},
		{
			MethodName: "SendSell",
			Handler:    _Msg_SendSell_Handler,
		},
		{
			MethodName: "SendBuy",
			Handler:    _Msg_SendBuy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bridge/dex/tx.proto",
}
