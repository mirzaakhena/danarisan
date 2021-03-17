package port

import (
	"github.com/mirzaakhena/danarisan/domain/repository"
	"github.com/mirzaakhena/danarisan/domain/service"
)

// SetoranTidakDibayarOutport ...
type SetoranTidakDibayarOutport interface {
	repository.FindOneTagihanRepo
	repository.FindOnePesertaRepo
	repository.SaveTagihanRepo
	repository.SavePesertaRepo
	service.TransactionDB
}
