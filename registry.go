package funcgistry

import (
	"fmt"
	"reflect"
)

type RegistryRecType map[string]interface{}

var g_registry RegistryRecType = make(RegistryRecType)

func AddElement(funcName string, funcImpl interface{}) bool {
	if existingValue, ok := g_registry[funcName]; ok || existingValue != nil {
		return false
	}

	if funcImpl == nil {
		return false
	}

	g_registry[funcName] = funcImpl
	return true
}

func DelElement(funcName string) bool {
	if _, ok := g_registry[funcName]; !ok {
		return true
	}
	g_registry[funcName] = nil
	return true
}

func Call(funcName string, params ...interface{}) []reflect.Value {
	existingFunc, ok := g_registry[funcName]
	if !ok {
		return nil
	}

	existingFuncValue := reflect.ValueOf(existingFunc)
	paramNum := len(params)
	if paramNum == 0 {
		return existingFuncValue.Call(nil)
	}

	paramValues := make([]reflect.Value, paramNum)
	for k, v := range params {
		paramValues[k] = reflect.ValueOf(v)
	}

	return existingFuncValue.Call(paramValues)
}

func Dump() {
	fmt.Println("Content of Registry:")
	fmt.Println("--------------------")
	for k, v := range g_registry {
		fmt.Println("r[", k, "]: ", v)
	}
}
