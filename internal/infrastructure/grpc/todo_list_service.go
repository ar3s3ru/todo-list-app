package grpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	connectgo "github.com/bufbuild/connect-go"
	"github.com/google/uuid"

	v1 "github.com/ar3s3ru/todo-list-app/gen/todolist/v1"
	"github.com/ar3s3ru/todo-list-app/gen/todolist/v1/todolistv1connect"
	"github.com/ar3s3ru/todo-list-app/internal/command"
	"github.com/ar3s3ru/todo-list-app/internal/domain/todolist"
	"github.com/ar3s3ru/todo-list-app/internal/query"
	"github.com/ar3s3ru/todo-list-app/lib/ddd"
)

//nolint:exhaustruct // This is just a guard to ensure the interface is implemented.
var _ todolistv1connect.TodoListServiceHandler = &TodoListService{}

type TodoListService struct {
	GenerateIDFunc func() uuid.UUID

	query.GetTodoListHandler

	command.CreateTodoListHandler
	command.AddItemToTodoListHandler
	command.ToggleTodoItemHandler
	command.DeleteTodoItemHandler
}

// CreateTodoList implements todolistv1connect.TodoListServiceHandler.
func (srv *TodoListService) CreateTodoList(
	ctx context.Context,
	req *connectgo.Request[v1.CreateTodoListRequest],
) (*connectgo.Response[v1.CreateTodoListResponse], error) {
	cmd := command.CreateTodoList{
		ID:    srv.GenerateIDFunc(),
		Title: req.Msg.Title,
		Owner: req.Msg.Owner,
	}

	switch err := srv.CreateTodoListHandler.Handle(ctx, cmd); {
	case err == nil:
		return connectgo.NewResponse(&v1.CreateTodoListResponse{
			TodoListId: cmd.ID.String(),
		}), nil
	case errors.Is(err, todolist.ErrEmptyTitle), errors.Is(err, todolist.ErrNoOwner):
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, err)
	default:
		return nil, connectgo.NewError(connectgo.CodeInternal, err)
	}
}

// AddTodoItem implements todolistv1connect.TodoListServiceHandler.
func (srv *TodoListService) AddTodoItem(
	ctx context.Context,
	req *connectgo.Request[v1.AddTodoItemRequest],
) (*connectgo.Response[v1.AddTodoItemResponse], error) {
	todoListID, err := uuid.Parse(req.Msg.TodoListId)
	if err != nil {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, fmt.Errorf("grpc.TodoListService: failed to parse todoListId, %v", err))
	}

	cmd := command.AddItemToTodoList{
		TodoListID:  todoListID,
		TodoItemID:  srv.GenerateIDFunc(),
		Description: req.Msg.Description,
		DueDate:     time.Time{},
	}

	if req.Msg.DueDate != nil {
		cmd.DueDate = req.Msg.DueDate.AsTime()
	}

	switch err := srv.AddItemToTodoListHandler.Handle(ctx, cmd); {
	case err == nil:
		return connectgo.NewResponse(&v1.AddTodoItemResponse{
			TodoItemId: cmd.TodoItemID.String(),
		}), nil
	case errors.Is(err, todolist.ErrEmptyItemDescription):
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, err)
	default:
		return nil, connectgo.NewError(connectgo.CodeInternal, err)
	}
}

// ToggleTodoItem implements todolistv1connect.TodoListServiceHandler.
func (srv *TodoListService) ToggleTodoItem(
	ctx context.Context,
	req *connectgo.Request[v1.ToggleTodoItemRequest],
) (*connectgo.Response[v1.ToggleTodoItemResponse], error) {
	todoListID, err := uuid.Parse(req.Msg.TodoListId)
	if err != nil {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, fmt.Errorf("grpc.TodoListService: failed to parse todoListId, %v", err))
	}

	todoItemID, err := uuid.Parse(req.Msg.TodoItemId)
	if err != nil {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, fmt.Errorf("grpc.TodoListService: failed to parse todoItemId, %v", err))
	}

	cmd := command.ToggleTodoItem{
		TodoListID: todoListID,
		TodoItemID: todoItemID,
	}

	switch err := srv.ToggleTodoItemHandler.Handle(ctx, cmd); {
	case err == nil:
		return connectgo.NewResponse(&v1.ToggleTodoItemResponse{}), nil
	default:
		return nil, connectgo.NewError(connectgo.CodeInternal, err)
	}
}

// DeleteTodoItem implements todolistv1connect.TodoListServiceHandler.
func (srv *TodoListService) DeleteTodoItem(
	ctx context.Context,
	req *connectgo.Request[v1.DeleteTodoItemRequest],
) (*connectgo.Response[v1.DeleteTodoItemResponse], error) {
	todoListID, err := uuid.Parse(req.Msg.TodoListId)
	if err != nil {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, fmt.Errorf("grpc.TodoListService: failed to parse todoListId, %v", err))
	}

	todoItemID, err := uuid.Parse(req.Msg.TodoItemId)
	if err != nil {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, fmt.Errorf("grpc.TodoListService: failed to parse todoItemId, %v", err))
	}

	cmd := command.DeleteTodoItem{
		TodoListID: todoListID,
		TodoItemID: todoItemID,
	}

	switch err := srv.DeleteTodoItemHandler.Handle(ctx, cmd); {
	case err == nil:
		return connectgo.NewResponse(&v1.DeleteTodoItemResponse{}), nil
	case errors.Is(err, todolist.ErrItemNotFound):
		return nil, connectgo.NewError(connectgo.CodeNotFound, err)
	default:
		return nil, connectgo.NewError(connectgo.CodeInternal, err)
	}
}

// GetTodoList implements todolistv1connect.TodoListServiceHandler.
func (srv *TodoListService) GetTodoList(
	ctx context.Context,
	req *connectgo.Request[v1.GetTodoListRequest],
) (*connectgo.Response[v1.GetTodoListResponse], error) {
	todoListID, err := uuid.Parse(req.Msg.TodoListId)
	if err != nil {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, fmt.Errorf("grpc.TodoListService: failed to parse todoListId, %v", err))
	}

	todoList, err := srv.GetTodoListHandler.Handle(ctx, query.GetTodoList{
		TodoListID: todoListID,
	})

	switch {
	case err == nil:
		return connectgo.NewResponse(&v1.GetTodoListResponse{
			TodoList: todoList,
		}), nil
	case errors.Is(err, ddd.ErrEntityNotFound):
		return nil, connectgo.NewError(connectgo.CodeNotFound, err)
	default:
		return nil, connectgo.NewError(connectgo.CodeInternal, err)
	}
}
