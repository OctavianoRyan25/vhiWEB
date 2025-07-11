package res

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(c *gin.Context, code int, message string, data interface{}) {
	res := &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	c.JSON(code, res)
}
