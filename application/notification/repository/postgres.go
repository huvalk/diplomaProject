package repository

import (
	"context"
	"diplomaProject/application/notification"
	"diplomaProject/pkg/channel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type NotificationRepository struct {
	conn *pgxpool.Pool
}

func NewNotificationRepository(db *pgxpool.Pool) notification.Repository {
	return &NotificationRepository{conn: db}
}

func (r *NotificationRepository) SaveNotification(n *channel.Notification) error {
	sql := `insert into notification 
			(type, user_id, message, created, watched, status) 
			values ($1, $2, $3, $4, $5, $6)`

	_, err := r.conn.Exec(context.Background(), sql, n.Type, n.UserID, n.Message, n.Created, n.Watched, n.Status)
	return err
}

func (r *NotificationRepository) GetEventName(eventID int) (name string, err error) {
	sql := `select distinct(name) from event where id=$1`

	err = r.conn.QueryRow(context.Background(), sql, eventID).Scan(&name)
	return name, err
}

func (r *NotificationRepository) MarkAsWatched(notificationID int) error {
	sql := `update notification 
			set watched = TRUE
			where id = $1`

	_, err := r.conn.Exec(context.Background(), sql, notificationID)
	return err
}

// TODO убрать
func (r *NotificationRepository) GetMoreLastNotification(userID int) (arr []channel.Notification, err error) {
	arr = []channel.Notification{}
	sql := `select type, status, message, created
			from notification 
			where user_id = $1
			and id in (
				select max(id)
				from notification
				where user_id=$1
				and watched = FALSE
				group by type, user_id, status
				)
			and watched = FALSE
			order by created desc
			limit 3`

	rows, err := r.conn.Query(context.Background(), sql, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var n channel.Notification

		err = rows.Scan(&n.Type, &n.Status, &n.Message, &n.Created)
		if err != nil {
			return nil, err
		}

		arr = append(arr, n)
	}
	return arr, err
}

func (r *NotificationRepository) GetLastNotification(userID int) (arr []channel.Notification, err error) {
	arr = []channel.Notification{}
	sqlRow := `select type, status, message, created
			from notification 
			where id in (
				select max(id)
				from notification
				where user_id=$1
				group by type, user_id, status
				)
			and user_id=$1
			order by created desc
			limit 10`


	rows, err := r.conn.Query(context.Background(), sqlRow, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var n channel.Notification

		err = rows.Scan(&n.Type, &n.Status, &n.Message, &n.Created)
		if err != nil {
			return nil, err
		}

		arr = append(arr, n)
	}
	return arr, err
}
