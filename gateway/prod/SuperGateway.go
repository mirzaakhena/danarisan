package prod

import (
	"github.com/mirzaakhena/danarisan/gateway/prod/repoimpl"
	"github.com/mirzaakhena/danarisan/gateway/prod/serviceimpl"
	"gorm.io/gorm"
)

type superGateway struct {
	*repoimpl.DatabaseImpl
	*repoimpl.TransactionImpl
	*repoimpl.RepositoryImplementation
	*serviceimpl.ServiceImplementation
}

// NewSuperGateway ...
func NewSuperGateway(db *gorm.DB) *superGateway {
	return &superGateway{
		DatabaseImpl:             repoimpl.SingletonDatabaseImpl(db),
		TransactionImpl:          repoimpl.SingletonTransactionImpl(db),
		RepositoryImplementation: repoimpl.SingletonRepositoryImplementation(db),
		ServiceImplementation:    serviceimpl.SingletonServiceImplementation(),
	}
}
