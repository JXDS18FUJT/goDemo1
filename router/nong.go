package router

import (
	"347613781qq.com/demo1/controller"
	"github.com/gin-gonic/gin"
)

func NongRouter(Router *gin.RouterGroup) {
	NongGroup := Router.Group("nong")
	{

		NongGroup.POST("/add", controller.NongAdd)
		NongGroup.POST("/update", controller.NongUpdate)
		NongGroup.POST("/del", controller.NongDel)
		NongGroup.POST("/get", controller.NongGet)

	}

}
