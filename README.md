# gopass
gopass 是一个简单优雅的golang参数验证器

## [English Readme](./README_en.md)

## 特点
- 支持常见的规则验证  
- 可以轻松的自定义验证规则  
- 可以轻松的自定义错误信息  
- 可以按需引入需要的规则  

## 安装
- go mod (vgo)
```sh
require github.com/fizzday/gopass master
```

- go get
```bash
go get github.com/fizzday/gopass
```

## 使用示例

```go
package main

import (
	"fmt"
	"github.com/gohouse/gopass"
	"github.com/gohouse/gopass/rules"
)

func main()  {
	data := gopass.Data{
		"mobile":1234567,
	}
	rule := gopass.Rules{
		"mobile":{"required","min:7","max:14","numberic"},
		"password":{"required","min:6","max:32"},
	}

	// 这里是按需引入
	//v := gopass.NewValidator().Use(rules.Required(),rules.Min(),rules.Max(),rules.Numberic())

	// 这里是引入默认的所有规则
	v := gopass.NewValidator().Use(rules.Default())
	err := v.Validate(data,rule)
	fmt.Println(err)
}
```

## 自定义验证规则

自定义一个规则叫`mustabc`，意思是，参数值必须为`abc`，否则验证不通过
```go
package main

import (
	"fmt"
	"github.com/gohouse/t"
	"github.com/gohouse/gopass"
	"github.com/gohouse/gopass/rules"
)

func main()  {
	data := gopass.Data{
		"mobile":1234567,
	}
	rule := gopass.Rules{
		"mobile":{"mustabc"},
	}
    
	v := gopass.NewValidator().Use(mustabc())
	err := v.Validate(data,rule)
	fmt.Println(err)
}


func mustabc() gopass.ValidatorHandler {
	return func(v *gopass.Validator) {
		v.Register("mustabc", func(data interface{}, rule ...string) error {
			// 只需要更改这里的验证逻辑即可
			if t.New(data).String() != "abc" {
				return errors.New("abc needed")
			}
			return nil
		})
	}
}
```
运行,会得到以下结果
```shell script
abc needed(mobile)
```

## 自定义错误信息
```go
package main

import (
	"fmt"
	"github.com/gohouse/gopass"
	"github.com/gohouse/gopass/rules"
)

func main()  {
	data := gopass.Data{
		"mobile":1234567,
	}
	rule := gopass.Rules{
		"mobile":{"required","min:7","max:14","numberic"},
		"password":{"required","min:6","max:32"},
	}

	// 自定义错误信息
	var msg = gopass.Data{"required":"参数缺失啊啊啊"}
	v := gopass.NewValidator(gopass.Message(msg)).Use(rules.Default())
	err := v.Validate(data,rule)
	fmt.Println(err)
}
```
执行，会输出
```shell script
参数缺失啊啊啊(password)
```