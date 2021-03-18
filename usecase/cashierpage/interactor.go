package cashierpage

import (
	"context"

	"github.com/mirzaakhena/danarisan/usecase/cashierpage/port"
)

//go:generate mockery --dir port/ --name CashierPageOutport -output mocks/

type cashierPageInteractor struct {
	outport port.CashierPageOutport
}

// NewCashierPageUsecase ...
func NewCashierPageUsecase(outputPort port.CashierPageOutport) port.CashierPageInport {
	return &cashierPageInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *cashierPageInteractor) Execute(ctx context.Context, req port.CashierPageRequest) (*port.CashierPageResponse, error) {

    res := &port.CashierPageResponse{}
    
    resCashierPage, err := r.outport.CashierPage(ctx, port.CashierPageRequest {})
    if err != nil {
        return nil, err
    }
    _ = resCashierPage
    
    return res, nil
}
