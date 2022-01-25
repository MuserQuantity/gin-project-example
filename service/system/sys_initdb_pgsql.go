package system

import (
	"github.com/MuserQuantity/gin-project-example/global"
	"github.com/MuserQuantity/gin-project-example/initialize/source/system"
	"github.com/MuserQuantity/gin-project-example/model/config"
	model "github.com/MuserQuantity/gin-project-example/model/system"
	"github.com/MuserQuantity/gin-project-example/model/system/request"
	"github.com/MuserQuantity/gin-project-example/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// writePgsqlConfig pgsql 回写配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (initDBService *InitDBService) writePgsqlConfig(pgsql config.Pgsql) error {
	global.SYS_CONFIG.System.DbType = "pgsql"
	global.SYS_CONFIG.Pgsql = pgsql
	cs := utils.StructToMap(global.SYS_CONFIG)
	for k, v := range cs {
		global.SYS_VP.Set(k, v)
	}
	global.SYS_VP.Set("jwt.signing-key", uuid.NewV4().String())
	return global.SYS_VP.WriteConfig()
}

func (initDBService *InitDBService) initPgsqlDB(conf request.InitDB) error {
	dsn := conf.PgsqlEmptyDsn()
	createSql := "CREATE DATABASE " + conf.DBName
	if err := initDBService.createDatabase(dsn, "pgx", createSql); err != nil {
		return err
	} // 创建数据库

	pgsqlConfig := conf.ToPgsqlConfig()
	if pgsqlConfig.Dbname == "" {
		return nil
	} // 如果没有数据库名, 则跳出初始化数据

	if db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  pgsqlConfig.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		global.SYS_DB = db
	}

	if err := initDBService.initTables(); err != nil {
		global.SYS_DB = nil
		return err
	}

	if err := initDBService.initPgsqlData(); err != nil {
		global.SYS_DB = nil
		return err
	}

	if err := initDBService.writePgsqlConfig(pgsqlConfig); err != nil {
		return err
	}

	return nil
}

// initPgsqlData pgsql 初始化数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (initDBService *InitDBService) initPgsqlData() error {
	return model.PgsqlDataInitialize(
		system.Api,
		system.User,
		system.Casbin,
		system.BaseMenu,
		system.Authority,
		system.Dictionary,
		system.UserAuthority,
		system.DataAuthorities,
		system.AuthoritiesMenus,
		system.DictionaryDetail,
		system.ViewAuthorityMenuPostgres,
	)
}
