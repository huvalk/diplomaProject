package infrastructure

import (
	"context"
	"diplomaProject/pkg/globalVars"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

const dsn = `pool_max_conns=30 host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`

func InitDatabase() (*pgxpool.Pool, error) {
	dsnFmt := fmt.Sprintf(dsn, globalVars.POSTGRES_HOST, globalVars.POSTGRES_PORT,
		globalVars.POSTGRES_USER, globalVars.POSTGRES_PASSWORD, globalVars.POSTGRES_DB)

	config, err := pgxpool.ParseConfig(dsnFmt)
	if err != nil {
		return nil, err
	}
	return pgxpool.ConnectConfig(context.Background(), config)
}
