package utilities

import "reflect"

func InstanceOf(object any, typeToCheck any) bool {
	return reflect.TypeOf(object) == reflect.TypeOf(typeToCheck)
}
