package main

import (
	"fmt"
	"time"

	"347613781qq.com/demo1/dao"
	"347613781qq.com/demo1/initialize"
	"347613781qq.com/demo1/model"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 初始化gin引擎
	var r = gin.New()
	logger, _ := zap.NewProduction()
	//添加一个ginzap中间件，它：
	//记录所有请求，如组合的访问和错误日志。
	//记录到stdout。
	//RFC3339，UTC时间格式。
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	//将所有死机记录到错误日志中
	//stack表示是否输出堆栈信息。
	r.Use(ginzap.RecoveryWithZap(logger, true))
	//跨域的相关配置
	// r.Use(middlewares.Cors())
	//读取相关的配置
	initialize.InitSetting()
	//挂载路由
	initialize.InitRouter(r)
	if err := initialize.InitSetting(); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	//初始化mysql
	err := dao.InitMySQL(initialize.Conf.MysqlConf.User, initialize.Conf.MysqlConf.Password, initialize.Conf.MysqlConf.Host, initialize.Conf.MysqlConf.Port, initialize.Conf.MysqlConf.Db)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	//取消表名字加s
	dao.DB.SingularTable(true)
	// 模型绑定
	dao.DB.AutoMigrate(&model.Nong{})

	r.Run(fmt.Sprintf(":%d", initialize.Conf.Port))
	//关闭mysql
	defer dao.Close() // 程序退出关闭数据库连接

}
