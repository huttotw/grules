package grules

import (
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
		testCase{args: []interface{}{1, 1}, expected: true},
		testCase{args: []interface{}{1, 0}, expected: false},
		testCase{args: []interface{}{1.1, 1.1}, expected: true},
		testCase{args: []interface{}{1.1, 0.1}, expected: false},
	}

	for i, c := range cases {
		res := equal(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func TestNotEqual(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{"a", "a"}, expected: false},
		testCase{args: []interface{}{"a", "b"}, expected: true},
		testCase{args: []interface{}{1, 1}, expected: false},
		testCase{args: []interface{}{1, 0}, expected: true},
		testCase{args: []interface{}{1.1, 1.1}, expected: false},
		testCase{args: []interface{}{1.1, 0.1}, expected: true},
	}

	for i, c := range cases {
		res := notEqual(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func TestLessThan(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{"a", "a"}, expected: false},
		testCase{args: []interface{}{"a", "b"}, expected: true},
		testCase{args: []interface{}{1, 1}, expected: false},
		testCase{args: []interface{}{0, 1}, expected: true},
		testCase{args: []interface{}{1.1, 1.1}, expected: false},
		testCase{args: []interface{}{1.1, 1.2}, expected: true},
	}

	for i, c := range cases {
		res := lessThan(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func TestLessThanEqual(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{"a", "a"}, expected: true},
		testCase{args: []interface{}{"a", "b"}, expected: true},
		testCase{args: []interface{}{"c", "b"}, expected: false},
		testCase{args: []interface{}{1, 1}, expected: true},
		testCase{args: []interface{}{0, 1}, expected: true},
		testCase{args: []interface{}{1, 0}, expected: false},
		testCase{args: []interface{}{1.1, 1.1}, expected: true},
		testCase{args: []interface{}{1.1, 1.2}, expected: true},
		testCase{args: []interface{}{1.2, 1.1}, expected: false},
	}

	for i, c := range cases {
		res := lessThanEqual(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func TestGreaterThan(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{"a", "a"}, expected: false},
		testCase{args: []interface{}{"b", "a"}, expected: true},
		testCase{args: []interface{}{1, 1}, expected: false},
		testCase{args: []interface{}{1, 0}, expected: true},
		testCase{args: []interface{}{1.1, 1.1}, expected: false},
		testCase{args: []interface{}{1.2, 1.1}, expected: true},
	}

	for i, c := range cases {
		res := greaterThan(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func TestGreaterThanEqual(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{"a", "a"}, expected: true},
		testCase{args: []interface{}{"a", "b"}, expected: false},
		testCase{args: []interface{}{"c", "b"}, expected: true},
		testCase{args: []interface{}{1, 1}, expected: true},
		testCase{args: []interface{}{0, 1}, expected: false},
		testCase{args: []interface{}{1, 0}, expected: true},
		testCase{args: []interface{}{1.1, 1.1}, expected: true},
		testCase{args: []interface{}{1.1, 1.2}, expected: false},
		testCase{args: []interface{}{1.2, 1.1}, expected: true},
	}

	for i, c := range cases {
		res := greaterThanEqual(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}

func TestContains(t *testing.T) {
	cases := []testCase{
		testCase{args: []interface{}{[]string{"a", "b"}, "a"}, expected: true},
		testCase{args: []interface{}{[]string{"a", "b"}, "c"}, expected: false},
		testCase{args: []interface{}{[]string{"a", "b"}, 1}, expected: false},
		testCase{args: []interface{}{[]int{1, 2}, 1}, expected: true},
		testCase{args: []interface{}{[]int{1, 2}, 3}, expected: false},
		testCase{args: []interface{}{[]float64{1.01, 1.02}, 1.01}, expected: true},
	}

	for i, c := range cases {
		res := contains(c.args[0], c.args[1])
		if res != c.expected {
			t.Fatalf("expected case %d to be %v, got %v", i, c.expected, res)
		}
	}
}
