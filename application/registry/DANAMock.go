package registry

import (
	"fmt"
	"github.com/mirzaakhena/danarisan/infrastructure/config"
	"github.com/mirzaakhena/danarisan/infrastructure/log"
	"github.com/mirzaakhena/danarisan/usecase/bayarsetoran"
	"github.com/mirzaakhena/danarisan/usecase/buatarisan"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/mirzaakhena/danarisan/application"
	"github.com/mirzaakhena/danarisan/controller"
	"github.com/mirzaakhena/danarisan/controller/restapi"
	"github.com/mirzaakhena/danarisan/gateway/prod"
	"github.com/mirzaakhena/danarisan/infrastructure/server"
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

	databaseConnectionString := config.GetString("database.connectionstring", "test.db")

	db, err := gorm.Open(sqlite.Open(databaseConnectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	gw := prod.NewSuperGateway(db)

	r.createOrderHandler(gw)
	r.topupHandler(gw)
}

func (r *danaMockRegistry) createOrderHandler(gw *prod.SuperGateway) {
	inport := bayarsetoran.NewUsecase(gw)
	r.Router.POST("/bayarsetoran", controller.Authorized(), restapi.BayarSetoranHandler(inport))
}

func (r *danaMockRegistry) topupHandler(gw *prod.SuperGateway) {
	inport := buatarisan.NewUsecase(gw)
	r.Router.POST("/buatarisan", controller.Authorized(), restapi.BuatArisanHandler(inport))
}


