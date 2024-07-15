package port

import (
	"context"
	"time"

	"github.com/tx-pgx/internal/domain"
	"github.com/tx-pgx/internal/persistence/model"
)

// CarRepository car persistence repository
type CarRepository interface {
	Transactor
	CreateCar(ctx context.Context, car *model.Car) error
	UpdateCar(ctx context.Context, car *model.Car) error
	CreateCarHistoryEvent(ctx context.Context, event *model.CarEvent) error
	UpdateCarHistoryEvent(ctx context.Context, event *model.CarEvent) error
	GetCar(ctx context.Context, id string) (*model.Car, error)
	GetCarsBatch(ctx context.Context, ids []string) ([]model.Car, error)
	GetCarsByTimePeriod(ctx context.Context, from, to time.Time) ([]model.Car, error)
	GetCarsByVersion(ctx context.Context, model string) ([]model.Car, error)
	DeleteCar(ctx context.Context, id string) error
}

//go:generate mockery --case=snake --outpkg=mocksRepository --output=../persistence/postgres/car/mock --name=CarRepository

// CarService car service
type CarService interface {
	Transactor
	GetNewCars(ctx context.Context, from time.Time) ([]domain.Car, error)
	SaveCar(ctx context.Context, car *domain.Car) error
	ReadCar(ctx context.Context, id string) (*domain.Car, error)
	Sell(ctx context.Context, carId, owner, buyer string, price int) error
}

//go:generate mockery --case=snake --outpkg=mocksService --output=../services/car/mock --name=CarService
