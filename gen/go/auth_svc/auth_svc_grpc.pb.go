// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: auth_svc/auth_svc.proto

package auth1

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

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	CheckToken(ctx context.Context, in *CheckTokenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error)
	CheckPermisson(ctx context.Context, in *CheckPermissonRequest, opts ...grpc.CallOption) (*CheckPermissonResponse, error)
	TakePin(ctx context.Context, in *TakePinRequest, opts ...grpc.CallOption) (*TakePinResponse, error)
	SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error)
	HasPassword(ctx context.Context, in *HasPasswordRequest, opts ...grpc.CallOption) (*HasPasswordResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*SetPasswordResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	CheckVersion(ctx context.Context, in *CheckVersionRequest, opts ...grpc.CallOption) (*CheckVersionResponse, error)
	CustomerInfo(ctx context.Context, in *CustomerInfoRequest, opts ...grpc.CallOption) (*CustomerInfoResponse, error)
	GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionResponse, error)
	AddPermissionCustomer(ctx context.Context, in *AddPermissionCustomerRequest, opts ...grpc.CallOption) (*AddPermissionCustomerResponse, error)
	AddPermissionGroup(ctx context.Context, in *AddPermissionGroupRequest, opts ...grpc.CallOption) (*AddPermissionGroupResponse, error)
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error)
	AddGroup(ctx context.Context, in *AddGroupRequest, opts ...grpc.CallOption) (*AddGroupResponse, error)
	AddCustomerToGroup(ctx context.Context, in *AddCustomerToGroupRequest, opts ...grpc.CallOption) (*AddCustomerToGroupResponse, error)
	CheckTOTP(ctx context.Context, in *CheckTOTPRequest, opts ...grpc.CallOption) (*CheckTOTPResponse, error)
	TOTPLinkAccountTmp(ctx context.Context, in *TOTPLinkAccTmpReq, opts ...grpc.CallOption) (*TOTPLinkAccTmpResp, error)
	TOTPLinkAccount(ctx context.Context, in *TOTPLinkAccReq, opts ...grpc.CallOption) (*TOTPLinkAccResp, error)
	TOTPUnlinkAccount(ctx context.Context, in *TOTPUnlinkAccReq, opts ...grpc.CallOption) (*TOTPUnlinkAccResp, error)
	TOTPHasLinkedAccount(ctx context.Context, in *TOTPHasLinkedAccReq, opts ...grpc.CallOption) (*TOTPHasLinkedAccResp, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckToken(ctx context.Context, in *CheckTokenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error) {
	out := new(CheckTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/CheckToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckPermisson(ctx context.Context, in *CheckPermissonRequest, opts ...grpc.CallOption) (*CheckPermissonResponse, error) {
	out := new(CheckPermissonResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/CheckPermisson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) TakePin(ctx context.Context, in *TakePinRequest, opts ...grpc.CallOption) (*TakePinResponse, error) {
	out := new(TakePinResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/TakePin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error) {
	out := new(SignOutResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/SignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) HasPassword(ctx context.Context, in *HasPasswordRequest, opts ...grpc.CallOption) (*HasPasswordResponse, error) {
	out := new(HasPasswordResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/HasPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*SetPasswordResponse, error) {
	out := new(SetPasswordResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/SetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckVersion(ctx context.Context, in *CheckVersionRequest, opts ...grpc.CallOption) (*CheckVersionResponse, error) {
	out := new(CheckVersionResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/CheckVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CustomerInfo(ctx context.Context, in *CustomerInfoRequest, opts ...grpc.CallOption) (*CustomerInfoResponse, error) {
	out := new(CustomerInfoResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/CustomerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionResponse, error) {
	out := new(GetActionResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AddPermissionCustomer(ctx context.Context, in *AddPermissionCustomerRequest, opts ...grpc.CallOption) (*AddPermissionCustomerResponse, error) {
	out := new(AddPermissionCustomerResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/AddPermissionCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AddPermissionGroup(ctx context.Context, in *AddPermissionGroupRequest, opts ...grpc.CallOption) (*AddPermissionGroupResponse, error) {
	out := new(AddPermissionGroupResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/AddPermissionGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error) {
	out := new(GetGroupResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AddGroup(ctx context.Context, in *AddGroupRequest, opts ...grpc.CallOption) (*AddGroupResponse, error) {
	out := new(AddGroupResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/AddGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AddCustomerToGroup(ctx context.Context, in *AddCustomerToGroupRequest, opts ...grpc.CallOption) (*AddCustomerToGroupResponse, error) {
	out := new(AddCustomerToGroupResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/AddCustomerToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckTOTP(ctx context.Context, in *CheckTOTPRequest, opts ...grpc.CallOption) (*CheckTOTPResponse, error) {
	out := new(CheckTOTPResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/CheckTOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) TOTPLinkAccountTmp(ctx context.Context, in *TOTPLinkAccTmpReq, opts ...grpc.CallOption) (*TOTPLinkAccTmpResp, error) {
	out := new(TOTPLinkAccTmpResp)
	err := c.cc.Invoke(ctx, "/auth.Auth/TOTPLinkAccountTmp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) TOTPLinkAccount(ctx context.Context, in *TOTPLinkAccReq, opts ...grpc.CallOption) (*TOTPLinkAccResp, error) {
	out := new(TOTPLinkAccResp)
	err := c.cc.Invoke(ctx, "/auth.Auth/TOTPLinkAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) TOTPUnlinkAccount(ctx context.Context, in *TOTPUnlinkAccReq, opts ...grpc.CallOption) (*TOTPUnlinkAccResp, error) {
	out := new(TOTPUnlinkAccResp)
	err := c.cc.Invoke(ctx, "/auth.Auth/TOTPUnlinkAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) TOTPHasLinkedAccount(ctx context.Context, in *TOTPHasLinkedAccReq, opts ...grpc.CallOption) (*TOTPHasLinkedAccResp, error) {
	out := new(TOTPHasLinkedAccResp)
	err := c.cc.Invoke(ctx, "/auth.Auth/TOTPHasLinkedAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	CheckToken(context.Context, *CheckTokenRequest) (*CheckTokenResponse, error)
	CheckPermisson(context.Context, *CheckPermissonRequest) (*CheckPermissonResponse, error)
	TakePin(context.Context, *TakePinRequest) (*TakePinResponse, error)
	SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error)
	HasPassword(context.Context, *HasPasswordRequest) (*HasPasswordResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	SetPassword(context.Context, *SetPasswordRequest) (*SetPasswordResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	CheckVersion(context.Context, *CheckVersionRequest) (*CheckVersionResponse, error)
	CustomerInfo(context.Context, *CustomerInfoRequest) (*CustomerInfoResponse, error)
	GetAction(context.Context, *GetActionRequest) (*GetActionResponse, error)
	AddPermissionCustomer(context.Context, *AddPermissionCustomerRequest) (*AddPermissionCustomerResponse, error)
	AddPermissionGroup(context.Context, *AddPermissionGroupRequest) (*AddPermissionGroupResponse, error)
	GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error)
	AddGroup(context.Context, *AddGroupRequest) (*AddGroupResponse, error)
	AddCustomerToGroup(context.Context, *AddCustomerToGroupRequest) (*AddCustomerToGroupResponse, error)
	CheckTOTP(context.Context, *CheckTOTPRequest) (*CheckTOTPResponse, error)
	TOTPLinkAccountTmp(context.Context, *TOTPLinkAccTmpReq) (*TOTPLinkAccTmpResp, error)
	TOTPLinkAccount(context.Context, *TOTPLinkAccReq) (*TOTPLinkAccResp, error)
	TOTPUnlinkAccount(context.Context, *TOTPUnlinkAccReq) (*TOTPUnlinkAccResp, error)
	TOTPHasLinkedAccount(context.Context, *TOTPHasLinkedAccReq) (*TOTPHasLinkedAccResp, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedAuthServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAuthServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAuthServer) CheckToken(context.Context, *CheckTokenRequest) (*CheckTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckToken not implemented")
}
func (UnimplementedAuthServer) CheckPermisson(context.Context, *CheckPermissonRequest) (*CheckPermissonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermisson not implemented")
}
func (UnimplementedAuthServer) TakePin(context.Context, *TakePinRequest) (*TakePinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakePin not implemented")
}
func (UnimplementedAuthServer) SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (UnimplementedAuthServer) HasPassword(context.Context, *HasPasswordRequest) (*HasPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasPassword not implemented")
}
func (UnimplementedAuthServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedAuthServer) SetPassword(context.Context, *SetPasswordRequest) (*SetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPassword not implemented")
}
func (UnimplementedAuthServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedAuthServer) CheckVersion(context.Context, *CheckVersionRequest) (*CheckVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckVersion not implemented")
}
func (UnimplementedAuthServer) CustomerInfo(context.Context, *CustomerInfoRequest) (*CustomerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomerInfo not implemented")
}
func (UnimplementedAuthServer) GetAction(context.Context, *GetActionRequest) (*GetActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAction not implemented")
}
func (UnimplementedAuthServer) AddPermissionCustomer(context.Context, *AddPermissionCustomerRequest) (*AddPermissionCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPermissionCustomer not implemented")
}
func (UnimplementedAuthServer) AddPermissionGroup(context.Context, *AddPermissionGroupRequest) (*AddPermissionGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPermissionGroup not implemented")
}
func (UnimplementedAuthServer) GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedAuthServer) AddGroup(context.Context, *AddGroupRequest) (*AddGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroup not implemented")
}
func (UnimplementedAuthServer) AddCustomerToGroup(context.Context, *AddCustomerToGroupRequest) (*AddCustomerToGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCustomerToGroup not implemented")
}
func (UnimplementedAuthServer) CheckTOTP(context.Context, *CheckTOTPRequest) (*CheckTOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTOTP not implemented")
}
func (UnimplementedAuthServer) TOTPLinkAccountTmp(context.Context, *TOTPLinkAccTmpReq) (*TOTPLinkAccTmpResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TOTPLinkAccountTmp not implemented")
}
func (UnimplementedAuthServer) TOTPLinkAccount(context.Context, *TOTPLinkAccReq) (*TOTPLinkAccResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TOTPLinkAccount not implemented")
}
func (UnimplementedAuthServer) TOTPUnlinkAccount(context.Context, *TOTPUnlinkAccReq) (*TOTPUnlinkAccResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TOTPUnlinkAccount not implemented")
}
func (UnimplementedAuthServer) TOTPHasLinkedAccount(context.Context, *TOTPHasLinkedAccReq) (*TOTPHasLinkedAccResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TOTPHasLinkedAccount not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/CheckToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckToken(ctx, req.(*CheckTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckPermisson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckPermisson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/CheckPermisson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckPermisson(ctx, req.(*CheckPermissonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_TakePin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TakePinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).TakePin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/TakePin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).TakePin(ctx, req.(*TakePinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SignOut(ctx, req.(*SignOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_HasPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).HasPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/HasPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).HasPassword(ctx, req.(*HasPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/SetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SetPassword(ctx, req.(*SetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/CheckVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckVersion(ctx, req.(*CheckVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CustomerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CustomerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/CustomerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CustomerInfo(ctx, req.(*CustomerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetAction(ctx, req.(*GetActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AddPermissionCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPermissionCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AddPermissionCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/AddPermissionCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AddPermissionCustomer(ctx, req.(*AddPermissionCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AddPermissionGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPermissionGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AddPermissionGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/AddPermissionGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AddPermissionGroup(ctx, req.(*AddPermissionGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AddGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AddGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/AddGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AddGroup(ctx, req.(*AddGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AddCustomerToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCustomerToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AddCustomerToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/AddCustomerToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AddCustomerToGroup(ctx, req.(*AddCustomerToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckTOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckTOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/CheckTOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckTOTP(ctx, req.(*CheckTOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_TOTPLinkAccountTmp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TOTPLinkAccTmpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).TOTPLinkAccountTmp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/TOTPLinkAccountTmp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).TOTPLinkAccountTmp(ctx, req.(*TOTPLinkAccTmpReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_TOTPLinkAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TOTPLinkAccReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).TOTPLinkAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/TOTPLinkAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).TOTPLinkAccount(ctx, req.(*TOTPLinkAccReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_TOTPUnlinkAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TOTPUnlinkAccReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).TOTPUnlinkAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/TOTPUnlinkAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).TOTPUnlinkAccount(ctx, req.(*TOTPUnlinkAccReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_TOTPHasLinkedAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TOTPHasLinkedAccReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).TOTPHasLinkedAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/TOTPHasLinkedAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).TOTPHasLinkedAccount(ctx, req.(*TOTPHasLinkedAccReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _Auth_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Auth_SignIn_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _Auth_RefreshToken_Handler,
		},
		{
			MethodName: "CheckToken",
			Handler:    _Auth_CheckToken_Handler,
		},
		{
			MethodName: "CheckPermisson",
			Handler:    _Auth_CheckPermisson_Handler,
		},
		{
			MethodName: "TakePin",
			Handler:    _Auth_TakePin_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _Auth_SignOut_Handler,
		},
		{
			MethodName: "HasPassword",
			Handler:    _Auth_HasPassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _Auth_ResetPassword_Handler,
		},
		{
			MethodName: "SetPassword",
			Handler:    _Auth_SetPassword_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _Auth_ChangePassword_Handler,
		},
		{
			MethodName: "CheckVersion",
			Handler:    _Auth_CheckVersion_Handler,
		},
		{
			MethodName: "CustomerInfo",
			Handler:    _Auth_CustomerInfo_Handler,
		},
		{
			MethodName: "GetAction",
			Handler:    _Auth_GetAction_Handler,
		},
		{
			MethodName: "AddPermissionCustomer",
			Handler:    _Auth_AddPermissionCustomer_Handler,
		},
		{
			MethodName: "AddPermissionGroup",
			Handler:    _Auth_AddPermissionGroup_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _Auth_GetGroup_Handler,
		},
		{
			MethodName: "AddGroup",
			Handler:    _Auth_AddGroup_Handler,
		},
		{
			MethodName: "AddCustomerToGroup",
			Handler:    _Auth_AddCustomerToGroup_Handler,
		},
		{
			MethodName: "CheckTOTP",
			Handler:    _Auth_CheckTOTP_Handler,
		},
		{
			MethodName: "TOTPLinkAccountTmp",
			Handler:    _Auth_TOTPLinkAccountTmp_Handler,
		},
		{
			MethodName: "TOTPLinkAccount",
			Handler:    _Auth_TOTPLinkAccount_Handler,
		},
		{
			MethodName: "TOTPUnlinkAccount",
			Handler:    _Auth_TOTPUnlinkAccount_Handler,
		},
		{
			MethodName: "TOTPHasLinkedAccount",
			Handler:    _Auth_TOTPHasLinkedAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_svc/auth_svc.proto",
}
