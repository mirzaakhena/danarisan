package repoimpl

import (
	"context"
	"github.com/mirzaakhena/danarisan/domain/entity"
	"github.com/mirzaakhena/danarisan/domain/vo"
	"github.com/mirzaakhena/danarisan/infrastructure/log"
	"github.com/mirzaakhena/danarisan/infrastructure/util"
	"gorm.io/gorm"
	"sync"
)

var onceRepoImpl sync.Once

var repoImplObj RepositoryImplementation

type RepositoryImplementation struct {
}

func SingletonRepositoryImplementation(db *gorm.DB) *RepositoryImplementation {

	onceRepoImpl.Do(func() {

		db.AutoMigrate(&entity.Arisan{})
		db.AutoMigrate(&entity.GroupSlot{})
		db.AutoMigrate(&entity.Jurnal{})
		db.AutoMigrate(&entity.Peserta{})
		db.AutoMigrate(&entity.SaldoAkun{})
		db.AutoMigrate(&entity.Slot{})
		db.AutoMigrate(&entity.Tagihan{})
		db.AutoMigrate(&entity.Undian{})

		repoImplObj = RepositoryImplementation{}
	})

	return &repoImplObj
}

func (r *RepositoryImplementation) SaveArisan(ctx context.Context, obj *entity.Arisan) (*entity.Arisan, error) {

	log.InfoRequest(ctx, util.MustJSON(obj))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	err = db.Save(obj).Error
	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, "Saved")

	return obj, nil
}

func (r *RepositoryImplementation) SavePeserta(ctx context.Context, obj *entity.Peserta) (*entity.Peserta, error) {

	log.InfoRequest(ctx, util.MustJSON(obj))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	err = db.Save(obj).Error
	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, "Saved")

	return obj, nil
}

func (r *RepositoryImplementation) SaveSlot(ctx context.Context, obj *entity.Slot) (*entity.Slot, error) {

	log.InfoRequest(ctx, util.MustJSON(obj))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	err = db.Save(obj).Error
	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, "Saved")

	return obj, nil
}

func (r *RepositoryImplementation) SaveUndian(ctx context.Context, obj *entity.Undian) (*entity.Undian, error) {
	log.InfoRequest(ctx, util.MustJSON(obj))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	err = db.Save(obj).Error
	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, "Saved")

	return obj, nil
}

func (r *RepositoryImplementation) SaveTagihan(ctx context.Context, obj *entity.Tagihan) (*entity.Tagihan, error) {
	log.InfoRequest(ctx, util.MustJSON(obj))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	err = db.Save(obj).Error
	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, "Saved")

	return obj, nil
}

func (r *RepositoryImplementation) SaveJurnal(ctx context.Context, obj *entity.Jurnal) (*entity.Jurnal, error) {
	log.InfoRequest(ctx, util.MustJSON(obj))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	err = db.Save(obj).Error
	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, "Saved")

	return obj, nil
}

func (r *RepositoryImplementation) SaveSaldoAkun(ctx context.Context, obj *entity.SaldoAkun) (*entity.SaldoAkun, error) {
	log.InfoRequest(ctx, util.MustJSON(obj))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	err = db.Save(obj).Error
	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, "Saved")

	return obj, nil
}

func (r *RepositoryImplementation) FindOneArisan(ctx context.Context, arisanID vo.ArisanID) (*entity.Arisan, error) {

	log.InfoRequest(ctx, util.MustJSON(map[string]interface{}{"arisanID": arisanID}))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	var arisanObj entity.Arisan
	err = db.
		Where("id = ?", arisanID).
		First(&arisanObj).Error

	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, util.MustJSON(arisanObj))

	return &arisanObj, nil
}

func (r *RepositoryImplementation) FindOneUndian(ctx context.Context, arisanID vo.ArisanID, putaranKe int) (*entity.Undian, error) {

	log.InfoRequest(ctx, util.MustJSON(map[string]interface{}{"arisanID": arisanID, "putaranKe": putaranKe}))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	var undianObj entity.Undian
	err = db.
		Where("arisan_id = ? AND putaran_ke = ?", arisanID, putaranKe).
		First(&undianObj).Error

	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, util.MustJSON(undianObj))

	return &undianObj, nil

}

func (r *RepositoryImplementation) FindOneUndianByID(ctx context.Context, undianID vo.UndianID) (*entity.Undian, error) {

	log.InfoRequest(ctx, util.MustJSON(map[string]interface{}{"undianID": undianID}))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	var undianObj entity.Undian
	err = db.
		Where("id = ?", undianID).
		First(&undianObj).Error

	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, util.MustJSON(undianObj))

	return &undianObj, nil
}

func (r *RepositoryImplementation) FindOneTagihan(ctx context.Context, tagihanID vo.TagihanID) (*entity.Tagihan, error) {

	log.InfoRequest(ctx, util.MustJSON(map[string]interface{}{"tagihanID": tagihanID}))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	var tagihanObj entity.Tagihan
	err = db.
		Where("id = ?", tagihanID).
		First(&tagihanObj).Error

	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, util.MustJSON(tagihanObj))

	return &tagihanObj, nil

}

func (r *RepositoryImplementation) FindOnePeserta(ctx context.Context, pesertaID vo.PesertaID) (*entity.Peserta, error) {

	log.InfoRequest(ctx, util.MustJSON(map[string]interface{}{"pesertaID": pesertaID}))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	var pesertaObj entity.Peserta
	err = db.
		Where("id = ?", pesertaID).
		First(&pesertaObj).Error

	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, util.MustJSON(pesertaObj))

	return &pesertaObj, nil
}

func (r *RepositoryImplementation) FindOneArisanByAdminID(ctx context.Context, adminID vo.PesertaID) (*entity.Arisan, error) {

	log.InfoRequest(ctx, util.MustJSON(map[string]interface{}{"adminID": adminID}))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	var arisanObj entity.Arisan
	err = db.
		Where("admin_id = ?", adminID).
		First(&arisanObj).Error

	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, util.MustJSON(arisanObj))

	return &arisanObj, nil

}

func (r *RepositoryImplementation) FindAllTagihan(ctx context.Context, undianID vo.UndianID) ([]entity.Tagihan, error) {

	log.InfoRequest(ctx, util.MustJSON(map[string]interface{}{"undianID": undianID}))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	var tagihanObjs []entity.Tagihan
	err = db.
		Where("undian_id = ?", undianID).
		Find(&tagihanObjs).Error

	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, util.MustJSON(tagihanObjs))

	return tagihanObjs, nil

}

func (r *RepositoryImplementation) FindAllSlot(ctx context.Context, arisanID vo.ArisanID) ([]entity.Slot, error) {

	log.InfoRequest(ctx, util.MustJSON(map[string]interface{}{"arisanID": arisanID}))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	var slotObjs []entity.Slot
	err = db.
		Where("arisan_id = ?", arisanID).
		Find(&slotObjs).Error

	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, util.MustJSON(slotObjs))

	return slotObjs, nil

}

func (r *RepositoryImplementation) FindAllSlotNotWinYet(ctx context.Context, arisanID vo.ArisanID) ([]entity.Slot, error) {

	log.InfoRequest(ctx, util.MustJSON(map[string]interface{}{"arisanID": arisanID}))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	var slotObjs []entity.Slot
	err = db.
		Where("arisan_id = ? AND tanggal_menang IS NULL", arisanID).
		Find(&slotObjs).Error

	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, util.MustJSON(slotObjs))

	return slotObjs, nil
}

func (r *RepositoryImplementation) FindLastSaldoAkun(ctx context.Context, arisanID vo.ArisanID, pesertaID vo.PesertaID, akunType vo.AkunType) (*entity.SaldoAkun, error) {

	log.InfoRequest(ctx, util.MustJSON(map[string]interface{}{"arisanID": arisanID, "pesertaID": pesertaID, "akunType": akunType}))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	var saldoAkunObj entity.SaldoAkun
	err = db.
		Where("arisan_id = ? AND peserta_id = ? AND akun_type = ?", arisanID, pesertaID, akunType).
		Order("tanggal desc, sequence desc").
		First(&saldoAkunObj).Error
	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, util.MustJSON(saldoAkunObj))

	return &saldoAkunObj, nil

}

func (r *RepositoryImplementation) FindAllPeserta(ctx context.Context, arisanID vo.ArisanID) ([]entity.Peserta, error) {

	log.InfoRequest(ctx, util.MustJSON(map[string]interface{}{"arisanID": arisanID}))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	var pesertaObjs []entity.Peserta
	err = db.
		Where("arisan_id = ?", arisanID).
		Find(&pesertaObjs).Error

	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, util.MustJSON(pesertaObjs))

	return pesertaObjs, nil

}

func (r *RepositoryImplementation) FindAllUndian(ctx context.Context, arisanID vo.ArisanID) ([]entity.Undian, error) {

	log.InfoRequest(ctx, util.MustJSON(map[string]interface{}{"arisanID": arisanID}))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	var undianObjs []entity.Undian
	err = db.
		Where("arisan_id = ?", arisanID).
		Find(&undianObjs).Error

	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, util.MustJSON(undianObjs))

	return undianObjs, nil

}

func (r *RepositoryImplementation) FindAllTagihanByArisanID(ctx context.Context, arisanID vo.ArisanID) ([]entity.Tagihan, error) {

	log.InfoRequest(ctx, util.MustJSON(map[string]interface{}{"arisanID": arisanID}))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	var tagihanObjs []entity.Tagihan
	err = db.
		Where("arisan_id = ?", arisanID).
		Find(&tagihanObjs).Error

	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, util.MustJSON(tagihanObjs))

	return tagihanObjs, nil

}

func (r *RepositoryImplementation) FindAllJurnal(ctx context.Context, arisanID vo.ArisanID) ([]entity.Jurnal, error) {

	log.InfoRequest(ctx, util.MustJSON(map[string]interface{}{"arisanID": arisanID}))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	var jurnalObjs []entity.Jurnal
	err = db.
		Where("arisan_id = ?", arisanID).
		Find(&jurnalObjs).Error

	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, util.MustJSON(jurnalObjs))

	return jurnalObjs, nil

}

func (r *RepositoryImplementation) FindAllSaldoAkun(ctx context.Context, arisanID vo.ArisanID) ([]entity.SaldoAkun, error) {

	log.InfoRequest(ctx, util.MustJSON(map[string]interface{}{"arisanID": arisanID}))

	db, err := extractDB(ctx)
	if err != nil {
		return nil, err
	}

	var saldoAkunObjs []entity.SaldoAkun
	err = db.
		Where("arisan_id = ?", arisanID).
		Find(&saldoAkunObjs).Error

	if err != nil {
		log.ErrorResponse(ctx, err)
		return nil, err
	}

	log.InfoResponse(ctx, util.MustJSON(saldoAkunObjs))

	return saldoAkunObjs, nil

}
