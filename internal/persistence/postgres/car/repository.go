package car

import (
	"context"

	"gorm.io/gorm"

	"github.com/tx-pgx/internal/persistence/postgres"
)

type Repository struct {
	Db postgres.Database
}

func (r *Repository) repository(ctx context.Context) *postgres.StandardRepository {
	return &postgres.StandardRepository{Ctx: ctx, Db: r.Db}
}

func (r *Repository) model(ctx context.Context, model interface{}) *gorm.DB {
	return r.repository(ctx).Model(model)
}

func (r *Repository) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	return r.repository(ctx).WithinTransaction(tFunc)
}
