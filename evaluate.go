package grules

import (
	jsonencoding "encoding/json"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

// NOTE the string returned will say why the rule did not evaluate. I am returning the string now
// as to not change the function signature for a future version when this feature comes out.
func Evaluate(json string, rule string) (bool, string) {
	object := gjson.Parse(json)

	var r Rule
	err := jsonencoding.NewDecoder(strings.NewReader(rule)).Decode(&r)
	if err != nil {
		return false, fmt.Sprintf("could not decode rule, %s", err.Error())
	}

	return evaluateObject(object, r), "REASON NOT YET IMPLEMENTED"
}

func evaluateObject(object gjson.Result, rule Rule) bool {
	if rule.HasChildren() {
		return evaluateMultiRule(object, rule.Rules, rule.Operator)
	}

	value := object.Get(rule.Path)
	if !value.Exists() {
		return false
	}

	if !typeMatches(value, rule) {
		return false
	}

	compare, found := defaultComparers[rule.Comparer]
	if !found {
		return false
	}

	if value.IsArray() {
		return evaluateArrayOfPrimitives(value.Array(), rule, compare)
	}

	return evaluatePrimitive(value, rule, compare)
}

func evaluateMultiRule(object gjson.Result, rules []Rule, operator Operator) bool {
	switch operator {
	case Or:
		for _, rule := range rules {
			evalTrue := evaluateObject(object, rule)
			if evalTrue {
				return true
			}
		}

		return false
	case And:
		fallthrough
	default:
		for _, rule := range rules {
			evalTrue := evaluateObject(object, rule)
			if !evalTrue {
				return false
			}
		}

		return true
	}
}

func evaluatePrimitive(value gjson.Result, rule Rule, compare Compare) bool {
	switch value.Type {
	case gjson.String:
		return compare(value.Str, rule.Value)
	case gjson.Number:
		return compare(value.Num, rule.Value)
	case gjson.True:
		fallthrough
	case gjson.False:
		return compare(value.Bool(), rule.Value)
	default:
		if value.IsArray() {
			slice := transformGJSONArrayToSlice(value.Array())
			return compare(slice, rule.Value)
		}
	}

	return false
}

func transformGJSONArrayToSlice(values []gjson.Result) []interface{} {
	var slice []interface{}
	for _, value := range values {
		switch value.Type {
		case gjson.String:
			slice = append(slice, value.Str)
		case gjson.Number:
			slice = append(slice, value.Num)
		case gjson.True:
			fallthrough
		case gjson.False:
			slice = append(slice, value.Bool())
		default:
			if value.IsArray() {
				slice = append(slice, transformGJSONArrayToSlice(value.Array()))
			}
		}
	}

	return slice
}

func evaluateArrayOfPrimitives(values []gjson.Result, rule Rule, compare Compare) bool {
	switch rule.Operator {
	case Or:
		for _, value := range values {
			if value.IsArray() {
				return evaluateArrayOfPrimitives(value.Array(), rule, compare)
			}

			evalTrue := evaluatePrimitive(value, rule, compare)
			if evalTrue {
				return true
			}
		}

		return false
	case And:
		fallthrough
	default:
		for _, value := range values {
			evalTrue := evaluatePrimitive(value, rule, compare)
			if !evalTrue {
				return false
			}
		}

		return true
	}
}

// NOTE you can see where there could be a problem with array of arrays. A problem for another day
func typeMatches(result gjson.Result, rule Rule) bool {
	if result.IsArray() {
		for _, eachResult := range result.Array() {
			match := typeMatches(eachResult, rule)
			if !match {
				return false
			}
		}

		return true
	}

	switch rule.Value.(type) {
	case float64:
		return result.Type == gjson.Number
	case string:
		return result.Type == gjson.String
	case bool:
		return (result.Type == gjson.True || result.Type == gjson.False)
	default:
		return false
	}
}
