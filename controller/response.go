package controller

import (
	"github.com/mirzaakhena/danarisan/infrastructure/util"
	"github.com/mirzaakhena/danarisan/application/apperror"
)

type Response struct {
	Success      bool        `json:"success"`
	ErrorCode    string      `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
	Data         interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) string {
	var res Response
	res.Success = true
	res.Data = data
	return util.MustJSON(res)
}

func NewErrorResponse(err error) string {
	var res Response
	res.Success = false

	et, ok := err.(apperror.ErrorWithCode)
	if !ok {
		res.ErrorCode = "UNDEFINED"
		res.ErrorMessage = err.Error()
		return util.MustJSON(res)
	}

	res.ErrorCode = et.Code()
	res.ErrorMessage = et.Error()
	return util.MustJSON(res)
}
