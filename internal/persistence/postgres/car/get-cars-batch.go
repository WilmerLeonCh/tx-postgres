package car

import (
	"context"

	"github.com/tx-pgx/internal/persistence/model"
)

func (r *Repository) GetCarsBatch(ctx context.Context, ids []string) ([]model.Car, error) {
	//TODO implement me
	panic("implement me")
}
