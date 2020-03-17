package go_userAgent

import (
	"os"
	"math/rand"
	"time"
)

var cachePath string = "./.cache/"
var cache []map[int]string
var fileConf = []string{"Chrome", "Safari", "Opera", /* "Android+Webkit+Browser"*/}

func init(){
	// 创建缓存目录
	os.MkdirAll(cachePath, os.ModePerm)

	for _, name := range fileConf{
		_, err := os.Stat(cachePath+name) 
		if !(err == nil || os.IsExist(err)) { // 不存在，下载
			Get(name)
		}
	}

	rand.Seed(time.Now().Unix())
}

type UserAgent struct{}

func NewUserAgent() *UserAgent {
	return &UserAgent{}
}

// 随机
func (ua *UserAgent)Random() string {
	m := Read(fileConf[rand.Intn(len(fileConf))])
	return m[rand.Intn(len(m))]
}

func (ua *UserAgent)Chrome() string {
	m := Read("Chrome")
	return m[rand.Intn(len(m))]
}

func (ua *UserAgent)Safari() string {
	m := Read("Safari")
	return m[rand.Intn(len(m))]
}

func (ua *UserAgent)Firefox() string {
	m := Read("Firefox")
	return m[rand.Intn(len(m))]
}

func (ua *UserAgent)Opera() string {
	m := Read("Opera")
	return m[rand.Intn(len(m))]
}

func (ua *UserAgent)Android() string {
	return ""
}


