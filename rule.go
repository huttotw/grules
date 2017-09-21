package grules

import (
	"reflect"
)

const (
	// OperatorAnd is what identifies the AND condition in a composite
	OperatorAnd = "and"
	// OperatorOr is what identifies the OR condition in a composite
	OperatorOr = "or"
)

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
		comparators: map[string]Comparator{
			"eq":       equal,
			"ne":       notEqual,
			"gt":       greaterThan,
			"gte":      greaterThanEqual,
			"lt":       lessThan,
			"lte":      lessThanEqual,
			"contains": contains,
		},
	}
	return e
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
	val := pluck(props, r.Path)
	if val == nil {
		return false
	}

	// Both values must be comparable
	t1 := reflect.TypeOf(r.Value)
	t2 := reflect.TypeOf(val)
	if !t1.Comparable() || !t2.Comparable() {
		return false
	}

	comp, ok := comps[r.Comparator]
	if !ok {
		return false
	}

	return comp(val, r.Value)
}
