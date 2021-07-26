package controller

import (
	"347613781qq.com/demo1/common"
	"347613781qq.com/demo1/model"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var name = c.PostForm("name")

	user, err := model.GetUserByName(name)
	// err := model.CreateUser(&user)

	if err != nil {

		c.JSON(200, gin.H{
			"data": "密码或者用户名错误",
			"code": "400",
			"msg":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":  "存在该用户",
		"code": "200",
		"data": user,
	})
}

func LoginRegister(c *gin.Context) {
	var user = new(model.User)
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "绑定参数失败",
		})

		return

	}
	var token, _ = common.ReleaseToken(*user)
	c.JSON(200, gin.H{
		"msg":  "成功绑定",
		"data": token,
	})

}
func LoginParse(c *gin.Context) {

	var _, claims, _ = common.ParseToken(c.GetHeader("token"))
	c.JSON(200, gin.H{
		"msg":  "成功绑定",
		"data": claims,
	})

}
