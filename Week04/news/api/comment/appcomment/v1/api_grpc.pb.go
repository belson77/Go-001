// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package news_api_comment_appcomment_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AppCommentClient is the client API for AppComment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppCommentClient interface {
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type appCommentClient struct {
	cc grpc.ClientConnInterface
}

func NewAppCommentClient(cc grpc.ClientConnInterface) AppCommentClient {
	return &appCommentClient{cc}
}

func (c *appCommentClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error) {
	out := new(SubmitResponse)
	err := c.cc.Invoke(ctx, "/news.api.comment.appcomment.v1.AppComment/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appCommentClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/news.api.comment.appcomment.v1.AppComment/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppCommentServer is the server API for AppComment service.
// All implementations must embed UnimplementedAppCommentServer
// for forward compatibility
type AppCommentServer interface {
	Submit(context.Context, *SubmitRequest) (*SubmitResponse, error)
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	mustEmbedUnimplementedAppCommentServer()
}

// UnimplementedAppCommentServer must be embedded to have forward compatible implementations.
type UnimplementedAppCommentServer struct {
}

func (UnimplementedAppCommentServer) Submit(context.Context, *SubmitRequest) (*SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedAppCommentServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedAppCommentServer) mustEmbedUnimplementedAppCommentServer() {}

// UnsafeAppCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppCommentServer will
// result in compilation errors.
type UnsafeAppCommentServer interface {
	mustEmbedUnimplementedAppCommentServer()
}

func RegisterAppCommentServer(s grpc.ServiceRegistrar, srv AppCommentServer) {
	s.RegisterService(&_AppComment_serviceDesc, srv)
}

func _AppComment_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppCommentServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.api.comment.appcomment.v1.AppComment/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppCommentServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppComment_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppCommentServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.api.comment.appcomment.v1.AppComment/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppCommentServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppComment_serviceDesc = grpc.ServiceDesc{
	ServiceName: "news.api.comment.appcomment.v1.AppComment",
	HandlerType: (*AppCommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _AppComment_Submit_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _AppComment_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}