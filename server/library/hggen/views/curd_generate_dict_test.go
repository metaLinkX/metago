package views

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gutil"
	"testing"
)

func TestName(t *testing.T) {
	content, err := Curd.GenerateDictConstFile(gctx.New(), "sys_user_sex")
	gutil.Dump(content, err)
}
