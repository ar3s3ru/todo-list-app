package command

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/ar3s3ru/todo-list-app/internal/domain/todolist"
	"github.com/ar3s3ru/todo-list-app/lib/ddd"
)

type CreateTodoList struct {
	ID    uuid.UUID
	Title string
	Owner string
}

var _ ddd.CommandHandler[CreateTodoList] = CreateTodoListHandler{}

type CreateTodoListHandler struct {
	Clock      func() time.Time
	Repository todolist.Adder
}

func (h CreateTodoListHandler) Handle(ctx context.Context, cmd CreateTodoList) error {
	now := h.Clock()

	todoList, err := todolist.Create(cmd.ID, cmd.Title, cmd.Owner, now)
	if err != nil {
		return fmt.Errorf("command.CreateTodoList: failed to create todolist, %w", err)
	}

	if err := h.Repository.Add(ctx, todoList); err != nil {
		return fmt.Errorf("command.CreateTodoList: failed to add todolist to repository, %w", err)
	}

	return nil
}
