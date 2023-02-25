package command

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/ar3s3ru/todo-list-app/internal/domain/todolist"
	"github.com/ar3s3ru/todo-list-app/lib/ddd"
)

type DeleteTodoItem struct {
	TodoListID uuid.UUID
	TodoItemID uuid.UUID
}

//nolint:exhaustruct // This is just a guard to ensure the interface is implemented.
var _ ddd.CommandHandler[CreateTodoList] = CreateTodoListHandler{}

type DeleteTodoItemHandler struct {
	Repository todolist.Repository
}

func (h DeleteTodoItemHandler) Handle(ctx context.Context, cmd DeleteTodoItem) error {
	todoList, err := h.Repository.Get(ctx, cmd.TodoListID)
	if err != nil {
		return fmt.Errorf("command.DeleteTodoItem: failed to get todolist from repository, %w", err)
	}

	if err := todoList.DeleteTodoItem(cmd.TodoItemID); err != nil {
		return fmt.Errorf("command.DeleteTodoItem: failed to delete todo item from todolist, %w", err)
	}

	if err := h.Repository.Add(ctx, todoList); err != nil {
		return fmt.Errorf("command.DeleteTodoItem: failed to add todolist to repository, %w", err)
	}

	return nil
}
