package controller

import (
  "github.com/mirzaakhena/danarisan/controller/restapi"
  "github.com/mirzaakhena/danarisan/infrastructure/server"
  "github.com/mirzaakhena/danarisan/usecase/bayarsetoran/port"
)

type MyController struct {
  server.GinHTTPHandler
  BayarSetoranUsecase port.BayarSetoranInport

}

func (r MyController) RouteRegister()  {
  r.Router.POST("", restapi.BayarSetoranHandler(nil))
  r.Router.POST("", restapi.JawabUndanganHandler(nil))
  r.Router.POST("", restapi.RegisterPesertaHandler(nil))
}
