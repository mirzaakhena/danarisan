package port

import (
	"context"
)

// JawabUndanganInport ...
type JawabUndanganInport interface {
	Execute(ctx context.Context, req JawabUndanganRequest) (*JawabUndanganResponse, error)
}

// JawabUndanganRequest ...
type JawabUndanganRequest struct {
	PesertaID string //
	Jawaban   string //
}

// JawabUndanganResponse ...
type JawabUndanganResponse struct {
}
