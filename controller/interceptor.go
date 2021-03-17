package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorized() gin.HandlerFunc {

	return func(c *gin.Context) {

		authorized := true

		if !authorized {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

	}

}
