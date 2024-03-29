// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: userinfogateway.proto

package gateway_api

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

const (
	UserInfoService_SignUpUserInfo_FullMethodName = "/api.UserInfoService/SignUpUserInfo"
	UserInfoService_GetUserInfo_FullMethodName    = "/api.UserInfoService/GetUserInfo"
	UserInfoService_GetUserList_FullMethodName    = "/api.UserInfoService/GetUserList"
)

// UserInfoServiceClient is the client API for UserInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInfoServiceClient interface {
	SignUpUserInfo(ctx context.Context, in *SignUpUserInfoRequest, opts ...grpc.CallOption) (*SignUpUserInfoResponse, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error)
}

type userInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInfoServiceClient(cc grpc.ClientConnInterface) UserInfoServiceClient {
	return &userInfoServiceClient{cc}
}

func (c *userInfoServiceClient) SignUpUserInfo(ctx context.Context, in *SignUpUserInfoRequest, opts ...grpc.CallOption) (*SignUpUserInfoResponse, error) {
	out := new(SignUpUserInfoResponse)
	err := c.cc.Invoke(ctx, UserInfoService_SignUpUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoServiceClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, UserInfoService_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoServiceClient) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error) {
	out := new(GetUserListResponse)
	err := c.cc.Invoke(ctx, UserInfoService_GetUserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoServiceServer is the server API for UserInfoService service.
// All implementations must embed UnimplementedUserInfoServiceServer
// for forward compatibility
type UserInfoServiceServer interface {
	SignUpUserInfo(context.Context, *SignUpUserInfoRequest) (*SignUpUserInfoResponse, error)
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error)
	mustEmbedUnimplementedUserInfoServiceServer()
}

// UnimplementedUserInfoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserInfoServiceServer struct {
}

func (UnimplementedUserInfoServiceServer) SignUpUserInfo(context.Context, *SignUpUserInfoRequest) (*SignUpUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUpUserInfo not implemented")
}
func (UnimplementedUserInfoServiceServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserInfoServiceServer) GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedUserInfoServiceServer) mustEmbedUnimplementedUserInfoServiceServer() {}

// UnsafeUserInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInfoServiceServer will
// result in compilation errors.
type UnsafeUserInfoServiceServer interface {
	mustEmbedUnimplementedUserInfoServiceServer()
}

func RegisterUserInfoServiceServer(s grpc.ServiceRegistrar, srv UserInfoServiceServer) {
	s.RegisterService(&UserInfoService_ServiceDesc, srv)
}

func _UserInfoService_SignUpUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServiceServer).SignUpUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInfoService_SignUpUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServiceServer).SignUpUserInfo(ctx, req.(*SignUpUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfoService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInfoService_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServiceServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfoService_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServiceServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInfoService_GetUserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServiceServer).GetUserList(ctx, req.(*GetUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInfoService_ServiceDesc is the grpc.ServiceDesc for UserInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserInfoService",
	HandlerType: (*UserInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUpUserInfo",
			Handler:    _UserInfoService_SignUpUserInfo_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserInfoService_GetUserInfo_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _UserInfoService_GetUserList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userinfogateway.proto",
}
