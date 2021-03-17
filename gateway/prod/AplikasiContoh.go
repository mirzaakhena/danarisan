package prod

import (
	"context"

	"github.com/mirzaakhena/danarisan/infrastructure/log"
	"github.com/mirzaakhena/danarisan/infrastructure/util"
	"github.com/mirzaakhena/danarisan/usecase/aplikasicontoh/port"
)

type aplikasiContoh struct {
}

// NewAplikasiContohGateway ...
func NewAplikasiContohGateway() port.AplikasiContohOutport {
	return &aplikasiContoh{}
}

// AplikasiContoh ...
func (r *aplikasiContoh) AplikasiContoh(ctx context.Context, req port.AplikasiContohRequest) (*port.AplikasiContohResponse, error) {
	log.InfoRequest(ctx, util.MustJSON(req))

	var res port.AplikasiContohResponse

	log.InfoResponse(ctx, util.MustJSON(res))
	return &res, nil
}

// SimpanData ...
func (r *aplikasiContoh) SimpanData(ctx context.Context, req port.SimpanDataRequest) (*port.SimpanDataResponse, error) {
	log.InfoRequest(ctx, util.MustJSON(req))

	var res port.SimpanDataResponse

	log.InfoResponse(ctx, util.MustJSON(res))
	return &res, nil
}
