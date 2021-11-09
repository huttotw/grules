package migrations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		desc          string
		v1            string
		expectedV2    string
		shouldError   bool
		expectedError string
	}{
		{
			desc: "basic",
			v1: `
			{
				"composites": [
					{
						"operator": "or",
						"rules": [
							{
								"comparator": "eq",
								"path": "user.firstName",
								"value": "stephen"
							}
						]
					}
				]
			}
			`,
			expectedV2: `
			{
				"comparator": "eq",
				"path": "user.firstName",
				"value": "stephen"
			}
			`,
		},
		{
			desc: "multiple rules",
			v1: `
			{
				"composites": [
					{
						"operator": "or",
						"rules": [
							{
								"comparator": "eq",
								"path": "user.firstName",
								"value": "stephen"
							},
							{
								"comparator": "eq",
								"path": "user.lastName",
								"value": "stanton"
							}
						]
					}
				]
			}
			`,
			expectedV2: `
			{
				"operator": "or",
				"rules": [
					{
						"comparator": "eq",
						"path": "user.firstName",
						"value": "stephen"
					},
					{
						"comparator": "eq",
						"path": "user.lastName",
						"value": "stanton"
					}
				]
			}
			`,
		},
		{
			desc: "multiple composites",
			v1: `
			{
				"composites": [
					{
						"operator": "or",
						"rules": [
							{
								"comparator": "eq",
								"path": "user.firstName",
								"value": "stephen"
							}
						]
					},
					{
						"operator": "or",
						"rules": [
							{
								"comparator": "eq",
								"path": "user.lastName",
								"value": "stanton"
							}
						]
					}
				]
			}
			`,
			expectedV2: `
			{
				"operator": "and",
				"rules": [
					{
						"comparator": "eq",
						"path": "user.firstName",
						"value": "stephen"
					},
					{
						"comparator": "eq",
						"path": "user.lastName",
						"value": "stanton"
					}
				]
			}
			`,
		},
		{
			desc: "multiple composites, multiple rules",
			v1: `
			{
				"composites": [
					{
						"operator": "or",
						"rules": [
							{
								"comparator": "eq",
								"path": "user.firstName",
								"value": "stephen"
							},
							{
								"comparator": "eq",
								"path": "user.lastName",
								"value": "stanton"
							}
						]
					},
					{
						"operator": "and",
						"rules": [
							{
								"comparator": "eq",
								"path": "user.firstName",
								"value": "stephen"
							},
							{
								"comparator": "eq",
								"path": "user.lastName",
								"value": "stanton"
							}
						]
					}
				]
			}
			`,
			expectedV2: `
			{
				"operator": "and",
				"rules": [
					{
						"operator": "or",
						"rules": [
							{
								"comparator": "eq",
								"path": "user.firstName",
								"value": "stephen"
							},
							{
								"comparator": "eq",
								"path": "user.lastName",
								"value": "stanton"
							}
						]
					},
					{
						"operator": "and",
						"rules": [
							{
								"comparator": "eq",
								"path": "user.firstName",
								"value": "stephen"
							},
							{
								"comparator": "eq",
								"path": "user.lastName",
								"value": "stanton"
							}
						]
					}
				]
			}
			`,
		},
		{
			desc: "no composites",
			v1: `
			{
				"composites": []
			}
			`,
			shouldError:   true,
			expectedError: "converting to v2: no composites found",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			v2, err := MigrateGrulesV1ToV2(tc.v1)
			if tc.shouldError && err == nil {
				assert.Fail(t, "should have errored but didn't")
			}
			if err != nil {
				assert.EqualError(t, err, tc.expectedError)
			} else {
				assert.JSONEq(t, tc.expectedV2, v2)
			}
		})
	}
}
