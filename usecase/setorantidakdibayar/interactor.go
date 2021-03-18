package setorantidakdibayar

import (
	"context"
	"github.com/mirzaakhena/danarisan/application/apperror"
	"github.com/mirzaakhena/danarisan/domain/service"
	"github.com/mirzaakhena/danarisan/usecase/setorantidakdibayar/port"
)

//go:generate mockery --dir port/ --name SetoranTidakDibayarOutport -output mocks/

type setoranTidakDibayarInteractor struct {
	outport port.SetoranTidakDibayarOutport
}

// NewUsecase ...
func NewUsecase(outputPort port.SetoranTidakDibayarOutport) port.SetoranTidakDibayarInport {
	return &setoranTidakDibayarInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *setoranTidakDibayarInteractor) Execute(ctx context.Context, req port.SetoranTidakDibayarRequest) (*port.SetoranTidakDibayarResponse, error) {

	res := &port.SetoranTidakDibayarResponse{}

	err := service.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

		tagihanObj, err := r.outport.FindOneTagihan(ctx, req.TagihanID)
		if err != nil {
			return err
		}

		if tagihanObj == nil {
			return apperror.TagihanTidakDitemukan
		}

		tagihanObj.TidakDiBayar()

		err = r.outport.SaveTagihan(ctx, tagihanObj)
		if err != nil {
			return err
		}

		pesertaObj, err := r.outport.FindOnePeserta(ctx, tagihanObj.PesertaID.String())
		if err != nil {
			return err
		}

		if pesertaObj == nil {
			return apperror.PesertaTidakDitemukan
		}

		pesertaObj.TidakMelakukanPembayaran()

		err = r.outport.SavePeserta(ctx, pesertaObj)
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
