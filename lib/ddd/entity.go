package ddd

import (
	"context"
	"errors"
)

var (
	ErrEntityNotFound = errors.New("ddd: entity not found")
)

type Entity[ID comparable] interface {
	ID() ID
}

type EntityAdder[ID comparable, T Entity[ID]] interface {
	Add(ctx context.Context, entity T) error
}

type EntityGetter[ID comparable, T Entity[ID]] interface {
	Get(ctx context.Context, id ID) (T, error)
}

type EntityRepository[ID comparable, T Entity[ID]] interface {
	EntityAdder[ID, T]
	EntityGetter[ID, T]
}
