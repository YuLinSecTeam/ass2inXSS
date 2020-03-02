package core

import (
	"bufio"
	"fmt"
	"github.com/chromedp/chromedp"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"github.com/gookit/color"
	"time"
	"context"
)

//=====读取并处理响应

func FetchResponse(url string, reqData string) string {

	var req *http.Request
	var err error
	//请求方式处理
	if ReqMethod == http.MethodPost || (ReqMethod == "GET" && reqData != "") {
		req, err = http.NewRequest(http.MethodPost,url,strings.NewReader(reqData))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	} else {
		req, err = http.NewRequest(ReqMethod,url,nil)
	}

	if err != nil || req == nil {
		fmt.Println(err)
		return ""
	}
	//添加请求头，不包含cookie头
	for k, v := range headers {
		req.Header.Set(k,v)
	}
	//添加cookie头
	for k, v := range cookies {
		cookie := &http.Cookie{Name:k,Value:reflect.TypeOf(v).String(),}
		req.AddCookie(cookie)
	}
	//利用客户端获取响应
	client := SetClient()
	resp, err := client.Do(req)
	if err != nil {
		color.Red.Println(err)
		return ""
	}
	defer resp.Body.Close()
	//判断并处理响应编码
	reader := bufio.NewReader(resp.Body)
	e := determineEncoding(reader)
	utf8Reader := transform.NewReader(reader, e.NewDecoder())
	//此时，utf8Reader即为decode之后的响应体。
	body, _ := ioutil.ReadAll(utf8Reader)
	return string(body)



}


//== chromedp fetchres
func FetchResponse1(url string,reqdata string) string {
	var html string
	// ctxt是chromedp的实例，用于执行网页操作
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	err := chromedp.Run(ctx,
		//GetServiceList(&html)，
		chromedp.Navigate(url),
		chromedp.WaitReady("*", chromedp.ByQuery),
		chromedp.Sleep(2 * time.Second),
		chromedp.OuterHTML("*", &html, chromedp.ByQuery),
	)
	if err != nil {
		fmt.Println(err)// error handle
	}

	// 成功取得数据
	return html+reqdata
}
