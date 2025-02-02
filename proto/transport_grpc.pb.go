// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RaftTransportClient is the client API for RaftTransport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaftTransportClient interface {
	// AppendEntriesPipeline opens an AppendEntries message stream.
	AppendEntriesPipeline(ctx context.Context, opts ...grpc.CallOption) (RaftTransport_AppendEntriesPipelineClient, error)
	// AppendEntries performs a single append entries request / response.
	AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error)
	// RequestVote is the command used by a candidate to ask a Raft peer for a vote in an election.
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error)
	// TimeoutNow is used to start a leadership transfer to the target node.
	TimeoutNow(ctx context.Context, in *TimeoutNowRequest, opts ...grpc.CallOption) (*TimeoutNowResponse, error)
	// InstallSnapshot is the command sent to a Raft peer to bootstrap its log (and state machine) from a snapshot on another peer.
	InstallSnapshot(ctx context.Context, opts ...grpc.CallOption) (RaftTransport_InstallSnapshotClient, error)
}

type raftTransportClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftTransportClient(cc grpc.ClientConnInterface) RaftTransportClient {
	return &raftTransportClient{cc}
}

func (c *raftTransportClient) AppendEntriesPipeline(ctx context.Context, opts ...grpc.CallOption) (RaftTransport_AppendEntriesPipelineClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RaftTransport_serviceDesc.Streams[0], "/RaftTransport/AppendEntriesPipeline", opts...)
	if err != nil {
		return nil, err
	}
	x := &raftTransportAppendEntriesPipelineClient{stream}
	return x, nil
}

type RaftTransport_AppendEntriesPipelineClient interface {
	Send(*AppendEntriesRequest) error
	Recv() (*AppendEntriesResponse, error)
	grpc.ClientStream
}

type raftTransportAppendEntriesPipelineClient struct {
	grpc.ClientStream
}

func (x *raftTransportAppendEntriesPipelineClient) Send(m *AppendEntriesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *raftTransportAppendEntriesPipelineClient) Recv() (*AppendEntriesResponse, error) {
	m := new(AppendEntriesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *raftTransportClient) AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error) {
	out := new(AppendEntriesResponse)
	err := c.cc.Invoke(ctx, "/RaftTransport/AppendEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftTransportClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error) {
	out := new(RequestVoteResponse)
	err := c.cc.Invoke(ctx, "/RaftTransport/RequestVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftTransportClient) TimeoutNow(ctx context.Context, in *TimeoutNowRequest, opts ...grpc.CallOption) (*TimeoutNowResponse, error) {
	out := new(TimeoutNowResponse)
	err := c.cc.Invoke(ctx, "/RaftTransport/TimeoutNow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftTransportClient) InstallSnapshot(ctx context.Context, opts ...grpc.CallOption) (RaftTransport_InstallSnapshotClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RaftTransport_serviceDesc.Streams[1], "/RaftTransport/InstallSnapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &raftTransportInstallSnapshotClient{stream}
	return x, nil
}

type RaftTransport_InstallSnapshotClient interface {
	Send(*InstallSnapshotRequest) error
	CloseAndRecv() (*InstallSnapshotResponse, error)
	grpc.ClientStream
}

type raftTransportInstallSnapshotClient struct {
	grpc.ClientStream
}

func (x *raftTransportInstallSnapshotClient) Send(m *InstallSnapshotRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *raftTransportInstallSnapshotClient) CloseAndRecv() (*InstallSnapshotResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(InstallSnapshotResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RaftTransportServer is the server API for RaftTransport service.
// All implementations must embed UnimplementedRaftTransportServer
// for forward compatibility
type RaftTransportServer interface {
	// AppendEntriesPipeline opens an AppendEntries message stream.
	AppendEntriesPipeline(RaftTransport_AppendEntriesPipelineServer) error
	// AppendEntries performs a single append entries request / response.
	AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error)
	// RequestVote is the command used by a candidate to ask a Raft peer for a vote in an election.
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error)
	// TimeoutNow is used to start a leadership transfer to the target node.
	TimeoutNow(context.Context, *TimeoutNowRequest) (*TimeoutNowResponse, error)
	// InstallSnapshot is the command sent to a Raft peer to bootstrap its log (and state machine) from a snapshot on another peer.
	InstallSnapshot(RaftTransport_InstallSnapshotServer) error
	mustEmbedUnimplementedRaftTransportServer()
}

// UnimplementedRaftTransportServer must be embedded to have forward compatible implementations.
type UnimplementedRaftTransportServer struct {
}

func (UnimplementedRaftTransportServer) AppendEntriesPipeline(RaftTransport_AppendEntriesPipelineServer) error {
	return status.Errorf(codes.Unimplemented, "method AppendEntriesPipeline not implemented")
}
func (UnimplementedRaftTransportServer) AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendEntries not implemented")
}
func (UnimplementedRaftTransportServer) RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVote not implemented")
}
func (UnimplementedRaftTransportServer) TimeoutNow(context.Context, *TimeoutNowRequest) (*TimeoutNowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TimeoutNow not implemented")
}
func (UnimplementedRaftTransportServer) InstallSnapshot(RaftTransport_InstallSnapshotServer) error {
	return status.Errorf(codes.Unimplemented, "method InstallSnapshot not implemented")
}
func (UnimplementedRaftTransportServer) mustEmbedUnimplementedRaftTransportServer() {}

// UnsafeRaftTransportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaftTransportServer will
// result in compilation errors.
type UnsafeRaftTransportServer interface {
	mustEmbedUnimplementedRaftTransportServer()
}

func RegisterRaftTransportServer(s grpc.ServiceRegistrar, srv RaftTransportServer) {
	s.RegisterService(&_RaftTransport_serviceDesc, srv)
}

func _RaftTransport_AppendEntriesPipeline_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RaftTransportServer).AppendEntriesPipeline(&raftTransportAppendEntriesPipelineServer{stream})
}

type RaftTransport_AppendEntriesPipelineServer interface {
	Send(*AppendEntriesResponse) error
	Recv() (*AppendEntriesRequest, error)
	grpc.ServerStream
}

type raftTransportAppendEntriesPipelineServer struct {
	grpc.ServerStream
}

func (x *raftTransportAppendEntriesPipelineServer) Send(m *AppendEntriesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *raftTransportAppendEntriesPipelineServer) Recv() (*AppendEntriesRequest, error) {
	m := new(AppendEntriesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RaftTransport_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftTransportServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RaftTransport/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftTransportServer).AppendEntries(ctx, req.(*AppendEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftTransport_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftTransportServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RaftTransport/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftTransportServer).RequestVote(ctx, req.(*RequestVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftTransport_TimeoutNow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeoutNowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftTransportServer).TimeoutNow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RaftTransport/TimeoutNow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftTransportServer).TimeoutNow(ctx, req.(*TimeoutNowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftTransport_InstallSnapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RaftTransportServer).InstallSnapshot(&raftTransportInstallSnapshotServer{stream})
}

type RaftTransport_InstallSnapshotServer interface {
	SendAndClose(*InstallSnapshotResponse) error
	Recv() (*InstallSnapshotRequest, error)
	grpc.ServerStream
}

type raftTransportInstallSnapshotServer struct {
	grpc.ServerStream
}

func (x *raftTransportInstallSnapshotServer) SendAndClose(m *InstallSnapshotResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *raftTransportInstallSnapshotServer) Recv() (*InstallSnapshotRequest, error) {
	m := new(InstallSnapshotRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RaftTransport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RaftTransport",
	HandlerType: (*RaftTransportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppendEntries",
			Handler:    _RaftTransport_AppendEntries_Handler,
		},
		{
			MethodName: "RequestVote",
			Handler:    _RaftTransport_RequestVote_Handler,
		},
		{
			MethodName: "TimeoutNow",
			Handler:    _RaftTransport_TimeoutNow_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AppendEntriesPipeline",
			Handler:       _RaftTransport_AppendEntriesPipeline_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "InstallSnapshot",
			Handler:       _RaftTransport_InstallSnapshot_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "transport.proto",
}
