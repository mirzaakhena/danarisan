package port

import (
	"context"
)

// AplikasiContohOutport ...
type AplikasiContohOutport interface {
	AplikasiContoh(ctx context.Context, req AplikasiContohRequest) (*AplikasiContohResponse, error)
	SimpanData(ctx context.Context, req SimpanDataRequest) (*SimpanDataResponse, error)
}

// SimpanDataRequest ...
type SimpanDataRequest struct {
}

// SimpanDataResponse ...
type SimpanDataResponse struct {
}
