package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

type MysqlConf struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Db       string `ini:"db"`
}

func InitMySQL(user string, password string, host string, port int, db string) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, db)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {

		return
	}
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
