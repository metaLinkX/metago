// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppMemberDao is the data access object for the table app_member.
type AppMemberDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AppMemberColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AppMemberColumns defines and stores column names for the table app_member.
type AppMemberColumns struct {
	Id                 string // 管理员ID
	RealName           string // 真实姓名
	Username           string // 帐号
	PasswordHash       string // 密码
	Salt               string // 密码盐
	PasswordResetToken string // 密码重置令牌
	Avatar             string // 头像
	Sex                string // 性别
	Email              string // 邮箱
	Mobile             string // 手机号码
	Birthday           string // 生日
	Address            string // 联系地址
	LastActiveAt       string // 最后活跃时间
	Remark             string // 备注
	Status             string // 状态
	CreatedAt          string // 创建时间
	UpdatedAt          string // 修改时间
}

// appMemberColumns holds the columns for the table app_member.
var appMemberColumns = AppMemberColumns{
	Id:                 "id",
	RealName:           "real_name",
	Username:           "username",
	PasswordHash:       "password_hash",
	Salt:               "salt",
	PasswordResetToken: "password_reset_token",
	Avatar:             "avatar",
	Sex:                "sex",
	Email:              "email",
	Mobile:             "mobile",
	Birthday:           "birthday",
	Address:            "address",
	LastActiveAt:       "last_active_at",
	Remark:             "remark",
	Status:             "status",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewAppMemberDao creates and returns a new DAO object for table data access.
func NewAppMemberDao(handlers ...gdb.ModelHandler) *AppMemberDao {
	return &AppMemberDao{
		group:    "default",
		table:    "app_member",
		columns:  appMemberColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppMemberDao) Columns() AppMemberColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppMemberDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AppMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
