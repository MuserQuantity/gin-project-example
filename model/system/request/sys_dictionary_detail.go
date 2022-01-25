package request

import (
	"github.com/MuserQuantity/gin-project-example/model/common/request"
	"github.com/MuserQuantity/gin-project-example/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
