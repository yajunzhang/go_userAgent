# go_userAgent
可以伪装生成headers请求头中的User Agent值

package main

import (
	"fmt"
	"github.com/yajunzhang/go_userAgent"
	"time"
)
var t1 = time.Now() 

func main() {
	ua := go_userAgent.NewUserAgent()
	fmt.Println(ua.Chrome())
	fmt.Println(time.Since(t1))
}
