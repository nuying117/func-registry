package funcgistry

import (
	"reflect"
	"testing"
)

func testSetup() {
	Clear()
}

func testTeardown() {

}

func testFunc1() {

}

func testFunc2(a, b int) int {
	return a + b
}

func TestAddFunction(t *testing.T) {
	testSetup()
	defer testTeardown()

	testData := []struct {
		inputKey   string
		inputValue interface{}
		expected   bool
	}{
		{"aString", "abcdefg", false},
		{"func1", testFunc1, true},
		{"empty", nil, false},
	}

	for _, data := range testData {
		result := AddElement(data.inputKey, data.inputValue)
		if result != data.expected {
			t.Errorf("AddElement(%q, %q) == %q, expected: %q", data.inputKey, data.inputValue, result, data.expected)
		}
	}
}
func TestDelFunction(t *testing.T) {
	testSetup()
	defer testTeardown()

	var result bool = false
	const funcName = "func1"
	var funcImpl = testFunc1
	const expectedAddResult = true
	const expectedExistsResultForAdd = true
	const expectedDelResult = true
	const expectedExistsResultForDel = false

	result = AddElement(funcName, funcImpl)
	if result != expectedAddResult {
		t.Errorf("AddElement(%q, %q) == %q, expected: %q", funcName, funcImpl, result, expectedAddResult)
	}

	result = ElementExists(funcName)
	if result != expectedExistsResultForAdd {
		t.Errorf("ElementExists(%q) == %q, expected: %q", funcName, result, expectedExistsResultForAdd)
	}

	result = DelElement(funcName)
	if result != expectedDelResult {
		t.Errorf("DelElement(%q) == %q, expected: %q", funcName, result, expectedDelResult)
	}

	result = ElementExists(funcName)
	if result != expectedExistsResultForDel {
		t.Errorf("ElementExists(%q) == %q, expected: %q", funcName, result, expectedExistsResultForDel)
	}
}
func TestCallInvalidFunction(t *testing.T) {
	testSetup()
	defer testTeardown()

	var result bool = false
	const funcName = "func1"
	const expectedResult = false

	_, result = Call(funcName)
	if result != expectedResult {
		t.Errorf("Call(%q) == %q, expected: %q", funcName, result, expectedResult)
	}
}
func TestCallValidFunction(t *testing.T) {
	testSetup()
	defer testTeardown()

	var result bool = false
	const expectedResult = true
	const func1Name = "func1"
	const func2Name = "func2"
	testFuncsData := []struct {
		funcName          string
		funcImpl          interface{}
		expectedAddResult bool
	}{
		{"func1", testFunc1, true},
		{"func2", testFunc2, true},
	}

	for _, testData := range testFuncsData {
		result = AddElement(testData.funcName, testData.funcImpl)
		if result != testData.expectedAddResult {
			t.Errorf("AddElement(%q, %q) == %q, expected: %q", testData.funcName, testData.funcImpl, result, testData.expectedAddResult)
		}
	}

	_, result = Call("func1")
	if result != expectedResult {
		t.Errorf("Call(%q) == %q, expected: %q", func1Name, result, expectedResult)
	}

	var rets []reflect.Value
	rets, result = Call("func2", 2, 3)
	if result != expectedResult {
		t.Errorf("Call(%q) == %q, expected: %q", func2Name, result, expectedResult)
	}

	if rets[0].Int() != 5 {
		t.Errorf("return value of func(%q)  == %q, expected: %q", func2Name, rets[0], 5)
	}
}
