package entity

import (
	"github.com/mirzaakhena/danarisan/domain/service"
	"github.com/mirzaakhena/danarisan/domain/vo"
	"time"
)

type Jurnal struct {
	BaseModel
	ID         vo.JurnalID   `gorm:"primaryKey"` //
	ArisanID   vo.ArisanID   `json:"-"`          //
	PesertaID  vo.PesertaID  //
	Tanggal    time.Time     //
	JurnalType vo.JurnalType //
}

type JurnalRequest struct {
	GenerateID service.IDGenerator //
	ArisanID   vo.ArisanID         //
	PesertaID  vo.PesertaID        //
	Tanggal    time.Time           //
	JurnalType vo.JurnalType       //
}

func NewJurnal(req JurnalRequest) (obj *Jurnal, err error) {

	obj = &Jurnal{}

	obj.ID, err = vo.NewJurnalID(vo.JurnalIDRequest{GenerateID: req.GenerateID})
	if err != nil {
		return nil, err
	}

	obj.ArisanID = req.ArisanID
	obj.PesertaID = req.PesertaID
	obj.JurnalType = req.JurnalType
	obj.Tanggal = req.Tanggal

	return obj, nil
}
