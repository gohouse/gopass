package main

//import (
//	"fmt"
//	"github.com/gohouse/gopass"
//	"github.com/gohouse/gopass/gopassCors"
//	"github.com/valyala/fasthttp"
//	"log"
//	"reflect"
//)
//func main() {
//	log.Println("star serve: http://localhost:8899")
//	log.Fatal(fasthttp.ListenAndServe(":8899", func(ctx *fasthttp.RequestCtx) {
//		switch string(ctx.Path()) {
//		case "/foo":
//			fooHandlerFunc(ctx)
//		default:
//			ctx.Error("not found !!!", fasthttp.StatusNotFound)
//		}
//	}))
//}
//type Foo struct {
//	A string `gopass:"required|email|length:8"` // 必传参数,且是长度为8的email地址
//	B string `gopass:"required|date"`           // 必传参数,且是日期格式
//	C string `gopass:"required|between:3,6"`    // 必传参数,且是(3~6)之间
//	D string `gopass:"in:a,b"`                  // 可选参数,如果传了值,则必须是(a,b)中的任意一个
//}
//func fooHandlerFunc(ctx *fasthttp.RequestCtx) {
//	var foo Foo
//	err := gopassCors.BindFasthttp(ctx, &foo)
//	fmt.Fprintf(ctx, "param: a:%s, b:%s, c:%s \nvalide:%v", foo.A, foo.B, foo.C, err)
//}
//
//func BindFasthttp(ctx *fasthttp.RequestCtx, bind interface{}) error {
//	var args *fasthttp.Args
//	switch string(ctx.Method()) {
//	case "POST","PUT","PATCH":
//		args = ctx.PostArgs()
//	case "GET","DELETE":
//		args = ctx.QueryArgs()
//	default:
//		return nil
//	}
//	var dataTmp = map[string]string{}
//	args.VisitAll(func(key, value []byte) {
//		//log.Printf("%s,%s", key, value)
//		dataTmp[string(key)] = string(value)
//	})
//
//	// 开始绑定数据到struct
//	vof := reflect.ValueOf(bind).Elem()
//	tof := vof.Type()
//	for i := 0; i < vof.NumField(); i++ {
//		fieldName := tof.Field(i).Name
//		if v, ok := dataTmp[fieldName]; ok {
//			vof.FieldByName(fieldName).Set(reflect.ValueOf(v))
//		}
//	}
//
//	// 开始校验参数
//	return gopass.ValidateStruct(bind)
//}
