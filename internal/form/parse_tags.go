package form

import (
	"log"
	"net/url"
	"reflect"
	"strconv"
)

func mapToString(v reflect.Value) []string {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Invalid:
		return nil

	case reflect.String:
		if len(v.String()) == 0 {
			return nil
		}
		return []string{v.String()}

	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		return []string{strconv.Itoa(int(v.Int()))}

	case reflect.Bool:
		val := "0"
		if v.Bool() {
			val = "1"
		}
		return []string{val}

	case reflect.Slice:
		arr := []string{}
		s := reflect.ValueOf(v.Interface())
		for i := 0; i < s.Len(); i++ {
			arr = append(arr, mapToString(s.Index(i))...)
		}

		return arr
	default:
		log.Printf("form.parse_tags: unmapped type %s", v.Kind())
		return nil
	}
}

func ParseFormTags(v interface{}) url.Values {
	uv := url.Values{}
	rType := reflect.TypeOf(v)
	rValue := reflect.ValueOf(v)

	for i := 0; i < rType.NumField(); i++ {
		f := rType.Field(i)
		tag := f.Tag.Get("form")
		if tag == "" {
			continue
		}

		values := mapToString(rValue.Field(i))
		if values == nil {
			continue
		}

		for _, val := range values {
			uv.Add(tag, val)
		}
	}
	log.Printf("form_data.ParseFormTags: %#v", uv)

	return uv
}
