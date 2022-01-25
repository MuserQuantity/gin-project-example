package request

import (
	"github.com/MuserQuantity/gin-project-example/model/common/request"
	"github.com/MuserQuantity/gin-project-example/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
