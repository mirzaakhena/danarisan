package vo

import "github.com/mirzaakhena/danarisan/domain/service"

type ArisanID string

type ArisanIDRequest struct {
	GenerateID service.IDGenerator
}

func NewArisanID(req ArisanIDRequest) (ArisanID, error) {
	obj := ArisanID(req.GenerateID.Generate())
	return obj, nil
}

func (r ArisanID) String() string {
	return string(r)
}