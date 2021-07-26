package router

import (
	"347613781qq.com/demo1/controller"
	"github.com/gin-gonic/gin"
)

func WssRouter(Router *gin.RouterGroup) {
	wssGroup := Router.Group("wss")
	{

		wssGroup.GET("/chat", controller.WssChat)

	}

}
