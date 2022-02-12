// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package appctlpb

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

// ClientLifecycleServiceClient is the client API for ClientLifecycleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientLifecycleServiceClient interface {
	// Fetch client application status.
	GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AppStatusMsg, error)
	// Quit client daemon.
	Exit(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Generate a thread dump of client daemon.
	GetThreadDump(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ThreadDump, error)
	// Start CPU profiling.
	StartCPUProfile(ctx context.Context, in *ProfileSavePath, opts ...grpc.CallOption) (*Empty, error)
	// Stop CPU profiling.
	StopCPUProfile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Generate a heap profile.
	GetHeapProfile(ctx context.Context, in *ProfileSavePath, opts ...grpc.CallOption) (*Empty, error)
}

type clientLifecycleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientLifecycleServiceClient(cc grpc.ClientConnInterface) ClientLifecycleServiceClient {
	return &clientLifecycleServiceClient{cc}
}

func (c *clientLifecycleServiceClient) GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AppStatusMsg, error) {
	out := new(AppStatusMsg)
	err := c.cc.Invoke(ctx, "/appctl.ClientLifecycleService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientLifecycleServiceClient) Exit(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/appctl.ClientLifecycleService/Exit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientLifecycleServiceClient) GetThreadDump(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ThreadDump, error) {
	out := new(ThreadDump)
	err := c.cc.Invoke(ctx, "/appctl.ClientLifecycleService/GetThreadDump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientLifecycleServiceClient) StartCPUProfile(ctx context.Context, in *ProfileSavePath, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/appctl.ClientLifecycleService/StartCPUProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientLifecycleServiceClient) StopCPUProfile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/appctl.ClientLifecycleService/StopCPUProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientLifecycleServiceClient) GetHeapProfile(ctx context.Context, in *ProfileSavePath, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/appctl.ClientLifecycleService/GetHeapProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientLifecycleServiceServer is the server API for ClientLifecycleService service.
// All implementations must embed UnimplementedClientLifecycleServiceServer
// for forward compatibility
type ClientLifecycleServiceServer interface {
	// Fetch client application status.
	GetStatus(context.Context, *Empty) (*AppStatusMsg, error)
	// Quit client daemon.
	Exit(context.Context, *Empty) (*Empty, error)
	// Generate a thread dump of client daemon.
	GetThreadDump(context.Context, *Empty) (*ThreadDump, error)
	// Start CPU profiling.
	StartCPUProfile(context.Context, *ProfileSavePath) (*Empty, error)
	// Stop CPU profiling.
	StopCPUProfile(context.Context, *Empty) (*Empty, error)
	// Generate a heap profile.
	GetHeapProfile(context.Context, *ProfileSavePath) (*Empty, error)
	mustEmbedUnimplementedClientLifecycleServiceServer()
}

// UnimplementedClientLifecycleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientLifecycleServiceServer struct {
}

func (UnimplementedClientLifecycleServiceServer) GetStatus(context.Context, *Empty) (*AppStatusMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedClientLifecycleServiceServer) Exit(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exit not implemented")
}
func (UnimplementedClientLifecycleServiceServer) GetThreadDump(context.Context, *Empty) (*ThreadDump, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThreadDump not implemented")
}
func (UnimplementedClientLifecycleServiceServer) StartCPUProfile(context.Context, *ProfileSavePath) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartCPUProfile not implemented")
}
func (UnimplementedClientLifecycleServiceServer) StopCPUProfile(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopCPUProfile not implemented")
}
func (UnimplementedClientLifecycleServiceServer) GetHeapProfile(context.Context, *ProfileSavePath) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHeapProfile not implemented")
}
func (UnimplementedClientLifecycleServiceServer) mustEmbedUnimplementedClientLifecycleServiceServer() {
}

// UnsafeClientLifecycleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientLifecycleServiceServer will
// result in compilation errors.
type UnsafeClientLifecycleServiceServer interface {
	mustEmbedUnimplementedClientLifecycleServiceServer()
}

func RegisterClientLifecycleServiceServer(s grpc.ServiceRegistrar, srv ClientLifecycleServiceServer) {
	s.RegisterService(&ClientLifecycleService_ServiceDesc, srv)
}

func _ClientLifecycleService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientLifecycleServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ClientLifecycleService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientLifecycleServiceServer).GetStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientLifecycleService_Exit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientLifecycleServiceServer).Exit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ClientLifecycleService/Exit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientLifecycleServiceServer).Exit(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientLifecycleService_GetThreadDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientLifecycleServiceServer).GetThreadDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ClientLifecycleService/GetThreadDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientLifecycleServiceServer).GetThreadDump(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientLifecycleService_StartCPUProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileSavePath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientLifecycleServiceServer).StartCPUProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ClientLifecycleService/StartCPUProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientLifecycleServiceServer).StartCPUProfile(ctx, req.(*ProfileSavePath))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientLifecycleService_StopCPUProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientLifecycleServiceServer).StopCPUProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ClientLifecycleService/StopCPUProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientLifecycleServiceServer).StopCPUProfile(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientLifecycleService_GetHeapProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileSavePath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientLifecycleServiceServer).GetHeapProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ClientLifecycleService/GetHeapProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientLifecycleServiceServer).GetHeapProfile(ctx, req.(*ProfileSavePath))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientLifecycleService_ServiceDesc is the grpc.ServiceDesc for ClientLifecycleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientLifecycleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appctl.ClientLifecycleService",
	HandlerType: (*ClientLifecycleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _ClientLifecycleService_GetStatus_Handler,
		},
		{
			MethodName: "Exit",
			Handler:    _ClientLifecycleService_Exit_Handler,
		},
		{
			MethodName: "GetThreadDump",
			Handler:    _ClientLifecycleService_GetThreadDump_Handler,
		},
		{
			MethodName: "StartCPUProfile",
			Handler:    _ClientLifecycleService_StartCPUProfile_Handler,
		},
		{
			MethodName: "StopCPUProfile",
			Handler:    _ClientLifecycleService_StopCPUProfile_Handler,
		},
		{
			MethodName: "GetHeapProfile",
			Handler:    _ClientLifecycleService_GetHeapProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lifecycle.proto",
}

// ServerLifecycleServiceClient is the client API for ServerLifecycleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerLifecycleServiceClient interface {
	// Fetch server application status.
	GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AppStatusMsg, error)
	// Start proxy in server application.
	Start(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Stop proxy in server application.
	Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Quit server daemon.
	Exit(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Generate a thread dump of server daemon.
	GetThreadDump(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ThreadDump, error)
	// Start CPU profiling.
	StartCPUProfile(ctx context.Context, in *ProfileSavePath, opts ...grpc.CallOption) (*Empty, error)
	// Stop CPU profiling.
	StopCPUProfile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Generate a heap profile.
	GetHeapProfile(ctx context.Context, in *ProfileSavePath, opts ...grpc.CallOption) (*Empty, error)
}

type serverLifecycleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerLifecycleServiceClient(cc grpc.ClientConnInterface) ServerLifecycleServiceClient {
	return &serverLifecycleServiceClient{cc}
}

func (c *serverLifecycleServiceClient) GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AppStatusMsg, error) {
	out := new(AppStatusMsg)
	err := c.cc.Invoke(ctx, "/appctl.ServerLifecycleService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverLifecycleServiceClient) Start(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/appctl.ServerLifecycleService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverLifecycleServiceClient) Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/appctl.ServerLifecycleService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverLifecycleServiceClient) Exit(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/appctl.ServerLifecycleService/Exit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverLifecycleServiceClient) GetThreadDump(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ThreadDump, error) {
	out := new(ThreadDump)
	err := c.cc.Invoke(ctx, "/appctl.ServerLifecycleService/GetThreadDump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverLifecycleServiceClient) StartCPUProfile(ctx context.Context, in *ProfileSavePath, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/appctl.ServerLifecycleService/StartCPUProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverLifecycleServiceClient) StopCPUProfile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/appctl.ServerLifecycleService/StopCPUProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverLifecycleServiceClient) GetHeapProfile(ctx context.Context, in *ProfileSavePath, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/appctl.ServerLifecycleService/GetHeapProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerLifecycleServiceServer is the server API for ServerLifecycleService service.
// All implementations must embed UnimplementedServerLifecycleServiceServer
// for forward compatibility
type ServerLifecycleServiceServer interface {
	// Fetch server application status.
	GetStatus(context.Context, *Empty) (*AppStatusMsg, error)
	// Start proxy in server application.
	Start(context.Context, *Empty) (*Empty, error)
	// Stop proxy in server application.
	Stop(context.Context, *Empty) (*Empty, error)
	// Quit server daemon.
	Exit(context.Context, *Empty) (*Empty, error)
	// Generate a thread dump of server daemon.
	GetThreadDump(context.Context, *Empty) (*ThreadDump, error)
	// Start CPU profiling.
	StartCPUProfile(context.Context, *ProfileSavePath) (*Empty, error)
	// Stop CPU profiling.
	StopCPUProfile(context.Context, *Empty) (*Empty, error)
	// Generate a heap profile.
	GetHeapProfile(context.Context, *ProfileSavePath) (*Empty, error)
	mustEmbedUnimplementedServerLifecycleServiceServer()
}

// UnimplementedServerLifecycleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServerLifecycleServiceServer struct {
}

func (UnimplementedServerLifecycleServiceServer) GetStatus(context.Context, *Empty) (*AppStatusMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedServerLifecycleServiceServer) Start(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedServerLifecycleServiceServer) Stop(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedServerLifecycleServiceServer) Exit(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exit not implemented")
}
func (UnimplementedServerLifecycleServiceServer) GetThreadDump(context.Context, *Empty) (*ThreadDump, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThreadDump not implemented")
}
func (UnimplementedServerLifecycleServiceServer) StartCPUProfile(context.Context, *ProfileSavePath) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartCPUProfile not implemented")
}
func (UnimplementedServerLifecycleServiceServer) StopCPUProfile(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopCPUProfile not implemented")
}
func (UnimplementedServerLifecycleServiceServer) GetHeapProfile(context.Context, *ProfileSavePath) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHeapProfile not implemented")
}
func (UnimplementedServerLifecycleServiceServer) mustEmbedUnimplementedServerLifecycleServiceServer() {
}

// UnsafeServerLifecycleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerLifecycleServiceServer will
// result in compilation errors.
type UnsafeServerLifecycleServiceServer interface {
	mustEmbedUnimplementedServerLifecycleServiceServer()
}

func RegisterServerLifecycleServiceServer(s grpc.ServiceRegistrar, srv ServerLifecycleServiceServer) {
	s.RegisterService(&ServerLifecycleService_ServiceDesc, srv)
}

func _ServerLifecycleService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerLifecycleServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ServerLifecycleService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerLifecycleServiceServer).GetStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerLifecycleService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerLifecycleServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ServerLifecycleService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerLifecycleServiceServer).Start(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerLifecycleService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerLifecycleServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ServerLifecycleService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerLifecycleServiceServer).Stop(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerLifecycleService_Exit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerLifecycleServiceServer).Exit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ServerLifecycleService/Exit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerLifecycleServiceServer).Exit(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerLifecycleService_GetThreadDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerLifecycleServiceServer).GetThreadDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ServerLifecycleService/GetThreadDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerLifecycleServiceServer).GetThreadDump(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerLifecycleService_StartCPUProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileSavePath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerLifecycleServiceServer).StartCPUProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ServerLifecycleService/StartCPUProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerLifecycleServiceServer).StartCPUProfile(ctx, req.(*ProfileSavePath))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerLifecycleService_StopCPUProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerLifecycleServiceServer).StopCPUProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ServerLifecycleService/StopCPUProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerLifecycleServiceServer).StopCPUProfile(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerLifecycleService_GetHeapProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileSavePath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerLifecycleServiceServer).GetHeapProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ServerLifecycleService/GetHeapProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerLifecycleServiceServer).GetHeapProfile(ctx, req.(*ProfileSavePath))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerLifecycleService_ServiceDesc is the grpc.ServiceDesc for ServerLifecycleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerLifecycleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appctl.ServerLifecycleService",
	HandlerType: (*ServerLifecycleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _ServerLifecycleService_GetStatus_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _ServerLifecycleService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _ServerLifecycleService_Stop_Handler,
		},
		{
			MethodName: "Exit",
			Handler:    _ServerLifecycleService_Exit_Handler,
		},
		{
			MethodName: "GetThreadDump",
			Handler:    _ServerLifecycleService_GetThreadDump_Handler,
		},
		{
			MethodName: "StartCPUProfile",
			Handler:    _ServerLifecycleService_StartCPUProfile_Handler,
		},
		{
			MethodName: "StopCPUProfile",
			Handler:    _ServerLifecycleService_StopCPUProfile_Handler,
		},
		{
			MethodName: "GetHeapProfile",
			Handler:    _ServerLifecycleService_GetHeapProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lifecycle.proto",
}
