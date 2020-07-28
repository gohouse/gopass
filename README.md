## 1. 关于gopass
gopass 是一个简单优雅的golang参数验证器

## 2. 特点
- 支持国际化多语言验证  
- 支持常见的规则验证  
- 可以轻松的自定义验证规则  
- 可以轻松的自定义错误信息  
- 支持struct绑定验证和自定义map验证  
- 支持fasthttp,gin等web框架参数绑定自动验证  

## 3. 使用示例
### 3.1 数据直接验证
```go
package main

import (
	"fmt"
	"github.com/gohouse/gopass"
)

func main() {
	var data = "13800138000"
	err := gopass.ValidateField(data, []string{gopass.Required(),gopass.ChineseMobile()})
	fmt.Println(err)
}
```

### 3.2 `struct`绑定验证示例
```go
package main

import (
	"fmt"
	"github.com/gohouse/gopass"
	//_ "github.com/gohouse/gopass/rules"
)

func main() {
	structDemo()
}

type User struct {
	Mobile   string `gopass:"required|numeric|length:11"` // 手機
	Password string `gopass:"required|length:6,16"`       // 密碼demo.go
	Email    string `gopass:"required|email"`             // Email
	Sex      string `gopass:"in:man,woman,other"`         // sex
}

func structDemo() {
	var u = User{
		Mobile:   "13800138000",
		Password: "123456",
		Email:    "aa@qq.com",
		Sex:      "xxx",
	}
	err := gopass.ValidateStruct(&u)
	if err != nil {
		fmt.Println("get err:", err.Error())
	} else {
		fmt.Println("struct test pass")
	}
}
```

### 3.3 `map`验证
```go
package main

import (
	"fmt"
	"github.com/gohouse/gopass"
)

func main() {
	var data = map[string]interface{}{
		"myKey":  "5",
		"fielda": "5",
	}
	err := gopass.Validate(data, gopass.Rules{
		"fielda": {"required", "between:1,5", "integer", "in:1,a,5"},
                // 方法快捷构建规则, 与上边是等效的
		"myKey":  {gopass.Required(), gopass.Between(1, 5), gopass.Integer(), gopass.In(1, "a", 5)},
		"test":   {gopass.Required()},
	})

	fmt.Println(err)
}
```

## 4. 自定义验证规则

自定义一个规则叫`mustabc`，意思是，参数值必须为`abc`，否则验证不通过

```go
package main

import (
	"fmt"
	"github.com/gohouse/gopass"
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

func main() {
	err := gopass.ValidateField("ABCd", []string{"mustABC"})

	fmt.Println(err)
}

func init() {
	// 注册 required 验证器
	gopassRuler.Register("mustABC", &MinValidator{})
}
// MinValidator 参数最小值验证器
type MinValidator struct{}
// Validate 实现当前规则的验证接口
func (rv *MinValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if t.New(data).String() != "ABC" {
		return e.New("参数值必须为 ABC")
	}
	return nil
}
```

运行,会得到以下结果

```shell script
参数值必须为 ABC
```

## 5. 自定义错误信息,以及国际化多语言支持
### 5.1 初始化自定义错误信息,全局有效  
```go
gp := gopass.NewGopass(&gopass.Option{DefaultMessage:map[string]interface{}{
    "zh_cn":map[string]string{
        "numberic":"参数{field}值必须为数字"
    },
    "en_us":map[string]string{
        "required":"param must be numberic, but {args} given"
    },
}})
gp.Validate(xxx)
```
- `{field}`为字段名字占位符,如果传入则会自动替换为实际字段名字  
- `{args}`为传入的具体参数占位符,如果传入则会自动替换为实际参数值    

### 5.2 验证时,临时定义错误信息,只作用于当前验证规则
```go
gopass.Validate(data, gopass.Rules{}, map[string]string{"required":"参数值还是必要的"})
```
该错误信息会覆盖默认的`required`提示信息

### 5.3 国际化多语言支持  
具体参考代码 [examples/i18n/demo2.go](./examples/i18n/demo2.go)

## 6. gopass提供的默认验证规则列表
规则 | 说明 | 示例  
--- | --- | ---  
require | 必传参数验证 | "required"  
between | 验证数字是否在给定范围之内 | "between:6,32"  
chinese_mobile | 中国手机号验证 | "chinese_mobile"  
date | 必须是日期格式(yyyy-mm-dd) | "date"  
datetime | 必须是时间格式({date} HH:ii:ss) | "datetime"  
email | 必须是邮件格式 | "email"  
float | 必须是浮点数 | "float"  
in | 在给定的值中(等同于enum) | "in:a,b,c"  
integer | 必须是整数 | "integer"  
json | 必须是json格式 | "json"  
length | 数据长度范围 | "length:8" 或者 "length:5,10"  
max | 最大值 | "max"  
min | 最小值 | "min"  
numeric | 必须为数字 | "numeric"  
regexp | 正则 | "regexp:^\w+\d*$"  
url | 必须是网址 | "url"  
xml | 必须是xml格式 | "xml"  


## 7. 在web框架中使用
### 7.1 gin
```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gopass"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", fooHandlerFunc)
	r.Run(":8080")
}

type Foo struct {
	A string `gopass:"required|email|length:8"` // 必传参数,且是长度为8的email地址
	B string `gopass:"required|date"`           // 必传参数,且是日期格式
	C string `gopass:"required|between:3,6"`    // 必传参数,且是(3~6)之间
	D string `gopass:"in:a,b"`                  // 可选参数,如果传了值,则必须是(a,b)中的任意一个
}

func fooHandlerFunc(ctx *gin.Context) {
	var foo Foo
	err := BindGin(ctx, &foo)
	ctx.String(http.StatusOK, "param: a:%s, b:%s, c:%s \nvalide:%v", foo.A, foo.B, foo.C, err)
}

func BindGin(ctx *gin.Context, bind interface{}) error {
	var err error
	switch ctx.Request.Method {
	case "POST", "PUT":
		err = ctx.Bind(bind)
	case "GET", "DELETE":
		err = ctx.BindQuery(bind)
	default:
		return nil
	}
	if err != nil {
		return err
	}

	// 开始校验参数
	return gopass.ValidateStruct(bind)
}
```

### 7.2 fasthttp
```go
package main

import (
	"fmt"
	"github.com/gohouse/gopass"
	"github.com/gohouse/gopass/gopassCors"
	"github.com/valyala/fasthttp"
	"log"
	"reflect"
)
func main() {
	log.Println("star serve: http://localhost:8899")
	log.Fatal(fasthttp.ListenAndServe(":8899", func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/foo":
			fooHandlerFunc(ctx)
		default:
			ctx.Error("not found !!!", fasthttp.StatusNotFound)
		}
	}))
}
type Foo struct {
	A string `gopass:"required|email|length:8"` // 必传参数,且是长度为8的email地址
	B string `gopass:"required|date"`           // 必传参数,且是日期格式
	C string `gopass:"required|between:3,6"`    // 必传参数,且是(3~6)之间
	D string `gopass:"in:a,b"`                  // 可选参数,如果传了值,则必须是(a,b)中的任意一个
}
func fooHandlerFunc(ctx *fasthttp.RequestCtx) {
	var foo Foo
	err := gopassCors.BindFasthttp(ctx, &foo)
	fmt.Fprintf(ctx, "param: a:%s, b:%s, c:%s \nvalide:%v", foo.A, foo.B, foo.C, err)
}

func BindFasthttp(ctx *fasthttp.RequestCtx, bind interface{}) error {
	var args *fasthttp.Args
	switch string(ctx.Method()) {
	case "POST","PUT":
		args = ctx.PostArgs()
	case "GET","DELETE":
		args = ctx.QueryArgs()
	default:
		return nil
	}
	var dataTmp = map[string]string{}
	args.VisitAll(func(key, value []byte) {
		//log.Printf("%s,%s", key, value)
		dataTmp[string(key)] = string(value)
	})

	// 开始绑定数据到struct
	vof := reflect.ValueOf(bind).Elem()
	tof := vof.Type()
	for i := 0; i < vof.NumField(); i++ {
		fieldName := tof.Field(i).Name
		if v, ok := dataTmp[fieldName]; ok {
			vof.FieldByName(fieldName).Set(reflect.ValueOf(v))
		}
	}

	// 开始校验参数
	return gopass.ValidateStruct(bind)
}
```
