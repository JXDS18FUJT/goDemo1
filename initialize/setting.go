package initialize

import (
	"reflect"

	"gopkg.in/ini.v1"
)

type AppConf struct {
	Port       int  `ini:"port"`
	Release    bool `ini:"release"`
	*MysqlConf `ini:"mysql"`
}
type MysqlConf struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Db       string `ini:"db"`
}

var Conf = new(AppConf)

func InitSetting() error {
	return ini.MapTo(Conf, "conf/conf.ini")
}
func InitSettingMapTo(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data

}
