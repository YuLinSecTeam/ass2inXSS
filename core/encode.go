package core

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

func determineEncoding(r *bufio.Reader) encoding.Encoding {//判断字符编码
	b, err := r.Peek(1024)
	if err != nil {
		// log.Error("get code error")
		return unicode.UTF8 //特别容易出错，出错就返回utf-8，能解决大部分问题
	}
	e, _, _ := charset.DetermineEncoding(b, "")
	return e
}
