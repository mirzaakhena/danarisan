package vo

import "github.com/mirzaakhena/danarisan/domain/service"

type JurnalID string

type JurnalIDRequest struct {
	GenerateID service.IDGenerator
}

func NewJurnalID(req JurnalIDRequest) (JurnalID, error) {
	obj := JurnalID(req.GenerateID.Generate())
	return obj, nil
}

func (r JurnalID) String() string {
	return string(r)
}
