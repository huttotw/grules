[![CircleCI](https://circleci.com/gh/huttotw/grules/tree/master.svg?style=svg)](https://circleci.com/gh/huttotw/grules/tree/master)

# Introduction
This package was created with inspiration from Thomas' [go-ruler](https://github.com/hopkinsth/go-ruler) to run a simple set of rules against an entity.

This version includes a couple more features including, AND and OR composites and the ability to add custom comparators.

**Note**: This package only compares two types: `string` and `float64`, this plays nicely with `encoding/json`.

# Example
```go
// Create a new instance of an engine with some default comparators
e := NewEngine()

// Add a new, custom comparator
e = e.AddComparator("always-false", func(a, b interface{}) bool {
    return false
})

// Create composites, with rules for the engine to evaluate
e.Composites = []Composite{
    Composite{
        Operator: OperatorOr,
        Rules: []Rule{
            Rule{
                Comparator: "always-false",
                Path: "user.name",
                Value: "Trevor",
            },
            Rule{
                Comparator: "eq",
                Path: "user.name",
                Value: "Trevor",
            },
        },
    },
}

// Give some properties, this map can be deeper and supports interfaces
props := map[string]interface{}{
    "user": map[string]interface{}{
        "name": "Trevor",
    }
}

// Run the engine on the props
res := e.Evaluate(props)
// res == true
```

# Comparators
* `eq` will return true if `a == b`
* `neq` will return true if `a != b`
* `lt` will return true if `a < b`
* `lte` will return true if `a <= b`
* `gt` will return true if `a > b`
* `gte` will return true if `a >= b`
* `contains` will return true if `a` contains `b`
* `oneof` will return true if `a` is one of `b`

`contains` is different than `oneof` in that `contains` expects the first argument to be a slice, and `oneof` expects the second argument to be a slice.

# Benchmarks

|Benchmark|N|Speed|Used|Allocs|
|---------|----------|-----|------|------|
|BenchmarkEqual-8|1000000000|7.26 ns/op|0 B/op|0 allocs/op|
BenchmarkNotEqual-8|1000000000|4.95 ns/op|0 B/op|0 allocs/op|
BenchmarkLessThan-8|500000000|10.9 ns/op|0 B/op|0 allocs/op|
BenchmarkLessThanEqual-8|1000000000|7.79 ns/op|0 B/op|0 allocs/op|
BenchmarkGreaterThan-8|200000000|18.1 ns/op|0 B/op|0 allocs/op|
BenchmarkGreaterThanEqual-8|300000000|13.9 ns/op|0 B/op|0 allocs/op|
BenchmarkContains-8|50000000|73.0 ns/op|64 B/op|2 allocs/op|
BenchmarkContainsLong50000-8|100000000|55.6 ns/op|32 B/op|1 allocs/op|
BenchmarkNotContains-8|50000000|75.1 ns/op|64 B/op|2 allocs/op|
BenchmarkNotContainsLong50000-8|100000000|56.2 ns/op|32 B/op|1 allocs/op|
BenchmarkOneOf|50000000|70.5 ns/op|64 B/op|2 allocs/op|
BenchmarkNoneOf|100000000|71.7 ns/op|64 B/op|2 allocs/op|
BenchmarkPluckShallow-8|100000000|60.2 ns/op|16 B/op|1 allocs/op|
BenchmarkPluckDeep-8|20000000|242 ns/op|112 B/op|1 allocs/op|

To run benchmarks:
```
go test -run none -bench . -benchtime 3s -benchmem
```

# License

Copyright &copy; 2018 Trevor Hutto

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this work except in compliance with the License. You may obtain a copy of the License in the LICENSE file, or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.