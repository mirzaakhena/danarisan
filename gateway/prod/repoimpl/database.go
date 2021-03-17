package repoimpl

import (
	"context"
	"github.com/mirzaakhena/danarisan/application/apperror"
	"gorm.io/gorm"
	"sync"

	"github.com/mirzaakhena/danarisan/infrastructure/log"
)

var onceDatabaseImpl sync.Once

var databaseImplObj DatabaseImpl

func SingletonDatabaseImpl(db *gorm.DB) *DatabaseImpl {
	onceDatabaseImpl.Do(func() {
		databaseImplObj = DatabaseImpl{db: db}
	})
	return &databaseImplObj
}

type DatabaseImpl struct {
	db *gorm.DB
}

// GetDatabase put the database object into context.
// From client view, it will return the context to user
func (r *DatabaseImpl) GetDatabase(ctx context.Context) context.Context {
	log.InfoRequest(ctx, "GetDB")

	trxCtx := context.WithValue(ctx, ContextDBValue, r.db)

	log.InfoResponse(ctx, "GetDB")
	return trxCtx
}

// extractDB is used by other repo to extract the database from context
func extractDB(ctx context.Context) (*gorm.DB, error) {

	db, ok := ctx.Value(ContextDBValue).(*gorm.DB)
	if !ok {
		return nil, apperror.DatabaseNotFoundInContextError
	}

	return db, nil
}
