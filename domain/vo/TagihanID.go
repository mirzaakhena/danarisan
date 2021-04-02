package vo

import "fmt"

type TagihanID string

type TagihanIDRequest struct {
	UndianID  UndianID  //
	PesertaID PesertaID //
}

func NewTagihanID(req TagihanIDRequest) (TagihanID, error) {
	obj := TagihanID(fmt.Sprintf("%s_%s", req.UndianID, req.PesertaID))
	return obj, nil
}

func (r TagihanID) String() string {
	return string(r)
}
