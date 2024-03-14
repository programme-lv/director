// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: msg/director.proto

package msg

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

// DirectorClient is the client API for Director service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirectorClient interface {
	EvaluateSubmission(ctx context.Context, in *EvaluationRequest, opts ...grpc.CallOption) (Director_EvaluateSubmissionClient, error)
}

type directorClient struct {
	cc grpc.ClientConnInterface
}

func NewDirectorClient(cc grpc.ClientConnInterface) DirectorClient {
	return &directorClient{cc}
}

func (c *directorClient) EvaluateSubmission(ctx context.Context, in *EvaluationRequest, opts ...grpc.CallOption) (Director_EvaluateSubmissionClient, error) {
	stream, err := c.cc.NewStream(ctx, &Director_ServiceDesc.Streams[0], "/msg.Director/EvaluateSubmission", opts...)
	if err != nil {
		return nil, err
	}
	x := &directorEvaluateSubmissionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Director_EvaluateSubmissionClient interface {
	Recv() (*EvaluationFeedback, error)
	grpc.ClientStream
}

type directorEvaluateSubmissionClient struct {
	grpc.ClientStream
}

func (x *directorEvaluateSubmissionClient) Recv() (*EvaluationFeedback, error) {
	m := new(EvaluationFeedback)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DirectorServer is the server API for Director service.
// All implementations must embed UnimplementedDirectorServer
// for forward compatibility
type DirectorServer interface {
	EvaluateSubmission(*EvaluationRequest, Director_EvaluateSubmissionServer) error
	mustEmbedUnimplementedDirectorServer()
}

// UnimplementedDirectorServer must be embedded to have forward compatible implementations.
type UnimplementedDirectorServer struct {
}

func (UnimplementedDirectorServer) EvaluateSubmission(*EvaluationRequest, Director_EvaluateSubmissionServer) error {
	return status.Errorf(codes.Unimplemented, "method EvaluateSubmission not implemented")
}
func (UnimplementedDirectorServer) mustEmbedUnimplementedDirectorServer() {}

// UnsafeDirectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirectorServer will
// result in compilation errors.
type UnsafeDirectorServer interface {
	mustEmbedUnimplementedDirectorServer()
}

func RegisterDirectorServer(s grpc.ServiceRegistrar, srv DirectorServer) {
	s.RegisterService(&Director_ServiceDesc, srv)
}

func _Director_EvaluateSubmission_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EvaluationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DirectorServer).EvaluateSubmission(m, &directorEvaluateSubmissionServer{stream})
}

type Director_EvaluateSubmissionServer interface {
	Send(*EvaluationFeedback) error
	grpc.ServerStream
}

type directorEvaluateSubmissionServer struct {
	grpc.ServerStream
}

func (x *directorEvaluateSubmissionServer) Send(m *EvaluationFeedback) error {
	return x.ServerStream.SendMsg(m)
}

// Director_ServiceDesc is the grpc.ServiceDesc for Director service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Director_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msg.Director",
	HandlerType: (*DirectorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EvaluateSubmission",
			Handler:       _Director_EvaluateSubmission_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "msg/director.proto",
}
