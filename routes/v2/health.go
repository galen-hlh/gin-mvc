package v2

import (
	"github.com/gin-gonic/gin"
	"go-restful-api/routes"
)

func init() {
	v1 := routes.RouterGroupV2
	v1.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": true,
		})
	})
}
