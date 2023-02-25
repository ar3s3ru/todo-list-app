package protoconv

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	todolistv1 "github.com/ar3s3ru/todo-list-app/gen/todolist/v1"
	"github.com/ar3s3ru/todo-list-app/internal/domain/todolist"
)

func FromTodoItem(todoItem todolist.TodoItem) *todolistv1.TodoItem {
	pb := &todolistv1.TodoItem{
		Description:  todoItem.Description,
		Completed:    todoItem.Completed,
		DueDate:      nil,
		CreationTime: nil,
	}

	if t := todoItem.CreationTime; !t.IsZero() {
		pb.CreationTime = timestamppb.New(t)
	}

	if t := todoItem.DueDate; !t.IsZero() {
		pb.DueDate = timestamppb.New(t)
	}

	return pb
}
