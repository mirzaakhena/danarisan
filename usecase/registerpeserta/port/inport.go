package port

import (
	"context"
)

// RegisterPesertaInport ...
type RegisterPesertaInport interface {
	Execute(ctx context.Context, req RegisterPesertaRequest) (*RegisterPesertaResponse, error)
}

// RegisterPesertaRequest ...
type RegisterPesertaRequest struct {
	PesertaID string
}

// RegisterPesertaResponse ...
type RegisterPesertaResponse struct {
}
