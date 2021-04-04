package jawabundangan2

import (
  "github.com/mirzaakhena/danarisan/domain/repository"
	"github.com/mirzaakhena/danarisan/domain/service"
)

// Outport ...
type Outport interface {
	repository.FindOnePesertaRepo
	repository.FindOneArisanRepo
	repository.SavePesertaRepo
	repository.SaveArisanRepo
	repository.SaveSlotRepo
	service.TransactionDB
	service.IDGenerator
}
