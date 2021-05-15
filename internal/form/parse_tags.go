package form

import (
	"fmt"
	"reflect"
	"strconv"
)

func mapToString(v reflect.Value) []string {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	// Handle string
	case reflect.String:
		if len(v.String()) == 0 {
			return nil
		}
		return []string{v.String()}

	// Handle integers
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Int:
		return []string{strconv.Itoa(int(v.Int()))}

	case reflect.Slice:
		arr := []string{}
		s := reflect.ValueOf(v.Interface())
		for i := 0; i < s.Len(); i++ {
			arr = append(arr, mapToString(s.Index(i))...)
		}

		return arr
	default:
		return nil
	}
}

func ParseFormTags(v interface{}) map[string][]string {
	rType := reflect.TypeOf(v)
	rValue := reflect.ValueOf(v)

	params := map[string][]string{}

	for i := 0; i < rType.NumField(); i++ {
		f := rType.Field(i)
		tag := f.Tag.Get("form")
		if tag == "" {
			continue
		}

		value := mapToString(rValue.Field(i))
		if value == nil {
			continue
		}
		params[tag] = value
	}

	fmt.Printf("form_data.ParseFormTags: %#v\n", params)

	return params
}
