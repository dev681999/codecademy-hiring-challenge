package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"

	"catinator-backend/pkg/config"
	"catinator-backend/pkg/db/ent"

	"entgo.io/ent/dialect/sql/schema"
)

func CreateConnStr(cfg config.DB) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode)
}

func openEntClient(cfg config.DB) (*ent.Client, error) {
	connStr := CreateConnStr(cfg)
	client, err := ent.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func OpenEntClient(ctx context.Context, cfg config.DB) (*ent.Client, error) {
	client, err := openEntClient(cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Debug {
		client = client.Debug()
	}

	if err := client.Schema.Create(ctx, schema.WithAtlas(true)); err != nil {
		return nil, err
	}

	return client, nil
}
