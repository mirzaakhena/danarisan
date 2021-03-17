package port

import (
	"github.com/mirzaakhena/danarisan/domain/repository"
	"github.com/mirzaakhena/danarisan/domain/service"
)

// RegisterPesertaOutport ...
type RegisterPesertaOutport interface {
	repository.SavePesertaRepo
	service.TransactionDB
}
