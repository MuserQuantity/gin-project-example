package system

import (
	"github.com/MuserQuantity/gin-project-example/model/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
