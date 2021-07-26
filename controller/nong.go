package controller

import (
	"347613781qq.com/demo1/model"
	"github.com/gin-gonic/gin"
)

func NongAdd(c *gin.Context) {
	var nong model.Nong
	c.ShouldBind(&nong)
	err := model.CreateNong(&nong)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "失败",
			"code":    "400",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "成功",
		"code":    "200",
	})
}

func NongUpdate(c *gin.Context) {
	var nong model.Nong
	c.ShouldBind(&nong)
	err := model.UpdateNong(&nong)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "失败",
			"code":    "400",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "成功",
		"code":    "200",
	})
}

func NongDel(c *gin.Context) {
	var nong model.Nong
	c.ShouldBind(&nong)
	err := model.DeleteNong(nong.Id)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "失败",
			"code":    "400",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "成功",
		"code":    "200",
	})
}
func NongGet(c *gin.Context) {
	var nong model.Nong
	c.ShouldBind(&nong)
	nong, err := model.GetNong(nong.Id)
	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
			"code":    "400",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "成功",
		"code":    "200",
		"data":    nong,
	})
}
