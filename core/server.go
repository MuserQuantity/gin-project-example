package core

import (
	"fmt"
	"time"

	"github.com/MuserQuantity/gin-project-example/global"
	"github.com/MuserQuantity/gin-project-example/initialize"
	"github.com/MuserQuantity/gin-project-example/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.SYS_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.SYS_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.SYS_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.SYS_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Println("正在运行")
	global.SYS_LOG.Error(s.ListenAndServe().Error())
}
