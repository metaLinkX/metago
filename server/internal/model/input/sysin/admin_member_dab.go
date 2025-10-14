// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 1.0.0
package sysin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"server/internal/model/entity"
	"server/internal/model/input/form"
)

// AdminMemberDabUpdateFields 修改管理员_用户表字段过滤
type AdminMemberDabUpdateFields struct {
	DeptId             int64       `json:"deptId"             dc:"部门ID"`
	RoleId             int64       `json:"roleId"             dc:"角色ID"`
	RealName           string      `json:"realName"           dc:"真实姓名"`
	Username           string      `json:"username"           dc:"帐号"`
	PasswordHash       string      `json:"passwordHash"       dc:"密码"`
	Salt               string      `json:"salt"               dc:"密码盐"`
	PasswordResetToken string      `json:"passwordResetToken" dc:"密码重置令牌"`
	Integral           float64     `json:"integral"           dc:"积分"`
	Balance            float64     `json:"balance"            dc:"余额"`
	Avatar             string      `json:"avatar"             dc:"头像"`
	Sex                int         `json:"sex"                dc:"性别"`
	Qq                 string      `json:"qq"                 dc:"qq"`
	Email              string      `json:"email"              dc:"邮箱"`
	Mobile             string      `json:"mobile"             dc:"手机号码"`
	Birthday           *gtime.Time `json:"birthday"           dc:"生日"`
	CityId             int64       `json:"cityId"             dc:"城市编码"`
	Address            string      `json:"address"            dc:"联系地址"`
	Pid                int64       `json:"pid"                dc:"上级管理员ID"`
	InviteCode         string      `json:"inviteCode"         dc:"邀请码"`
	Cash               string      `json:"cash"               dc:"提现配置"`
	LastActiveAt       *gtime.Time `json:"lastActiveAt"       dc:"最后活跃时间"`
	Remark             string      `json:"remark"             dc:"备注"`
}

// AdminMemberDabInsertFields 新增管理员_用户表字段过滤
type AdminMemberDabInsertFields struct {
	DeptId             int64       `json:"deptId"             dc:"部门ID"`
	RoleId             int64       `json:"roleId"             dc:"角色ID"`
	RealName           string      `json:"realName"           dc:"真实姓名"`
	Username           string      `json:"username"           dc:"帐号"`
	PasswordHash       string      `json:"passwordHash"       dc:"密码"`
	Salt               string      `json:"salt"               dc:"密码盐"`
	PasswordResetToken string      `json:"passwordResetToken" dc:"密码重置令牌"`
	Integral           float64     `json:"integral"           dc:"积分"`
	Balance            float64     `json:"balance"            dc:"余额"`
	Avatar             string      `json:"avatar"             dc:"头像"`
	Sex                int         `json:"sex"                dc:"性别"`
	Qq                 string      `json:"qq"                 dc:"qq"`
	Email              string      `json:"email"              dc:"邮箱"`
	Mobile             string      `json:"mobile"             dc:"手机号码"`
	Birthday           *gtime.Time `json:"birthday"           dc:"生日"`
	CityId             int64       `json:"cityId"             dc:"城市编码"`
	Address            string      `json:"address"            dc:"联系地址"`
	Pid                int64       `json:"pid"                dc:"上级管理员ID"`
	InviteCode         string      `json:"inviteCode"         dc:"邀请码"`
	Cash               string      `json:"cash"               dc:"提现配置"`
	LastActiveAt       *gtime.Time `json:"lastActiveAt"       dc:"最后活跃时间"`
	Remark             string      `json:"remark"             dc:"备注"`
}

// AdminMemberDabEditInp 修改/新增管理员_用户表
type AdminMemberDabEditInp struct {
	entity.AdminMember
}

func (in *AdminMemberDabEditInp) Filter(ctx context.Context) (err error) {
	// 验证帐号
	if err := g.Validator().Rules("required").Data(in.Username).Messages("帐号不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证密码
	if err := g.Validator().Rules("required").Data(in.PasswordHash).Messages("密码不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证密码盐
	if err := g.Validator().Rules("required").Data(in.Salt).Messages("密码盐不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证qq
	if err := g.Validator().Rules("qq").Data(in.Qq).Messages("qq不是QQ号码").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证邮箱
	if err := g.Validator().Rules("email").Data(in.Email).Messages("邮箱不是邮箱地址").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证手机号码
	if err := g.Validator().Rules("phone").Data(in.Mobile).Messages("手机号码不是手机号码").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证上级管理员ID
	if err := g.Validator().Rules("required").Data(in.Pid).Messages("上级管理员ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type AdminMemberDabEditModel struct{}

// AdminMemberDabDeleteInp 删除管理员_用户表
type AdminMemberDabDeleteInp struct {
	Id interface{} `json:"id" v:"required#管理员ID不能为空" dc:"管理员ID"`
}

func (in *AdminMemberDabDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type AdminMemberDabDeleteModel struct{}

// AdminMemberDabViewInp 获取指定管理员_用户表信息
type AdminMemberDabViewInp struct {
	Id int64 `json:"id" v:"required#管理员ID不能为空" dc:"管理员ID"`
}

func (in *AdminMemberDabViewInp) Filter(ctx context.Context) (err error) {
	return
}

type AdminMemberDabViewModel struct {
	entity.AdminMember
}

// AdminMemberDabListInp 获取管理员_用户表列表
type AdminMemberDabListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"管理员ID"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *AdminMemberDabListInp) Filter(ctx context.Context) (err error) {
	return
}

type AdminMemberDabListModel struct {
	Id                 int64       `json:"id"                 dc:"管理员ID"`
	DeptId             int64       `json:"deptId"             dc:"部门ID"`
	RoleId             int64       `json:"roleId"             dc:"角色ID"`
	RealName           string      `json:"realName"           dc:"真实姓名"`
	Username           string      `json:"username"           dc:"帐号"`
	PasswordHash       string      `json:"passwordHash"       dc:"密码"`
	Salt               string      `json:"salt"               dc:"密码盐"`
	PasswordResetToken string      `json:"passwordResetToken" dc:"密码重置令牌"`
	Integral           float64     `json:"integral"           dc:"积分"`
	Balance            float64     `json:"balance"            dc:"余额"`
	Avatar             string      `json:"avatar"             dc:"头像"`
	Sex                int         `json:"sex"                dc:"性别"`
	Qq                 string      `json:"qq"                 dc:"qq"`
	Email              string      `json:"email"              dc:"邮箱"`
	Mobile             string      `json:"mobile"             dc:"手机号码"`
	Birthday           *gtime.Time `json:"birthday"           dc:"生日"`
	CityId             int64       `json:"cityId"             dc:"城市编码"`
	Address            string      `json:"address"            dc:"联系地址"`
	Pid                int64       `json:"pid"                dc:"上级管理员ID"`
	InviteCode         string      `json:"inviteCode"         dc:"邀请码"`
	Cash               string      `json:"cash"               dc:"提现配置"`
	LastActiveAt       *gtime.Time `json:"lastActiveAt"       dc:"最后活跃时间"`
	Remark             string      `json:"remark"             dc:"备注"`
	Status             int         `json:"status"             dc:"状态"`
	CreatedAt          *gtime.Time `json:"createdAt"          dc:"创建时间"`
	UpdatedAt          *gtime.Time `json:"updatedAt"          dc:"修改时间"`
}

// AdminMemberDabExportModel 导出管理员_用户表
type AdminMemberDabExportModel struct {
	Id                 int64       `json:"id"                 dc:"管理员ID"`
	DeptId             int64       `json:"deptId"             dc:"部门ID"`
	RoleId             int64       `json:"roleId"             dc:"角色ID"`
	RealName           string      `json:"realName"           dc:"真实姓名"`
	Username           string      `json:"username"           dc:"帐号"`
	PasswordHash       string      `json:"passwordHash"       dc:"密码"`
	Salt               string      `json:"salt"               dc:"密码盐"`
	PasswordResetToken string      `json:"passwordResetToken" dc:"密码重置令牌"`
	Integral           float64     `json:"integral"           dc:"积分"`
	Balance            float64     `json:"balance"            dc:"余额"`
	Avatar             string      `json:"avatar"             dc:"头像"`
	Sex                int         `json:"sex"                dc:"性别"`
	Qq                 string      `json:"qq"                 dc:"qq"`
	Email              string      `json:"email"              dc:"邮箱"`
	Mobile             string      `json:"mobile"             dc:"手机号码"`
	Birthday           *gtime.Time `json:"birthday"           dc:"生日"`
	CityId             int64       `json:"cityId"             dc:"城市编码"`
	Address            string      `json:"address"            dc:"联系地址"`
	Pid                int64       `json:"pid"                dc:"上级管理员ID"`
	InviteCode         string      `json:"inviteCode"         dc:"邀请码"`
	Cash               string      `json:"cash"               dc:"提现配置"`
	LastActiveAt       *gtime.Time `json:"lastActiveAt"       dc:"最后活跃时间"`
	Remark             string      `json:"remark"             dc:"备注"`
	Status             int         `json:"status"             dc:"状态"`
	CreatedAt          *gtime.Time `json:"createdAt"          dc:"创建时间"`
	UpdatedAt          *gtime.Time `json:"updatedAt"          dc:"修改时间"`
}

// AdminMemberDabStatusInp 更新管理员_用户表状态
type AdminMemberDabStatusInp struct {
	Id     int64 `json:"id" v:"required#管理员ID不能为空" dc:"管理员ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *AdminMemberDabStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("管理员ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	//if !validate.InSlice(consts.StatusSlice, in.Status) {
	//	err = gerror.New("状态不正确")
	//	return
	//}
	return
}

type AdminMemberDabStatusModel struct{}
