package vo

import "fmt"

type GroupSlotID string

type GroupSlotIDRequest struct {
	ArisanID ArisanID //
	Index    int      //
}

func NewGroupSlotID(req GroupSlotIDRequest) (GroupSlotID, error) {
	obj := GroupSlotID(fmt.Sprintf("%s_%d", req.ArisanID, req.Index))
	return obj, nil
}

func (r GroupSlotID) String() string {
	return string(r)
}
