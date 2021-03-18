package entity

import (
	"github.com/mirzaakhena/danarisan/domain/vo"
	"time"
)

type Undian struct {
	BaseModel
	ID             vo.UndianID `gorm:"primaryKey"` //
	ArisanID       vo.ArisanID `json:"-"`          //
	PutaranKe      int         //
	TanggalTagihan time.Time   //
	TanggalUndian  time.Time   //
	BiayaAdmin     float64     //
	BiayaArisan    float64     //
}

type UndianRequest struct {
	ArisanID       vo.ArisanID //
	PutaranKe      int         //
	TanggalTagihan string      //
	TanggalUndian  string      //
	BiayaAdmin     float64     //
	BiayaArisan    float64     //
}

func NewUndian(req UndianRequest) (obj *Undian, err error) {

	obj = &Undian{}

	obj.ArisanID = req.ArisanID
	obj.PutaranKe = req.PutaranKe
	obj.BiayaAdmin = req.BiayaAdmin
	obj.BiayaArisan = req.BiayaArisan

	obj.ID, err = vo.NewUndianID(vo.UndianIDRequest{
		ArisanID:  req.ArisanID,
		PutaranKe: req.PutaranKe,
	})
	if err != nil {
		return nil, err
	}

	obj.TanggalTagihan, err = time.Parse("2006-01-02", req.TanggalTagihan)
	if err != nil {
		return nil, err
	}

	obj.TanggalUndian, err = time.Parse("2006-01-02", req.TanggalUndian)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
