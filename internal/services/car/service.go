package car

import (
	"context"
	"time"

	"github.com/tx-pgx/internal/domain"
	"github.com/tx-pgx/internal/persistence/model"
	"github.com/tx-pgx/internal/port"
)

type Service struct {
	CarRepo port.CarRepository
}

func (s *Service) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	return s.CarRepo.WithinTransaction(ctx, tFunc)
}

func (s *Service) ToDomainModels(cars []model.Car) []domain.Car {
	var res []domain.Car
	for _, car := range cars {
		var manufactureDate int64
		if car.ManufactureDate != nil {
			manufactureDate = car.ManufactureDate.UnixMilli()
		}
		res = append(res, domain.Car{
			ID:              car.ID,
			Name:            car.Name,
			Version:         car.Version,
			ManufactureDate: manufactureDate,
			Owner:           car.Owner,
		})
	}
	return res
}

func (s *Service) ToDomainModel(car *model.Car) *domain.Car {
	var manufactureDate int64
	if car.ManufactureDate != nil {
		manufactureDate = car.ManufactureDate.UnixMilli()
	}
	var res = &domain.Car{
		ID:              car.ID,
		Name:            car.Name,
		Version:         car.Version,
		ManufactureDate: manufactureDate,
		Owner:           car.Owner,
	}
	return res
}

func (s *Service) ToPersistenceModel(car *domain.Car) *model.Car {
	var manufactureDate *time.Time
	if car.ManufactureDate != 0 {
		date := time.UnixMilli(car.ManufactureDate)
		manufactureDate = &date
	}
	var res = &model.Car{
		ID:              car.ID,
		Name:            car.Name,
		Version:         car.Version,
		ManufactureDate: manufactureDate,
		Owner:           car.Owner,
	}
	return res
}
