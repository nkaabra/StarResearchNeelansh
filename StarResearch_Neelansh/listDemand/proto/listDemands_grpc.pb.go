// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: listDemands.proto

package __

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
	DemandListingService_ListDemands_FullMethodName = "/listDemand.DemandListingService/ListDemands"
)

// DemandListingServiceClient is the client API for DemandListingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemandListingServiceClient interface {
	// ListDemands returns a list of current customer demands.
	ListDemands(ctx context.Context, in *ListDemandsRequest, opts ...grpc.CallOption) (*ListDemandsResponse, error)
}

type demandListingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDemandListingServiceClient(cc grpc.ClientConnInterface) DemandListingServiceClient {
	return &demandListingServiceClient{cc}
}

func (c *demandListingServiceClient) ListDemands(ctx context.Context, in *ListDemandsRequest, opts ...grpc.CallOption) (*ListDemandsResponse, error) {
	out := new(ListDemandsResponse)
	err := c.cc.Invoke(ctx, DemandListingService_ListDemands_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemandListingServiceServer is the server API for DemandListingService service.
// All implementations must embed UnimplementedDemandListingServiceServer
// for forward compatibility
type DemandListingServiceServer interface {
	// ListDemands returns a list of current customer demands.
	ListDemands(context.Context, *ListDemandsRequest) (*ListDemandsResponse, error)
	mustEmbedUnimplementedDemandListingServiceServer()
}

// UnimplementedDemandListingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDemandListingServiceServer struct {
}

func (UnimplementedDemandListingServiceServer) ListDemands(context.Context, *ListDemandsRequest) (*ListDemandsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDemands not implemented")
}
func (UnimplementedDemandListingServiceServer) mustEmbedUnimplementedDemandListingServiceServer() {}

// UnsafeDemandListingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemandListingServiceServer will
// result in compilation errors.
type UnsafeDemandListingServiceServer interface {
	mustEmbedUnimplementedDemandListingServiceServer()
}

func RegisterDemandListingServiceServer(s grpc.ServiceRegistrar, srv DemandListingServiceServer) {
	s.RegisterService(&DemandListingService_ServiceDesc, srv)
}

func _DemandListingService_ListDemands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDemandsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemandListingServiceServer).ListDemands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DemandListingService_ListDemands_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemandListingServiceServer).ListDemands(ctx, req.(*ListDemandsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DemandListingService_ServiceDesc is the grpc.ServiceDesc for DemandListingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DemandListingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "listDemand.DemandListingService",
	HandlerType: (*DemandListingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDemands",
			Handler:    _DemandListingService_ListDemands_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "listDemands.proto",
}
