package core

import "strings"

// 添加请求头
func AddHeaders() {
	// 处理请求头
	for _, line := range ReqHeaders {
		pair := strings.SplitN(line, ":", 2)
		if len(pair) == 2 {
			k, v := pair[0], strings.Trim(pair[1], " ")//处理值空格
			//之后配置headers的时候防止因为空格使请求头错误
			headers[k] = v
		}
	}
}
//添加cookie
func AddCookies()  {
	for _, line := range ReqCookies {
		pair := strings.SplitN(line,"=",2)
		if len(pair) == 2 {
			k, v := strings.Trim(pair[0]," "), pair[1]//用于处理key的空格
			cookies[k] = v
			//fmt.Println(cookies) 调试
		}
	}
}
