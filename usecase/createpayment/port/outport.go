package port

import (
	"context"
	"time"
)

// CreatePaymentOutport ...
type CreatePaymentOutport interface {
	CreatePayment(ctx context.Context, req CreatePaymentRequest) (*CreatePaymentResponse, error)
	PaymentFinishNotify(ctx context.Context, req PaymentFinishNotifyRequest) (*PaymentFinishNotifyResponse, error)
}

// PaymentFinishNotifyRequest ...
type PaymentFinishNotifyRequest struct {
	Delay     time.Duration
	TagihanID string
}

// PaymentFinishNotifyResponse ...
type PaymentFinishNotifyResponse struct {
}
