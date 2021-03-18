package port

import (
	"context"
)

// BukaAplikasiInport ...
type BukaAplikasiInport interface {
	Execute(ctx context.Context, req BukaAplikasiRequest) (*BukaAplikasiResponse, error)
}

// BukaAplikasiRequest ...
type BukaAplikasiRequest struct {
	PesertaID string
}

// BukaAplikasiResponse ...
type BukaAplikasiResponse struct {
	Data interface{}
	User interface{}
}
