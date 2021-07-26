package controller

import (
	"347613781qq.com/demo1/model"
	"github.com/gin-gonic/gin"
)

func V1Test(c *gin.Context) {
	var nong model.Nong
	c.BindJSON(&nong)
	model.CreateNong(&nong)
	c.JSON(200, gin.H{
		"message": "err",
		"code":    "200",
	})
}
func V1Gorou(c *gin.Context) {
	var nong model.Nong
	c.BindJSON(&nong)
	model.CreateNong(&nong)
	c.JSON(200, gin.H{
		"message": "err",
		"code":    "200",
	})
}
