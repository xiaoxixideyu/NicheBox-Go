// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.1
// source: user.proto

package user

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
	User_Register_FullMethodName               = "/userclient.User/Register"
	User_CheckEmail_FullMethodName             = "/userclient.User/CheckEmail"
	User_GetUidByEmailAndPwd_FullMethodName    = "/userclient.User/GetUidByEmailAndPwd"
	User_CheckUid_FullMethodName               = "/userclient.User/CheckUid"
	User_SetUserBaseInfo_FullMethodName        = "/userclient.User/SetUserBaseInfo"
	User_GetUserBaseInfo_FullMethodName        = "/userclient.User/GetUserBaseInfo"
	User_SetVerificationCode_FullMethodName    = "/userclient.User/SetVerificationCode"
	User_GetVerificationCode_FullMethodName    = "/userclient.User/GetVerificationCode"
	User_RemoveVerificationCode_FullMethodName = "/userclient.User/RemoveVerificationCode"
	User_ForgetPassword_FullMethodName         = "/userclient.User/ForgetPassword"
	User_SetCriticalUserInfo_FullMethodName    = "/userclient.User/SetCriticalUserInfo"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	CheckEmail(ctx context.Context, in *CheckEmailRequest, opts ...grpc.CallOption) (*CheckEmailResponse, error)
	GetUidByEmailAndPwd(ctx context.Context, in *GetUidByEmailAndPwdRequest, opts ...grpc.CallOption) (*GetUidByEmailAndPwdResponse, error)
	CheckUid(ctx context.Context, in *CheckUidRequest, opts ...grpc.CallOption) (*CheckUidResponse, error)
	SetUserBaseInfo(ctx context.Context, in *SetUserBaseInfoRequest, opts ...grpc.CallOption) (*SetUserBaseInfoResponse, error)
	GetUserBaseInfo(ctx context.Context, in *GetUserBaseInfoRequest, opts ...grpc.CallOption) (*GetUserBaseInfoResponse, error)
	SetVerificationCode(ctx context.Context, in *SetVerificationCodeRequest, opts ...grpc.CallOption) (*SetVerificationCodeResponse, error)
	GetVerificationCode(ctx context.Context, in *GetVerificationCodeRequest, opts ...grpc.CallOption) (*GetVerificationCodeResponse, error)
	RemoveVerificationCode(ctx context.Context, in *RemoveVerificationCodeRequest, opts ...grpc.CallOption) (*RemoveVerificationCodeResponse, error)
	ForgetPassword(ctx context.Context, in *ForgetPasswordRequest, opts ...grpc.CallOption) (*ForgetPasswordResponse, error)
	SetCriticalUserInfo(ctx context.Context, in *SetCriticalUserInfoRequest, opts ...grpc.CallOption) (*SetCriticalUserInfoResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, User_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CheckEmail(ctx context.Context, in *CheckEmailRequest, opts ...grpc.CallOption) (*CheckEmailResponse, error) {
	out := new(CheckEmailResponse)
	err := c.cc.Invoke(ctx, User_CheckEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUidByEmailAndPwd(ctx context.Context, in *GetUidByEmailAndPwdRequest, opts ...grpc.CallOption) (*GetUidByEmailAndPwdResponse, error) {
	out := new(GetUidByEmailAndPwdResponse)
	err := c.cc.Invoke(ctx, User_GetUidByEmailAndPwd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CheckUid(ctx context.Context, in *CheckUidRequest, opts ...grpc.CallOption) (*CheckUidResponse, error) {
	out := new(CheckUidResponse)
	err := c.cc.Invoke(ctx, User_CheckUid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetUserBaseInfo(ctx context.Context, in *SetUserBaseInfoRequest, opts ...grpc.CallOption) (*SetUserBaseInfoResponse, error) {
	out := new(SetUserBaseInfoResponse)
	err := c.cc.Invoke(ctx, User_SetUserBaseInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserBaseInfo(ctx context.Context, in *GetUserBaseInfoRequest, opts ...grpc.CallOption) (*GetUserBaseInfoResponse, error) {
	out := new(GetUserBaseInfoResponse)
	err := c.cc.Invoke(ctx, User_GetUserBaseInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetVerificationCode(ctx context.Context, in *SetVerificationCodeRequest, opts ...grpc.CallOption) (*SetVerificationCodeResponse, error) {
	out := new(SetVerificationCodeResponse)
	err := c.cc.Invoke(ctx, User_SetVerificationCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetVerificationCode(ctx context.Context, in *GetVerificationCodeRequest, opts ...grpc.CallOption) (*GetVerificationCodeResponse, error) {
	out := new(GetVerificationCodeResponse)
	err := c.cc.Invoke(ctx, User_GetVerificationCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RemoveVerificationCode(ctx context.Context, in *RemoveVerificationCodeRequest, opts ...grpc.CallOption) (*RemoveVerificationCodeResponse, error) {
	out := new(RemoveVerificationCodeResponse)
	err := c.cc.Invoke(ctx, User_RemoveVerificationCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ForgetPassword(ctx context.Context, in *ForgetPasswordRequest, opts ...grpc.CallOption) (*ForgetPasswordResponse, error) {
	out := new(ForgetPasswordResponse)
	err := c.cc.Invoke(ctx, User_ForgetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetCriticalUserInfo(ctx context.Context, in *SetCriticalUserInfoRequest, opts ...grpc.CallOption) (*SetCriticalUserInfoResponse, error) {
	out := new(SetCriticalUserInfoResponse)
	err := c.cc.Invoke(ctx, User_SetCriticalUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	CheckEmail(context.Context, *CheckEmailRequest) (*CheckEmailResponse, error)
	GetUidByEmailAndPwd(context.Context, *GetUidByEmailAndPwdRequest) (*GetUidByEmailAndPwdResponse, error)
	CheckUid(context.Context, *CheckUidRequest) (*CheckUidResponse, error)
	SetUserBaseInfo(context.Context, *SetUserBaseInfoRequest) (*SetUserBaseInfoResponse, error)
	GetUserBaseInfo(context.Context, *GetUserBaseInfoRequest) (*GetUserBaseInfoResponse, error)
	SetVerificationCode(context.Context, *SetVerificationCodeRequest) (*SetVerificationCodeResponse, error)
	GetVerificationCode(context.Context, *GetVerificationCodeRequest) (*GetVerificationCodeResponse, error)
	RemoveVerificationCode(context.Context, *RemoveVerificationCodeRequest) (*RemoveVerificationCodeResponse, error)
	ForgetPassword(context.Context, *ForgetPasswordRequest) (*ForgetPasswordResponse, error)
	SetCriticalUserInfo(context.Context, *SetCriticalUserInfoRequest) (*SetCriticalUserInfoResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) CheckEmail(context.Context, *CheckEmailRequest) (*CheckEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckEmail not implemented")
}
func (UnimplementedUserServer) GetUidByEmailAndPwd(context.Context, *GetUidByEmailAndPwdRequest) (*GetUidByEmailAndPwdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUidByEmailAndPwd not implemented")
}
func (UnimplementedUserServer) CheckUid(context.Context, *CheckUidRequest) (*CheckUidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUid not implemented")
}
func (UnimplementedUserServer) SetUserBaseInfo(context.Context, *SetUserBaseInfoRequest) (*SetUserBaseInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserBaseInfo not implemented")
}
func (UnimplementedUserServer) GetUserBaseInfo(context.Context, *GetUserBaseInfoRequest) (*GetUserBaseInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBaseInfo not implemented")
}
func (UnimplementedUserServer) SetVerificationCode(context.Context, *SetVerificationCodeRequest) (*SetVerificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetVerificationCode not implemented")
}
func (UnimplementedUserServer) GetVerificationCode(context.Context, *GetVerificationCodeRequest) (*GetVerificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVerificationCode not implemented")
}
func (UnimplementedUserServer) RemoveVerificationCode(context.Context, *RemoveVerificationCodeRequest) (*RemoveVerificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveVerificationCode not implemented")
}
func (UnimplementedUserServer) ForgetPassword(context.Context, *ForgetPasswordRequest) (*ForgetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgetPassword not implemented")
}
func (UnimplementedUserServer) SetCriticalUserInfo(context.Context, *SetCriticalUserInfoRequest) (*SetCriticalUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCriticalUserInfo not implemented")
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

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CheckEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_CheckEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CheckEmail(ctx, req.(*CheckEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUidByEmailAndPwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUidByEmailAndPwdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUidByEmailAndPwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUidByEmailAndPwd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUidByEmailAndPwd(ctx, req.(*GetUidByEmailAndPwdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CheckUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_CheckUid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CheckUid(ctx, req.(*CheckUidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetUserBaseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserBaseInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetUserBaseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SetUserBaseInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetUserBaseInfo(ctx, req.(*SetUserBaseInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserBaseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserBaseInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserBaseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserBaseInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserBaseInfo(ctx, req.(*GetUserBaseInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVerificationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SetVerificationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetVerificationCode(ctx, req.(*SetVerificationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVerificationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetVerificationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetVerificationCode(ctx, req.(*GetVerificationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RemoveVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveVerificationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RemoveVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_RemoveVerificationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RemoveVerificationCode(ctx, req.(*RemoveVerificationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ForgetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ForgetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ForgetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ForgetPassword(ctx, req.(*ForgetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetCriticalUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCriticalUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetCriticalUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SetCriticalUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetCriticalUserInfo(ctx, req.(*SetCriticalUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userclient.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "CheckEmail",
			Handler:    _User_CheckEmail_Handler,
		},
		{
			MethodName: "GetUidByEmailAndPwd",
			Handler:    _User_GetUidByEmailAndPwd_Handler,
		},
		{
			MethodName: "CheckUid",
			Handler:    _User_CheckUid_Handler,
		},
		{
			MethodName: "SetUserBaseInfo",
			Handler:    _User_SetUserBaseInfo_Handler,
		},
		{
			MethodName: "GetUserBaseInfo",
			Handler:    _User_GetUserBaseInfo_Handler,
		},
		{
			MethodName: "SetVerificationCode",
			Handler:    _User_SetVerificationCode_Handler,
		},
		{
			MethodName: "GetVerificationCode",
			Handler:    _User_GetVerificationCode_Handler,
		},
		{
			MethodName: "RemoveVerificationCode",
			Handler:    _User_RemoveVerificationCode_Handler,
		},
		{
			MethodName: "ForgetPassword",
			Handler:    _User_ForgetPassword_Handler,
		},
		{
			MethodName: "SetCriticalUserInfo",
			Handler:    _User_SetCriticalUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
