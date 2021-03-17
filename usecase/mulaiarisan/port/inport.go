package port

import (
	"context"
)

// MulaiArisanInport ...
type MulaiArisanInport interface {
	Execute(ctx context.Context, req MulaiArisanRequest) (*MulaiArisanResponse, error)
}

// MulaiArisanRequest ...
type MulaiArisanRequest struct {
	ArisanID string
}

// MulaiArisanResponse ...
type MulaiArisanResponse struct {
}
