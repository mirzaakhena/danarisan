package port

import "github.com/mirzaakhena/danarisan/domain/repository"

// AktivasiPesertaOutport ...
type AktivasiPesertaOutport interface {
	repository.FindOnePeserta2Repo
	repository.SavePesertaRepo
}
