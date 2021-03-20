package repository

import (
	"context"
	"diplomaProject/application/event"
	"diplomaProject/application/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type EventDatabase struct {
	conn *pgxpool.Pool
}

func NewEventDatabase(db *pgxpool.Pool) event.Repository {
	return &EventDatabase{conn: db}
}

func (e EventDatabase) Get(id int) (*models.EventDB, error) {
	evt := models.EventDB{}
	sql := `select id, name, description, founder, date_start, date_end, state, place
			from event where id = $1`

	queryResult := e.conn.QueryRow(context.Background(), sql, id)
	err := queryResult.Scan(&evt.Id, &evt.Name, &evt.Description, &evt.Founder,
		&evt.DateStart, &evt.DateEnd, &evt.State, &evt.Place)
	if err != nil {
		return nil, err
	}
	return &evt, err
}

func (e EventDatabase) Create(newEvent *models.Event) (*models.EventDB, error) {
	sql := `INSERT INTO event VALUES(default,$1,$2,$3,$4,$5,$6)  RETURNING id`
	id := 0
	err := e.conn.QueryRow(context.Background(), sql, newEvent.Name, newEvent.Description,
		newEvent.Founder, newEvent.DateStart, newEvent.DateEnd, newEvent.Place).Scan(&id)
	if err != nil {
		return nil, err
	}
	return e.Get(id)
}

func (e EventDatabase) CheckUser(evtID, uid int) bool {
	sql := `select * from event_users where event_id = $1 AND user_id = $2`

	queryResult := e.conn.QueryRow(context.Background(), sql, evtID, uid)
	err := queryResult.Scan(&evtID, &uid)

	return err != nil
}
