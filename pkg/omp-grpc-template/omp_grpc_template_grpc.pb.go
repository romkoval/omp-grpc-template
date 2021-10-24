// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package omp_grpc_template

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

// OmpGrpcTemplateServiceClient is the client API for OmpGrpcTemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OmpGrpcTemplateServiceClient interface {
	// DescribeTemplateV1 - Describe a template
	DescribeTemplateV1(ctx context.Context, in *DescribeTemplateV1Request, opts ...grpc.CallOption) (*DescribeTemplateV1Response, error)
}

type ompGrpcTemplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOmpGrpcTemplateServiceClient(cc grpc.ClientConnInterface) OmpGrpcTemplateServiceClient {
	return &ompGrpcTemplateServiceClient{cc}
}

func (c *ompGrpcTemplateServiceClient) DescribeTemplateV1(ctx context.Context, in *DescribeTemplateV1Request, opts ...grpc.CallOption) (*DescribeTemplateV1Response, error) {
	out := new(DescribeTemplateV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.omp_grpc_template.v1.OmpGrpcTemplateService/DescribeTemplateV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OmpGrpcTemplateServiceServer is the server API for OmpGrpcTemplateService service.
// All implementations must embed UnimplementedOmpGrpcTemplateServiceServer
// for forward compatibility
type OmpGrpcTemplateServiceServer interface {
	// DescribeTemplateV1 - Describe a template
	DescribeTemplateV1(context.Context, *DescribeTemplateV1Request) (*DescribeTemplateV1Response, error)
	mustEmbedUnimplementedOmpGrpcTemplateServiceServer()
}

// UnimplementedOmpGrpcTemplateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOmpGrpcTemplateServiceServer struct {
}

func (UnimplementedOmpGrpcTemplateServiceServer) DescribeTemplateV1(context.Context, *DescribeTemplateV1Request) (*DescribeTemplateV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTemplateV1 not implemented")
}
func (UnimplementedOmpGrpcTemplateServiceServer) mustEmbedUnimplementedOmpGrpcTemplateServiceServer() {
}

// UnsafeOmpGrpcTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OmpGrpcTemplateServiceServer will
// result in compilation errors.
type UnsafeOmpGrpcTemplateServiceServer interface {
	mustEmbedUnimplementedOmpGrpcTemplateServiceServer()
}

func RegisterOmpGrpcTemplateServiceServer(s grpc.ServiceRegistrar, srv OmpGrpcTemplateServiceServer) {
	s.RegisterService(&OmpGrpcTemplateService_ServiceDesc, srv)
}

func _OmpGrpcTemplateService_DescribeTemplateV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTemplateV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmpGrpcTemplateServiceServer).DescribeTemplateV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.omp_grpc_template.v1.OmpGrpcTemplateService/DescribeTemplateV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmpGrpcTemplateServiceServer).DescribeTemplateV1(ctx, req.(*DescribeTemplateV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OmpGrpcTemplateService_ServiceDesc is the grpc.ServiceDesc for OmpGrpcTemplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OmpGrpcTemplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozonmp.omp_grpc_template.v1.OmpGrpcTemplateService",
	HandlerType: (*OmpGrpcTemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeTemplateV1",
			Handler:    _OmpGrpcTemplateService_DescribeTemplateV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/omp_grpc_template/omp_grpc_template.proto",
}
