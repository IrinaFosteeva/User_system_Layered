package repository

import (
	"context"
	"errors"
	"time"

	"github.com/IrinaFosteeva/User_system_Layered/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("not found")

type UserRepository interface {
	GetAll(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id int) (models.User, error)
	Create(ctx context.Context, u models.User) (models.User, error)
	Update(ctx context.Context, u models.User) (models.User, error)
	Delete(ctx context.Context, id int) error
}

type PostgresUserRepo struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepo(db *pgxpool.Pool) *PostgresUserRepo {
	return &PostgresUserRepo{db: db}
}

func (r *PostgresUserRepo) GetAll(ctx context.Context) ([]models.User, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, email, created_at, updated_at FROM users ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, u)
	}
	return res, nil
}

func (r *PostgresUserRepo) GetByID(ctx context.Context, id int) (models.User, error) {
	var u models.User
	row := r.db.QueryRow(ctx, `SELECT id, name, email, created_at, updated_at FROM users WHERE id=$1`, id)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return models.User{}, ErrNotFound
	}
	return u, nil
}

func (r *PostgresUserRepo) Create(ctx context.Context, u models.User) (models.User, error) {
	now := time.Now()
	var id int
	err := r.db.QueryRow(ctx, `INSERT INTO users (name, email, created_at, updated_at) VALUES ($1,$2,$3,$4) RETURNING id`, u.Name, u.Email, now, now).Scan(&id)
	if err != nil {
		return models.User{}, err
	}
	u.ID = id
	u.CreatedAt = now
	u.UpdatedAt = now
	return u, nil
}

func (r *PostgresUserRepo) Update(ctx context.Context, u models.User) (models.User, error) {
	now := time.Now()
	cmdTag, err := r.db.Exec(ctx, `UPDATE users SET name=$1, email=$2, updated_at=$3 WHERE id=$4`, u.Name, u.Email, now, u.ID)
	if err != nil {
		return models.User{}, err
	}
	if cmdTag.RowsAffected() == 0 {
		return models.User{}, ErrNotFound
	}
	u.UpdatedAt = now
	return u, nil
}

func (r *PostgresUserRepo) Delete(ctx context.Context, id int) error {
	cmdTag, err := r.db.Exec(ctx, `DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
