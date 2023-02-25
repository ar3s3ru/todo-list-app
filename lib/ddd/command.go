package ddd

import "context"

type CommandHandler[T any] interface {
	Handle(ctx context.Context, cmd T) error
}
