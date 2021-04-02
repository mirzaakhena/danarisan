package undangpeserta

import (
	"context"
	"github.com/mirzaakhena/danarisan/application/apperror"
	"github.com/mirzaakhena/danarisan/domain/service"
	"github.com/mirzaakhena/danarisan/domain/vo"

	"github.com/mirzaakhena/danarisan/usecase/undangpeserta/port"
)

//go:generate mockery --dir port/ --name UndangPesertaOutport -output mocks/

type undangPesertaInteractor struct {
	outport port.UndangPesertaOutport
}

// NewUsecase ...
func NewUsecase(outputPort port.UndangPesertaOutport) port.UndangPesertaInport {
	return &undangPesertaInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *undangPesertaInteractor) Execute(ctx context.Context, req port.UndangPesertaRequest) (*port.UndangPesertaResponse, error) {

	res := &port.UndangPesertaResponse{}

	err := service.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

		arisanObj, err := r.outport.FindOneArisanByAdminID(ctx, req.AdminID)
		if err != nil {
			return err
		}

		if arisanObj == nil {
			return apperror.ArisanTidakDitemukan
		}

		pesertaObjs, err := r.outport.FindPesertaByIDs(ctx, req.PesertaIDs)
		if err != nil {
			return err
		}

		if pesertaObjs == nil || len(pesertaObjs) == 0 {
			return apperror.PesertaTidakDitemukan
		}

		for _, pesertaObj := range pesertaObjs {

			if pesertaObj.StateUndangan == vo.DitawarkanUndanganStateEnum {
				return apperror.PesertaSudahDiundang
			}

			if pesertaObj.StateUndangan == vo.TerimaUndanganStateEnum {
				return apperror.PesertaSudahJoinUndangan
			}

			pesertaObj.DitawarkanIkutArisan(arisanObj.ID)

		}

		err = r.outport.SaveListOfPeserta(ctx, pesertaObjs)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
