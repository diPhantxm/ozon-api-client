package ozon

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func compareJsonResponse(t *testing.T, expectedJSON string, response interface{}) {
	err := json.Unmarshal([]byte(expectedJSON), response)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	after, err := json.Marshal(response)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	var j1, j2 map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(expectedJSON)).Decode(&j1); err != nil {
		t.Errorf("got error: %s", err)
	}
	if err := json.NewDecoder(bytes.NewReader(after)).Decode(&j2); err != nil {
		t.Errorf("got error: %s", err)
	}

	for k, v := range j2 {
		if k == "StatusCode" {
			continue
		}
		if k == "code" {
			continue
		}
		if k == "message" {
			continue
		}
		if k == "details" {
			continue
		}
		if !reflect.DeepEqual(j1[k], v) {
			t.Errorf("jsons are not equal for key %s: got: %s\nexpected: %s", k, j2[k], j1[k])
		}
	}
}
