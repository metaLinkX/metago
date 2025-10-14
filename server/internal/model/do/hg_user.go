// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HgUser is the golang structure of table hg_user for DAO operations like Where/Data.
type HgUser struct {
	g.Meta        `orm:"table:hg_user, do:true"`
	Id            interface{} // 用户ID
	Pid           interface{} // 父级ID
	ComId         interface{} // 小区ID
	HouseId       interface{} // 房屋ID
	Username      interface{} // 姓名
	Mobile        interface{} // 手机号
	Password      interface{} // 密码
	PasswordSalt  interface{} // salt
	IdCard        interface{} // 身份证
	IsAuthed      interface{} // 是否已认证
	AuthedAt      *gtime.Time // 最新认证时间
	TotalParkTime interface{} // 累计租住时长
	UserType      interface{} // 租户类型
	UserStatus    interface{} // 用户状态
	StartAt       *gtime.Time // 开始日期
	EndAt         *gtime.Time // 结束日期
	Remark        interface{} // 备注
	NowIsPay      interface{} // 本期是否已支付
	IdCardFront   interface{} // 身份证正面
	IdCardBack    interface{} // 身份证背面
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 修改时间
	RefoundRs     interface{} // 退租理由
	IsChangedPass interface{} // 是否已更换密码
	OpenId        interface{} // openId
}
