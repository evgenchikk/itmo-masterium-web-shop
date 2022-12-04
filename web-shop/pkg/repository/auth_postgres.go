package repository

import (
	"context"

	"github.com/evgenchikk/itmo-web-shop/models"
	"github.com/jackc/pgx/v4"
)

type AuthPostgres struct {
	db *pgx.Conn
}

func NewAuthPostgres(db *pgx.Conn) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) SignUp(user models.User) (int, error) {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	var id int

	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return -1, err
	}
	defer tx.Rollback(context.Background())

	if err := tx.QueryRow(context.Background(), query, user.Name, user.Email, user.Password).Scan(&id); err != nil {
		return -1, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return -1, err
	}

	return id, nil
}
