// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: product/product.proto

package product

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Product_CreateProduct_FullMethodName = "/Product/CreateProduct"
)

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	CreateProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) CreateProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, Product_CreateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations should embed UnimplementedProductServer
// for forward compatibility
type ProductServer interface {
	CreateProduct(context.Context, *ProductRequest) (*ProductResponse, error)
}

// UnimplementedProductServer should be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (UnimplementedProductServer) CreateProduct(context.Context, *ProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CreateProduct(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _Product_CreateProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product/product.proto",
}
