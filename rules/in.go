package rules

import (
	"fmt"
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
	"strings"
)

// InString 定义规则名字
const InString = "in"

func init() {
	// 注册 required 验证器
	gopassRuler.Register(InString, &InValidator{})
	// 注册 required 错误提示消息
	// gopassRuler.RegisterMessageMulti(defaultMessages)
}

// InValidator 参数范围验证器
type InValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *InValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	split := t.New(rule).SpliteAndTrimSpace(",")
	if len(split) == 0 {
		return e.New("rule error, eg: (in:1,a,3,b)")
	}
	if !t.New(data).InArrayString(split) {
		return gopassRuler.GetMsgByRule(InString, field, rule, msgs...)
	}
	return nil
}

// In 参数范围
func In(args ...interface{}) string {
	var tmp []string
	for _, v := range args {
		tmp = append(tmp, t.New(v).String())
	}
	return fmt.Sprintf("in:%s", strings.Join(tmp, ","))
}
