package rules

import "github.com/gohouse/gopass"

func Default() gopass.ValidatorHandler {
	return func(v *gopass.Validator) {
		Required()(v)
		Min()(v)
		Max()(v)
		Numberic()(v)
	}
}
