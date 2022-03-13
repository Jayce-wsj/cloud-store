package db

import (
	"database/sql"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Str2Bytes 字符串转字节切片
func Str2Bytes(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

// Bytes2Str 字节切片转字符串
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Btoi 布尔值转整形
func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Atoi8 字符串转换成 int8
func Atoi8(s string, d ...int8) int8 {
	i, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}
	return int8(i)
}

// Atoi16 字符串转换成 int8
func Atoi16(s string, d ...int16) int16 {
	i, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}
	return int16(i)
}

// Atoi 字符串转换成 int
func Atoi(s string, d ...int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}

	return i
}

// Atoi32 字符串转换成 int32
func Atoi32(s string, d ...int32) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}

	return int32(i)
}

// Atoi64 字符串转换成 int64
func Atoi64(s string, d ...int64) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}

	return i
}

// AtoUi8 字符串转换成 uint8
func AtoUi8(s string, d ...uint8) uint8 {
	i, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}
	return uint8(i)
}

// AtoUi16 字符串转换成 uint16
func AtoUi16(s string, d ...uint16) uint16 {
	i, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}
	return uint16(i)
}

// AtoUi 字符串转换成 uint
func AtoUi(s string, d ...uint) uint {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}
	return uint(i)
}

// AtoUi32 字符串转换成 uint32
func AtoUi32(s string, d ...uint32) uint32 {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}
	return uint32(i)
}

// AtoUi64 字符串转换成 uint64
func AtoUi64(s string, d ...uint64) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}

	return i
}

// Atof 字符串转换成 float32
func Atof(s string, d ...float32) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}

	return float32(f)
}

// Atof64 字符串转换成 float64
func Atof64(s string, d ...float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}

	return f
}

// I16toA int8 转字符串
func I8toA(i int8) string {
	return strconv.FormatInt(int64(i), 10)
}

// I16toA int16 转字符串
func I16toA(i int16) string {
	return strconv.FormatInt(int64(i), 10)
}

// Itoa int 转字符串
func Itoa(i int) string {
	return strconv.Itoa(i)
}

// I32toA int32 转字符串
func I32toA(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

// I64toA int64 转字符串
func I64toA(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Ui8toA uint8 转字符串
func Ui8toA(i uint8) string {
	return strconv.FormatUint(uint64(i), 10)
}

// Ui16toA uint16 转字符串
func Ui16toA(i uint16) string {
	return strconv.FormatUint(uint64(i), 10)
}

// UitoA uint 转字符串
func UitoA(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}

// Ui32toA uint32 转字符串
func Ui32toA(i uint32) string {
	return strconv.FormatUint(uint64(i), 10)
}

// Ui64toA uint64 转字符串
func Ui64toA(i uint64) string {
	return strconv.FormatUint(i, 10)
}

// F32toA float32 转字符串
func F32toA(f float32) string {
	return F64toA(float64(f))
}

// F64toA float64 转字符串
func F64toA(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// 返回一个带有Null值的数据库字符串
func NewNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// 返回一个带有Null值的数据库整形
func NewNullInt32(s int32, isNull bool) sql.NullInt32 {
	return sql.NullInt32{
		Int32: s,
		Valid: !isNull,
	}
}

// 返回一个带有Null值的数据库整形
func NewNullInt64(s int64, isNull bool) sql.NullInt64 {
	return sql.NullInt64{
		Int64: s,
		Valid: !isNull,
	}
}

// Ternary 模拟三元操作符
func Ternary(b bool, trueVal, falseVal interface{}) interface{} {
	if b {
		return trueVal
	}
	return falseVal
}

// Substr 截取字符串
// 例: abc你好1234
// Substr(str, 0) == abc你好1234
// Substr(str, 2) == c你好1234
// Substr(str, -2) == 34
// Substr(str, 2, 3) == c你好
// Substr(str, 0, -2) == 34
// Substr(str, 2, -1) == b
// Substr(str, -3, 2) == 23
// Substr(str, -3, -2) == 好1
func Substr(str string, start int, length ...int) string {
	rs := []rune(str)
	lth := len(rs)
	end := 0

	if start > lth {
		return ""
	}

	if len(length) == 1 {
		end = length[0]
	}

	// 从后数的某个位置向后截取
	if start < 0 {
		if -start >= lth {
			start = 0
		} else {
			start = lth + start
		}
	}

	if end == 0 {
		end = lth
	} else if end > 0 {
		end += start
		if end > lth {
			end = lth
		}
	} else { // 从指定位置向前截取
		if start == 0 {
			start = lth
		}
		start, end = start+end, start
	}
	if start < 0 {
		start = 0
	}

	return string(rs[start:end])
}

// Quote 对参数转码
func Quote(s string) string {
	return strings.Replace(strings.Replace(s, "'", "", -1), `\`, `\\`, -1)
}

// FullSQL 返回绑定完参数的完整的SQL语句
func FullSQL(str string, args ...interface{}) (string, error) {
	if !strings.Contains(str, "?") {
		return str, nil
	}
	sons := strings.Split(str, "?")

	var ret string
	var argIndex int
	var maxArgIndex = len(args)

	for _, son := range sons {
		ret += son

		if argIndex < maxArgIndex {
			switch v := args[argIndex].(type) {
			case int:
				ret += strconv.Itoa(v)
			case int8:
				ret += strconv.Itoa(int(v))
			case int16:
				ret += strconv.Itoa(int(v))
			case int32:
				ret += I64toA(int64(v))
			case int64:
				ret += I64toA(v)
			case uint:
				ret += UitoA(v)
			case uint8:
				ret += UitoA(uint(v))
			case uint16:
				ret += UitoA(uint(v))
			case uint32:
				ret += Ui32toA(v)
			case uint64:
				ret += Ui64toA(v)
			case float32:
				ret += F32toA(v)
			case float64:
				ret += F64toA(v)
			case *big.Int:
				ret += v.String()
			case bool:
				if v {
					ret += "true"
				} else {
					ret += "false"
				}
			case string:
				ret += "'" + Quote(v) + "'"
			case nil:
				ret += "NULL"
			default:
				return "", fmt.Errorf(
					"invalid sql argument type: %v => %v (sql: %s)",
					reflect.TypeOf(v).String(), v, str)
			}

			argIndex++
		}
	}

	return ret, nil
}
