package repository

import (
	"context"
	"diplomaProject/application/feed"
	"diplomaProject/application/models"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type FeedDatabase struct {
	conn *pgxpool.Pool
}

func NewFeedDatabase(db *pgxpool.Pool) feed.Repository {
	return &FeedDatabase{conn: db}
}

func (f FeedDatabase) Get(feedID int) (*models.Feed, error) {
	fd := models.Feed{}
	sql := `select * from feed where id = $1`
	queryResult := f.conn.QueryRow(context.Background(), sql, feedID)
	err := queryResult.Scan(&fd.Id, &fd.Event)
	if err != nil {
		return nil, err
	}
	return &fd, err
}

func (f FeedDatabase) GetByEvent(eventID int) (*models.Feed, error) {
	fd := models.Feed{}
	sql := `select * from feed where event = $1`
	queryResult := f.conn.QueryRow(context.Background(), sql, eventID)
	err := queryResult.Scan(&fd.Id, &fd.Event)
	if err != nil {
		return nil, err
	}
	return &fd, err
}

func (f FeedDatabase) Create(eventID int) (*models.Feed, error) {
	sql := `insert into feed values(default,$1) returning id`
	id := 0
	err := f.conn.QueryRow(context.Background(), sql, eventID).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &models.Feed{Id: id, Users: nil, Event: eventID}, nil
}

func (f FeedDatabase) AddUser(uid, eventID int) error {
	fd, err := f.GetByEvent(eventID)
	if err != nil {
		return err
	}
	sql := `insert into feed_users values($1,$2)`
	queryResult, err := f.conn.Exec(context.Background(), sql, fd.Id, uid)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != 1 {
		return errors.New("already in feed")
	}
	return nil
}

func (f FeedDatabase) RemoveUser(uid, eventID int) error {
	fd, err := f.GetByEvent(eventID)
	if err != nil {
		return err
	}
	sql := `delete from feed_users where feed_id=$1 AND user_id=$2`
	queryResult, err := f.conn.Exec(context.Background(), sql, fd.Id, uid)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != 1 {
		return errors.New("not in feed")
	}
	return nil
}

func (f FeedDatabase) GetFeedUsers(feedID int) ([]models.User, error) {
	var us []models.User
	u := models.User{}
	sql := `select u1.id,u1.firstname,u1.lastname,u1.email from feed_users f1
join users u1  on f1.user_id=u1.id where f1.feed_id=$1`
	queryResult, err := f.conn.Query(context.Background(), sql, feedID)
	if err != nil {
		return nil, err
	}
	for queryResult.Next() {
		err = queryResult.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	queryResult.Close()

	return us, err
}
