package grules

import "github.com/tidwall/gjson"

// specialSlice is just syntactic sugar. We need a 2d slice when doing
// compares that involve slices. For instance, we could have a json like:
// { "nums": [1,2,3] }
// and if we wanted to see if there was a 2 in the 'nums' that would be easy,
// we just simply do a 'contains' comparator. However, what if we have json:
// {
//  	"group": [
//  		{ "nums": [1,2,3] },
//  		{ "nums": [4,5,6] }
//  	]
// }
// and we wanted to ask does ALL 'nums' in 'group' contain 2 or even
// does ANY 'nums' in 'group' contain 2. So we transform that group into
// a 2 deminional arraya and loop through each slice of interfaces and
// run the contains on each and applying the 'Operator' logic on each
type specialSlice [][]interface{}

func newSpecialSlice(values []gjson.Result) specialSlice {
	var specialSlice specialSlice
	var slice []interface{}

	for _, value := range values {
		switch value.Type {
		case gjson.String:
			slice = append(slice, value.Str)
		case gjson.Number:
			slice = append(slice, value.Num)
		case gjson.True:
			fallthrough
		case gjson.False:
			slice = append(slice, value.Bool())
		default:
			if value.IsArray() {
				specialSlice = append(specialSlice, newSpecialSlice(value.Array())...)
			}
		}
	}

	if slice != nil {
		specialSlice = append(specialSlice, slice)
	}

	return specialSlice
}

func (s specialSlice) evalualte(comparator Comparator, rule Rule) bool {
	switch rule.Operator {
	case Or:
		for _, slice := range s {
			ok := comparator(slice, rule.Value)
			if ok {
				return true
			}
		}

		return false
	case And:
		fallthrough
	default:
		for _, slice := range s {
			ok := comparator(slice, rule.Value)
			if !ok {
				return false
			}
		}

		return true
	}
}
