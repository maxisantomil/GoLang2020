package database

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // adding sqlite driver support
	"github.com/maxisantomil/GoLang2020.git/internal/config"
)

// NewDatabase ...
func NewDatabase(conf *config.Config) (*sqlx.DB, error) {
	switch conf.Db.Type {
	case "sqlite3":
		db, err := sqlx.Open(conf.Db.Driver, conf.Db.Conn)
		if err != nil {
			return nil, err
		}
		err = db.Ping()
		if err != nil {
			return nil, err
		}
		return db, nil

	default:
		return nil, errors.New("invalid db type")
	}
}
