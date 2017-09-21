# Introduction
This package was created with inspiration from @hopkinsth's [go-ruler](https://github.com/hopkinsth/go-ruler) to run a simple set of rules against an entity.

This version includes a couple more features including, AND and OR composites and the ability to add custom comparators.

# Example
```go
e := NewEngine()
e = e.AddComparator("always-false", func(a, b interface{}) bool {
    return false
})

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

props := map[string]interface{}{
    "user": map[string]interface{}{
        "name": "Trevor",
    }
}

res := e.Evaluate(props)
// res == true
```

# License

Copyright 2017 Trevor Hutto

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this work except in compliance with the License. You may obtain a copy of the License in the LICENSE file, or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.