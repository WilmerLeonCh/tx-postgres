package car

import (
	"context"
	"fmt"
	"time"

	"github.com/tx-pgx/internal/domain"
)

func (s *Service) GetNewCars(ctx context.Context, from time.Time) ([]domain.Car, error) {
	if err := s.someValidation(); err != nil {
		return nil, fmt.Errorf("invalid request: %v", err)
	}

	// read the latest cars
	cars, err := s.CarRepo.GetCarsByTimePeriod(ctx, from, time.Now())
	if err != nil {
		return nil, fmt.Errorf("newest cars read: %v", err)
	}

	return s.ToDomainModels(cars), nil
}

func (s *Service) someValidation() error {
	return nil
}
