package core

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DefaultStructure struct {
	EmptyField string `json:"empty_field" default:"empty_string"`
	Field      string `json:"field" default:"string"`
}

type DefaultRequest struct {
	Field                  int                `json:"field" default:"100"`
	EmptyField             int                `json:"empty_field" default:"14"`
	Structure              DefaultStructure   `json:"structure"`
	Slice                  []DefaultStructure `json:"slice"`
	OptionalStructure      *DefaultStructure  `json:"optional_structure"`
	EmptyOptionalStructure *DefaultStructure  `json:"empty_optional_structure"`
}

func TestDefaultValues(t *testing.T) {
	req := &DefaultRequest{
		Field: 50,
		Structure: DefaultStructure{
			Field: "something",
		},
		Slice: []DefaultStructure{
			{
				Field: "something",
			},
			{
				Field: "something",
			},
		},
		OptionalStructure: &DefaultStructure{
			Field: "something",
		},
	}
	err := getDefaultValues(reflect.ValueOf(req))
	assert.Nil(t, err)

	assert.Equal(t, 50, req.Field)
	assert.Equal(t, 14, req.EmptyField)
	assert.Equal(t, "something", req.Structure.Field)
	assert.Equal(t, "empty_string", req.Structure.EmptyField)
	assert.Equal(t, "something", req.Slice[0].Field)
	assert.Equal(t, "something", req.Slice[1].Field)
	assert.Equal(t, "empty_string", req.Slice[1].EmptyField)
	assert.Equal(t, "empty_string", req.Slice[1].EmptyField)
	assert.Equal(t, "something", req.OptionalStructure.Field)
	assert.Equal(t, "empty_string", req.OptionalStructure.EmptyField)
	assert.Equal(t, (*DefaultStructure)(nil), req.EmptyOptionalStructure)
}
