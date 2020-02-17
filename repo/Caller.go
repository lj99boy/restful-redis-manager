package repo

import (
	"errors"
	"reflect"
)

func Call(m map[string]interface{}, name string, params ... interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	//inParamsLen := f.Type().NumIn()
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}

	in := make([]reflect.Value, len(params))
	//in := make([]reflect.Value, inParamsLen)
	//for i:=0;i<inParamsLen;i++{
	//	in[i] = reflect.ValueOf(params[i])
	//}
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}