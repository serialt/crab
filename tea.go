package crab

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Bool returns the truthy value of anything.
// If the value's type has a Bool() bool method, the method is called and returned.
// If the type has an IsZero() bool method, the opposite value is returned.
// Slices and maps are truthy if they have a length greater than zero.
// All other types are truthy if they are not their zero value.
// Play: https://go.dev/play/p/ETzeDJRSvhm
func Bool[T any](value T) bool {
	switch m := any(value).(type) {
	case interface{ Bool() bool }:
		return m.Bool()
	case interface{ IsZero() bool }:
		return !m.IsZero()
	}
	return reflectValue(&value)
}

func reflectValue(vp any) bool {
	switch rv := reflect.ValueOf(vp).Elem(); rv.Kind() {
	case reflect.Map, reflect.Slice:
		return rv.Len() != 0
	default:
		is := rv.IsZero()
		return !is
	}
}

// And returns true if both a and b are truthy.
// Play: https://go.dev/play/p/W1SSUmt6pvr
func And[T, U any](a T, b U) bool {
	return Bool(a) && Bool(b)
}

// Or returns false if neither a nor b is truthy.
// Play: https://go.dev/play/p/UlQTxHaeEkq
func Or[T, U any](a T, b U) bool {
	return Bool(a) || Bool(b)
}

// Xor returns true if a or b but not both is truthy.
// Play: https://go.dev/play/p/gObZrW7ZbG8
func Xor[T, U any](a T, b U) bool {
	valA := Bool(a)
	valB := Bool(b)
	return (valA || valB) && valA != valB
}

// Nor returns true if neither a nor b is truthy.
// Play: https://go.dev/play/p/g2j08F_zZky
func Nor[T, U any](a T, b U) bool {
	return !(Bool(a) || Bool(b))
}

// Xnor returns true if both a and b or neither a nor b are truthy.
// Play: https://go.dev/play/p/OuDB9g51643
func Xnor[T, U any](a T, b U) bool {
	valA := Bool(a)
	valB := Bool(b)
	return (valA && valB) || (!valA && !valB)
}

// Nand returns false if both a and b are truthy.
// Play: https://go.dev/play/p/vSRMLxLIbq8
func Nand[T, U any](a T, b U) bool {
	return !Bool(a) || !Bool(b)
}

// TernaryOperator checks the value of param `isTrue`, if true return ifValue else return elseValue.
// Play: https://go.dev/play/p/ElllPZY0guT
func TernaryOperator[T, U any](isTrue T, ifValue U, elseValue U) U {
	if Bool(isTrue) {
		return ifValue
	} else {
		return elseValue
	}
}

// ToBool convert string to boolean.
// Play: https://go.dev/play/p/ARht2WnGdIN
func ToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// ToBytes convert value to byte slice.
// Play: https://go.dev/play/p/fAMXYFDvOvr
func ToBytes(value any) ([]byte, error) {
	v := reflect.ValueOf(value)

	switch value.(type) {
	case int, int8, int16, int32, int64:
		number := v.Int()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case uint, uint8, uint16, uint32, uint64:
		number := v.Uint()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case float32:
		number := float32(v.Float())
		bits := math.Float32bits(number)
		bytes := make([]byte, 4)
		binary.BigEndian.PutUint32(bytes, bits)
		return bytes, nil
	case float64:
		number := v.Float()
		bits := math.Float64bits(number)
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, bits)
		return bytes, nil
	case bool:
		return strconv.AppendBool([]byte{}, v.Bool()), nil
	case string:
		return []byte(v.String()), nil
	case []byte:
		return v.Bytes(), nil
	default:
		newValue, err := json.Marshal(value)
		return newValue, err
	}
}

// ToChar convert string to char slice.
// Play: https://go.dev/play/p/JJ1SvbFkVdM
func ToChar(s string) []string {
	c := make([]string, 0)
	if len(s) == 0 {
		c = append(c, "")
	}
	for _, v := range s {
		c = append(c, string(v))
	}
	return c
}

// ToChannel convert a slice of elements to a read-only channel.
// Play: https://go.dev/play/p/hOx_oYZbAnL
func ToChannel[T any](array []T) <-chan T {
	ch := make(chan T)

	go func() {
		for _, item := range array {
			ch <- item
		}
		close(ch)
	}()

	return ch
}

// ToString convert value to string
// for number, string, []byte, will convert to string
// for other type (slice, map, array, struct) will call json.Marshal.
// Play: https://go.dev/play/p/nF1zOOslpQq
func ToString(value any) string {
	if value == nil {
		return ""
	}

	switch val := value.(type) {
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case int:
		return strconv.FormatInt(int64(val), 10)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case string:
		return val
	case []byte:
		return string(val)
	default:
		b, err := json.Marshal(val)
		if err != nil {
			return ""
		}
		return string(b)

		// todo: maybe we should't supprt other type conversion
		// v := reflect.ValueOf(value)
		// log.Panicf("Unsupported data type: %s ", v.String())
		// return ""
	}
}

// ToJson convert value to a json string.
// Play: https://go.dev/play/p/2rLIkMmXWvR
func ToJson(value any) (string, error) {
	result, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

// ToFloat convert value to float64, if input is not a float return 0.0 and error.
// Play: https://go.dev/play/p/4YTmPCibqHJ
func ToFloat(value any) (float64, error) {
	v := reflect.ValueOf(value)

	result := 0.0
	err := fmt.Errorf("ToInt: unvalid interface type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		result = float64(v.Int())
		return result, nil
	case uint, uint8, uint16, uint32, uint64:
		result = float64(v.Uint())
		return result, nil
	case float32, float64:
		result = v.Float()
		return result, nil
	case string:
		result, err = strconv.ParseFloat(v.String(), 64)
		if err != nil {
			result = 0.0
		}
		return result, err
	default:
		return result, err
	}
}

// ToInt convert value to int64 value, if input is not numerical, return 0 and error.
// Play: https://go.dev/play/p/9_h9vIt-QZ_b
func ToInt(value any) (int64, error) {
	v := reflect.ValueOf(value)

	var result int64
	err := fmt.Errorf("ToInt: invalid value type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		result = v.Int()
		return result, nil
	case uint, uint8, uint16, uint32, uint64:
		result = int64(v.Uint())
		return result, nil
	case float32, float64:
		result = int64(v.Float())
		return result, nil
	case string:
		result, err = strconv.ParseInt(v.String(), 0, 64)
		if err != nil {
			result = 0
		}
		return result, err
	default:
		return result, err
	}
}

// ToPointer returns a pointer to passed value.
// Play: https://go.dev/play/p/ASf_etHNlw1
func ToPointer[T any](value T) *T {
	return &value
}

// ColorHexToRGB convert hex color to rgb color.
// Play: https://go.dev/play/p/o7_ft-JCJBV
func ColorHexToRGB(colorHex string) (red, green, blue int) {
	colorHex = strings.TrimPrefix(colorHex, "#")
	color64, err := strconv.ParseInt(colorHex, 16, 32)
	if err != nil {
		return
	}
	color := int(color64)
	return color >> 16, (color & 0x00FF00) >> 8, color & 0x0000FF
}

// ColorRGBToHex convert rgb color to hex color.
// Play: https://go.dev/play/p/nzKS2Ro87J1
func ColorRGBToHex(red, green, blue int) string {
	r := strconv.FormatInt(int64(red), 16)
	g := strconv.FormatInt(int64(green), 16)
	b := strconv.FormatInt(int64(blue), 16)

	if len(r) == 1 {
		r = "0" + r
	}
	if len(g) == 1 {
		g = "0" + g
	}
	if len(b) == 1 {
		b = "0" + b
	}

	return "#" + r + g + b
}

// StringAllLetter 判断字符串是否只由字母组成
func StringIsLetter(str string) (bool, error) {
	return regexp.MatchString(`^[A-Za-z]+$`, str)
}

// StringTrim 去除字符串中的空格和换行符
func StringTrim(str string) string {
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	return StringTrimN(str)
}

// StringTrimN 去除字符串中的换行符
func StringTrimN(str string) string {
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	return str
}

// SubString 截取字符串
func SubString(res string, start, end int) string {
	if start > end || start < 0 || end > len(res) {
		return ""
	}
	return res[start:end]
}

// StringBuild 拼接字符串
func StringBuild(arrString ...string) string {
	return strings.Join(arrString, "")
}

// StringBuildSep 拼接字符串
func StringBuildSep(sep string, arrString ...string) string {
	return strings.Join(arrString, sep)
}

// FilterPrefix 根据前缀过滤slice
func FilterPrefix(strs []string, s string) (r []string) {
	for _, v := range strs {
		if len(v) >= len(s) {
			if v[:len(s)] == s {
				r = append(r, v)
			}
		}
	}

	return r
}

// FindLongestStr 查询最长字符串
func FindLongestStr(strs []string) string {
	longestStr := ""
	for _, str := range strs {
		if len(str) >= len(longestStr) {
			longestStr = str
		}
	}

	return longestStr
}

// ArrayToString 数字切片变字符串
func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}
