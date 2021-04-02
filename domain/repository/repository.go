package repository

import (
	"context"

	"github.com/mirzaakhena/danarisan/domain/entity"
)

type SaveArisanRepo interface {
	SaveArisan(ctx context.Context, obj *entity.Arisan) error
}

type SavePesertaRepo interface {
	SavePeserta(ctx context.Context, obj *entity.Peserta) error
}

type SaveListOfPesertaRepo interface {
	SaveListOfPeserta(ctx context.Context, objs []*entity.Peserta) error
}

type SaveSlotRepo interface {
	SaveSlot(ctx context.Context, obj *entity.Slot) error
}

type SaveUndianRepo interface {
	SaveUndian(ctx context.Context, obj *entity.Undian) error
}

type SaveTagihanRepo interface {
	SaveTagihan(ctx context.Context, obj *entity.Tagihan) error
}

type SaveJurnalRepo interface {
	SaveJurnal(ctx context.Context, obj *entity.Jurnal) error
}

type SaveSaldoAkunRepo interface {
	SaveSaldoAkun(ctx context.Context, obj *entity.SaldoAkun) error
}

type FindOneArisanRepo interface {
	FindOneArisan(ctx context.Context, arisanID string) (*entity.Arisan, error)
}

type FindOneUndianRepo interface {
	FindOneUndian(ctx context.Context, arisanID string, putaranKe int) (*entity.Undian, error)
}

type FindOneUndianByIDRepo interface {
	FindOneUndianByID(ctx context.Context, undianID string) (*entity.Undian, error)
}

type FindAllTagihanRepo interface {
	FindAllTagihan(ctx context.Context, undianID string) ([]*entity.Tagihan, error)
}

type FindAllTagihanByArisanIDRepo interface {
	FindAllTagihanByArisanID(ctx context.Context, arisanID string) ([]*entity.Tagihan, error)
}

type FindOneTagihanRepo interface {
	FindOneTagihan(ctx context.Context, tagihanID string) (*entity.Tagihan, error)
}

type FindAllSlotRepo interface {
	FindAllSlot(ctx context.Context, arisanID string) ([]*entity.Slot, error)
}

type FindOnePesertaRepo interface {
	FindOnePeserta(ctx context.Context, pesertaID string) (*entity.Peserta, error)
}

type FindAllSlotNotWinYetRepo interface {
	FindAllSlotNotWinYet(ctx context.Context, arisanID string) ([]*entity.Slot, error)
}

type FindLastSaldoAkunRepo interface {
	FindLastSaldoAkun(ctx context.Context, arisanID string, pesertaID string, akunType string) (*entity.SaldoAkun, error)
}

type FindOneArisanByAdminIDRepo interface {
	FindOneArisanByAdminID(ctx context.Context, adminID string) (*entity.Arisan, error)
}

type FindPesertaByIDsRepo interface {
	FindPesertaByIDs(ctx context.Context, pesertaID []string) ([]*entity.Peserta, error)
}

type FindAllPesertaRepo interface {
	FindAllPeserta(ctx context.Context, arisanID string) ([]*entity.Peserta, error)
}

type FindAllUndianRepo interface {
	FindAllUndian(ctx context.Context, arisanID string) ([]*entity.Undian, error)
}

type FindAllSaldoAkunRepo interface {
	FindAllSaldoAkun(ctx context.Context, arisanID string) ([]*entity.SaldoAkun, error)
}

type FindAllJurnalRepo interface {
	FindAllJurnal(ctx context.Context, arisanID string) ([]*entity.Jurnal, error)
}

type FindOnePeserta2Repo interface {
	FindOnePeserta2(ctx context.Context, pesertaID string) (*entity.Peserta, error)
}
