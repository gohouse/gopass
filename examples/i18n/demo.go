package main

import (
	"fmt"
	"github.com/gohouse/gopass"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/gopass/rules"
	//_ "github.com/gohouse/gopass/rules"
)

func main() {
	demo()
	demo2()
}

// demo 自定义错误消息
func demo() {
	var data = map[string]interface{}{
		"myKey":  "5",
		"fielda": "5",
	}
	v := gopass.NewGov(&gopass.Option{
		DefaultLanguage: "en_us",
		DefaultMessage:  GetDefaultMessages(),
	})
	err := v.Validate(data, gopass.Rules{
		"fielda": {"required", "between:1,5", "integer", "in:1,a,5"},
		"myKey":  {gopass.Required(), gopass.Between(1, 5), gopass.Integer(), gopass.In(1, "a", 5)},
		"test":   {gopass.Required()},
	})

	fmt.Println(err)
}

// defaultMessages 验证失败默认提示语
var defaultMessages = map[string]interface{}{
	// 必填参数: required
	gopassRuler.LangZHCN: map[string]interface{}{
		rules.RequiredString:      "参数{field}为必须参数",
		rules.BetweenString:       "参数{field}的值必须在({args})之间",
		rules.LengthString:        "参数{field}的长度必须为{args}",
		rules.MinString:           "参数{field}的值不能小于{args}",
		rules.MaxString:           "参数{field}的值不能大于{args}",
		rules.InString:            "参数{field}的值必须在给定的值({args})中",
		rules.NumericString:       "参数{field}的值必须是数字",
		rules.IntegerString:       "参数{field}的值必须是整数",
		rules.FloatString:         "参数{field}的值必须是浮点数",
		rules.UrlString:           "参数{field}的值必须是 url 地址",
		rules.EmaillString:        "参数{field}的值必须是邮箱地址",
		rules.ChineseMobileString: "参数{field}的值必须是中国大陆手机号",
		rules.JsonString:          "参数{field}的值必须是 json 格式",
		rules.XmlString:           "参数{field}的值必须是 xml 格式",
		rules.DateString:          "参数{field}的值必须是日期格式",
		rules.RegexpString:        "参数{field}的值不匹配",
	},
	gopassRuler.LangENUS: map[string]interface{}{
		rules.RequiredString:      "param {field} needed",
		rules.BetweenString:       "param {field} allow range ({args})",
		rules.LengthString:        "param {field} length must {args}",
		rules.MinString:           "param {field} must greater or equal than {args}",
		rules.MaxString:           "param {field} must less or equal than {args}",
		rules.InString:            "param {field} must in ({args})",
		rules.NumericString:       "param {field} must be numeric",
		rules.IntegerString:       "param {field} must be integer",
		rules.FloatString:         "param {field} must be float",
		rules.UrlString:           "param {field} must be a url",
		rules.EmaillString:        "param {field} must be a email address",
		rules.ChineseMobileString: "param {field} must be a chinese mobile number",
		rules.JsonString:          "param {field} must be json format",
		rules.XmlString:           "param {field} must be xml format",
		rules.DateString:          "param {field} must be date format",
		rules.RegexpString:        "param {field} value is not matched",
	},
	gopassRuler.LangZHTR: map[string]interface{}{
		rules.RequiredString: "參數{field}不能為空",
		rules.BetweenString:  "參數{field}必須在({args})之間",
	},
}

// GetDefaultMessages 获取默认错误提示语
func GetDefaultMessages() map[string]interface{} {
	return defaultMessages
}
