package helper

import (
	"encoding/json"
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

func Sleep(m time.Duration) {
	time.Sleep(m * time.Second)
}
