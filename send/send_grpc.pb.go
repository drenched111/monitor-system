// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package send

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

// SendServiceClient is the client API for SendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendServiceClient interface {
	Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendRsp, error)
}

type sendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSendServiceClient(cc grpc.ClientConnInterface) SendServiceClient {
	return &sendServiceClient{cc}
}

func (c *sendServiceClient) Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendRsp, error) {
	out := new(SendRsp)
	err := c.cc.Invoke(ctx, "/SendService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendServiceServer is the server API for SendService service.
// All implementations must embed UnimplementedSendServiceServer
// for forward compatibility
type SendServiceServer interface {
	Send(context.Context, *SendReq) (*SendRsp, error)
	mustEmbedUnimplementedSendServiceServer()
}

// UnimplementedSendServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSendServiceServer struct {
}

func (UnimplementedSendServiceServer) Send(context.Context, *SendReq) (*SendRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedSendServiceServer) mustEmbedUnimplementedSendServiceServer() {}

// UnsafeSendServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendServiceServer will
// result in compilation errors.
type UnsafeSendServiceServer interface {
	mustEmbedUnimplementedSendServiceServer()
}

func RegisterSendServiceServer(s grpc.ServiceRegistrar, srv SendServiceServer) {
	s.RegisterService(&SendService_ServiceDesc, srv)
}

func _SendService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SendService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendServiceServer).Send(ctx, req.(*SendReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SendService_ServiceDesc is the grpc.ServiceDesc for SendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SendService",
	HandlerType: (*SendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _SendService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "send.proto",
}