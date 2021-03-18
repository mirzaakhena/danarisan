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

	_, err := service.ReadOnly(ctx, r.outport, func(ctx context.Context) (interface{}, error) {

		pesertaObj, err := r.outport.FindOnePeserta(ctx, vo.PesertaID(req.PesertaID))
		if err != nil {
			return nil, err
		}

		if pesertaObj == nil {
			return nil, apperror.PesertaTidakDitemukan
		}

		res.User = pesertaObj

		arisanObj, err := r.outport.FindOneArisan(ctx, pesertaObj.ArisanID)
		if err != nil {
			return nil, nil
		}

		if arisanObj == nil {
			return nil, nil
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

		res.Data = arisanObj

		return arisanObj, nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
