package core

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

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

func TestTimeFormat(t *testing.T) {
	t.Run("Time Format Marshalling", func(t *testing.T) {
		tests := []struct {
			ft           *TimeFormat
			layout       string
			expectedJSON string
			diff         time.Duration
		}{
			{
				ft:           NewTimeFormat(time.Date(2024, 4, 30, 15, 42, 12, 55, time.FixedZone("Test Zone", 0)), ShortDateLayout),
				layout:       ShortDateLayout,
				expectedJSON: `"2024-04-30"`,
				diff:         time.Hour * 24,
			},
			{
				ft:           NewTimeFormat(time.Date(2024, 4, 30, 0, 0, 0, 0, time.FixedZone("Test Zone", 0)), ShortDateLayout),
				layout:       ShortDateLayout,
				expectedJSON: `"2024-04-30"`,
				diff:         time.Hour * 24,
			},
			{
				ft:           NewTimeFormat(time.Time{}, ShortDateLayout),
				layout:       ShortDateLayout,
				expectedJSON: `"0001-01-01"`,
				diff:         time.Hour * 24,
			},
			{
				ft:           nil,
				layout:       ShortDateLayout,
				expectedJSON: `null`,
				diff:         time.Hour * 24,
			},
		}

		for _, tc := range tests {
			marshaled, err := json.Marshal(tc.ft)
			assert.Equal(t, nil, err)
			assert.Equal(t, tc.expectedJSON, string(marshaled))

			unmarshaled := newTimeLayout(tc.layout)
			err = json.Unmarshal(marshaled, unmarshaled)
			assert.Equal(t, nil, err)

			if tc.ft != nil {
				diffedTime := tc.ft.Add(-tc.diff)
				assert.Equal(t, true, diffedTime.Before(unmarshaled.Time) || diffedTime.Equal(unmarshaled.Time))
				assert.Equal(t, true, tc.ft.After(unmarshaled.Time) || tc.ft.Equal(unmarshaled.Time))
			}
		}
	})

	t.Run("Time Format in structure Marshalling", func(t *testing.T) {
		type test struct {
			Date *TimeFormat `json:"date"`
		}

		tests := []struct {
			structure    *test
			layout       string
			expectedJSON string
			diff         time.Duration
		}{
			{
				structure:    &test{Date: NewTimeFormat(time.Date(2024, 4, 30, 5, 4, 7, 20, time.FixedZone("Test Zone", 0)), ShortDateLayout)},
				layout:       ShortDateLayout,
				expectedJSON: `{"date":"2024-04-30"}`,
				diff:         time.Hour * 24,
			},
			{
				structure:    &test{Date: nil},
				layout:       ShortDateLayout,
				expectedJSON: `{"date":null}`,
				diff:         time.Hour * 24,
			},
		}

		for _, tc := range tests {
			marshaled, err := json.Marshal(tc.structure)
			assert.Equal(t, nil, err)
			assert.Equal(t, tc.expectedJSON, string(marshaled))

			unmarshaled := &test{Date: newTimeLayout(tc.layout)}
			err = json.Unmarshal(marshaled, unmarshaled)
			assert.Equal(t, nil, err)

			if tc.structure != nil && tc.structure.Date != nil {
				diffedTime := tc.structure.Date.Add(-tc.diff)
				assert.Equal(t, true, diffedTime.Before(unmarshaled.Date.Time) || diffedTime.Equal(unmarshaled.Date.Time))
				assert.Equal(t, true, tc.structure.Date.After(unmarshaled.Date.Time) || tc.structure.Date.Equal(unmarshaled.Date.Time))
			}
		}
	})
}
