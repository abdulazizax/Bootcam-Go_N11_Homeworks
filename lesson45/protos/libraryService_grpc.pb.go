// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: protos/libraryService.proto

package Book

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
	LibraryService_SearchBookByID_FullMethodName    = "/LibraryService.LibraryService/SearchBookByID"
	LibraryService_SearchBookByTitle_FullMethodName = "/LibraryService.LibraryService/SearchBookByTitle"
	LibraryService_AddBook_FullMethodName           = "/LibraryService.LibraryService/AddBook"
	LibraryService_BorrowBook_FullMethodName        = "/LibraryService.LibraryService/BorrowBook"
)

// LibraryServiceClient is the client API for LibraryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibraryServiceClient interface {
	SearchBookByID(ctx context.Context, in *SearchBookRequest, opts ...grpc.CallOption) (*SearchBookResponse, error)
	SearchBookByTitle(ctx context.Context, in *SearchBookRequest, opts ...grpc.CallOption) (*SearchBookResponse, error)
	AddBook(ctx context.Context, in *AddBookRequest, opts ...grpc.CallOption) (*AddBookResponse, error)
	BorrowBook(ctx context.Context, in *BorrowBookRequest, opts ...grpc.CallOption) (*BorrowBookResponse, error)
}

type libraryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLibraryServiceClient(cc grpc.ClientConnInterface) LibraryServiceClient {
	return &libraryServiceClient{cc}
}

func (c *libraryServiceClient) SearchBookByID(ctx context.Context, in *SearchBookRequest, opts ...grpc.CallOption) (*SearchBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchBookResponse)
	err := c.cc.Invoke(ctx, LibraryService_SearchBookByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) SearchBookByTitle(ctx context.Context, in *SearchBookRequest, opts ...grpc.CallOption) (*SearchBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchBookResponse)
	err := c.cc.Invoke(ctx, LibraryService_SearchBookByTitle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) AddBook(ctx context.Context, in *AddBookRequest, opts ...grpc.CallOption) (*AddBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddBookResponse)
	err := c.cc.Invoke(ctx, LibraryService_AddBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) BorrowBook(ctx context.Context, in *BorrowBookRequest, opts ...grpc.CallOption) (*BorrowBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BorrowBookResponse)
	err := c.cc.Invoke(ctx, LibraryService_BorrowBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibraryServiceServer is the server API for LibraryService service.
// All implementations must embed UnimplementedLibraryServiceServer
// for forward compatibility
type LibraryServiceServer interface {
	SearchBookByID(context.Context, *SearchBookRequest) (*SearchBookResponse, error)
	SearchBookByTitle(context.Context, *SearchBookRequest) (*SearchBookResponse, error)
	AddBook(context.Context, *AddBookRequest) (*AddBookResponse, error)
	BorrowBook(context.Context, *BorrowBookRequest) (*BorrowBookResponse, error)
	mustEmbedUnimplementedLibraryServiceServer()
}

// UnimplementedLibraryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLibraryServiceServer struct {
}

func (UnimplementedLibraryServiceServer) SearchBookByID(context.Context, *SearchBookRequest) (*SearchBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBookByID not implemented")
}
func (UnimplementedLibraryServiceServer) SearchBookByTitle(context.Context, *SearchBookRequest) (*SearchBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBookByTitle not implemented")
}
func (UnimplementedLibraryServiceServer) AddBook(context.Context, *AddBookRequest) (*AddBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (UnimplementedLibraryServiceServer) BorrowBook(context.Context, *BorrowBookRequest) (*BorrowBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BorrowBook not implemented")
}
func (UnimplementedLibraryServiceServer) mustEmbedUnimplementedLibraryServiceServer() {}

// UnsafeLibraryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibraryServiceServer will
// result in compilation errors.
type UnsafeLibraryServiceServer interface {
	mustEmbedUnimplementedLibraryServiceServer()
}

func RegisterLibraryServiceServer(s grpc.ServiceRegistrar, srv LibraryServiceServer) {
	s.RegisterService(&LibraryService_ServiceDesc, srv)
}

func _LibraryService_SearchBookByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).SearchBookByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_SearchBookByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).SearchBookByID(ctx, req.(*SearchBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_SearchBookByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).SearchBookByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_SearchBookByTitle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).SearchBookByTitle(ctx, req.(*SearchBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_AddBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).AddBook(ctx, req.(*AddBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_BorrowBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BorrowBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).BorrowBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_BorrowBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).BorrowBook(ctx, req.(*BorrowBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LibraryService_ServiceDesc is the grpc.ServiceDesc for LibraryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LibraryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LibraryService.LibraryService",
	HandlerType: (*LibraryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchBookByID",
			Handler:    _LibraryService_SearchBookByID_Handler,
		},
		{
			MethodName: "SearchBookByTitle",
			Handler:    _LibraryService_SearchBookByTitle_Handler,
		},
		{
			MethodName: "AddBook",
			Handler:    _LibraryService_AddBook_Handler,
		},
		{
			MethodName: "BorrowBook",
			Handler:    _LibraryService_BorrowBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/libraryService.proto",
}
