package repository

import (
	"context"
	"diplomaProject/application/notification"
	"diplomaProject/pkg/channel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type NotificationDatabase struct {
	conn *pgxpool.Pool
}

func NewNotificationDatabase(db *pgxpool.Pool) notification.Repository {
	return &NotificationDatabase{conn: db}
}

func (r *NotificationDatabase) SaveNotification(n *channel.Notification) error {
	sql := `insert into notification 
			(type, user_id, message, created, watched) 
			values ($1, $2, $3, $4, $5)`

	_, err := r.conn.Exec(context.Background(), sql, n.Type, n.UserID, n.Message, n.Created, n.Watched)
	return err
}

func (r *NotificationDatabase) MarkAsWatched(notificationID int) error {
	sql := `update notification 
			set watched = TRUE
			where id = $1`

	_, err := r.conn.Exec(context.Background(), sql, notificationID)
	return err
}

func (r *NotificationDatabase) GetPendingNotification(userID int) (arr []channel.Notification, err error) {
	arr = []channel.Notification{}
	sql := `select (id, type, user_id, message, created, watched)
			from notification 
			where user_id = $1
			and watched = FALSE`

	rows, err := r.conn.Query(context.Background(), sql, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var n channel.Notification

		err = rows.Scan(&n.ID, &n.Type, &n.UserID, &n.Message, &n.Created, &n.Watched)
		if err != nil {
			return nil , err
		}

		arr = append(arr, n)
	}
	return arr, err
}
