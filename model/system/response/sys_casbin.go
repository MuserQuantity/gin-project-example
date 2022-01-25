package response

import (
	"github.com/MuserQuantity/gin-project-example/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
