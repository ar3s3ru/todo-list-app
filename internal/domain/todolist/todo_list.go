package todolist

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/ar3s3ru/todo-list-app/lib/ddd"
)

var (
	ErrEmptyID           = errors.New("todolist: id is empty")
	ErrEmptyTitle        = errors.New("todolist: title is empty")
	ErrNoOwner           = errors.New("todolist: no owner specified")
	ErrItemAlreadyExists = errors.New("todolist: item already exists")
	ErrEmptyItemID       = errors.New("todolist: item id is empty")
	ErrItemNotFound      = errors.New("todolist: item not found")
)

//nolint:exhaustruct // This is just a guard to ensure the interface is implemented.
var _ ddd.Entity[uuid.UUID] = &TodoList{}

type TodoList struct {
	TodoListID   uuid.UUID
	Title        string
	Owner        string
	Items        map[uuid.UUID]*TodoItem
	CreationTime time.Time
}

func (tl *TodoList) ID() uuid.UUID {
	return tl.TodoListID
}

func (tl *TodoList) guardItsValidState() error {
	if tl.TodoListID == uuid.Nil {
		return ErrEmptyID
	}

	if tl.Title == "" {
		return ErrEmptyTitle
	}

	if tl.Owner == "" {
		return ErrNoOwner
	}

	return nil
}

func Create(id uuid.UUID, title, owner string, now time.Time) (*TodoList, error) {
	todoList := TodoList{
		TodoListID:   id,
		Title:        title,
		Owner:        owner,
		Items:        nil, // To avoid unnecessary allocations.
		CreationTime: now,
	}

	if err := todoList.guardItsValidState(); err != nil {
		return nil, fmt.Errorf("todolist.Create: failed to create new list, %w", err)
	}

	return &todoList, nil
}

func (tl *TodoList) AddTodoItem(id uuid.UUID, description string, dueDate, now time.Time) error {
	wrapErr := func(err error) error {
		return fmt.Errorf("todolist.AddTodoItem: failed to add new item, %w", err)
	}

	if id == uuid.Nil {
		return wrapErr(ErrEmptyItemID)
	}

	if _, ok := tl.Items[id]; ok {
		return wrapErr(ErrItemAlreadyExists)
	}

	todoItem, err := CreateNewItem(description, dueDate, now)
	if err != nil {
		return wrapErr(err)
	}

	if tl.Items == nil {
		tl.Items = make(map[uuid.UUID]*TodoItem)
	}

	tl.Items[id] = todoItem

	return nil
}

func (tl *TodoList) ToggleTodoItem(id uuid.UUID) error {
	wrapErr := func(err error) error {
		return fmt.Errorf("todolist.ToggleTodoItem: failed to toggle todo item, %w", err)
	}

	if id == uuid.Nil {
		return wrapErr(ErrEmptyItemID)
	}

	item, ok := tl.Items[id]
	if !ok {
		return wrapErr(ErrItemNotFound)
	}

	item.Toggle()

	return nil
}

func (tl *TodoList) DeleteTodoItem(id uuid.UUID) error {
	wrapErr := func(err error) error {
		return fmt.Errorf("todolist.DeleteTodoItem: failed to delete todo item, %w", err)
	}

	if id == uuid.Nil {
		return wrapErr(ErrEmptyItemID)
	}

	if _, ok := tl.Items[id]; !ok {
		return wrapErr(ErrItemNotFound)
	}

	delete(tl.Items, id)

	return nil
}
