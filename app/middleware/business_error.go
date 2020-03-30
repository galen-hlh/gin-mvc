package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go-restful-api/app/compoments/errors"
	"go-restful-api/app/compoments/response"
	"net/http"
	"runtime/debug"
)

// 404 处理
func HandleNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, response.HttpResponse{Code: http.StatusNotFound, Msg: "not found"})
}

// 405 处理
func HandleNotMethod(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, response.HttpResponse{Code: http.StatusMethodNotAllowed, Msg: "method not allowed"})
}

// 400 业务异常
func HandleBusinessError(c *gin.Context, b *errors.Business) {
	c.JSON(http.StatusBadRequest, response.HttpResponse{Code: b.Code(), Msg: b.Error()})
}

// 500 处理
func HandleInternalServerError(c *gin.Context, stack []string) {
	c.JSON(http.StatusInternalServerError, response.HttpResponse{Code: http.StatusInternalServerError, Msg: "internal server error"})
}

// 代码异常处理中间件
func HandleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {

				//如果是业务异常
				b, ok := interface{}(r).(*errors.Business)
				if ok {
					HandleBusinessError(c, b)
					return
				}

				//获取错误堆栈
				buf := debug.Stack()
				var errorStack []string
				stack := bytes.Split(buf, []byte("\n"))
				for _, bt := range stack {
					info := string(bt)
					if info != "" {
						errorStack = append(errorStack, info)
					}
				}

				HandleInternalServerError(c, errorStack)
			}
		}()
		c.Next()
	}
}
