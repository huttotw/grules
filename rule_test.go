package grules

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestRule_evaluate(t *testing.T) {
	comparators := map[string]Comparator{
		"eq": equal,
	}
	props := map[string]interface{}{
		"first_name": "Trevor",
	}
	t.Run("basic rule", func(t *testing.T) {
		r := rule{
			Comparator: "eq",
			Path:       "first_name",
			Value:      "Trevor",
		}
		res := r.evaluate(props, comparators)
		if res != true {
			t.Fatal("expected rule to be true")
		}
	})

	t.Run("unknown path", func(t *testing.T) {
		r := rule{
			Comparator: "eq",
			Path:       "email",
			Value:      "Trevor",
		}
		res := r.evaluate(props, comparators)
		if res != false {
			t.Fatal("expected rule to be false")
		}
	})

	t.Run("non comparable types", func(t *testing.T) {
		r := rule{
			Comparator: "eq",
			Path:       "name",
			Value:      func() {},
		}
		res := r.evaluate(props, comparators)
		if res != false {
			t.Fatal("expected rule to be false")
		}
	})

	t.Run("unknown comparator", func(t *testing.T) {
		r := rule{
			Comparator: "unknown",
			Path:       "name",
			Value:      "Trevor",
		}
		res := r.evaluate(props, comparators)
		if res != false {
			t.Fatal("expected rule to be false")
		}
	})
}

func BenchmarkRule_evaluate(b *testing.B) {
	r := rule{
		Comparator: "unit",
		Path:       "name",
		Value:      "Trevor",
	}
	props := map[string]interface{}{
		"name": "Trevor",
	}
	comps := map[string]Comparator{
		"unit": func(a, b interface{}) bool {
			return true
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.evaluate(props, comps)
	}
}

func TestRule_MarshalJSON(t *testing.T) {
	t.Run("simple engine", func(t *testing.T) {
		j := []byte(`{"composites":[{"operator":"and","rules":[{"comparator":"eq","path":"first_name","value":"Trevor"}]}]}`)
		e, err := NewJSONEngine(j)
		if err != nil {
			t.Fatal(err)
		}

		b, err := json.Marshal(e)
		if err != nil {
			t.Fatal(err)
		}

		if string(b) != string(j) {
			t.Fatal("expected json to be same")
		}
	})

	t.Run("list to map", func(t *testing.T) {
		j := []byte(`{"composites":[{"operator":"and","rules":[{"comparator":"oneof","path":"first_name","value":["Trevor"]}]}]}`)
		e, err := NewJSONEngine(j)
		if err != nil {
			t.Fatal(err)
		}

		b, err := json.Marshal(e)
		if err != nil {
			t.Fatal(err)
		}

		if string(b) != string(j) {
			t.Fatal("expected json to be same")
		}
	})
}

func TestComposite_evaluate(t *testing.T) {
	comparators := map[string]Comparator{
		"eq": equal,
		"gt": greaterThan,
		"lt": lessThan,
	}
	props := map[string]interface{}{
		"name": "Trevor",
		"age":  float64(23),
	}

	t.Run("and", func(t *testing.T) {
		c := composite{
			Operator: OperatorAnd,
			Rules: []rule{
				rule{
					Comparator: "eq",
					Path:       "name",
					Value:      "Trevor",
				},
				rule{
					Comparator: "eq",
					Path:       "age",
					Value:      float64(23),
				},
			},
		}
		res := c.evaluate(props, comparators)
		if res != true {
			t.Fatal("expected composite to be true")
		}
	})

	t.Run("or", func(t *testing.T) {
		c := composite{
			Operator: OperatorOr,
			Rules: []rule{
				rule{
					Comparator: "eq",
					Path:       "name",
					Value:      "John",
				},
				rule{
					Comparator: "eq",
					Path:       "age",
					Value:      float64(23),
				},
			},
		}
		res := c.evaluate(props, comparators)
		if res != true {
			t.Fatal("expected composite to be true")
		}
	})

	t.Run("nested and - or", func(t *testing.T) {
		c := Composite{
			Operator: OperatorAnd,
			Rules: []Rule{
				Rule{
					Comparator: "eq",
					Path:       "name",
					Value:      "Trevor",
				},
			},
			Composites: []Composite{
				Composite{
					Operator: OperatorOr,
					Rules: []Rule{
						Rule{
							Comparator: "gt",
							Path:       "age",
							Value:      float64(20),
						},
						Rule{
							Comparator: "lt",
							Path:       "age",
							Value:      float64(20),
						},
					},
				},
			},
		}
		res := c.evaluate(props, comparators)
		if res != true {
			t.Fatal("expected composite to be true")
		}
	})

	t.Run("nested or - and", func(t *testing.T) {
		c := Composite{
			Operator: OperatorOr,
			Rules: []Rule{
				Rule{
					Comparator: "eq",
					Path:       "name",
					Value:      "John",
				},
			},
			Composites: []Composite{
				Composite{
					Operator: OperatorAnd,
					Rules: []Rule{
						Rule{
							Comparator: "gt",
							Path:       "age",
							Value:      float64(20),
						},
						Rule{
							Comparator: "lt",
							Path:       "age",
							Value:      float64(30),
						},
					},
				},
			},
		}
		res := c.evaluate(props, comparators)
		if res != true {
			t.Fatal("expected composite to be true")
		}
	})

	t.Run("unknown operator", func(t *testing.T) {
		c := composite{
			Operator: "unknown",
			Rules: []rule{
				rule{
					Comparator: "eq",
					Path:       "name",
					Value:      "John",
				},
				rule{
					Comparator: "eq",
					Path:       "age",
					Value:      float64(23),
				},
			},
		}
		res := c.evaluate(props, comparators)
		if res != false {
			t.Fatal("expected composite to be true")
		}
	})
}

func BenchmarkComposite_evaluate(b *testing.B) {
	c := composite{
		Operator: "or",
		Rules: []rule{
			rule{
				Comparator: "unit",
				Path:       "name",
				Value:      "Trevor",
			},
		},
	}

	props := map[string]interface{}{
		"name": "Trevor",
	}
	comps := map[string]Comparator{
		"unit": func(a, b interface{}) bool {
			return true
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.evaluate(props, comps)
	}
}

func TestAddComparator(t *testing.T) {
	comp := func(a, b interface{}) bool {
		return false
	}
	e, err := NewJSONEngine(json.RawMessage(`{}`))
	if err != nil {
		t.Fatal(err)
	}
	e = e.AddComparator("always-false", comp)
	if e.comparators["always-false"] == nil {
		t.Fatal("expected comparator to be added under key always-false")
	}

	e.Composites = []composite{
		composite{
			Operator: OperatorAnd,
			Rules: []rule{
				rule{
					Comparator: "always-false",
					Path:       "user.name",
					Value:      "Trevor",
				},
			},
		},
	}

	props := map[string]interface{}{
		"user": map[string]interface{}{
			"name": "Trevor",
		},
	}

	res := e.Evaluate(props)
	if res != false {
		t.Fatal("expected engine to be false")
	}
}

func TestNewJSONEngine(t *testing.T) {
	t.Run("simple engine", func(t *testing.T) {
		j := []byte(`{"composites":[{"operator":"and","rules":[{"comparator":"eq","path":"first_name","value":"Trevor"}]}]}`)
		e, err := NewJSONEngine(j)
		if err != nil {
			t.Fatal(err)
		}
		if len(e.Composites) != 1 {
			t.Fatal("expected 1 composite")
		}
		if len(e.Composites[0].Rules) != 1 {
			t.Fatal("expected 1 rule in first composite")
		}
	})

	t.Run("list to map", func(t *testing.T) {
		j := []byte(`{"composites":[{"operator":"and","rules":[{"comparator":"oneof","path":"first_name","value":["Trevor"]}]}]}`)
		e, err := NewJSONEngine(j)
		if err != nil {
			t.Fatal(err)
		}
		if len(e.Composites) != 1 {
			t.Fatal("expected 1 composite")
		}
		if len(e.Composites[0].Rules) != 1 {
			t.Fatal("expected 1 rule in first composite")
		}

		if reflect.TypeOf(e.Composites[0].Rules[0].Value).Kind() != reflect.Map {
			t.Fatal("expected list to be transformed to map")
		}
	})
}

func TestEngineEvaluate(t *testing.T) {
	t.Run("no composites", func(t *testing.T) {
		props := map[string]interface{}{
			"user": map[string]interface{}{
				"email": "test@test.com",
				"name":  "Trevor",
				"id":    float64(1234),
			},
		}
		e, err := NewJSONEngine(json.RawMessage(`{}`))
		if err != nil {
			t.Fatal(err)
		}
		res := e.Evaluate(props)
		if res != true {
			t.Fatal("expected engine to pass")
		}
	})

	t.Run("1 composite, 1 rule", func(t *testing.T) {
		props := map[string]interface{}{
			"address": map[string]interface{}{
				"bedroom": map[string]interface{}{
					"furniture": []interface{}{
						"bed",
						"tv",
						"dresser",
					},
				},
			},
		}
		e, err := NewJSONEngine(json.RawMessage(`{"composites":[{"operator":"and","rules":[{"comparator":"contains","path":"address.bedroom.furniture","value":"tv"}]}]}`))
		if err != nil {
			t.Fatal(err)
		}
		res := e.Evaluate(props)
		if res != true {
			t.Fatal("expected engine to pass")
		}
	})

	t.Run("2 composites, 1 rule", func(t *testing.T) {
		props := map[string]interface{}{
			"user": map[string]interface{}{
				"email": "test@test.com",
				"name":  "Trevor",
				"id":    float64(1234),
			},
		}
		e, err := NewJSONEngine(json.RawMessage(`{"composites":[{"operator":"and","rules":[{"comparator":"eq","path":"user.name","value":"Trevor"},{"comparator":"eq","path":"user.id","value":1234}]},{"operator":"or","rules":[{"comparator":"eq","path":"user.name","value":"Trevor"},{"comparator":"eq","path":"user.id","value":7}]}]}`))
		if err != nil {
			t.Fatal(err)
		}
		res := e.Evaluate(props)
		if res != true {
			t.Fatal("expected engine to pass")
		}
	})

	t.Run("1 composites, 1 rule, strictly typed list", func(t *testing.T) {
		props := map[string]interface{}{
			"user": map[string]interface{}{
				"email": "test@test.com",
				"name":  "Trevor",
				"id":    float64(1234),
				"favorites": []string{
					"golang",
					"javascript",
				},
			},
		}
		e, err := NewJSONEngine(json.RawMessage(`{"composites":[{"operator":"and","rules":[{"comparator":"contains","path":"user.favorites","value":"golang"}]}]}`))
		if err != nil {
			t.Fatal(err)
		}
		res := e.Evaluate(props)
		if res != true {
			t.Fatal("expected engine to pass")
		}
	})
}

func BenchmarkEngine_Evaluate(b *testing.B) {
	e, err := NewJSONEngine(json.RawMessage(`{"composites":[{"operator":"and","rules":[{"comparator":"unit","path":"name","value":"Trevor"}]}]}`))
	if err != nil {
		b.Fatal(err)
	}
	e.AddComparator("unit", func(a, b interface{}) bool { return true })
	props := map[string]interface{}{
		"name": "Trevor",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Evaluate(props)
	}
}
