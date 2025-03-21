// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.1
// source: proto/driver/metadata.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DriverMetadata_DirectoryList_FullMethodName = "/proto.DriverMetadata/DirectoryList"
)

// DriverMetadataClient is the client API for DriverMetadata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverMetadataClient interface {
	DirectoryList(ctx context.Context, in *DirectoryListRequest, opts ...grpc.CallOption) (*DirectoryListResponse, error)
}

type driverMetadataClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverMetadataClient(cc grpc.ClientConnInterface) DriverMetadataClient {
	return &driverMetadataClient{cc}
}

func (c *driverMetadataClient) DirectoryList(ctx context.Context, in *DirectoryListRequest, opts ...grpc.CallOption) (*DirectoryListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DirectoryListResponse)
	err := c.cc.Invoke(ctx, DriverMetadata_DirectoryList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverMetadataServer is the server API for DriverMetadata service.
// All implementations must embed UnimplementedDriverMetadataServer
// for forward compatibility.
type DriverMetadataServer interface {
	DirectoryList(context.Context, *DirectoryListRequest) (*DirectoryListResponse, error)
	mustEmbedUnimplementedDriverMetadataServer()
}

// UnimplementedDriverMetadataServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDriverMetadataServer struct{}

func (UnimplementedDriverMetadataServer) DirectoryList(context.Context, *DirectoryListRequest) (*DirectoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DirectoryList not implemented")
}
func (UnimplementedDriverMetadataServer) mustEmbedUnimplementedDriverMetadataServer() {}
func (UnimplementedDriverMetadataServer) testEmbeddedByValue()                        {}

// UnsafeDriverMetadataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverMetadataServer will
// result in compilation errors.
type UnsafeDriverMetadataServer interface {
	mustEmbedUnimplementedDriverMetadataServer()
}

func RegisterDriverMetadataServer(s grpc.ServiceRegistrar, srv DriverMetadataServer) {
	// If the following call pancis, it indicates UnimplementedDriverMetadataServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DriverMetadata_ServiceDesc, srv)
}

func _DriverMetadata_DirectoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DirectoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverMetadataServer).DirectoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverMetadata_DirectoryList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverMetadataServer).DirectoryList(ctx, req.(*DirectoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverMetadata_ServiceDesc is the grpc.ServiceDesc for DriverMetadata service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverMetadata_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DriverMetadata",
	HandlerType: (*DriverMetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DirectoryList",
			Handler:    _DriverMetadata_DirectoryList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/driver/metadata.proto",
}
