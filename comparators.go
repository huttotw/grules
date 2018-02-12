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
	// We need to check the types here because, lessThanEqual will
	// return false if the types do not match.
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}
	return !lessThanEqual(a, b)
}

// greaterThanEqual will return true if a >= b
func greaterThanEqual(a, b interface{}) bool {
	// We need to check the types here because, lessThan will
	// return false if the types do not match.
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}
	return !lessThan(a, b)
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
	case reflect.Int:
		return containsInt(a, b)
	case reflect.Int8:
		return containsInt8(a, b)
	case reflect.Int16:
		return containsInt16(a, b)
	case reflect.Int32:
		return containsInt32(a, b)
	case reflect.Int64:
		return containsInt64(a, b)
	case reflect.Uint:
		return containsUint(a, b)
	case reflect.Uint8:
		return containsUint8(a, b)
	case reflect.Uint16:
		return containsUint16(a, b)
	case reflect.Uint32:
		return containsUint32(a, b)
	case reflect.Uint64:
		return containsUint64(a, b)
	case reflect.Float32:
		return containsFloat32(a, b)
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

func containsInt(a, b interface{}) bool {
	as, ok := a.([]interface{})
	if !ok {
		return false
	}
	for _, elem := range as {
		if val, ok := elem.(int); ok && val == b.(int) {
			return true
		}
	}
	return false
}

func containsInt8(a, b interface{}) bool {
	as, ok := a.([]interface{})
	if !ok {
		return false
	}
	for _, elem := range as {
		if val, ok := elem.(int8); ok && val == b.(int8) {
			return true
		}
	}
	return false
}

func containsInt16(a, b interface{}) bool {
	as, ok := a.([]interface{})
	if !ok {
		return false
	}
	for _, elem := range as {
		if val, ok := elem.(int16); ok && val == b.(int16) {
			return true
		}
	}
	return false
}

func containsInt32(a, b interface{}) bool {
	as, ok := a.([]interface{})
	if !ok {
		return false
	}
	for _, elem := range as {
		if val, ok := elem.(int32); ok && val == b.(int32) {
			return true
		}
	}
	return false
}

func containsInt64(a, b interface{}) bool {
	as, ok := a.([]interface{})
	if !ok {
		return false
	}
	for _, elem := range as {
		if val, ok := elem.(int64); ok && val == b.(int64) {
			return true
		}
	}
	return false
}

func containsUint(a, b interface{}) bool {
	as, ok := a.([]interface{})
	if !ok {
		return false
	}
	for _, elem := range as {
		if val, ok := elem.(uint); ok && val == b.(uint) {
			return true
		}
	}
	return false
}

func containsUint8(a, b interface{}) bool {
	as, ok := a.([]interface{})
	if !ok {
		return false
	}
	for _, elem := range as {
		if val, ok := elem.(uint8); ok && val == b.(uint8) {
			return true
		}
	}
	return false
}

func containsUint16(a, b interface{}) bool {
	as, ok := a.([]interface{})
	if !ok {
		return false
	}
	for _, elem := range as {
		if val, ok := elem.(uint16); ok && val == b.(uint16) {
			return true
		}
	}
	return false
}

func containsUint32(a, b interface{}) bool {
	as, ok := a.([]interface{})
	if !ok {
		return false
	}
	for _, elem := range as {
		if val, ok := elem.(uint32); ok && val == b.(uint32) {
			return true
		}
	}
	return false
}

func containsUint64(a, b interface{}) bool {
	as, ok := a.([]interface{})
	if !ok {
		return false
	}
	for _, elem := range as {
		if val, ok := elem.(uint64); ok && val == b.(uint64) {
			return true
		}
	}
	return false
}

func containsFloat32(a, b interface{}) bool {
	as, ok := a.([]interface{})
	if !ok {
		return false
	}
	for _, elem := range as {
		if val, ok := elem.(float32); ok && val == b.(float32) {
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

// oneOf will return true if b contains a
func oneOf(a, b interface{}) bool {
	return contains(b, a)
}
