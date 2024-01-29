package core

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
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

func getDefaultValues(v reflect.Value) error {
	vValue := v.Elem()
	vType := vValue.Type()

	for i := 0; i < vType.NumField(); i++ {
		field := vType.Field(i)

		switch field.Type.Kind() {
		case reflect.Slice:
			for j := 0; j < vValue.Field(i).Len(); j++ {
				// skip if slice type is primitive
				if vValue.Field(i).Index(j).Kind() != reflect.Struct {
					continue
				}

				// Attach any slices as query params
				err := getDefaultValues(vValue.Field(i).Index(j).Addr())
				if err != nil {
					return err
				}
			}
		case reflect.String:
			isNil, err := isZero(vValue.Field(i).Addr())
			if err != nil {
				return err
			}
			if !isNil {
				continue
			}

			defaultValue, ok := field.Tag.Lookup("default")
			if !ok {
				continue
			}

			vValue.Field(i).SetString(defaultValue)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			isNil, err := isZero(vValue.Field(i).Addr())
			if err != nil {
				return err
			}
			if !isNil {
				continue
			}

			defaultValue, ok := field.Tag.Lookup("default")
			if !ok {
				continue
			}
			defaultValueInt, err := strconv.Atoi(defaultValue)
			if err != nil {
				return err
			}

			vValue.Field(i).SetInt(int64(defaultValueInt))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			isNil, err := isZero(vValue.Field(i).Addr())
			if err != nil {
				return err
			}
			if !isNil {
				continue
			}

			defaultValue, ok := field.Tag.Lookup("default")
			if !ok {
				continue
			}
			defaultValueUint, err := strconv.ParseUint(defaultValue, 10, 64)
			if err != nil {
				return err
			}

			vValue.Field(i).SetUint(uint64(defaultValueUint))
		case reflect.Struct:
			err := getDefaultValues(vValue.Field(i).Addr())
			if err != nil {
				return err
			}
		case reflect.Ptr:
			isNil, err := isZero(vValue.Field(i).Addr())
			if err != nil {
				return err
			}
			if isNil {
				continue
			}

			if err := getDefaultValues(vValue.Field(i)); err != nil {
				return err
			}
		default:
			continue
		}
	}

	return nil
}

func buildRawQuery(req *http.Request, v interface{}) (string, error) {
	query := req.URL.Query()

	if v == nil {
		return query.Encode(), nil
	}

	err := getDefaultValues(reflect.ValueOf(v))
	if err != nil {
		return "", err
	}

	return query.Encode(), nil
}

func isZero(v reflect.Value) (bool, error) {
	t := v.Elem().Type()
	if !t.Comparable() {
		return false, fmt.Errorf("type is not comparable: %v", t)
	}
	return reflect.Zero(t).Equal(v.Elem()), nil
}

func TimeFromString(t *testing.T, format, datetime string) time.Time {
	dt, err := time.Parse(format, datetime)
	if err != nil {
		t.Errorf("error when parsing time: %s", err)
	}
	return dt
}
