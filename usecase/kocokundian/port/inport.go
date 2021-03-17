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
	UndianID       string
	TanggalHariIni time.Time
}

// KocokUndianResponse ...
type KocokUndianResponse struct {
}
