// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AccountHandlerClient is the client API for AccountHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountHandlerClient interface {
	InqbalanceRq(ctx context.Context, in *InquiryBalanceRq, opts ...grpc.CallOption) (*InquiryBalanceRp, error)
	InqMutationRq(ctx context.Context, in *InquiryMutationRq, opts ...grpc.CallOption) (*InquiryMutationRp, error)
}

type accountHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountHandlerClient(cc grpc.ClientConnInterface) AccountHandlerClient {
	return &accountHandlerClient{cc}
}

func (c *accountHandlerClient) InqbalanceRq(ctx context.Context, in *InquiryBalanceRq, opts ...grpc.CallOption) (*InquiryBalanceRp, error) {
	out := new(InquiryBalanceRp)
	err := c.cc.Invoke(ctx, "/service.AccountHandler/InqbalanceRq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountHandlerClient) InqMutationRq(ctx context.Context, in *InquiryMutationRq, opts ...grpc.CallOption) (*InquiryMutationRp, error) {
	out := new(InquiryMutationRp)
	err := c.cc.Invoke(ctx, "/service.AccountHandler/InqMutationRq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountHandlerServer is the server API for AccountHandler service.
// All implementations must embed UnimplementedAccountHandlerServer
// for forward compatibility
type AccountHandlerServer interface {
	InqbalanceRq(context.Context, *InquiryBalanceRq) (*InquiryBalanceRp, error)
	InqMutationRq(context.Context, *InquiryMutationRq) (*InquiryMutationRp, error)
	mustEmbedUnimplementedAccountHandlerServer()
}

// UnimplementedAccountHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedAccountHandlerServer struct {
}

func (UnimplementedAccountHandlerServer) InqbalanceRq(context.Context, *InquiryBalanceRq) (*InquiryBalanceRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InqbalanceRq not implemented")
}
func (UnimplementedAccountHandlerServer) InqMutationRq(context.Context, *InquiryMutationRq) (*InquiryMutationRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InqMutationRq not implemented")
}
func (UnimplementedAccountHandlerServer) mustEmbedUnimplementedAccountHandlerServer() {}

// UnsafeAccountHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountHandlerServer will
// result in compilation errors.
type UnsafeAccountHandlerServer interface {
	mustEmbedUnimplementedAccountHandlerServer()
}

func RegisterAccountHandlerServer(s grpc.ServiceRegistrar, srv AccountHandlerServer) {
	s.RegisterService(&_AccountHandler_serviceDesc, srv)
}

func _AccountHandler_InqbalanceRq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InquiryBalanceRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountHandlerServer).InqbalanceRq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AccountHandler/InqbalanceRq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountHandlerServer).InqbalanceRq(ctx, req.(*InquiryBalanceRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountHandler_InqMutationRq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InquiryMutationRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountHandlerServer).InqMutationRq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AccountHandler/InqMutationRq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountHandlerServer).InqMutationRq(ctx, req.(*InquiryMutationRq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.AccountHandler",
	HandlerType: (*AccountHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InqbalanceRq",
			Handler:    _AccountHandler_InqbalanceRq_Handler,
		},
		{
			MethodName: "InqMutationRq",
			Handler:    _AccountHandler_InqMutationRq_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

// TransferHandlerClient is the client API for TransferHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransferHandlerClient interface {
	TransferOffUsRq(ctx context.Context, in *TrfOffUsRq, opts ...grpc.CallOption) (*TrfOffUsRp, error)
	TransferOnUsRq(ctx context.Context, in *TrfOnUsRq, opts ...grpc.CallOption) (*TrfOnUsRp, error)
	InqTransferOffUsRq(ctx context.Context, in *InqTrfOffUsRq, opts ...grpc.CallOption) (*InqTrfOffUsRp, error)
	InqTransferOnUsRq(ctx context.Context, in *InqTrfOnUsRq, opts ...grpc.CallOption) (*InqTrfOnUsRp, error)
}

type transferHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferHandlerClient(cc grpc.ClientConnInterface) TransferHandlerClient {
	return &transferHandlerClient{cc}
}

func (c *transferHandlerClient) TransferOffUsRq(ctx context.Context, in *TrfOffUsRq, opts ...grpc.CallOption) (*TrfOffUsRp, error) {
	out := new(TrfOffUsRp)
	err := c.cc.Invoke(ctx, "/service.TransferHandler/TransferOffUsRq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferHandlerClient) TransferOnUsRq(ctx context.Context, in *TrfOnUsRq, opts ...grpc.CallOption) (*TrfOnUsRp, error) {
	out := new(TrfOnUsRp)
	err := c.cc.Invoke(ctx, "/service.TransferHandler/TransferOnUsRq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferHandlerClient) InqTransferOffUsRq(ctx context.Context, in *InqTrfOffUsRq, opts ...grpc.CallOption) (*InqTrfOffUsRp, error) {
	out := new(InqTrfOffUsRp)
	err := c.cc.Invoke(ctx, "/service.TransferHandler/InqTransferOffUsRq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferHandlerClient) InqTransferOnUsRq(ctx context.Context, in *InqTrfOnUsRq, opts ...grpc.CallOption) (*InqTrfOnUsRp, error) {
	out := new(InqTrfOnUsRp)
	err := c.cc.Invoke(ctx, "/service.TransferHandler/InqTransferOnUsRq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransferHandlerServer is the server API for TransferHandler service.
// All implementations must embed UnimplementedTransferHandlerServer
// for forward compatibility
type TransferHandlerServer interface {
	TransferOffUsRq(context.Context, *TrfOffUsRq) (*TrfOffUsRp, error)
	TransferOnUsRq(context.Context, *TrfOnUsRq) (*TrfOnUsRp, error)
	InqTransferOffUsRq(context.Context, *InqTrfOffUsRq) (*InqTrfOffUsRp, error)
	InqTransferOnUsRq(context.Context, *InqTrfOnUsRq) (*InqTrfOnUsRp, error)
	mustEmbedUnimplementedTransferHandlerServer()
}

// UnimplementedTransferHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedTransferHandlerServer struct {
}

func (UnimplementedTransferHandlerServer) TransferOffUsRq(context.Context, *TrfOffUsRq) (*TrfOffUsRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferOffUsRq not implemented")
}
func (UnimplementedTransferHandlerServer) TransferOnUsRq(context.Context, *TrfOnUsRq) (*TrfOnUsRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferOnUsRq not implemented")
}
func (UnimplementedTransferHandlerServer) InqTransferOffUsRq(context.Context, *InqTrfOffUsRq) (*InqTrfOffUsRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InqTransferOffUsRq not implemented")
}
func (UnimplementedTransferHandlerServer) InqTransferOnUsRq(context.Context, *InqTrfOnUsRq) (*InqTrfOnUsRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InqTransferOnUsRq not implemented")
}
func (UnimplementedTransferHandlerServer) mustEmbedUnimplementedTransferHandlerServer() {}

// UnsafeTransferHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransferHandlerServer will
// result in compilation errors.
type UnsafeTransferHandlerServer interface {
	mustEmbedUnimplementedTransferHandlerServer()
}

func RegisterTransferHandlerServer(s grpc.ServiceRegistrar, srv TransferHandlerServer) {
	s.RegisterService(&_TransferHandler_serviceDesc, srv)
}

func _TransferHandler_TransferOffUsRq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrfOffUsRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferHandlerServer).TransferOffUsRq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.TransferHandler/TransferOffUsRq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferHandlerServer).TransferOffUsRq(ctx, req.(*TrfOffUsRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferHandler_TransferOnUsRq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrfOnUsRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferHandlerServer).TransferOnUsRq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.TransferHandler/TransferOnUsRq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferHandlerServer).TransferOnUsRq(ctx, req.(*TrfOnUsRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferHandler_InqTransferOffUsRq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InqTrfOffUsRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferHandlerServer).InqTransferOffUsRq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.TransferHandler/InqTransferOffUsRq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferHandlerServer).InqTransferOffUsRq(ctx, req.(*InqTrfOffUsRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferHandler_InqTransferOnUsRq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InqTrfOnUsRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferHandlerServer).InqTransferOnUsRq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.TransferHandler/InqTransferOnUsRq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferHandlerServer).InqTransferOnUsRq(ctx, req.(*InqTrfOnUsRq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransferHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.TransferHandler",
	HandlerType: (*TransferHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransferOffUsRq",
			Handler:    _TransferHandler_TransferOffUsRq_Handler,
		},
		{
			MethodName: "TransferOnUsRq",
			Handler:    _TransferHandler_TransferOnUsRq_Handler,
		},
		{
			MethodName: "InqTransferOffUsRq",
			Handler:    _TransferHandler_InqTransferOffUsRq_Handler,
		},
		{
			MethodName: "InqTransferOnUsRq",
			Handler:    _TransferHandler_InqTransferOnUsRq_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

// SecurityHandlerClient is the client API for SecurityHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecurityHandlerClient interface {
	GeneratePartnerKey(ctx context.Context, in *GenPartnerKeyRq, opts ...grpc.CallOption) (*GenPartnerKeyRp, error)
	VerifySessionId(ctx context.Context, in *VerifySessionIdRq, opts ...grpc.CallOption) (*VerifySessionIdRp, error)
	DecryptValue(ctx context.Context, in *DecryptRq, opts ...grpc.CallOption) (*DecryptRp, error)
	EncryptValue(ctx context.Context, in *EncryptRq, opts ...grpc.CallOption) (*EncryptRp, error)
	Login(ctx context.Context, in *LoginRq, opts ...grpc.CallOption) (*LoginRp, error)
	GeneratePartnerOnboardingKey(ctx context.Context, in *GenPartnerOnboardKeyRq, opts ...grpc.CallOption) (*GenPartnerOnboardKeyRp, error)
}

type securityHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewSecurityHandlerClient(cc grpc.ClientConnInterface) SecurityHandlerClient {
	return &securityHandlerClient{cc}
}

func (c *securityHandlerClient) GeneratePartnerKey(ctx context.Context, in *GenPartnerKeyRq, opts ...grpc.CallOption) (*GenPartnerKeyRp, error) {
	out := new(GenPartnerKeyRp)
	err := c.cc.Invoke(ctx, "/service.SecurityHandler/GeneratePartnerKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityHandlerClient) VerifySessionId(ctx context.Context, in *VerifySessionIdRq, opts ...grpc.CallOption) (*VerifySessionIdRp, error) {
	out := new(VerifySessionIdRp)
	err := c.cc.Invoke(ctx, "/service.SecurityHandler/VerifySessionId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityHandlerClient) DecryptValue(ctx context.Context, in *DecryptRq, opts ...grpc.CallOption) (*DecryptRp, error) {
	out := new(DecryptRp)
	err := c.cc.Invoke(ctx, "/service.SecurityHandler/DecryptValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityHandlerClient) EncryptValue(ctx context.Context, in *EncryptRq, opts ...grpc.CallOption) (*EncryptRp, error) {
	out := new(EncryptRp)
	err := c.cc.Invoke(ctx, "/service.SecurityHandler/EncryptValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityHandlerClient) Login(ctx context.Context, in *LoginRq, opts ...grpc.CallOption) (*LoginRp, error) {
	out := new(LoginRp)
	err := c.cc.Invoke(ctx, "/service.SecurityHandler/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityHandlerClient) GeneratePartnerOnboardingKey(ctx context.Context, in *GenPartnerOnboardKeyRq, opts ...grpc.CallOption) (*GenPartnerOnboardKeyRp, error) {
	out := new(GenPartnerOnboardKeyRp)
	err := c.cc.Invoke(ctx, "/service.SecurityHandler/GeneratePartnerOnboardingKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityHandlerServer is the server API for SecurityHandler service.
// All implementations must embed UnimplementedSecurityHandlerServer
// for forward compatibility
type SecurityHandlerServer interface {
	GeneratePartnerKey(context.Context, *GenPartnerKeyRq) (*GenPartnerKeyRp, error)
	VerifySessionId(context.Context, *VerifySessionIdRq) (*VerifySessionIdRp, error)
	DecryptValue(context.Context, *DecryptRq) (*DecryptRp, error)
	EncryptValue(context.Context, *EncryptRq) (*EncryptRp, error)
	Login(context.Context, *LoginRq) (*LoginRp, error)
	GeneratePartnerOnboardingKey(context.Context, *GenPartnerOnboardKeyRq) (*GenPartnerOnboardKeyRp, error)
	mustEmbedUnimplementedSecurityHandlerServer()
}

// UnimplementedSecurityHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedSecurityHandlerServer struct {
}

func (UnimplementedSecurityHandlerServer) GeneratePartnerKey(context.Context, *GenPartnerKeyRq) (*GenPartnerKeyRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePartnerKey not implemented")
}
func (UnimplementedSecurityHandlerServer) VerifySessionId(context.Context, *VerifySessionIdRq) (*VerifySessionIdRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifySessionId not implemented")
}
func (UnimplementedSecurityHandlerServer) DecryptValue(context.Context, *DecryptRq) (*DecryptRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecryptValue not implemented")
}
func (UnimplementedSecurityHandlerServer) EncryptValue(context.Context, *EncryptRq) (*EncryptRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncryptValue not implemented")
}
func (UnimplementedSecurityHandlerServer) Login(context.Context, *LoginRq) (*LoginRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSecurityHandlerServer) GeneratePartnerOnboardingKey(context.Context, *GenPartnerOnboardKeyRq) (*GenPartnerOnboardKeyRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePartnerOnboardingKey not implemented")
}
func (UnimplementedSecurityHandlerServer) mustEmbedUnimplementedSecurityHandlerServer() {}

// UnsafeSecurityHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecurityHandlerServer will
// result in compilation errors.
type UnsafeSecurityHandlerServer interface {
	mustEmbedUnimplementedSecurityHandlerServer()
}

func RegisterSecurityHandlerServer(s grpc.ServiceRegistrar, srv SecurityHandlerServer) {
	s.RegisterService(&_SecurityHandler_serviceDesc, srv)
}

func _SecurityHandler_GeneratePartnerKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenPartnerKeyRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityHandlerServer).GeneratePartnerKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.SecurityHandler/GeneratePartnerKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityHandlerServer).GeneratePartnerKey(ctx, req.(*GenPartnerKeyRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityHandler_VerifySessionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifySessionIdRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityHandlerServer).VerifySessionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.SecurityHandler/VerifySessionId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityHandlerServer).VerifySessionId(ctx, req.(*VerifySessionIdRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityHandler_DecryptValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityHandlerServer).DecryptValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.SecurityHandler/DecryptValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityHandlerServer).DecryptValue(ctx, req.(*DecryptRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityHandler_EncryptValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityHandlerServer).EncryptValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.SecurityHandler/EncryptValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityHandlerServer).EncryptValue(ctx, req.(*EncryptRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityHandler_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityHandlerServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.SecurityHandler/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityHandlerServer).Login(ctx, req.(*LoginRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityHandler_GeneratePartnerOnboardingKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenPartnerOnboardKeyRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityHandlerServer).GeneratePartnerOnboardingKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.SecurityHandler/GeneratePartnerOnboardingKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityHandlerServer).GeneratePartnerOnboardingKey(ctx, req.(*GenPartnerOnboardKeyRq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecurityHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.SecurityHandler",
	HandlerType: (*SecurityHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeneratePartnerKey",
			Handler:    _SecurityHandler_GeneratePartnerKey_Handler,
		},
		{
			MethodName: "VerifySessionId",
			Handler:    _SecurityHandler_VerifySessionId_Handler,
		},
		{
			MethodName: "DecryptValue",
			Handler:    _SecurityHandler_DecryptValue_Handler,
		},
		{
			MethodName: "EncryptValue",
			Handler:    _SecurityHandler_EncryptValue_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _SecurityHandler_Login_Handler,
		},
		{
			MethodName: "GeneratePartnerOnboardingKey",
			Handler:    _SecurityHandler_GeneratePartnerOnboardingKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

// UserHandlerClient is the client API for UserHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserHandlerClient interface {
	UserInfo(ctx context.Context, in *UserInfoRq, opts ...grpc.CallOption) (*UserInfoRp, error)
	UserData(ctx context.Context, in *UserDataRq, opts ...grpc.CallOption) (*UserDataRp, error)
}

type userHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserHandlerClient(cc grpc.ClientConnInterface) UserHandlerClient {
	return &userHandlerClient{cc}
}

func (c *userHandlerClient) UserInfo(ctx context.Context, in *UserInfoRq, opts ...grpc.CallOption) (*UserInfoRp, error) {
	out := new(UserInfoRp)
	err := c.cc.Invoke(ctx, "/service.UserHandler/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userHandlerClient) UserData(ctx context.Context, in *UserDataRq, opts ...grpc.CallOption) (*UserDataRp, error) {
	out := new(UserDataRp)
	err := c.cc.Invoke(ctx, "/service.UserHandler/UserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserHandlerServer is the server API for UserHandler service.
// All implementations must embed UnimplementedUserHandlerServer
// for forward compatibility
type UserHandlerServer interface {
	UserInfo(context.Context, *UserInfoRq) (*UserInfoRp, error)
	UserData(context.Context, *UserDataRq) (*UserDataRp, error)
	mustEmbedUnimplementedUserHandlerServer()
}

// UnimplementedUserHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedUserHandlerServer struct {
}

func (UnimplementedUserHandlerServer) UserInfo(context.Context, *UserInfoRq) (*UserInfoRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedUserHandlerServer) UserData(context.Context, *UserDataRq) (*UserDataRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserData not implemented")
}
func (UnimplementedUserHandlerServer) mustEmbedUnimplementedUserHandlerServer() {}

// UnsafeUserHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserHandlerServer will
// result in compilation errors.
type UnsafeUserHandlerServer interface {
	mustEmbedUnimplementedUserHandlerServer()
}

func RegisterUserHandlerServer(s grpc.ServiceRegistrar, srv UserHandlerServer) {
	s.RegisterService(&_UserHandler_serviceDesc, srv)
}

func _UserHandler_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHandlerServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserHandler/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHandlerServer).UserInfo(ctx, req.(*UserInfoRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserHandler_UserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDataRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHandlerServer).UserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserHandler/UserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHandlerServer).UserData(ctx, req.(*UserDataRq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.UserHandler",
	HandlerType: (*UserHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserInfo",
			Handler:    _UserHandler_UserInfo_Handler,
		},
		{
			MethodName: "UserData",
			Handler:    _UserHandler_UserData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

// ParamHandlerClient is the client API for ParamHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParamHandlerClient interface {
	Province(ctx context.Context, in *ProvinceRq, opts ...grpc.CallOption) (*ProvinceRp, error)
	City(ctx context.Context, in *CityRq, opts ...grpc.CallOption) (*CityRp, error)
	Degree(ctx context.Context, in *DegreeRq, opts ...grpc.CallOption) (*DegreeRp, error)
	Income(ctx context.Context, in *IncomeRq, opts ...grpc.CallOption) (*IncomeRp, error)
	IndustrialSector(ctx context.Context, in *IndustrialSectorRq, opts ...grpc.CallOption) (*IndustrialSectorRp, error)
	OpeningPurpose(ctx context.Context, in *OpeningPurposeRq, opts ...grpc.CallOption) (*OpeningPurposeRp, error)
	SourceIncome(ctx context.Context, in *SourceIncomeRq, opts ...grpc.CallOption) (*SourceIncomeRp, error)
	JobTypes(ctx context.Context, in *JobTypesRq, opts ...grpc.CallOption) (*JobTypesRp, error)
}

type paramHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewParamHandlerClient(cc grpc.ClientConnInterface) ParamHandlerClient {
	return &paramHandlerClient{cc}
}

func (c *paramHandlerClient) Province(ctx context.Context, in *ProvinceRq, opts ...grpc.CallOption) (*ProvinceRp, error) {
	out := new(ProvinceRp)
	err := c.cc.Invoke(ctx, "/service.ParamHandler/Province", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paramHandlerClient) City(ctx context.Context, in *CityRq, opts ...grpc.CallOption) (*CityRp, error) {
	out := new(CityRp)
	err := c.cc.Invoke(ctx, "/service.ParamHandler/City", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paramHandlerClient) Degree(ctx context.Context, in *DegreeRq, opts ...grpc.CallOption) (*DegreeRp, error) {
	out := new(DegreeRp)
	err := c.cc.Invoke(ctx, "/service.ParamHandler/Degree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paramHandlerClient) Income(ctx context.Context, in *IncomeRq, opts ...grpc.CallOption) (*IncomeRp, error) {
	out := new(IncomeRp)
	err := c.cc.Invoke(ctx, "/service.ParamHandler/Income", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paramHandlerClient) IndustrialSector(ctx context.Context, in *IndustrialSectorRq, opts ...grpc.CallOption) (*IndustrialSectorRp, error) {
	out := new(IndustrialSectorRp)
	err := c.cc.Invoke(ctx, "/service.ParamHandler/IndustrialSector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paramHandlerClient) OpeningPurpose(ctx context.Context, in *OpeningPurposeRq, opts ...grpc.CallOption) (*OpeningPurposeRp, error) {
	out := new(OpeningPurposeRp)
	err := c.cc.Invoke(ctx, "/service.ParamHandler/OpeningPurpose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paramHandlerClient) SourceIncome(ctx context.Context, in *SourceIncomeRq, opts ...grpc.CallOption) (*SourceIncomeRp, error) {
	out := new(SourceIncomeRp)
	err := c.cc.Invoke(ctx, "/service.ParamHandler/SourceIncome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paramHandlerClient) JobTypes(ctx context.Context, in *JobTypesRq, opts ...grpc.CallOption) (*JobTypesRp, error) {
	out := new(JobTypesRp)
	err := c.cc.Invoke(ctx, "/service.ParamHandler/JobTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParamHandlerServer is the server API for ParamHandler service.
// All implementations must embed UnimplementedParamHandlerServer
// for forward compatibility
type ParamHandlerServer interface {
	Province(context.Context, *ProvinceRq) (*ProvinceRp, error)
	City(context.Context, *CityRq) (*CityRp, error)
	Degree(context.Context, *DegreeRq) (*DegreeRp, error)
	Income(context.Context, *IncomeRq) (*IncomeRp, error)
	IndustrialSector(context.Context, *IndustrialSectorRq) (*IndustrialSectorRp, error)
	OpeningPurpose(context.Context, *OpeningPurposeRq) (*OpeningPurposeRp, error)
	SourceIncome(context.Context, *SourceIncomeRq) (*SourceIncomeRp, error)
	JobTypes(context.Context, *JobTypesRq) (*JobTypesRp, error)
	mustEmbedUnimplementedParamHandlerServer()
}

// UnimplementedParamHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedParamHandlerServer struct {
}

func (UnimplementedParamHandlerServer) Province(context.Context, *ProvinceRq) (*ProvinceRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Province not implemented")
}
func (UnimplementedParamHandlerServer) City(context.Context, *CityRq) (*CityRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method City not implemented")
}
func (UnimplementedParamHandlerServer) Degree(context.Context, *DegreeRq) (*DegreeRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Degree not implemented")
}
func (UnimplementedParamHandlerServer) Income(context.Context, *IncomeRq) (*IncomeRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Income not implemented")
}
func (UnimplementedParamHandlerServer) IndustrialSector(context.Context, *IndustrialSectorRq) (*IndustrialSectorRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IndustrialSector not implemented")
}
func (UnimplementedParamHandlerServer) OpeningPurpose(context.Context, *OpeningPurposeRq) (*OpeningPurposeRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpeningPurpose not implemented")
}
func (UnimplementedParamHandlerServer) SourceIncome(context.Context, *SourceIncomeRq) (*SourceIncomeRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SourceIncome not implemented")
}
func (UnimplementedParamHandlerServer) JobTypes(context.Context, *JobTypesRq) (*JobTypesRp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobTypes not implemented")
}
func (UnimplementedParamHandlerServer) mustEmbedUnimplementedParamHandlerServer() {}

// UnsafeParamHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParamHandlerServer will
// result in compilation errors.
type UnsafeParamHandlerServer interface {
	mustEmbedUnimplementedParamHandlerServer()
}

func RegisterParamHandlerServer(s grpc.ServiceRegistrar, srv ParamHandlerServer) {
	s.RegisterService(&_ParamHandler_serviceDesc, srv)
}

func _ParamHandler_Province_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvinceRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamHandlerServer).Province(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ParamHandler/Province",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamHandlerServer).Province(ctx, req.(*ProvinceRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParamHandler_City_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CityRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamHandlerServer).City(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ParamHandler/City",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamHandlerServer).City(ctx, req.(*CityRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParamHandler_Degree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DegreeRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamHandlerServer).Degree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ParamHandler/Degree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamHandlerServer).Degree(ctx, req.(*DegreeRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParamHandler_Income_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncomeRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamHandlerServer).Income(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ParamHandler/Income",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamHandlerServer).Income(ctx, req.(*IncomeRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParamHandler_IndustrialSector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndustrialSectorRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamHandlerServer).IndustrialSector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ParamHandler/IndustrialSector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamHandlerServer).IndustrialSector(ctx, req.(*IndustrialSectorRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParamHandler_OpeningPurpose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpeningPurposeRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamHandlerServer).OpeningPurpose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ParamHandler/OpeningPurpose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamHandlerServer).OpeningPurpose(ctx, req.(*OpeningPurposeRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParamHandler_SourceIncome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SourceIncomeRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamHandlerServer).SourceIncome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ParamHandler/SourceIncome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamHandlerServer).SourceIncome(ctx, req.(*SourceIncomeRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParamHandler_JobTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobTypesRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamHandlerServer).JobTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ParamHandler/JobTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamHandlerServer).JobTypes(ctx, req.(*JobTypesRq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ParamHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.ParamHandler",
	HandlerType: (*ParamHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Province",
			Handler:    _ParamHandler_Province_Handler,
		},
		{
			MethodName: "City",
			Handler:    _ParamHandler_City_Handler,
		},
		{
			MethodName: "Degree",
			Handler:    _ParamHandler_Degree_Handler,
		},
		{
			MethodName: "Income",
			Handler:    _ParamHandler_Income_Handler,
		},
		{
			MethodName: "IndustrialSector",
			Handler:    _ParamHandler_IndustrialSector_Handler,
		},
		{
			MethodName: "OpeningPurpose",
			Handler:    _ParamHandler_OpeningPurpose_Handler,
		},
		{
			MethodName: "SourceIncome",
			Handler:    _ParamHandler_SourceIncome_Handler,
		},
		{
			MethodName: "JobTypes",
			Handler:    _ParamHandler_JobTypes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
