package grules

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	// OperatorAnd is what identifies the AND condition in a composite
	OperatorAnd = "and"
	// OperatorOr is what identifies the OR condition in a composite
	OperatorOr = "or"
)

// defaultComparators is a map of all the default comparators that
// a new engine should include
var defaultComparators = map[string]Comparator{
	"eq":       equal,
	"ne":       notEqual,
	"gt":       greaterThan,
	"gte":      greaterThanEqual,
	"lt":       lessThan,
	"lte":      lessThanEqual,
	"contains": contains,
	"oneof":    oneOf,
}

// Rule is a our smallest unit of measure, each rule will be
// evaluated separately. The comparator is the logical operation to be
// performed, the path is the path into a map, delimited by '.', and
// the value is the value that we expect to match the value at the
// path
type Rule struct {
	Comparator string      `json:"comparator"`
	Path       string      `json:"path"`
	Value      interface{} `json:"value"`
}

// Composite is a group of rules that are joined by a logical operator
// AND or OR. If the operator is AND all of the rules must be true,
// if the operator is OR, one of the rules must be true.
type Composite struct {
	Operator string `json:"operator"`
	Rules    []Rule `json:"rules"`
}

// Engine is a group of composites. All of the composites must be
// true for the engine's evaluate function to return true.
type Engine struct {
	Composites  []Composite `json:"composites"`
	comparators map[string]Comparator
}

// NewEngine will create a new engine with the default comparators
func NewEngine() Engine {
	e := Engine{
		comparators: defaultComparators,
	}
	return e
}

// NewJSONEngine will create a new engine from it's JSON representation
func NewJSONEngine(raw json.RawMessage) (Engine, error) {
	var e Engine
	err := json.Unmarshal(raw, &e)
	if err != nil {
		return Engine{}, err
	}
	e.comparators = defaultComparators
	return e, nil
}

// AddComparator will add a new comparator that can be used in the
// engine's evaluation
func (e Engine) AddComparator(name string, c Comparator) Engine {
	e.comparators[name] = c
	return e
}

// Evaluate will ensure all of the composites in the engine are true
func (e Engine) Evaluate(props map[string]interface{}) bool {
	for _, c := range e.Composites {
		res := c.evaluate(props, e.comparators)
		if res == false {
			return false
		}
	}
	return true
}

// Evaluate will ensure all either all of the rules are true, if given
// the AND operator, or that one of the rules is true if given the OR
// operator.
func (c Composite) evaluate(props map[string]interface{}, comps map[string]Comparator) bool {
	switch c.Operator {
	case OperatorAnd:
		for _, r := range c.Rules {
			res := r.evaluate(props, comps)
			if res == false {
				return false
			}
		}
		return true
	case OperatorOr:
		for _, r := range c.Rules {
			res := r.evaluate(props, comps)
			if res == true {
				return true
			}
		}
		return false
	}

	return false
}

// Evaluate will return true if the rule is true, false otherwise
func (r Rule) evaluate(props map[string]interface{}, comps map[string]Comparator) bool {
	// Make sure we can get a value from the props
	inter := pluck(props, r.Path)
	if inter == nil {
		return false
	}

	// This is an important step, we need to get numbers to their most
	// precise type, because that is how the mapper works
	var val interface{}
	switch inter.(type) {
	case uint:
		val = float64(inter.(uint))
	case uint8:
		val = float64(inter.(uint8))
	case uint16:
		val = float64(inter.(uint16))
	case uint32:
		val = float64(inter.(uint32))
	case uint64:
		val = float64(inter.(uint64))
	case int:
		val = float64(inter.(int))
	case int8:
		val = float64(inter.(int8))
	case int16:
		val = float64(inter.(int16))
	case int32:
		val = float64(inter.(int32))
	case int64:
		val = float64(inter.(int64))
	case float32:
		val = float64(inter.(float32))
	default:
		val = inter
	}

	comp, ok := comps[r.Comparator]
	if !ok {
		return false
	}

	fmt.Println(reflect.TypeOf(val), reflect.TypeOf(r.Value))
	return comp(val, r.Value)
}
