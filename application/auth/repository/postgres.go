package repository

import (
	"diplomaProject/application/auth"
	"diplomaProject/application/invite"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	conn *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) auth.Repository {
	return &Repository{conn: db}
}
