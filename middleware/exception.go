package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

// 404 页面处理
func PageNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": "Page not found"})
}

// 代码异常处理中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//获取错误堆栈
				buf := debug.Stack()

				// 转为数组
				var errorStack []string
				errorMsg := r.(string)
				stack := bytes.Split(buf, []byte("\n"))
				for _, bt := range stack {
					info := string(bt)
					if info != "" {
						errorStack = append(errorStack, info)
					}
				}

				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  "",
					"debug": gin.H{
						"msg":  errorMsg,
						"info": errorStack,
					},
				})
			}
		}()
		c.Next()
	}
}
