package entity

import (
	"github.com/mirzaakhena/danarisan/domain/vo"
)

type Peserta struct {
	ID              vo.PesertaID     //
	Nama            string           //
	Membayar        int              //
	TidakMembayar   int              //
	ArisanYgDiikuti vo.ArisanID      //
	StateUndangan   vo.UndanganState //
	IsAdmin         bool             //
}

type PesertaRequest struct {
	GenerateID func() string //
	Nama       string        //
}

func NewPeserta(req PesertaRequest) (*Peserta, error) {

	pesertaID, err := vo.NewPesertaID(vo.PesertaIDRequest{
		GenerateID: req.GenerateID,
	})
	if err != nil {
		return nil, err
	}

	var obj Peserta
	obj.ID = pesertaID
	obj.Nama = req.Nama
	obj.Membayar = 0
	obj.TidakMembayar = 0
	obj.StateUndangan = vo.NganggurUndanganStateEnum
	obj.IsAdmin = false
	obj.ArisanYgDiikuti = ""

	return &obj, nil
}

func (r *Peserta) JadiAdmin(arisanID vo.ArisanID) {
	r.IsAdmin = true
	r.ArisanYgDiikuti = arisanID
	r.StateUndangan = vo.TerimaUndanganStateEnum
}

func (r *Peserta) MelakukanPembayaran() {
	r.Membayar++
}

func (r *Peserta) TidakMelakukanPembayaran() {
	r.TidakMembayar++
}

func (r *Peserta) DitawarkanIkutArisan(arisanID vo.ArisanID) {
	r.StateUndangan = vo.DitawarkanUndanganStateEnum
	r.ArisanYgDiikuti = arisanID
}
