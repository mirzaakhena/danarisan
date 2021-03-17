package bukaaplikasi

import (
	"context"
	"github.com/mirzaakhena/danarisan/domain/service"

	"github.com/mirzaakhena/danarisan/usecase/bukaaplikasi/port"
)

//go:generate mockery --dir port/ --name BukaAplikasiOutport -output mocks/

type bukaAplikasiInteractor struct {
	outport port.BukaAplikasiOutport
}

// NewBukaAplikasiUsecase ...
func NewBukaAplikasiUsecase(outputPort port.BukaAplikasiOutport) port.BukaAplikasiInport {
	return &bukaAplikasiInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *bukaAplikasiInteractor) Execute(ctx context.Context, req port.BukaAplikasiRequest) (*port.BukaAplikasiResponse, error) {

    res := &port.BukaAplikasiResponse{}

	err := service.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

		_, err := r.outport.FindOnePeserta(ctx, )
		if err != nil {
			return nil, err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
