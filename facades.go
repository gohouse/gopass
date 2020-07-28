package gopass

import "github.com/gohouse/gopass/rules"

func Required() string {
	return rules.Required()
}
func Numeric() string {
	return rules.Numeric()
}
func Integer() string {
	return rules.Integer()
}
func Float() string {
	return rules.Float()
}
func Url() string {
	return rules.Url()
}
func Email() string {
	return rules.Emaill()
}
func ChineseMobile() string {
	return rules.ChineseMobile()
}
func Json() string {
	return rules.Json()
}
func Xml() string {
	return rules.Xml()
}
func Date() string {
	return rules.Date()
}
func Between(min, max interface{}) string {
	return rules.Between(min, max)
}
func Length(arg interface{}) string {
	return rules.Length(arg)
}
func Min(arg interface{}) string {
	return rules.Min(arg)
}
func Max(arg interface{}) string {
	return rules.Max(arg)
}
func In(args ...interface{}) string {
	return rules.In(args...)
}
func Regexp(arg string) string {
	return rules.Regexp(arg)
}
