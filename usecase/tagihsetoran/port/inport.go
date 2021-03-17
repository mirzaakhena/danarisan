package port

import (
	"context"
	"time"
)

// TagihSetoranInport ...
type TagihSetoranInport interface {
	Execute(ctx context.Context, req TagihSetoranRequest) (*TagihSetoranResponse, error)
}

// TagihSetoranRequest ...
type TagihSetoranRequest struct {
	ArisanID string
	HariIni  time.Time
}

// TagihSetoranResponse ...
type TagihSetoranResponse struct {
}
