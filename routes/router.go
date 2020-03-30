package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-restful-api/app/middleware"
	"go-restful-api/app/modules/v1/controllers/site"
	_ "go-restful-api/app/validate"
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

	Router.Use(middleware.HandleErrors())       // 异常处理中间件
	Router.NoRoute(middleware.HandleNotFound)   // 404 页面处理
	Router.NoMethod(middleware.HandleNotMethod) // 405 页面处理

	Router.GET("/", site.Index)
	RouterGroupV1 = Router.Group("/v1")
	RouterGroupV2 = Router.Group("/v2")
}
