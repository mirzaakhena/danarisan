package port

import (
	"context"
)

// TopupOutport ...
type TopupOutport interface {
	Topup(ctx context.Context, req TopupRequest) (*TopupResponse, error)
}
