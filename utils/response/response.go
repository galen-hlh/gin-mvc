package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequest(c *gin.Context, code int, msg string) {
	Response(c, http.StatusBadRequest, code, msg, gin.H{})
}

func OkRequest(c *gin.Context, code int, msg string, data map[string]interface{}) {
	Response(c, http.StatusOK, code, msg, data)
}

func ServerErrorRequest(c *gin.Context, code int, msg string, data map[string]interface{}) {
	Response(c, http.StatusInternalServerError, code, msg, data)
}

func Response(c *gin.Context, statusCode int, code int, msg string, data map[string]interface{}) {
	res := gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	//if code == http.StatusInternalServerError {
	//	res["debug"] = gin.H{
	//		"msg":  errorMsg,
	//		"info": errorStack,
	//	}
	//}
	c.JSON(statusCode, res)
}
