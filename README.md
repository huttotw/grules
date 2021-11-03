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
    "comparer": "eq",
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
    "comparer": "lt",
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
                    "comparer": "eq",
                    "value": "darth"
                },
                {
                    "path": "name.last",
                    "comparer": "eq",
                    "value": "vader"
                }
            ]
        },
        {
            "operator": "or",
            "rules": [
                {
                    "path": "order",
                    "comparer": "eq",
                    "value": "first world order"
                },
                {
                    "operator": "or",
                    "path": "friends.#.order",
                    "comparer": "contains",
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
| BenchmarkEqual-12                | 650602549  | 5.52 ns/op   | 0 B/op    | 0 allocs/op  |
| BenchmarkNotEqual-12             | 876894124  | 4.09 ns/op   | 0 B/op    | 0 allocs/op  |
| BenchmarkLessThan-12             | 1000000000 | 2.84 ns/op   | 0 B/op    | 0 allocs/op  |
| BenchmarkLessThanEqual-12        | 1000000000 | 2.57 ns/op   | 0 B/op    | 0 allocs/op  |
| BenchmarkGreaterThan-12          | 1000000000 | 2.07 ns/op   | 0 B/op    | 0 allocs/op  |
| BenchmarkGreaterThanEqual-12     | 1000000000 | 2.86 ns/op   | 0 B/op    | 0 allocs/op  |
| BenchmarkRegex-12                | 4524237    | 793 ns/op    | 753 B/op  | 11 allocs/op |
| BenchmarkRegexPhone-12           | 1000000    | 3338 ns/op   | 3199 B/op | 30 allocs/op |
| BenchmarkContains-12             | 499627219  | 7.16 ns/op   | 0 B/op    | 0 allocs/op  |
| BenchmarkStringContains-12       | 405497102  | 8.87 ns/op   | 0 B/op    | 0 allocs/op  |
| BenchmarkContainsLong50000-12    | 18992      | 184016 ns/op | 0 B/op    | 0 allocs/op  |
| BenchmarkNotContains-12          | 292932907  | 12.3 ns/op   | 0 B/op    | 0 allocs/op  |
| BenchmarkStringNotContains-12    | 392618857  | 9.14 ns/op   | 0 B/op    | 0 allocs/op  |
| BenchmarkNotContainsLong50000-12 | 19243      | 191787 ns/op | 0 B/op    | 0 allocs/op  |
| BenchmarkOneOf-12                | 1000000000 | 1.80 ns/op   | 0 B/op    | 0 allocs/op  |
| BenchmarkNoneOf-12               | 1000000000 | 1.79 ns/op   | 0 B/op    | 0 allocs/op  |
| BenchmarkPluckShallow-12         | 85997188   | 41.6 ns/op   | 16 B/op   | 1 allocs/op  |
| BenchmarkPluckDeep-12            | 18789103   | 194 ns/op    | 112 B/op  | 1 allocs/op  |
| BenchmarkRule_evaluate-12        | 69558996   | 51.1 ns/op   | 16 B/op   | 1 allocs/op  |
| BenchmarkComposite_evaluate-12   | 59484760   | 55.7 ns/op   | 16 B/op   | 1 allocs/op  |
| BenchmarkEngine_Evaluate-12      | 47892318   | 75.0 ns/op   | 16 B/op   | 1 allocs/op  |

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
