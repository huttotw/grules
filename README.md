# Introduction
This package was created with inspiration from Thomas' [go-ruler](https://github.com/hopkinsth/go-ruler) to run a simple set of rules against an entity.

This version includes a couple more features including, AND and OR composites and the ability to add custom comparators.

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

# Benchmarks

|Benchmark|N|Speed|
|---------|----------|-----|
|BenchmarkEqual-8|200000000|8.19 ns/op|
|BenchmarkNotEqual-8|300000000|5.82 ns/op|
|BenchmarkLessThan-8|50000000|31.8 ns/op|
|BenchmarkLessThanEqual-8|50000000|26.9 ns/op|
|BenchmarkGreaterThan-8|30000000|37.7 ns/op|
|BenchmarkGreaterThanEqual-8|50000000|33.2 ns/op|
|BenchmarkContains-8|10000000|131 ns/op|
|BenchmarkContainsLong50000-8|500|2944142 ns/op|
|BenchmarkPluckShallow-8|30000000|53.8 ns/op|
|BenchmarkPluckDeep-8|5000000|311 ns/op|

# License

Copyright 2017 Trevor Hutto

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this work except in compliance with the License. You may obtain a copy of the License in the LICENSE file, or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.