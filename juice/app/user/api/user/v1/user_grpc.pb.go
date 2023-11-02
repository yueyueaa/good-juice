// 定义项目 API 的 proto 文件 可以同时描述 gRPC 和 HTTP API
// protobuf 文件参考:
//  - https://developers.google.com/protocol-buffers/

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: app/user/api/user/v1/user.proto

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

const (
	UserBasic_UserLogin_FullMethodName      = "/v1.UserBasic/UserLogin"
	UserBasic_UserRegister_FullMethodName   = "/v1.UserBasic/UserRegister"
	UserBasic_UpdateUserInfo_FullMethodName = "/v1.UserBasic/UpdateUserInfo"
	UserBasic_GetUserInfo_FullMethodName    = "/v1.UserBasic/GetUserInfo"
	UserBasic_SearchUserList_FullMethodName = "/v1.UserBasic/SearchUserList"
)

// UserBasicClient is the client API for UserBasic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserBasicClient interface {
	// 用户登陆
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	// 用户注册
	UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
	// 用户信息更新接口
	UpdateUserInfo(ctx context.Context, in *UserUpdateInfoRequest, opts ...grpc.CallOption) (*UserUpdateInfoResponse, error)
	// 获取用户信息
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	// 搜索用户
	SearchUserList(ctx context.Context, in *SeaechUserListRequest, opts ...grpc.CallOption) (*UserListResponse, error)
}

type userBasicClient struct {
	cc grpc.ClientConnInterface
}

func NewUserBasicClient(cc grpc.ClientConnInterface) UserBasicClient {
	return &userBasicClient{cc}
}

func (c *userBasicClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, UserBasic_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userBasicClient) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	out := new(UserRegisterResponse)
	err := c.cc.Invoke(ctx, UserBasic_UserRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userBasicClient) UpdateUserInfo(ctx context.Context, in *UserUpdateInfoRequest, opts ...grpc.CallOption) (*UserUpdateInfoResponse, error) {
	out := new(UserUpdateInfoResponse)
	err := c.cc.Invoke(ctx, UserBasic_UpdateUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userBasicClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, UserBasic_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userBasicClient) SearchUserList(ctx context.Context, in *SeaechUserListRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, UserBasic_SearchUserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserBasicServer is the server API for UserBasic service.
// All implementations must embed UnimplementedUserBasicServer
// for forward compatibility
type UserBasicServer interface {
	// 用户登陆
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	// 用户注册
	UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error)
	// 用户信息更新接口
	UpdateUserInfo(context.Context, *UserUpdateInfoRequest) (*UserUpdateInfoResponse, error)
	// 获取用户信息
	GetUserInfo(context.Context, *GetUserInfoRequest) (*UserInfoResponse, error)
	// 搜索用户
	SearchUserList(context.Context, *SeaechUserListRequest) (*UserListResponse, error)
	mustEmbedUnimplementedUserBasicServer()
}

// UnimplementedUserBasicServer must be embedded to have forward compatible implementations.
type UnimplementedUserBasicServer struct {
}

func (UnimplementedUserBasicServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserBasicServer) UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedUserBasicServer) UpdateUserInfo(context.Context, *UserUpdateInfoRequest) (*UserUpdateInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUserBasicServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserBasicServer) SearchUserList(context.Context, *SeaechUserListRequest) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUserList not implemented")
}
func (UnimplementedUserBasicServer) mustEmbedUnimplementedUserBasicServer() {}

// UnsafeUserBasicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserBasicServer will
// result in compilation errors.
type UnsafeUserBasicServer interface {
	mustEmbedUnimplementedUserBasicServer()
}

func RegisterUserBasicServer(s grpc.ServiceRegistrar, srv UserBasicServer) {
	s.RegisterService(&UserBasic_ServiceDesc, srv)
}

func _UserBasic_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBasicServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserBasic_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBasicServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserBasic_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBasicServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserBasic_UserRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBasicServer).UserRegister(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserBasic_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBasicServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserBasic_UpdateUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBasicServer).UpdateUserInfo(ctx, req.(*UserUpdateInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserBasic_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBasicServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserBasic_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBasicServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserBasic_SearchUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeaechUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBasicServer).SearchUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserBasic_SearchUserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBasicServer).SearchUserList(ctx, req.(*SeaechUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserBasic_ServiceDesc is the grpc.ServiceDesc for UserBasic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserBasic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UserBasic",
	HandlerType: (*UserBasicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _UserBasic_UserLogin_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _UserBasic_UserRegister_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _UserBasic_UpdateUserInfo_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserBasic_GetUserInfo_Handler,
		},
		{
			MethodName: "SearchUserList",
			Handler:    _UserBasic_SearchUserList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/user/api/user/v1/user.proto",
}

const (
	UserRelation_SetFollow_FullMethodName       = "/v1.UserRelation/SetFollow"
	UserRelation_GetFollowedList_FullMethodName = "/v1.UserRelation/GetFollowedList"
	UserRelation_GetFollowerList_FullMethodName = "/v1.UserRelation/GetFollowerList"
)

// UserRelationClient is the client API for UserRelation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRelationClient interface {
	// 关注（互动）
	SetFollow(ctx context.Context, in *SetFollowRequest, opts ...grpc.CallOption) (*SetFollowResponse, error)
	// 关注列表
	GetFollowedList(ctx context.Context, in *GetFollowedListRequest, opts ...grpc.CallOption) (*FollowedListResponse, error)
	// 粉丝列表
	GetFollowerList(ctx context.Context, in *GetFollowerListRequest, opts ...grpc.CallOption) (*FollowerListResponse, error)
}

type userRelationClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRelationClient(cc grpc.ClientConnInterface) UserRelationClient {
	return &userRelationClient{cc}
}

func (c *userRelationClient) SetFollow(ctx context.Context, in *SetFollowRequest, opts ...grpc.CallOption) (*SetFollowResponse, error) {
	out := new(SetFollowResponse)
	err := c.cc.Invoke(ctx, UserRelation_SetFollow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationClient) GetFollowedList(ctx context.Context, in *GetFollowedListRequest, opts ...grpc.CallOption) (*FollowedListResponse, error) {
	out := new(FollowedListResponse)
	err := c.cc.Invoke(ctx, UserRelation_GetFollowedList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationClient) GetFollowerList(ctx context.Context, in *GetFollowerListRequest, opts ...grpc.CallOption) (*FollowerListResponse, error) {
	out := new(FollowerListResponse)
	err := c.cc.Invoke(ctx, UserRelation_GetFollowerList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRelationServer is the server API for UserRelation service.
// All implementations must embed UnimplementedUserRelationServer
// for forward compatibility
type UserRelationServer interface {
	// 关注（互动）
	SetFollow(context.Context, *SetFollowRequest) (*SetFollowResponse, error)
	// 关注列表
	GetFollowedList(context.Context, *GetFollowedListRequest) (*FollowedListResponse, error)
	// 粉丝列表
	GetFollowerList(context.Context, *GetFollowerListRequest) (*FollowerListResponse, error)
	mustEmbedUnimplementedUserRelationServer()
}

// UnimplementedUserRelationServer must be embedded to have forward compatible implementations.
type UnimplementedUserRelationServer struct {
}

func (UnimplementedUserRelationServer) SetFollow(context.Context, *SetFollowRequest) (*SetFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFollow not implemented")
}
func (UnimplementedUserRelationServer) GetFollowedList(context.Context, *GetFollowedListRequest) (*FollowedListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowedList not implemented")
}
func (UnimplementedUserRelationServer) GetFollowerList(context.Context, *GetFollowerListRequest) (*FollowerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowerList not implemented")
}
func (UnimplementedUserRelationServer) mustEmbedUnimplementedUserRelationServer() {}

// UnsafeUserRelationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRelationServer will
// result in compilation errors.
type UnsafeUserRelationServer interface {
	mustEmbedUnimplementedUserRelationServer()
}

func RegisterUserRelationServer(s grpc.ServiceRegistrar, srv UserRelationServer) {
	s.RegisterService(&UserRelation_ServiceDesc, srv)
}

func _UserRelation_SetFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServer).SetFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelation_SetFollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServer).SetFollow(ctx, req.(*SetFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelation_GetFollowedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowedListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServer).GetFollowedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelation_GetFollowedList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServer).GetFollowedList(ctx, req.(*GetFollowedListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelation_GetFollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServer).GetFollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelation_GetFollowerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServer).GetFollowerList(ctx, req.(*GetFollowerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRelation_ServiceDesc is the grpc.ServiceDesc for UserRelation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRelation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UserRelation",
	HandlerType: (*UserRelationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetFollow",
			Handler:    _UserRelation_SetFollow_Handler,
		},
		{
			MethodName: "GetFollowedList",
			Handler:    _UserRelation_GetFollowedList_Handler,
		},
		{
			MethodName: "GetFollowerList",
			Handler:    _UserRelation_GetFollowerList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/user/api/user/v1/user.proto",
}
