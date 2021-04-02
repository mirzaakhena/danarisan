package vo

import "fmt"

type UndianID string

type UndianIDRequest struct {
	ArisanID  ArisanID
	PutaranKe int
}

func NewUndianID(req UndianIDRequest) (UndianID, error) {
	obj := UndianID(fmt.Sprintf("%s_%d", req.ArisanID, req.PutaranKe))
	return obj, nil
}

func (r UndianID) String() string {
	return string(r)
}
