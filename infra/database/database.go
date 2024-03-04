package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/anneau/go-template/config"
	errctrl "github.com/anneau/go-template/errorctrl"
	_ "github.com/lib/pq"
)

func NewConnection(config *config.DatabaseConfig) (*sql.DB, error) {
	dns := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
	)

	db, err := sql.Open("postgres", dns)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func NewTestConnection() *sql.DB {
	c := config.DatabaseConfig{
		Host:     "db-test",
		User:     "postgres",
		Port:     5432,
		Password: "postgres",
		Database: "postgres",
	}

	return errctrl.Must(NewConnection(&c))
}
