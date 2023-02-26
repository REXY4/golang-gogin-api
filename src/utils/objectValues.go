package utils

import (
	"reflect"
)

func ObjectValues(obj interface{}) []interface{} {
	values := []interface{}{}
	reflectValue := reflect.ValueOf(obj)
	for i := 0; i < reflectValue.NumField(); i++ {
		value := reflectValue.Field(i).Interface()
		values = append(values, value)
	}

	return values
}

func FieldName(data interface{}) []string {
	var fieldNames []string
	for i := 0; i < reflect.TypeOf(data).NumField(); i++ {
		field := reflect.TypeOf(data).Field(i)
		fieldNames = append(fieldNames, field.Name)
	}
	return fieldNames
}
