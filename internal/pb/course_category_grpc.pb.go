// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: proto/course_category.proto

package pb

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

// CateogryServiceClient is the client API for CateogryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CateogryServiceClient interface {
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
}

type cateogryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCateogryServiceClient(cc grpc.ClientConnInterface) CateogryServiceClient {
	return &cateogryServiceClient{cc}
}

func (c *cateogryServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, "/pb.CateogryService/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CateogryServiceServer is the server API for CateogryService service.
// All implementations must embed UnimplementedCateogryServiceServer
// for forward compatibility
type CateogryServiceServer interface {
	CreateCategory(context.Context, *CreateCategoryRequest) (*CategoryResponse, error)
	mustEmbedUnimplementedCateogryServiceServer()
}

// UnimplementedCateogryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCateogryServiceServer struct {
}

func (UnimplementedCateogryServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedCateogryServiceServer) mustEmbedUnimplementedCateogryServiceServer() {}

// UnsafeCateogryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CateogryServiceServer will
// result in compilation errors.
type UnsafeCateogryServiceServer interface {
	mustEmbedUnimplementedCateogryServiceServer()
}

func RegisterCateogryServiceServer(s grpc.ServiceRegistrar, srv CateogryServiceServer) {
	s.RegisterService(&CateogryService_ServiceDesc, srv)
}

func _CateogryService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CateogryServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CateogryService/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CateogryServiceServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CateogryService_ServiceDesc is the grpc.ServiceDesc for CateogryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CateogryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CateogryService",
	HandlerType: (*CateogryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCategory",
			Handler:    _CateogryService_CreateCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/course_category.proto",
}
