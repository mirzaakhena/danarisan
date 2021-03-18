package port

import (
	"context"
	"time"
)

// CreatePaymentInport ...
type CreatePaymentInport interface {
	Execute(ctx context.Context, req CreatePaymentRequest) (*CreatePaymentResponse, error)
}

// CreatePaymentRequest ...
type CreatePaymentRequest struct {
	TagihanID          string
	Nominal            float64
	PesertaID          string
	TanggalKadaluwarsa time.Time
	TanggalHariIni     time.Time `json:"-"`
}

// CreatePaymentResponse ...
type CreatePaymentResponse struct {
}
