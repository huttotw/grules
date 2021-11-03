package grules

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		desc          string
		rule          string
		shouldError   bool
		expectedError string
	}{
		{
			desc: "valid rule",
			rule: `
			{
				"path": "foo.bar",
				"comparer": "eq",
				"value": 42
			}
			`,
			shouldError: false,
		},
		{
			desc: "valid rule with children",
			rule: `
			{
				"operator": "or",
				"rules": [
					{
						"path": "foo.bar",
						"comparer": "eq",
						"value": 42
					},
					{
						"path": "fizz.buzz",
						"comparer": "eq",
						"value": 24
					}
				]
			}
			`,
			shouldError: false,
		},
		{
			desc:          "nothing set",
			rule:          `{}`,
			shouldError:   true,
			expectedError: "must set either path, comparer, and value OR operator and rules",
		},
		{
			desc: "setting both base rule and child rules",
			rule: `
			{
				"path": "foo.bar",
				"comparer": "eq",
				"value": 42,
				"operator": "or",
				"rules": [
					{
						"path": "foo.bar",
						"comparer": "eq",
						"value": 42
					},
					{
						"path": "fizz.buzz",
						"comparer": "eq",
						"value": 24
					}
				]
			}
			`,
			shouldError:   true,
			expectedError: "setting path, comparer, and value AS WELL AS operator and rules is not valid",
		},
		{
			desc: "missing value for basic rule",
			rule: `
			{
				"path": "foo.bar",
				"comparer": "eq"
			}
			`,
			shouldError:   true,
			expectedError: "must set either path, comparer, and value OR operator and rules",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			var rule Rule
			err := json.NewDecoder(strings.NewReader(tc.rule)).Decode(&rule)
			if err != nil {
				t.Error("test case rule is not valid:", err.Error())
				t.FailNow()
			}

			err = rule.Validate()
			if tc.shouldError && err == nil {
				assert.Fail(t, "should have errored but didn't")
			}
			if err != nil {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}
