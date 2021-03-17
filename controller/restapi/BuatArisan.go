package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/mirzaakhena/danarisan/controller"
	"github.com/mirzaakhena/danarisan/infrastructure/log"
	"github.com/mirzaakhena/danarisan/infrastructure/util"
	"github.com/mirzaakhena/danarisan/usecase/buatarisan/port"
	"net/http"
)

// BuatArisan ...

// BuatArisanHandler ...
func BuatArisanHandler(inputPort port.BuatArisanInport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.ContextWithOperationID(c.Request.Context())

		var req port.BuatArisanRequest
		if err := c.BindJSON(&req); err != nil {
			log.ErrorResponse(ctx, err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

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
