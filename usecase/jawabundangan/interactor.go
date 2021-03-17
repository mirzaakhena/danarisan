package jawabundangan

import (
	"context"

	"github.com/mirzaakhena/danarisan/usecase/jawabundangan/port"
)

//go:generate mockery --dir port/ --name JawabUndanganOutport -output mocks/

type jawabUndanganInteractor struct {
	outport port.JawabUndanganOutport
}

// NewJawabUndanganUsecase ...
func NewJawabUndanganUsecase(outputPort port.JawabUndanganOutport) port.JawabUndanganInport {
	return &jawabUndanganInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *jawabUndanganInteractor) Execute(ctx context.Context, req port.JawabUndanganRequest) (*port.JawabUndanganResponse, error) {

    res := &port.JawabUndanganResponse{}
    
    resJawabUndangan, err := r.outport.JawabUndangan(ctx, port.JawabUndanganRequest {})
    if err != nil {
        return nil, err
    }
    _ = resJawabUndangan
    
    return res, nil
}
