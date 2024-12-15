// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrganizationServiceClient is the client API for OrganizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationServiceClient interface {
	CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*CreateOrganizationResponse, error)
	UpdateOrganization(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*UpdateOrganizationResponse, error)
	ListOrganizations(ctx context.Context, in *ListOrganizationsRequest, opts ...grpc.CallOption) (*ListOrganizationsResponse, error)
	CreateOrganizationDoc(ctx context.Context, in *CreateOrganizationDocRequest, opts ...grpc.CallOption) (*CreateOrganizationDocResponse, error)
	ListOrganizationDocs(ctx context.Context, in *ListOrganizationDocsRequest, opts ...grpc.CallOption) (*ListOrganizationDocsResponse, error)
	CreateApartment(ctx context.Context, in *CreateApartmentRequest, opts ...grpc.CallOption) (*CreateApartmentResponse, error)
	UpdateApartment(ctx context.Context, in *UpdateApartmentRequest, opts ...grpc.CallOption) (*UpdateApartmentResponse, error)
	DeleteApartment(ctx context.Context, in *DeleteApartmentRequest, opts ...grpc.CallOption) (*DeleteApartmentResponse, error)
	ListApartments(ctx context.Context, in *ListApartmentsRequest, opts ...grpc.CallOption) (*ListApartmentsResponse, error)
	ShowOrganizationApartments(ctx context.Context, in *ShowOrganizationApartmentsRequest, opts ...grpc.CallOption) (*ShowOrganizationApartmentsResponse, error)
}

type organizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationServiceClient(cc grpc.ClientConnInterface) OrganizationServiceClient {
	return &organizationServiceClient{cc}
}

func (c *organizationServiceClient) CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*CreateOrganizationResponse, error) {
	out := new(CreateOrganizationResponse)
	err := c.cc.Invoke(ctx, "/organization.OrganizationService/CreateOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) UpdateOrganization(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*UpdateOrganizationResponse, error) {
	out := new(UpdateOrganizationResponse)
	err := c.cc.Invoke(ctx, "/organization.OrganizationService/UpdateOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) ListOrganizations(ctx context.Context, in *ListOrganizationsRequest, opts ...grpc.CallOption) (*ListOrganizationsResponse, error) {
	out := new(ListOrganizationsResponse)
	err := c.cc.Invoke(ctx, "/organization.OrganizationService/ListOrganizations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) CreateOrganizationDoc(ctx context.Context, in *CreateOrganizationDocRequest, opts ...grpc.CallOption) (*CreateOrganizationDocResponse, error) {
	out := new(CreateOrganizationDocResponse)
	err := c.cc.Invoke(ctx, "/organization.OrganizationService/CreateOrganizationDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) ListOrganizationDocs(ctx context.Context, in *ListOrganizationDocsRequest, opts ...grpc.CallOption) (*ListOrganizationDocsResponse, error) {
	out := new(ListOrganizationDocsResponse)
	err := c.cc.Invoke(ctx, "/organization.OrganizationService/ListOrganizationDocs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) CreateApartment(ctx context.Context, in *CreateApartmentRequest, opts ...grpc.CallOption) (*CreateApartmentResponse, error) {
	out := new(CreateApartmentResponse)
	err := c.cc.Invoke(ctx, "/organization.OrganizationService/CreateApartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) UpdateApartment(ctx context.Context, in *UpdateApartmentRequest, opts ...grpc.CallOption) (*UpdateApartmentResponse, error) {
	out := new(UpdateApartmentResponse)
	err := c.cc.Invoke(ctx, "/organization.OrganizationService/UpdateApartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) DeleteApartment(ctx context.Context, in *DeleteApartmentRequest, opts ...grpc.CallOption) (*DeleteApartmentResponse, error) {
	out := new(DeleteApartmentResponse)
	err := c.cc.Invoke(ctx, "/organization.OrganizationService/DeleteApartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) ListApartments(ctx context.Context, in *ListApartmentsRequest, opts ...grpc.CallOption) (*ListApartmentsResponse, error) {
	out := new(ListApartmentsResponse)
	err := c.cc.Invoke(ctx, "/organization.OrganizationService/ListApartments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) ShowOrganizationApartments(ctx context.Context, in *ShowOrganizationApartmentsRequest, opts ...grpc.CallOption) (*ShowOrganizationApartmentsResponse, error) {
	out := new(ShowOrganizationApartmentsResponse)
	err := c.cc.Invoke(ctx, "/organization.OrganizationService/ShowOrganizationApartments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationServiceServer is the server API for OrganizationService service.
// All implementations must embed UnimplementedOrganizationServiceServer
// for forward compatibility
type OrganizationServiceServer interface {
	CreateOrganization(context.Context, *CreateOrganizationRequest) (*CreateOrganizationResponse, error)
	UpdateOrganization(context.Context, *UpdateOrganizationRequest) (*UpdateOrganizationResponse, error)
	ListOrganizations(context.Context, *ListOrganizationsRequest) (*ListOrganizationsResponse, error)
	CreateOrganizationDoc(context.Context, *CreateOrganizationDocRequest) (*CreateOrganizationDocResponse, error)
	ListOrganizationDocs(context.Context, *ListOrganizationDocsRequest) (*ListOrganizationDocsResponse, error)
	CreateApartment(context.Context, *CreateApartmentRequest) (*CreateApartmentResponse, error)
	UpdateApartment(context.Context, *UpdateApartmentRequest) (*UpdateApartmentResponse, error)
	DeleteApartment(context.Context, *DeleteApartmentRequest) (*DeleteApartmentResponse, error)
	ListApartments(context.Context, *ListApartmentsRequest) (*ListApartmentsResponse, error)
	ShowOrganizationApartments(context.Context, *ShowOrganizationApartmentsRequest) (*ShowOrganizationApartmentsResponse, error)
	mustEmbedUnimplementedOrganizationServiceServer()
}

// UnimplementedOrganizationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrganizationServiceServer struct {
}

func (UnimplementedOrganizationServiceServer) CreateOrganization(context.Context, *CreateOrganizationRequest) (*CreateOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrganization not implemented")
}
func (UnimplementedOrganizationServiceServer) UpdateOrganization(context.Context, *UpdateOrganizationRequest) (*UpdateOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrganization not implemented")
}
func (UnimplementedOrganizationServiceServer) ListOrganizations(context.Context, *ListOrganizationsRequest) (*ListOrganizationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrganizations not implemented")
}
func (UnimplementedOrganizationServiceServer) CreateOrganizationDoc(context.Context, *CreateOrganizationDocRequest) (*CreateOrganizationDocResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrganizationDoc not implemented")
}
func (UnimplementedOrganizationServiceServer) ListOrganizationDocs(context.Context, *ListOrganizationDocsRequest) (*ListOrganizationDocsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrganizationDocs not implemented")
}
func (UnimplementedOrganizationServiceServer) CreateApartment(context.Context, *CreateApartmentRequest) (*CreateApartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApartment not implemented")
}
func (UnimplementedOrganizationServiceServer) UpdateApartment(context.Context, *UpdateApartmentRequest) (*UpdateApartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApartment not implemented")
}
func (UnimplementedOrganizationServiceServer) DeleteApartment(context.Context, *DeleteApartmentRequest) (*DeleteApartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApartment not implemented")
}
func (UnimplementedOrganizationServiceServer) ListApartments(context.Context, *ListApartmentsRequest) (*ListApartmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApartments not implemented")
}
func (UnimplementedOrganizationServiceServer) ShowOrganizationApartments(context.Context, *ShowOrganizationApartmentsRequest) (*ShowOrganizationApartmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowOrganizationApartments not implemented")
}
func (UnimplementedOrganizationServiceServer) mustEmbedUnimplementedOrganizationServiceServer() {}

// UnsafeOrganizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationServiceServer will
// result in compilation errors.
type UnsafeOrganizationServiceServer interface {
	mustEmbedUnimplementedOrganizationServiceServer()
}

func RegisterOrganizationServiceServer(s grpc.ServiceRegistrar, srv OrganizationServiceServer) {
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
		FullMethod: "/organization.OrganizationService/CreateOrganization",
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
		FullMethod: "/organization.OrganizationService/UpdateOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).UpdateOrganization(ctx, req.(*UpdateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_ListOrganizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).ListOrganizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization.OrganizationService/ListOrganizations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).ListOrganizations(ctx, req.(*ListOrganizationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_CreateOrganizationDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationDocRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).CreateOrganizationDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization.OrganizationService/CreateOrganizationDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).CreateOrganizationDoc(ctx, req.(*CreateOrganizationDocRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_ListOrganizationDocs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationDocsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).ListOrganizationDocs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization.OrganizationService/ListOrganizationDocs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).ListOrganizationDocs(ctx, req.(*ListOrganizationDocsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_CreateApartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).CreateApartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization.OrganizationService/CreateApartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).CreateApartment(ctx, req.(*CreateApartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_UpdateApartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).UpdateApartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization.OrganizationService/UpdateApartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).UpdateApartment(ctx, req.(*UpdateApartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_DeleteApartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).DeleteApartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization.OrganizationService/DeleteApartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).DeleteApartment(ctx, req.(*DeleteApartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_ListApartments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApartmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).ListApartments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization.OrganizationService/ListApartments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).ListApartments(ctx, req.(*ListApartmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_ShowOrganizationApartments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowOrganizationApartmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).ShowOrganizationApartments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization.OrganizationService/ShowOrganizationApartments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).ShowOrganizationApartments(ctx, req.(*ShowOrganizationApartmentsRequest))
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
			MethodName: "ListOrganizations",
			Handler:    _OrganizationService_ListOrganizations_Handler,
		},
		{
			MethodName: "CreateOrganizationDoc",
			Handler:    _OrganizationService_CreateOrganizationDoc_Handler,
		},
		{
			MethodName: "ListOrganizationDocs",
			Handler:    _OrganizationService_ListOrganizationDocs_Handler,
		},
		{
			MethodName: "CreateApartment",
			Handler:    _OrganizationService_CreateApartment_Handler,
		},
		{
			MethodName: "UpdateApartment",
			Handler:    _OrganizationService_UpdateApartment_Handler,
		},
		{
			MethodName: "DeleteApartment",
			Handler:    _OrganizationService_DeleteApartment_Handler,
		},
		{
			MethodName: "ListApartments",
			Handler:    _OrganizationService_ListApartments_Handler,
		},
		{
			MethodName: "ShowOrganizationApartments",
			Handler:    _OrganizationService_ShowOrganizationApartments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "organization/organization.proto",
}
