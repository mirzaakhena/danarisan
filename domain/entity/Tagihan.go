package entity

import (
	"github.com/mirzaakhena/danarisan/domain/vo"
	"time"
)

type Tagihan struct {
	BaseModel
	ID               vo.TagihanID     //
	ArisanID         vo.ArisanID      //
	UndianID         vo.UndianID      //
	PesertaID        vo.PesertaID     //
	AcquirementID    vo.AcquirementID //
	CheckoutURL      string           //
	Nominal          float64          //
	JumlahSlot       int              //
	TanggalPelunasan *time.Time       //
	Status           vo.TagihanStatus //
}

type TagihanRequest struct {
	ArisanID   vo.ArisanID  //
	UndianID   vo.UndianID  //
	PesertaID  vo.PesertaID //
	Nominal    float64      //
	JumlahSlot int          //
}

func NewTagihan(req TagihanRequest) (obj *Tagihan, err error) {

	obj = &Tagihan{}

	obj.ID, err = vo.NewTagihanID(vo.TagihanIDRequest{
		UndianID:  req.UndianID,
		PesertaID: req.PesertaID,
	})
	if err != nil {
		return nil, err
	}

	obj.PesertaID = req.PesertaID
	obj.UndianID = req.UndianID
	obj.PesertaID = req.PesertaID
	obj.Nominal = req.Nominal
	obj.JumlahSlot = req.JumlahSlot
	obj.AcquirementID = ""
	obj.CheckoutURL = ""
	obj.TanggalPelunasan = nil
	obj.Status = vo.BelumDitagihTagihanStatusEnum

	return obj, nil
}

func (r *Tagihan) SimpanPenagihan(acqID string, checkoutURL string) (err error) {
	r.AcquirementID, err = vo.NewAcquirementID(vo.AcquirementIDRequest{AcquirementID: acqID})
	if err != nil {
		return err
	}
	r.CheckoutURL = checkoutURL
	r.Status = vo.MenungguPembayaranTagihanStatusEnum
	return nil
}

func (r *Tagihan) Bayar(tanggal time.Time) {
	r.TanggalPelunasan = &tanggal
	r.Status = vo.LunasTagihanStatusEnum
}

func (r *Tagihan) TidakDiBayar() {
	r.Status = vo.KadaluwarsaTagihanStatusEnum
}
