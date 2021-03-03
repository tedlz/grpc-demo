package sign

import (
	"fmt"
	"reflect"
	"strings"
)

// StructToMap 结构体转为 map[string]interface{}
func StructToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("toMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		tagValue := fi.Tag.Get(tagName)
		if strings.Contains(tagValue, ",") {
			tagValue = strings.Split(tagValue, ",")[0]
		}
		if tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}
