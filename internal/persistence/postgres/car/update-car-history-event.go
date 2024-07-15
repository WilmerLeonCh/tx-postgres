package car

import (
	"context"

	"github.com/tx-pgx/internal/persistence/model"
)

func (r *Repository) UpdateCarHistoryEvent(ctx context.Context, event *model.CarEvent) error {
	return r.model(ctx, &model.CarEvent{}).
		Where("id = ?", event.ID).
		Updates(event).Error
}
