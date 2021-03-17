package port

import (
	"context"
	"time"
)

// KocokUndianInport ...
type KocokUndianInport interface {
	Execute(ctx context.Context, req KocokUndianRequest) (*KocokUndianResponse, error)
}

// KocokUndianRequest ...
type KocokUndianRequest struct {
	PesertaID      string
	TanggalHariIni time.Time `json:"-"`
}

// KocokUndianResponse ...
type KocokUndianResponse struct {
}
