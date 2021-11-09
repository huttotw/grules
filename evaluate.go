package grules

import (
	"github.com/tidwall/gjson"
)

// NOTE the string returned will say why the rule did not evaluate. I am returning the string now
// as to not change the function signature for a future version when this feature comes out.
func Evaluate(json string, rule Rule) (bool, string) {
	object := gjson.Parse(json)

	return evaluateObject(object, rule), "REASON NOT YET IMPLEMENTED"
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

	comparator, found := defaultComparators[rule.Comparator]
	if !found {
		return false
	}

	if value.IsArray() {
		return newSpecialSlice(value.Array()).evalualte(comparator, rule)
	}

	return evaluatePrimitive(value, rule, comparator)
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

func evaluatePrimitive(value gjson.Result, rule Rule, comparator Comparator) bool {
	switch value.Type {
	case gjson.String:
		return comparator(value.Str, rule.Value)
	case gjson.Number:
		return comparator(value.Num, rule.Value)
	case gjson.True:
		fallthrough
	case gjson.False:
		return comparator(value.Bool(), rule.Value)
	default:
		if value.IsArray() {
			slice := transformGJSONArrayToSlice(value.Array())
			return comparator(slice, rule.Value)
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
				slice = append(slice, transformGJSONArrayToSlice(value.Array())...)
			}
		}
	}

	return slice
}

func evaluateArrayOfPrimitives(values []gjson.Result, rule Rule, comparator Comparator) bool {
	switch rule.Operator {
	case Or:
		for _, value := range values {
			if value.IsArray() {
				return evaluateArrayOfPrimitives(value.Array(), rule, comparator)
			}

			evalTrue := evaluatePrimitive(value, rule, comparator)
			if evalTrue {
				return true
			}
		}

		return false
	case And:
		fallthrough
	default:
		for _, value := range values {
			evalTrue := evaluatePrimitive(value, rule, comparator)
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
