package model

import (
	"347613781qq.com/demo1/common/diymodel"
	"347613781qq.com/demo1/dao"
)

//设置主键
type RecomGame struct {
	RecomGameId int    `gorm:"json:recom_game_id;primary_key" form:"recomGame_id" json:"recom_game_id"`
	Title       string `gorm:"json:title" form:"title" json:"title" validate:"required"`
	SubTitle    string `gorm:"json:subTitle" form:"subTitle" json:"subTitle" validate:"required"`

	Url         string         `gorm:"json:url" form:"url" json:"url" validate:"required"`
	UpdateTime  diymodel.XTime `json:"update_time" gorm:"default:null"`
	Insert_Time diymodel.XTime `json:"insert_time" gorm:"default:null"`
}

func CreateRecomGame(recomGame *RecomGame) (err error) {
	err = dao.DB.Create(&recomGame).Error
	return
}
func UpdateRecomGame(recomGame *RecomGame, id int) (err error) {
	// var newrecomGame = new(recomGame)
	err = dao.DB.Where("recomGame_id=?", id).Updates(&recomGame).Error
	return
}
func DeleteRecomGame(id int) (err error) {
	err = dao.DB.Where("recomGame_id=?", id).Delete(&RecomGame{}).Error
	return
}
func GetRecomGame(id int) (recomGame RecomGame, err error) {

	err = dao.DB.Debug().Where("recomGame_id=?", id).First(&recomGame).Error
	return recomGame, err
}
func GetAllRecomGame() (recomGameList []*RecomGame, err error) {
	if err = dao.DB.Find(&recomGameList).Error; err != nil {
		return nil, err
	}
	return
}
