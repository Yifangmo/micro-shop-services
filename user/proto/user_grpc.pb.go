// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: user/proto/user.proto

package proto

import (
	context "context"
	common "github.com/Yifangmo/micro-shop-services/common"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// 用户个人信息
	GetUserList(ctx context.Context, in *common.PageInfo, opts ...grpc.CallOption) (*UserListResponse, error)
	GetUserByMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	GetUserById(ctx context.Context, in *UserIDRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*IDResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CheckPassWord(ctx context.Context, in *CheckPasswordRequest, opts ...grpc.CallOption) (*CheckPasswordResponse, error)
	// 收货人信息
	GetConsigneeAddressList(ctx context.Context, in *ConsigneeAddressRequest, opts ...grpc.CallOption) (*ConsigneeAddressListResponse, error)
	CreateConsigneeAddress(ctx context.Context, in *ConsigneeAddressRequest, opts ...grpc.CallOption) (*IDResponse, error)
	DeleteConsigneeAddress(ctx context.Context, in *ConsigneeAddressRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateConsigneeAddress(ctx context.Context, in *ConsigneeAddressRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// 用户留言
	GetUserMessageList(ctx context.Context, in *UserMessageRequest, opts ...grpc.CallOption) (*UserMessageListResponse, error)
	CreateUserMessage(ctx context.Context, in *UserMessageRequest, opts ...grpc.CallOption) (*IDResponse, error)
	// 用户收藏
	GetUserFavList(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*UserFavListResponse, error)
	AddUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetUserFavDetail(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUserList(ctx context.Context, in *common.PageInfo, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserByMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserByMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserById(ctx context.Context, in *UserIDRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*IDResponse, error) {
	out := new(IDResponse)
	err := c.cc.Invoke(ctx, "/User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CheckPassWord(ctx context.Context, in *CheckPasswordRequest, opts ...grpc.CallOption) (*CheckPasswordResponse, error) {
	out := new(CheckPasswordResponse)
	err := c.cc.Invoke(ctx, "/User/CheckPassWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetConsigneeAddressList(ctx context.Context, in *ConsigneeAddressRequest, opts ...grpc.CallOption) (*ConsigneeAddressListResponse, error) {
	out := new(ConsigneeAddressListResponse)
	err := c.cc.Invoke(ctx, "/User/GetConsigneeAddressList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateConsigneeAddress(ctx context.Context, in *ConsigneeAddressRequest, opts ...grpc.CallOption) (*IDResponse, error) {
	out := new(IDResponse)
	err := c.cc.Invoke(ctx, "/User/CreateConsigneeAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteConsigneeAddress(ctx context.Context, in *ConsigneeAddressRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/User/DeleteConsigneeAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateConsigneeAddress(ctx context.Context, in *ConsigneeAddressRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/User/UpdateConsigneeAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserMessageList(ctx context.Context, in *UserMessageRequest, opts ...grpc.CallOption) (*UserMessageListResponse, error) {
	out := new(UserMessageListResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserMessageList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUserMessage(ctx context.Context, in *UserMessageRequest, opts ...grpc.CallOption) (*IDResponse, error) {
	out := new(IDResponse)
	err := c.cc.Invoke(ctx, "/User/CreateUserMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserFavList(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*UserFavListResponse, error) {
	out := new(UserFavListResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserFavList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/User/AddUserFav", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/User/DeleteUserFav", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserFavDetail(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/User/GetUserFavDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	// 用户个人信息
	GetUserList(context.Context, *common.PageInfo) (*UserListResponse, error)
	GetUserByMobile(context.Context, *MobileRequest) (*UserInfoResponse, error)
	GetUserById(context.Context, *UserIDRequest) (*UserInfoResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*IDResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*empty.Empty, error)
	CheckPassWord(context.Context, *CheckPasswordRequest) (*CheckPasswordResponse, error)
	// 收货人信息
	GetConsigneeAddressList(context.Context, *ConsigneeAddressRequest) (*ConsigneeAddressListResponse, error)
	CreateConsigneeAddress(context.Context, *ConsigneeAddressRequest) (*IDResponse, error)
	DeleteConsigneeAddress(context.Context, *ConsigneeAddressRequest) (*empty.Empty, error)
	UpdateConsigneeAddress(context.Context, *ConsigneeAddressRequest) (*empty.Empty, error)
	// 用户留言
	GetUserMessageList(context.Context, *UserMessageRequest) (*UserMessageListResponse, error)
	CreateUserMessage(context.Context, *UserMessageRequest) (*IDResponse, error)
	// 用户收藏
	GetUserFavList(context.Context, *UserFavRequest) (*UserFavListResponse, error)
	AddUserFav(context.Context, *UserFavRequest) (*empty.Empty, error)
	DeleteUserFav(context.Context, *UserFavRequest) (*empty.Empty, error)
	GetUserFavDetail(context.Context, *UserFavRequest) (*empty.Empty, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) GetUserList(context.Context, *common.PageInfo) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedUserServer) GetUserByMobile(context.Context, *MobileRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByMobile not implemented")
}
func (UnimplementedUserServer) GetUserById(context.Context, *UserIDRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServer) CreateUser(context.Context, *CreateUserRequest) (*IDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServer) UpdateUser(context.Context, *UpdateUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServer) CheckPassWord(context.Context, *CheckPasswordRequest) (*CheckPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPassWord not implemented")
}
func (UnimplementedUserServer) GetConsigneeAddressList(context.Context, *ConsigneeAddressRequest) (*ConsigneeAddressListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsigneeAddressList not implemented")
}
func (UnimplementedUserServer) CreateConsigneeAddress(context.Context, *ConsigneeAddressRequest) (*IDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConsigneeAddress not implemented")
}
func (UnimplementedUserServer) DeleteConsigneeAddress(context.Context, *ConsigneeAddressRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConsigneeAddress not implemented")
}
func (UnimplementedUserServer) UpdateConsigneeAddress(context.Context, *ConsigneeAddressRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConsigneeAddress not implemented")
}
func (UnimplementedUserServer) GetUserMessageList(context.Context, *UserMessageRequest) (*UserMessageListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMessageList not implemented")
}
func (UnimplementedUserServer) CreateUserMessage(context.Context, *UserMessageRequest) (*IDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserMessage not implemented")
}
func (UnimplementedUserServer) GetUserFavList(context.Context, *UserFavRequest) (*UserFavListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFavList not implemented")
}
func (UnimplementedUserServer) AddUserFav(context.Context, *UserFavRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserFav not implemented")
}
func (UnimplementedUserServer) DeleteUserFav(context.Context, *UserFavRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserFav not implemented")
}
func (UnimplementedUserServer) GetUserFavDetail(context.Context, *UserFavRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFavDetail not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserList(ctx, req.(*common.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserByMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserByMobile(ctx, req.(*MobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserById(ctx, req.(*UserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CheckPassWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckPassWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CheckPassWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CheckPassWord(ctx, req.(*CheckPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetConsigneeAddressList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsigneeAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetConsigneeAddressList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetConsigneeAddressList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetConsigneeAddressList(ctx, req.(*ConsigneeAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateConsigneeAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsigneeAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateConsigneeAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CreateConsigneeAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateConsigneeAddress(ctx, req.(*ConsigneeAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteConsigneeAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsigneeAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteConsigneeAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/DeleteConsigneeAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteConsigneeAddress(ctx, req.(*ConsigneeAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateConsigneeAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsigneeAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateConsigneeAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdateConsigneeAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateConsigneeAddress(ctx, req.(*ConsigneeAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserMessageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserMessageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserMessageList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserMessageList(ctx, req.(*UserMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CreateUserMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUserMessage(ctx, req.(*UserMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserFavList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserFavList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserFavList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserFavList(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddUserFav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddUserFav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/AddUserFav",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddUserFav(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUserFav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUserFav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/DeleteUserFav",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUserFav(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserFavDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserFavDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserFavDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserFavDetail(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserList",
			Handler:    _User_GetUserList_Handler,
		},
		{
			MethodName: "GetUserByMobile",
			Handler:    _User_GetUserByMobile_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _User_GetUserById_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "CheckPassWord",
			Handler:    _User_CheckPassWord_Handler,
		},
		{
			MethodName: "GetConsigneeAddressList",
			Handler:    _User_GetConsigneeAddressList_Handler,
		},
		{
			MethodName: "CreateConsigneeAddress",
			Handler:    _User_CreateConsigneeAddress_Handler,
		},
		{
			MethodName: "DeleteConsigneeAddress",
			Handler:    _User_DeleteConsigneeAddress_Handler,
		},
		{
			MethodName: "UpdateConsigneeAddress",
			Handler:    _User_UpdateConsigneeAddress_Handler,
		},
		{
			MethodName: "GetUserMessageList",
			Handler:    _User_GetUserMessageList_Handler,
		},
		{
			MethodName: "CreateUserMessage",
			Handler:    _User_CreateUserMessage_Handler,
		},
		{
			MethodName: "GetUserFavList",
			Handler:    _User_GetUserFavList_Handler,
		},
		{
			MethodName: "AddUserFav",
			Handler:    _User_AddUserFav_Handler,
		},
		{
			MethodName: "DeleteUserFav",
			Handler:    _User_DeleteUserFav_Handler,
		},
		{
			MethodName: "GetUserFavDetail",
			Handler:    _User_GetUserFavDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/proto/user.proto",
}
