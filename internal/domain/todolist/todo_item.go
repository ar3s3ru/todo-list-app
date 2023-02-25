package todolist

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrEmptyItemDescription = errors.New("todolist.TodoItem: description is empty")
)

type TodoItem struct {
	Description  string
	Completed    bool
	DueDate      time.Time
	CreationTime time.Time
}

func (ti TodoItem) guardItsValidState() error {
	if ti.Description == "" {
		return ErrEmptyItemDescription
	}

	return nil
}

func CreateNewItem(description string, dueDate, now time.Time) (*TodoItem, error) {
	ti := &TodoItem{
		Description:  description,
		Completed:    false,
		DueDate:      dueDate,
		CreationTime: now,
	}

	if err := ti.guardItsValidState(); err != nil {
		return nil, fmt.Errorf("todolist.CreateNewItem: invalid item state, %w", err)
	}

	return ti, nil
}

func (ti *TodoItem) Toggle() {
	ti.Completed = !ti.Completed
}

func (ti TodoItem) IsDue(now time.Time) bool {
	return !ti.DueDate.IsZero() && now.After(ti.DueDate)
}
