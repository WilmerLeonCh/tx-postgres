package car

import (
	"context"

	"github.com/google/uuid"

	"github.com/tx-pgx/internal/domain"
)

func (s *Service) SaveCar(ctx context.Context, registerCarReq *domain.Car) error {
	id := registerCarReq.ID
	if id == "" {
		id = uuid.NewString()
	}
	var car = domain.Car{
		ID:              id,
		Name:            registerCarReq.Name,
		Version:         registerCarReq.Version,
		ManufactureDate: registerCarReq.ManufactureDate,
		Owner:           registerCarReq.Owner,
	}
	mCar := s.ToPersistenceModel(&car)
	if err := s.CarRepo.CreateCar(ctx, mCar); err != nil {
		return err
	}
	return nil
}
