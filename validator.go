package gopass

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

// Data ...
type Data map[string]interface{}

// Rules ...
type Rules map[string][]string

// RuleHandler ...
type RuleHandler func(data interface{}, rule ...string) error

// Options ...
type Options struct {
	msg Data
}

// OptionHandler ...
type OptionHandler func(options *Options)

// Validator ...
type Validator struct {
	rule *sync.Map
	option *Options
}

// ValidatorHandler ...
type ValidatorHandler func(validator *Validator)

var once sync.Once
var vd *Validator

// NewValidator ...
func NewValidator(opts ...OptionHandler) *Validator {
	once.Do(func() {
		vd = &Validator{rule:&sync.Map{},option:&Options{}}
	})
	for _, o := range opts {
		o(vd.option)
	}
	return vd
}

// Message ...
func Message(msg Data) OptionHandler {
	return func(o *Options) {
		o.msg = msg
	}
}

// Use ...
func (v *Validator) Use(opts ...ValidatorHandler) *Validator {
	for _, o := range opts {
		o(v)
	}
	return v
}

// Validate ...
func (v *Validator) Validate(data map[string]interface{}, rules Rules, msgs ...Data) error {
	if len(msgs) > 0 {
		v.option.msg = msgs[0]
	}
	// 验证rules
	//todo 开启多线程，同时验证多个规则
	for field, allRules := range rules {
		// 解析rules.val
		//todo 开启多线程，同时解析多个规则
		for _, rule := range allRules {
			// 如果有冒号，则取分割后的索引为0的数据作为验证对象
			var ruleReal = rule
			if strings.Contains(rule, ":") {
				ruleReals := strings.Split(rule, ":")
				ruleReal = ruleReals[0]
			}
			// 获取要验证的数据
			var dataReal interface{}
			if v2, ok := data[field]; ok {
				dataReal = v2
			}
			// 开始验证
			ruleHandler := v.Getter(ruleReal)
			if ruleHandler == nil {
				return errors.New(fmt.Sprintf("%s rule not exists", ruleReal))
			}
			err := ruleHandler(dataReal, rule)
			if err != nil {
				// 判断是否有自定义错误信息
				if v.option.msg != nil {
					// 返回自定义了错误信息
					if v, ok := v.option.msg[ruleReal]; ok {
						return errors.New(fmt.Sprintf("%s(%s)", v, field))
					}
				}
				// 返回默认错误信息
				return errors.New(fmt.Sprintf("%s(%s)", err.Error(), field))
			}
		}
	}
	return nil
}

// Register ...
func (v *Validator) Register(rule string, obj RuleHandler) {
	v.rule.Store(rule, obj)

	// 初始化的时候做一次getter测试
	// 如果有问题在初始化阶段就暴露出来
	validate := v.Getter(rule)
	if validate == nil {
		panic(fmt.Sprintf("规则(%s)注册失败：", rule))
	}
	var _ = validate("")
}

// Getter ...
func (v *Validator) Getter(rule string) RuleHandler {
	if value, ok := v.rule.Load(rule); ok {
		return value.(RuleHandler)
	}
	return nil
}
