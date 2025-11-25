package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {

	type TestCase struct {
		input    string
		expected []string
	}

	testCases := []TestCase{
		{input: "hello world", expected: []string{"hello", "world"}},
		{input: "  leading and trailing  ", expected: []string{"leading", "and", "trailing"}},
		{input: "", expected: []string{}},
		{input: "singleword", expected: []string{"singleword"}},
	}

	for _, tc := range testCases {
		result := cleanInput(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("clean_input(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}
