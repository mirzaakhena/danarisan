package port

import (
	"github.com/mirzaakhena/danarisan/domain/repository"
	"github.com/mirzaakhena/danarisan/domain/service"
)

// BukaAplikasiOutport ...
type BukaAplikasiOutport interface {
	repository.FindOnePesertaRepo
	repository.FindOneArisanRepo
	repository.FindAllPesertaRepo
	repository.FindAllSlotRepo
	repository.FindAllUndianRepo
	repository.FindAllTagihanByArisanIDRepo
	repository.FindAllJurnalRepo
	repository.FindAllSaldoAkunRepo
	service.ReadOnlyDB
}
