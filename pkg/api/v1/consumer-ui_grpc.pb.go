// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// DigitalConsumerUIClient is the client API for DigitalConsumerUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DigitalConsumerUIClient interface {
	// ListEngineSpecs returns a list of Consumer Engine(s) that can be started through the Digital Consumer UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (DigitalConsumerUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the Digital Consumer UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type digitalConsumerUIClient struct {
	cc grpc.ClientConnInterface
}

func NewDigitalConsumerUIClient(cc grpc.ClientConnInterface) DigitalConsumerUIClient {
	return &digitalConsumerUIClient{cc}
}

func (c *digitalConsumerUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (DigitalConsumerUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &DigitalConsumerUI_ServiceDesc.Streams[0], "/v1.DigitalConsumerUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &digitalConsumerUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DigitalConsumerUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type digitalConsumerUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *digitalConsumerUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *digitalConsumerUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.DigitalConsumerUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DigitalConsumerUIServer is the server API for DigitalConsumerUI service.
// All implementations must embed UnimplementedDigitalConsumerUIServer
// for forward compatibility
type DigitalConsumerUIServer interface {
	// ListEngineSpecs returns a list of Consumer Engine(s) that can be started through the Digital Consumer UI.
	ListEngineSpecs(*ListEngineSpecsRequest, DigitalConsumerUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the Digital Consumer UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedDigitalConsumerUIServer()
}

// UnimplementedDigitalConsumerUIServer must be embedded to have forward compatible implementations.
type UnimplementedDigitalConsumerUIServer struct {
}

func (UnimplementedDigitalConsumerUIServer) ListEngineSpecs(*ListEngineSpecsRequest, DigitalConsumerUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedDigitalConsumerUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedDigitalConsumerUIServer) mustEmbedUnimplementedDigitalConsumerUIServer() {}

// UnsafeDigitalConsumerUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DigitalConsumerUIServer will
// result in compilation errors.
type UnsafeDigitalConsumerUIServer interface {
	mustEmbedUnimplementedDigitalConsumerUIServer()
}

func RegisterDigitalConsumerUIServer(s grpc.ServiceRegistrar, srv DigitalConsumerUIServer) {
	s.RegisterService(&DigitalConsumerUI_ServiceDesc, srv)
}

func _DigitalConsumerUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DigitalConsumerUIServer).ListEngineSpecs(m, &digitalConsumerUIListEngineSpecsServer{stream})
}

type DigitalConsumerUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type digitalConsumerUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *digitalConsumerUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DigitalConsumerUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DigitalConsumerUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DigitalConsumerUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DigitalConsumerUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DigitalConsumerUI_ServiceDesc is the grpc.ServiceDesc for DigitalConsumerUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DigitalConsumerUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.DigitalConsumerUI",
	HandlerType: (*DigitalConsumerUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _DigitalConsumerUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _DigitalConsumerUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "consumer-ui.proto",
}
