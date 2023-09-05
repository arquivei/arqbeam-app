package app

import (
	"fmt"
	"reflect"
	"strings"
)

// nestedFieldValueString returns the value of a nested field inside a struct.
// The field is denoted by a dot separated string, like "Field1.Field2.Field3".
func nestedFieldValueString(data any, field string) string {
	fields := strings.Split(field, ".")
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < len(fields)-1; i++ {
		field := v.FieldByName(fields[i])
		if !field.IsValid() {
			return ""
		}

		v = field
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
	}

	lastField := v.FieldByName(fields[len(fields)-1])
	if !lastField.IsValid() {
		return ""
	}

	return fmt.Sprintf("%v", lastField.Interface())
}
