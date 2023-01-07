// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: table/v1/table.proto

package v1

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

// TableServiceClient is the client API for TableService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TableServiceClient interface {
	CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*CreateTableResponse, error)
}

type tableServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTableServiceClient(cc grpc.ClientConnInterface) TableServiceClient {
	return &tableServiceClient{cc}
}

func (c *tableServiceClient) CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*CreateTableResponse, error) {
	out := new(CreateTableResponse)
	err := c.cc.Invoke(ctx, "/singhhp1069.go_sass.TableService/CreateTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TableServiceServer is the server API for TableService service.
// All implementations should embed UnimplementedTableServiceServer
// for forward compatibility
type TableServiceServer interface {
	CreateTable(context.Context, *CreateTableRequest) (*CreateTableResponse, error)
}

// UnimplementedTableServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTableServiceServer struct {
}

func (UnimplementedTableServiceServer) CreateTable(context.Context, *CreateTableRequest) (*CreateTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTable not implemented")
}

// UnsafeTableServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TableServiceServer will
// result in compilation errors.
type UnsafeTableServiceServer interface {
	mustEmbedUnimplementedTableServiceServer()
}

func RegisterTableServiceServer(s grpc.ServiceRegistrar, srv TableServiceServer) {
	s.RegisterService(&TableService_ServiceDesc, srv)
}

func _TableService_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServiceServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/singhhp1069.go_sass.TableService/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServiceServer).CreateTable(ctx, req.(*CreateTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TableService_ServiceDesc is the grpc.ServiceDesc for TableService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TableService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "singhhp1069.go_sass.TableService",
	HandlerType: (*TableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTable",
			Handler:    _TableService_CreateTable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "table/v1/table.proto",
}