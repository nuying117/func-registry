package funcgistry

import (
	"testing"
)

func testSetup() {
	Clear()
}

func testTeardown() {

}

func testFunc1() {

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
		{"func1", testFunc1, false},
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
