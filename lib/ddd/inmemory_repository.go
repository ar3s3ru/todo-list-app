package ddd

import (
	"context"
	"sync"
)

type InmemoryEntityRepository[ID comparable, T Entity[ID]] struct {
	mx       *sync.RWMutex
	entities map[ID]T
}

func NewInmemoryEntityRepository[ID comparable, T Entity[ID]]() InmemoryEntityRepository[ID, T] {
	return InmemoryEntityRepository[ID, T]{
		mx:       new(sync.RWMutex),
		entities: make(map[ID]T),
	}
}

func (r InmemoryEntityRepository[ID, T]) Add(_ context.Context, entity T) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.entities[entity.ID()] = entity

	return nil
}

func (r InmemoryEntityRepository[ID, T]) Get(_ context.Context, id ID) (T, error) {
	var zero T

	r.mx.RLock()
	defer r.mx.RUnlock()

	if entity, ok := r.entities[id]; ok {
		return entity, nil
	}

	return zero, ErrEntityNotFound
}
