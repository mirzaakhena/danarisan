package restapi

import (
  "context"
  "flag"
  "github.com/mirzaakhena/danarisan/infrastructure/log"
  "github.com/mirzaakhena/danarisan/infrastructure/util"
  "github.com/mirzaakhena/danarisan/usecase/bayarsetoran2"
  "time"
)

// BayarSetoran2Handler ...
func BayarSetoranHandler2(inputPort bayarsetoran2.Inport) func(msg []byte ) {

  return func(msg []byte) {

    a := flag.Arg(1)

    ctx := log.ContextWithOperationID(context.Background())

    var req bayarsetoran2.InportRequest
    req.TagihanID = a
    req.TanggalHariIni = time.Now()

    log.InfoRequest(ctx, util.MustJSON(req))

    res, err := inputPort.Execute(ctx, req)

    if err != nil {
      log.ErrorResponse(ctx, err)
      return
    }

    log.InfoResponse(ctx, util.MustJSON(res))

  }
}
