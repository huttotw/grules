package grules

import (
	"encoding/json"
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
	"eq":        equal,
	"neq":       notEqual,
	"gt":        greaterThan,
	"gte":       greaterThanEqual,
	"lt":        lessThan,
	"lte":       lessThanEqual,
	"contains":  contains,
	"ncontains": notContains,
	"oneof":     oneOf,
	"noneof":    noneOf,
	"regex":     regex,
}

// Rule is a our smallest unit of measure, each rule will be
// evaluated separately. The comparator is the logical operation to be
// performed, the path is the path into a map, delimited by '.', and
// the value is the value that we expect to match the value at the
// path
type rule struct {
	Comparator string      `json:"comparator"`
	Path       string      `json:"path"`
	Value      interface{} `json:"value"`
}

// MarshalJSON is important because it will put maps back into arrays, we used maps
// to speed up one of
func (r *rule) MarshalJSON() ([]byte, error) {
	type unmappedRule struct {
		Comparator string      `json:"comparator"`
		Path       string      `json:"path"`
		Value      interface{} `json:"value"`
	}

	switch t := r.Value.(type) {
	case map[interface{}]struct{}:
		var s []interface{}
		for k := range t {
			s = append(s, k)
		}
		r.Value = s
	}

	umr := unmappedRule{
		Comparator: r.Comparator,
		Path:       r.Path,
		Value:      r.Value,
	}

	return json.Marshal(umr)
}

// UnmarshalJSON is important because it will convert arrays in a rule set to a map
// to provide faster lookups
func (r *rule) UnmarshalJSON(data []byte) error {
	type mapRule struct {
		Comparator string      `json:"comparator"`
		Path       string      `json:"path"`
		Value      interface{} `json:"value"`
	}

	var mr mapRule
	err := json.Unmarshal(data, &mr)
	if err != nil {
		return err
	}

	switch t := mr.Value.(type) {
	case []interface{}:
		var m = make(map[interface{}]struct{})
		for _, v := range t {
			m[v] = struct{}{}
		}

		mr.Value = m
	}

	*r = rule{
		Comparator: mr.Comparator,
		Path:       mr.Path,
		Value:      mr.Value,
	}

	return nil
}

// Composite is a group of rules that are joined by a logical operator
// AND or OR. If the operator is AND all of the rules must be true,
// if the operator is OR, one of the rules must be true.

type composite struct {
	Operator   string      `json:"operator"`
	Rules      []Rule      `json:"rules"`
	Composites []Composite `json:"composites"`
}

// Engine is a group of composites. All of the composites must be
// true for the engine's evaluate function to return true.
type Engine struct {
	Composites  []composite `json:"composites"`
	comparators map[string]Comparator
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
func (c composite) evaluate(props map[string]interface{}, comps map[string]Comparator) bool {
	switch c.Operator {
	case OperatorAnd:
		for _, r := range c.Rules {
			res := r.evaluate(props, comps)
			if res == false {
				return false
			}
		}
		for _, cc := range c.Composites {
			res := cc.evaluate(props, comps)
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
		for _, cc := range c.Composites {
			res := cc.evaluate(props, comps)
			if res == true {
				return true
			}
		}
		return false
	}

	return false
}

// Evaluate will return true if the rule is true, false otherwise
func (r rule) evaluate(props map[string]interface{}, comps map[string]Comparator) bool {
	// Make sure we can get a value from the props
	val := pluck(props, r.Path)
	if val == nil {
		return false
	}

	comp, ok := comps[r.Comparator]
	if !ok {
		return false
	}

	return comp(val, r.Value)
}
