package testcase

import (
	"context"
	"errors"
	"github.com/mirzaakhena/danarisan/domain/entity"
	"github.com/mirzaakhena/danarisan/infrastructure/log"
	"github.com/mirzaakhena/danarisan/usecase/registerpeserta"
	"github.com/mirzaakhena/danarisan/usecase/registerpeserta/port"
)

type mockOutport struct{}

func (m mockOutport) SavePeserta(ctx context.Context, obj *entity.Peserta) error {
	panic("implement me")
}

func (m mockOutport) BeginTransaction(ctx context.Context) (context.Context, error) {
	panic("implement me")
}

func (m mockOutport) CommitTransaction(ctx context.Context) error {
	panic("implement me")
}

func (m mockOutport) RollbackTransaction(ctx context.Context) error {
	panic("implement me")
}

func newMockOutport() port.RegisterPesertaOutport {
	return &mockOutport{}
}

func Register() {

	usecases := map[string]func(interface{}) error{

		"registerpeserta": func(data interface{}) error {

			outport := newMockOutport()
			inport := registerpeserta.NewUsecase(outport)

			ctx := log.ContextWithOperationID(context.Background())

			request, ok := data.(port.RegisterPesertaRequest)
			if !ok {
				return errors.New("cannot assert data to RegisterPesertaRequest")
			}

			res, err := inport.Execute(ctx, request)
			if err != nil {
				return err
			}
			_ = res

			return nil
		},
	}

	_ = usecases

}
