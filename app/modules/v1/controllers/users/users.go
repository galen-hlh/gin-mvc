package users

import (
	"github.com/gin-gonic/gin"
	"go-restful-api/app/modules/v1/rules/users"
	"go-restful-api/app/modules/v1/service"
	"net/http"
)

func List(c *gin.Context) {
	// 表单参数绑定
	var request users.ListRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "msg": err.Error()})
		return
	}

	// 成功
	c.JSON(http.StatusOK, gin.H{"code": "0", "msg": "success"})
}

func Add(c *gin.Context) {

	var request users.AddRequest
	_ = c.ShouldBind(&request)
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "msg": err.Error()})
		return
	}

	s := service.UsersService{}
	s.Add(request)
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
