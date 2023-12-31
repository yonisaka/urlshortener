// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/service.proto

package grpc

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

// URLShortenerServiceClient is the client API for URLShortenerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type URLShortenerServiceClient interface {
	// CreateURLShortener creates a new record for URL Shortener.
	// Only single transaction will create by this RPC for a specific User.
	CreateURLShortener(ctx context.Context, in *CreateURLShortenerRequest, opts ...grpc.CallOption) (*URLShortener, error)
	// ListURLShortener get the list of records for URL Shortener.
	// The record can be filtered by specific User.
	ListURLShortener(ctx context.Context, in *ListURLShortenerRequest, opts ...grpc.CallOption) (*ListURLShortenerResponse, error)
	// GetShortenedURL get the shortened URL.
	GetShortenedURL(ctx context.Context, in *GetShortenedURLRequest, opts ...grpc.CallOption) (*URLShortener, error)
}

type uRLShortenerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewURLShortenerServiceClient(cc grpc.ClientConnInterface) URLShortenerServiceClient {
	return &uRLShortenerServiceClient{cc}
}

func (c *uRLShortenerServiceClient) CreateURLShortener(ctx context.Context, in *CreateURLShortenerRequest, opts ...grpc.CallOption) (*URLShortener, error) {
	out := new(URLShortener)
	err := c.cc.Invoke(ctx, "/URLShortenerService/CreateURLShortener", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLShortenerServiceClient) ListURLShortener(ctx context.Context, in *ListURLShortenerRequest, opts ...grpc.CallOption) (*ListURLShortenerResponse, error) {
	out := new(ListURLShortenerResponse)
	err := c.cc.Invoke(ctx, "/URLShortenerService/ListURLShortener", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLShortenerServiceClient) GetShortenedURL(ctx context.Context, in *GetShortenedURLRequest, opts ...grpc.CallOption) (*URLShortener, error) {
	out := new(URLShortener)
	err := c.cc.Invoke(ctx, "/URLShortenerService/GetShortenedURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// URLShortenerServiceServer is the server API for URLShortenerService service.
// All implementations must embed UnimplementedURLShortenerServiceServer
// for forward compatibility
type URLShortenerServiceServer interface {
	// CreateURLShortener creates a new record for URL Shortener.
	// Only single transaction will create by this RPC for a specific User.
	CreateURLShortener(context.Context, *CreateURLShortenerRequest) (*URLShortener, error)
	// ListURLShortener get the list of records for URL Shortener.
	// The record can be filtered by specific User.
	ListURLShortener(context.Context, *ListURLShortenerRequest) (*ListURLShortenerResponse, error)
	// GetShortenedURL get the shortened URL.
	GetShortenedURL(context.Context, *GetShortenedURLRequest) (*URLShortener, error)
	mustEmbedUnimplementedURLShortenerServiceServer()
}

// UnimplementedURLShortenerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedURLShortenerServiceServer struct {
}

func (UnimplementedURLShortenerServiceServer) CreateURLShortener(context.Context, *CreateURLShortenerRequest) (*URLShortener, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateURLShortener not implemented")
}
func (UnimplementedURLShortenerServiceServer) ListURLShortener(context.Context, *ListURLShortenerRequest) (*ListURLShortenerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListURLShortener not implemented")
}
func (UnimplementedURLShortenerServiceServer) GetShortenedURL(context.Context, *GetShortenedURLRequest) (*URLShortener, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShortenedURL not implemented")
}
func (UnimplementedURLShortenerServiceServer) mustEmbedUnimplementedURLShortenerServiceServer() {}

// UnsafeURLShortenerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to URLShortenerServiceServer will
// result in compilation errors.
type UnsafeURLShortenerServiceServer interface {
	mustEmbedUnimplementedURLShortenerServiceServer()
}

func RegisterURLShortenerServiceServer(s grpc.ServiceRegistrar, srv URLShortenerServiceServer) {
	s.RegisterService(&URLShortenerService_ServiceDesc, srv)
}

func _URLShortenerService_CreateURLShortener_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateURLShortenerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLShortenerServiceServer).CreateURLShortener(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/URLShortenerService/CreateURLShortener",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLShortenerServiceServer).CreateURLShortener(ctx, req.(*CreateURLShortenerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLShortenerService_ListURLShortener_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListURLShortenerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLShortenerServiceServer).ListURLShortener(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/URLShortenerService/ListURLShortener",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLShortenerServiceServer).ListURLShortener(ctx, req.(*ListURLShortenerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLShortenerService_GetShortenedURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShortenedURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLShortenerServiceServer).GetShortenedURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/URLShortenerService/GetShortenedURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLShortenerServiceServer).GetShortenedURL(ctx, req.(*GetShortenedURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// URLShortenerService_ServiceDesc is the grpc.ServiceDesc for URLShortenerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var URLShortenerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "URLShortenerService",
	HandlerType: (*URLShortenerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateURLShortener",
			Handler:    _URLShortenerService_CreateURLShortener_Handler,
		},
		{
			MethodName: "ListURLShortener",
			Handler:    _URLShortenerService_ListURLShortener_Handler,
		},
		{
			MethodName: "GetShortenedURL",
			Handler:    _URLShortenerService_GetShortenedURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
