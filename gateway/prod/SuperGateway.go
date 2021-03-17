package prod

import (
	"context"
	port2 "github.com/mirzaakhena/danarisan/usecase/kocokundian/port"
	port1 "github.com/mirzaakhena/danarisan/usecase/tagihsetoran/port"
	"math/rand"

	"github.com/mirzaakhena/danarisan/gateway/prod/repoimpl"
	"github.com/mirzaakhena/danarisan/gateway/prod/serviceimpl"
	"gorm.io/gorm"
)

type SuperGateway struct {
	*repoimpl.DatabaseImpl
	*repoimpl.TransactionImpl
	*repoimpl.RepositoryImplementation
	*serviceimpl.ServiceImplementation
}



// NewSuperGateway ...
func NewSuperGateway(db *gorm.DB) *SuperGateway {
	return &SuperGateway{
		DatabaseImpl:             repoimpl.SingletonDatabaseImpl(db),
		TransactionImpl:          repoimpl.SingletonTransactionImpl(db),
		RepositoryImplementation: repoimpl.SingletonRepositoryImplementation(db),
		ServiceImplementation:    serviceimpl.SingletonServiceImplementation(),
	}
}

func (r *SuperGateway) CreatePayment(ctx context.Context, req port1.CreatePaymentRequest) (*port1.CreatePaymentResponse, error) {
	// TODO create mock server
	return nil, nil
}

func (r *SuperGateway) NotifyPeserta(ctx context.Context, req port1.NotifyPesertaRequest) (*port1.NotifyPesertaResponse, error) {

	// TODO create mock server

	return nil, nil
}

func (r *SuperGateway) GetRandomNumber(ctx context.Context, req port2.GetRandomNumberRequest) (*port2.GetRandomNumberResponse, error) {

	res := port2.GetRandomNumberResponse{
		RandomNumber: rand.Intn(req.Length),
	}
	return &res, nil
}
func (r *SuperGateway) TopupPeserta(ctx context.Context, req port2.TopupPesertaRequest) (*port2.TopupPesertaResponse, error) {

	// TODO create mock server

	return nil, nil
}
