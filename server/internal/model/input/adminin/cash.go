// Package adminin

package adminin

import "server/internal/model/input/form"

// CashViewInp 获取信息
type CashViewInp struct {
	Id int64 `json:"id" v:"required#提现ID不能为空" dc:"提现ID"`
}

type CashViewModel struct {
	MemberCash
}

// CashListInp 获取列表
type CashListInp struct {
	form.PageReq
	form.StatusReq
	MemberId        int64    `json:"memberId"`
	CreatedAt       []int64  `json:"created_at"`
	ComplexMemberId []string `json:"complexMemberId" dc:"申请人"`
}

// CashApplyInp 申请提现
type CashApplyInp struct {
	Money    float64 `json:"money"     description:"提现金额"`
	MemberId int64
}

// CashPaymentInp 提现打款处理
type CashPaymentInp struct {
	Id     int64  `json:"id"        description:"ID"`
	Status int64  `json:"status"    description:"状态码"`
	Msg    string `json:"msg"       description:"处理结果"`
}
