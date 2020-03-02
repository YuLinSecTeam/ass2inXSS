package core

import "regexp"

var (
	Probe	= "qpskdjcn" //探测值
	ProbeRe = regexp.MustCompile(Probe)

	Context string

	TargetUrl string //目标url
	xss = 0 //测试flag
	//请求处理
	ReqMethod string //方法
	ReqData string //请求数据

	ReqHeaders Headers //请求头
	ReqCookies Cookies
	headers    = map[string]string{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36",
	}//请求头
	cookies = map[string]interface{}{

	}


	Timeout int
	//echo漏洞详细信息
	VulVerboseInfo = map[int]string{
		0:"' can use in js context",
		1:"\" can use in js context",
		2:"\\' can use in  js context",
		3:"\\\" can use in js context",
		4:"reflect in js" + "you can use </script> to break out js context",
		5:"use normal html tag",
		10:"' or \" can use in html tag",
		11:"'> or \"> can break out html tag",
	}
	VulContext = map[int]string{
		0:"script",
		1:"script",
		2:"script",
		3:"script",
		4:"scripto",
		5:"html",
		10:"htmli",
		11:"htmlo",
	}
	//探测filter
	/*
	<sihuo// <sihuo>
	<sihuo
	<sihuo//

	FuzzersHead = {"<sihuo","xx"}
	FuzzersBetTagAndEvent = {"%20","/","%09","%09%09","%0a","%0d","/s/"}
	FuzzersEvent = {"x","x=","x=y","onx=y"}
	FuzzersTail = {"%20","//",">","%0a","%0d","%09"}

	*/



)

