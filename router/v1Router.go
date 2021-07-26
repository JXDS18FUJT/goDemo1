package router

import (
	"347613781qq.com/demo1/controller"
	"github.com/gin-gonic/gin"
)

func V1Router(Router *gin.RouterGroup) {
	v1Group := Router.Group("v1")
	{

		v1Group.POST("/test", controller.V1Test)

	}

}
