package model

import (
	"time"

	"347613781qq.com/demo1/dao"
)

type Nong struct {
	Id         int       `gorm:"json:id"`
	Name       string    `gorm:"json:name"`
	Content    string    `gorm:"json:content"`
	InsertTime time.Time `gorm:"json:insert_time"`
	UpdateTime time.Time `gorm:"json:update_time"`
}

func CreateNong(nong *Nong) (err error) {
	err = dao.DB.Create(&nong).Error
	return
}
func UpdateNong(nong *Nong) (err error) {
	err = dao.DB.Save(&nong).Error
	return
}
func DeleteNong(id int) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Nong{}).Error
	return
}
func GetNong(id int) (nong Nong, err error) {

	err = dao.DB.Where("id = ?", id).First(&nong).Error
	return nong, err

}
func GetAllNong() (nongList []*Nong, err error) {
	if err = dao.DB.Find(&nongList).Error; err != nil {
		return nil, err
	}
	return
}
