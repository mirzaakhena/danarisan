package buatarisan

import (
	"context"
	"github.com/mirzaakhena/danarisan/application/apperror"
	"github.com/mirzaakhena/danarisan/domain/entity"
	"github.com/mirzaakhena/danarisan/domain/service"
	"github.com/mirzaakhena/danarisan/domain/vo"
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

		pesertaObj, err := r.outport.FindOnePeserta(ctx, vo.PesertaID(req.PesertaID))
		if err != nil {
			return err
		}

		if pesertaObj == nil {
			return apperror.PesertaTidakDitemukan
		}

		if pesertaObj.IsAdmin && pesertaObj.ArisanYgDiikuti != "" {
			return apperror.PesertaSudahMenjadiAdmin
		}

		arisanObj, err := entity.NewArisan(entity.ArisanRequest{
			GenerateID:      r.outport,
			Nama:            req.NamaArisan,
			SetoranTiapSlot: req.SetoranTiapSlot,
			AdminID:         vo.PesertaID(req.PesertaID),
		})
		if err != nil {
			return err
		}

		_, err = r.outport.SaveArisan(ctx, arisanObj)
		if err != nil {
			return err
		}

		pesertaObj.JadiAdmin(arisanObj.ID)

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
