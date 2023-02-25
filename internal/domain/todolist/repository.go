package todolist

import (
	"github.com/google/uuid"

	"github.com/ar3s3ru/todo-list-app/lib/ddd"
)

type (
	Getter     = ddd.EntityGetter[uuid.UUID, *TodoList]
	Adder      = ddd.EntityAdder[uuid.UUID, *TodoList]
	Repository = ddd.EntityRepository[uuid.UUID, *TodoList]
)
