package repository

import (
	"diplomaProject/application/notification"
	"github.com/jackc/pgx/v4/pgxpool"
)

type NotificationDatabase struct {
	conn *pgxpool.Pool
}

func NewNotificationDatabase(db *pgxpool.Pool) notification.Repository {
	return &NotificationDatabase{conn: db}
}


//func (e NotificationDatabase) Get(id int) (*models.EventDB, error) {
//	evt := models.EventDB{}
//	sql := `select * from event where id = $1`
//
//	queryResult := e.conn.QueryRow(context.Background(), sql, id)
//	err := queryResult.Scan(&evt.Id, &evt.Name, &evt.Description, &evt.Founder,
//		&evt.DateStart, &evt.DateEnd, &evt.Place)
//	if err != nil {
//		return nil, err
//	}
//	return &evt, err
//}
//
//func (e NotificationDatabase) Create(newEvent *models.Event) (*models.EventDB, error) {
//	sql := `INSERT INTO event VALUES(default,$1,$2,$3,$4,$5,$6)  RETURNING id`
//	id := 0
//	err := e.conn.QueryRow(context.Background(), sql, newEvent.Name, newEvent.Description,
//		newEvent.Founder, newEvent.DateStart, newEvent.DateEnd, newEvent.Place).Scan(&id)
//	if err != nil {
//		return nil, err
//	}
//	return e.Get(id)
//}
//
//func (e NotificationDatabase) CheckUser(evtID, uid int) bool {
//	sql := `select * from event_users where event_id = $1 AND user_id = $2`
//
//	queryResult := e.conn.QueryRow(context.Background(), sql, evtID, uid)
//	err := queryResult.Scan(&evtID, &uid)
//	if err != nil {
//		return false
//	}
//	return true
//}
