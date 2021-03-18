package registry

import (
	"fmt"
	"github.com/mirzaakhena/danarisan/application"
	"github.com/mirzaakhena/danarisan/controller"
	"github.com/mirzaakhena/danarisan/controller/danamock"
	"github.com/mirzaakhena/danarisan/gateway/mockserver"
	"github.com/mirzaakhena/danarisan/infrastructure/config"
	"github.com/mirzaakhena/danarisan/infrastructure/log"
	"github.com/mirzaakhena/danarisan/infrastructure/server"
	"github.com/mirzaakhena/danarisan/usecase/createpayment"
	"github.com/mirzaakhena/danarisan/usecase/topup"
)

type danaMockRegistry struct {
	server.GinHTTPHandler
}

func NewDANAMockRegistry() application.RegistryContract {

	config.InitConfig("config-danamock", ".")
	serverPort := config.GetInt("server.port", 8081)

	log.UseRotateFile(
		config.GetString("logfile.path", "."),
		config.GetString("logfile.name", "defaultservice"),
		config.GetInt("logfile.age", 14),
	)

	app := danaMockRegistry{ //
		GinHTTPHandler: server.NewGinHTTPHandler(fmt.Sprintf(":%d", serverPort)),
	}

	return &app

}

// RegisterUsecase is implementation of RegistryContract.RegisterUsecase()
func (r *danaMockRegistry) RegisterUsecase() {
	r.createOrderHandler()
	r.topupHandler()
}


func (r *danaMockRegistry) createOrderHandler() {
	outport := mockserver.NewCreatePaymentGateway()
	inport := createpayment.NewUsecase(outport)
	r.Router.POST("/createpayment", controller.Authorized(), danamock.CreatePaymentHandler(inport))
}

func (r *danaMockRegistry) topupHandler() {
	outport := mockserver.NewTopupGateway()
	inport := topup.NewUsecase(outport)
	r.Router.POST("/topup", controller.Authorized(), danamock.TopupHandler(inport))
}
