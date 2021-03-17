package entity

import (
	"github.com/mirzaakhena/danarisan/domain/service"
	"time"

	"github.com/mirzaakhena/danarisan/domain/vo"
)

type Slot struct {
	BaseModel
	ID            vo.SlotID      `gorm:"primaryKey"` //
	ArisanID      vo.ArisanID    `json:"-"`//
	GroupSlotID   vo.GroupSlotID `json:"-"`//
	PesertaID     vo.PesertaID   //
	TanggalMenang *time.Time     //
}

type SlotRequest struct {
	GenerateID  service.IDGenerator //
	ArisanID    vo.ArisanID         //
	GroupSlotID vo.GroupSlotID      //
	PesertaID   vo.PesertaID        //
}

func NewSlot(req SlotRequest) (obj *Slot, err error) {

	obj = &Slot{}

	slotID, err := vo.NewSlotID(vo.SlotIDRequest{
		GenerateID: req.GenerateID,
	})
	if err != nil {
		return nil, err
	}

	obj.ID = slotID
	obj.ArisanID = req.ArisanID
	obj.GroupSlotID = req.GroupSlotID
	obj.PesertaID = req.PesertaID
	obj.TanggalMenang = nil

	return obj, nil
}

func (r *Slot) Terpilih(tanggal time.Time) error {
	r.TanggalMenang = &tanggal
	return nil
}
