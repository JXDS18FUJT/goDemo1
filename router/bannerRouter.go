package router

import (
	"347613781qq.com/demo1/controller"
	"github.com/gin-gonic/gin"
)

func BannerRouter(Router *gin.RouterGroup) {
	BannerRouter := Router.Group("banner")
	{

		BannerRouter.POST("/add", controller.BannerAdd)
		BannerRouter.GET("/get", controller.BannerGet)
		BannerRouter.GET("/del", controller.BannerDel)
		BannerRouter.POST("/update", controller.BannerUpdate)

	}

}
