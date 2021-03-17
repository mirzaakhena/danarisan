package aplikasicontoh

import (
	"context"

	"github.com/mirzaakhena/danarisan/usecase/aplikasicontoh/port"
)

//go:generate mockery --dir port/ --name AplikasiContohOutport -output mocks/

type aplikasiContohInteractor struct {
	outport port.AplikasiContohOutport
}

// NewUsecase ...
func NewUsecase(outputPort port.AplikasiContohOutport) port.AplikasiContohInport {
	return &aplikasiContohInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *aplikasiContohInteractor) Execute(ctx context.Context, req port.AplikasiContohRequest) (*port.AplikasiContohResponse, error) {

	res := &port.AplikasiContohResponse{}

	resAplikasiContoh, err := r.outport.AplikasiContoh(ctx, port.AplikasiContohRequest{})
	if err != nil {
		return nil, err
	}
	_ = resAplikasiContoh

	return res, nil
}
