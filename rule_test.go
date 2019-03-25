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
		r := Rule{
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
		r := Rule{
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
		r := Rule{
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
		r := Rule{
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
	r := Rule{
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
	}
	props := map[string]interface{}{
		"name": "Trevor",
		"age":  float64(23),
	}

	t.Run("and", func(t *testing.T) {
		c := Composite{
			Operator: OperatorAnd,
			Rules: []Rule{
				Rule{
					Comparator: "eq",
					Path:       "name",
					Value:      "Trevor",
				},
				Rule{
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
		c := Composite{
			Operator: OperatorOr,
			Rules: []Rule{
				Rule{
					Comparator: "eq",
					Path:       "name",
					Value:      "John",
				},
				Rule{
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

	t.Run("unknown operator", func(t *testing.T) {
		c := Composite{
			Operator: "unknown",
			Rules: []Rule{
				Rule{
					Comparator: "eq",
					Path:       "name",
					Value:      "John",
				},
				Rule{
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
	c := Composite{
		Operator: "or",
		Rules: []Rule{
			Rule{
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
	e := NewEngine()
	e = e.AddComparator("always-false", comp)
	if e.comparators["always-false"] == nil {
		t.Fatal("expected comparator to be added under key always-false")
	}

	e.Composites = []Composite{
		Composite{
			Operator: OperatorAnd,
			Rules: []Rule{
				Rule{
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
		e := NewEngine()
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
		e := NewEngine()
		e.Composites = []Composite{
			Composite{
				Operator: OperatorAnd,
				Rules: []Rule{
					Rule{
						Comparator: "contains",
						Path:       "address.bedroom.furniture",
						Value:      "tv",
					},
				},
			},
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
		e := NewEngine()
		e.Composites = []Composite{
			Composite{
				Operator: OperatorAnd,
				Rules: []Rule{
					Rule{
						Comparator: "eq",
						Path:       "user.name",
						Value:      "Trevor",
					},
					Rule{
						Comparator: "eq",
						Path:       "user.id",
						Value:      float64(1234),
					},
				},
			},
			Composite{
				Operator: OperatorOr,
				Rules: []Rule{
					Rule{
						Comparator: "eq",
						Path:       "user.name",
						Value:      "Trevor",
					},
					Rule{
						Comparator: "eq",
						Path:       "user.id",
						Value:      float64(7),
					},
				},
			},
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
		e := NewEngine()
		e.Composites = []Composite{
			Composite{
				Operator: OperatorAnd,
				Rules: []Rule{
					Rule{
						Comparator: "contains",
						Path:       "user.favorites",
						Value:      "golang",
					},
				},
			},
		}
		res := e.Evaluate(props)
		if res != true {
			t.Fatal("expected engine to pass")
		}
	})
}

func BenchmarkEngine_Evaluate(b *testing.B) {
	e := NewEngine()
	e.Composites = []Composite{
		Composite{
			Operator: "or",
			Rules: []Rule{
				Rule{
					Comparator: "unit",
					Path:       "name",
					Value:      "Trevor",
				},
			},
		},
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
