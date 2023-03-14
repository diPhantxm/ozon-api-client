package core

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

type CommonResponse struct {
	StatusCode int
	Code       int                    `json:"code"`
	Details    []CommonResponseDetail `json:"details"`
	Message    string                 `json:"message"`
}

type CommonResponseDetail struct {
	TypeUrl string `json:"typeUrl"`
	Value   string `json:"value"`
}

type Response struct {
	CommonResponse
	Data interface{}
}

func (r Response) CopyCommonResponse(rhs *CommonResponse) {
	rhs.Code = r.Code
	rhs.Details = r.Details
	rhs.StatusCode = r.StatusCode
	rhs.Message = r.Message
}

func getDefaultValues(v interface{}) (map[string]string, error) {
	isNil, err := isZero(v)
	if err != nil {
		return make(map[string]string), err
	}

	if isNil {
		return make(map[string]string), nil
	}

	out := make(map[string]string)

	vType := reflect.TypeOf(v).Elem()
	vValue := reflect.ValueOf(v).Elem()

	for i := 0; i < vType.NumField(); i++ {
		field := vType.Field(i)
		tag := field.Tag.Get("json")
		defaultValue := field.Tag.Get("default")

		if field.Type.Kind() == reflect.Slice {
			// Attach any slices as query params
			fieldVal := vValue.Field(i)
			for j := 0; j < fieldVal.Len(); j++ {
				out[tag] = fmt.Sprintf("%v", fieldVal.Index(j))
			}
		} else {
			// Add any scalar values as query params
			fieldVal := fmt.Sprintf("%v", vValue.Field(i))

			// If no value was set by the user, use the default
			// value specified in the struct tag.
			if fieldVal == "" || fieldVal == "0" {
				if defaultValue == "" {
					continue
				}

				fieldVal = defaultValue
			}

			out[tag] = fmt.Sprintf("%v", fieldVal)
		}
	}

	return out, nil
}

func buildRawQuery(req *http.Request, v interface{}) (string, error) {
	query := req.URL.Query()

	if v == nil {
		return query.Encode(), nil
	}

	values, err := getDefaultValues(v)
	if err != nil {
		return "", err
	}
	for k, v := range values {
		query.Add(k, v)
	}

	return query.Encode(), nil
}

func isZero(v interface{}) (bool, error) {
	t := reflect.TypeOf(v)
	if !t.Comparable() {
		return false, fmt.Errorf("type is not comparable: %v", t)
	}
	return v == reflect.Zero(t).Interface(), nil
}

func TimeFromString(t *testing.T, format, datetime string) time.Time {
	dt, err := time.Parse(format, datetime)
	if err != nil {
		t.Errorf("error when parsing time: %s", err)
	}
	return dt
}
