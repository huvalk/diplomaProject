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

func (r *NotificationRepository) MarkAsWatched(notificationID int) error {
	sql := `update notification 
			set watched = TRUE
			where id = $1`

	_, err := r.conn.Exec(context.Background(), sql, notificationID)
	return err
}

func (r *NotificationRepository) GetPendingNotification(userID int) (arr []channel.Notification, err error) {
	arr = []channel.Notification{}
	sql := `select id, type, user_id, message, created, watched, status
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

		err = rows.Scan(&n.ID, &n.Type, &n.UserID, &n.Message, &n.Created, &n.Watched, &n.Status)
		if err != nil {
			return nil, err
		}

		arr = append(arr, n)
	}
	return arr, err
}
