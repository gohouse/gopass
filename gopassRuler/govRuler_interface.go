package gopassRuler

import "github.com/gohouse/e"

type IValidator interface {
	Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error
}
