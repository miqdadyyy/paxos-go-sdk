package util

import (
	"fmt"
	"reflect"
	"strings"
)

func GenerateQueryFromStruct(queryRequest interface{}, excludes ...string) string {
	v := reflect.ValueOf(queryRequest)
	var res []string

	for i := 0; i < v.NumField(); i++ {
		val := fmt.Sprintf("%v", v.Field(i).Interface())
		if val != "" {
			isExclude := false
			tag := reflect.TypeOf(queryRequest).Field(i).Tag.Get("json")
			for _, excludeItem := range excludes {
				if tag == excludeItem {
					isExclude = true
					break
				}
			}

			if !isExclude {
				res = append(res, tag+"="+val)
			}
		}
	}

	return strings.Join(res, "&")
}

func GenerateBodyFromStruct(requestBody interface{}, excludes ...string) map[string]string {
	v := reflect.ValueOf(requestBody)
	res := make(map[string]string)

	for i := 0; i < v.NumField(); i++ {
		val := fmt.Sprintf("%v", v.Field(i).Interface())
		if val != "" {
			isExclude := false
			tag := reflect.TypeOf(requestBody).Field(i).Tag.Get("json")
			for _, excludeItem := range excludes {
				if tag == excludeItem {
					isExclude = true
					break
				}
			}

			if !isExclude {
				res[tag] = val
			}
		}
	}
	return res
}
