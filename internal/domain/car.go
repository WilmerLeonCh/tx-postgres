package domain

import (
	"time"
)

type Car struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Version         string `json:"version"`
	ManufactureDate int64  `json:"manufacture_date"`
	Owner           string `json:"owner"`
}

type CreateCarRequest struct {
	Name            string `json:"name"`
	Version         string `json:"version"`
	ManufactureDate int64  `json:"manufacture_date"`
	Owner           string `json:"owner"`
}

type BuyCarRequest struct {
	CarID string `json:"car_id" binding:"required"`
	Buyer string `json:"buyer" binding:"required"`
	Price int    `json:"price" binding:"required"`
}

type CarEvent struct {
	ID     string    `json:"id"`
	CarID  int       `json:"car_id"`
	Seller string    `json:"seller"`
	Buyer  string    `json:"buyer"`
	Price  int       `json:"price"`
	SoldAt time.Time `json:"sold_at"`
}
