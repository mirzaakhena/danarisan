package registerpeserta

import (
	"context"
	"github.com/mirzaakhena/danarisan/domain/entity"

	"github.com/mirzaakhena/danarisan/usecase/registerpeserta/port"
)

//go:generate mockery --dir port/ --name RegisterPesertaOutport -output mocks/

type registerPesertaInteractor struct {
	outport port.RegisterPesertaOutport
}

// NewRegisterPesertaUsecase ...
func NewRegisterPesertaUsecase(outputPort port.RegisterPesertaOutport) port.RegisterPesertaInport {
	return &registerPesertaInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *registerPesertaInteractor) Execute(ctx context.Context, req port.RegisterPesertaRequest) (*port.RegisterPesertaResponse, error) {

	res := &port.RegisterPesertaResponse{}



	pesertaObj, err := entity.NewPeserta(entity.PesertaRequest{
		GenerateID: func() string {
			return req.PesertaID
		},
		Nama: req.PesertaID,
	})
	if err != nil {
		return nil, err
	}

	_, err = r.outport.SavePeserta(ctx, pesertaObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
