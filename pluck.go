package grules

import (
	"strings"
)

func pluck(props map[string]interface{}, path string) interface{} {
	parts := strings.Split(path, ".")
	for i := 0; i < len(parts)-1; i++ {
		var ok bool
		props, ok = props[parts[i]].(map[string]interface{})
		if !ok {
			return nil
		}
	}
	return props[parts[len(parts)-1]]
}
