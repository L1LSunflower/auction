package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

type SqlConnection struct {
	DB *sql.DB
}

var onceSqlConnectionInstance sync.Once
var sqlConnectionInstance *SqlConnection

func SqlInstance(driver, dbConnectionString string) *SqlConnection {
	if sqlConnectionInstance == nil {
		onceSqlConnectionInstance.Do(func() {
			conn, err := SqlConnect(driver, dbConnectionString)

			if err != nil {
				log.Println("sql connect error_response:", err)
			}

			sqlConnectionInstance = conn
		})
	} else {
		err := sqlConnectionInstance.Ping(driver, dbConnectionString)

		if err != nil {
			log.Println("sql ping error_response:", err)
		}
	}

	return sqlConnectionInstance
}

func SqlConnect(driver, dbConnectionString string) (*SqlConnection, error) {
	db, err := sql.Open(driver, dbConnectionString)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(0)

	return &SqlConnection{
		DB: db,
	}, nil
}

func (c *SqlConnection) Ping(driver, dbConnectionString string) error {
	if err := c.DB.Ping(); err == nil {
		return nil
	}

	sq, err := SqlConnect(driver, dbConnectionString)

	if err != nil {
		return err
	}

	c.DB = sq.DB
	return nil
}
