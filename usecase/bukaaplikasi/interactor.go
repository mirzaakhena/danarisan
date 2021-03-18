package bukaaplikasi

import (
	"context"
	"github.com/mirzaakhena/danarisan/application/apperror"
	"github.com/mirzaakhena/danarisan/domain/service"
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

	err := service.ReadOnly(ctx, r.outport, func(ctx context.Context) (error) {

		pesertaObj, err := r.outport.FindOnePeserta(ctx, req.PesertaID)
		if err != nil {
			return err
		}

		if pesertaObj == nil {
			return apperror.PesertaTidakDitemukan
		}

		res.User = pesertaObj

		arisanObj, err := r.outport.FindOneArisan(ctx, pesertaObj.ArisanID.String())
		if err != nil {
			return nil
		}

		if arisanObj == nil {
			return nil
		}

		listPeserta, err := r.outport.FindAllPeserta(ctx, arisanObj.ID.String())
		if err != nil {
			return err
		}

		arisanObj.ListPeserta = listPeserta

		listSlot, err := r.outport.FindAllSlot(ctx, arisanObj.ID.String())
		if err != nil {
			return err
		}

		arisanObj.ListSlot = listSlot

		listUndian, err := r.outport.FindAllUndian(ctx, arisanObj.ID.String())
		if err != nil {
			return err
		}

		arisanObj.ListUndian = listUndian

		listTagihan, err := r.outport.FindAllTagihanByArisanID(ctx, arisanObj.ID.String())
		if err != nil {
			return err
		}

		arisanObj.ListTagihan = listTagihan

		listJurnal, err := r.outport.FindAllJurnal(ctx, arisanObj.ID.String())
		if err != nil {
			return err
		}

		arisanObj.ListJurnal = listJurnal

		listSaldoAkun, err := r.outport.FindAllSaldoAkun(ctx, arisanObj.ID.String())
		if err != nil {
			return err
		}

		arisanObj.ListSaldoAkun = listSaldoAkun

		res.Arisan = arisanObj

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
