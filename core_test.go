package core

import (
	"log"
	"testing"
)

type TestTagDefaultValueStruct struct {
	TestString string `json:"test_string" default:"something"`
	TestNumber int    `json:"test_number" default:"12"`
}

func TestTagDefaultValue(t *testing.T) {
	testStruct := &TestTagDefaultValueStruct{}

	values, err := getDefaultValues(testStruct)
	if err != nil {
		log.Fatalf("error when getting default values from tags: %s", err)
	}

	expected := map[string]string{
		"test_string": "something",
		"test_number": "12",
	}

	if len(values) != len(expected) {
		log.Fatalf("expected equal length of values and expected: expected: %d, got: %d", len(expected), len(values))
	}
	for expKey, expValue := range expected {
		if expValue != values[expKey] {
			log.Fatalf("not equal values for key %s", expKey)
		}
	}
}
