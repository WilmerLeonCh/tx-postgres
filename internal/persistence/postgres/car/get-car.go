package car

import (
	"context"

	"github.com/tx-pgx/internal/persistence/model"
)

func (r *Repository) GetCar(ctx context.Context, id string) (*model.Car, error) {
	var res *model.Car
	if err := r.model(ctx, &model.Car{}).
		Where("id = ?", id).
		First(&res).
		Error; err != nil {
		return nil, err
	}
	return res, nil
}
