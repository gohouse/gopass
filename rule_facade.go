package gopass

import "fmt"

// RuleFacade ...
type RuleFacade struct{}

// NewRuleFacade ...
func NewRuleFacade() *RuleFacade {
	return &RuleFacade{}
}

// Required ...
func (*RuleFacade) Required() string {
	return "required"
}

// Numberic ...
func (*RuleFacade) Numberic() string {
	return "numberic"
}

// Min ...
func (*RuleFacade) Min(arg interface{}) string {
	return fmt.Sprintf("min:%v", arg)
}

// Max ...
func (*RuleFacade) Max(arg interface{}) string {
	return fmt.Sprintf("max:%v", arg)
}

// Length ...
func (*RuleFacade) Length(arg interface{}) string {
	return fmt.Sprintf("length:%v", arg)
}
