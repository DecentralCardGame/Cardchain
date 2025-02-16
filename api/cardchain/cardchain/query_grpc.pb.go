// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cardchain/cardchain/query.proto

package cardchain

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
	Query_Params_FullMethodName            = "/cardchain.cardchain.Query/Params"
	Query_ProductDetails_FullMethodName    = "/cardchain.cardchain.Query/ProductDetails"
	Query_ProductDetailsAll_FullMethodName = "/cardchain.cardchain.Query/ProductDetailsAll"
	Query_Card_FullMethodName              = "/cardchain.cardchain.Query/Card"
	Query_User_FullMethodName              = "/cardchain.cardchain.Query/User"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of ProductDetails items.
	ProductDetails(ctx context.Context, in *QueryGetProductDetailsRequest, opts ...grpc.CallOption) (*QueryGetProductDetailsResponse, error)
	ProductDetailsAll(ctx context.Context, in *QueryAllProductDetailsRequest, opts ...grpc.CallOption) (*QueryAllProductDetailsResponse, error)
	// Queries a list of Card items.
	Card(ctx context.Context, in *QueryCardRequest, opts ...grpc.CallOption) (*QueryCardResponse, error)
	// Queries a list of User items.
	User(ctx context.Context, in *QueryUserRequest, opts ...grpc.CallOption) (*QueryUserResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ProductDetails(ctx context.Context, in *QueryGetProductDetailsRequest, opts ...grpc.CallOption) (*QueryGetProductDetailsResponse, error) {
	out := new(QueryGetProductDetailsResponse)
	err := c.cc.Invoke(ctx, Query_ProductDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ProductDetailsAll(ctx context.Context, in *QueryAllProductDetailsRequest, opts ...grpc.CallOption) (*QueryAllProductDetailsResponse, error) {
	out := new(QueryAllProductDetailsResponse)
	err := c.cc.Invoke(ctx, Query_ProductDetailsAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Card(ctx context.Context, in *QueryCardRequest, opts ...grpc.CallOption) (*QueryCardResponse, error) {
	out := new(QueryCardResponse)
	err := c.cc.Invoke(ctx, Query_Card_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) User(ctx context.Context, in *QueryUserRequest, opts ...grpc.CallOption) (*QueryUserResponse, error) {
	out := new(QueryUserResponse)
	err := c.cc.Invoke(ctx, Query_User_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of ProductDetails items.
	ProductDetails(context.Context, *QueryGetProductDetailsRequest) (*QueryGetProductDetailsResponse, error)
	ProductDetailsAll(context.Context, *QueryAllProductDetailsRequest) (*QueryAllProductDetailsResponse, error)
	// Queries a list of Card items.
	Card(context.Context, *QueryCardRequest) (*QueryCardResponse, error)
	// Queries a list of User items.
	User(context.Context, *QueryUserRequest) (*QueryUserResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) ProductDetails(context.Context, *QueryGetProductDetailsRequest) (*QueryGetProductDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductDetails not implemented")
}
func (UnimplementedQueryServer) ProductDetailsAll(context.Context, *QueryAllProductDetailsRequest) (*QueryAllProductDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductDetailsAll not implemented")
}
func (UnimplementedQueryServer) Card(context.Context, *QueryCardRequest) (*QueryCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Card not implemented")
}
func (UnimplementedQueryServer) User(context.Context, *QueryUserRequest) (*QueryUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method User not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ProductDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetProductDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ProductDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ProductDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ProductDetails(ctx, req.(*QueryGetProductDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ProductDetailsAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllProductDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ProductDetailsAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ProductDetailsAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ProductDetailsAll(ctx, req.(*QueryAllProductDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Card_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Card(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Card_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Card(ctx, req.(*QueryCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_User_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).User(ctx, req.(*QueryUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cardchain.cardchain.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ProductDetails",
			Handler:    _Query_ProductDetails_Handler,
		},
		{
			MethodName: "ProductDetailsAll",
			Handler:    _Query_ProductDetailsAll_Handler,
		},
		{
			MethodName: "Card",
			Handler:    _Query_Card_Handler,
		},
		{
			MethodName: "User",
			Handler:    _Query_User_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cardchain/cardchain/query.proto",
}
