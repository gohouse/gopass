# gopass
a simple and beautiful validator for golang

## feature
Supports common validation rules, also，You can easily add new validation rules.   

## install
- go mod (vgo)
```sh
require github.com/fizzday/gopass master
```

- go get
```bash
go get github.com/fizzday/gopass
```

## usage

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
    
	// here you can only import some rules you need
	//v := gopass.NewValidator().Use(rules.Required(),rules.Min(),rules.Max(),rules.Numberic())

	// here you can import all default rules
	v := gopass.NewValidator().Use(rules.Default())
	err := v.Validate(data,rule)
	fmt.Println(err)
}
```

## Define validation rules yourself 

Customize a rule called `mustabc`, which means that the parameter value must be` abc`, otherwise the verification fails
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
			if t.New(data).String() != "abc" {
				return errors.New("abc needed")
			}
			return nil
		})
	}
}
```
Running it, you will get the following results
```shell script
abc needed(mobile)
```