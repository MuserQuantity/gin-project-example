package request

// PageInfo 页面请求
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// IdReq 请求单id
type IdReq struct {
	Id int `json:"id" form:"id"` // 主键ID
}

func (r *IdReq) Uint() uint {
	return uint(r.Id)
}

// IdsReq 请求多id
type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId 请求角色id
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
