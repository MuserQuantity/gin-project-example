package v1

import (
	"github.com/MuserQuantity/gin-project-example/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
