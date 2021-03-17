package bukaaplikasi

import (
	"context"
	"github.com/mirzaakhena/danarisan/application/apperror"
	"github.com/mirzaakhena/danarisan/domain/service"
	"github.com/mirzaakhena/danarisan/domain/vo"

	"github.com/mirzaakhena/danarisan/usecase/bukaaplikasi/port"
)

//go:generate mockery --dir port/ --name BukaAplikasiOutport -output mocks/

type bukaAplikasiInteractor struct {
	outport port.BukaAplikasiOutport
}

// NewUsecase ...
func NewUsecase(outputPort port.BukaAplikasiOutport) port.BukaAplikasiInport {
	return &bukaAplikasiInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *bukaAplikasiInteractor) Execute(ctx context.Context, req port.BukaAplikasiRequest) (*port.BukaAplikasiResponse, error) {

	res := &port.BukaAplikasiResponse{}

	data, err := service.ReadOnly(ctx, r.outport, func(ctx context.Context) (interface{}, error) {

		pesertaObj, err := r.outport.FindOnePeserta(ctx, vo.PesertaID(req.PesertaID))
		if err != nil {
			return nil, err
		}

		if pesertaObj == nil {
			return nil, apperror.PesertaTidakDitemukan
		}

		if pesertaObj.StateUndangan == vo.TolakUndanganStateEnum {

			// peserta tidak join arisan dan bisa bikin arisan sendiri
			return nil, nil
		}

		if pesertaObj.StateUndangan == vo.DitawarkanUndanganStateEnum {
			arisanObj, err := r.outport.FindOneArisan(ctx, pesertaObj.ArisanYgDiikuti)
			if err != nil {
				return nil, err
			}

			if arisanObj == nil {
				return nil, apperror.ArisanTidakDitemukan
			}

			// peserta belum join arisan dan arisan belum dimulai
			// TODO return list of all peserta with their status and arisanID and state
			return nil, nil
		}

		if pesertaObj.StateUndangan == vo.TerimaUndanganStateEnum {
			arisanObj, err := r.outport.FindOneArisan(ctx, pesertaObj.ArisanYgDiikuti)
			if err != nil {
				return nil, err
			}

			if arisanObj == nil {
				return nil, apperror.ArisanTidakDitemukan
			}

			if arisanObj.SudahSelesai() {

				// arisan sudah selesai
				return nil, apperror.ArisanSudahSelesai
			}

			if arisanObj.MasihTerimaPeserta() {

				// peserta join arisan tapi arisan belum dimulai
				// TODO return list of all peserta with their status and arisanID and state
				return nil, nil
			}

			// peserta join arisan dan arisan sudah dimulai
			// TODO return list of all peserta with their status and arisanID and state
			return nil, nil
		}

		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	_ = data

	return res, nil
}
