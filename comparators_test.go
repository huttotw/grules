package grules

import (
	"fmt"
	"testing"
)

type testCase struct {
	args     []interface{}
	expected bool
}

func TestEqual(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{"a", "a"}, expected: true},
		testCase{args: []interface{}{"a", "b"}, expected: false},
		testCase{args: []interface{}{float64(1), float64(1)}, expected: true},
		testCase{args: []interface{}{float64(1), float64(0)}, expected: false},
		testCase{args: []interface{}{float64(1.1), float64(1.1)}, expected: true},
		testCase{args: []interface{}{float64(1.1), float64(0.1)}, expected: false},
	}

	for i, c := range cases {
		res := equal(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func BenchmarkEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		equal("benchmark", "benchmark")
	}
}

func TestNotEqual(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{"a", "a"}, expected: false},
		testCase{args: []interface{}{"a", "b"}, expected: true},
		testCase{args: []interface{}{float64(1), float64(1)}, expected: false},
		testCase{args: []interface{}{float64(1), float64(0)}, expected: true},
		testCase{args: []interface{}{float64(1.1), float64(1.1)}, expected: false},
		testCase{args: []interface{}{float64(1.1), float64(0.1)}, expected: true},
	}

	for i, c := range cases {
		res := notEqual(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func BenchmarkNotEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		notEqual("benchmark", "not-benchmark")
	}
}

func TestLessThan(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{"a", "a"}, expected: false},
		testCase{args: []interface{}{"a", "b"}, expected: true},
		testCase{args: []interface{}{float64(1), float64(1)}, expected: false},
		testCase{args: []interface{}{float64(0), float64(1)}, expected: true},
		testCase{args: []interface{}{float64(1.1), float64(1.1)}, expected: false},
		testCase{args: []interface{}{float64(1.1), float64(1.2)}, expected: true},
	}

	for i, c := range cases {
		res := lessThan(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func BenchmarkLessThan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lessThan(0, 1)
	}
}

func TestLessThanEqual(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{"a", "a"}, expected: true},
		testCase{args: []interface{}{"a", "b"}, expected: true},
		testCase{args: []interface{}{"c", "b"}, expected: false},
		testCase{args: []interface{}{float64(1), float64(1)}, expected: true},
		testCase{args: []interface{}{float64(0), float64(1)}, expected: true},
		testCase{args: []interface{}{float64(1), float64(0)}, expected: false},
		testCase{args: []interface{}{float64(1.1), float64(1.1)}, expected: true},
		testCase{args: []interface{}{float64(1.1), float64(1.2)}, expected: true},
		testCase{args: []interface{}{float64(1.2), float64(1.1)}, expected: false},
	}

	for i, c := range cases {
		res := lessThanEqual(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func BenchmarkLessThanEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lessThanEqual(0, 0)
	}
}

func TestGreaterThan(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{"a", "a"}, expected: false},
		testCase{args: []interface{}{"b", "a"}, expected: true},
		testCase{args: []interface{}{float64(1), float64(1)}, expected: false},
		testCase{args: []interface{}{float64(1), float64(0)}, expected: true},
		testCase{args: []interface{}{float64(1.1), float64(1.1)}, expected: false},
		testCase{args: []interface{}{float64(1.2), float64(1.1)}, expected: true},
	}

	for i, c := range cases {
		res := greaterThan(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func BenchmarkGreaterThan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		greaterThan(1, 0)
	}
}

func TestGreaterThanEqual(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{"a", "a"}, expected: true},
		testCase{args: []interface{}{"a", "b"}, expected: false},
		testCase{args: []interface{}{"c", "b"}, expected: true},
		testCase{args: []interface{}{float64(1), float64(1)}, expected: true},
		testCase{args: []interface{}{float64(0), float64(1)}, expected: false},
		testCase{args: []interface{}{float64(1), float64(0)}, expected: true},
		testCase{args: []interface{}{float64(1.1), float64(1.1)}, expected: true},
		testCase{args: []interface{}{float64(1.1), float64(1.2)}, expected: false},
		testCase{args: []interface{}{float64(1.2), float64(1.1)}, expected: true},
	}

	for i, c := range cases {
		res := greaterThanEqual(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func BenchmarkGreaterThanEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		greaterThanEqual(0, 0)
	}
}

func TestRegex(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{"a", "a"}, expected: true},
		testCase{args: []interface{}{"a", "[ab]"}, expected: true},
		testCase{args: []interface{}{"c", "[ab]"}, expected: false},
		testCase{args: []interface{}{"abc", "c$"}, expected: true},
		testCase{args: []interface{}{float64(1), float64(1)}, expected: false},
	}

	for i, c := range cases {
		res := regex(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func BenchmarkRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		regex("a", "a")
	}
}

func BenchmarkRegexPhone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		regex("+1(555) 555-5555", "\\+\\d\\(\\d+\\) \\d+-\\d+")
	}
}

func TestContains(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{[]interface{}{"a", "b"}, "a"}, expected: true},
		testCase{args: []interface{}{[]interface{}{"a", "b"}, "c"}, expected: false},
		testCase{args: []interface{}{[]interface{}{"a", "b"}, float64(1)}, expected: false},
		testCase{args: []interface{}{[]interface{}{float64(1), float64(2)}, float64(1)}, expected: true},
		testCase{args: []interface{}{[]interface{}{float64(1), float64(2)}, float64(3)}, expected: false},
		testCase{args: []interface{}{[]interface{}{float64(1.01), float64(1.02)}, float64(1.01)}, expected: true},
		testCase{args: []interface{}{"abc", "bc"}, expected: true},
		testCase{args: []interface{}{"abc", "de"}, expected: false},
	}

	for i, c := range cases {
		res := contains(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func BenchmarkContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		contains([]string{"1", "2"}, "1")
	}
}

func BenchmarkContainsLong50000(b *testing.B) {
	var list []interface{}

	// Simulate a list of postal codes
	for i := 0; i < 50000; i++ {
		list = append(list, fmt.Sprintf("%d", i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		contains(list, "49999")
	}
}

func TestNotContains(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{[]interface{}{}, "a"}, expected: true},
		testCase{args: []interface{}{[]interface{}{"a", "b"}, "a"}, expected: false},
		testCase{args: []interface{}{[]interface{}{"a", "b"}, "c"}, expected: true},
		testCase{args: []interface{}{[]interface{}{"a", "b"}, float64(1)}, expected: true},
		testCase{args: []interface{}{[]interface{}{float64(1), float64(2)}, float64(1)}, expected: false},
		testCase{args: []interface{}{[]interface{}{float64(1), float64(2)}, float64(3)}, expected: true},
		testCase{args: []interface{}{[]interface{}{float64(1.01), float64(1.02)}, float64(1.01)}, expected: false},
		testCase{args: []interface{}{"abc", "bc"}, expected: false},
		testCase{args: []interface{}{"abc", "de"}, expected: true},
	}

	for i, c := range cases {
		res := notContains(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func BenchmarkNotContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		contains([]string{"1", "2"}, "3")
	}
}

func BenchmarkNotContainsLong50000(b *testing.B) {
	var list []interface{}

	// Simulate a list of postal codes
	for i := 0; i < 50000; i++ {
		list = append(list, fmt.Sprintf("%d", i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		contains(list, "50000")
	}
}

func BenchmarkOneOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oneOf("2", []string{"1", "2"})
	}
}

func TestOneOf(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{"a", map[interface{}]struct{}{"a": struct{}{}, "b": struct{}{}}}, expected: true},
		testCase{args: []interface{}{"c", map[interface{}]struct{}{"a": struct{}{}, "b": struct{}{}}}, expected: false},
		testCase{args: []interface{}{float64(1), map[interface{}]struct{}{"a": struct{}{}, "b": struct{}{}}}, expected: false},
		testCase{args: []interface{}{float64(1), map[interface{}]struct{}{float64(1): struct{}{}, float64(2): struct{}{}}}, expected: true},
		testCase{args: []interface{}{float64(3), map[interface{}]struct{}{float64(1): struct{}{}, float64(2): struct{}{}}}, expected: false},
		testCase{args: []interface{}{float64(1.01), map[interface{}]struct{}{1.01: struct{}{}, 1.02: struct{}{}}}, expected: true},
	}
	for i, c := range cases {
		res := oneOf(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func BenchmarkNoneOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noneOf("2", []string{"1", "2"})
	}
}

func TestNoneOf(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{"a", map[interface{}]struct{}{"a": struct{}{}, "b": struct{}{}}}, expected: false},
		testCase{args: []interface{}{"c", map[interface{}]struct{}{"a": struct{}{}, "b": struct{}{}}}, expected: true},
		testCase{args: []interface{}{float64(1), map[interface{}]struct{}{"a": struct{}{}, "b": struct{}{}}}, expected: true},
		testCase{args: []interface{}{float64(1), map[interface{}]struct{}{float64(1): struct{}{}, float64(2): struct{}{}}}, expected: false},
		testCase{args: []interface{}{float64(3), map[interface{}]struct{}{float64(1): struct{}{}, float64(2): struct{}{}}}, expected: true},
		testCase{args: []interface{}{float64(1.01), map[interface{}]struct{}{1.01: struct{}{}, 1.02: struct{}{}}}, expected: false},
		testCase{args: []interface{}{float64(1.03), map[interface{}]struct{}{1.01: struct{}{}, 1.02: struct{}{}}}, expected: true},
	}

	for i, c := range cases {
		res := noneOf(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}
