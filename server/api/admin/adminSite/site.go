package adminSite

import "github.com/gogf/gf/v2/frame/g"

// SiteConfigReq 获取配置
type SiteConfigReq struct {
	g.Meta `path:"/site/config" method:"get" tags:"后台基础" summary:"获取配置"`
}

type SiteConfigRes struct {
	Version string `json:"version"        dc:"系统版本"`
	WsAddr  string `json:"wsAddr"         dc:"客户端websocket地址"`
	Domain  string `json:"domain"         dc:"对外域名"`
	Mode    string `json:"mode"           dc:"运行模式"`
}
