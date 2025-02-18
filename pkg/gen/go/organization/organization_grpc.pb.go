// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: organization/organization.proto

package organizationv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	OrganizationService_CreateOrganization_FullMethodName  = "/organization.OrganizationService/CreateOrganization"
	OrganizationService_UpdateOrganization_FullMethodName  = "/organization.OrganizationService/UpdateOrganization"
	OrganizationService_PatchOrganization_FullMethodName   = "/organization.OrganizationService/PatchOrganization"
	OrganizationService_DeleteOrganization_FullMethodName  = "/organization.OrganizationService/DeleteOrganization"
	OrganizationService_GetOrganizationById_FullMethodName = "/organization.OrganizationService/GetOrganizationById"
	OrganizationService_CreateProperty_FullMethodName      = "/organization.OrganizationService/CreateProperty"
	OrganizationService_UpdateProperty_FullMethodName      = "/organization.OrganizationService/UpdateProperty"
	OrganizationService_DeleteProperty_FullMethodName      = "/organization.OrganizationService/DeleteProperty"
)

// OrganizationServiceClient is the client API for OrganizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationServiceClient interface {
	CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*CreateOrganizationResponse, error)
	UpdateOrganization(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*UpdateOrganizationResponse, error)
	PatchOrganization(ctx context.Context, in *PatchOrganizationRequest, opts ...grpc.CallOption) (*PatchOrganizationResponse, error)
	DeleteOrganization(ctx context.Context, in *DeleteOrganizationRequest, opts ...grpc.CallOption) (*DeleteOrganizationResponse, error)
	GetOrganizationById(ctx context.Context, in *GetOrganizationByIdRequest, opts ...grpc.CallOption) (*GetOrganizationByIdResponse, error)
	CreateProperty(ctx context.Context, in *CreatePropertyRequest, opts ...grpc.CallOption) (*CreatePropertyResponse, error)
	UpdateProperty(ctx context.Context, in *UpdatePropertyRequest, opts ...grpc.CallOption) (*UpdatePropertyResponse, error)
	DeleteProperty(ctx context.Context, in *DeletePropertyRequest, opts ...grpc.CallOption) (*DeletePropertyResponse, error)
}

type organizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationServiceClient(cc grpc.ClientConnInterface) OrganizationServiceClient {
	return &organizationServiceClient{cc}
}

func (c *organizationServiceClient) CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*CreateOrganizationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrganizationResponse)
	err := c.cc.Invoke(ctx, OrganizationService_CreateOrganization_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) UpdateOrganization(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*UpdateOrganizationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateOrganizationResponse)
	err := c.cc.Invoke(ctx, OrganizationService_UpdateOrganization_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) PatchOrganization(ctx context.Context, in *PatchOrganizationRequest, opts ...grpc.CallOption) (*PatchOrganizationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PatchOrganizationResponse)
	err := c.cc.Invoke(ctx, OrganizationService_PatchOrganization_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) DeleteOrganization(ctx context.Context, in *DeleteOrganizationRequest, opts ...grpc.CallOption) (*DeleteOrganizationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteOrganizationResponse)
	err := c.cc.Invoke(ctx, OrganizationService_DeleteOrganization_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) GetOrganizationById(ctx context.Context, in *GetOrganizationByIdRequest, opts ...grpc.CallOption) (*GetOrganizationByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrganizationByIdResponse)
	err := c.cc.Invoke(ctx, OrganizationService_GetOrganizationById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) CreateProperty(ctx context.Context, in *CreatePropertyRequest, opts ...grpc.CallOption) (*CreatePropertyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePropertyResponse)
	err := c.cc.Invoke(ctx, OrganizationService_CreateProperty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) UpdateProperty(ctx context.Context, in *UpdatePropertyRequest, opts ...grpc.CallOption) (*UpdatePropertyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePropertyResponse)
	err := c.cc.Invoke(ctx, OrganizationService_UpdateProperty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) DeleteProperty(ctx context.Context, in *DeletePropertyRequest, opts ...grpc.CallOption) (*DeletePropertyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePropertyResponse)
	err := c.cc.Invoke(ctx, OrganizationService_DeleteProperty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationServiceServer is the server API for OrganizationService service.
// All implementations must embed UnimplementedOrganizationServiceServer
// for forward compatibility.
type OrganizationServiceServer interface {
	CreateOrganization(context.Context, *CreateOrganizationRequest) (*CreateOrganizationResponse, error)
	UpdateOrganization(context.Context, *UpdateOrganizationRequest) (*UpdateOrganizationResponse, error)
	PatchOrganization(context.Context, *PatchOrganizationRequest) (*PatchOrganizationResponse, error)
	DeleteOrganization(context.Context, *DeleteOrganizationRequest) (*DeleteOrganizationResponse, error)
	GetOrganizationById(context.Context, *GetOrganizationByIdRequest) (*GetOrganizationByIdResponse, error)
	CreateProperty(context.Context, *CreatePropertyRequest) (*CreatePropertyResponse, error)
	UpdateProperty(context.Context, *UpdatePropertyRequest) (*UpdatePropertyResponse, error)
	DeleteProperty(context.Context, *DeletePropertyRequest) (*DeletePropertyResponse, error)
	mustEmbedUnimplementedOrganizationServiceServer()
}

// UnimplementedOrganizationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrganizationServiceServer struct{}

func (UnimplementedOrganizationServiceServer) CreateOrganization(context.Context, *CreateOrganizationRequest) (*CreateOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrganization not implemented")
}
func (UnimplementedOrganizationServiceServer) UpdateOrganization(context.Context, *UpdateOrganizationRequest) (*UpdateOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrganization not implemented")
}
func (UnimplementedOrganizationServiceServer) PatchOrganization(context.Context, *PatchOrganizationRequest) (*PatchOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchOrganization not implemented")
}
func (UnimplementedOrganizationServiceServer) DeleteOrganization(context.Context, *DeleteOrganizationRequest) (*DeleteOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrganization not implemented")
}
func (UnimplementedOrganizationServiceServer) GetOrganizationById(context.Context, *GetOrganizationByIdRequest) (*GetOrganizationByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganizationById not implemented")
}
func (UnimplementedOrganizationServiceServer) CreateProperty(context.Context, *CreatePropertyRequest) (*CreatePropertyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProperty not implemented")
}
func (UnimplementedOrganizationServiceServer) UpdateProperty(context.Context, *UpdatePropertyRequest) (*UpdatePropertyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProperty not implemented")
}
func (UnimplementedOrganizationServiceServer) DeleteProperty(context.Context, *DeletePropertyRequest) (*DeletePropertyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProperty not implemented")
}
func (UnimplementedOrganizationServiceServer) mustEmbedUnimplementedOrganizationServiceServer() {}
func (UnimplementedOrganizationServiceServer) testEmbeddedByValue()                             {}

// UnsafeOrganizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationServiceServer will
// result in compilation errors.
type UnsafeOrganizationServiceServer interface {
	mustEmbedUnimplementedOrganizationServiceServer()
}

func RegisterOrganizationServiceServer(s grpc.ServiceRegistrar, srv OrganizationServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrganizationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrganizationService_ServiceDesc, srv)
}

func _OrganizationService_CreateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).CreateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrganizationService_CreateOrganization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).CreateOrganization(ctx, req.(*CreateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_UpdateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).UpdateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrganizationService_UpdateOrganization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).UpdateOrganization(ctx, req.(*UpdateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_PatchOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).PatchOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrganizationService_PatchOrganization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).PatchOrganization(ctx, req.(*PatchOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_DeleteOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).DeleteOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrganizationService_DeleteOrganization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).DeleteOrganization(ctx, req.(*DeleteOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_GetOrganizationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).GetOrganizationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrganizationService_GetOrganizationById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).GetOrganizationById(ctx, req.(*GetOrganizationByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_CreateProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePropertyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).CreateProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrganizationService_CreateProperty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).CreateProperty(ctx, req.(*CreatePropertyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_UpdateProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePropertyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).UpdateProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrganizationService_UpdateProperty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).UpdateProperty(ctx, req.(*UpdatePropertyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_DeleteProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePropertyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).DeleteProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrganizationService_DeleteProperty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).DeleteProperty(ctx, req.(*DeletePropertyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrganizationService_ServiceDesc is the grpc.ServiceDesc for OrganizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrganizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "organization.OrganizationService",
	HandlerType: (*OrganizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrganization",
			Handler:    _OrganizationService_CreateOrganization_Handler,
		},
		{
			MethodName: "UpdateOrganization",
			Handler:    _OrganizationService_UpdateOrganization_Handler,
		},
		{
			MethodName: "PatchOrganization",
			Handler:    _OrganizationService_PatchOrganization_Handler,
		},
		{
			MethodName: "DeleteOrganization",
			Handler:    _OrganizationService_DeleteOrganization_Handler,
		},
		{
			MethodName: "GetOrganizationById",
			Handler:    _OrganizationService_GetOrganizationById_Handler,
		},
		{
			MethodName: "CreateProperty",
			Handler:    _OrganizationService_CreateProperty_Handler,
		},
		{
			MethodName: "UpdateProperty",
			Handler:    _OrganizationService_UpdateProperty_Handler,
		},
		{
			MethodName: "DeleteProperty",
			Handler:    _OrganizationService_DeleteProperty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "organization/organization.proto",
}
