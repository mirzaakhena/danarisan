package controller

import (
  "github.com/gin-gonic/gin"
  "github.com/mirzaakhena/danarisan/controller/restapi"
  "github.com/mirzaakhena/danarisan/usecase/bayarsetoran2"
  "github.com/mirzaakhena/danarisan/usecase/jawabundangan2"
)

type MyRouter struct {
  Router gin.IRouter
  BayarSetoranUsecase  bayarsetoran2.Inport
  JawabUndanganUsecase jawabundangan2.Inport
}

func (r MyRouter) Register() {
  r.Router.POST("/sometjing", restapi.BayarSetoranHandler(r.BayarSetoranUsecase))
  r.Router.POST("/whatever", restapi.JawabUndanganHandler(r.JawabUndanganUsecase))
}
