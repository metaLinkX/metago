package consts

const (
	DictRoleDataAuthAll = 1 // 全部权限

	// 通过部门划分
	DictRoleDataAuthNowDept    = 2 // 当前部门
	DictRoleDataAuthDeptAndSub = 3 // 当前部门及以下部门
	DictRoleDataAuthDeptCustom = 4 // 自定义部门

	// 通过上下级关系划分
	DictRoleDataAuthSelf          = 5 // 仅自己
	DictRoleDataAuthSelfAndSub    = 6 // 自己和直属下级
	DictRoleDataAuthSelfAndAllSub = 7 // 自己和全部下级
)
