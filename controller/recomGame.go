package controller

import (
	"fmt"

	"347613781qq.com/demo1/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

func RecomGameAdd(c *gin.Context) {
	var RecomGame model.RecomGame
	c.ShouldBind(&RecomGame)
	err := model.CreateRecomGame(&RecomGame)
	var validate = validator.New()
	err1 := validate.Struct(&RecomGame)
	if err1 != nil {
		c.JSON(200, gin.H{
			"message": err1.Error(),
			"code":    "400",
		})
		return

	}

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

func RecomGameUpdate(c *gin.Context) {
	var RecomGame model.RecomGame
	c.ShouldBind(&RecomGame)
	if RecomGame.RecomGameId <= 0 {
		c.JSON(200, gin.H{
			"message": "缺少部分参数RecomGame_id",
			"code":    "400",
		})
		return

	}
	err := model.UpdateRecomGame(&RecomGame, RecomGame.RecomGameId)
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

func RecomGameDel(c *gin.Context) {
	var RecomGame model.RecomGame
	c.ShouldBind(&RecomGame)
	err := model.DeleteRecomGame(RecomGame.RecomGameId)
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
func RecomGameGet(c *gin.Context) {
	var emptyData = make([]string, 0)
	var resData = make([]model.RecomGame, 1, 30)
	var RecomGame model.RecomGame
	c.ShouldBind(&RecomGame)
	RecomGame, err := model.GetRecomGame(RecomGame.RecomGameId)
	fmt.Print(err)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(200, gin.H{
				"message": "成功",
				"code":    "200",
				"data":    emptyData,
			})
			return

		}
		c.JSON(200, gin.H{
			"message": "错误",
			"code":    "400",
		})

		return
	}
	resData[0] = RecomGame
	c.JSON(200, gin.H{
		"message": "成功",
		"code":    "200",
		"data":    resData,
	})
}
