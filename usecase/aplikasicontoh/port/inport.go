package port

import (
	"context"
)

// AplikasiContohInport ...
type AplikasiContohInport interface {
	Execute(ctx context.Context, req AplikasiContohRequest) (*AplikasiContohResponse, error)
}

// AplikasiContohRequest ...
type AplikasiContohRequest struct {
}

// AplikasiContohResponse ...
type AplikasiContohResponse struct {
}
