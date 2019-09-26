package site

import "github.com/gin-gonic/gin"

func Index(context *gin.Context) {
	context.JSON(200, gin.H{
		"computer": "MacBook Pro (13-inch, 2017, Two Thunderbolt 3 ports)",
		"cpu":      "2.3 GHz Intel Core i5",
		"memory":   "8 GB 2133 MHz LPDDR3",
		"graphics": "Intel Iris Plus Graphics 640 1536 MB",
	})
}
