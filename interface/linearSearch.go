package a1

import (
	"reflect"
)

func linearSearch(x interface{}, lst interface{}) int {
	type_x := reflect.TypeOf(x)
	underlying_lst := reflect.ValueOf(lst)
	for i := 0; i < underlying_lst.Len(); i++ {
		element := underlying_lst.Index(i)
		if element.Type() != type_x {
			panic("type of x is different from type of list element")
		}
		if element.Interface() == x {
			return i
		}
	}
	return -1
}
