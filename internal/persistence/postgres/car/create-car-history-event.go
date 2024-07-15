package car

import (
	"context"

	"github.com/tx-pgx/internal/persistence/model"
)

func (r *Repository) CreateCarHistoryEvent(ctx context.Context, event *model.CarEvent) error {
	return r.model(ctx, &model.CarEvent{}).Create(event).Error
}
