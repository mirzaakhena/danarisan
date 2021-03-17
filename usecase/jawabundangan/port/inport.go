package port

import (
	"context"
)

// JawabUndanganInport ...
type JawabUndanganInport interface {
	Execute(ctx context.Context, req JawabUndanganRequest) (*JawabUndanganResponse, error)
}

// JawabUndanganRequest ...
type JawabUndanganRequest struct { 
}

// JawabUndanganResponse ...
type JawabUndanganResponse struct { 
}
