package store

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	config *Config
	db     *pgxpool.Pool
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open(ctx context.Context) error {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", s.config.Username, s.config.Password, s.config.Host, s.config.Port, s.config.Database)
	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return err
	}
	if err = dbpool.Ping(ctx); err != nil {
		return err
	}

	s.db = dbpool
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
