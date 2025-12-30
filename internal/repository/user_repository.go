// Package repository defines repository interfaces used by the application.
package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/stepanov/postgress-debezium-kafka-app/internal/model"
)

// UserRepository defines CRUD operations for users.
type UserRepository interface {
	Create(ctx context.Context, u *model.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	Update(ctx context.Context, u *model.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context) ([]*model.User, error)
}
