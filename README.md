<img src="https://socialify.git.ci/huttotw/grules/image?description=1&font=Raleway&language=1&owner=1&pattern=Overlapping%20Hexagons&stargazers=1&theme=Dark" alt="grules" />

# Introduction

A simple rules engine with great flexibility capable of building binary decision trees of any depth.

Utilizes [gjson](https://github.com/tidwall/gjson) for pathing for extra flexibility.

# Simple Example (pass)

```go
json := `
{
    "name": {"first": "anakin", "last": "skywalker"},
    "age": 22,
    "children": ["luke", "leia"],
    "order": "jedi",
    "friends": [
        {"first": "r2d2",  "last": "droid",      "order": "republic", "age": 13, "episodes": [1,2,3,4,5,6,7,8,9]},
        {"first": "ben",   "last": "kenobi",     "order": "jedi",     "age": 38, "episodes": [1,2,3,4,5,6]},
        {"first": "c3po",  "last": "droid",      "order": "republic", "age": 13, "episodes": [1,2,3,4,5,6,7,8,9]},
        {"first": "sheev", "last": "palpatine",  "order": "sith",     "age": 63, "episodes": [1,2,3,5,6,9]}
    ]
}
`

rule := `
{
    "comparators": "eq",
    "path": "name.first",
    "value": "anakin"
}
`

pass, failReason := grules.Evaluate(json, rule)
if !pass {
    fmt.Println("FAILED: ", failreason)
}

fmt.Println(pass)

```

Output: `true`

# Simple Example (fail)

```go
json := `
{
    "name": {"first": "anakin", "last": "skywalker"},
    "age": 22,
    "children": ["luke", "leia"],
    "order": "jedi",
    "friends": [
        {"first": "r2d2",  "last": "droid",      "order": "republic", "age": 13, "episodes": [1,2,3,4,5,6,7,8,9]},
        {"first": "ben",   "last": "kenobi",     "order": "jedi",     "age": 38, "episodes": [1,2,3,4,5,6]},
        {"first": "c3po",  "last": "droid",      "order": "republic", "age": 13, "episodes": [1,2,3,4,5,6,7,8,9]},
        {"first": "sheev", "last": "palpatine",  "order": "sith",     "age": 63, "episodes": [1,2,3,5,6,9]}
    ]
}
`

rule := `
{
    "comparator": "lt",
    "path": "age",
    "value": "20"
}
`

passed, failReason := grules.Evaluate(json, rule)
if !passed {
    fmt.Println("FAILED: ", failreason)
}

fmt.Println(passed)

```

Output: `FAILED: value '22' at 'age' is not 'less than' rule value '20'`

# Complicated Example (pass)

```go
json := `
{
    "name": {"first": "anakin", "last": "skywalker"},
    "age": 22,
    "children": ["luke", "leia"],
    "order": "jedi",
    "friends": [
        {"first": "r2d2",  "last": "droid",      "order": "republic", "age": 13, "episodes": [1,2,3,4,5,6,7,8,9]},
        {"first": "ben",   "last": "kenobi",     "order": "jedi",     "age": 38, "episodes": [1,2,3,4,5,6]},
        {"first": "c3po",  "last": "droid",      "order": "republic", "age": 13, "episodes": [1,2,3,4,5,6,7,8,9]},
        {"first": "sheev", "last": "palpatine",  "order": "sith",     "age": 63, "episodes": [1,2,3,5,6,9]}
    ]
}
`

rule := `
{
    "operator": "or",
    "rules": [
        {
            "operator": "and",
            "rules": [
                {
                    "path": "name.first",
                    "comparator": "eq",
                    "value": "darth"
                },
                {
                    "path": "name.last",
                    "comparator": "eq",
                    "value": "vader"
                }
            ]
        },
        {
            "operator": "or",
            "rules": [
                {
                    "path": "order",
                    "comparator": "eq",
                    "value": "first world order"
                },
                {
                    "operator": "or",
                    "path": "friends.#.order",
                    "comparator": "contains",
                    "value": "sith"
                }
            ]
        }
    ]
}
`

pass, failReason := grules.Evaluate(json, rule)
if !pass {
    fmt.Println("FAILED: ", failreason)
}

fmt.Println(pass)

```

Output: `true`

# Comparators

- `eq` will return true if `a == b`
- `neq` will return true if `a != b`
- `lt` will return true if `a < b`
- `lte` will return true if `a <= b`
- `gt` will return true if `a > b`
- `gte` will return true if `a >= b`
- `contains` will return true if `a` contains `b`
- `ncontains` will return true if `a` does not contain `b`
- `oneof` will return true if `a` is one of `b`
- `noneof` will return true if `a` is not one of `b`
- `regex` will return true if `a` matches `b`

`contains` and `ncontains` work for substring comparisons as well as item-in-collection comparisons.

When used for item-in-collection comparisons, `contains` expects the first argument to be a slice. `contains` is different than `oneof` in that `oneof` expects the second argument to be a slice.

# Benchmarks

| Benchmark                        | N          | Speed        | Used      | Allocs       |
| -------------------------------- | ---------- | ------------ | --------- | ------------ |
| BenchmarkEqual-12                | 805088716  | 4.588 ns/op  | 0 B/op    | 0 allocs/op
| BenchmarkNotEqual-12             | 1000000000 | 3.203 ns/op  | 0 B/op    | 0 allocs/op
| BenchmarkLessThan-12             | 1000000000 | 0.2521 ns/op | 0 B/op    | 0 allocs/op
| BenchmarkLessThanEqual-12        | 1000000000 | 0.2639 ns/op | 0 B/op    | 0 allocs/op
| BenchmarkGreaterThan-12          | 1000000000 | 0.2550 ns/op | 0 B/op    | 0 allocs/op
| BenchmarkGreaterThanEqual-12     | 1000000000 | 0.2542 ns/op | 0 B/op    | 0 allocs/op
| BenchmarkRegex-12                | 4384954    | 813.7 ns/op  | 754 B/op  | 11 allocs/op
| BenchmarkRegexPhone-12           | 1000000    | 3378 ns/op   | 3201 B/op | 30 allocs/op
| BenchmarkContains-12             | 525908002  | 6.169 ns/op  | 0 B/op    | 0 allocs/op
| BenchmarkStringContains-12       | 497205760  | 7.790 ns/op  | 0 B/op    | 0 allocs/op
| BenchmarkContainsLong50000-12    | 17085      | 217622 ns/op | 0 B/op    | 0 allocs/op
| BenchmarkNotContains-12          | 314265795  | 12.05 ns/op  | 0 B/op    | 0 allocs/op
| BenchmarkStringNotContains-12    | 470831580  | 7.750 ns/op  | 0 B/op    | 0 allocs/op
| BenchmarkNotContainsLong50000-12 | 17414      | 224783 ns/op | 0 B/op    | 0 allocs/op
| BenchmarkOneOf-12                | 1000000000 | 1.582 ns/op  | 0 B/op    | 0 allocs/op
| BenchmarkNoneOf-12               | 1000000000 | 1.632 ns/op  | 0 B/op    | 0 allocs/op
| BenchmarkEvaluate-12             | 35343126   | 95.07 ns/op  | 0 B/op    | 0 allocs/op

To run benchmarks:

```
go test -run none -bench . -benchtime 3s -benchmem
```

All benchmarks were run on:

MacOS Big Sur 2.6Ghz Intel Core i7 16 GB 2400 MHz DDR4

# License

Copyright &copy; 2019 Trevor Hutto

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this work except in compliance with the License. You may obtain a copy of the License in the LICENSE file, or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
