package entity

import (
	"time"

	"github.com/mirzaakhena/danarisan/domain/vo"
)

type GroupSlot struct {
	BaseModel
	ID            vo.GroupSlotID `gorm:"primaryKey"` //
	ArisanID      vo.ArisanID    `json:"-"`          //
	TanggalMenang *time.Time     //
}

type GroupSlotRequest struct {
	ArisanID vo.ArisanID //
	Index    int         //
}

func NewGroupSlot(req GroupSlotRequest) (obj *GroupSlot, err error) {

	obj = &GroupSlot{}

	obj.ID, err = vo.NewGroupSlotID(vo.GroupSlotIDRequest{
		ArisanID: req.ArisanID,
		Index:    req.Index,
	})
	if err != nil {
		return nil, err
	}

	obj.ArisanID = req.ArisanID
	obj.TanggalMenang = nil

	return obj, nil
}
