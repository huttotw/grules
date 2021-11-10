package grules

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestNewSpecialSlice(t *testing.T) {
	testCases := []struct {
		desc     string
		value    gjson.Result
		expected specialSlice
	}{
		{
			desc:  "array of strings",
			value: gjson.Get(`{"strings":["foo","bar"]`, "strings"),
			expected: specialSlice{
				{"foo", "bar"},
			},
		},
		{
			desc:  "array of numbers",
			value: gjson.Get(`{"numbers":[1,2,3]`, "numbers"),
			expected: specialSlice{
				{float64(1), float64(2), float64(3)},
			},
		},
		{
			desc:  "array of bools",
			value: gjson.Get(`{"bools":[true, false, true]`, "bools"),
			expected: specialSlice{
				{true, false, true},
			},
		},
		{
			desc:  "non array",
			value: gjson.Get(`{"fizz":"buzz"`, "fizz"),
			expected: specialSlice{
				{"buzz"}, // this is what will happen if you don't pass in an actual array
			},
		},
		{
			desc:  "mixed types",
			value: gjson.Get(`{"mixedTypes":[true, 42, "foo"]`, "mixedTypes"),
			expected: specialSlice{
				{true, float64(42), "foo"},
			},
		},
		{
			desc:  "array of arrays",
			value: gjson.Get(`{"arrayOfArrays":[["foo", "bar"], ["fizz","buzz"]]`, "arrayOfArrays"),
			expected: specialSlice{
				{"foo", "bar"},
				{"fizz", "buzz"},
			},
		},
		{
			desc: "depth of 1",
			value: gjson.Get(`
			{
				"children": [
					{ "age": 21 },
					{ "age": 18 }
				]
			}
			`, "children.#.age"),
			expected: specialSlice{
				{float64(21), float64(18)},
			},
		},
		{
			desc: "depth of 2",
			value: gjson.Get(`
			{
				"parents": [
					{
						"children": [ { "age": 21 }, { "age": 18 } ]
					},
					{
						"children": [ { "age": 31 }, { "age": 12 } ]
					},
				]
			}
			`, "parents.#.children.#.age"),
			expected: specialSlice{
				{float64(21), float64(18)},
				{float64(31), float64(12)},
			},
		},
		{
			desc: "depth of 3",
			value: gjson.Get(`
			{
				"grandparents": [
					{
						"parents": [
							{
								"children": [ { "age": 21 }, { "age": 18 } ]
							},
							{
								"children": [ { "age": 31 }, { "age": 12 } ]
							},
						]
					},
					{
						"parents": [
							{
								"children": [ { "age": 23 }, { "age": 21 } ]
							},
							{
								"children": [ { "age": 35 }, { "age": 35 } ]
							},
						]
					},
				]
			}
			`, "grandparents.#.parents.#.children.#.age"),
			expected: specialSlice{
				{float64(21), float64(18)},
				{float64(31), float64(12)},
				{float64(23), float64(21)},
				{float64(35), float64(35)},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := newSpecialSlice(tc.value.Array())

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestSpecialSliceEvaluate(t *testing.T) {
	testCases := []struct {
		desc         string
		specialSlice specialSlice
		comparator   Comparator
		rule         Rule
		expected     bool
	}{
		{
			desc: "standard 'and' evaluation - true",
			specialSlice: specialSlice{
				{"foo", "bar", "fubar"},
				{"fizz", "buzz", "fubar"},
			},
			comparator: contains,
			rule: Rule{
				Operator: And,
				Value:    "fubar",
			},
			expected: true,
		},
		{
			desc: "standard 'and' evaluation - false",
			specialSlice: specialSlice{
				{"foo", "bar", "fubar"},
				{"fizz", "buzz"},
			},
			comparator: contains,
			rule: Rule{
				Operator: And,
				Value:    "fubar",
			},
			expected: false,
		},
		{
			desc: "standard 'or' evaluation - true",
			specialSlice: specialSlice{
				{"foo", "bar", "fubar"},
				{"fizz", "buzz"},
			},
			comparator: contains,
			rule: Rule{
				Operator: Or,
				Value:    "fubar",
			},
			expected: true,
		},
		{
			desc: "standard 'or' evaluation - false",
			specialSlice: specialSlice{
				{"foo", "bar"},
				{"fizz", "buzz"},
			},
			comparator: contains,
			rule: Rule{
				Operator: Or,
				Value:    "fubar",
			},
			expected: false,
		},
		{
			desc: "default to 'and'",
			specialSlice: specialSlice{
				{"foo", "bar"},
				{"fizz", "buzz", "fubar"},
			},
			comparator: contains,
			rule: Rule{
				// Operator not set
				Value: "fubar",
			},
			expected: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := tc.specialSlice.evalualte(tc.comparator, tc.rule)

			assert.Equal(t, tc.expected, result)
		})
	}
}
