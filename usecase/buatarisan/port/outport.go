package port

import (
	"github.com/mirzaakhena/danarisan/domain/repository"
	"github.com/mirzaakhena/danarisan/domain/service"
)

// BuatArisanOutport ...
type BuatArisanOutport interface {
	service.TransactionDB
	service.IDGenerator
	repository.SaveArisanRepo
	repository.SavePesertaRepo
	repository.SaveSlotRepo
	repository.SaveUndianRepo
}
