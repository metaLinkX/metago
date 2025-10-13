// Package handler

package handler

// handler.
import (
	"github.com/gogf/gf/v2/database/gdb"
)

// Option 预处理选项
type Option struct {
	FilterAuth bool // 过滤权限
}

// DefaultOption 默认预处理选项
var DefaultOption = &Option{
	FilterAuth: true,
}

func AdminModel(m *gdb.Model, opt ...*Option) *gdb.Model {
	var option *Option
	if len(opt) > 0 {
		option = opt[0]
	} else {
		option = DefaultOption
	}
	if option.FilterAuth {
		m = m.Handler(FilterAuth)
	}
	return m
}
