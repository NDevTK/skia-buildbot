// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: service.proto

// Working in progress protobuf and service definition.
//

package pinpointpb

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
	Pinpoint_ScheduleBisection_FullMethodName     = "/pinpoint.v1.Pinpoint/ScheduleBisection"
	Pinpoint_CancelJob_FullMethodName             = "/pinpoint.v1.Pinpoint/CancelJob"
	Pinpoint_QueryBisection_FullMethodName        = "/pinpoint.v1.Pinpoint/QueryBisection"
	Pinpoint_LegacyJobQuery_FullMethodName        = "/pinpoint.v1.Pinpoint/LegacyJobQuery"
	Pinpoint_SchedulePairwise_FullMethodName      = "/pinpoint.v1.Pinpoint/SchedulePairwise"
	Pinpoint_ScheduleCulpritFinder_FullMethodName = "/pinpoint.v1.Pinpoint/ScheduleCulpritFinder"
)

// PinpointClient is the client API for Pinpoint service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PinpointClient interface {
	ScheduleBisection(ctx context.Context, in *ScheduleBisectRequest, opts ...grpc.CallOption) (*BisectExecution, error)
	CancelJob(ctx context.Context, in *CancelJobRequest, opts ...grpc.CallOption) (*CancelJobResponse, error)
	QueryBisection(ctx context.Context, in *QueryBisectRequest, opts ...grpc.CallOption) (*BisectExecution, error)
	LegacyJobQuery(ctx context.Context, in *LegacyJobRequest, opts ...grpc.CallOption) (*LegacyJobResponse, error)
	SchedulePairwise(ctx context.Context, in *SchedulePairwiseRequest, opts ...grpc.CallOption) (*PairwiseExecution, error)
	ScheduleCulpritFinder(ctx context.Context, in *ScheduleCulpritFinderRequest, opts ...grpc.CallOption) (*CulpritFinderExecution, error)
}

type pinpointClient struct {
	cc grpc.ClientConnInterface
}

func NewPinpointClient(cc grpc.ClientConnInterface) PinpointClient {
	return &pinpointClient{cc}
}

func (c *pinpointClient) ScheduleBisection(ctx context.Context, in *ScheduleBisectRequest, opts ...grpc.CallOption) (*BisectExecution, error) {
	out := new(BisectExecution)
	err := c.cc.Invoke(ctx, Pinpoint_ScheduleBisection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pinpointClient) CancelJob(ctx context.Context, in *CancelJobRequest, opts ...grpc.CallOption) (*CancelJobResponse, error) {
	out := new(CancelJobResponse)
	err := c.cc.Invoke(ctx, Pinpoint_CancelJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pinpointClient) QueryBisection(ctx context.Context, in *QueryBisectRequest, opts ...grpc.CallOption) (*BisectExecution, error) {
	out := new(BisectExecution)
	err := c.cc.Invoke(ctx, Pinpoint_QueryBisection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pinpointClient) LegacyJobQuery(ctx context.Context, in *LegacyJobRequest, opts ...grpc.CallOption) (*LegacyJobResponse, error) {
	out := new(LegacyJobResponse)
	err := c.cc.Invoke(ctx, Pinpoint_LegacyJobQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pinpointClient) SchedulePairwise(ctx context.Context, in *SchedulePairwiseRequest, opts ...grpc.CallOption) (*PairwiseExecution, error) {
	out := new(PairwiseExecution)
	err := c.cc.Invoke(ctx, Pinpoint_SchedulePairwise_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pinpointClient) ScheduleCulpritFinder(ctx context.Context, in *ScheduleCulpritFinderRequest, opts ...grpc.CallOption) (*CulpritFinderExecution, error) {
	out := new(CulpritFinderExecution)
	err := c.cc.Invoke(ctx, Pinpoint_ScheduleCulpritFinder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PinpointServer is the server API for Pinpoint service.
// All implementations must embed UnimplementedPinpointServer
// for forward compatibility
type PinpointServer interface {
	ScheduleBisection(context.Context, *ScheduleBisectRequest) (*BisectExecution, error)
	CancelJob(context.Context, *CancelJobRequest) (*CancelJobResponse, error)
	QueryBisection(context.Context, *QueryBisectRequest) (*BisectExecution, error)
	LegacyJobQuery(context.Context, *LegacyJobRequest) (*LegacyJobResponse, error)
	SchedulePairwise(context.Context, *SchedulePairwiseRequest) (*PairwiseExecution, error)
	ScheduleCulpritFinder(context.Context, *ScheduleCulpritFinderRequest) (*CulpritFinderExecution, error)
	mustEmbedUnimplementedPinpointServer()
}

// UnimplementedPinpointServer must be embedded to have forward compatible implementations.
type UnimplementedPinpointServer struct {
}

func (UnimplementedPinpointServer) ScheduleBisection(context.Context, *ScheduleBisectRequest) (*BisectExecution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleBisection not implemented")
}
func (UnimplementedPinpointServer) CancelJob(context.Context, *CancelJobRequest) (*CancelJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelJob not implemented")
}
func (UnimplementedPinpointServer) QueryBisection(context.Context, *QueryBisectRequest) (*BisectExecution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBisection not implemented")
}
func (UnimplementedPinpointServer) LegacyJobQuery(context.Context, *LegacyJobRequest) (*LegacyJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LegacyJobQuery not implemented")
}
func (UnimplementedPinpointServer) SchedulePairwise(context.Context, *SchedulePairwiseRequest) (*PairwiseExecution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchedulePairwise not implemented")
}
func (UnimplementedPinpointServer) ScheduleCulpritFinder(context.Context, *ScheduleCulpritFinderRequest) (*CulpritFinderExecution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleCulpritFinder not implemented")
}
func (UnimplementedPinpointServer) mustEmbedUnimplementedPinpointServer() {}

// UnsafePinpointServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PinpointServer will
// result in compilation errors.
type UnsafePinpointServer interface {
	mustEmbedUnimplementedPinpointServer()
}

func RegisterPinpointServer(s grpc.ServiceRegistrar, srv PinpointServer) {
	s.RegisterService(&Pinpoint_ServiceDesc, srv)
}

func _Pinpoint_ScheduleBisection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleBisectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PinpointServer).ScheduleBisection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pinpoint_ScheduleBisection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PinpointServer).ScheduleBisection(ctx, req.(*ScheduleBisectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pinpoint_CancelJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PinpointServer).CancelJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pinpoint_CancelJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PinpointServer).CancelJob(ctx, req.(*CancelJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pinpoint_QueryBisection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBisectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PinpointServer).QueryBisection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pinpoint_QueryBisection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PinpointServer).QueryBisection(ctx, req.(*QueryBisectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pinpoint_LegacyJobQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LegacyJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PinpointServer).LegacyJobQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pinpoint_LegacyJobQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PinpointServer).LegacyJobQuery(ctx, req.(*LegacyJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pinpoint_SchedulePairwise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulePairwiseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PinpointServer).SchedulePairwise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pinpoint_SchedulePairwise_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PinpointServer).SchedulePairwise(ctx, req.(*SchedulePairwiseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pinpoint_ScheduleCulpritFinder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleCulpritFinderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PinpointServer).ScheduleCulpritFinder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pinpoint_ScheduleCulpritFinder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PinpointServer).ScheduleCulpritFinder(ctx, req.(*ScheduleCulpritFinderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Pinpoint_ServiceDesc is the grpc.ServiceDesc for Pinpoint service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pinpoint_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pinpoint.v1.Pinpoint",
	HandlerType: (*PinpointServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScheduleBisection",
			Handler:    _Pinpoint_ScheduleBisection_Handler,
		},
		{
			MethodName: "CancelJob",
			Handler:    _Pinpoint_CancelJob_Handler,
		},
		{
			MethodName: "QueryBisection",
			Handler:    _Pinpoint_QueryBisection_Handler,
		},
		{
			MethodName: "LegacyJobQuery",
			Handler:    _Pinpoint_LegacyJobQuery_Handler,
		},
		{
			MethodName: "SchedulePairwise",
			Handler:    _Pinpoint_SchedulePairwise_Handler,
		},
		{
			MethodName: "ScheduleCulpritFinder",
			Handler:    _Pinpoint_ScheduleCulpritFinder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
