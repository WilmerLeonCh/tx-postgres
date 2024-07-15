package model

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	ID              string `gorm:"type:uuid;primaryKey"`
	Name            string
	Version         string
	ManufactureDate *time.Time
	Owner           string
}

// TableName overrides the table name used by Car to `car.cars`
func (Car) TableName() string {
	return "car.cars"
}

type CarEvent struct {
	gorm.Model
	ID     string `gorm:"type:uuid;primaryKey"`
	CarID  string `gorm:"type:uuid"`
	Seller string
	Buyer  string
	Price  int
	SoldAt time.Time
}

// TableName overrides the table name used by CarEvent to `car.car_events`
func (CarEvent) TableName() string {
	return "car.car_events"
}
