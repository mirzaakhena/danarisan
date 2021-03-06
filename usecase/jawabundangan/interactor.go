package jawabundangan

import (
	"context"
	"github.com/mirzaakhena/danarisan/application/apperror"
	"github.com/mirzaakhena/danarisan/domain/entity"
	"github.com/mirzaakhena/danarisan/domain/service"
	"github.com/mirzaakhena/danarisan/domain/vo"

	"github.com/mirzaakhena/danarisan/usecase/jawabundangan/port"
)

//go:generate mockery --dir port/ --name JawabUndanganOutport -output mocks/

type jawabUndanganInteractor struct {
	outport port.JawabUndanganOutport
}

// NewUsecase ...
func NewUsecase(outputPort port.JawabUndanganOutport) port.JawabUndanganInport {
	return &jawabUndanganInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *jawabUndanganInteractor) Execute(ctx context.Context, req port.JawabUndanganRequest) (*port.JawabUndanganResponse, error) {

	res := &port.JawabUndanganResponse{}

	err := service.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

		pesertaObj, err := r.outport.FindOnePeserta(ctx, req.PesertaID)
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

		err = r.outport.SavePeserta(ctx, pesertaObj)
		if err != nil {
			return err
		}

		arisanObj, err := r.outport.FindOneArisan(ctx, pesertaObj.ArisanID.String())
		if err != nil {
			return err
		}

		if arisanObj == nil {
			return apperror.ArisanTidakDitemukan
		}

		arisanObj.TambahPeserta()

		err = r.outport.SaveArisan(ctx, arisanObj)
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

		err = r.outport.SaveSlot(ctx, slotObj)
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
