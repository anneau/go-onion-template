package database

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/anneau/go-template/config"
	"github.com/go-sql-driver/mysql"
)

func NewConnection(config *config.DatabaseConfig) (*sql.DB, error) {
	c := mysql.Config{
		User: config.User,
		Passwd: config.Password,
		Net: "tcp",
		Addr: config.Host + ":" + strconv.Itoa(config.Port),
	}

	conn, err := sql.Open("mysql", c.FormatDSN())

	if err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	return conn, nil
}
