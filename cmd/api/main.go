package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tx-pgx/cmd/api/handlers/car"
	"github.com/tx-pgx/internal/persistence/config"
	"github.com/tx-pgx/internal/persistence/postgres"
	carRepository "github.com/tx-pgx/internal/persistence/postgres/car"
	carService "github.com/tx-pgx/internal/services/car"
)

func main() {
	var db postgres.Database
	dsn := postgres.BuildDSNPostgres(
		config.Parsed.PostgresHost,
		config.Parsed.PostgresPort,
		config.Parsed.PostgresUser,
		config.Parsed.PostgresPassword,
		config.Parsed.PostgresDB)
	db.ConnectClient(dsn)
	db.RegisterSchemas()
	db.RegisterMigration()

	carRepo := carRepository.Repository{
		Db: db,
	}
	carServ := carService.Service{
		CarRepo: &carRepo,
	}
	carHand := car.Handler{
		CarService: &carServ,
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "pong"}) })
	r.GET("/car/:id", carHand.ReadCar)
	r.POST("/car", carHand.Register)
	r.POST("/car/buy", carHand.BuyCar)

	r.Run()
}
