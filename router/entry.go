package router

import (
	"github.com/MuserQuantity/gin-project-example/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
