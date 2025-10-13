package system

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"server/internal/consts"
	"server/library/validate"
)

func InitServerRoot(ctx context.Context) {
	serverRoot := g.Cfg().MustGet(ctx, "server.serverRoot", "").String()
	if serverRoot != "" {
		err := gfile.Mkdir(serverRoot)
		if err != nil {
			panic(err)
		}
	}
}

// InitGFMode 设置gf运行模式
func InitGFMode(ctx context.Context) {
	mode := g.Cfg().MustGet(ctx, "system.mode").String()
	if len(mode) == 0 {
		mode = gmode.NOT_SET
	}

	var modes = []string{gmode.DEVELOP, gmode.TESTING, gmode.STAGING, gmode.PRODUCT}

	// 如果是有效的运行模式，就进行设置
	if validate.InSlice(modes, mode) {
		gmode.Set(mode)
	}
}

func Debug(ctx context.Context) bool {
	return g.Cfg().MustGet(ctx, "system.debug", true).Bool()
}

func DefaultErrorTplContent(ctx context.Context) string {
	return gfile.GetContents(g.Cfg().MustGet(ctx, "viewer.paths").String() + "/error/default.html")
}

// ConvType 类型转换
func ConvType(val interface{}, t string) interface{} {
	switch t {
	case consts.TypeString:
		val = gconv.String(val)
	case consts.TypeInt:
		val = gconv.Int(val)
	case consts.TypeInt8:
		val = gconv.Int8(val)
	case consts.TypeInt16:
		val = gconv.Int16(val)
	case consts.TypeInt32:
		val = gconv.Int32(val)
	case consts.TypeInt64:
		val = gconv.Int64(val)
	case consts.TypeUint:
		val = gconv.Uint(val)
	case consts.TypeUint8:
		val = gconv.Uint8(val)
	case consts.TypeUint16:
		val = gconv.Uint16(val)
	case consts.TypeUint32:
		val = gconv.Uint32(val)
	case consts.TypeUint64:
		val = gconv.Uint64(val)
	case consts.TypeFloat32:
		val = gconv.Float32(val)
	case consts.TypeFloat64:
		val = gconv.Float64(val)
	case consts.TypeBool:
		val = gconv.Bool(val)
	case consts.TypeDate:
		val = gconv.Time(val, "Y-m-d")
	case consts.TypeDateTime:
		val = gconv.Time(val, "Y-m-d H:i:s")
	case consts.TypeSliceInt:
		val = gconv.SliceInt(val)
	case consts.TypeSliceInt64:
		val = gconv.SliceInt64(val)
	case consts.TypeSliceString:
		val = gconv.SliceStr(val)
	default:
		val = gconv.String(val)
	}
	return val
}
