package port

import (
	"github.com/mirzaakhena/danarisan/domain/repository"
	"github.com/mirzaakhena/danarisan/domain/service"
)

// UndangPesertaOutport ...
type UndangPesertaOutport interface {
	repository.FindOneArisanRepo
	repository.FindOneArisanByAdminIDRepo
	repository.SavePesertaRepo
	repository.SaveListOfPesertaRepo
	repository.FindPesertaByIDsRepo
	service.TransactionDB
}
