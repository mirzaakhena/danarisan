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
}

type PesertaRequest struct {
	GenerateID   func() string //
	Nama         string        //
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

	return &obj, nil
}

func (r *Peserta) MelakukanPembayaran() {
	r.Membayar++
}

func (r *Peserta) TidakMelakukanPembayaran() {
	r.TidakMembayar++
}
