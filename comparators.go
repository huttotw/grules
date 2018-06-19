package grules

import (
	"reflect"
)

// Comparator is a function that should evaluate two values and return
// the true if the comparison is true, or false if the comparison is
// false
type Comparator func(a, b interface{}) bool

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
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	// Make sure the types are the same
	if ta != tb {
		return false
	}

	// We have already checked that each argument is the same type
	// so it is safe to only check the first argument
	switch ta.Kind() {
	case reflect.String:
		return a.(string) < b.(string)
	case reflect.Float64:
		return a.(float64) < b.(float64)
	}

	return false
}

// lessThanEqual will return true if a <= b
func lessThanEqual(a, b interface{}) bool {
	// If the values are equal, no more work necessary
	if a == b {
		return true
	}

	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	// Make sure the types are the same
	if ta != tb {
		return false
	}

	// We have already checked that each argument is the same type
	// so it is safe to only check the first argument
	switch ta.Kind() {
	case reflect.String:
		return a.(string) <= b.(string)
	case reflect.Float64:
		return a.(float64) <= b.(float64)
	}

	return false
}

// greaterThan will return true if a > b
func greaterThan(a, b interface{}) bool {
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	// Make sure the types are the same
	if ta != tb {
		return false
	}

	// We have already checked that each argument is the same type
	// so it is safe to only check the first argument
	switch ta.Kind() {
	case reflect.String:
		return a.(string) > b.(string)
	case reflect.Float64:
		return a.(float64) > b.(float64)
	}

	return false
}

// greaterThanEqual will return true if a >= b
func greaterThanEqual(a, b interface{}) bool {
	// If the values are equal, no more work necessary
	if a == b {
		return true
	}

	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	// Make sure the types are the same
	if ta != tb {
		return false
	}

	// We have already checked that each argument is the same type
	// so it is safe to only check the first argument
	switch ta.Kind() {
	case reflect.String:
		return a.(string) >= b.(string)
	case reflect.Float64:
		return a.(float64) >= b.(float64)
	}

	return false
}

// contains will return true if a contains b. We assume
// that the first interface is a slice. If you need b to be a slice
// consider using oneOf
func contains(a, b interface{}) bool {
	t1 := reflect.TypeOf(a)
	t2 := reflect.TypeOf(b)

	if t1.Kind() != reflect.Slice {
		return false
	}

	switch t2.Kind() {
	case reflect.String:
		return containsString(a, b)
	case reflect.Float64:
		return containsFloat64(a, b)
	default:
		return false
	}
}

func containsString(a, b interface{}) bool {
	as, ok := a.([]interface{})
	if !ok {
		return false
	}
	for _, elem := range as {
		if val, ok := elem.(string); ok && val == b.(string) {
			return true
		}
	}
	return false
}

func containsFloat64(a, b interface{}) bool {
	as, ok := a.([]interface{})
	if !ok {
		return false
	}
	for _, elem := range as {
		if val, ok := elem.(float64); ok && val == b.(float64) {
			return true
		}
	}
	return false
}

// notContains will return true if the b is not contained a. This will also return
// true if a is a slice of different types than b. It will return false if a
// is not a slice.
func notContains(a, b interface{}) bool {
	t1 := reflect.TypeOf(a)
	t2 := reflect.TypeOf(b)

	if t1.Kind() != reflect.Slice {
		return false
	}

	switch t2.Kind() {
	case reflect.String:
		return notContainsString(a, b)
	case reflect.Float64:
		return notContainsFloat64(a, b)
	default:
		return false
	}
}

func notContainsString(a, b interface{}) bool {
	as, ok := a.([]interface{})
	if !ok {
		return false
	}
	for _, elem := range as {
		if val, ok := elem.(string); ok && val == b.(string) {
			return false
		}
	}
	return true
}

func notContainsFloat64(a, b interface{}) bool {
	as, ok := a.([]interface{})
	if !ok {
		return false
	}

	for _, elem := range as {
		if val, ok := elem.(float64); ok && val == b.(float64) {
			return false
		}
	}
	return true
}

// oneOf will return true if b contains a
func oneOf(a, b interface{}) bool {
	return contains(b, a)
}

// noneOf will return true if b does not contain a
func noneOf(a, b interface{}) bool {
	return notContains(b, a)
}
