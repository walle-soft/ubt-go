package log

import (
	"fmt"
	"github.com/guonaihong/gout"
	"log"
	"os"
)

var logUrl string

type H = gout.H

func base(logLevel string, jsonData H)  {
	err := gout.
		// 设置POST方法和url
		POST(logUrl).
		//打开debug模式
		Debug(false).
		SetHeader(gout.H{"LogLevel": logLevel}).
		// 设置非结构化数据到http body里面
		// 设置json需使用SetJSON
		SetJSON(jsonData).
		//结束函数
		Do()

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func SendLog(jsonData H) {
	base("log", jsonData)
}
func SendError(jsonData H)  {
	base("error", jsonData)
}

func init()  {
	logUrl = os.Getenv("UBT_URL")
	if logUrl == "" {
		log.Fatal("[ubt-go]请在环境变量中提供日志url, 环境变量字段名: UBT_URL")
	}
}