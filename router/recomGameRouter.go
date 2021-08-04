package router

import (
	"347613781qq.com/demo1/controller"
	"github.com/gin-gonic/gin"
)

func RecomGameRouter(Router *gin.RouterGroup) {
	BannerRouter := Router.Group("recomGame")
	{

		BannerRouter.POST("/add", controller.RecomGameAdd)
		BannerRouter.GET("/get", controller.RecomGameGet)
		BannerRouter.GET("/del", controller.RecomGameDel)
		BannerRouter.POST("/update", controller.RecomGameUpdate)

	}

}
