package entity

import (
	"github.com/mirzaakhena/danarisan/application/apperror"
	"github.com/mirzaakhena/danarisan/domain/vo"
	"time"
)

type Arah string

const (
	ArahBertambah = "BERTAMBAH"
	ArahBerkurang = "BERKURANG"
)

type SaldoAkun struct {
	ArisanID  vo.ArisanID  //
	PesertaID vo.PesertaID //
	JurnalID  vo.JurnalID  //
	AkunType  vo.AkunType  //
	Tanggal   time.Time    //
	Sequence  int          //
	Amount    float64      //
	Balance   float64      //
}

type SaldoAkunRequest struct {
	Jurnal              *Jurnal
	AkunType            vo.AkunType //
	Nominal             float64     //
	Arah                Arah        //
	SaldoAkunSebelumnya *SaldoAkun  //
	Sequence            int         //
}

func NewSaldoAkun(req SaldoAkunRequest) (*SaldoAkun, error) {

	var obj SaldoAkun

	obj.ArisanID = req.Jurnal.ArisanID
	obj.PesertaID = req.Jurnal.PesertaID
	obj.JurnalID = req.Jurnal.ID
	obj.Tanggal = req.Jurnal.Tanggal
	obj.AkunType = req.AkunType
	obj.Sequence = req.Sequence

	if req.Nominal > 0 {
		return nil, apperror.NominalHarusLebihBesarDariNol
	}

	if req.Arah == ArahBertambah {
		obj.Amount = req.Nominal

	} else if req.Arah == ArahBerkurang {
		obj.Amount = -req.Nominal

	}

	if req.SaldoAkunSebelumnya != nil {
		obj.Balance = req.SaldoAkunSebelumnya.Balance + obj.Amount
	}

	return &obj, nil
}
