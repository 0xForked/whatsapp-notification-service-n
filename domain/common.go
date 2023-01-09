package domain

import "context"

type FindWith int

const (
	FindWithID FindWith = iota
	// FindWithRelationID
)

type (
	// IReadOneRepository - ReadOne/Show
	IReadOneRepository[T any] interface {
		Find(ctx context.Context, key FindWith, val any) (data *T, err error)
	}

	// ICreateRepository - Create
	ICreateRepository[T any] interface {
		Create(ctx context.Context, param *T) error
	}

	// IUpdateRepository - Update
	IUpdateRepository[T any] interface {
		Update(ctx context.Context, param *T) error
	}

	// IDeleteRepository - Delete
	IDeleteRepository[T any] interface {
		Delete(ctx context.Context, param *T) error
	}
)
