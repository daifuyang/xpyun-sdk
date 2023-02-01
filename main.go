/**
** @创建时间: 2021/4/11 12:52 上午
** @作者　　: return
** @描述　　:
 */
package xpyunSdk

import "reflect"

type baseOptions struct {
	User string `json:"user"`
	Ukey string `json:"ukey"`
	Sn   string `json:"sn"`
}

var options *baseOptions

func NewOptions(params map[string]string) baseOptions {
	options = &baseOptions{
		User: params["user"],
		Ukey: params["ukey"],
	}
	return *options
}

func SetOption(key string, val string) {
	oPoint := reflect.ValueOf(options)
	field := oPoint.Elem().FieldByName(key)
	field.SetString(val)
}

func Options() baseOptions {
	if options == nil {
		panic("配置未初始化！")
	}
	return *options
}
