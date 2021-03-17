package repository

import (
	"context"

	"github.com/mirzaakhena/danarisan/domain/vo"

	"github.com/mirzaakhena/danarisan/domain/entity"
)

type SaveArisanRepo interface {
	SaveArisan(ctx context.Context, obj *entity.Arisan) (*entity.Arisan, error)
}

type SavePesertaRepo interface {
	SavePeserta(ctx context.Context, obj *entity.Peserta) (*entity.Peserta, error)
}

type SaveSlotRepo interface {
	SaveSlot(ctx context.Context, obj *entity.Slot) (*entity.Slot, error)
}

type SaveUndianRepo interface {
	SaveUndian(ctx context.Context, obj *entity.Undian) (*entity.Undian, error)
}

type SaveTagihanRepo interface {
	SaveTagihan(ctx context.Context, obj *entity.Tagihan) (*entity.Tagihan, error)
}

type SaveJurnalRepo interface {
	SaveJurnal(ctx context.Context, obj *entity.Jurnal) (*entity.Jurnal, error)
}

type SaveSaldoAkunRepo interface {
	SaveSaldoAkun(ctx context.Context, obj *entity.SaldoAkun) (*entity.SaldoAkun, error)
}

//type CountPesertaRepo interface {
//	CountPeserta(ctx context.Context, obj *entity.Arisan) (*entity.Arisan, error)
//}

type FindOneArisanRepo interface {
	FindOneArisan(ctx context.Context, arisanID vo.ArisanID) (*entity.Arisan, error)
}

type FindOneUndianRepo interface {
	FindOneUndian(ctx context.Context, arisanID vo.ArisanID, putaranKe int) (*entity.Undian, error)
}

type FindOneUndianByIDRepo interface {
	FindOneUndianByID(ctx context.Context, undianID vo.UndianID) (*entity.Undian, error)
}

type FindAllTagihanRepo interface {
	FindAllTagihan(ctx context.Context, undianID vo.UndianID) ([]entity.Tagihan, error)
}

type FindAllTagihanByArisanIDRepo interface {
	FindAllTagihanByArisanID(ctx context.Context, arisanID vo.ArisanID) ([]entity.Tagihan, error)
}

type FindOneTagihanRepo interface {
	FindOneTagihan(ctx context.Context, tagihanID vo.TagihanID) (*entity.Tagihan, error)
}

type FindAllSlotRepo interface {
	FindAllSlot(ctx context.Context, arisanID vo.ArisanID) ([]entity.Slot, error)
}

type FindOnePesertaRepo interface {
	FindOnePeserta(ctx context.Context, pesertaID vo.PesertaID) (*entity.Peserta, error)
}

type FindAllSlotNotWinYetRepo interface {
	FindAllSlotNotWinYet(ctx context.Context, arisanID vo.ArisanID) ([]entity.Slot, error)
}

type FindLastSaldoAkunRepo interface {
	FindLastSaldoAkun(ctx context.Context, arisanID vo.ArisanID, pesertaID vo.PesertaID, akunType vo.AkunType) (*entity.SaldoAkun, error)
}

type FindOneArisanByAdminIDRepo interface {
	FindOneArisanByAdminID(ctx context.Context, adminID vo.PesertaID) (*entity.Arisan, error)
}

type FindAllPesertaRepo interface {
	FindAllPeserta(ctx context.Context, arisanID vo.ArisanID) ([]entity.Peserta, error)
}

type FindAllUndianRepo interface {
	FindAllUndian(ctx context.Context, arisanID vo.ArisanID) ([]entity.Undian, error)
}
