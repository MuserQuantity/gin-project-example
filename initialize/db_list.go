package initialize

import (
	"github.com/MuserQuantity/gin-project-example/global"
	"gorm.io/gorm"
)

const sys = "system"

func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.SYS_CONFIG.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.Dbname] = GormMysqlByConfig(info)
		case "pgsql":
			dbMap[info.Dbname] = GormPgSqlByConfig(info)
		default:
			continue
		}
	}
	// 做特殊判断,是否有迁移
	// 适配低版本迁移多数据库版本
	if sysDB, ok := dbMap[sys]; ok {
		global.SYS_DB = sysDB
	}
	global.SYS_DBList = dbMap
}
