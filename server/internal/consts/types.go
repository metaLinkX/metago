package consts

import "github.com/gogf/gf/v2/util/gconv"

const (
	TypeString      = "string"
	TypeInt         = "int"
	TypeInt8        = "int8"
	TypeInt16       = "int16"
	TypeInt32       = "int32"
	TypeInt64       = "int64"
	TypeUint        = "uint"
	TypeUint8       = "uint8"
	TypeUint16      = "uint16"
	TypeUint32      = "uint32"
	TypeUint64      = "uint64"
	TypeFloat32     = "float32"
	TypeFloat64     = "float64"
	TypeBool        = "bool"
	TypeDate        = "date"
	TypeDateTime    = "datetime"
	TypeSliceString = "[]string"
	TypeSliceInt    = "[]int"
	TypeSliceInt64  = "[]int64"
)

var Types = []string{TypeString,
	TypeInt, TypeInt8, TypeInt16, TypeInt32, TypeInt64,
	TypeUint, TypeUint8, TypeUint16, TypeUint32, TypeUint64,
	TypeFloat32, TypeFloat64,
	TypeBool,
	TypeDate, TypeDateTime,
	TypeSliceString, TypeSliceInt, TypeSliceInt64,
}

// ConvType 类型转换
func ConvType(val interface{}, t string) interface{} {
	switch t {
	case TypeString:
		val = gconv.String(val)
	case TypeInt:
		val = gconv.Int(val)
	case TypeInt8:
		val = gconv.Int8(val)
	case TypeInt16:
		val = gconv.Int16(val)
	case TypeInt32:
		val = gconv.Int32(val)
	case TypeInt64:
		val = gconv.Int64(val)
	case TypeUint:
		val = gconv.Uint(val)
	case TypeUint8:
		val = gconv.Uint8(val)
	case TypeUint16:
		val = gconv.Uint16(val)
	case TypeUint32:
		val = gconv.Uint32(val)
	case TypeUint64:
		val = gconv.Uint64(val)
	case TypeFloat32:
		val = gconv.Float32(val)
	case TypeFloat64:
		val = gconv.Float64(val)
	case TypeBool:
		val = gconv.Bool(val)
	case TypeDate:
		val = gconv.Time(val, "Y-m-d")
	case TypeDateTime:
		val = gconv.Time(val, "Y-m-d H:i:s")
	case TypeSliceInt:
		val = gconv.SliceInt(val)
	case TypeSliceInt64:
		val = gconv.SliceInt64(val)
	case TypeSliceString:
		val = gconv.SliceStr(val)
	default:
		val = gconv.String(val)
	}
	return val
}
