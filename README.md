[![CircleCI](https://circleci.com/gh/huttotw/grules/tree/master.svg?style=svg)](https://circleci.com/gh/huttotw/grules/tree/master)

# Introduction

This package was created with inspiration from Thomas' [go-ruler](https://github.com/hopkinsth/go-ruler) to run a simple set of rules against an entity.

This version includes a couple more features including, AND and OR composites and the ability to add custom comparators.

**Note**: This package only compares two types: `string` and `float64`, this plays nicely with `encoding/json`.

# Example

```go
// Create a new instance of an engine with some default comparators
e, err := NewJSONEngine(json.RawMessage(`{"composites":[{"operator":"or","rules":[{"comparator":"always-false","path":"user.name","value":"Trevor"},{"comparator":"eq","path":"user.name","value":"Trevor"}]}]}`))
if err != nil {
    panic(err)
}

// Add a new, custom comparator
e = e.AddComparator("always-false", func(a, b interface{}) bool {
    return false
})

// Give some properties, this map can be deeper and supports interfaces
props := map[string]interface{}{
    "user": map[string]interface{}{
        "name": "Trevor",
        "age": float64(25),
    }
}

// Run the engine on the props
res := e.Evaluate(props)
// res == true
```

# Comparators

- `eq` will return true if `a == b`
- `neq` will return true if `a != b`
- `lt` will return true if `a < b`
- `lte` will return true if `a <= b`
- `gt` will return true if `a > b`
- `gte` will return true if `a >= b`
- `contains` will return true if `a` contains `b`
- `oneof` will return true if `a` is one of `b`

`contains` is different than `oneof` in that `contains` expects the first argument to be a slice, and `oneof` expects the second argument to be a slice.

# Benchmarks

| Benchmark                        | N          | Speed        | Used     | Allocs      |
| -------------------------------- | ---------- | ------------ | -------- | ----------- |
| BenchmarkEqual-12                | 1000000000 | 5.22 ns/op   | 0 B/op   | 0 allocs/op |
| BenchmarkNotEqual-12             | 2000000000 | 3.77 ns/op   | 0 B/op   | 0 allocs/op |
| BenchmarkLessThan-12             | 2000000000 | 2.20 ns/op   | 0 B/op   | 0 allocs/op |
| BenchmarkLessThanEqual-12        | 2000000000 | 1.95 ns/op   | 0 B/op   | 0 allocs/op |
| BenchmarkGreaterThan-12          | 5000000000 | 1.95 ns/op   | 0 B/op   | 0 allocs/op |
| BenchmarkGreaterThanEqual-12     | 2000000000 | 1.97 ns/op   | 0 B/op   | 0 allocs/op |
| BenchmarkContains-12             | 1000000000 | 5.66 ns/op   | 0 B/op   | 0 allocs/op |
| BenchmarkContainsLong50000-12    | 30000      | 157679 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkNotContains-12          | 500000000  | 11.5 ns/op   | 0 B/op   | 0 allocs/op |
| BenchmarkNotContainsLong50000-12 | 30000      | 157437 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkOneOf-12                | 500000000  | 0.53 ns/op   | 0 B/op   | 0 allocs/op |
| BenchmarkNoneOf-12               | 500000000  | 0.53 ns/op   | 0 B/op   | 0 allocs/op |
| BenchmarkPluckShallow-12         | 100000000  | 42.4 ns/op   | 16 B/op  | 1 allocs/op |
| BenchmarkPluckDeep-12            | 30000000   | 174 ns/op    | 112 B/op | 1 allocs/op |
| BenchmarkRule_evaluate-12        | 100000000  | 51.7 ns/op   | 16 B/op  | 1 allocs/op |
| BenchmarkComposite_evaluate-12   | 100000000  | 58.9 ns/op   | 16 B/op  | 1 allocs/op |
| BenchmarkEngine_Evaluate-12      | 100000000  | 69.9 ns/op   | 16 B/op  | 1 allocs/op |

To run benchmarks:

```
go test -run none -bench . -benchtime 3s -benchmem
```

All benchmarks were run on:

MacOS High Sierra 2.6Ghz Intel Core i7 16 GB 2400 MHz DDR4

# License

Copyright &copy; 2019 Trevor Hutto

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this work except in compliance with the License. You may obtain a copy of the License in the LICENSE file, or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
