package utils

import (
	"reflect"
	"strings"
	"unicode"
)

func ConvertCamelToSnake(str string) string {
	if str == "ID" {
		return strings.ToLower(str)
	}

	var result []rune
	for i, r := range str {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, '_')
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, unicode.ToLower(r))
		}
	}
	return string(result)
}

func ConvertToJson(value reflect.Value) map[string]interface{} {
	fields := make(map[string]interface{})
	for i := 0; i < value.NumField(); i++ {
		fieldName := value.Type().Field(i).Name
		if fieldName == "Model" {
			for j := 0; j < value.Field(i).NumField(); j++ {
				subField := value.Field(i).Type().Field(j).Name
				snakeCaseName := ConvertCamelToSnake(subField)
				fields[snakeCaseName] = value.Field(i).Field(j).Interface()
			}
		} else {
			snakeCaseName := ConvertCamelToSnake(fieldName)
			fields[snakeCaseName] = value.Field(i).Interface()
		}
	}
	return fields
}
