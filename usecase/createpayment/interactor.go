package createpayment

import (
	"context"
	"time"

	"github.com/mirzaakhena/danarisan/usecase/createpayment/port"
)

//go:generate mockery --dir port/ --name CreatePaymentOutport -output mocks/

type createPaymentInteractor struct {
	outport port.CreatePaymentOutport
}

// NewCreateOrderUsecase ...
func NewUsecase(outputPort port.CreatePaymentOutport) port.CreatePaymentInport {
	return &createPaymentInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *createPaymentInteractor) Execute(ctx context.Context, req port.CreatePaymentRequest) (*port.CreatePaymentResponse, error) {

	res := &port.CreatePaymentResponse{}

	paymentFinishNotifyResponse, err := r.outport.PaymentFinishNotify(ctx, port.PaymentFinishNotifyRequest{
		TagihanID: req.TagihanID,
		Delay:     2 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	_ = paymentFinishNotifyResponse

	return res, nil
}
