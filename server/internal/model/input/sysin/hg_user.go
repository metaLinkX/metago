// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 1.0.0
package sysin

import (
	"context"
	"server/internal/model/entity"
	"server/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HgUserUpdateFields 修改租户字段过滤
type HgUserUpdateFields struct {
	Pid           int         `json:"pid"           dc:"父级ID"`
	ComId         int         `json:"comId"         dc:"小区ID"`
	HouseId       int         `json:"houseId"       dc:"房屋ID"`
	Username      string      `json:"username"      dc:"姓名"`
	Mobile        string      `json:"mobile"        dc:"手机号"`
	Password      string      `json:"password"      dc:"密码"`
	PasswordSalt  string      `json:"passwordSalt"  dc:"salt"`
	IdCard        string      `json:"idCard"        dc:"身份证"`
	IsAuthed      int         `json:"isAuthed"      dc:"是否已认证"`
	AuthedAt      *gtime.Time `json:"authedAt"      dc:"最新认证时间"`
	TotalParkTime int         `json:"totalParkTime" dc:"累计租住时长"`
	UserType      int         `json:"userType"      dc:"租户类型"`
	UserStatus    int         `json:"userStatus"    dc:"用户状态"`
	StartAt       *gtime.Time `json:"startAt"       dc:"开始日期"`
	EndAt         *gtime.Time `json:"endAt"         dc:"结束日期"`
	Remark        string      `json:"remark"        dc:"备注"`
	NowIsPay      int         `json:"nowIsPay"      dc:"本期是否已支付"`
	IdCardFront   string      `json:"idCardFront"   dc:"身份证正面"`
	IdCardBack    string      `json:"idCardBack"    dc:"身份证背面"`
	RefoundRs     string      `json:"refoundRs"     dc:"退租理由"`
	IsChangedPass int         `json:"isChangedPass" dc:"是否已更换密码"`
	OpenId        string      `json:"openId"        dc:"openId"`
}

// HgUserInsertFields 新增租户字段过滤
type HgUserInsertFields struct {
	Pid           int         `json:"pid"           dc:"父级ID"`
	ComId         int         `json:"comId"         dc:"小区ID"`
	HouseId       int         `json:"houseId"       dc:"房屋ID"`
	Username      string      `json:"username"      dc:"姓名"`
	Mobile        string      `json:"mobile"        dc:"手机号"`
	Password      string      `json:"password"      dc:"密码"`
	PasswordSalt  string      `json:"passwordSalt"  dc:"salt"`
	IdCard        string      `json:"idCard"        dc:"身份证"`
	IsAuthed      int         `json:"isAuthed"      dc:"是否已认证"`
	AuthedAt      *gtime.Time `json:"authedAt"      dc:"最新认证时间"`
	TotalParkTime int         `json:"totalParkTime" dc:"累计租住时长"`
	UserType      int         `json:"userType"      dc:"租户类型"`
	UserStatus    int         `json:"userStatus"    dc:"用户状态"`
	StartAt       *gtime.Time `json:"startAt"       dc:"开始日期"`
	EndAt         *gtime.Time `json:"endAt"         dc:"结束日期"`
	Remark        string      `json:"remark"        dc:"备注"`
	NowIsPay      int         `json:"nowIsPay"      dc:"本期是否已支付"`
	IdCardFront   string      `json:"idCardFront"   dc:"身份证正面"`
	IdCardBack    string      `json:"idCardBack"    dc:"身份证背面"`
	RefoundRs     string      `json:"refoundRs"     dc:"退租理由"`
	IsChangedPass int         `json:"isChangedPass" dc:"是否已更换密码"`
	OpenId        string      `json:"openId"        dc:"openId"`
}

// HgUserEditInp 修改/新增租户
type HgUserEditInp struct {
	entity.HgUser
}

func (in *HgUserEditInp) Filter(ctx context.Context) (err error) {
	// 验证小区ID
	if err := g.Validator().Rules("required").Data(in.ComId).Messages("小区ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证房屋ID
	if err := g.Validator().Rules("required").Data(in.HouseId).Messages("房屋ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证姓名
	if err := g.Validator().Rules("required").Data(in.Username).Messages("姓名不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证手机号
	if err := g.Validator().Rules("phone").Data(in.Mobile).Messages("手机号不是手机号码").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证密码
	if err := g.Validator().Rules("regex:^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,18}$").Data(in.Password).Messages("密码必须包含6-18为字母和数字").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证salt
	if err := g.Validator().Rules("required").Data(in.PasswordSalt).Messages("salt不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证身份证
	if err := g.Validator().Rules("resident-id").Data(in.IdCard).Messages("身份证不是身份证号码").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证是否已认证
	if err := g.Validator().Rules("required").Data(in.IsAuthed).Messages("是否已认证不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证累计租住时长
	if err := g.Validator().Rules("required").Data(in.TotalParkTime).Messages("累计租住时长不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证租户类型
	if err := g.Validator().Rules("required").Data(in.UserType).Messages("租户类型不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证用户状态
	if err := g.Validator().Rules("required").Data(in.UserStatus).Messages("用户状态不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.UserStatus).Messages("用户状态值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证本期是否已支付
	if err := g.Validator().Rules("required").Data(in.NowIsPay).Messages("本期是否已支付不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证身份证正面
	if err := g.Validator().Rules("required").Data(in.IdCardFront).Messages("身份证正面不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证身份证背面
	if err := g.Validator().Rules("required").Data(in.IdCardBack).Messages("身份证背面不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证是否已更换密码
	if err := g.Validator().Rules("required").Data(in.IsChangedPass).Messages("是否已更换密码不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.IsChangedPass).Messages("是否已更换密码值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type HgUserEditModel struct{}

// HgUserDeleteInp 删除租户
type HgUserDeleteInp struct {
	Id interface{} `json:"id" v:"required#用户ID不能为空" dc:"用户ID"`
}

func (in *HgUserDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type HgUserDeleteModel struct{}

// HgUserViewInp 获取指定租户信息
type HgUserViewInp struct {
	Id int64 `json:"id" v:"required#用户ID不能为空" dc:"用户ID"`
}

func (in *HgUserViewInp) Filter(ctx context.Context) (err error) {
	return
}

type HgUserViewModel struct {
	entity.HgUser
}

// HgUserListInp 获取租户列表
type HgUserListInp struct {
	form.PageReq
	Id         int64         `json:"id"         dc:"用户ID"`
	UserStatus int           `json:"userStatus" dc:"用户状态"`
	CreatedAt  []*gtime.Time `json:"createdAt"  dc:"创建时间"`
}

func (in *HgUserListInp) Filter(ctx context.Context) (err error) {
	return
}

type HgUserListModel struct {
	Id            int64       `json:"id"            dc:"用户ID"`
	Pid           int         `json:"pid"           dc:"父级ID"`
	ComId         int         `json:"comId"         dc:"小区ID"`
	HouseId       int         `json:"houseId"       dc:"房屋ID"`
	Username      string      `json:"username"      dc:"姓名"`
	Mobile        string      `json:"mobile"        dc:"手机号"`
	Password      string      `json:"password"      dc:"密码"`
	PasswordSalt  string      `json:"passwordSalt"  dc:"salt"`
	IdCard        string      `json:"idCard"        dc:"身份证"`
	IsAuthed      int         `json:"isAuthed"      dc:"是否已认证"`
	AuthedAt      *gtime.Time `json:"authedAt"      dc:"最新认证时间"`
	TotalParkTime int         `json:"totalParkTime" dc:"累计租住时长"`
	UserType      int         `json:"userType"      dc:"租户类型"`
	UserStatus    int         `json:"userStatus"    dc:"用户状态"`
	StartAt       *gtime.Time `json:"startAt"       dc:"开始日期"`
	EndAt         *gtime.Time `json:"endAt"         dc:"结束日期"`
	Remark        string      `json:"remark"        dc:"备注"`
	NowIsPay      int         `json:"nowIsPay"      dc:"本期是否已支付"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"修改时间"`
	RefoundRs     string      `json:"refoundRs"     dc:"退租理由"`
	IsChangedPass int         `json:"isChangedPass" dc:"是否已更换密码"`
	OpenId        string      `json:"openId"        dc:"openId"`
}

// HgUserExportModel 导出租户
type HgUserExportModel struct {
	Id            int64       `json:"id"            dc:"用户ID"`
	Pid           int         `json:"pid"           dc:"父级ID"`
	ComId         int         `json:"comId"         dc:"小区ID"`
	HouseId       int         `json:"houseId"       dc:"房屋ID"`
	Username      string      `json:"username"      dc:"姓名"`
	Mobile        string      `json:"mobile"        dc:"手机号"`
	Password      string      `json:"password"      dc:"密码"`
	PasswordSalt  string      `json:"passwordSalt"  dc:"salt"`
	IdCard        string      `json:"idCard"        dc:"身份证"`
	IsAuthed      int         `json:"isAuthed"      dc:"是否已认证"`
	AuthedAt      *gtime.Time `json:"authedAt"      dc:"最新认证时间"`
	TotalParkTime int         `json:"totalParkTime" dc:"累计租住时长"`
	UserType      int         `json:"userType"      dc:"租户类型"`
	UserStatus    int         `json:"userStatus"    dc:"用户状态"`
	StartAt       *gtime.Time `json:"startAt"       dc:"开始日期"`
	EndAt         *gtime.Time `json:"endAt"         dc:"结束日期"`
	Remark        string      `json:"remark"        dc:"备注"`
	NowIsPay      int         `json:"nowIsPay"      dc:"本期是否已支付"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"修改时间"`
	RefoundRs     string      `json:"refoundRs"     dc:"退租理由"`
	IsChangedPass int         `json:"isChangedPass" dc:"是否已更换密码"`
	OpenId        string      `json:"openId"        dc:"openId"`
}