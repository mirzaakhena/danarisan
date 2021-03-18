package port 

import (
	"context"
) 

// CashierPageOutport ...
type CashierPageOutport interface { 
	CashierPage(ctx context.Context, req CashierPageRequest) (*CashierPageResponse, error) 
}

