package grules

import "errors"

type Rule struct {
	Path       string      `json:"path,omitempty"`
	Comparator string      `json:"comparator,omitempty"`
	Value      interface{} `json:"value,omitempty"`

	Operator Operator `json:"operator,omitempty"`
	Rules    []Rule   `json:"rules,omitempty"`
}

func (r Rule) HasChildren() bool {
	return len(r.Rules) > 0
}

func (r Rule) Validate() error {
	pathIsSet := r.Path != ""
	comparatorIsSet := r.Comparator != ""
	valueIsSet := r.Value != nil

	standardRuleValid := pathIsSet && comparatorIsSet && valueIsSet

	operatorIsSet := r.Operator != ""
	rulesIsSet := len(r.Rules) != 0

	multiRuleValid := operatorIsSet && rulesIsSet

	if standardRuleValid && multiRuleValid {
		return errors.New("setting path, comparator, and value AS WELL AS operator and rules is not valid")
	}

	if !standardRuleValid && !multiRuleValid {
		return errors.New("must set either path, comparator, and value OR operator and rules")
	}

	return nil
}
