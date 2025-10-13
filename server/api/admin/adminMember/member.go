package adminMember

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model/input/adminin"
)

// InfoReq 获取登录用户信息
type InfoReq struct {
	g.Meta `path:"/member/info" method:"get" tags:"用户" summary:"获取登录用户信息"`
}

type InfoRes struct {
	*adminin.LoginMemberInfoModel
}
