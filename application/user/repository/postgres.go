package repository

import (
	"context"
	"diplomaProject/application/models"
	"diplomaProject/application/user"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserDatabase struct {
	conn *pgxpool.Pool
}

func NewUserDatabase(db *pgxpool.Pool) user.Repository {
	return &UserDatabase{conn: db}
}

func (ud *UserDatabase) JoinEvent(uid, evtID int) error {
	sql := `insert into event_users values($1,$2)`
	queryResult, err := ud.conn.Exec(context.Background(), sql, evtID, uid)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != 1 {
		return errors.New("already join event")
	}
	return nil
}

func (ud *UserDatabase) GetByID(uid int) (*models.User, error) {
	u := models.User{}
	sql := `select * from users where id = $1`
	queryResult := ud.conn.QueryRow(context.Background(), sql, uid)
	err := queryResult.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, err
}

func (ud *UserDatabase) GetByName(name string) (*models.User, error) {
	u := models.User{}
	sql := `select * from users where name = $1`
	queryResult := ud.conn.QueryRow(context.Background(), sql, name)
	err := queryResult.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, err
}
