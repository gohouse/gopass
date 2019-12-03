package rules

import "github.com/gohouse/gopass"

// DateFormat ...
const DateFormat = "2006-01-02"

// Default ...
func Default() gopass.ValidatorHandler {
	return func(v *gopass.Validator) {
		Required()(v)
		Min()(v)
		Max()(v)
		Numberic()(v)
		Length()(v)
		Before()(v)
		BeforeOrEqual()(v)
		DateEqual()(v)
		After()(v)
		AfterOrEqual()(v)
	}
}
