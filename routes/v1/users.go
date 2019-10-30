package v1

import (
	"go-restful-api/modules/v1/controllers/users"
	"go-restful-api/routes"
)

func init() {
	v1 := routes.RouterGroupV1
	v1.GET("/users", users.List)
	v1.POST("/users", users.Add)
	v1.PUT("/users/:id", users.List)
	v1.DELETE("/users/:id", users.List)
}
