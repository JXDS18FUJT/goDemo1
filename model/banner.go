package model

import (
	"347613781qq.com/demo1/common/diymodel"
	"347613781qq.com/demo1/dao"
)

//设置主键
type Banner struct {
	BannerId    int            `gorm:"json:banner_id;primary_key" form:"banner_id" json:"banner_id"`
	Name        string         `gorm:"json:name" form:"name" json:"name" validate:"required"`
	Desc        string         `gorm:"json:desc" form:"desc" json:"desc" name:"描述"`
	Url         string         `gorm:"json:url" form:"url" json:"url" validate:"required,email"`
	UpdateTime  diymodel.XTime `json:"update_time" gorm:"default:null"`
	Insert_Time diymodel.XTime `json:"insert_time" gorm:"default:null"`
}

func CreateBanner(banner *Banner) (err error) {
	err = dao.DB.Create(&banner).Error
	return
}
func UpdateBanner(banner *Banner, id int) (err error) {
	// var newBanner = new(Banner)
	err = dao.DB.Where("banner_id=?", id).Updates(&banner).Error
	return
}
func DeleteBanner(id int) (err error) {
	err = dao.DB.Where("banner_id=?", id).Delete(&Banner{}).Error
	return
}
func GetBanner(id int) (banner Banner, err error) {

	err = dao.DB.Debug().Where("banner_id=?", id).First(&banner).Error
	return banner, err
}
func GetAllBanner() (bannerList []*Banner, err error) {
	if err = dao.DB.Find(&bannerList).Error; err != nil {
		return nil, err
	}
	return
}
