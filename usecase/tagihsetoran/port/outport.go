package port

import (
	"context"
	"time"

	"github.com/mirzaakhena/danarisan/domain/entity"

	"github.com/mirzaakhena/danarisan/domain/repository"
	"github.com/mirzaakhena/danarisan/domain/service"
)

// TagihSetoranOutport ...
type TagihSetoranOutport interface {
	repository.FindOneArisanRepo
	repository.FindOneUndianRepo
	repository.FindAllTagihanRepo
	repository.SaveTagihanRepo
	service.TransactionDB
	CreatePayment(ctx context.Context, req CreatePaymentRequest) (*CreatePaymentResponse, error)
	NotifyPeserta(ctx context.Context, req NotifyPesertaRequest) (*NotifyPesertaResponse, error)
}

// CreatePaymentRequest ...
type CreatePaymentRequest struct {
	TagihanID          string
	Nominal            float64
	ArisanID           string
	TanggalKadaluwarsa time.Time
	Tagihan            entity.Tagihan
}

// CreatePaymentResponse ...
type CreatePaymentResponse struct {
	CheckoutURL   string
	AcquirementID string
}

// NotifyPesertaRequest ...
type NotifyPesertaRequest struct {
	PesertaID string
}

// NotifyPesertaResponse ...
type NotifyPesertaResponse struct {
}
