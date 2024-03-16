// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0
// source: aspirante.proto

package pruebagrpc

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

// AspiranteServiceClient is the client API for AspiranteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AspiranteServiceClient interface {
	// Obtiene un solo aspirante por ID.
	GetAspirante(ctx context.Context, in *GetAspiranteRequest, opts ...grpc.CallOption) (*AspiranteResponse, error)
	// Obtiene una lista de todos los aspirantes.
	GetAspirantes(ctx context.Context, in *AspiranteRequest, opts ...grpc.CallOption) (*AspirantesResponse, error)
}

type aspiranteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAspiranteServiceClient(cc grpc.ClientConnInterface) AspiranteServiceClient {
	return &aspiranteServiceClient{cc}
}

func (c *aspiranteServiceClient) GetAspirante(ctx context.Context, in *GetAspiranteRequest, opts ...grpc.CallOption) (*AspiranteResponse, error) {
	out := new(AspiranteResponse)
	err := c.cc.Invoke(ctx, "/AspiranteService/GetAspirante", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aspiranteServiceClient) GetAspirantes(ctx context.Context, in *AspiranteRequest, opts ...grpc.CallOption) (*AspirantesResponse, error) {
	out := new(AspirantesResponse)
	err := c.cc.Invoke(ctx, "/AspiranteService/GetAspirantes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AspiranteServiceServer is the server API for AspiranteService service.
// All implementations must embed UnimplementedAspiranteServiceServer
// for forward compatibility
type AspiranteServiceServer interface {
	// Obtiene un solo aspirante por ID.
	GetAspirante(context.Context, *GetAspiranteRequest) (*AspiranteResponse, error)
	// Obtiene una lista de todos los aspirantes.
	GetAspirantes(context.Context, *AspiranteRequest) (*AspirantesResponse, error)
	mustEmbedUnimplementedAspiranteServiceServer()
}

// UnimplementedAspiranteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAspiranteServiceServer struct {
}

func (UnimplementedAspiranteServiceServer) GetAspirante(context.Context, *GetAspiranteRequest) (*AspiranteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAspirante not implemented")
}
func (UnimplementedAspiranteServiceServer) GetAspirantes(context.Context, *AspiranteRequest) (*AspirantesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAspirantes not implemented")
}
func (UnimplementedAspiranteServiceServer) mustEmbedUnimplementedAspiranteServiceServer() {}

// UnsafeAspiranteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AspiranteServiceServer will
// result in compilation errors.
type UnsafeAspiranteServiceServer interface {
	mustEmbedUnimplementedAspiranteServiceServer()
}

func RegisterAspiranteServiceServer(s grpc.ServiceRegistrar, srv AspiranteServiceServer) {
	s.RegisterService(&AspiranteService_ServiceDesc, srv)
}

func _AspiranteService_GetAspirante_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAspiranteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AspiranteServiceServer).GetAspirante(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AspiranteService/GetAspirante",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AspiranteServiceServer).GetAspirante(ctx, req.(*GetAspiranteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AspiranteService_GetAspirantes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AspiranteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AspiranteServiceServer).GetAspirantes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AspiranteService/GetAspirantes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AspiranteServiceServer).GetAspirantes(ctx, req.(*AspiranteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AspiranteService_ServiceDesc is the grpc.ServiceDesc for AspiranteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AspiranteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AspiranteService",
	HandlerType: (*AspiranteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAspirante",
			Handler:    _AspiranteService_GetAspirante_Handler,
		},
		{
			MethodName: "GetAspirantes",
			Handler:    _AspiranteService_GetAspirantes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aspirante.proto",
}
