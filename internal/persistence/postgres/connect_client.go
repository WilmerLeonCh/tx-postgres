package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Conn *gorm.DB
}

func BuildDSNPostgres(host, port, user, password, dbname string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func (d *Database) ConnectClient(dsn string) {
	db, errOpen := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableAutomaticPing: false,
		Logger:               logger.Default.LogMode(logger.Info),
	})
	if errOpen != nil {
		panic(errOpen)
	}
	d.Conn = db
}
