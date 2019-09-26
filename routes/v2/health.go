package v2

import (
	"gin-mvc/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	v1 := routes.RouterGroupV2
	v1.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": true,
		})
	})
}
