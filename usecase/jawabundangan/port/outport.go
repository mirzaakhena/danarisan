package port 

import (
	"context"
) 

// JawabUndanganOutport ...
type JawabUndanganOutport interface { 
	JawabUndangan(ctx context.Context, req JawabUndanganRequest) (*JawabUndanganResponse, error) 
}

