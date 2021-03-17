package port

import (
	"context"
)

// SetoranTidakDibayarInport ...
type SetoranTidakDibayarInport interface {
	Execute(ctx context.Context, req SetoranTidakDibayarRequest) (*SetoranTidakDibayarResponse, error)
}

// SetoranTidakDibayarRequest ...
type SetoranTidakDibayarRequest struct {
	TagihanID string
}

// SetoranTidakDibayarResponse ...
type SetoranTidakDibayarResponse struct {
}
