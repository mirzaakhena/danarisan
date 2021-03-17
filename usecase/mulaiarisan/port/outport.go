package port

import (
	"github.com/mirzaakhena/danarisan/domain/repository"
	"github.com/mirzaakhena/danarisan/domain/service"
)

// MulaiArisanOutport ...
type MulaiArisanOutport interface {
	repository.FindOnePesertaRepo
	repository.FindOneArisanRepo
	repository.FindAllSlotRepo
	repository.FindOneUndianRepo
	repository.SaveArisanRepo
	repository.SaveTagihanRepo
	service.TransactionDB
}
