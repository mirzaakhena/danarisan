package port

import (
	"context"
	"time"
)

// BayarSetoranInport ...
type BayarSetoranInport interface {
	Execute(ctx context.Context, req BayarSetoranRequest) (*BayarSetoranResponse, error)
}

// BayarSetoranRequest ...
type BayarSetoranRequest struct {
	TagihanID      string
	TanggalHariIni time.Time `json:"-"`
}

// BayarSetoranResponse ...
type BayarSetoranResponse struct {
}
