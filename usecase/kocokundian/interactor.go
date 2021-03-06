package kocokundian

import (
	"context"
	"github.com/mirzaakhena/danarisan/application/apperror"
	"github.com/mirzaakhena/danarisan/domain/entity"
	"github.com/mirzaakhena/danarisan/domain/service"
	"github.com/mirzaakhena/danarisan/domain/vo"

	"github.com/mirzaakhena/danarisan/usecase/kocokundian/port"
)

//go:generate mockery --dir port/ --name KocokUndianOutport -output mocks/

type kocokUndianInteractor struct {
	outport port.KocokUndianOutport
}

// NewUsecase ...
func NewUsecase(outputPort port.KocokUndianOutport) port.KocokUndianInport {
	return &kocokUndianInteractor{
		outport: outputPort,
	}
}

// Execute ...
func (r *kocokUndianInteractor) Execute(ctx context.Context, req port.KocokUndianRequest) (*port.KocokUndianResponse, error) {

	res := &port.KocokUndianResponse{}

	err := service.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

		adminObj, err := r.outport.FindOnePeserta(ctx, req.PesertaID)
		if err != nil {
			return err
		}

		if adminObj == nil {
			return apperror.PesertaTidakDitemukan
		}

		if !adminObj.IsAdmin {
			return apperror.PesertaBukanAdmin
		}

		arisanObj, err := r.outport.FindOneArisan(ctx, adminObj.ArisanID.String())

		if arisanObj == nil {
			return apperror.ArisanTidakDitemukan
		}

		undianObj, err := r.outport.FindOneUndian(ctx, arisanObj.ID.String(), arisanObj.PutaranKe)
		if err != nil {
			return err
		}

		if undianObj == nil {
			return apperror.UndianTidakDitemukan
		}

		slotsObj, err := r.outport.FindAllSlotNotWinYet(ctx, arisanObj.ID.String())
		if err != nil {
			return err
		}

		if slotsObj == nil || len(slotsObj) == 0 {
			return apperror.SemuaPesertaSudahMenang
		}

		resRandomNumber, err := r.outport.GetRandomNumber(ctx, port.GetRandomNumberRequest{Length: len(slotsObj)})
		if err != nil {
			return err
		}

		winnerSlot := slotsObj[resRandomNumber.RandomNumber]

		err = winnerSlot.Terpilih(req.TanggalHariIni)
		if err != nil {
			return err
		}

		pesertaObj, err := r.outport.FindOnePeserta(ctx, winnerSlot.PesertaID.String())
		if err != nil {
			return err
		}

		if pesertaObj == nil {
			return apperror.PesertaTidakDitemukan
		}

		totalNilaiUndian := arisanObj.GetTotalNilaiUndian()

		_, err = r.outport.TopupPeserta(ctx, port.TopupPesertaRequest{
			PesertaID:  string(pesertaObj.ID),
			TotalTopup: totalNilaiUndian,
		})
		if err != nil {
			// TODO Jika gagal harus ada mekanisme retry
			return err
		}

		err = r.outport.SaveSlot(ctx, winnerSlot)
		if err != nil {
			return err
		}

		arisanObj.SiapkanArisanBerikutnya()

		err = r.outport.SaveArisan(ctx, arisanObj)
		if err != nil {
			return err
		}

		// JURNAL MENANG UNDIAN
		{
			jurnalObj, err := entity.NewJurnal(entity.JurnalRequest{
				GenerateID: r.outport,
				ArisanID:   arisanObj.ID,
				PesertaID:  pesertaObj.ID,
				Tanggal:    req.TanggalHariIni,
				JurnalType: vo.MenangUndianJurnalTypeEnum,
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

				lastSaldoAkun, err := r.outport.FindLastSaldoAkun(ctx, arisanObj.ID.String(), pesertaObj.ID.String(), vo.HartaAkunTypeEnum.String())
				if err != nil {
					return err
				}

				saldoAkunHartaObj, err := entity.NewSaldoAkun(entity.SaldoAkunRequest{
					Jurnal:              jurnalObj,
					AkunType:            vo.HartaAkunTypeEnum,
					Arah:                entity.ArahBertambah,
					Nominal:             totalNilaiUndian,
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

			// UTANG BERTAMBAH
			{
				lastSaldoAkun, err := r.outport.FindLastSaldoAkun(ctx, arisanObj.ID.String(), pesertaObj.ID.String(), vo.UtangAkunTypeEnum.String())
				if err != nil {
					return err
				}

				saldoAkunModalObj, err := entity.NewSaldoAkun(entity.SaldoAkunRequest{
					Jurnal:              jurnalObj,
					AkunType:            vo.UtangAkunTypeEnum,
					Arah:                entity.ArahBertambah,
					Nominal:             totalNilaiUndian,
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

		// JURNAL PENYESUAIAN
		{

			jurnalObj, err := entity.NewJurnal(entity.JurnalRequest{
				GenerateID: r.outport,
				ArisanID:   arisanObj.ID,
				PesertaID:  pesertaObj.ID,
				Tanggal:    req.TanggalHariIni,
				JurnalType: vo.PenyesuaianJurnalTypeEnum,
			})
			if err != nil {
				return err
			}

			err = r.outport.SaveJurnal(ctx, jurnalObj)
			if err != nil {
				return err
			}

			lastSaldoAkun, err := r.outport.FindLastSaldoAkun(ctx, arisanObj.ID.String(), pesertaObj.ID.String(), vo.PiutangAkunTypeEnum.String())
			if err != nil {
				return err
			}

			nilaiPenyesuaian := lastSaldoAkun.Balance

			// UTANG BERKURANG
			{
				lastSaldoAkun, err := r.outport.FindLastSaldoAkun(ctx, arisanObj.ID.String(), pesertaObj.ID.String(), vo.UtangAkunTypeEnum.String())
				if err != nil {
					return err
				}

				saldoAkunHartaObj, err := entity.NewSaldoAkun(entity.SaldoAkunRequest{
					Jurnal:              jurnalObj,
					AkunType:            vo.UtangAkunTypeEnum,
					Arah:                entity.ArahBerkurang,
					Nominal:             nilaiPenyesuaian,
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

			// PIUTANG BERKURANG
			{
				lastSaldoAkun, err := r.outport.FindLastSaldoAkun(ctx, arisanObj.ID.String(), pesertaObj.ID.String(), vo.PiutangAkunTypeEnum.String())
				if err != nil {
					return err
				}

				saldoAkunModalObj, err := entity.NewSaldoAkun(entity.SaldoAkunRequest{
					Jurnal:              jurnalObj,
					AkunType:            vo.PiutangAkunTypeEnum,
					Arah:                entity.ArahBerkurang,
					Nominal:             nilaiPenyesuaian,
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

		// ARISAN MASIH BERLANJUT
		if !arisanObj.SudahSelesai() {

			tanggalTagihanBerikutnya := undianObj.TanggalTagihan.AddDate(0, 1, 0)
			tanggalUndianBerikutnya := undianObj.TanggalUndian.AddDate(0, 1, 0)

			undianObj, err := entity.NewUndian(entity.UndianRequest{
				ArisanID:       arisanObj.ID,
				PutaranKe:      arisanObj.PutaranKe,
				TanggalTagihan: tanggalTagihanBerikutnya.Format("2006-01-02"),
				TanggalUndian:  tanggalUndianBerikutnya.Format("2006-01-02"),
				BiayaAdmin:     0,
				BiayaArisan:    undianObj.BiayaArisan,
			})
			if err != nil {
				return err
			}

			err = r.outport.SaveUndian(ctx, undianObj)
			if err != nil {
				return err
			}

			slots, err := r.outport.FindAllSlot(ctx, arisanObj.ID.String())
			if err != nil {
				return err
			}

			if slots == nil {
				return nil
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

				err = r.outport.SaveTagihan(ctx, tagihanObj)
				if err != nil {
					return err
				}

			}

		} else

		// ARISAN SUDAH SELESAI
		{
			for _, slot := range slotsObj {

				pesertaObj, err := r.outport.FindOnePeserta(ctx, slot.PesertaID.String())
				if err != nil {
					return err
				}

				if pesertaObj == nil {
					return apperror.PesertaTidakDitemukan
				}

				pesertaObj.ResetPeserta()

				err = r.outport.SavePeserta(ctx, pesertaObj)
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
