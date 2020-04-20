package helper

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"strings"
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

//生成md5
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//随机字符串
func RandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzQWERTYUIOPASDFGHJKLZXCVBNM"
	bytes := []byte(str)
	slen := len(bytes)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(slen)])
	}
	return string(result)
}

//读取文件
func ReadFile(fileName string) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return []byte{}, err
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	return fd, err
}

//写入文件
func WriteFile(fileName string, data []byte) error {
	//创建目录
	if err := os.MkdirAll(filepath.Dir(fileName)+`/`, 0666); err != nil {
		return err
	}
	//写入文件
	err := ioutil.WriteFile(fileName, data, 0666)
	return err
}

//追加写入文件
func WriteFileAppend(fileName string, data []byte) error {
	var f *os.File
	var err error
	//创建目录
	if err = os.MkdirAll(filepath.Dir(fileName), 0666); err != nil {
		return err
	}
	if _, err := os.Stat(fileName); err == nil { //文件存在
		f, err = os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
	} else { //文件不存在
		f, err = os.Create(fileName) //创建文件
	}
	if err != nil {
		return err
	}
	//将文件写进去
	_, err = f.Write(data) //追加
	return err
}

//获取byte中某个字符的索引
func ByteIndex(arr []byte, s byte) int {
	for i, v := range arr {
		if v == s {
			return i
		}
	}
	return -1
}

//读取key=value类型的配置文件
func GetConfig(path string) (map[string]string, error) {
	config := make(map[string]string)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return config, err
	}
	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config, nil
}

//拼接byte
func ByteJoin(b ...[]byte) []byte {
	var buf bytes.Buffer
	for _, v := range b {
		buf.Write(v)
	}
	return buf.Bytes()
}
