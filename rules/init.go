package rules

import "github.com/gohouse/gopass/gopassRuler"

// defaultMessages 验证失败默认提示语
var defaultMessages = map[string]interface{}{
	// 必填参数: required
	gopassRuler.LangZHCN: map[string]interface{}{
		RequiredString:      "参数{field}为必须参数",
		BetweenString:       "参数{field}的值必须在({args})之间",
		LengthString:        "参数{field}的长度必须为{args}",
		MinString:           "参数{field}的值不能小于{args}",
		MaxString:           "参数{field}的值不能大于{args}",
		InString:            "参数{field}的值必须在给定的值({args})中",
		NumericString:       "参数{field}的值必须是数字",
		IntegerString:       "参数{field}的值必须是整数",
		FloatString:         "参数{field}的值必须是浮点数",
		UrlString:           "参数{field}的值必须是 url 地址",
		EmaillString:        "参数{field}的值必须是邮箱地址",
		ChineseMobileString: "参数{field}的值必须是中国大陆手机号",
		JsonString:          "参数{field}的值必须是 json 格式",
		XmlString:           "参数{field}的值必须是 xml 格式",
		DateString:          "参数{field}的值必须是日期格式",
		RegexpString:        "参数{field}的值不匹配",
	},
	gopassRuler.LangENUS: map[string]interface{}{
		RequiredString:      "param {field} needed",
		BetweenString:       "param {field} allow range ({args})",
		LengthString:        "param {field} length must {args}",
		MinString:           "param {field} must greater or equal than {args}",
		MaxString:           "param {field} must less or equal than {args}",
		InString:            "param {field} must in ({args})",
		NumericString:       "param {field} must be numeric",
		IntegerString:       "param {field} must be integer",
		FloatString:         "param {field} must be float",
		UrlString:           "param {field} must be a url",
		EmaillString:        "param {field} must be a email address",
		ChineseMobileString: "param {field} must be a chinese mobile number",
		JsonString:          "param {field} must be json format",
		XmlString:           "param {field} must be xml format",
		DateString:          "param {field} must be date format",
		RegexpString:        "param {field} value is not matched",
	},
	gopassRuler.LangZHTR: map[string]interface{}{
		RequiredString: "參數{field}不能為空",
		BetweenString:  "參數{field}必須在({args})之間",
	},
}

func init() {
	// 注册所有默认提示信息
	gopassRuler.RegisterMessageMulti(defaultMessages)
}

// GetDefaultMessages 获取默认错误提示语
func GetDefaultMessages() map[string]interface{} {
	return defaultMessages
}
