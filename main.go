package main

import (
	"fmt"
	"github.com/MuserQuantity/gin-project-example/core"
	"github.com/MuserQuantity/gin-project-example/global"
	"github.com/MuserQuantity/gin-project-example/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.SYS_VP = core.Viper()      // 初始化Viper
	global.SYS_LOG = core.Zap()       // 初始化zap日志库
	global.SYS_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.SYS_DB != nil {
		initialize.RegisterTables(global.SYS_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.SYS_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
	fmt.Println("正在运行")
}
