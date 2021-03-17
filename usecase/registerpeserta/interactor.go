package registerpeserta

import (
	"context"
	"github.com/mirzaakhena/danarisan/domain/entity"
	"github.com/mirzaakhena/danarisan/domain/service"
	"github.com/mirzaakhena/danarisan/usecase/registerpeserta/port"
)

//go:generate mockery --dir port/ --name RegisterPesertaOutport -output mocks/

type registerPesertaInteractor struct {
	outport port.RegisterPesertaOutport
}

// NewUsecase ...
func NewUsecase(outputPort port.RegisterPesertaOutport) port.RegisterPesertaInport {
	return &registerPesertaInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *registerPesertaInteractor) Execute(ctx context.Context, req port.RegisterPesertaRequest) (*port.RegisterPesertaResponse, error) {

	res := &port.RegisterPesertaResponse{}

	err := service.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

		pesertaObj, err := entity.NewPeserta(entity.PesertaRequest{
			GenerateID: func() string {
				return req.PesertaID
			},
			Nama: req.PesertaID,
		})
		if err != nil {
			return err
		}

		_, err = r.outport.SavePeserta(ctx, pesertaObj)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
