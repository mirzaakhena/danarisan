package port

import (
	"context"
)

// UndangPesertaOutport ...
type UndangPesertaOutport interface {
	UndangPeserta(ctx context.Context, req UndangPesertaRequest) (*UndangPesertaResponse, error)
}
