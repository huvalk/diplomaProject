package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

const dsn = `pool_max_conns=30 host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable`

func InitDatabase() (*pgxpool.Pool, error) {
	dsnFmt := fmt.Sprintf(dsn, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	config, err := pgxpool.ParseConfig(dsnFmt)
	if err != nil {
		return nil, err
	}
	return pgxpool.ConnectConfig(context.Background(), config)
}
