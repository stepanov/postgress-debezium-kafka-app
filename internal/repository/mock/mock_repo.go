// Package mockrepo provides an in-memory UserRepository implementation for tests.
package mockrepo

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/stepanov/postgress-debezium-kafka-app/internal/model"
	"github.com/stepanov/postgress-debezium-kafka-app/internal/repository"
)

var _ repository.UserRepository = (*InMemoryUserRepo)(nil)

// InMemoryUserRepo is a simple in-memory implementation of UserRepository for tests.
type InMemoryUserRepo struct {
	mu sync.RWMutex
	m  map[uuid.UUID]*model.User
}

// New returns a new InMemoryUserRepo.
func New() *InMemoryUserRepo {
	return &InMemoryUserRepo{m: make(map[uuid.UUID]*model.User)}
}

// Create stores a user in the repository.
func (r *InMemoryUserRepo) Create(ctx context.Context, u *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now().UTC()
	}
	r.m[u.ID] = u
	return nil
}

// GetByID returns a user by its ID.
func (r *InMemoryUserRepo) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.m[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return u, nil
}

// Update replaces an existing user.
func (r *InMemoryUserRepo) Update(ctx context.Context, u *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.m[u.ID]; !ok {
		return errors.New("not found")
	}
	r.m[u.ID] = u
	return nil
}

// Delete removes a user by ID.
func (r *InMemoryUserRepo) Delete(ctx context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.m, id)
	return nil
}

// List returns all users.
func (r *InMemoryUserRepo) List(ctx context.Context) ([]*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*model.User, 0, len(r.m))
	for _, u := range r.m {
		out = append(out, u)
	}
	return out, nil
}
