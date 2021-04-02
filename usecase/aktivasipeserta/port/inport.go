package port

import (
	"context"
)

// AktivasiPesertaInport ...
type AktivasiPesertaInport interface {
	Execute(ctx context.Context, req AktivasiPesertaRequest) (*AktivasiPesertaResponse, error)
}

// AktivasiPesertaRequest ...
type AktivasiPesertaRequest struct {
	PesertaID string
}

// AktivasiPesertaResponse ...
type AktivasiPesertaResponse struct {
}
