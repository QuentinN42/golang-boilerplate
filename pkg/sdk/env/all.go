package env

import (
	"fmt"
	"reflect"
)

func GetAll[T any](container *T) (*T, error) {
	ctn := reflect.Indirect(reflect.ValueOf(container))
	if ctn.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Container must be a struct")
	}

	for i := 0; i < ctn.NumField(); i++ {
		// Skip unexported fields
		if !ctn.Field(i).CanSet() {
			continue
		}

		field := reflect.Indirect(reflect.ValueOf(container)).Field(i)
		key := ctn.Type().Field(i).Name
		typ := ctn.Type().Field(i).Type

		if typ.Kind() == reflect.Ptr {
			switch typ.Elem().Kind() {
			case reflect.String:
				val, _ := Get(key)
				field.Set(reflect.ValueOf(&val))
			case reflect.Int:
				val, _ := GetInt(key)
				field.Set(reflect.ValueOf(&val))
			case reflect.Bool:
				val := GetBool(key)
				field.Set(reflect.ValueOf(&val))
			default:
				return nil, fmt.Errorf("Unsupported type %s for value %s", typ.Elem().Kind(), key)
			}
		} else {
			switch field.Kind() {
			case reflect.String:
				val, err := Get(key)
				if err != nil {
					return nil, err
				}
				field.SetString(val)
			case reflect.Int:
				val, err := GetInt(key)
				if err != nil {
					return nil, err
				}
				field.SetInt(int64(val))
			case reflect.Bool:
				field.SetBool(GetBool(key))
			default:
				return nil, fmt.Errorf("Unsupported type %s for value %s", field.Kind(), key)
			}
		}
	}
	return container, nil
}
