package starter

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func NewDatabase(ctx context.Context, apiServer *ApiServer) error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		apiServer.Сonfig.Database.Username,
		apiServer.Сonfig.Database.Password,
		apiServer.Сonfig.Database.Host,
		apiServer.Сonfig.Database.Port,
		apiServer.Сonfig.Database.DatabaseName,
	)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		fmt.Errorf("unable to parse database config: %w", err)
	}

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		fmt.Errorf("unable to connect to database: %w", err)
	}

	logrus.Info("Connected to PostgreSQL")
	apiServer.Db = pool
	return nil
}
