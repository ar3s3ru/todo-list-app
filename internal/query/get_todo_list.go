package query

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	todolistv1 "bitbucket.org/chronomics/todo-list-app/gen/todolist/v1"
	"bitbucket.org/chronomics/todo-list-app/internal/domain/todolist"
	"bitbucket.org/chronomics/todo-list-app/internal/infrastructure/protoconv"
	"bitbucket.org/chronomics/todo-list-app/lib/ddd"
)

type GetTodoList struct {
	TodoListID uuid.UUID
}

var _ ddd.QueryHandler[GetTodoList, *todolistv1.TodoList] = GetTodoListHandler{}

type GetTodoListHandler struct {
	Repository todolist.Getter
}

func (qh GetTodoListHandler) Handle(ctx context.Context, query GetTodoList) (*todolistv1.TodoList, error) {
	if query.TodoListID == uuid.Nil {
		return nil, fmt.Errorf("query.GetTodoList: invalid id provided, %w", todolist.ErrEmptyID)
	}

	todoList, err := qh.Repository.Get(ctx, query.TodoListID)
	if err != nil {
		return nil, fmt.Errorf("query.GetTodoList: failed to fetch TodoList from repository, %w", err)
	}

	return protoconv.FromTodoList(todoList), nil
}
