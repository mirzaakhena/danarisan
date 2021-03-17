package restapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mirzaakhena/danarisan/application/apperror"
	"github.com/mirzaakhena/danarisan/controller"
	"github.com/mirzaakhena/danarisan/infrastructure/log"
	"github.com/mirzaakhena/danarisan/infrastructure/util"
	"github.com/mirzaakhena/danarisan/usecase/aplikasicontoh/port"
)

// AplikasiContoh ...
func AplikasiContohHandler(inputPort port.AplikasiContohInport) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := log.ContextWithOperationID(r.Context())

		jsonReq, _ := ioutil.ReadAll(r.Body)

		log.InfoRequest(ctx, string(jsonReq))

		var req port.AplikasiContohRequest

		if err := json.Unmarshal(jsonReq, &req); err != nil {
			newErr := apperror.FailUnmarshalResponseBodyError
			log.ErrorResponse(ctx, newErr)
			http.Error(w, controller.NewErrorResponse(newErr), http.StatusBadRequest)
			return
		}

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.ErrorResponse(ctx, err)
			http.Error(w, controller.NewErrorResponse(err), http.StatusBadRequest)
			return
		}

		log.InfoResponse(ctx, util.MustJSON(res))
		fmt.Fprint(w, controller.NewSuccessResponse(res))

	}
}
