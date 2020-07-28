package gopass

import (
	"fmt"
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"reflect"
	"strings"
)

var RuleSep = "|"

// Rules 构建规则数据结构
type Rules map[string][]string

// 初始化可以自定义的选项
type Option struct {
	// 设置语言,不设置, 默认zh_cn
	DefaultLanguage string
	// 自定义提示语,优先级最高
	DefaultMessage map[string]interface{}
	// 自定义多条规则分割服, 默认 |
	RuleSep string
}

// Gov 主结构体
type Gov struct {
	*Option
}

// NewGov 初始化Gov设置
func NewGov(opt *Option) *Gov {
	if opt != nil {
		if opt.DefaultLanguage != "" {
			gopassRuler.SetDefaultLang(opt.DefaultLanguage)
		}
		if opt.DefaultMessage != nil {
			gopassRuler.RegisterMessageMulti(opt.DefaultMessage)
		}
		if opt.RuleSep == "" {
			opt.RuleSep = RuleSep
		}
	} else {
		opt = &Option{
			RuleSep: RuleSep,
		}
	}
	return &Gov{opt}
}

func (*Gov) Validate(data map[string]interface{}, rules Rules, msgs ...map[string]string) (err e.Error) {
	// 遍历规则, 进行校验
	for field, item := range rules {
		var val interface{}
		if v, ok := data[field]; ok {
			val = v
		}
		for _, rule := range item {
			// 对rule进行解析
			ruleSlice := strings.Split(rule, ":")
			var ruleReal string
			switch len(ruleSlice) {
			case 0:
				return e.New(fmt.Sprintf("规则不存在: %s", rule))
			case 1: // 如: required; numeric
				ruleReal = ""
			default: // 如: min:3; between:3,6; in:3,4,5; regexp:\d+\:\w+
				// 兼容正则内有冒号的情况, 只截取第一个冒号前的字符串为规则名字
				ruleReal = strings.TrimSpace(strings.Join(ruleSlice[1:], ":"))
			}

			v := gopassRuler.Getter(ruleSlice[0])
			if v != nil {
				err = v.Validate(val, field, ruleReal, msgs...)
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func (g *Gov) ValidateField(value interface{}, rules []string, msg ...map[string]string) (err e.Error) {
	var field = "field"
	return g.Validate(map[string]interface{}{field: value}, Rules{field: rules}, msg...)
}

func (g *Gov) ValidateStruct(bind interface{}, msgs ...map[string]string) (err e.Error) {
	rv := reflect.ValueOf(bind)
	kind := rv.Kind()
	switch kind {
	case reflect.Ptr:
		var data = map[string]interface{}{}
		var rules = Rules{}
		// 获取tag,解析规则
		rtelem := rv.Type().Elem()
		rvelem := rv.Elem()
		for i := 0; i < rtelem.NumField(); i++ {
			var field = rtelem.Field(i).Name
			// 组装data数据
			// 检查字段是否是零值
			fieldValue := rvelem.Field(i)
			if !fieldValue.IsZero() {
				data[field] = rvelem.Field(i).Interface()
			}

			// 解析tag, 组装rules规则
			tag := rtelem.Field(i).Tag.Get("gopass")
			if tag != "" {
				tagRules := strings.Split(strings.TrimRight(strings.TrimSpace(tag), g.RuleSep), g.RuleSep)
				if len(tagRules) > 0 {
					rules[field] = tagRules
				}
			}
		}
		// 筛选出 有required无数据的字段
		var rulesReal = Rules{}
		for key, item := range rules {
			if _, ok := data[key]; ok {
				rulesReal[key] = item
			} else {
				if len(item) > 0 && item[0] == "required" {
					rulesReal[key] = item
				}
			}
		}
		err = Validate(data, rulesReal, msgs...)
	default:
		err = e.New("请传入结构体")
	}
	return
}

// Validate 执行验证
func Validate(data map[string]interface{}, rules Rules, msgs ...map[string]string) (err e.Error) {
	return NewGov(nil).Validate(data, rules, msgs...)
}

// ValidateField 验证单一值
func ValidateField(value interface{}, rules []string, msg ...map[string]string) (err e.Error)  {
	return NewGov(nil).ValidateField(value, rules, msg...)
}

// ValidateStruct 绑定 struct 验证
func ValidateStruct(bind interface{}, msgs ...map[string]string) (err e.Error) {
	return NewGov(nil).ValidateStruct(bind, msgs...)
}
