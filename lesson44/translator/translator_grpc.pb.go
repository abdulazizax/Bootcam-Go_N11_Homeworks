// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: translator/translator.proto

package translator

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Translator_Translate_FullMethodName = "/translator.Translator/Translate"
)

// TranslatorClient is the client API for Translator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranslatorClient interface {
	Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error)
}

type translatorClient struct {
	cc grpc.ClientConnInterface
}

func NewTranslatorClient(cc grpc.ClientConnInterface) TranslatorClient {
	return &translatorClient{cc}
}

func (c *translatorClient) Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TranslateResponse)
	err := c.cc.Invoke(ctx, Translator_Translate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslatorServer is the server API for Translator service.
// All implementations must embed UnimplementedTranslatorServer
// for forward compatibility
type TranslatorServer interface {
	Translate(context.Context, *TranslateRequest) (*TranslateResponse, error)
	mustEmbedUnimplementedTranslatorServer()
}

// UnimplementedTranslatorServer must be embedded to have forward compatible implementations.
type UnimplementedTranslatorServer struct {
}

func (UnimplementedTranslatorServer) Translate(context.Context, *TranslateRequest) (*TranslateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Translate not implemented")
}
func (UnimplementedTranslatorServer) mustEmbedUnimplementedTranslatorServer() {}

// UnsafeTranslatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranslatorServer will
// result in compilation errors.
type UnsafeTranslatorServer interface {
	mustEmbedUnimplementedTranslatorServer()
}

func RegisterTranslatorServer(s grpc.ServiceRegistrar, srv TranslatorServer) {
	s.RegisterService(&Translator_ServiceDesc, srv)
}

func _Translator_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Translator_Translate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).Translate(ctx, req.(*TranslateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Translator_ServiceDesc is the grpc.ServiceDesc for Translator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Translator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "translator.Translator",
	HandlerType: (*TranslatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Translate",
			Handler:    _Translator_Translate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "translator/translator.proto",
}
