package car

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/tx-pgx/internal/persistence/model"
)

func (s *Service) Sell(ctx context.Context, carId, owner, buyer string, price int) error {
	if err := s.validateBuyer(ctx, owner, price); err != nil {
		return err
	}
	if err := s.validatePurchase(ctx, owner); err != nil {
		return err
	}
	event := &model.CarEvent{
		ID:     uuid.NewString(),
		CarID:  carId,
		Seller: owner,
		Buyer:  buyer,
		Price:  price,
		SoldAt: time.Now(),
	}

	if err := s.CarRepo.CreateCarHistoryEvent(ctx, event); err != nil {
		return err
	}

	carUpdates := &model.Car{
		ID:    carId,
		Owner: buyer,
	}
	if err := s.CarRepo.UpdateCar(ctx, carUpdates); err != nil {
		return err
	}

	log.Printf("car %s sold to %s for %d", event.CarID, event.Buyer, price)
	return errors.New("nuevo error jaja")
	return nil
}

func (s *Service) validateBuyer(ctx context.Context, id string, price int) error {
	return nil
}

func (s *Service) validatePurchase(ctx context.Context, id string) error {
	return nil
}
