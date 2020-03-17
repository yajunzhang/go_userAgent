# go_userAgent
可以伪装生成headers请求头中的User Agent值

```
// 安装
go get github.com/yajunzhang/go_userAgent
```

```go
package main

import (
	"fmt"
	"github.com/yajunzhang/go_userAgent"
	"time"
)
var t1 = time.Now() 

func main() {
	ua := go_userAgent.NewUserAgent()
	// 随机生成
	fmt.Println(ua.Random())
	// Chrome 浏览器
	fmt.Println(ua.Chrome())
	// Safari 浏览器
	fmt.Println(ua.Safari())
	// 火狐浏览器
	fmt.Println(ua.Firefox())
	// Opera 浏览器
	fmt.Println(ua.Opera())
	
	fmt.Println(time.Since(t1))
}

```
