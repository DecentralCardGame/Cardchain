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
	Query_Params_FullMethodName                = "/cardchain.cardchain.Query/Params"
	Query_ProductDetails_FullMethodName        = "/cardchain.cardchain.Query/ProductDetails"
	Query_ProductDetailsAll_FullMethodName     = "/cardchain.cardchain.Query/ProductDetailsAll"
	Query_Card_FullMethodName                  = "/cardchain.cardchain.Query/Card"
	Query_User_FullMethodName                  = "/cardchain.cardchain.Query/User"
	Query_Cards_FullMethodName                 = "/cardchain.cardchain.Query/Cards"
	Query_Match_FullMethodName                 = "/cardchain.cardchain.Query/Match"
	Query_Set_FullMethodName                   = "/cardchain.cardchain.Query/Set"
	Query_SellOffer_FullMethodName             = "/cardchain.cardchain.Query/SellOffer"
	Query_Council_FullMethodName               = "/cardchain.cardchain.Query/Council"
	Query_Server_FullMethodName                = "/cardchain.cardchain.Query/Server"
	Query_Encounter_FullMethodName             = "/cardchain.cardchain.Query/Encounter"
	Query_Encounters_FullMethodName            = "/cardchain.cardchain.Query/Encounters"
	Query_EncounterWithImage_FullMethodName    = "/cardchain.cardchain.Query/EncounterWithImage"
	Query_EncountersWithImage_FullMethodName   = "/cardchain.cardchain.Query/EncountersWithImage"
	Query_CardchainInfo_FullMethodName         = "/cardchain.cardchain.Query/CardchainInfo"
	Query_SetRarityDistribution_FullMethodName = "/cardchain.cardchain.Query/SetRarityDistribution"
	Query_AccountFromZealy_FullMethodName      = "/cardchain.cardchain.Query/AccountFromZealy"
	Query_VotingResults_FullMethodName         = "/cardchain.cardchain.Query/VotingResults"
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
	// Queries a list of Cards items.
	Cards(ctx context.Context, in *QueryCardsRequest, opts ...grpc.CallOption) (*QueryCardsResponse, error)
	// Queries a list of Match items.
	Match(ctx context.Context, in *QueryMatchRequest, opts ...grpc.CallOption) (*QueryMatchResponse, error)
	// Queries a list of Set items.
	Set(ctx context.Context, in *QuerySetRequest, opts ...grpc.CallOption) (*QuerySetResponse, error)
	// Queries a list of SellOffer items.
	SellOffer(ctx context.Context, in *QuerySellOfferRequest, opts ...grpc.CallOption) (*QuerySellOfferResponse, error)
	// Queries a list of Council items.
	Council(ctx context.Context, in *QueryCouncilRequest, opts ...grpc.CallOption) (*QueryCouncilResponse, error)
	// Queries a list of Server items.
	Server(ctx context.Context, in *QueryServerRequest, opts ...grpc.CallOption) (*QueryServerResponse, error)
	// Queries a list of Encounter items.
	Encounter(ctx context.Context, in *QueryEncounterRequest, opts ...grpc.CallOption) (*QueryEncounterResponse, error)
	// Queries a list of Encounters items.
	Encounters(ctx context.Context, in *QueryEncountersRequest, opts ...grpc.CallOption) (*QueryEncountersResponse, error)
	// Queries a list of EncounterWithImage items.
	EncounterWithImage(ctx context.Context, in *QueryEncounterWithImageRequest, opts ...grpc.CallOption) (*QueryEncounterWithImageResponse, error)
	// Queries a list of EncountersWithImage items.
	EncountersWithImage(ctx context.Context, in *QueryEncountersWithImageRequest, opts ...grpc.CallOption) (*QueryEncountersWithImageResponse, error)
	// Queries a list of CardchainInfo items.
	CardchainInfo(ctx context.Context, in *QueryCardchainInfoRequest, opts ...grpc.CallOption) (*QueryCardchainInfoResponse, error)
	// Queries a list of SetRarityDistribution items.
	SetRarityDistribution(ctx context.Context, in *QuerySetRarityDistributionRequest, opts ...grpc.CallOption) (*QuerySetRarityDistributionResponse, error)
	// Queries a list of AccountFromZealy items.
	AccountFromZealy(ctx context.Context, in *QueryAccountFromZealyRequest, opts ...grpc.CallOption) (*QueryAccountFromZealyResponse, error)
	// Queries a list of VotingResults items.
	VotingResults(ctx context.Context, in *QueryVotingResultsRequest, opts ...grpc.CallOption) (*QueryVotingResultsResponse, error)
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

func (c *queryClient) Cards(ctx context.Context, in *QueryCardsRequest, opts ...grpc.CallOption) (*QueryCardsResponse, error) {
	out := new(QueryCardsResponse)
	err := c.cc.Invoke(ctx, Query_Cards_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Match(ctx context.Context, in *QueryMatchRequest, opts ...grpc.CallOption) (*QueryMatchResponse, error) {
	out := new(QueryMatchResponse)
	err := c.cc.Invoke(ctx, Query_Match_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Set(ctx context.Context, in *QuerySetRequest, opts ...grpc.CallOption) (*QuerySetResponse, error) {
	out := new(QuerySetResponse)
	err := c.cc.Invoke(ctx, Query_Set_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SellOffer(ctx context.Context, in *QuerySellOfferRequest, opts ...grpc.CallOption) (*QuerySellOfferResponse, error) {
	out := new(QuerySellOfferResponse)
	err := c.cc.Invoke(ctx, Query_SellOffer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Council(ctx context.Context, in *QueryCouncilRequest, opts ...grpc.CallOption) (*QueryCouncilResponse, error) {
	out := new(QueryCouncilResponse)
	err := c.cc.Invoke(ctx, Query_Council_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Server(ctx context.Context, in *QueryServerRequest, opts ...grpc.CallOption) (*QueryServerResponse, error) {
	out := new(QueryServerResponse)
	err := c.cc.Invoke(ctx, Query_Server_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Encounter(ctx context.Context, in *QueryEncounterRequest, opts ...grpc.CallOption) (*QueryEncounterResponse, error) {
	out := new(QueryEncounterResponse)
	err := c.cc.Invoke(ctx, Query_Encounter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Encounters(ctx context.Context, in *QueryEncountersRequest, opts ...grpc.CallOption) (*QueryEncountersResponse, error) {
	out := new(QueryEncountersResponse)
	err := c.cc.Invoke(ctx, Query_Encounters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EncounterWithImage(ctx context.Context, in *QueryEncounterWithImageRequest, opts ...grpc.CallOption) (*QueryEncounterWithImageResponse, error) {
	out := new(QueryEncounterWithImageResponse)
	err := c.cc.Invoke(ctx, Query_EncounterWithImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EncountersWithImage(ctx context.Context, in *QueryEncountersWithImageRequest, opts ...grpc.CallOption) (*QueryEncountersWithImageResponse, error) {
	out := new(QueryEncountersWithImageResponse)
	err := c.cc.Invoke(ctx, Query_EncountersWithImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CardchainInfo(ctx context.Context, in *QueryCardchainInfoRequest, opts ...grpc.CallOption) (*QueryCardchainInfoResponse, error) {
	out := new(QueryCardchainInfoResponse)
	err := c.cc.Invoke(ctx, Query_CardchainInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SetRarityDistribution(ctx context.Context, in *QuerySetRarityDistributionRequest, opts ...grpc.CallOption) (*QuerySetRarityDistributionResponse, error) {
	out := new(QuerySetRarityDistributionResponse)
	err := c.cc.Invoke(ctx, Query_SetRarityDistribution_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AccountFromZealy(ctx context.Context, in *QueryAccountFromZealyRequest, opts ...grpc.CallOption) (*QueryAccountFromZealyResponse, error) {
	out := new(QueryAccountFromZealyResponse)
	err := c.cc.Invoke(ctx, Query_AccountFromZealy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VotingResults(ctx context.Context, in *QueryVotingResultsRequest, opts ...grpc.CallOption) (*QueryVotingResultsResponse, error) {
	out := new(QueryVotingResultsResponse)
	err := c.cc.Invoke(ctx, Query_VotingResults_FullMethodName, in, out, opts...)
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
	// Queries a list of Cards items.
	Cards(context.Context, *QueryCardsRequest) (*QueryCardsResponse, error)
	// Queries a list of Match items.
	Match(context.Context, *QueryMatchRequest) (*QueryMatchResponse, error)
	// Queries a list of Set items.
	Set(context.Context, *QuerySetRequest) (*QuerySetResponse, error)
	// Queries a list of SellOffer items.
	SellOffer(context.Context, *QuerySellOfferRequest) (*QuerySellOfferResponse, error)
	// Queries a list of Council items.
	Council(context.Context, *QueryCouncilRequest) (*QueryCouncilResponse, error)
	// Queries a list of Server items.
	Server(context.Context, *QueryServerRequest) (*QueryServerResponse, error)
	// Queries a list of Encounter items.
	Encounter(context.Context, *QueryEncounterRequest) (*QueryEncounterResponse, error)
	// Queries a list of Encounters items.
	Encounters(context.Context, *QueryEncountersRequest) (*QueryEncountersResponse, error)
	// Queries a list of EncounterWithImage items.
	EncounterWithImage(context.Context, *QueryEncounterWithImageRequest) (*QueryEncounterWithImageResponse, error)
	// Queries a list of EncountersWithImage items.
	EncountersWithImage(context.Context, *QueryEncountersWithImageRequest) (*QueryEncountersWithImageResponse, error)
	// Queries a list of CardchainInfo items.
	CardchainInfo(context.Context, *QueryCardchainInfoRequest) (*QueryCardchainInfoResponse, error)
	// Queries a list of SetRarityDistribution items.
	SetRarityDistribution(context.Context, *QuerySetRarityDistributionRequest) (*QuerySetRarityDistributionResponse, error)
	// Queries a list of AccountFromZealy items.
	AccountFromZealy(context.Context, *QueryAccountFromZealyRequest) (*QueryAccountFromZealyResponse, error)
	// Queries a list of VotingResults items.
	VotingResults(context.Context, *QueryVotingResultsRequest) (*QueryVotingResultsResponse, error)
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
func (UnimplementedQueryServer) Cards(context.Context, *QueryCardsRequest) (*QueryCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cards not implemented")
}
func (UnimplementedQueryServer) Match(context.Context, *QueryMatchRequest) (*QueryMatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Match not implemented")
}
func (UnimplementedQueryServer) Set(context.Context, *QuerySetRequest) (*QuerySetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedQueryServer) SellOffer(context.Context, *QuerySellOfferRequest) (*QuerySellOfferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellOffer not implemented")
}
func (UnimplementedQueryServer) Council(context.Context, *QueryCouncilRequest) (*QueryCouncilResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Council not implemented")
}
func (UnimplementedQueryServer) Server(context.Context, *QueryServerRequest) (*QueryServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Server not implemented")
}
func (UnimplementedQueryServer) Encounter(context.Context, *QueryEncounterRequest) (*QueryEncounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encounter not implemented")
}
func (UnimplementedQueryServer) Encounters(context.Context, *QueryEncountersRequest) (*QueryEncountersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encounters not implemented")
}
func (UnimplementedQueryServer) EncounterWithImage(context.Context, *QueryEncounterWithImageRequest) (*QueryEncounterWithImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncounterWithImage not implemented")
}
func (UnimplementedQueryServer) EncountersWithImage(context.Context, *QueryEncountersWithImageRequest) (*QueryEncountersWithImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncountersWithImage not implemented")
}
func (UnimplementedQueryServer) CardchainInfo(context.Context, *QueryCardchainInfoRequest) (*QueryCardchainInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CardchainInfo not implemented")
}
func (UnimplementedQueryServer) SetRarityDistribution(context.Context, *QuerySetRarityDistributionRequest) (*QuerySetRarityDistributionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRarityDistribution not implemented")
}
func (UnimplementedQueryServer) AccountFromZealy(context.Context, *QueryAccountFromZealyRequest) (*QueryAccountFromZealyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountFromZealy not implemented")
}
func (UnimplementedQueryServer) VotingResults(context.Context, *QueryVotingResultsRequest) (*QueryVotingResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VotingResults not implemented")
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

func _Query_Cards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Cards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Cards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Cards(ctx, req.(*QueryCardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Match_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Match(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Match_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Match(ctx, req.(*QueryMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Set(ctx, req.(*QuerySetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SellOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySellOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SellOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_SellOffer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SellOffer(ctx, req.(*QuerySellOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Council_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCouncilRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Council(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Council_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Council(ctx, req.(*QueryCouncilRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Server_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Server(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Server_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Server(ctx, req.(*QueryServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Encounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEncounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Encounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Encounter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Encounter(ctx, req.(*QueryEncounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Encounters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEncountersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Encounters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Encounters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Encounters(ctx, req.(*QueryEncountersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EncounterWithImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEncounterWithImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EncounterWithImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_EncounterWithImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EncounterWithImage(ctx, req.(*QueryEncounterWithImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EncountersWithImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEncountersWithImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EncountersWithImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_EncountersWithImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EncountersWithImage(ctx, req.(*QueryEncountersWithImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CardchainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCardchainInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CardchainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_CardchainInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CardchainInfo(ctx, req.(*QueryCardchainInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SetRarityDistribution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySetRarityDistributionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SetRarityDistribution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_SetRarityDistribution_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SetRarityDistribution(ctx, req.(*QuerySetRarityDistributionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AccountFromZealy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountFromZealyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AccountFromZealy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AccountFromZealy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AccountFromZealy(ctx, req.(*QueryAccountFromZealyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VotingResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVotingResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VotingResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_VotingResults_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VotingResults(ctx, req.(*QueryVotingResultsRequest))
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
		{
			MethodName: "Cards",
			Handler:    _Query_Cards_Handler,
		},
		{
			MethodName: "Match",
			Handler:    _Query_Match_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Query_Set_Handler,
		},
		{
			MethodName: "SellOffer",
			Handler:    _Query_SellOffer_Handler,
		},
		{
			MethodName: "Council",
			Handler:    _Query_Council_Handler,
		},
		{
			MethodName: "Server",
			Handler:    _Query_Server_Handler,
		},
		{
			MethodName: "Encounter",
			Handler:    _Query_Encounter_Handler,
		},
		{
			MethodName: "Encounters",
			Handler:    _Query_Encounters_Handler,
		},
		{
			MethodName: "EncounterWithImage",
			Handler:    _Query_EncounterWithImage_Handler,
		},
		{
			MethodName: "EncountersWithImage",
			Handler:    _Query_EncountersWithImage_Handler,
		},
		{
			MethodName: "CardchainInfo",
			Handler:    _Query_CardchainInfo_Handler,
		},
		{
			MethodName: "SetRarityDistribution",
			Handler:    _Query_SetRarityDistribution_Handler,
		},
		{
			MethodName: "AccountFromZealy",
			Handler:    _Query_AccountFromZealy_Handler,
		},
		{
			MethodName: "VotingResults",
			Handler:    _Query_VotingResults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cardchain/cardchain/query.proto",
}
