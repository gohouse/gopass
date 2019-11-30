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

// Validator ...
type Validator struct {
	sync.Map
}

// ValidatorHandler ...
type ValidatorHandler func(validator *Validator)

var once sync.Once
var vd *Validator

// NewValidator ...
func NewValidator() *Validator {
	once.Do(func() {
		vd = &Validator{}
	})
	return vd
}

// Use ...
func (v *Validator) Use(opts ...ValidatorHandler) *Validator {
	for _, o := range opts {
		o(v)
	}
	return v
}

// Validate ...
func (v *Validator) Validate(data map[string]interface{}, rules Rules, msgs ...interface{}) error {
	// 检查rules
	for key, val := range rules {
		// 解析rules.val
		for _, rule := range val {
			// 如果有冒号，则取分割后的索引为0的数据作为验证对象
			var ruleReal = rule
			if strings.Contains(rule, ":") {
				ruleReals := strings.Split(rule, ":")
				ruleReal = ruleReals[0]
			}
			// 获取要验证的数据
			var dataReal interface{}
			if v2, ok := data[key]; ok {
				dataReal = v2
			}
			// 开始验证
			rh := v.Getter(ruleReal)
			if rh == nil {
				return errors.New(fmt.Sprintf("%s验证规则不存在", ruleReal))
			}
			err := rh(dataReal, rule)
			if err != nil {
				return errors.New(fmt.Sprintf("%s(%s)", err.Error(), key))
			}
		}
	}
	return nil
}

// Register ...
func (v *Validator) Register(rule string, obj RuleHandler) {
	v.Store(rule, obj)

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
	if value, ok := v.Load(rule); ok {
		return value.(RuleHandler)
	}
	return nil
}
