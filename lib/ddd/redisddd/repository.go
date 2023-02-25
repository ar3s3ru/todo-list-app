package redisddd

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/ar3s3ru/todo-list-app/lib/ddd"
)

type id interface {
	comparable
	fmt.Stringer
}

type Repository[ID id, T ddd.Entity[ID]] struct {
	Prefix string
	Client *redis.Client
}

func (r Repository[ID, T]) makeKey(id string) string {
	key := id

	if r.Prefix != "" {
		key = fmt.Sprintf("%s-%s", r.Prefix, key)
	}

	return key
}

func (r Repository[ID, T]) encode(entity T) ([]byte, error) {
	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(entity); err != nil {
		return nil, fmt.Errorf("redisddd.Repository: failed to encode entity in gob, %w", err)
	}

	return buf.Bytes(), nil
}

func (r Repository[ID, T]) Add(ctx context.Context, entity T) error {
	key := r.makeKey(entity.ID().String())

	data, err := r.encode(entity)
	if err != nil {
		return err
	}

	if err := r.Client.Set(ctx, key, data, 0).Err(); err != nil {
		return fmt.Errorf("redisddd.Repository: failed to set entity in redis, %w", err)
	}

	return nil
}

func (r Repository[ID, T]) decode(data []byte) (T, error) {
	var entity T

	buf := bytes.NewBuffer(data)
	if err := gob.NewDecoder(buf).Decode(&entity); err != nil {
		return entity, fmt.Errorf("redisddd.Repository: failed to decode entity from gob, %w", err)
	}

	return entity, nil
}

func (r Repository[ID, T]) Get(ctx context.Context, id ID) (T, error) {
	var zero T

	key := r.makeKey(id.String())

	data, err := r.Client.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		err = ddd.ErrEntityNotFound
	}
	if err != nil {
		return zero, fmt.Errorf("redisddd.Repository: failed to get entity, %w", err)
	}

	entity, err := r.decode(data)
	if err != nil {
		return zero, err
	}

	return entity, nil
}
