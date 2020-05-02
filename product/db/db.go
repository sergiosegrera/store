package db

import (
	"github.com/go-pg/pg/v9"
)

func NewConnection(options *pg.Options) (*pg.DB, error) {
	db := pg.Connect(options)

	_, err := db.Exec("SELECT 1")
	if err != nil {
		return db, err
	}

	return db, err
}
