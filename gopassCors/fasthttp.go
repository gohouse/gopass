package gopassCors

//import (
//	"github.com/gin-gonic/gin"
//	"github.com/gohouse/gopass"
//	"github.com/valyala/fasthttp"
//)
//
//func BindGin(ctx *gin.Context, bind interface{}) error {
//	var err error
//	switch ctx.Request.Method {
//	case "POST", "PUT", "PATCH":
//		err = ctx.Bind(bind)
//	case "GET", "DELETE":
//		err = ctx.BindQuery(bind)
//	default:
//		return nil
//	}
//	if err != nil {
//		return err
//	}
//
//	// 开始校验参数
//	return gopass.ValidateStruct(bind)
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