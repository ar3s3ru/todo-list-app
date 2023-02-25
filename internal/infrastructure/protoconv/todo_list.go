package protoconv

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	todolistv1 "bitbucket.org/chronomics/todo-list-app/gen/todolist/v1"
	"bitbucket.org/chronomics/todo-list-app/internal/domain/todolist"
)

func FromTodoList(todoList *todolist.TodoList) *todolistv1.TodoList {
	if todoList == nil {
		return nil
	}

	pb := &todolistv1.TodoList{
		Title: todoList.Title,
		Owner: todoList.Owner,
	}

	if t := todoList.CreationTime; !t.IsZero() {
		pb.CreationTime = timestamppb.New(t)
	}

	if todoList.Items != nil {
		pb.Items = map[string]*todolistv1.TodoItem{}
	}

	for id, item := range todoList.Items {
		pb.Items[id.String()] = FromTodoItem(*item)
	}

	return pb
}
