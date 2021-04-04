package registry

import (
	"fmt"
	"github.com/mirzaakhena/danarisan/infrastructure/config"
	"github.com/mirzaakhena/danarisan/infrastructure/log"
	"github.com/mirzaakhena/danarisan/usecase/buatarisan"
	"github.com/mirzaakhena/danarisan/usecase/bukaaplikasi"
	"github.com/mirzaakhena/danarisan/usecase/kocokundian"
	"github.com/mirzaakhena/danarisan/usecase/mulaiarisan"
	"github.com/mirzaakhena/danarisan/usecase/registerpeserta"
	"github.com/mirzaakhena/danarisan/usecase/setorantidakdibayar"
	"github.com/mirzaakhena/danarisan/usecase/tagihsetoran"
	"github.com/mirzaakhena/danarisan/usecase/undangpeserta"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/mirzaakhena/danarisan/application"
	"github.com/mirzaakhena/danarisan/controller"
	"github.com/mirzaakhena/danarisan/controller/restapi"
	"github.com/mirzaakhena/danarisan/gateway/prod"
	"github.com/mirzaakhena/danarisan/infrastructure/server"
)

type arisanSystemRegistry struct {
	server.GinHTTPHandler
}

func NewArisanSystemRegistry() application.RegistryContract {

	config.InitConfig("config-arisansystem", ".")
	serverPort := config.GetInt("server.port", 8080)

	log.UseRotateFile(
		config.GetString("logfile.path", "."),
		config.GetString("logfile.name", "defaultservice"),
		config.GetInt("logfile.age", 14),
	)

	app := arisanSystemRegistry{ //
		GinHTTPHandler: server.NewGinHTTPHandler(fmt.Sprintf(":%d", serverPort)),
	}

	return &app

}

// RegisterUsecase is implementation of RegistryContract.RegisterUsecase()
func (r *arisanSystemRegistry) RegisterUsecase() {

	databaseConnectionString := config.GetString("database.connectionstring", "test.db")

	db, err := gorm.Open(sqlite.Open(databaseConnectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	gw := prod.NewSuperGateway(db)

	//r.bayarSetoranHandler(gw)
	r.buatArisanHandler(gw)
	r.bukaAplikasiHandler(gw)
	//r.jawabUndanganHandler(gw)
	r.kocokUndianHandler(gw)
	r.mulaiArisanHandler(gw)
	r.registerPesertaHandler(gw)
	r.setoranTidakDibayarHandler(gw)
	r.tagihSetoranHandler(gw)
	r.undangPesertaHandler(gw)
}

func (r *arisanSystemRegistry) bukaAplikasiHandler(gw *prod.SuperGateway) {
	inport := bukaaplikasi.NewUsecase(gw)
	r.Router.GET("/arisan/:pesertaID", controller.Authorized(), restapi.BukaAplikasiHandler(inport))
}

//func (r *arisanSystemRegistry) bayarSetoranHandler(gw *prod.SuperGateway) {
//	inport := bayarsetoran.NewUsecase(gw)
//	r.Router.POST("/bayarsetoran", controller.Authorized(), restapi.BayarSetoranHandler(inport))
//}

func (r *arisanSystemRegistry) buatArisanHandler(gw *prod.SuperGateway) {
	inport := buatarisan.NewUsecase(gw)
	r.Router.POST("/buatarisan", controller.Authorized(), restapi.BuatArisanHandler(inport))
}

//func (r *arisanSystemRegistry) jawabUndanganHandler(gw *prod.SuperGateway) {
//	inport := jawabundangan.NewUsecase(gw)
//	r.Router.POST("/jawabundangan", controller.Authorized(), restapi.JawabUndanganHandler(inport))
//}

func (r *arisanSystemRegistry) kocokUndianHandler(gw *prod.SuperGateway) {
	inport := kocokundian.NewUsecase(gw)
	r.Router.POST("/kocokundian", controller.Authorized(), restapi.KocokUndianHandler(inport))
}

func (r *arisanSystemRegistry) mulaiArisanHandler(gw *prod.SuperGateway) {
	inport := mulaiarisan.NewUsecase(gw)
	r.Router.POST("/mulaiarisan", controller.Authorized(), restapi.MulaiArisanHandler(inport))
}

func (r *arisanSystemRegistry) registerPesertaHandler(gw *prod.SuperGateway) {
	inport := registerpeserta.NewUsecase(gw)
	r.Router.POST("/registerpeserta", controller.Authorized(), restapi.RegisterPesertaHandler(inport))
}

func (r *arisanSystemRegistry) setoranTidakDibayarHandler(gw *prod.SuperGateway) {
	inport := setorantidakdibayar.NewUsecase(gw)
	r.Router.POST("/setorantidakdibayar", controller.Authorized(), restapi.SetoranTidakDibayarHandler(inport))
}

func (r *arisanSystemRegistry) tagihSetoranHandler(gw *prod.SuperGateway) {
	inport := tagihsetoran.NewUsecase(gw)
	r.Router.POST("/tagihsetoran", controller.Authorized(), restapi.TagihSetoranHandler(inport))
}

func (r *arisanSystemRegistry) undangPesertaHandler(gw *prod.SuperGateway) {
	inport := undangpeserta.NewUsecase(gw)
	r.Router.POST("/undangpeserta", controller.Authorized(), restapi.UndangPesertaHandler(inport))
}
