package danamock

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/mirzaakhena/danarisan/controller"
	"github.com/mirzaakhena/danarisan/infrastructure/log"
	"github.com/mirzaakhena/danarisan/infrastructure/util"
	"github.com/mirzaakhena/danarisan/usecase/topup/port"
)

// TopupHandler ...
func TopupHandler(inputPort port.TopupInport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.ContextWithOperationID(c.Request.Context())

		var req port.TopupRequest
		if err := c.BindJSON(&req); err != nil {
			log.ErrorResponse(ctx, err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		req.TanggalHariIni = time.Now()

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