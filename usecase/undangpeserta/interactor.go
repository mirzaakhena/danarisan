package undangpeserta

import (
	"context"

	"github.com/mirzaakhena/danarisan/usecase/undangpeserta/port"
)

//go:generate mockery --dir port/ --name UndangPesertaOutport -output mocks/

type undangPesertaInteractor struct {
	outport port.UndangPesertaOutport
}

// NewUndangPesertaUsecase ...
func NewUndangPesertaUsecase(outputPort port.UndangPesertaOutport) port.UndangPesertaInport {
	return &undangPesertaInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *undangPesertaInteractor) Execute(ctx context.Context, req port.UndangPesertaRequest) (*port.UndangPesertaResponse, error) {

	res := &port.UndangPesertaResponse{}

	resUndangPeserta, err := r.outport.UndangPeserta(ctx, port.UndangPesertaRequest{})
	if err != nil {
		return nil, err
	}
	_ = resUndangPeserta

	return res, nil
}
