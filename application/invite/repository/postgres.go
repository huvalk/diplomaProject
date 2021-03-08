package repository

import (
	"diplomaProject/application/invite"
	"github.com/jackc/pgx/v4/pgxpool"
)

type InviteRepository struct {
	conn *pgxpool.Pool
}

func NewInviteRepository(db *pgxpool.Pool) invite.Repository {
	return &InviteRepository{conn: db}
}