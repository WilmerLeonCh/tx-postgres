package postgres

import (
	"context"
	"log"

	"gorm.io/gorm"
)

type StandardRepository struct {
	Ctx context.Context
	Db  Database
}

// txKey is a key for transaction in context
type txKey struct{}

// injectTx injects transaction to context
func injectTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// extractTx extracts transaction from context
func extractTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return nil
}

// WithinTransaction runs function within transaction
// The transaction commits when function were finished without error
func (r StandardRepository) WithinTransaction(tFunc func(ctx context.Context) error) error {
	tx := r.Db.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			if errRollback := tx.Rollback().Error; errRollback != nil {
				log.Printf("rollback transaction: %v", errRollback)
			}
			panic(r)
		}
	}()
	err := tFunc(injectTx(r.Ctx, tx))
	if err != nil {
		if errRollback := tx.Rollback().Error; errRollback != nil {
			log.Printf("rollback transaction: %v", errRollback)
		}
		return err
	}
	if errCommit := tx.Commit().Error; errCommit != nil {
		log.Printf("commit transaction: %v", errCommit)
	}
	return nil
}

// Model extract transaction from context if exist or return default connection
func (r StandardRepository) Model(model interface{}) *gorm.DB {
	tx := extractTx(r.Ctx)
	if tx != nil {
		return tx.Model(model)
	}
	return r.Db.Conn.Model(model)
}
