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

		arisanObj, err := r.outport.FindOneArisan(ctx, pesertaObj.ArisanID)
		if err != nil {
			return nil, err
		}

		listPeserta, err := r.outport.FindAllPeserta(ctx, arisanObj.ID)
		if err != nil {
			return nil, err
		}

		arisanObj.ListPeserta = listPeserta

		listSlot, err := r.outport.FindAllSlot(ctx, arisanObj.ID)
		if err != nil {
			return nil, err
		}

		arisanObj.ListSlot = listSlot

		listUndian, err := r.outport.FindAllUndian(ctx, arisanObj.ID)
		if err != nil {
			return nil, err
		}

		arisanObj.ListUndian = listUndian

		listTagihan, err := r.outport.FindAllTagihanByArisanID(ctx, arisanObj.ID)
		if err != nil {
			return nil, err
		}

		arisanObj.ListTagihan = listTagihan

		listJurnal, err := r.outport.FindAllJurnal(ctx, arisanObj.ID)
		if err != nil {
			return nil, err
		}

		arisanObj.ListJurnal = listJurnal

		listSaldoAkun, err := r.outport.FindAllSaldoAkun(ctx, arisanObj.ID)
		if err != nil {
			return nil, err
		}

		arisanObj.ListSaldoAkun = listSaldoAkun

		//if pesertaObj.StateUndangan == vo.DitawarkanUndanganStateEnum {
		//	arisanObj, err := r.outport.FindOneArisan(ctx, pesertaObj.ArisanID)
		//	if err != nil {
		//		return nil, err
		//	}
		//
		//	if arisanObj == nil {
		//		return nil, apperror.ArisanTidakDitemukan
		//	}
		//
		//	// peserta belum join arisan dan arisan belum dimulai
		//	// TODO return list of all peserta with their status and arisanID and state
		//	return nil, nil
		//}
		//
		//if pesertaObj.StateUndangan == vo.TerimaUndanganStateEnum {
		//	arisanObj, err := r.outport.FindOneArisan(ctx, pesertaObj.ArisanID)
		//	if err != nil {
		//		return nil, err
		//	}
		//
		//	if arisanObj == nil {
		//		return nil, apperror.ArisanTidakDitemukan
		//	}
		//
		//	if arisanObj.SudahSelesai() {
		//
		//		// arisan sudah selesai
		//		return nil, apperror.ArisanSudahSelesai
		//	}
		//
		//	if arisanObj.MasihTerimaPeserta() {
		//
		//		// peserta join arisan tapi arisan belum dimulai
		//		// TODO return list of all peserta with their status and arisanID and state
		//		return nil, nil
		//	}
		//
		//	// peserta join arisan dan arisan sudah dimulai
		//	// TODO return list of all peserta with their status and arisanID and state
		//	return nil, nil
		//}

		return arisanObj, nil
	})
	if err != nil {
		return nil, err
	}

	res.Data = data

	return res, nil
}
