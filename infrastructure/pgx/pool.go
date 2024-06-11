package pgx

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func SetupPool(connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	// Set maximum number of connections in the pool
	config.MaxConns = 60
	// Set maximum number of idle connections in the pool
	config.MaxConnIdleTime = time.Minute

	// Additional configurations can be set here
	config.HealthCheckPeriod = time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
