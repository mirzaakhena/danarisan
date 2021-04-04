package bayarsetoran2

import (
  "github.com/mirzaakhena/danarisan/domain/repository"
	"github.com/mirzaakhena/danarisan/domain/service"
)

// Outport ...
type Outport interface {
	repository.FindOneTagihanRepo
	repository.FindOnePesertaRepo
	repository.FindLastSaldoAkunRepo
	repository.SaveTagihanRepo
	repository.SavePesertaRepo
	repository.SaveJurnalRepo
	repository.SaveSaldoAkunRepo
	service.IDGenerator
	service.TransactionDB
}
