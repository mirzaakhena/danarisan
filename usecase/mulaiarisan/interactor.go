package mulaiarisan

import (
	"context"
	"github.com/mirzaakhena/danarisan/application/apperror"
	"github.com/mirzaakhena/danarisan/domain/entity"
	"github.com/mirzaakhena/danarisan/domain/service"
	"github.com/mirzaakhena/danarisan/domain/vo"

	"github.com/mirzaakhena/danarisan/usecase/mulaiarisan/port"
)

//go:generate mockery --dir port/ --name MulaiArisanOutport -output mocks/

type mulaiArisanInteractor struct {
	outport port.MulaiArisanOutport
}

// NewMulaiArisanUsecase ...
func NewMulaiArisanUsecase(outputPort port.MulaiArisanOutport) port.MulaiArisanInport {
	return &mulaiArisanInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *mulaiArisanInteractor) Execute(ctx context.Context, req port.MulaiArisanRequest) (*port.MulaiArisanResponse, error) {

	res := &port.MulaiArisanResponse{}

	err := service.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

		arisanObj, err := r.outport.FindOneArisan(ctx, vo.ArisanID(req.ArisanID))
		if err != nil {
			return err
		}

		if arisanObj == nil {
			return apperror.ArisanTidakDitemukan
		}

		err = arisanObj.Mulai()
		if err != nil {
			return err
		}

		_, err = r.outport.SaveArisan(ctx, arisanObj)
		if err != nil {
			return err
		}

		undianObj, err := r.outport.FindOneUndian(ctx, arisanObj.ID, arisanObj.PutaranKe)
		if err != nil {
			return err
		}

		if undianObj == nil {
			return apperror.UndianTidakDitemukan
		}

		slots, err := r.outport.FindAllSlot(ctx, arisanObj.ID)
		if err != nil {
			return err
		}

		for _, sp := range slots {
			tagihanObj, err := entity.NewTagihan(entity.TagihanRequest{
				ArisanID:   arisanObj.ID,
				UndianID:   undianObj.ID,
				PesertaID:  sp.PesertaID,
				Nominal:    arisanObj.SetoranTiapSlot,
				JumlahSlot: 1,
			})
			if err != nil {
				return err
			}

			_, err = r.outport.SaveTagihan(ctx, tagihanObj)
			if err != nil {
				return err
			}

		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

type tagihanPeserta struct {
	Nominal    float64
	JumlahSlot int
}
