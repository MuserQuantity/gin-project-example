package system

import (
	"github.com/MuserQuantity/gin-project-example/global"
)

type JwtBlacklist struct {
	global.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
