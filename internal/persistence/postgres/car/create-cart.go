package car

import (
	"context"

	"github.com/tx-pgx/internal/persistence/model"
)

func (r *Repository) CreateCar(ctx context.Context, car *model.Car) error {
	return r.model(ctx, &model.Car{}).Create(car).Error
}
