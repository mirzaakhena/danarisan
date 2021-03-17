package repoimpl

import (
	"context"
	"gorm.io/gorm"
	"sync"

	"github.com/mirzaakhena/danarisan/infrastructure/log"
)

type contextDBType string

var ContextDBValue contextDBType = "DB"

var onceTransactionImpl sync.Once

var transactionImplObj TransactionImpl

func SingletonTransactionImpl(db *gorm.DB) *TransactionImpl {
	onceTransactionImpl.Do(func() {
		transactionImplObj = TransactionImpl{db: db}
	})
	return &transactionImplObj
}

type TransactionImpl struct {
	db *gorm.DB
}

func (r *TransactionImpl) BeginTransaction(ctx context.Context) (context.Context, error) {
	log.InfoRequest(ctx, "Begin")

	dbTrx := r.db.Begin()

	trxCtx := context.WithValue(ctx, ContextDBValue, dbTrx)

	log.InfoResponse(ctx, "Begin")
	return trxCtx, nil
}

func (r *TransactionImpl) CommitTransaction(ctx context.Context) error {
	log.InfoRequest(ctx, "Commit")

	db, err := extractDB(ctx)
	if err != nil {
		return err
	}

	db.Commit()

	log.InfoResponse(ctx, "Commit")
	return nil
}

func (r *TransactionImpl) RollbackTransaction(ctx context.Context) error {
	log.InfoRequest(ctx, "Rollback")

	db, err := extractDB(ctx)
	if err != nil {
		return err
	}

	db.Rollback()

	log.InfoResponse(ctx, "Rollback")
	return err
}
