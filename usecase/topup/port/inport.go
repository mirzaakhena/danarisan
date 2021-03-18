package port

import (
	"context"
	"time"
)

// TopupInport ...
type TopupInport interface {
	Execute(ctx context.Context, req TopupRequest) (*TopupResponse, error)
}

// TopupRequest ...
type TopupRequest struct {
	PesertaID      string
	TotalTopup     float64
	TanggalHariIni time.Time `json:"-"`
}

// TopupResponse ...
type TopupResponse struct {
}
