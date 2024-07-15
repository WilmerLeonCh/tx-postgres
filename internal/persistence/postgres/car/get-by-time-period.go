package car

import (
	"context"
	"time"

	"github.com/tx-pgx/internal/persistence/model"
)

func (r *Repository) GetCarsByTimePeriod(ctx context.Context, from, to time.Time) ([]model.Car, error) {
	var res []model.Car
	if err := r.model(ctx, &model.Car{}).
		Where("manufacture_date BETWEEN ? AND ?", from, to).
		Order("version").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
