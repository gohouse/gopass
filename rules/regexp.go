package rules

import (
	"fmt"
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
	"regexp"
)

// RegexpString 定义规则名字
const RegexpString = "regexp"

func init() {
	// 注册 url 验证器
	gopassRuler.Register(RegexpString, &RegexpValidator{})
	// 注册 url 错误提示消息
	// gopassRuler.RegisterMessageMulti(defaultMessages)
}

// RegexpValidator 必传参数验证器
type RegexpValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *RegexpValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {

	if !regexp.MustCompile(rule).Match(t.New(data).Bytes()) {
		return gopassRuler.GetMsgByRule(RegexpString, field, rule, msgs...)
	}
	return nil
}

// Regexp 参数需为正则
func Regexp(exp string) string {
	return fmt.Sprintf("regexp:%s", exp)
}
