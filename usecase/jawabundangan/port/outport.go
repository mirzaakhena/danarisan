package port

import (
	"github.com/mirzaakhena/danarisan/domain/repository"
	"github.com/mirzaakhena/danarisan/domain/service"
)

// JawabUndanganOutport ...
type JawabUndanganOutport interface {
	repository.FindOnePesertaRepo
	repository.FindOneArisanRepo
	repository.SavePesertaRepo
	repository.SaveArisanRepo
	service.TransactionDB
}
