package db

import (
	"database/sql"
	"fmt"
	"github.com/Kisesaiyajin/testApi/config"
	_ "github.com/lib/pq"
)

func NewPostgreSQLDB(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password =%s sslmode =%s", cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err

	}
	return db, nil

}
