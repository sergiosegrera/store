package db

import (
	"github.com/sergiosegrera/store/product-manager/models"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

func NewConnection(options *pg.Options) (*pg.DB, error) {
	db := pg.Connect(options)

	err := createSchema(db)
	if err != nil {
		return db, err
	}

	return db, err
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*models.Product)(nil), (*models.Option)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
