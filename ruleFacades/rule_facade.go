package ruleFacades

import "fmt"

// RuleFacades ...
type RuleFacades struct{}

// NewRuleFacades ...
func NewRuleFacades() *RuleFacades {
	return &RuleFacades{}
}

// Required ...
func (*RuleFacades) Required() string {
	return "required"
}

// Numberic ...
func (*RuleFacades) Numberic() string {
	return "numberic"
}

// Min ...
func (*RuleFacades) Min(arg interface{}) string {
	return fmt.Sprintf("min:%v", arg)
}

// Max ...
func (*RuleFacades) Max(arg interface{}) string {
	return fmt.Sprintf("max:%v", arg)
}

// Length ...
func (*RuleFacades) Length(arg interface{}) string {
	return fmt.Sprintf("length:%v", arg)
}
