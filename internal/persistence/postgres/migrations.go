package postgres

import (
	"log"

	"github.com/tx-pgx/internal/persistence/model"
)

func (d *Database) RegisterMigration() {
	d.migrate(&model.Car{})
	d.migrate(&model.CarEvent{})
}
func (d *Database) RegisterSchemas() {
	throw(d.Conn.Exec("CREATE SCHEMA IF NOT EXISTS car").Error)
}
func (d *Database) migrate(model interface{}) {
	throw(d.Conn.AutoMigrate(model))
}

func throw(err error) {
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
}
