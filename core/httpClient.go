package core

import (
	"crypto/tls"
	"net/http"
	"time"
)

func SetClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig:&tls.Config{InsecureSkipVerify:true},
	}//跳过证书验证

	client := &http.Client{
		Timeout:time.Duration(Timeout) * time.Second,
		Transport:tr,
	}//建立http客户端
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}//重写checkredirect函数，强制让跳转报错并返回上一个响应
	return client

}
