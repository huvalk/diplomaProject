package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

const dsn = `pool_max_conns=30 host=localhost port=8081 user=%s password=%s dbname=%s sslmode=disable`

func InitDatabase() (*pgxpool.Pool, error) {
	dsnFmt := fmt.Sprintf(dsn, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	config, err := pgxpool.ParseConfig(dsnFmt)
	if err != nil {
		return nil, err
	}
	return pgxpool.ConnectConfig(context.Background(), config)
}
