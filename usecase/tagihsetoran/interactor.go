package tagihsetoran

import (
	"context"
	"github.com/mirzaakhena/danarisan/domain/service"
	"github.com/mirzaakhena/danarisan/domain/vo"
	"github.com/mirzaakhena/danarisan/usecase/tagihsetoran/port"
)

//go:generate mockery --dir port/ --name TagihSetoranOutport -output mocks/

type tagihSetoranInteractor struct {
	outport port.TagihSetoranOutport
}

// NewTagihSetoranUsecase ...
func NewTagihSetoranUsecase(outputPort port.TagihSetoranOutport) port.TagihSetoranInport {
	return &tagihSetoranInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *tagihSetoranInteractor) Execute(ctx context.Context, req port.TagihSetoranRequest) (*port.TagihSetoranResponse, error) {

	res := &port.TagihSetoranResponse{}

	err := service.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

		arisanObj, err := r.outport.FindOneArisan(ctx, vo.ArisanID(req.ArisanID))
		if err != nil {
			return err
		}

		undianObj, err := r.outport.FindOneUndian(ctx, arisanObj.ID, arisanObj.PutaranKe)
		if err != nil {
			return err
		}

		tagihanObjs, err := r.outport.FindAllTagihan(ctx, undianObj.ID)
		if err != nil {
			return err
		}

		for _, tagihanObj := range tagihanObjs {

			totalTagihan := tagihanObj.Nominal + undianObj.BiayaArisan + (undianObj.BiayaAdmin * float64(tagihanObj.JumlahSlot))

			createPaymentReq := port.CreatePaymentRequest{
				ArisanID:           req.ArisanID,
				TagihanID:          string(tagihanObj.ID),
				Nominal:            totalTagihan,
				TanggalKadaluwarsa: undianObj.TanggalUndian,
				Tagihan:            tagihanObj,
			}

			go func(innerReq port.CreatePaymentRequest) {

				createPaymentRes, err := r.outport.CreatePayment(ctx, innerReq)
				if err != nil {
					// TODO jika gagal ada harus mekanisme retry disini
					return
				}

				tagihanObj2 := innerReq.Tagihan

				_ = tagihanObj2.SimpanPenagihan(createPaymentRes.AcquirementID, createPaymentRes.CheckoutURL)

				_, _ = r.outport.SaveTagihan(ctx, &tagihanObj2)

				r.outport.NotifyPeserta(ctx, port.NotifyPesertaRequest{PesertaID: string(tagihanObj.PesertaID)})

			}(createPaymentReq)

		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
