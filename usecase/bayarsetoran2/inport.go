package bayarsetoran2

import (
	"context"
	"time"
)

// Inport ...
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest ...
type InportRequest struct {
	TagihanID      string
	TanggalHariIni time.Time `json:"-"`
}

// InportResponse ...
type InportResponse struct {
}
