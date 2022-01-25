package global

import (
	"github.com/go-redis/redis/v8"
	"sync"

	"github.com/MuserQuantity/gin-project-example/model/config"
	"github.com/MuserQuantity/gin-project-example/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	SYS_DB     *gorm.DB
	SYS_DBList map[string]*gorm.DB
	SYS_REDIS  *redis.Client
	SYS_CONFIG config.Server
	SYS_VP     *viper.Viper
	// SYS_LOG    *oplogging.Logger
	SYS_LOG                 *zap.Logger
	SYS_Timer               timer.Timer = timer.NewTimerTask()
	SYS_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return SYS_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := SYS_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
