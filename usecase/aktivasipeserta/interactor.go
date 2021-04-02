package aktivasipeserta

import (
  "context"
  "github.com/mirzaakhena/danarisan/application/apperror"

  "github.com/mirzaakhena/danarisan/usecase/aktivasipeserta/port"
)

//go:generate mockery --dir port/ --name AktivasiPesertaOutport -output mocks/

type aktivasiPesertaInteractor struct {
  outport port.AktivasiPesertaOutport
}

// NewUsecase ...
func NewUsecase(outputPort port.AktivasiPesertaOutport) port.AktivasiPesertaInport {
  return &aktivasiPesertaInteractor{
    outport: outputPort,
  }
}

// Execute ...
func (r *aktivasiPesertaInteractor) Execute(ctx context.Context, req port.AktivasiPesertaRequest) (*port.AktivasiPesertaResponse, error) {

  res := &port.AktivasiPesertaResponse{}

  pesertaObj, err := r.outport.FindOnePeserta2(ctx, req.PesertaID)
  if err != nil {
    return nil, err
  }
  if pesertaObj == nil {
    return nil, apperror.Peserta2TidkDitemukan
  }

  pesertaObj.Aktivasi()

  err = r.outport.SavePeserta(ctx, pesertaObj)
  if err != nil {
    return nil, err
  }

  return res, nil
}
