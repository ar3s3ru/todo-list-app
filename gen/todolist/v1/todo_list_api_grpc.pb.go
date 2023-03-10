// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package todolistv1

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

// TodoListServiceClient is the client API for TodoListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoListServiceClient interface {
	CreateTodoList(ctx context.Context, in *CreateTodoListRequest, opts ...grpc.CallOption) (*CreateTodoListResponse, error)
	GetTodoList(ctx context.Context, in *GetTodoListRequest, opts ...grpc.CallOption) (*GetTodoListResponse, error)
	AddTodoItem(ctx context.Context, in *AddTodoItemRequest, opts ...grpc.CallOption) (*AddTodoItemResponse, error)
	ToggleTodoItem(ctx context.Context, in *ToggleTodoItemRequest, opts ...grpc.CallOption) (*ToggleTodoItemResponse, error)
	DeleteTodoItem(ctx context.Context, in *DeleteTodoItemRequest, opts ...grpc.CallOption) (*DeleteTodoItemResponse, error)
}

type todoListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoListServiceClient(cc grpc.ClientConnInterface) TodoListServiceClient {
	return &todoListServiceClient{cc}
}

func (c *todoListServiceClient) CreateTodoList(ctx context.Context, in *CreateTodoListRequest, opts ...grpc.CallOption) (*CreateTodoListResponse, error) {
	out := new(CreateTodoListResponse)
	err := c.cc.Invoke(ctx, "/todolist.v1.TodoListService/CreateTodoList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListServiceClient) GetTodoList(ctx context.Context, in *GetTodoListRequest, opts ...grpc.CallOption) (*GetTodoListResponse, error) {
	out := new(GetTodoListResponse)
	err := c.cc.Invoke(ctx, "/todolist.v1.TodoListService/GetTodoList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListServiceClient) AddTodoItem(ctx context.Context, in *AddTodoItemRequest, opts ...grpc.CallOption) (*AddTodoItemResponse, error) {
	out := new(AddTodoItemResponse)
	err := c.cc.Invoke(ctx, "/todolist.v1.TodoListService/AddTodoItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListServiceClient) ToggleTodoItem(ctx context.Context, in *ToggleTodoItemRequest, opts ...grpc.CallOption) (*ToggleTodoItemResponse, error) {
	out := new(ToggleTodoItemResponse)
	err := c.cc.Invoke(ctx, "/todolist.v1.TodoListService/ToggleTodoItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListServiceClient) DeleteTodoItem(ctx context.Context, in *DeleteTodoItemRequest, opts ...grpc.CallOption) (*DeleteTodoItemResponse, error) {
	out := new(DeleteTodoItemResponse)
	err := c.cc.Invoke(ctx, "/todolist.v1.TodoListService/DeleteTodoItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoListServiceServer is the server API for TodoListService service.
// All implementations must embed UnimplementedTodoListServiceServer
// for forward compatibility
type TodoListServiceServer interface {
	CreateTodoList(context.Context, *CreateTodoListRequest) (*CreateTodoListResponse, error)
	GetTodoList(context.Context, *GetTodoListRequest) (*GetTodoListResponse, error)
	AddTodoItem(context.Context, *AddTodoItemRequest) (*AddTodoItemResponse, error)
	ToggleTodoItem(context.Context, *ToggleTodoItemRequest) (*ToggleTodoItemResponse, error)
	DeleteTodoItem(context.Context, *DeleteTodoItemRequest) (*DeleteTodoItemResponse, error)
	mustEmbedUnimplementedTodoListServiceServer()
}

// UnimplementedTodoListServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoListServiceServer struct {
}

func (UnimplementedTodoListServiceServer) CreateTodoList(context.Context, *CreateTodoListRequest) (*CreateTodoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodoList not implemented")
}
func (UnimplementedTodoListServiceServer) GetTodoList(context.Context, *GetTodoListRequest) (*GetTodoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoList not implemented")
}
func (UnimplementedTodoListServiceServer) AddTodoItem(context.Context, *AddTodoItemRequest) (*AddTodoItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTodoItem not implemented")
}
func (UnimplementedTodoListServiceServer) ToggleTodoItem(context.Context, *ToggleTodoItemRequest) (*ToggleTodoItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleTodoItem not implemented")
}
func (UnimplementedTodoListServiceServer) DeleteTodoItem(context.Context, *DeleteTodoItemRequest) (*DeleteTodoItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodoItem not implemented")
}
func (UnimplementedTodoListServiceServer) mustEmbedUnimplementedTodoListServiceServer() {}

// UnsafeTodoListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoListServiceServer will
// result in compilation errors.
type UnsafeTodoListServiceServer interface {
	mustEmbedUnimplementedTodoListServiceServer()
}

func RegisterTodoListServiceServer(s grpc.ServiceRegistrar, srv TodoListServiceServer) {
	s.RegisterService(&TodoListService_ServiceDesc, srv)
}

func _TodoListService_CreateTodoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).CreateTodoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.v1.TodoListService/CreateTodoList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).CreateTodoList(ctx, req.(*CreateTodoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoListService_GetTodoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).GetTodoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.v1.TodoListService/GetTodoList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).GetTodoList(ctx, req.(*GetTodoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoListService_AddTodoItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTodoItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).AddTodoItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.v1.TodoListService/AddTodoItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).AddTodoItem(ctx, req.(*AddTodoItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoListService_ToggleTodoItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleTodoItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).ToggleTodoItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.v1.TodoListService/ToggleTodoItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).ToggleTodoItem(ctx, req.(*ToggleTodoItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoListService_DeleteTodoItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).DeleteTodoItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.v1.TodoListService/DeleteTodoItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).DeleteTodoItem(ctx, req.(*DeleteTodoItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoListService_ServiceDesc is the grpc.ServiceDesc for TodoListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todolist.v1.TodoListService",
	HandlerType: (*TodoListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodoList",
			Handler:    _TodoListService_CreateTodoList_Handler,
		},
		{
			MethodName: "GetTodoList",
			Handler:    _TodoListService_GetTodoList_Handler,
		},
		{
			MethodName: "AddTodoItem",
			Handler:    _TodoListService_AddTodoItem_Handler,
		},
		{
			MethodName: "ToggleTodoItem",
			Handler:    _TodoListService_ToggleTodoItem_Handler,
		},
		{
			MethodName: "DeleteTodoItem",
			Handler:    _TodoListService_DeleteTodoItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todolist/v1/todo_list_api.proto",
}
