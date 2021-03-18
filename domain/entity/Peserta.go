package entity

import (
  "github.com/mirzaakhena/danarisan/domain/vo"
)

type Peserta struct {
  BaseModel
  ID            vo.PesertaID     `gorm:"primaryKey"` //
  Nama          string           ``                  //
  Membayar      int              `json:"-"`          //
  TidakMembayar int              `json:"-"`          //
  ArisanID      vo.ArisanID      ``                  //
  StateUndangan vo.UndanganState ``                  //
  IsAdmin       bool             ``                  //
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
  obj.ArisanID = ""

  return &obj, nil
}

func (r *Peserta) JadiAdmin(arisanID vo.ArisanID) {
  r.IsAdmin = true
  r.ArisanID = arisanID
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
  r.ArisanID = arisanID
}

func (r *Peserta) ResetPeserta() {
  r.IsAdmin = false
  r.StateUndangan = vo.NganggurUndanganStateEnum
  r.ArisanID = ""
}
