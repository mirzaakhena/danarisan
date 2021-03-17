package port

import (
	"context"
)

// UndangPesertaInport ...
type UndangPesertaInport interface {
	Execute(ctx context.Context, req UndangPesertaRequest) (*UndangPesertaResponse, error)
}

// UndangPesertaRequest ...
type UndangPesertaRequest struct {
	AdminID               string
	ArisanID              string
	PesertaYangDiundangID string
}

// UndangPesertaResponse ...
type UndangPesertaResponse struct {
}
