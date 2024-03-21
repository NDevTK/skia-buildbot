// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: service.proto

package v1

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
	AnomalyGroupService_CreateAnomalyGroup_FullMethodName = "/anomalygroup.v1.AnomalyGroupService/CreateAnomalyGroup"
	AnomalyGroupService_ReadAnomalyGroup_FullMethodName   = "/anomalygroup.v1.AnomalyGroupService/ReadAnomalyGroup"
	AnomalyGroupService_AddIssueIdToGroup_FullMethodName  = "/anomalygroup.v1.AnomalyGroupService/AddIssueIdToGroup"
	AnomalyGroupService_AddBisectIdToGroup_FullMethodName = "/anomalygroup.v1.AnomalyGroupService/AddBisectIdToGroup"
	AnomalyGroupService_AddAnomalyToGroup_FullMethodName  = "/anomalygroup.v1.AnomalyGroupService/AddAnomalyToGroup"
	AnomalyGroupService_AddCulpritsToGroup_FullMethodName = "/anomalygroup.v1.AnomalyGroupService/AddCulpritsToGroup"
	AnomalyGroupService_FindExistingGroups_FullMethodName = "/anomalygroup.v1.AnomalyGroupService/FindExistingGroups"
)

// AnomalyGroupServiceClient is the client API for AnomalyGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnomalyGroupServiceClient interface {
	// Create a new anomaly group based on a set of criterias.
	// Avoid binding it to a specific regression.
	CreateAnomalyGroup(ctx context.Context, in *CreateAnomalyGroupRequest, opts ...grpc.CallOption) (*CreateAnomalyGroupResponse, error)
	// Read info for an anomaly group.
	ReadAnomalyGroup(ctx context.Context, in *ReadAnomalyGroupRequest, opts ...grpc.CallOption) (*ReadAnomalyGroupResponse, error)
	// Add the filed issue ID for the anomaly group which has action as 'report'.
	AddIssueIdToGroup(ctx context.Context, in *AddIssueIdToGroupRequest, opts ...grpc.CallOption) (*SimpleGroupUpdateResponse, error)
	// Add the launched bisection ID for the anomaly group which has action as 'bisect'.
	AddBisectIdToGroup(ctx context.Context, in *AddBisectIdToGroupRequest, opts ...grpc.CallOption) (*SimpleGroupUpdateResponse, error)
	// Add a new anomaly to the group.
	AddAnomalyToGroup(ctx context.Context, in *AddAnomalyToGroupRequest, opts ...grpc.CallOption) (*AddAnomalyToGroupResponse, error)
	// Add culprits found by a bisection to the group.
	// (Invoked during persisting culprits)
	AddCulpritsToGroup(ctx context.Context, in *AddCulpritsToGroupRequest, opts ...grpc.CallOption) (*SimpleGroupUpdateResponse, error)
	// Find matching anomaly groups based on the criterias.
	// (e.g., from a newly found anomaly).
	FindExistingGroups(ctx context.Context, in *FindExistingGroupsRequest, opts ...grpc.CallOption) (*FindExistingGroupsResponse, error)
}

type anomalyGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnomalyGroupServiceClient(cc grpc.ClientConnInterface) AnomalyGroupServiceClient {
	return &anomalyGroupServiceClient{cc}
}

func (c *anomalyGroupServiceClient) CreateAnomalyGroup(ctx context.Context, in *CreateAnomalyGroupRequest, opts ...grpc.CallOption) (*CreateAnomalyGroupResponse, error) {
	out := new(CreateAnomalyGroupResponse)
	err := c.cc.Invoke(ctx, AnomalyGroupService_CreateAnomalyGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyGroupServiceClient) ReadAnomalyGroup(ctx context.Context, in *ReadAnomalyGroupRequest, opts ...grpc.CallOption) (*ReadAnomalyGroupResponse, error) {
	out := new(ReadAnomalyGroupResponse)
	err := c.cc.Invoke(ctx, AnomalyGroupService_ReadAnomalyGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyGroupServiceClient) AddIssueIdToGroup(ctx context.Context, in *AddIssueIdToGroupRequest, opts ...grpc.CallOption) (*SimpleGroupUpdateResponse, error) {
	out := new(SimpleGroupUpdateResponse)
	err := c.cc.Invoke(ctx, AnomalyGroupService_AddIssueIdToGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyGroupServiceClient) AddBisectIdToGroup(ctx context.Context, in *AddBisectIdToGroupRequest, opts ...grpc.CallOption) (*SimpleGroupUpdateResponse, error) {
	out := new(SimpleGroupUpdateResponse)
	err := c.cc.Invoke(ctx, AnomalyGroupService_AddBisectIdToGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyGroupServiceClient) AddAnomalyToGroup(ctx context.Context, in *AddAnomalyToGroupRequest, opts ...grpc.CallOption) (*AddAnomalyToGroupResponse, error) {
	out := new(AddAnomalyToGroupResponse)
	err := c.cc.Invoke(ctx, AnomalyGroupService_AddAnomalyToGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyGroupServiceClient) AddCulpritsToGroup(ctx context.Context, in *AddCulpritsToGroupRequest, opts ...grpc.CallOption) (*SimpleGroupUpdateResponse, error) {
	out := new(SimpleGroupUpdateResponse)
	err := c.cc.Invoke(ctx, AnomalyGroupService_AddCulpritsToGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyGroupServiceClient) FindExistingGroups(ctx context.Context, in *FindExistingGroupsRequest, opts ...grpc.CallOption) (*FindExistingGroupsResponse, error) {
	out := new(FindExistingGroupsResponse)
	err := c.cc.Invoke(ctx, AnomalyGroupService_FindExistingGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnomalyGroupServiceServer is the server API for AnomalyGroupService service.
// All implementations must embed UnimplementedAnomalyGroupServiceServer
// for forward compatibility
type AnomalyGroupServiceServer interface {
	// Create a new anomaly group based on a set of criterias.
	// Avoid binding it to a specific regression.
	CreateAnomalyGroup(context.Context, *CreateAnomalyGroupRequest) (*CreateAnomalyGroupResponse, error)
	// Read info for an anomaly group.
	ReadAnomalyGroup(context.Context, *ReadAnomalyGroupRequest) (*ReadAnomalyGroupResponse, error)
	// Add the filed issue ID for the anomaly group which has action as 'report'.
	AddIssueIdToGroup(context.Context, *AddIssueIdToGroupRequest) (*SimpleGroupUpdateResponse, error)
	// Add the launched bisection ID for the anomaly group which has action as 'bisect'.
	AddBisectIdToGroup(context.Context, *AddBisectIdToGroupRequest) (*SimpleGroupUpdateResponse, error)
	// Add a new anomaly to the group.
	AddAnomalyToGroup(context.Context, *AddAnomalyToGroupRequest) (*AddAnomalyToGroupResponse, error)
	// Add culprits found by a bisection to the group.
	// (Invoked during persisting culprits)
	AddCulpritsToGroup(context.Context, *AddCulpritsToGroupRequest) (*SimpleGroupUpdateResponse, error)
	// Find matching anomaly groups based on the criterias.
	// (e.g., from a newly found anomaly).
	FindExistingGroups(context.Context, *FindExistingGroupsRequest) (*FindExistingGroupsResponse, error)
	mustEmbedUnimplementedAnomalyGroupServiceServer()
}

// UnimplementedAnomalyGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnomalyGroupServiceServer struct {
}

func (UnimplementedAnomalyGroupServiceServer) CreateAnomalyGroup(context.Context, *CreateAnomalyGroupRequest) (*CreateAnomalyGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnomalyGroup not implemented")
}
func (UnimplementedAnomalyGroupServiceServer) ReadAnomalyGroup(context.Context, *ReadAnomalyGroupRequest) (*ReadAnomalyGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAnomalyGroup not implemented")
}
func (UnimplementedAnomalyGroupServiceServer) AddIssueIdToGroup(context.Context, *AddIssueIdToGroupRequest) (*SimpleGroupUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddIssueIdToGroup not implemented")
}
func (UnimplementedAnomalyGroupServiceServer) AddBisectIdToGroup(context.Context, *AddBisectIdToGroupRequest) (*SimpleGroupUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBisectIdToGroup not implemented")
}
func (UnimplementedAnomalyGroupServiceServer) AddAnomalyToGroup(context.Context, *AddAnomalyToGroupRequest) (*AddAnomalyToGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAnomalyToGroup not implemented")
}
func (UnimplementedAnomalyGroupServiceServer) AddCulpritsToGroup(context.Context, *AddCulpritsToGroupRequest) (*SimpleGroupUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCulpritsToGroup not implemented")
}
func (UnimplementedAnomalyGroupServiceServer) FindExistingGroups(context.Context, *FindExistingGroupsRequest) (*FindExistingGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindExistingGroups not implemented")
}
func (UnimplementedAnomalyGroupServiceServer) mustEmbedUnimplementedAnomalyGroupServiceServer() {}

// UnsafeAnomalyGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnomalyGroupServiceServer will
// result in compilation errors.
type UnsafeAnomalyGroupServiceServer interface {
	mustEmbedUnimplementedAnomalyGroupServiceServer()
}

func RegisterAnomalyGroupServiceServer(s grpc.ServiceRegistrar, srv AnomalyGroupServiceServer) {
	s.RegisterService(&AnomalyGroupService_ServiceDesc, srv)
}

func _AnomalyGroupService_CreateAnomalyGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnomalyGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyGroupServiceServer).CreateAnomalyGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnomalyGroupService_CreateAnomalyGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyGroupServiceServer).CreateAnomalyGroup(ctx, req.(*CreateAnomalyGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnomalyGroupService_ReadAnomalyGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAnomalyGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyGroupServiceServer).ReadAnomalyGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnomalyGroupService_ReadAnomalyGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyGroupServiceServer).ReadAnomalyGroup(ctx, req.(*ReadAnomalyGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnomalyGroupService_AddIssueIdToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddIssueIdToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyGroupServiceServer).AddIssueIdToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnomalyGroupService_AddIssueIdToGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyGroupServiceServer).AddIssueIdToGroup(ctx, req.(*AddIssueIdToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnomalyGroupService_AddBisectIdToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBisectIdToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyGroupServiceServer).AddBisectIdToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnomalyGroupService_AddBisectIdToGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyGroupServiceServer).AddBisectIdToGroup(ctx, req.(*AddBisectIdToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnomalyGroupService_AddAnomalyToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAnomalyToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyGroupServiceServer).AddAnomalyToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnomalyGroupService_AddAnomalyToGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyGroupServiceServer).AddAnomalyToGroup(ctx, req.(*AddAnomalyToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnomalyGroupService_AddCulpritsToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCulpritsToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyGroupServiceServer).AddCulpritsToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnomalyGroupService_AddCulpritsToGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyGroupServiceServer).AddCulpritsToGroup(ctx, req.(*AddCulpritsToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnomalyGroupService_FindExistingGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindExistingGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyGroupServiceServer).FindExistingGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnomalyGroupService_FindExistingGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyGroupServiceServer).FindExistingGroups(ctx, req.(*FindExistingGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AnomalyGroupService_ServiceDesc is the grpc.ServiceDesc for AnomalyGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnomalyGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "anomalygroup.v1.AnomalyGroupService",
	HandlerType: (*AnomalyGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAnomalyGroup",
			Handler:    _AnomalyGroupService_CreateAnomalyGroup_Handler,
		},
		{
			MethodName: "ReadAnomalyGroup",
			Handler:    _AnomalyGroupService_ReadAnomalyGroup_Handler,
		},
		{
			MethodName: "AddIssueIdToGroup",
			Handler:    _AnomalyGroupService_AddIssueIdToGroup_Handler,
		},
		{
			MethodName: "AddBisectIdToGroup",
			Handler:    _AnomalyGroupService_AddBisectIdToGroup_Handler,
		},
		{
			MethodName: "AddAnomalyToGroup",
			Handler:    _AnomalyGroupService_AddAnomalyToGroup_Handler,
		},
		{
			MethodName: "AddCulpritsToGroup",
			Handler:    _AnomalyGroupService_AddCulpritsToGroup_Handler,
		},
		{
			MethodName: "FindExistingGroups",
			Handler:    _AnomalyGroupService_FindExistingGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
