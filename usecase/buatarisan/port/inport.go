package port

import (
	"context"
)

// BuatArisanInport ...
type BuatArisanInport interface {
	Execute(ctx context.Context, req BuatArisanRequest) (*BuatArisanResponse, error)
}

// BuatArisanRequest ...
type BuatArisanRequest struct {
	PesertaID       string  //
	NamaAdmin       string  //
	NamaArisan      string  //
	SetoranTiapSlot float64 //
	TanggalTagihan  string  //
	TanggalUndian   string  //
	BiayaArisan     float64 //
}

// BuatArisanResponse ...
type BuatArisanResponse struct {
	ArisanID string
}
