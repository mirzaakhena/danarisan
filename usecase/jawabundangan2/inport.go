package jawabundangan2

import (
	"context"
)

// Inport ...
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest ...
type InportRequest struct {
	PesertaID string //
	Jawaban   string //
}

// InportResponse ...
type InportResponse struct {
}
