package vo

type PesertaID string

type PesertaIDRequest struct {
	GenerateID func() string
}

func NewPesertaID(req PesertaIDRequest) (PesertaID, error) {
	obj := PesertaID(req.GenerateID())
	return obj, nil
}
