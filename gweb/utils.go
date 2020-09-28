package gweb

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type serverResponse struct {
	Success bool
	Data    interface{} `json:",omitempty"`
	Error   string      `json:",omitempty"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, serverResponse{
		Success: true,
		Data:    data,
	})
}

func ResponseError(c *gin.Context, err error) {
	ResponseFailure(c, err, http.StatusInternalServerError)
}

func ResponseBadRequest(c *gin.Context, err error) {
	ResponseFailure(c, err, http.StatusBadRequest)
}

func ResponseUnauthorized(c *gin.Context, err error) {
	ResponseFailure(c, err, http.StatusUnauthorized)
}

func ResponseFailure(c *gin.Context, err error, code int) {
	resp := serverResponse{
		Success: false,
	}
	if err != nil {
		resp.Error = err.Error()
	}
	c.JSON(code, resp)
}
