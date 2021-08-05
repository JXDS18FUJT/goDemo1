package controller

import (
	"fmt"

	"347613781qq.com/demo1/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func BannerAdd(c *gin.Context) {
	var Banner model.Banner
	c.ShouldBind(&Banner)
	err := model.CreateBanner(&Banner)

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

func BannerUpdate(c *gin.Context) {
	var Banner model.Banner
	c.ShouldBind(&Banner)
	if Banner.BannerId <= 0 {
		c.JSON(200, gin.H{
			"message": "缺少部分参数banner_id",
			"code":    "400",
		})
		return

	}
	err := model.UpdateBanner(&Banner, Banner.BannerId)
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

func BannerDel(c *gin.Context) {
	var Banner model.Banner
	c.ShouldBind(&Banner)
	err := model.DeleteBanner(Banner.BannerId)
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
func BannerGet(c *gin.Context) {
	var emptyData = make([]string, 0)
	var resData = make([]model.Banner, 1, 30)
	var Banner model.Banner
	c.ShouldBind(&Banner)
	Banner, err := model.GetBanner(Banner.BannerId)
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
	resData[0] = Banner
	c.JSON(200, gin.H{
		"message": "成功",
		"code":    "200",
		"data":    resData,
	})
}
