package car

import (
	"context"

	"github.com/tx-pgx/internal/persistence/model"
)

func (r *Repository) GetCarsByVersion(ctx context.Context, version string) ([]model.Car, error) {
	var res []model.Car
	if err := r.model(ctx, &model.Car{}).
		Where("version = ?", version).
		Find(&res).
		Error; err != nil {
		return nil, err
	}
	return res, nil
}
