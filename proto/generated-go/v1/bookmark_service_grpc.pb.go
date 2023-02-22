// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1/bookmark_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookmarkServiceClient is the client API for BookmarkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookmarkServiceClient interface {
	// CreateBookmark creates a new bookmark.
	CreateBookmark(ctx context.Context, in *CreateBookmarkRequest, opts ...grpc.CallOption) (*Bookmark, error)
	// DeleteBookmark deletes a bookmark.
	DeleteBookmark(ctx context.Context, in *DeleteBookmarkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListBookmark lists bookmarks.
	ListBookmarks(ctx context.Context, in *ListBookmarksRequest, opts ...grpc.CallOption) (*ListBookmarksResponse, error)
}

type bookmarkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookmarkServiceClient(cc grpc.ClientConnInterface) BookmarkServiceClient {
	return &bookmarkServiceClient{cc}
}

func (c *bookmarkServiceClient) CreateBookmark(ctx context.Context, in *CreateBookmarkRequest, opts ...grpc.CallOption) (*Bookmark, error) {
	out := new(Bookmark)
	err := c.cc.Invoke(ctx, "/bytebase.v1.BookmarkService/CreateBookmark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookmarkServiceClient) DeleteBookmark(ctx context.Context, in *DeleteBookmarkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/bytebase.v1.BookmarkService/DeleteBookmark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookmarkServiceClient) ListBookmarks(ctx context.Context, in *ListBookmarksRequest, opts ...grpc.CallOption) (*ListBookmarksResponse, error) {
	out := new(ListBookmarksResponse)
	err := c.cc.Invoke(ctx, "/bytebase.v1.BookmarkService/ListBookmarks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookmarkServiceServer is the server API for BookmarkService service.
// All implementations must embed UnimplementedBookmarkServiceServer
// for forward compatibility
type BookmarkServiceServer interface {
	// CreateBookmark creates a new bookmark.
	CreateBookmark(context.Context, *CreateBookmarkRequest) (*Bookmark, error)
	// DeleteBookmark deletes a bookmark.
	DeleteBookmark(context.Context, *DeleteBookmarkRequest) (*emptypb.Empty, error)
	// ListBookmark lists bookmarks.
	ListBookmarks(context.Context, *ListBookmarksRequest) (*ListBookmarksResponse, error)
	mustEmbedUnimplementedBookmarkServiceServer()
}

// UnimplementedBookmarkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookmarkServiceServer struct {
}

func (UnimplementedBookmarkServiceServer) CreateBookmark(context.Context, *CreateBookmarkRequest) (*Bookmark, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBookmark not implemented")
}
func (UnimplementedBookmarkServiceServer) DeleteBookmark(context.Context, *DeleteBookmarkRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBookmark not implemented")
}
func (UnimplementedBookmarkServiceServer) ListBookmarks(context.Context, *ListBookmarksRequest) (*ListBookmarksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBookmarks not implemented")
}
func (UnimplementedBookmarkServiceServer) mustEmbedUnimplementedBookmarkServiceServer() {}

// UnsafeBookmarkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookmarkServiceServer will
// result in compilation errors.
type UnsafeBookmarkServiceServer interface {
	mustEmbedUnimplementedBookmarkServiceServer()
}

func RegisterBookmarkServiceServer(s grpc.ServiceRegistrar, srv BookmarkServiceServer) {
	s.RegisterService(&BookmarkService_ServiceDesc, srv)
}

func _BookmarkService_CreateBookmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookmarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookmarkServiceServer).CreateBookmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bytebase.v1.BookmarkService/CreateBookmark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookmarkServiceServer).CreateBookmark(ctx, req.(*CreateBookmarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookmarkService_DeleteBookmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookmarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookmarkServiceServer).DeleteBookmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bytebase.v1.BookmarkService/DeleteBookmark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookmarkServiceServer).DeleteBookmark(ctx, req.(*DeleteBookmarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookmarkService_ListBookmarks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBookmarksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookmarkServiceServer).ListBookmarks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bytebase.v1.BookmarkService/ListBookmarks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookmarkServiceServer).ListBookmarks(ctx, req.(*ListBookmarksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookmarkService_ServiceDesc is the grpc.ServiceDesc for BookmarkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookmarkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bytebase.v1.BookmarkService",
	HandlerType: (*BookmarkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBookmark",
			Handler:    _BookmarkService_CreateBookmark_Handler,
		},
		{
			MethodName: "DeleteBookmark",
			Handler:    _BookmarkService_DeleteBookmark_Handler,
		},
		{
			MethodName: "ListBookmarks",
			Handler:    _BookmarkService_ListBookmarks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/bookmark_service.proto",
}
