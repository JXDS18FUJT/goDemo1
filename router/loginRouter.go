package router

import (
	"347613781qq.com/demo1/controller"
	"github.com/gin-gonic/gin"
)

func LoginRouter(Router *gin.RouterGroup) {
	LoginRouter := Router.Group("login")
	{

		LoginRouter.POST("/", controller.Login)
		LoginRouter.POST("/register", controller.LoginRegister)
		LoginRouter.POST("/parse", controller.LoginParse)

	}

}
