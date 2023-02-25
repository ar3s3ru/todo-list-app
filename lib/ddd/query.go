package ddd

import "context"

type QueryHandler[T any, R any] interface {
	Handle(ctx context.Context, query T) (R, error)
}
