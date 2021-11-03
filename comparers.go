package grules

import (
	"regexp"
	"strings"
)

// Compare is a function that should evaluate two values and return
// the true if the comparison is true, or false if the comparison is
// false
type Compare func(a, b interface{}) bool

// defaultComparers is a map of all the default comparators that
// a new engine should include
var defaultComparers = map[string]Compare{
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

// equal will return true if a == b
func equal(a, b interface{}) bool {
	return a == b
}

// notEqual will return true if a != b
func notEqual(a, b interface{}) bool {
	return !equal(a, b)
}

// lessThan will return true if a < b
func lessThan(a, b interface{}) bool {
	switch a.(type) {
	case string:
		at, ok := a.(string)
		if !ok {
			return false
		}
		bt, ok := b.(string)
		if !ok {
			return false
		}
		return at < bt
	case float64:
		at, ok := a.(float64)
		if !ok {
			return false
		}
		bt, ok := b.(float64)
		if !ok {
			return false
		}
		return at < bt
	default:
		return false
	}
}

func lessThanEqual(a, b interface{}) bool {
	switch a.(type) {
	case string:
		at, ok := a.(string)
		if !ok {
			return false
		}
		bt, ok := b.(string)
		if !ok {
			return false
		}
		return at <= bt
	case float64:
		at, ok := a.(float64)
		if !ok {
			return false
		}
		bt, ok := b.(float64)
		if !ok {
			return false
		}
		return at <= bt
	default:
		return false
	}
}

// greaterThan will return true if a > b
func greaterThan(a, b interface{}) bool {
	switch a.(type) {
	case string:
		at, ok := a.(string)
		if !ok {
			return false
		}
		bt, ok := b.(string)
		if !ok {
			return false
		}
		return at > bt
	case float64:
		at, ok := a.(float64)
		if !ok {
			return false
		}
		bt, ok := b.(float64)
		if !ok {
			return false
		}
		return at > bt
	default:
		return false
	}
}

// greaterThanEqual will return true if a >= b
func greaterThanEqual(a, b interface{}) bool {
	switch a.(type) {
	case string:
		at, ok := a.(string)
		if !ok {
			return false
		}
		bt, ok := b.(string)
		if !ok {
			return false
		}
		return at >= bt
	case float64:
		at, ok := a.(float64)
		if !ok {
			return false
		}
		bt, ok := b.(float64)
		if !ok {
			return false
		}
		return at >= bt
	default:
		return false
	}
}

func regex(a, b interface{}) bool {
	switch a.(type) {
	case string:
		at, ok := a.(string)
		if !ok {
			return false
		}
		bt, ok := b.(string)
		if !ok {
			return false
		}

		r, err := regexp.Compile(bt)
		if err != nil {
			return false
		}

		return r.MatchString(at)
	default:
		return false
	}
}

// contains will return true if a contains b. a can be a slice
// or a string.  If you need b to be a slice consider using oneOf
func contains(a, b interface{}) bool {
	switch bt := b.(type) {
	case string:
		switch at := a.(type) {
		case []interface{}:
			for _, v := range at {
				if elem, ok := v.(string); ok && elem == bt {
					return true
				}
			}
			return false
		case []string:
			for _, v := range at {
				if v == bt {
					return true
				}
			}
			return false
		case string:
			return strings.Contains(a.(string), b.(string))
		default:
			return false
		}
	case float64:
		switch at := a.(type) {
		case []interface{}:
			for _, v := range at {
				if elem, ok := v.(float64); ok && elem == bt {
					return true
				}
			}
			return false
		case []float64:
			for _, v := range at {
				if v == bt {
					return true
				}
			}
		default:
			return false
		}
	default:
		return false
	}

	return false
}

// notContains will return true if the b is not contained a. This will also return
// true if a is a slice of different types than b. It will return false if a
// is not a slice or a string.
func notContains(a, b interface{}) bool {
	switch bt := b.(type) {
	case string:
		switch at := a.(type) {
		case []interface{}:
			for _, v := range at {
				if elem, ok := v.(string); ok && elem == bt {
					return false
				}
			}
			return true
		case []string:
			for _, v := range at {
				if v == bt {
					return false
				}
			}
			return true
		case string:
			return !strings.Contains(a.(string), b.(string))
		default:
			return false
		}
	case float64:
		switch at := a.(type) {
		case []interface{}:
			for _, v := range at {
				if elem, ok := v.(float64); ok && elem == bt {
					return false
				}
			}
			return true
		case []float64:
			for _, v := range at {
				if v == bt {
					return false
				}
			}
			return true
		default:
			return false
		}
	default:
		return false
	}
}

// oneOf will return true if b (slice) contains a
func oneOf(a, b interface{}) bool {
	m, ok := b.(map[interface{}]struct{})
	if !ok {
		return false
	}

	_, found := m[a]

	return found
}

// noneOf will return true if b (slice) does not contain a
func noneOf(a, b interface{}) bool {
	m, ok := b.(map[interface{}]struct{})
	if !ok {
		return false
	}

	_, found := m[a]

	return found
}
