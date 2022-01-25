package initialize

import (
	"fmt"

	"github.com/MuserQuantity/gin-project-example/global"
	"github.com/MuserQuantity/gin-project-example/model/config"
	"github.com/MuserQuantity/gin-project-example/utils"
)

func Timer() {
	if global.SYS_CONFIG.Timer.Start {
		for i := range global.SYS_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.SYS_Timer.AddTaskByFunc("ClearDB", global.SYS_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.SYS_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(global.SYS_CONFIG.Timer.Detail[i])
		}
	}
}
