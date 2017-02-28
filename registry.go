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

	if reflect.ValueOf(funcImpl).Kind() != reflect.Func {
		return false
	}

	g_registry[funcName] = funcImpl
	return true
}

func DelElement(funcName string) bool {
	if _, ok := g_registry[funcName]; !ok {
		return true
	}
	delete(g_registry, funcName)
	return true
}

func ElementExists(funcName string) bool {
	if _, ok := g_registry[funcName]; ok {
		return true
	}

	return false
}

func Clear() bool {
	g_registry = nil
	g_registry = make(RegistryRecType)
	return true
}

func Call(funcName string, params ...interface{}) ([]reflect.Value, bool) {
	existingFunc, ok := g_registry[funcName]
	if !ok {
		return nil, false
	}

	if reflect.ValueOf(existingFunc).Kind() != reflect.Func {
		return nil, false
	}

	existingFuncValue := reflect.ValueOf(existingFunc)
	paramNum := len(params)
	if paramNum == 0 {
		return existingFuncValue.Call(nil), true
	}

	paramValues := make([]reflect.Value, paramNum)
	for k, v := range params {
		paramValues[k] = reflect.ValueOf(v)
	}

	return existingFuncValue.Call(paramValues), true
}

func Dump() {
	fmt.Println("Content of Registry:")
	fmt.Println("--------------------")
	for k, v := range g_registry {
		fmt.Println("r[", k, "]: ", v)
	}
}
