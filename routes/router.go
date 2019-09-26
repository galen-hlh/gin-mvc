package routes

import (
	"fmt"
	"gin-mvc/middleware"
	"gin-mvc/modules/v1/controllers/site"
	_ "gin-mvc/validate"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine
var RouterGroupV1 *gin.RouterGroup
var RouterGroupV2 *gin.RouterGroup

func init() {
	Router = gin.New()

	// 设置运行模式
	gin.SetMode(gin.DebugMode)

	//设置日志格式
	Router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[HTTP] %s | %d |   %s|  %s  | %s     %s    %s\n",
			param.TimeStamp.Format("01-02 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.ErrorMessage,
		)
	}))

	Router.Use(middleware.Recovery())       // 异常处理中间件
	Router.NoRoute(middleware.PageNotFound) // 404 页面处理

	Router.GET("/", site.Index)
	RouterGroupV1 = Router.Group("/v1")
	RouterGroupV2 = Router.Group("/v2")
}
