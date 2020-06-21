package utils

import (
	"github.com/guonaihong/gout"
	User "os/user"
	"strings"
	"time"
)

func GetIP() (ip string) {
	data := ""
	err := gout.
		GET("http://119.29.29.29/d?dn=qq.com&clientip=1").
		SetTimeout(10*time.Second).
		Debug(false).
		BindBody(&data).
		Do()
	if err != nil {
		return ip
	}
	ip = strings.Split(data, "|")[1]
	return ip
}

func GetComputeUserName() (userName string) {
	user, err := User.Current()
	if err == nil {
		userName = user.Username
	}
	return userName
}