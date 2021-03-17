package jawabundangan

import (
	"context"
	"github.com/mirzaakhena/danarisan/application/apperror"
	"github.com/mirzaakhena/danarisan/domain/service"
	"github.com/mirzaakhena/danarisan/domain/vo"

	"github.com/mirzaakhena/danarisan/usecase/jawabundangan/port"
)

//go:generate mockery --dir port/ --name JawabUndanganOutport -output mocks/

type jawabUndanganInteractor struct {
	outport port.JawabUndanganOutport
}

// NewJawabUndanganUsecase ...
func NewJawabUndanganUsecase(outputPort port.JawabUndanganOutport) port.JawabUndanganInport {
	return &jawabUndanganInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *jawabUndanganInteractor) Execute(ctx context.Context, req port.JawabUndanganRequest) (*port.JawabUndanganResponse, error) {

	res := &port.JawabUndanganResponse{}

	err := service.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

		pesertaObj, err := r.outport.FindOnePeserta(ctx, vo.PesertaID(req.PesertaID))
		if err != nil {
			return err
		}

		if pesertaObj == nil {
			return apperror.PesertaTidakDitemukan
		}

		if pesertaObj.StateUndangan == vo.TerimaUndanganStateEnum {
			return apperror.PesertaSudahJoinUndangan
		}

		if pesertaObj.StateUndangan == vo.TolakUndanganStateEnum {
			return apperror.PesertaSudahMenolakUndangan
		}

		pesertaObj.StateUndangan, err = vo.NewUndanganState(req.Jawaban)
		if err != nil {
			return err
		}

		_, err = r.outport.SavePeserta(ctx, pesertaObj)
		if err != nil {
			return err
		}

		arisanObj, err := r.outport.FindOneArisan(ctx, pesertaObj.ArisanYgDiikuti)
		if err != nil {
			return err
		}

		if arisanObj == nil {
			return apperror.ArisanTidakDitemukan
		}

		arisanObj.TambahPeserta()

		_, err = r.outport.SaveArisan(ctx, arisanObj)
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
