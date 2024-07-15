package car

import (
	"context"

	"github.com/tx-pgx/internal/persistence/model"
)

func (r *Repository) UpdateCar(ctx context.Context, car *model.Car) error {
	return r.model(ctx, &model.Car{}).
		Where("id = ?", car.ID).
		Updates(car).Error
}
