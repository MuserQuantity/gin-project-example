package request

import (
	"github.com/MuserQuantity/gin-project-example/model/autocode"
	"github.com/MuserQuantity/gin-project-example/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}