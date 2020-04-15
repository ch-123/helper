package helper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func Json(data interface{}) string {
	tmp, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "json_encode error"
	}
	return string(tmp)
}

//获取时间戳  单位微秒
func Time() int64 {
	return time.Now().UnixNano() / 1000000
}

//延迟毫秒
func Sleep(m time.Duration) {
	time.Sleep(m * time.Second / 1000)
}

//空函数
func Nullf(d interface{}) {}

//打印数据
func Print(data interface{}) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println(t.Name())
		fmt.Println(v)
		return
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f)
	}
	fmt.Println(v)
}
