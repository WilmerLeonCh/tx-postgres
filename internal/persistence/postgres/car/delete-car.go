package car

import (
	"context"

	"github.com/tx-pgx/internal/persistence/model"
)

func (r *Repository) DeleteCar(ctx context.Context, id string) error {
	return r.model(ctx, &model.Car{}).Where("id = ?", id).Delete(&model.Car{}).Error
}
