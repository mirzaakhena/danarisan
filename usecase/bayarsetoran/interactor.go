package bayarsetoran

import (
	"context"
	"github.com/mirzaakhena/danarisan/application/apperror"
	"github.com/mirzaakhena/danarisan/domain/entity"
	"github.com/mirzaakhena/danarisan/domain/service"
	"github.com/mirzaakhena/danarisan/domain/vo"
	"github.com/mirzaakhena/danarisan/usecase/bayarsetoran/port"
)

//go:generate mockery --dir port/ --name BayarSetoranOutport -output mocks/

type bayarSetoranInteractor struct {
	outport port.BayarSetoranOutport
}

// NewUsecase ...
func NewUsecase(outputPort port.BayarSetoranOutport) port.BayarSetoranInport {
	return &bayarSetoranInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *bayarSetoranInteractor) Execute(ctx context.Context, req port.BayarSetoranRequest) (*port.BayarSetoranResponse, error) {

	res := &port.BayarSetoranResponse{}

	err := service.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

		tagihanObj, err := r.outport.FindOneTagihan(ctx, req.TagihanID)
		if err != nil {
			return err
		}

		if tagihanObj == nil {
			return apperror.TagihanTidakDitemukan
		}

		tagihanObj.Bayar(req.TanggalHariIni)

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

		pesertaObj.MelakukanPembayaran()

		err = r.outport.SavePeserta(ctx, pesertaObj)
		if err != nil {
			return err
		}

		// JURNAL TAMBAH MODAL
		{

			jurnalObj, err := entity.NewJurnal(entity.JurnalRequest{
				GenerateID: r.outport,
				ArisanID:   tagihanObj.ArisanID,
				PesertaID:  pesertaObj.ID,
				Tanggal:    req.TanggalHariIni,
				JurnalType: vo.TambahModalJurnalTypeEnum,
			})
			if err != nil {
				return err
			}

			err = r.outport.SaveJurnal(ctx, jurnalObj)
			if err != nil {
				return err
			}

			// HARTA BERTAMBAH
			{
				lastSaldoAkun, err := r.outport.FindLastSaldoAkun(ctx, tagihanObj.ArisanID.String(), pesertaObj.ID.String(), vo.HartaAkunTypeEnum.String())
				if err != nil {
					return err
				}

				saldoAkunHartaObj, err := entity.NewSaldoAkun(entity.SaldoAkunRequest{
					Jurnal:              jurnalObj,
					AkunType:            vo.HartaAkunTypeEnum,
					Arah:                entity.ArahBertambah,
					Nominal:             tagihanObj.Nominal,
					SaldoAkunSebelumnya: lastSaldoAkun,
					Sequence:            1,
				})
				if err != nil {
					return err
				}

				err = r.outport.SaveSaldoAkun(ctx, saldoAkunHartaObj)
				if err != nil {
					return err
				}
			}

			// MODAL BERTAMBAH
			{
				lastSaldoAkun, err := r.outport.FindLastSaldoAkun(ctx, tagihanObj.ArisanID.String(), pesertaObj.ID.String(), vo.ModalAkunTypeEnum.String())
				if err != nil {
					return err
				}

				saldoAkunModalObj, err := entity.NewSaldoAkun(entity.SaldoAkunRequest{
					Jurnal:              jurnalObj,
					AkunType:            vo.ModalAkunTypeEnum,
					Arah:                entity.ArahBertambah,
					Nominal:             tagihanObj.Nominal,
					SaldoAkunSebelumnya: lastSaldoAkun,
					Sequence:            2,
				})
				if err != nil {
					return err
				}

				err = r.outport.SaveSaldoAkun(ctx, saldoAkunModalObj)
				if err != nil {
					return err
				}
			}

		}

		// JURNAL SETOR TAGIHAN
		{

			jurnalObj, err := entity.NewJurnal(entity.JurnalRequest{
				GenerateID: r.outport,
				ArisanID:   tagihanObj.ArisanID,
				PesertaID:  pesertaObj.ID,
				Tanggal:    req.TanggalHariIni,
				JurnalType: vo.SetorTagihanJurnalTypeEnum,
			})
			if err != nil {
				return err
			}

			err = r.outport.SaveJurnal(ctx, jurnalObj)
			if err != nil {
				return err
			}

			// HARTA BERTAMBAH
			{
				lastSaldoAkun, err := r.outport.FindLastSaldoAkun(ctx, tagihanObj.ArisanID.String(), pesertaObj.ID.String(), vo.HartaAkunTypeEnum.String())
				if err != nil {
					return err
				}

				saldoAkunHartaObj, err := entity.NewSaldoAkun(entity.SaldoAkunRequest{
					Jurnal:              jurnalObj,
					AkunType:            vo.HartaAkunTypeEnum,
					Arah:                entity.ArahBerkurang,
					Nominal:             tagihanObj.Nominal,
					SaldoAkunSebelumnya: lastSaldoAkun,
					Sequence:            3,
				})
				if err != nil {
					return err
				}

				err = r.outport.SaveSaldoAkun(ctx, saldoAkunHartaObj)
				if err != nil {
					return err
				}
			}

			// PIUTANG BERTAMBAH
			{
				lastSaldoAkun, err := r.outport.FindLastSaldoAkun(ctx, tagihanObj.ArisanID.String(), pesertaObj.ID.String(), vo.PiutangAkunTypeEnum.String())
				if err != nil {
					return err
				}

				saldoAkunModalObj, err := entity.NewSaldoAkun(entity.SaldoAkunRequest{
					Jurnal:              jurnalObj,
					AkunType:            vo.PiutangAkunTypeEnum,
					Arah:                entity.ArahBertambah,
					Nominal:             tagihanObj.Nominal,
					SaldoAkunSebelumnya: lastSaldoAkun,
					Sequence:            4,
				})
				if err != nil {
					return err
				}

				err = r.outport.SaveSaldoAkun(ctx, saldoAkunModalObj)
				if err != nil {
					return err
				}
			}

		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
