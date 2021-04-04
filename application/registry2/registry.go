package registry2

import (
  "github.com/mirzaakhena/danarisan/application"
  "github.com/mirzaakhena/danarisan/controller"
  "github.com/mirzaakhena/danarisan/gateway/prod"
  "github.com/mirzaakhena/danarisan/infrastructure/server"
  "github.com/mirzaakhena/danarisan/usecase/bayarsetoran2"
  "github.com/mirzaakhena/danarisan/usecase/jawabundangan2"
)

type MyApplicationOne struct {
  Ctrls []application.RegistryContract2
}

func NewMyApplicationOne() application.Runner {

  handler := server.NewGinHTTPHandler(":8080")
  gateway := prod.NewSuperGateway(nil)

  return &MyApplicationOne{
    Ctrls: []application.RegistryContract2{
      controller.MyRouter{
        Router:               handler.Router,
        BayarSetoranUsecase:  bayarsetoran2.NewUsecase(gateway),
        JawabUndanganUsecase: jawabundangan2.NewUsecase(gateway),
      },
      controller.MyRouter{
        Router:               handler.Router,
        BayarSetoranUsecase:  bayarsetoran2.NewUsecase(gateway),
        JawabUndanganUsecase: jawabundangan2.NewUsecase(gateway),
      },
      controller.MyRouter{
        Router:               handler.Router,
        BayarSetoranUsecase:  bayarsetoran2.NewUsecase(gateway),
        JawabUndanganUsecase: jawabundangan2.NewUsecase(gateway),
      },
    },
  }

}

func (r *MyApplicationOne) Start() {

  // TODO registrer by for loop interface
  for _, ctrl := range r.Ctrls {
    ctrl.Register()
  }
}
