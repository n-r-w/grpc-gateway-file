// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: service.proto

package proto

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Service_DownloadFile_FullMethodName           = "/collision.servicepb.Service/DownloadFile"
	Service_UploadFile_FullMethodName             = "/collision.servicepb.Service/UploadFile"
	Service_UploadMultipleFiles_FullMethodName    = "/collision.servicepb.Service/UploadMultipleFiles"
	Service_UploadToAnotherService_FullMethodName = "/collision.servicepb.Service/UploadToAnotherService"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// download file
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (Service_DownloadFileClient, error)
	// upload file
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (Service_UploadFileClient, error)
	// upload multiple files
	UploadMultipleFiles(ctx context.Context, opts ...grpc.CallOption) (Service_UploadMultipleFilesClient, error)
	// upload with stream to another service
	UploadToAnotherService(ctx context.Context, opts ...grpc.CallOption) (Service_UploadToAnotherServiceClient, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (Service_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[0], Service_DownloadFile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_DownloadFileClient interface {
	Recv() (*httpbody.HttpBody, error)
	grpc.ClientStream
}

type serviceDownloadFileClient struct {
	grpc.ClientStream
}

func (x *serviceDownloadFileClient) Recv() (*httpbody.HttpBody, error) {
	m := new(httpbody.HttpBody)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (Service_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[1], Service_UploadFile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceUploadFileClient{stream}
	return x, nil
}

type Service_UploadFileClient interface {
	Send(*httpbody.HttpBody) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type serviceUploadFileClient struct {
	grpc.ClientStream
}

func (x *serviceUploadFileClient) Send(m *httpbody.HttpBody) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceUploadFileClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) UploadMultipleFiles(ctx context.Context, opts ...grpc.CallOption) (Service_UploadMultipleFilesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[2], Service_UploadMultipleFiles_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceUploadMultipleFilesClient{stream}
	return x, nil
}

type Service_UploadMultipleFilesClient interface {
	Send(*httpbody.HttpBody) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type serviceUploadMultipleFilesClient struct {
	grpc.ClientStream
}

func (x *serviceUploadMultipleFilesClient) Send(m *httpbody.HttpBody) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceUploadMultipleFilesClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) UploadToAnotherService(ctx context.Context, opts ...grpc.CallOption) (Service_UploadToAnotherServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[3], Service_UploadToAnotherService_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceUploadToAnotherServiceClient{stream}
	return x, nil
}

type Service_UploadToAnotherServiceClient interface {
	Send(*httpbody.HttpBody) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type serviceUploadToAnotherServiceClient struct {
	grpc.ClientStream
}

func (x *serviceUploadToAnotherServiceClient) Send(m *httpbody.HttpBody) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceUploadToAnotherServiceClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// download file
	DownloadFile(*DownloadFileRequest, Service_DownloadFileServer) error
	// upload file
	UploadFile(Service_UploadFileServer) error
	// upload multiple files
	UploadMultipleFiles(Service_UploadMultipleFilesServer) error
	// upload with stream to another service
	UploadToAnotherService(Service_UploadToAnotherServiceServer) error
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) DownloadFile(*DownloadFileRequest, Service_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedServiceServer) UploadFile(Service_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedServiceServer) UploadMultipleFiles(Service_UploadMultipleFilesServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadMultipleFiles not implemented")
}
func (UnimplementedServiceServer) UploadToAnotherService(Service_UploadToAnotherServiceServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadToAnotherService not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).DownloadFile(m, &serviceDownloadFileServer{stream})
}

type Service_DownloadFileServer interface {
	Send(*httpbody.HttpBody) error
	grpc.ServerStream
}

type serviceDownloadFileServer struct {
	grpc.ServerStream
}

func (x *serviceDownloadFileServer) Send(m *httpbody.HttpBody) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).UploadFile(&serviceUploadFileServer{stream})
}

type Service_UploadFileServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*httpbody.HttpBody, error)
	grpc.ServerStream
}

type serviceUploadFileServer struct {
	grpc.ServerStream
}

func (x *serviceUploadFileServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceUploadFileServer) Recv() (*httpbody.HttpBody, error) {
	m := new(httpbody.HttpBody)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Service_UploadMultipleFiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).UploadMultipleFiles(&serviceUploadMultipleFilesServer{stream})
}

type Service_UploadMultipleFilesServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*httpbody.HttpBody, error)
	grpc.ServerStream
}

type serviceUploadMultipleFilesServer struct {
	grpc.ServerStream
}

func (x *serviceUploadMultipleFilesServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceUploadMultipleFilesServer) Recv() (*httpbody.HttpBody, error) {
	m := new(httpbody.HttpBody)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Service_UploadToAnotherService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).UploadToAnotherService(&serviceUploadToAnotherServiceServer{stream})
}

type Service_UploadToAnotherServiceServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*httpbody.HttpBody, error)
	grpc.ServerStream
}

type serviceUploadToAnotherServiceServer struct {
	grpc.ServerStream
}

func (x *serviceUploadToAnotherServiceServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceUploadToAnotherServiceServer) Recv() (*httpbody.HttpBody, error) {
	m := new(httpbody.HttpBody)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "collision.servicepb.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadFile",
			Handler:       _Service_DownloadFile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadFile",
			Handler:       _Service_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UploadMultipleFiles",
			Handler:       _Service_UploadMultipleFiles_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UploadToAnotherService",
			Handler:       _Service_UploadToAnotherService_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
