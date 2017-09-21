package grules

import (
	"testing"
)

func TestPluck(t *testing.T) {
	t.Run("key does not exist", func(t *testing.T) {
		props := map[string]interface{}{}
		val := pluck(props, "email")
		if val != nil {
			t.Fatal("expected value to be nil")
		}
	})

	t.Run("1 level", func(t *testing.T) {
		props := map[string]interface{}{
			"email": "test@test.com",
		}
		val := pluck(props, "email")
		if val.(string) != "test@test.com" {
			t.Fatal("expected value to match the given")
		}
	})

	t.Run("2 levels", func(t *testing.T) {
		props := map[string]interface{}{
			"user": map[string]interface{}{
				"name": "Trevor",
			},
		}
		val := pluck(props, "user.name")
		if val.(string) != "Trevor" {
			t.Fatal("expected value to match the given")
		}
	})

	t.Run("2 levels, key does not exist", func(t *testing.T) {
		props := map[string]interface{}{
			"user": map[string]interface{}{
				"name": "Trevor",
			},
		}
		val := pluck(props, "user.last_name")
		if val != nil {
			t.Fatal("expected value to be nil")
		}
	})
}
