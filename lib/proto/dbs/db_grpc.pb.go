// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: lib/proto/dbs/db.proto

package dbs

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

// FlyDBServiceClient is the client API for FlyDBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlyDBServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Put(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelResponse, error)
	Type(ctx context.Context, in *TypeRequest, opts ...grpc.CallOption) (*TypeResponse, error)
	StrLen(ctx context.Context, in *StrLenRequest, opts ...grpc.CallOption) (*StrLenResponse, error)
	GetSet(ctx context.Context, in *GetSetRequest, opts ...grpc.CallOption) (*GetSetResponse, error)
	Append(ctx context.Context, in *AppendRequest, opts ...grpc.CallOption) (*AppendResponse, error)
	Incr(ctx context.Context, in *IncrRequest, opts ...grpc.CallOption) (*IncrResponse, error)
	IncrBy(ctx context.Context, in *IncrByRequest, opts ...grpc.CallOption) (*IncrByResponse, error)
	IncrByFloat(ctx context.Context, in *IncrByFloatRequest, opts ...grpc.CallOption) (*IncrByFloatResponse, error)
	Decr(ctx context.Context, in *DecrRequest, opts ...grpc.CallOption) (*DecrResponse, error)
	DecrBy(ctx context.Context, in *DecrByRequest, opts ...grpc.CallOption) (*DecrByResponse, error)
	Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error)
	Expire(ctx context.Context, in *ExpireRequest, opts ...grpc.CallOption) (*ExpireResponse, error)
	Persist(ctx context.Context, in *PersistRequest, opts ...grpc.CallOption) (*PersistResponse, error)
}

type flyDBServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlyDBServiceClient(cc grpc.ClientConnInterface) FlyDBServiceClient {
	return &flyDBServiceClient{cc}
}

func (c *flyDBServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/dbs.FlyDBService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flyDBServiceClient) Put(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/dbs.FlyDBService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flyDBServiceClient) Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelResponse, error) {
	out := new(DelResponse)
	err := c.cc.Invoke(ctx, "/dbs.FlyDBService/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flyDBServiceClient) Type(ctx context.Context, in *TypeRequest, opts ...grpc.CallOption) (*TypeResponse, error) {
	out := new(TypeResponse)
	err := c.cc.Invoke(ctx, "/dbs.FlyDBService/Type", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flyDBServiceClient) StrLen(ctx context.Context, in *StrLenRequest, opts ...grpc.CallOption) (*StrLenResponse, error) {
	out := new(StrLenResponse)
	err := c.cc.Invoke(ctx, "/dbs.FlyDBService/StrLen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flyDBServiceClient) GetSet(ctx context.Context, in *GetSetRequest, opts ...grpc.CallOption) (*GetSetResponse, error) {
	out := new(GetSetResponse)
	err := c.cc.Invoke(ctx, "/dbs.FlyDBService/GetSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flyDBServiceClient) Append(ctx context.Context, in *AppendRequest, opts ...grpc.CallOption) (*AppendResponse, error) {
	out := new(AppendResponse)
	err := c.cc.Invoke(ctx, "/dbs.FlyDBService/Append", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flyDBServiceClient) Incr(ctx context.Context, in *IncrRequest, opts ...grpc.CallOption) (*IncrResponse, error) {
	out := new(IncrResponse)
	err := c.cc.Invoke(ctx, "/dbs.FlyDBService/Incr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flyDBServiceClient) IncrBy(ctx context.Context, in *IncrByRequest, opts ...grpc.CallOption) (*IncrByResponse, error) {
	out := new(IncrByResponse)
	err := c.cc.Invoke(ctx, "/dbs.FlyDBService/IncrBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flyDBServiceClient) IncrByFloat(ctx context.Context, in *IncrByFloatRequest, opts ...grpc.CallOption) (*IncrByFloatResponse, error) {
	out := new(IncrByFloatResponse)
	err := c.cc.Invoke(ctx, "/dbs.FlyDBService/IncrByFloat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flyDBServiceClient) Decr(ctx context.Context, in *DecrRequest, opts ...grpc.CallOption) (*DecrResponse, error) {
	out := new(DecrResponse)
	err := c.cc.Invoke(ctx, "/dbs.FlyDBService/Decr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flyDBServiceClient) DecrBy(ctx context.Context, in *DecrByRequest, opts ...grpc.CallOption) (*DecrByResponse, error) {
	out := new(DecrByResponse)
	err := c.cc.Invoke(ctx, "/dbs.FlyDBService/DecrBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flyDBServiceClient) Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error) {
	out := new(ExistsResponse)
	err := c.cc.Invoke(ctx, "/dbs.FlyDBService/Exists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flyDBServiceClient) Expire(ctx context.Context, in *ExpireRequest, opts ...grpc.CallOption) (*ExpireResponse, error) {
	out := new(ExpireResponse)
	err := c.cc.Invoke(ctx, "/dbs.FlyDBService/Expire", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flyDBServiceClient) Persist(ctx context.Context, in *PersistRequest, opts ...grpc.CallOption) (*PersistResponse, error) {
	out := new(PersistResponse)
	err := c.cc.Invoke(ctx, "/dbs.FlyDBService/Persist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlyDBServiceServer is the server API for FlyDBService service.
// All implementations must embed UnimplementedFlyDBServiceServer
// for forward compatibility
type FlyDBServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Put(context.Context, *SetRequest) (*SetResponse, error)
	Del(context.Context, *DelRequest) (*DelResponse, error)
	Type(context.Context, *TypeRequest) (*TypeResponse, error)
	StrLen(context.Context, *StrLenRequest) (*StrLenResponse, error)
	GetSet(context.Context, *GetSetRequest) (*GetSetResponse, error)
	Append(context.Context, *AppendRequest) (*AppendResponse, error)
	Incr(context.Context, *IncrRequest) (*IncrResponse, error)
	IncrBy(context.Context, *IncrByRequest) (*IncrByResponse, error)
	IncrByFloat(context.Context, *IncrByFloatRequest) (*IncrByFloatResponse, error)
	Decr(context.Context, *DecrRequest) (*DecrResponse, error)
	DecrBy(context.Context, *DecrByRequest) (*DecrByResponse, error)
	Exists(context.Context, *ExistsRequest) (*ExistsResponse, error)
	Expire(context.Context, *ExpireRequest) (*ExpireResponse, error)
	Persist(context.Context, *PersistRequest) (*PersistResponse, error)
	mustEmbedUnimplementedFlyDBServiceServer()
}

// UnimplementedFlyDBServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlyDBServiceServer struct {
}

func (UnimplementedFlyDBServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFlyDBServiceServer) Put(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedFlyDBServiceServer) Del(context.Context, *DelRequest) (*DelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (UnimplementedFlyDBServiceServer) Type(context.Context, *TypeRequest) (*TypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Type not implemented")
}
func (UnimplementedFlyDBServiceServer) StrLen(context.Context, *StrLenRequest) (*StrLenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StrLen not implemented")
}
func (UnimplementedFlyDBServiceServer) GetSet(context.Context, *GetSetRequest) (*GetSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSet not implemented")
}
func (UnimplementedFlyDBServiceServer) Append(context.Context, *AppendRequest) (*AppendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Append not implemented")
}
func (UnimplementedFlyDBServiceServer) Incr(context.Context, *IncrRequest) (*IncrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Incr not implemented")
}
func (UnimplementedFlyDBServiceServer) IncrBy(context.Context, *IncrByRequest) (*IncrByResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrBy not implemented")
}
func (UnimplementedFlyDBServiceServer) IncrByFloat(context.Context, *IncrByFloatRequest) (*IncrByFloatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrByFloat not implemented")
}
func (UnimplementedFlyDBServiceServer) Decr(context.Context, *DecrRequest) (*DecrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decr not implemented")
}
func (UnimplementedFlyDBServiceServer) DecrBy(context.Context, *DecrByRequest) (*DecrByResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecrBy not implemented")
}
func (UnimplementedFlyDBServiceServer) Exists(context.Context, *ExistsRequest) (*ExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exists not implemented")
}
func (UnimplementedFlyDBServiceServer) Expire(context.Context, *ExpireRequest) (*ExpireResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Expire not implemented")
}
func (UnimplementedFlyDBServiceServer) Persist(context.Context, *PersistRequest) (*PersistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Persist not implemented")
}
func (UnimplementedFlyDBServiceServer) mustEmbedUnimplementedFlyDBServiceServer() {}

// UnsafeFlyDBServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlyDBServiceServer will
// result in compilation errors.
type UnsafeFlyDBServiceServer interface {
	mustEmbedUnimplementedFlyDBServiceServer()
}

func RegisterFlyDBServiceServer(s grpc.ServiceRegistrar, srv FlyDBServiceServer) {
	s.RegisterService(&FlyDBService_ServiceDesc, srv)
}

func _FlyDBService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyDBServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.FlyDBService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyDBServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlyDBService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyDBServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.FlyDBService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyDBServiceServer).Put(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlyDBService_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyDBServiceServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.FlyDBService/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyDBServiceServer).Del(ctx, req.(*DelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlyDBService_Type_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyDBServiceServer).Type(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.FlyDBService/Type",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyDBServiceServer).Type(ctx, req.(*TypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlyDBService_StrLen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StrLenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyDBServiceServer).StrLen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.FlyDBService/StrLen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyDBServiceServer).StrLen(ctx, req.(*StrLenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlyDBService_GetSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyDBServiceServer).GetSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.FlyDBService/GetSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyDBServiceServer).GetSet(ctx, req.(*GetSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlyDBService_Append_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyDBServiceServer).Append(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.FlyDBService/Append",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyDBServiceServer).Append(ctx, req.(*AppendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlyDBService_Incr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyDBServiceServer).Incr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.FlyDBService/Incr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyDBServiceServer).Incr(ctx, req.(*IncrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlyDBService_IncrBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrByRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyDBServiceServer).IncrBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.FlyDBService/IncrBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyDBServiceServer).IncrBy(ctx, req.(*IncrByRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlyDBService_IncrByFloat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrByFloatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyDBServiceServer).IncrByFloat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.FlyDBService/IncrByFloat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyDBServiceServer).IncrByFloat(ctx, req.(*IncrByFloatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlyDBService_Decr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyDBServiceServer).Decr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.FlyDBService/Decr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyDBServiceServer).Decr(ctx, req.(*DecrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlyDBService_DecrBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecrByRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyDBServiceServer).DecrBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.FlyDBService/DecrBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyDBServiceServer).DecrBy(ctx, req.(*DecrByRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlyDBService_Exists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyDBServiceServer).Exists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.FlyDBService/Exists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyDBServiceServer).Exists(ctx, req.(*ExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlyDBService_Expire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpireRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyDBServiceServer).Expire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.FlyDBService/Expire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyDBServiceServer).Expire(ctx, req.(*ExpireRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlyDBService_Persist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyDBServiceServer).Persist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.FlyDBService/Persist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyDBServiceServer).Persist(ctx, req.(*PersistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FlyDBService_ServiceDesc is the grpc.ServiceDesc for FlyDBService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlyDBService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dbs.FlyDBService",
	HandlerType: (*FlyDBServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _FlyDBService_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _FlyDBService_Put_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _FlyDBService_Del_Handler,
		},
		{
			MethodName: "Type",
			Handler:    _FlyDBService_Type_Handler,
		},
		{
			MethodName: "StrLen",
			Handler:    _FlyDBService_StrLen_Handler,
		},
		{
			MethodName: "GetSet",
			Handler:    _FlyDBService_GetSet_Handler,
		},
		{
			MethodName: "Append",
			Handler:    _FlyDBService_Append_Handler,
		},
		{
			MethodName: "Incr",
			Handler:    _FlyDBService_Incr_Handler,
		},
		{
			MethodName: "IncrBy",
			Handler:    _FlyDBService_IncrBy_Handler,
		},
		{
			MethodName: "IncrByFloat",
			Handler:    _FlyDBService_IncrByFloat_Handler,
		},
		{
			MethodName: "Decr",
			Handler:    _FlyDBService_Decr_Handler,
		},
		{
			MethodName: "DecrBy",
			Handler:    _FlyDBService_DecrBy_Handler,
		},
		{
			MethodName: "Exists",
			Handler:    _FlyDBService_Exists_Handler,
		},
		{
			MethodName: "Expire",
			Handler:    _FlyDBService_Expire_Handler,
		},
		{
			MethodName: "Persist",
			Handler:    _FlyDBService_Persist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lib/proto/dbs/db.proto",
}
