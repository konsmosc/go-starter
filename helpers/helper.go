package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, message string, payload interface{}) {
	var res ResponseData

	res.Message = message
	res.Data = payload

	c.JSON(http.StatusOK, res)
}
