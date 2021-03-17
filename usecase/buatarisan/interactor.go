package buatarisan

import (
	"context"
	"github.com/mirzaakhena/danarisan/domain/entity"
	"github.com/mirzaakhena/danarisan/domain/service"
	"github.com/mirzaakhena/danarisan/usecase/buatarisan/port"
)

//go:generate mockery --dir port/ --name BuatArisanOutport -output mocks/

type buatArisanInteractor struct {
	outport port.BuatArisanOutport
}

// NewBuatArisanUsecase ...
func NewBuatArisanUsecase(outputPort port.BuatArisanOutport) port.BuatArisanInport {
	return &buatArisanInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *buatArisanInteractor) Execute(ctx context.Context, req port.BuatArisanRequest) (res *port.BuatArisanResponse, err error) {

	res = &port.BuatArisanResponse{}

	err = service.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

		arisanObj, err := entity.NewArisan(entity.ArisanRequest{
			GenerateID:      r.outport,
			Nama:            req.NamaArisan,
			SetoranTiapSlot: req.SetoranTiapSlot,
		})
		if err != nil {
			return err
		}

		_, err = r.outport.SaveArisan(ctx, arisanObj)
		if err != nil {
			return err
		}

		pesertaObj, err := entity.NewPeserta(entity.PesertaRequest{
			GenerateID: r.outport,
			Nama:       req.NamaAdmin,
		})
		if err != nil {
			return err
		}

		_, err = r.outport.SavePeserta(ctx, pesertaObj)
		if err != nil {
			return err
		}

		slotObj, err := entity.NewSlot(entity.SlotRequest{
			GenerateID:  r.outport,
			ArisanID:    arisanObj.ID,
			GroupSlotID: "",
			PesertaID:   pesertaObj.ID,
		})
		if err != nil {
			return err
		}

		_, err = r.outport.SaveSlot(ctx, slotObj)
		if err != nil {
			return err
		}

		undianObj, err := entity.NewUndian(entity.UndianRequest{
			ArisanID:       arisanObj.ID,
			PutaranKe:      1,
			TanggalTagihan: req.TanggalTagihan,
			TanggalUndian:  req.TanggalUndian,
			BiayaAdmin:     0,
			BiayaArisan:    req.BiayaArisan,
		})
		if err != nil {
			return err
		}

		_, err = r.outport.SaveUndian(ctx, undianObj)
		if err != nil {
			return err
		}

		res.ArisanID = string(arisanObj.ID)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
