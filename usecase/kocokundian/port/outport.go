package port

import (
	"context"

	"github.com/mirzaakhena/danarisan/domain/repository"
	"github.com/mirzaakhena/danarisan/domain/service"
)

// KocokUndianOutport ...
type KocokUndianOutport interface {
	repository.FindOneArisanRepo
	repository.FindOneUndianByIDRepo
	repository.FindAllSlotNotWinYetRepo
	repository.FindAllSlotRepo
	repository.FindOnePesertaRepo
	repository.FindLastSaldoAkunRepo
	repository.SaveSlotRepo
	repository.SaveArisanRepo
	repository.SaveUndianRepo
	repository.SaveTagihanRepo
	repository.SaveJurnalRepo
	repository.SaveSaldoAkunRepo
	repository.SavePesertaRepo
	service.TransactionDB
	service.IDGenerator
	GetRandomNumber(ctx context.Context, req GetRandomNumberRequest) (*GetRandomNumberResponse, error)
	TopupPeserta(ctx context.Context, req TopupPesertaRequest) (*TopupPesertaResponse, error)
}

// GetRandomNumberRequest ...
type GetRandomNumberRequest struct {
	Length int
}

// GetRandomNumberResponse ...
type GetRandomNumberResponse struct {
	RandomNumber int
}

// TopupPesertaRequest ...
type TopupPesertaRequest struct {
	PesertaID  string
	TotalTopup float64
}

// TopupPesertaResponse ...
type TopupPesertaResponse struct {
}
