package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/mirzaakhena/danarisan/controller"
	"github.com/mirzaakhena/danarisan/infrastructure/log"
	"github.com/mirzaakhena/danarisan/infrastructure/util"
	"github.com/mirzaakhena/danarisan/usecase/bukaaplikasi/port"
	"net/http"
)

// BukaAplikasiHandler ...
func BukaAplikasiHandler(inputPort port.BukaAplikasiInport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.ContextWithOperationID(c.Request.Context())

		var req port.BukaAplikasiRequest
		req.PesertaID = c.Param("pesertaID")

		log.InfoRequest(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)

		if err != nil {
			log.ErrorResponse(ctx, err)
			c.JSON(http.StatusBadRequest, controller.NewErrorResponse(err))
			return
		}

		log.InfoResponse(ctx, util.MustJSON(res))
		c.JSON(http.StatusOK, res)
	}
}
