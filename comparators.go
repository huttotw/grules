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
	case reflect.Int:
		return a.(int) < b.(int)
	case reflect.Int8:
		return a.(int8) < b.(int8)
	case reflect.Int16:
		return a.(int16) < b.(int16)
	case reflect.Int32:
		return a.(int32) < b.(int32)
	case reflect.Int64:
		return a.(int64) < b.(int64)
	case reflect.Uint:
		return a.(uint) < b.(uint)
	case reflect.Uint8:
		return a.(uint8) < b.(uint8)
	case reflect.Uint16:
		return a.(uint16) < b.(uint16)
	case reflect.Uint32:
		return a.(uint32) < b.(uint32)
	case reflect.Uint64:
		return a.(uint64) < b.(uint64)
	case reflect.Float32:
		return a.(float32) < b.(float32)
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
	case reflect.Int:
		return a.(int) <= b.(int)
	case reflect.Int8:
		return a.(int8) <= b.(int8)
	case reflect.Int16:
		return a.(int16) <= b.(int16)
	case reflect.Int32:
		return a.(int32) <= b.(int32)
	case reflect.Int64:
		return a.(int64) <= b.(int64)
	case reflect.Uint:
		return a.(uint) <= b.(uint)
	case reflect.Uint8:
		return a.(uint8) <= b.(uint8)
	case reflect.Uint16:
		return a.(uint16) <= b.(uint16)
	case reflect.Uint32:
		return a.(uint32) <= b.(uint32)
	case reflect.Uint64:
		return a.(uint64) <= b.(uint64)
	case reflect.Float32:
		return a.(float32) <= b.(float32)
	case reflect.Float64:
		return a.(float64) <= b.(float64)
	}

	return false
}

// greaterThan will return true if a > b
func greaterThan(a, b interface{}) bool {
	return !lessThanEqual(a, b)
}

// greaterThanEqual will return true if a >= b
func greaterThanEqual(a, b interface{}) bool {
	return !lessThan(a, b)
}

// contains will return true if a is contained in b
func contains(a, b interface{}) bool {
	t1 := reflect.TypeOf(a)

	if t1.Kind() != reflect.Slice {
		return false
	}

	s := reflect.ValueOf(a)
	n := s.Len()
	for i := 0; i < n; i++ {
		if s.Index(i).Interface() == b {
			return true
		}
	}

	return false
}
