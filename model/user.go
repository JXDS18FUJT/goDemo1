package model

import (
	"time"

	"347613781qq.com/demo1/dao"
)

type User struct {
	Id        int       `gorm:"json:id" form:"id"`
	Name      string    `gorm:"json:name" form:"name"`
	Password  string    `gorm:"json:content" form:"password"`
	CreatedAt time.Time `gorm:"json:insert_time"`
	UpdatedAt time.Time `gorm:"json:update_time"`
}

func CreateUser(User *User) (err error) {
	// dao.DB.Model(&User).Update("CreatedAt", time.Now())
	err = dao.DB.Create(&User).Error
	return
}
func UpdateUser(User *User) (err error) {
	err = dao.DB.Save(&User).Error
	return
}
func GetUser(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&User{}).Error
	return err

}
func GetUserByName(name string) (user *User, err error) {
	user = new(User)
	if err = dao.DB.Debug().Where("name=?", name).First(user).Error; err != nil {
		return user, err
	}
	return

}
func GetAllUser() (UserList []*User, err error) {
	if err = dao.DB.Find(&UserList).Error; err != nil {
		return nil, err
	}
	return
}
