// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.1
// source: proto/driver/feed.proto

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
	DriverFeed_Upload_FullMethodName = "/proto.DriverFeed/Upload"
)

// DriverFeedClient is the client API for DriverFeed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverFeedClient interface {
	Upload(ctx context.Context, in *UploadFeedRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[UploadFeedResponse], error)
}

type driverFeedClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverFeedClient(cc grpc.ClientConnInterface) DriverFeedClient {
	return &driverFeedClient{cc}
}

func (c *driverFeedClient) Upload(ctx context.Context, in *UploadFeedRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[UploadFeedResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DriverFeed_ServiceDesc.Streams[0], DriverFeed_Upload_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UploadFeedRequest, UploadFeedResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DriverFeed_UploadClient = grpc.ServerStreamingClient[UploadFeedResponse]

// DriverFeedServer is the server API for DriverFeed service.
// All implementations must embed UnimplementedDriverFeedServer
// for forward compatibility.
type DriverFeedServer interface {
	Upload(*UploadFeedRequest, grpc.ServerStreamingServer[UploadFeedResponse]) error
	mustEmbedUnimplementedDriverFeedServer()
}

// UnimplementedDriverFeedServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDriverFeedServer struct{}

func (UnimplementedDriverFeedServer) Upload(*UploadFeedRequest, grpc.ServerStreamingServer[UploadFeedResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedDriverFeedServer) mustEmbedUnimplementedDriverFeedServer() {}
func (UnimplementedDriverFeedServer) testEmbeddedByValue()                    {}

// UnsafeDriverFeedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverFeedServer will
// result in compilation errors.
type UnsafeDriverFeedServer interface {
	mustEmbedUnimplementedDriverFeedServer()
}

func RegisterDriverFeedServer(s grpc.ServiceRegistrar, srv DriverFeedServer) {
	// If the following call pancis, it indicates UnimplementedDriverFeedServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DriverFeed_ServiceDesc, srv)
}

func _DriverFeed_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UploadFeedRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DriverFeedServer).Upload(m, &grpc.GenericServerStream[UploadFeedRequest, UploadFeedResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DriverFeed_UploadServer = grpc.ServerStreamingServer[UploadFeedResponse]

// DriverFeed_ServiceDesc is the grpc.ServiceDesc for DriverFeed service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverFeed_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DriverFeed",
	HandlerType: (*DriverFeedServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _DriverFeed_Upload_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/driver/feed.proto",
}
