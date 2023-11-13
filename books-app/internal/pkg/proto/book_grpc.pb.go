// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: books-app/internal/pkg/proto/book.proto

package proto

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

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	// Unary
	AddBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*AddBookResponse, error)
	ListBooks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListBooksRespose, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	RemoveBook(ctx context.Context, in *RemoveBookRequest, opts ...grpc.CallOption) (*RemoveBookResponse, error)
	UpdateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*UpdateBookResponse, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) AddBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*AddBookResponse, error) {
	out := new(AddBookResponse)
	err := c.cc.Invoke(ctx, "/prot.BookService/AddBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) ListBooks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListBooksRespose, error) {
	out := new(ListBooksRespose)
	err := c.cc.Invoke(ctx, "/prot.BookService/ListBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/prot.BookService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) RemoveBook(ctx context.Context, in *RemoveBookRequest, opts ...grpc.CallOption) (*RemoveBookResponse, error) {
	out := new(RemoveBookResponse)
	err := c.cc.Invoke(ctx, "/prot.BookService/RemoveBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) UpdateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*UpdateBookResponse, error) {
	out := new(UpdateBookResponse)
	err := c.cc.Invoke(ctx, "/prot.BookService/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility
type BookServiceServer interface {
	// Unary
	AddBook(context.Context, *Book) (*AddBookResponse, error)
	ListBooks(context.Context, *Empty) (*ListBooksRespose, error)
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	RemoveBook(context.Context, *RemoveBookRequest) (*RemoveBookResponse, error)
	UpdateBook(context.Context, *Book) (*UpdateBookResponse, error)
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (UnimplementedBookServiceServer) AddBook(context.Context, *Book) (*AddBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (UnimplementedBookServiceServer) ListBooks(context.Context, *Empty) (*ListBooksRespose, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (UnimplementedBookServiceServer) GetBook(context.Context, *GetBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookServiceServer) RemoveBook(context.Context, *RemoveBookRequest) (*RemoveBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBook not implemented")
}
func (UnimplementedBookServiceServer) UpdateBook(context.Context, *Book) (*UpdateBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prot.BookService/AddBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).AddBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prot.BookService/ListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).ListBooks(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prot.BookService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_RemoveBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).RemoveBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prot.BookService/RemoveBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).RemoveBook(ctx, req.(*RemoveBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prot.BookService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).UpdateBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prot.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBook",
			Handler:    _BookService_AddBook_Handler,
		},
		{
			MethodName: "ListBooks",
			Handler:    _BookService_ListBooks_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BookService_GetBook_Handler,
		},
		{
			MethodName: "RemoveBook",
			Handler:    _BookService_RemoveBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BookService_UpdateBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "books-app/internal/pkg/proto/book.proto",
}

// ReviewServiceClient is the client API for ReviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewServiceClient interface {
	// Unary RPC to submit a book review
	SubmitReviews(ctx context.Context, in *SubmitReviewRequest, opts ...grpc.CallOption) (*SubmitReviewResponse, error)
}

type reviewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewServiceClient(cc grpc.ClientConnInterface) ReviewServiceClient {
	return &reviewServiceClient{cc}
}

func (c *reviewServiceClient) SubmitReviews(ctx context.Context, in *SubmitReviewRequest, opts ...grpc.CallOption) (*SubmitReviewResponse, error) {
	out := new(SubmitReviewResponse)
	err := c.cc.Invoke(ctx, "/prot.ReviewService/SubmitReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServiceServer is the server API for ReviewService service.
// All implementations must embed UnimplementedReviewServiceServer
// for forward compatibility
type ReviewServiceServer interface {
	// Unary RPC to submit a book review
	SubmitReviews(context.Context, *SubmitReviewRequest) (*SubmitReviewResponse, error)
	mustEmbedUnimplementedReviewServiceServer()
}

// UnimplementedReviewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReviewServiceServer struct {
}

func (UnimplementedReviewServiceServer) SubmitReviews(context.Context, *SubmitReviewRequest) (*SubmitReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitReviews not implemented")
}
func (UnimplementedReviewServiceServer) mustEmbedUnimplementedReviewServiceServer() {}

// UnsafeReviewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewServiceServer will
// result in compilation errors.
type UnsafeReviewServiceServer interface {
	mustEmbedUnimplementedReviewServiceServer()
}

func RegisterReviewServiceServer(s grpc.ServiceRegistrar, srv ReviewServiceServer) {
	s.RegisterService(&ReviewService_ServiceDesc, srv)
}

func _ReviewService_SubmitReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).SubmitReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prot.ReviewService/SubmitReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).SubmitReviews(ctx, req.(*SubmitReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReviewService_ServiceDesc is the grpc.ServiceDesc for ReviewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReviewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prot.ReviewService",
	HandlerType: (*ReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitReviews",
			Handler:    _ReviewService_SubmitReviews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "books-app/internal/pkg/proto/book.proto",
}

// BookInfoServiceClient is the client API for BookInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookInfoServiceClient interface {
	// Unary RPC to get book information with reviews by ID
	GetBookInfoWithReviews(ctx context.Context, in *GetBookInfoRequest, opts ...grpc.CallOption) (*GetBookInfoResponse, error)
}

type bookInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookInfoServiceClient(cc grpc.ClientConnInterface) BookInfoServiceClient {
	return &bookInfoServiceClient{cc}
}

func (c *bookInfoServiceClient) GetBookInfoWithReviews(ctx context.Context, in *GetBookInfoRequest, opts ...grpc.CallOption) (*GetBookInfoResponse, error) {
	out := new(GetBookInfoResponse)
	err := c.cc.Invoke(ctx, "/prot.BookInfoService/GetBookInfoWithReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookInfoServiceServer is the server API for BookInfoService service.
// All implementations must embed UnimplementedBookInfoServiceServer
// for forward compatibility
type BookInfoServiceServer interface {
	// Unary RPC to get book information with reviews by ID
	GetBookInfoWithReviews(context.Context, *GetBookInfoRequest) (*GetBookInfoResponse, error)
	mustEmbedUnimplementedBookInfoServiceServer()
}

// UnimplementedBookInfoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookInfoServiceServer struct {
}

func (UnimplementedBookInfoServiceServer) GetBookInfoWithReviews(context.Context, *GetBookInfoRequest) (*GetBookInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookInfoWithReviews not implemented")
}
func (UnimplementedBookInfoServiceServer) mustEmbedUnimplementedBookInfoServiceServer() {}

// UnsafeBookInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookInfoServiceServer will
// result in compilation errors.
type UnsafeBookInfoServiceServer interface {
	mustEmbedUnimplementedBookInfoServiceServer()
}

func RegisterBookInfoServiceServer(s grpc.ServiceRegistrar, srv BookInfoServiceServer) {
	s.RegisterService(&BookInfoService_ServiceDesc, srv)
}

func _BookInfoService_GetBookInfoWithReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookInfoServiceServer).GetBookInfoWithReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prot.BookInfoService/GetBookInfoWithReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookInfoServiceServer).GetBookInfoWithReviews(ctx, req.(*GetBookInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookInfoService_ServiceDesc is the grpc.ServiceDesc for BookInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prot.BookInfoService",
	HandlerType: (*BookInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookInfoWithReviews",
			Handler:    _BookInfoService_GetBookInfoWithReviews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "books-app/internal/pkg/proto/book.proto",
}
