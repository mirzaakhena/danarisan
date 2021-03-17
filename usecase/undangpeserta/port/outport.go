package port

import (
	"github.com/mirzaakhena/danarisan/domain/repository"
	"github.com/mirzaakhena/danarisan/domain/service"
)

// UndangPesertaOutport ...
type UndangPesertaOutport interface {
	repository.FindOneArisanRepo
	repository.FindOnePesertaRepo
	repository.FindOneArisanByAdminIDRepo
	repository.SavePesertaRepo
	service.TransactionDB
}
