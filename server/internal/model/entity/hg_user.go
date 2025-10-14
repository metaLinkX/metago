// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HgUser is the golang structure for table hg_user.
type HgUser struct {
	Id            int64       `json:"id"            orm:"id"            description:"用户ID"`
	Pid           int         `json:"pid"           orm:"pid"           description:"父级ID"`
	ComId         int         `json:"comId"         orm:"comId"         description:"小区ID"`
	HouseId       int         `json:"houseId"       orm:"houseId"       description:"房屋ID"`
	Username      string      `json:"username"      orm:"username"      description:"姓名"`
	Mobile        string      `json:"mobile"        orm:"mobile"        description:"手机号"`
	Password      string      `json:"password"      orm:"password"      description:"密码"`
	PasswordSalt  string      `json:"passwordSalt"  orm:"password_salt" description:"salt"`
	IdCard        string      `json:"idCard"        orm:"id_card"       description:"身份证"`
	IsAuthed      int         `json:"isAuthed"      orm:"isAuthed"      description:"是否已认证"`
	AuthedAt      *gtime.Time `json:"authedAt"      orm:"authedAt"      description:"最新认证时间"`
	TotalParkTime int         `json:"totalParkTime" orm:"totalParkTime" description:"累计租住时长"`
	UserType      int         `json:"userType"      orm:"user_type"     description:"租户类型"`
	UserStatus    int         `json:"userStatus"    orm:"user_status"   description:"用户状态"`
	StartAt       *gtime.Time `json:"startAt"       orm:"start_at"      description:"开始日期"`
	EndAt         *gtime.Time `json:"endAt"         orm:"end_at"        description:"结束日期"`
	Remark        string      `json:"remark"        orm:"remark"        description:"备注"`
	NowIsPay      int         `json:"nowIsPay"      orm:"nowIsPay"      description:"本期是否已支付"`
	IdCardFront   string      `json:"idCardFront"   orm:"idCardFront"   description:"身份证正面"`
	IdCardBack    string      `json:"idCardBack"    orm:"idCardBack"    description:"身份证背面"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"    description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"    description:"修改时间"`
	RefoundRs     string      `json:"refoundRs"     orm:"refoundRs"     description:"退租理由"`
	IsChangedPass int         `json:"isChangedPass" orm:"isChangedPass" description:"是否已更换密码"`
	OpenId        string      `json:"openId"        orm:"openId"        description:"openId"`
}
