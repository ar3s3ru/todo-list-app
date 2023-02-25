package command

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/ar3s3ru/todo-list-app/internal/domain/todolist"
	"github.com/ar3s3ru/todo-list-app/lib/ddd"
)

type AddItemToTodoList struct {
	TodoListID  uuid.UUID
	TodoItemID  uuid.UUID
	Description string
	DueDate     time.Time
}

//nolint:exhaustruct // This is just a guard to ensure the interface is implemented.
var _ ddd.CommandHandler[CreateTodoList] = CreateTodoListHandler{}

type AddItemToTodoListHandler struct {
	Clock      func() time.Time
	Repository todolist.Repository
}

func (h AddItemToTodoListHandler) Handle(ctx context.Context, cmd AddItemToTodoList) error {
	now := h.Clock()

	todoList, err := h.Repository.Get(ctx, cmd.TodoListID)
	if err != nil {
		return fmt.Errorf("command.AddItemToTodoList: failed to get todolist from repository, %w", err)
	}

	if err := todoList.AddTodoItem(cmd.TodoItemID, cmd.Description, cmd.DueDate, now); err != nil {
		return fmt.Errorf("command.AddItemToTodoList: failed to add item to todolist, %w", err)
	}

	if err := h.Repository.Add(ctx, todoList); err != nil {
		return fmt.Errorf("command.AddItemToTodoList: failed to add todolist to repository, %w", err)
	}

	return nil
}
