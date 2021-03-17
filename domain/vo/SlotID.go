package vo

import "github.com/mirzaakhena/danarisan/domain/service"

type SlotID string

type SlotIDRequest struct {
	GenerateID service.IDGenerator
}

func NewSlotID(req SlotIDRequest) (SlotID, error) {
	obj := SlotID(req.GenerateID.Generate())
	return obj, nil
}
