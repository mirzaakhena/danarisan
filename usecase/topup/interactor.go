package topup

import (
	"context"

	"github.com/mirzaakhena/danarisan/usecase/topup/port"
)

//go:generate mockery --dir port/ --name TopupOutport -output mocks/

type topupInteractor struct {
	outport port.TopupOutport
}

// NewTopupUsecase ...
func NewUsecase(outputPort port.TopupOutport) port.TopupInport {
	return &topupInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *topupInteractor) Execute(ctx context.Context, req port.TopupRequest) (*port.TopupResponse, error) {

	res := &port.TopupResponse{}

	resTopup, err := r.outport.Topup(ctx, port.TopupRequest{})
	if err != nil {
		return nil, err
	}
	_ = resTopup

	return res, nil
}
