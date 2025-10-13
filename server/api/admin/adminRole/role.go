package adminRole

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model/input/adminin"
)

// DynamicReq 动态路由
type DynamicReq struct {
	g.Meta `path:"/role/dynamic" method:"get" tags:"角色" summary:"获取动态路由" description:"获取登录用户动态路由"`
}

type DynamicRes struct {
	List []*adminin.MenuRoute `json:"list"   description:"数据列表"`
}
