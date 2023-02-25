package protoconv

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	todolistv1 "bitbucket.org/chronomics/todo-list-app/gen/todolist/v1"
	"bitbucket.org/chronomics/todo-list-app/internal/domain/todolist"
)

func FromTodoItem(todoItem todolist.TodoItem) *todolistv1.TodoItem {
	pb := &todolistv1.TodoItem{
		Description: todoItem.Description,
		Completed:   todoItem.Completed,
	}

	if t := todoItem.CreationTime; !t.IsZero() {
		pb.CreationTime = timestamppb.New(t)
	}

	if t := todoItem.DueDate; !t.IsZero() {
		pb.DueDate = timestamppb.New(t)
	}

	return pb
}