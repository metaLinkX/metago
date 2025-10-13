// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppMember is the golang structure of table app_member for DAO operations like Where/Data.
type AppMember struct {
	g.Meta             `orm:"table:app_member, do:true"`
	Id                 any         // 管理员ID
	RealName           any         // 真实姓名
	Username           any         // 帐号
	PasswordHash       any         // 密码
	Salt               any         // 密码盐
	PasswordResetToken any         // 密码重置令牌
	Avatar             any         // 头像
	Sex                any         // 性别
	Email              any         // 邮箱
	Mobile             any         // 手机号码
	Birthday           *gtime.Time // 生日
	Address            any         // 联系地址
	LastActiveAt       any         // 最后活跃时间
	Remark             any         // 备注
	Status             any         // 状态
	CreatedAt          any         // 创建时间
	UpdatedAt          any         // 修改时间
}
