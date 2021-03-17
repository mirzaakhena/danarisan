package database

import "context"

// ReadOnlyDB used to get database object from any database implementation.
// For consistency reason both TransactionDB and ReadOnlyDB will seek database object under the context params
type ReadOnlyDB interface {
	GetDatabase(ctx context.Context) (context.Context, error)
}
