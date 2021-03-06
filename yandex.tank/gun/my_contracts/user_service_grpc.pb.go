// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: user_service.proto

package my_contracts

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

// KionServiceClient is the client API for KionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KionServiceClient interface {
	// Создание записи
	CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error)
	// Получение записи
	GetLatestRecord(ctx context.Context, in *GetLatestRecordRequest, opts ...grpc.CallOption) (*GetLatestRecordResponse, error)
}

type kionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKionServiceClient(cc grpc.ClientConnInterface) KionServiceClient {
	return &kionServiceClient{cc}
}

func (c *kionServiceClient) CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error) {
	out := new(CreateRecordResponse)
	err := c.cc.Invoke(ctx, "/KionService/CreateRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kionServiceClient) GetLatestRecord(ctx context.Context, in *GetLatestRecordRequest, opts ...grpc.CallOption) (*GetLatestRecordResponse, error) {
	out := new(GetLatestRecordResponse)
	err := c.cc.Invoke(ctx, "/KionService/GetLatestRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KionServiceServer is the server API for KionService service.
// All implementations must embed UnimplementedKionServiceServer
// for forward compatibility
type KionServiceServer interface {
	// Создание записи
	CreateRecord(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error)
	// Получение записи
	GetLatestRecord(context.Context, *GetLatestRecordRequest) (*GetLatestRecordResponse, error)
	mustEmbedUnimplementedKionServiceServer()
}

// UnimplementedKionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKionServiceServer struct {
}

func (UnimplementedKionServiceServer) CreateRecord(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecord not implemented")
}
func (UnimplementedKionServiceServer) GetLatestRecord(context.Context, *GetLatestRecordRequest) (*GetLatestRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestRecord not implemented")
}
func (UnimplementedKionServiceServer) mustEmbedUnimplementedKionServiceServer() {}

// UnsafeKionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KionServiceServer will
// result in compilation errors.
type UnsafeKionServiceServer interface {
	mustEmbedUnimplementedKionServiceServer()
}

func RegisterKionServiceServer(s grpc.ServiceRegistrar, srv KionServiceServer) {
	s.RegisterService(&KionService_ServiceDesc, srv)
}

func _KionService_CreateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KionServiceServer).CreateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KionService/CreateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KionServiceServer).CreateRecord(ctx, req.(*CreateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KionService_GetLatestRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KionServiceServer).GetLatestRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KionService/GetLatestRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KionServiceServer).GetLatestRecord(ctx, req.(*GetLatestRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KionService_ServiceDesc is the grpc.ServiceDesc for KionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "KionService",
	HandlerType: (*KionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRecord",
			Handler:    _KionService_CreateRecord_Handler,
		},
		{
			MethodName: "GetLatestRecord",
			Handler:    _KionService_GetLatestRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}
