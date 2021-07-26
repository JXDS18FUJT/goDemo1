package initialize

import (
	"347613781qq.com/demo1/router"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	//一级分组如api,bpi
	var PublicGroup = r.Group("/api")

	if Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	{
		router.V1Router(PublicGroup)
		router.LoginRouter(PublicGroup)
		router.NongRouter(PublicGroup)
		router.WssRouter(PublicGroup)

	}

}
