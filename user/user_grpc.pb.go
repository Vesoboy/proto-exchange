// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ExchangeService_GetExchangeRates_FullMethodName           = "/user.ExchangeService/GetExchangeRates"
	ExchangeService_GetExchangeRateForCurrency_FullMethodName = "/user.ExchangeService/GetExchangeRateForCurrency"
	ExchangeService_ExchangeCurrency_FullMethodName           = "/user.ExchangeService/ExchangeCurrency"
)

// ExchangeServiceClient is the client API for ExchangeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Определение сервиса
type ExchangeServiceClient interface {
	// Получение курсов обмена всех валют
	GetExchangeRates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ExchangeRatesResponse, error)
	// Получение курса обмена для конкретной валюты
	GetExchangeRateForCurrency(ctx context.Context, in *CurrencyRequest, opts ...grpc.CallOption) (*ExchangeRateResponse, error)
	// Обмен валют
	ExchangeCurrency(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type exchangeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExchangeServiceClient(cc grpc.ClientConnInterface) ExchangeServiceClient {
	return &exchangeServiceClient{cc}
}

func (c *exchangeServiceClient) GetExchangeRates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ExchangeRatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangeRatesResponse)
	err := c.cc.Invoke(ctx, ExchangeService_GetExchangeRates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) GetExchangeRateForCurrency(ctx context.Context, in *CurrencyRequest, opts ...grpc.CallOption) (*ExchangeRateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangeRateResponse)
	err := c.cc.Invoke(ctx, ExchangeService_GetExchangeRateForCurrency_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) ExchangeCurrency(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, ExchangeService_ExchangeCurrency_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangeServiceServer is the server API for ExchangeService service.
// All implementations must embed UnimplementedExchangeServiceServer
// for forward compatibility.
//
// Определение сервиса
type ExchangeServiceServer interface {
	// Получение курсов обмена всех валют
	GetExchangeRates(context.Context, *Empty) (*ExchangeRatesResponse, error)
	// Получение курса обмена для конкретной валюты
	GetExchangeRateForCurrency(context.Context, *CurrencyRequest) (*ExchangeRateResponse, error)
	// Обмен валют
	ExchangeCurrency(context.Context, *ExchangeRequest) (*TransactionResponse, error)
	mustEmbedUnimplementedExchangeServiceServer()
}

// UnimplementedExchangeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExchangeServiceServer struct{}

func (UnimplementedExchangeServiceServer) GetExchangeRates(context.Context, *Empty) (*ExchangeRatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExchangeRates not implemented")
}
func (UnimplementedExchangeServiceServer) GetExchangeRateForCurrency(context.Context, *CurrencyRequest) (*ExchangeRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExchangeRateForCurrency not implemented")
}
func (UnimplementedExchangeServiceServer) ExchangeCurrency(context.Context, *ExchangeRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExchangeCurrency not implemented")
}
func (UnimplementedExchangeServiceServer) mustEmbedUnimplementedExchangeServiceServer() {}
func (UnimplementedExchangeServiceServer) testEmbeddedByValue()                         {}

// UnsafeExchangeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExchangeServiceServer will
// result in compilation errors.
type UnsafeExchangeServiceServer interface {
	mustEmbedUnimplementedExchangeServiceServer()
}

func RegisterExchangeServiceServer(s grpc.ServiceRegistrar, srv ExchangeServiceServer) {
	// If the following call pancis, it indicates UnimplementedExchangeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExchangeService_ServiceDesc, srv)
}

func _ExchangeService_GetExchangeRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).GetExchangeRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_GetExchangeRates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).GetExchangeRates(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_GetExchangeRateForCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).GetExchangeRateForCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_GetExchangeRateForCurrency_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).GetExchangeRateForCurrency(ctx, req.(*CurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_ExchangeCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).ExchangeCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_ExchangeCurrency_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).ExchangeCurrency(ctx, req.(*ExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExchangeService_ServiceDesc is the grpc.ServiceDesc for ExchangeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExchangeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.ExchangeService",
	HandlerType: (*ExchangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExchangeRates",
			Handler:    _ExchangeService_GetExchangeRates_Handler,
		},
		{
			MethodName: "GetExchangeRateForCurrency",
			Handler:    _ExchangeService_GetExchangeRateForCurrency_Handler,
		},
		{
			MethodName: "ExchangeCurrency",
			Handler:    _ExchangeService_ExchangeCurrency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

const (
	Auth_RegisterUser_FullMethodName = "/user.Auth/RegisterUser"
	Auth_LoginUser_FullMethodName    = "/user.Auth/LoginUser"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Auth сервиса
type AuthClient interface {
	// Регистрация пользователя
	RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// Авторизация пользователя
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, Auth_RegisterUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Auth_LoginUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility.
//
// Auth сервиса
type AuthServer interface {
	// Регистрация пользователя
	RegisterUser(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// Авторизация пользователя
	LoginUser(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServer struct{}

func (UnimplementedAuthServer) RegisterUser(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedAuthServer) LoginUser(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}
func (UnimplementedAuthServer) testEmbeddedByValue()              {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	// If the following call pancis, it indicates UnimplementedAuthServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RegisterUser(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _Auth_RegisterUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Auth_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

const (
	FinancialService_GetBalance_FullMethodName = "/user.FinancialService/GetBalance"
	FinancialService_Deposit_FullMethodName    = "/user.FinancialService/Deposit"
	FinancialService_Withdraw_FullMethodName   = "/user.FinancialService/Withdraw"
)

// FinancialServiceClient is the client API for FinancialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Финансовые операции (Баланс, пополнение, вывод)
type FinancialServiceClient interface {
	// Получение баланса пользователя
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	// Пополнение счета
	Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Вывод средств
	Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type financialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFinancialServiceClient(cc grpc.ClientConnInterface) FinancialServiceClient {
	return &financialServiceClient{cc}
}

func (c *financialServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, FinancialService_GetBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialServiceClient) Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, FinancialService_Deposit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialServiceClient) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, FinancialService_Withdraw_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinancialServiceServer is the server API for FinancialService service.
// All implementations must embed UnimplementedFinancialServiceServer
// for forward compatibility.
//
// Финансовые операции (Баланс, пополнение, вывод)
type FinancialServiceServer interface {
	// Получение баланса пользователя
	GetBalance(context.Context, *GetBalanceRequest) (*BalanceResponse, error)
	// Пополнение счета
	Deposit(context.Context, *DepositRequest) (*TransactionResponse, error)
	// Вывод средств
	Withdraw(context.Context, *WithdrawRequest) (*TransactionResponse, error)
	mustEmbedUnimplementedFinancialServiceServer()
}

// UnimplementedFinancialServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFinancialServiceServer struct{}

func (UnimplementedFinancialServiceServer) GetBalance(context.Context, *GetBalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedFinancialServiceServer) Deposit(context.Context, *DepositRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedFinancialServiceServer) Withdraw(context.Context, *WithdrawRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedFinancialServiceServer) mustEmbedUnimplementedFinancialServiceServer() {}
func (UnimplementedFinancialServiceServer) testEmbeddedByValue()                          {}

// UnsafeFinancialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FinancialServiceServer will
// result in compilation errors.
type UnsafeFinancialServiceServer interface {
	mustEmbedUnimplementedFinancialServiceServer()
}

func RegisterFinancialServiceServer(s grpc.ServiceRegistrar, srv FinancialServiceServer) {
	// If the following call pancis, it indicates UnimplementedFinancialServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FinancialService_ServiceDesc, srv)
}

func _FinancialService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialService_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialService_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialService_Deposit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).Deposit(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialService_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialService_Withdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).Withdraw(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FinancialService_ServiceDesc is the grpc.ServiceDesc for FinancialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FinancialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.FinancialService",
	HandlerType: (*FinancialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _FinancialService_GetBalance_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _FinancialService_Deposit_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _FinancialService_Withdraw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
