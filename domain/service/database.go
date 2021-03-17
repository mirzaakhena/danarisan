package service

import "context"

// ReadOnlyDB used to get database object from any database implementation.
// For consistency reason both TransactionDB and ReadOnlyDB will seek database object under the context params
type ReadOnlyDB interface {
	GetDatabase(ctx context.Context) (context.Context, error)
}

// ReadOnly is helper function that simplify the readonly db
func ReadOnly(ctx context.Context, trx ReadOnlyDB, trxFunc func(dbCtx context.Context) (interface{}, error)) (interface{}, error) {
	dbCtx, err := trx.GetDatabase(ctx)
	if err != nil {
		return nil, err
	}
	return trxFunc(dbCtx)
}
