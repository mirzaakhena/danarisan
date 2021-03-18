package entity

import (
	"github.com/mirzaakhena/danarisan/application/apperror"
	"github.com/mirzaakhena/danarisan/domain/service"
	"github.com/mirzaakhena/danarisan/domain/vo"
)

const (
	MinimalPesertaArisan = 4
)

type Arisan struct {
	BaseModel
	ID                       vo.ArisanID     `gorm:"primaryKey"` //
	Nama                     string          //
	Status                   vo.ArisanStatus //
	ArisanType               vo.ArisanType   `json:"-"` //
	SetoranTiapSlot          float64         //
	JumlahGroup              int             `json:"-"` //
	JumlahSlotTiapGroup      int             `json:"-"` //
	JumlahMaxSlotTiapPeserta int             `json:"-"` //
	TotalPutaran             int             //
	PutaranKe                int             //
	JumlahPeserta            int             //
	AdminID                  vo.PesertaID    //

	ListPeserta   []*Peserta   `gorm:"-"` //
	ListSlot      []*Slot      `gorm:"-"` //
	ListUndian    []*Undian    `gorm:"-"` //
	ListTagihan   []*Tagihan   `gorm:"-"` //
	ListJurnal    []*Jurnal    `gorm:"-"` //
	ListSaldoAkun []*SaldoAkun `gorm:"-"` //
}

type ArisanRequest struct {
	GenerateID      service.IDGenerator //
	Nama            string              //
	SetoranTiapSlot float64             //
	AdminID         vo.PesertaID        //
}

func NewArisan(req ArisanRequest) (obj *Arisan, err error) {

	obj = &Arisan{}

	arisanID, err := vo.NewArisanID(vo.ArisanIDRequest{
		GenerateID: req.GenerateID,
	})
	if err != nil {
		return nil, err
	}

	obj.ID = arisanID
	obj.Nama = req.Nama
	obj.Status = vo.TerimaPesertaArisanStatusEnum
	obj.SetoranTiapSlot = req.SetoranTiapSlot
	obj.JumlahGroup = 0
	obj.JumlahSlotTiapGroup = 0
	obj.JumlahMaxSlotTiapPeserta = 1
	obj.TotalPutaran = 0
	obj.PutaranKe = 0
	obj.ArisanType = vo.SingleSlotArisanTypeEnum
	obj.JumlahPeserta = 0
	obj.AdminID = req.AdminID

	return obj, nil
}

func (r *Arisan) Mulai() error {

	if r.Status == vo.MulaiArisanStatusEnum {
		return apperror.ArisanSudahDimulai
	}

	if r.Status == vo.SelesaiArisanStatusEnum {
		return apperror.ArisanSudahSelesai
	}

	if r.JumlahPeserta < MinimalPesertaArisan {
		return apperror.PesertaArisanMasihKurang
	}

	r.Status = vo.MulaiArisanStatusEnum

	r.PutaranKe = 1

	return nil
}

func (r *Arisan) GetTotalNilaiUndian() float64 {
	return float64(r.JumlahPeserta) * r.SetoranTiapSlot
}

func (r *Arisan) SiapkanArisanBerikutnya() {
	nextPutaran := r.PutaranKe + 1
	if nextPutaran > r.TotalPutaran {
		r.Status = vo.SelesaiArisanStatusEnum
		return
	}
	r.PutaranKe = nextPutaran
}

func (r *Arisan) SudahSelesai() bool {
	return r.Status == vo.SelesaiArisanStatusEnum
}

func (r *Arisan) MasihTerimaPeserta() bool {
	return r.Status == vo.TerimaPesertaArisanStatusEnum
}

func (r *Arisan) TambahPeserta() {
	r.JumlahPeserta++
}
