package car

import (
	"context"
	"fmt"

	"github.com/tx-pgx/internal/domain"
)

func (s *Service) ReadCar(ctx context.Context, id string) (*domain.Car, error) {
	car, err := s.CarRepo.GetCar(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("read car by id: %v", err)
	}
	return s.ToDomainModel(car), nil
}
