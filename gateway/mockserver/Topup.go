package mockserver

import (
	"context"

	"github.com/mirzaakhena/danarisan/infrastructure/log"
	"github.com/mirzaakhena/danarisan/infrastructure/util"
	"github.com/mirzaakhena/danarisan/usecase/topup/port"
)

type topup struct {
}

// NewTopupGateway ...
func NewTopupGateway() port.TopupOutport {
	return &topup{}
}

// Topup ...
func (r *topup) Topup(ctx context.Context, req port.TopupRequest) (*port.TopupResponse, error) {
	log.InfoRequest(ctx, util.MustJSON(req))
	// _ = req.PesertaID
	// _ = req.TotalTopup
	// _ = req.TanggalHariIni

	var res port.TopupResponse

	log.InfoResponse(ctx, util.MustJSON(res))
	return &res, nil
}
