package port

import (
	"context"
)

// CashierPageInport ...
type CashierPageInport interface {
	Execute(ctx context.Context, req CashierPageRequest) (*CashierPageResponse, error)
}

// CashierPageRequest ...
type CashierPageRequest struct { 
}

// CashierPageResponse ...
type CashierPageResponse struct { 
}
