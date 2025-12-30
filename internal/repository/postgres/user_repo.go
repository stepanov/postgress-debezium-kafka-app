package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stepanov/postgress-debezium-kafka-app/internal/model"
	"github.com/stepanov/postgress-debezium-kafka-app/internal/repository"
)

var _ repository.UserRepository = (*UserRepo)(nil)

// UserRepo implements UserRepository using Postgres.
type UserRepo struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{pool: pool}
}

func (r *UserRepo) Create(ctx context.Context, u *model.User) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now().UTC()
	}
	_, err := r.pool.Exec(ctx, `INSERT INTO users (id, name, email, created_at) VALUES ($1,$2,$3,$4)`, u.ID, u.Name, u.Email, u.CreatedAt)
	return err
}

func (r *UserRepo) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	u := &model.User{}
	row := r.pool.QueryRow(ctx, `SELECT id, name, email, created_at FROM users WHERE id=$1`, id)
	var idVal uuid.UUID
	if err := row.Scan(&idVal, &u.Name, &u.Email, &u.CreatedAt); err != nil {
		return nil, err
	}
	u.ID = idVal
	return u, nil
}

func (r *UserRepo) Update(ctx context.Context, u *model.User) error {
	_, err := r.pool.Exec(ctx, `UPDATE users SET name=$1, email=$2 WHERE id=$3`, u.Name, u.Email, u.ID)
	return err
}

func (r *UserRepo) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM users WHERE id=$1`, id)
	return err
}

func (r *UserRepo) List(ctx context.Context) ([]*model.User, error) {
	rows, err := r.pool.Query(ctx, `SELECT id, name, email, created_at FROM users ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*model.User
	for rows.Next() {
		u := &model.User{}
		var idVal uuid.UUID
		if err := rows.Scan(&idVal, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		u.ID = idVal
		out = append(out, u)
	}
	return out, nil
}
