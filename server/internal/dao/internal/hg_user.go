// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HgUserDao is the data access object for the table hg_user.
type HgUserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  HgUserColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// HgUserColumns defines and stores column names for the table hg_user.
type HgUserColumns struct {
	Id            string // 用户ID
	Pid           string // 父级ID
	ComId         string // 小区ID
	HouseId       string // 房屋ID
	Username      string // 姓名
	Mobile        string // 手机号
	Password      string // 密码
	PasswordSalt  string // salt
	IdCard        string // 身份证
	IsAuthed      string // 是否已认证
	AuthedAt      string // 最新认证时间
	TotalParkTime string // 累计租住时长
	UserType      string // 租户类型
	UserStatus    string // 用户状态
	StartAt       string // 开始日期
	EndAt         string // 结束日期
	Remark        string // 备注
	NowIsPay      string // 本期是否已支付
	IdCardFront   string // 身份证正面
	IdCardBack    string // 身份证背面
	CreatedAt     string // 创建时间
	UpdatedAt     string // 修改时间
	RefoundRs     string // 退租理由
	IsChangedPass string // 是否已更换密码
	OpenId        string // openId
}

// hgUserColumns holds the columns for the table hg_user.
var hgUserColumns = HgUserColumns{
	Id:            "id",
	Pid:           "pid",
	ComId:         "comId",
	HouseId:       "houseId",
	Username:      "username",
	Mobile:        "mobile",
	Password:      "password",
	PasswordSalt:  "password_salt",
	IdCard:        "id_card",
	IsAuthed:      "isAuthed",
	AuthedAt:      "authedAt",
	TotalParkTime: "totalParkTime",
	UserType:      "user_type",
	UserStatus:    "user_status",
	StartAt:       "start_at",
	EndAt:         "end_at",
	Remark:        "remark",
	NowIsPay:      "nowIsPay",
	IdCardFront:   "idCardFront",
	IdCardBack:    "idCardBack",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	RefoundRs:     "refoundRs",
	IsChangedPass: "isChangedPass",
	OpenId:        "openId",
}

// NewHgUserDao creates and returns a new DAO object for table data access.
func NewHgUserDao(handlers ...gdb.ModelHandler) *HgUserDao {
	return &HgUserDao{
		group:    "default",
		table:    "hg_user",
		columns:  hgUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *HgUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *HgUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *HgUserDao) Columns() HgUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *HgUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *HgUserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *HgUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
