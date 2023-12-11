package ozon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func compareJsonResponse(t *testing.T, expectedJSON string, response interface{}) {
	err := json.Unmarshal([]byte(expectedJSON), response)
	if err != nil {
		t.Errorf("got error: %s", err)
		return
	}

	after, err := json.Marshal(response)
	if err != nil {
		t.Errorf("got error: %s", err)
		return
	}

	var j1, j2 map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(expectedJSON)).Decode(&j1); err != nil {
		t.Errorf("got error: %s", err)
		return
	}
	if err := json.NewDecoder(bytes.NewReader(after)).Decode(&j2); err != nil {
		t.Errorf("got error: %s", err)
		return
	}

	if err := compareJson(j1, j2, ""); err != nil {
		t.Errorf("jsons are not equal: %s", err)
		return
	}
}

func compareJson(expected interface{}, actual interface{}, prefix string) error {
	if expected == nil {
		return nil
	}

	expectedType := reflect.TypeOf(expected).Kind()
	actualType := reflect.TypeOf(actual).Kind()
	if expectedType != actualType {
		return fmt.Errorf("type for key %s is different: expected: %s, \ngot: %s", prefix, expectedType, actualType)
	}

	switch expected.(type) {
	case map[string]interface{}:
		expectedMap := expected.(map[string]interface{})
		actualMap := actual.(map[string]interface{})
		for k, v := range expectedMap {
			key := fmt.Sprintf("%s.%s", prefix, k)

			actualValue, ok := actualMap[k]
			if !ok {
				return fmt.Errorf("key %s is absent", key)
			}
			if err := compareJson(v, actualValue, key); err != nil {
				return err
			}
		}
	case []interface{}:
		expectedSlice := expected.([]interface{})
		actualSlice := actual.([]interface{})
		for i := range expectedSlice {
			key := fmt.Sprintf("%s.%d", prefix, i)
			if err := compareJson(expectedSlice[i], actualSlice[i], key); err != nil {
				return err
			}
		}
	default:
		if !reflect.DeepEqual(expected, actual) {
			return fmt.Errorf("value for key %s is different: expected: %s, \ngot: %s", prefix, expected, actual)
		}
	}

	return nil
}
