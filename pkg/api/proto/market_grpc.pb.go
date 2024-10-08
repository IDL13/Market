// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// MarketClient is the client API for Market service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketClient interface {
	ResetStatistics(ctx context.Context, in *ResetStatisticsRequest, opts ...grpc.CallOption) (*ResetStatisticsResponse, error)
	SaveStatistics(ctx context.Context, in *SaveStatisticsRequest, opts ...grpc.CallOption) (*SaveStatisticsResponse, error)
	ShowStatistics(ctx context.Context, in *ShowStatisticsRequest, opts ...grpc.CallOption) (*ShowStatisticsResponse, error)
}

type marketClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketClient(cc grpc.ClientConnInterface) MarketClient {
	return &marketClient{cc}
}

func (c *marketClient) ResetStatistics(ctx context.Context, in *ResetStatisticsRequest, opts ...grpc.CallOption) (*ResetStatisticsResponse, error) {
	out := new(ResetStatisticsResponse)
	err := c.cc.Invoke(ctx, "/api.Market/ResetStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) SaveStatistics(ctx context.Context, in *SaveStatisticsRequest, opts ...grpc.CallOption) (*SaveStatisticsResponse, error) {
	out := new(SaveStatisticsResponse)
	err := c.cc.Invoke(ctx, "/api.Market/SaveStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) ShowStatistics(ctx context.Context, in *ShowStatisticsRequest, opts ...grpc.CallOption) (*ShowStatisticsResponse, error) {
	out := new(ShowStatisticsResponse)
	err := c.cc.Invoke(ctx, "/api.Market/ShowStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketServer is the server API for Market service.
// All implementations must embed UnimplementedMarketServer
// for forward compatibility
type MarketServer interface {
	ResetStatistics(context.Context, *ResetStatisticsRequest) (*ResetStatisticsResponse, error)
	SaveStatistics(context.Context, *SaveStatisticsRequest) (*SaveStatisticsResponse, error)
	ShowStatistics(context.Context, *ShowStatisticsRequest) (*ShowStatisticsResponse, error)
	mustEmbedUnimplementedMarketServer()
}

// UnimplementedMarketServer must be embedded to have forward compatible implementations.
type UnimplementedMarketServer struct {
}

func (UnimplementedMarketServer) ResetStatistics(context.Context, *ResetStatisticsRequest) (*ResetStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetStatistics not implemented")
}
func (UnimplementedMarketServer) SaveStatistics(context.Context, *SaveStatisticsRequest) (*SaveStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveStatistics not implemented")
}
func (UnimplementedMarketServer) ShowStatistics(context.Context, *ShowStatisticsRequest) (*ShowStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowStatistics not implemented")
}
func (UnimplementedMarketServer) mustEmbedUnimplementedMarketServer() {}

// UnsafeMarketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketServer will
// result in compilation errors.
type UnsafeMarketServer interface {
	mustEmbedUnimplementedMarketServer()
}

func RegisterMarketServer(s grpc.ServiceRegistrar, srv MarketServer) {
	s.RegisterService(&Market_ServiceDesc, srv)
}

func _Market_ResetStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).ResetStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Market/ResetStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).ResetStatistics(ctx, req.(*ResetStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_SaveStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).SaveStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Market/SaveStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).SaveStatistics(ctx, req.(*SaveStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_ShowStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).ShowStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Market/ShowStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).ShowStatistics(ctx, req.(*ShowStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Market_ServiceDesc is the grpc.ServiceDesc for Market service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Market_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Market",
	HandlerType: (*MarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResetStatistics",
			Handler:    _Market_ResetStatistics_Handler,
		},
		{
			MethodName: "SaveStatistics",
			Handler:    _Market_SaveStatistics_Handler,
		},
		{
			MethodName: "ShowStatistics",
			Handler:    _Market_ShowStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/market.proto",
}
