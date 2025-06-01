package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:    500,
		Message: "请求失败！服务器内部错误！",
		Data:    data,
	})
}

func ResponseError400(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:    400,
		Message: "请求失败！无法找到请求的资源！",
		Data:    data,
	})
}

func ResponseAuthError(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:    407,
		Message: "请求失败！访问的资源需要代理身份验证！",
		Data:    data,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:    200,
		Message: "服务器响应正常",
		Data:    data,
	})
}
