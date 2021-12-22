// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package orchestratorserviceproto

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

// OrchestratorServiceClient is the client API for OrchestratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrchestratorServiceClient interface {
	GetUserByName(ctx context.Context, in *GetUserByNameRequest, opts ...grpc.CallOption) (*GetUserByNameResponse, error)
}

type orchestratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrchestratorServiceClient(cc grpc.ClientConnInterface) OrchestratorServiceClient {
	return &orchestratorServiceClient{cc}
}

func (c *orchestratorServiceClient) GetUserByName(ctx context.Context, in *GetUserByNameRequest, opts ...grpc.CallOption) (*GetUserByNameResponse, error) {
	out := new(GetUserByNameResponse)
	err := c.cc.Invoke(ctx, "/orchestratorserviceproto.OrchestratorService/GetUserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrchestratorServiceServer is the server API for OrchestratorService service.
// All implementations must embed UnimplementedOrchestratorServiceServer
// for forward compatibility
type OrchestratorServiceServer interface {
	GetUserByName(context.Context, *GetUserByNameRequest) (*GetUserByNameResponse, error)
	mustEmbedUnimplementedOrchestratorServiceServer()
}

// UnimplementedOrchestratorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrchestratorServiceServer struct {
}

func (UnimplementedOrchestratorServiceServer) GetUserByName(context.Context, *GetUserByNameRequest) (*GetUserByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByName not implemented")
}
func (UnimplementedOrchestratorServiceServer) mustEmbedUnimplementedOrchestratorServiceServer() {}

// UnsafeOrchestratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrchestratorServiceServer will
// result in compilation errors.
type UnsafeOrchestratorServiceServer interface {
	mustEmbedUnimplementedOrchestratorServiceServer()
}

func RegisterOrchestratorServiceServer(s grpc.ServiceRegistrar, srv OrchestratorServiceServer) {
	s.RegisterService(&OrchestratorService_ServiceDesc, srv)
}

func _OrchestratorService_GetUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServiceServer).GetUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orchestratorserviceproto.OrchestratorService/GetUserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServiceServer).GetUserByName(ctx, req.(*GetUserByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrchestratorService_ServiceDesc is the grpc.ServiceDesc for OrchestratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrchestratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orchestratorserviceproto.OrchestratorService",
	HandlerType: (*OrchestratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByName",
			Handler:    _OrchestratorService_GetUserByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orchestratorService.proto",
}
