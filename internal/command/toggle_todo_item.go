package command

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"bitbucket.org/chronomics/todo-list-app/internal/domain/todolist"
	"bitbucket.org/chronomics/todo-list-app/lib/ddd"
)

type ToggleTodoItem struct {
	TodoListID uuid.UUID
	TodoItemID uuid.UUID
}

var _ ddd.CommandHandler[CreateTodoList] = CreateTodoListHandler{}

type ToggleTodoItemHandler struct {
	Repository todolist.Repository
}

func (h ToggleTodoItemHandler) Handle(ctx context.Context, cmd ToggleTodoItem) error {
	todoList, err := h.Repository.Get(ctx, cmd.TodoListID)
	if err != nil {
		return fmt.Errorf("command.ToggleTodoItem: failed to get todolist from repository, %w", err)
	}

	if err := todoList.ToggleTodoItem(cmd.TodoItemID); err != nil {
		return fmt.Errorf("command.ToggleTodoItem: failed to toggle todo item to todolist, %w", err)
	}

	if err := h.Repository.Add(ctx, todoList); err != nil {
		return fmt.Errorf("command.ToggleTodoItem: failed to add todolist to repository, %w", err)
	}

	return nil
}
