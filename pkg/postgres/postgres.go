package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Config struct {
	Host     string `yaml:"POSTGRES_HOST" env:"POSTGRES_HOST" env-default:"localhost"`
	Port     string `yaml:"POSTGRES_PORT" env:"POSTGRES_PORT" env-default:"5432"`
	Username string `yaml:"POSTGRES_USERNAME" env:"POSTGRES_USERNAME" env-default:"root"`
	Password string `yaml:"POSTGRES_PASS" env:"POSTGRES_PASS" env-default:"1234"`
	Database string `yaml:"POSTGRES_DB" env:"POSTGRES_DB " env-default:"postgres"`
}

func New(cfg Config) (*pgx.Conn, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}
	return conn, nil
}
