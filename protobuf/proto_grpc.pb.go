// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: proto.proto

package protobuf

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

// CustomerServicesClient is the client API for CustomerServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerServicesClient interface {
	CreateCustomer(ctx context.Context, in *NewCustomer, opts ...grpc.CallOption) (*Customer, error)
	GetCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*Customer, error)
	GetAllCustomers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllCustomersList, error)
	UpdateCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error)
	DeleteCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*Customer, error)
}

type customerServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServicesClient(cc grpc.ClientConnInterface) CustomerServicesClient {
	return &customerServicesClient{cc}
}

func (c *customerServicesClient) CreateCustomer(ctx context.Context, in *NewCustomer, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/cms.CustomerServices/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServicesClient) GetCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/cms.CustomerServices/GetCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServicesClient) GetAllCustomers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllCustomersList, error) {
	out := new(AllCustomersList)
	err := c.cc.Invoke(ctx, "/cms.CustomerServices/GetAllCustomers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServicesClient) UpdateCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/cms.CustomerServices/UpdateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServicesClient) DeleteCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/cms.CustomerServices/DeleteCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServicesServer is the server API for CustomerServices service.
// All implementations must embed UnimplementedCustomerServicesServer
// for forward compatibility
type CustomerServicesServer interface {
	CreateCustomer(context.Context, *NewCustomer) (*Customer, error)
	GetCustomer(context.Context, *CustomerRequest) (*Customer, error)
	GetAllCustomers(context.Context, *Request) (*AllCustomersList, error)
	UpdateCustomer(context.Context, *Customer) (*Customer, error)
	DeleteCustomer(context.Context, *CustomerRequest) (*Customer, error)
	mustEmbedUnimplementedCustomerServicesServer()
}

// UnimplementedCustomerServicesServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServicesServer struct {
}

func (UnimplementedCustomerServicesServer) CreateCustomer(context.Context, *NewCustomer) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedCustomerServicesServer) GetCustomer(context.Context, *CustomerRequest) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}
func (UnimplementedCustomerServicesServer) GetAllCustomers(context.Context, *Request) (*AllCustomersList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCustomers not implemented")
}
func (UnimplementedCustomerServicesServer) UpdateCustomer(context.Context, *Customer) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (UnimplementedCustomerServicesServer) DeleteCustomer(context.Context, *CustomerRequest) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomer not implemented")
}
func (UnimplementedCustomerServicesServer) mustEmbedUnimplementedCustomerServicesServer() {}

// UnsafeCustomerServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServicesServer will
// result in compilation errors.
type UnsafeCustomerServicesServer interface {
	mustEmbedUnimplementedCustomerServicesServer()
}

func RegisterCustomerServicesServer(s grpc.ServiceRegistrar, srv CustomerServicesServer) {
	s.RegisterService(&CustomerServices_ServiceDesc, srv)
}

func _CustomerServices_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCustomer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServicesServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.CustomerServices/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServicesServer).CreateCustomer(ctx, req.(*NewCustomer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerServices_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServicesServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.CustomerServices/GetCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServicesServer).GetCustomer(ctx, req.(*CustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerServices_GetAllCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServicesServer).GetAllCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.CustomerServices/GetAllCustomers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServicesServer).GetAllCustomers(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerServices_UpdateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServicesServer).UpdateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.CustomerServices/UpdateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServicesServer).UpdateCustomer(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerServices_DeleteCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServicesServer).DeleteCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.CustomerServices/DeleteCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServicesServer).DeleteCustomer(ctx, req.(*CustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerServices_ServiceDesc is the grpc.ServiceDesc for CustomerServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cms.CustomerServices",
	HandlerType: (*CustomerServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _CustomerServices_CreateCustomer_Handler,
		},
		{
			MethodName: "GetCustomer",
			Handler:    _CustomerServices_GetCustomer_Handler,
		},
		{
			MethodName: "GetAllCustomers",
			Handler:    _CustomerServices_GetAllCustomers_Handler,
		},
		{
			MethodName: "UpdateCustomer",
			Handler:    _CustomerServices_UpdateCustomer_Handler,
		},
		{
			MethodName: "DeleteCustomer",
			Handler:    _CustomerServices_DeleteCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}
