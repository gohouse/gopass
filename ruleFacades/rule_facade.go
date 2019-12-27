package ruleFacades

import (
	"fmt"
	"github.com/gohouse/t"
	"strings"
)

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

// Before ...
func (*RuleFacades) Before(arg interface{}) string {
	return fmt.Sprintf("before:%v", arg)
}

// BeforeOrEqual ...
func (*RuleFacades) BeforeOrEqual(arg interface{}) string {
	return fmt.Sprintf("before_or_equal:%v", arg)
}

// DateEqual ...
func (*RuleFacades) DateEqual(arg interface{}) string {
	return fmt.Sprintf("date_equal:%v", arg)
}

// After ...
func (*RuleFacades) After(arg interface{}) string {
	return fmt.Sprintf("after:%v", arg)
}

// AfterOrEqual ...
func (*RuleFacades) AfterOrEqual(arg interface{}) string {
	return fmt.Sprintf("after_or_equal:%v", arg)
}

// In ...
func (*RuleFacades) In(arg []interface{}) string {
	return fmt.Sprintf("in:%v", strings.Join(t.New(arg).SliceString(), "|"))
}
