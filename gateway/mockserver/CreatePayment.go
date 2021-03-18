package mockserver

import (
	"context"

	"github.com/mirzaakhena/danarisan/infrastructure/log"
	"github.com/mirzaakhena/danarisan/infrastructure/util"
	"github.com/mirzaakhena/danarisan/usecase/createpayment/port"
)

type createPayment struct {
}

// NewCreatePaymentGateway ...
func NewCreatePaymentGateway() port.CreatePaymentOutport {
	return &createPayment{}
}

// CreatePayment ...
func (r *createPayment) CreatePayment(ctx context.Context, req port.CreatePaymentRequest) (*port.CreatePaymentResponse, error) {
	log.InfoRequest(ctx, util.MustJSON(req))
	// _ = req.TagihanID
	// _ = req.Nominal
	// _ = req.PesertaID
	// _ = req.TanggalKadaluwarsa
	// _ = req.TanggalHariIni

	var res port.CreatePaymentResponse

	log.InfoResponse(ctx, util.MustJSON(res))
	return &res, nil
}

// PaymentFinishNotify ...
func (r *createPayment) PaymentFinishNotify(ctx context.Context, req port.PaymentFinishNotifyRequest) (*port.PaymentFinishNotifyResponse, error) {
	log.InfoRequest(ctx, util.MustJSON(req))
	// _ = req.Delay
	// _ = req.TagihanID

	var res port.PaymentFinishNotifyResponse

	log.InfoResponse(ctx, util.MustJSON(res))
	return &res, nil
}
