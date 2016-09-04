// Package params provides a reflection-based parser for URL parameters.
package params

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func Pack(url string, ptr interface{}) string {
	var buf bytes.Buffer
	buf.WriteByte('?')

	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			buf.WriteByte('&')
		}
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		pack(&buf, name, v.Field(i))
	}

	return url + buf.String()
}

func pack(buf *bytes.Buffer, key string, val reflect.Value) {
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			if i > 0 {
				buf.WriteByte('&')
			}
			fmt.Fprintf(buf, "%s=", key)
			populate(buf, val.Index(i))
		}
	} else {
		fmt.Fprintf(buf, "%s=", key)
		populate(buf, val)
	}
}

func populate(buf *bytes.Buffer, v reflect.Value) {
	switch v.Kind() {
	case reflect.String:
		fmt.Fprintf(buf, "%s", v.String())

	case reflect.Int:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Bool:
		fmt.Fprintf(buf, "%t", v.Bool())

	default:
		panic(fmt.Sprintf("unsupported kind %s", v.Type()))
	}
}
