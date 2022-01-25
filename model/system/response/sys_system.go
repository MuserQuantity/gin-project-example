package response

import "github.com/MuserQuantity/gin-project-example/model/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
