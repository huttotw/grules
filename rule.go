package grules

import "errors"

type Rule struct {
	Path     string      `json:"path,omitempty"`
	Comparer string      `json:"comparer,omitempty"`
	Value    interface{} `json:"value,omitempty"`

	Operator Operator `json:"operator,omitempty"`
	Rules    []Rule   `json:"rules,omitempty"`
}

func (r Rule) HasChildren() bool {
	return len(r.Rules) > 0
}

func (r Rule) Validate() error {
	pathIsSet := r.Path != ""
	comparerIsSet := r.Comparer != ""
	valueIsSet := r.Value != nil

	standardRuleValid := pathIsSet && comparerIsSet && valueIsSet

	operatorIsSet := r.Operator != ""
	rulesIsSet := len(r.Rules) != 0

	multiRuleValid := operatorIsSet && rulesIsSet

	if standardRuleValid && multiRuleValid {
		return errors.New("setting path, comparer, and value AS WELL AS operator and rules is not valid")
	}

	if !standardRuleValid && !multiRuleValid {
		return errors.New("must set either path, comparer, and value OR operator and rules")
	}

	return nil
}
