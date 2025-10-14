// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/internal/model"
	"server/internal/model/entity"
	"server/internal/model/input/sysin"
	"server/library/hgorm/handler"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysAdminMemberDab interface {
		// Model 管理员_用户表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取管理员_用户表列表
		List(ctx context.Context, in *sysin.AdminMemberDabListInp) (list []*sysin.AdminMemberDabListModel, totalCount int, err error)
		// Export 导出管理员_用户表
		Export(ctx context.Context, in *sysin.AdminMemberDabListInp) (err error)
		// Edit 修改/新增管理员_用户表
		Edit(ctx context.Context, in *sysin.AdminMemberDabEditInp) (err error)
		// Delete 删除管理员_用户表
		Delete(ctx context.Context, in *sysin.AdminMemberDabDeleteInp) (err error)
		// View 获取管理员_用户表指定信息
		View(ctx context.Context, in *sysin.AdminMemberDabViewInp) (res *sysin.AdminMemberDabViewModel, err error)
		// Status 更新管理员_用户表状态
		Status(ctx context.Context, in *sysin.AdminMemberDabStatusInp) (err error)
	}
	ISysConfig interface {
		// GetLogin 获取登录配置
		GetLogin(ctx context.Context) (conf *model.LoginConfig, err error)
		// GetWechat 获取微信配置
		GetWechat(ctx context.Context) (conf *model.WechatConfig, err error)
		// GetPay 获取支付配置
		GetPay(ctx context.Context) (conf *model.PayConfig, err error)
		// GetSms 获取短信配置
		GetSms(ctx context.Context) (conf *model.SmsConfig, err error)
		// GetUpload 获取上传配置
		GetUpload(ctx context.Context) (conf *model.UploadConfig, err error)
		// GetSmtp 获取邮件配置
		GetSmtp(ctx context.Context) (conf *model.EmailConfig, err error)
		// GetBasic 获取基础配置
		GetBasic(ctx context.Context) (conf *model.BasicConfig, err error)
		// GetApiJwtToken 获取本地token配置
		GetApiJwtToken(ctx context.Context) (conf *model.TokenConfig, err error)
		// GetAdminJwtToken 获取本地token配置
		GetAdminJwtToken(ctx context.Context) (conf *model.TokenConfig, err error)
		// GetConfigByGroup 获取指定分组的配置
		GetConfigByGroup(ctx context.Context, in *sysin.GetConfigInp) (res *sysin.GetConfigModel, err error)
		// ConversionType 转换类型
		ConversionType(ctx context.Context, models *entity.SysConfig) (value interface{}, err error)
		// GetLoadGenerate 获取本地生成配置
		GetLoadGenerate(ctx context.Context) (conf *model.GenerateConfig, err error)
	}
	ISysDictType interface {
		// Tree 树
		Tree(ctx context.Context) (list []*sysin.DictTypeTree, err error)
		// Delete 删除
		Delete(ctx context.Context, in *sysin.DictTypeDeleteInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.DictTypeEditInp) (err error)
		// TreeSelect 获取类型关系树选项
		TreeSelect(ctx context.Context, in *sysin.DictTreeSelectInp) (list []*sysin.DictTypeTree, err error)
		// BuiltinSelect 内置字典选项
		BuiltinSelect() (list []*sysin.DictTypeTree)
	}
	ISysGenCodes interface {
		// Delete 删除
		Delete(ctx context.Context, in *sysin.GenCodesDeleteInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.GenCodesEditInp) (res *sysin.GenCodesEditModel, err error)
		// Status 更新部门状态
		Status(ctx context.Context, in *sysin.GenCodesStatusInp) (err error)
		// MaxSort 最大排序
		MaxSort(ctx context.Context, in *sysin.GenCodesMaxSortInp) (res *sysin.GenCodesMaxSortModel, err error)
		// View 获取指定字典类型信息
		View(ctx context.Context, in *sysin.GenCodesViewInp) (res *sysin.GenCodesViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *sysin.GenCodesListInp) (list []*sysin.GenCodesListModel, totalCount int, err error)
		// Selects 选项
		Selects(ctx context.Context, in *sysin.GenCodesSelectsInp) (res *sysin.GenCodesSelectsModel, err error)
		// TableSelect 表选项
		TableSelect(ctx context.Context, in *sysin.GenCodesTableSelectInp) (res []*sysin.GenCodesTableSelectModel, err error)
		// ColumnSelect 表字段选项
		ColumnSelect(ctx context.Context, in *sysin.GenCodesColumnSelectInp) (res []*sysin.GenCodesColumnSelectModel, err error)
		// ColumnList 表字段列表
		ColumnList(ctx context.Context, in *sysin.GenCodesColumnListInp) (res []*sysin.GenCodesColumnListModel, err error)
		// Preview 生成预览
		Preview(ctx context.Context, in *sysin.GenCodesPreviewInp) (res *sysin.GenCodesPreviewModel, err error)
		// Build 提交生成
		Build(ctx context.Context, in *sysin.GenCodesBuildInp) (err error)
	}
	ISysHgUser interface {
		// Model 租户ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取租户列表
		List(ctx context.Context, in *sysin.HgUserListInp) (list []*sysin.HgUserListModel, totalCount int, err error)
		// Export 导出租户
		Export(ctx context.Context, in *sysin.HgUserListInp) (err error)
		// Edit 修改/新增租户
		Edit(ctx context.Context, in *sysin.HgUserEditInp) (err error)
		// Delete 删除租户
		Delete(ctx context.Context, in *sysin.HgUserDeleteInp) (err error)
		// View 获取租户指定信息
		View(ctx context.Context, in *sysin.HgUserViewInp) (res *sysin.HgUserViewModel, err error)
	}
	ISysLogger interface {
		// PushAdminLoginLog 推送adminLog
		PushAdminLoginLog(ctx context.Context)
		PushApiLoginLog(ctx context.Context)
	}
)

var (
	localSysAdminMemberDab ISysAdminMemberDab
	localSysConfig         ISysConfig
	localSysDictType       ISysDictType
	localSysGenCodes       ISysGenCodes
	localSysHgUser         ISysHgUser
	localSysLogger         ISysLogger
)

func SysAdminMemberDab() ISysAdminMemberDab {
	if localSysAdminMemberDab == nil {
		panic("implement not found for interface ISysAdminMemberDab, forgot register?")
	}
	return localSysAdminMemberDab
}

func RegisterSysAdminMemberDab(i ISysAdminMemberDab) {
	localSysAdminMemberDab = i
}

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}

func SysDictType() ISysDictType {
	if localSysDictType == nil {
		panic("implement not found for interface ISysDictType, forgot register?")
	}
	return localSysDictType
}

func RegisterSysDictType(i ISysDictType) {
	localSysDictType = i
}

func SysGenCodes() ISysGenCodes {
	if localSysGenCodes == nil {
		panic("implement not found for interface ISysGenCodes, forgot register?")
	}
	return localSysGenCodes
}

func RegisterSysGenCodes(i ISysGenCodes) {
	localSysGenCodes = i
}

func SysHgUser() ISysHgUser {
	if localSysHgUser == nil {
		panic("implement not found for interface ISysHgUser, forgot register?")
	}
	return localSysHgUser
}

func RegisterSysHgUser(i ISysHgUser) {
	localSysHgUser = i
}

func SysLogger() ISysLogger {
	if localSysLogger == nil {
		panic("implement not found for interface ISysLogger, forgot register?")
	}
	return localSysLogger
}

func RegisterSysLogger(i ISysLogger) {
	localSysLogger = i
}
