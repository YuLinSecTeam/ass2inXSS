package main

import (
	"flag"
	"github.com/YuLinSecTeam/ass2inXSS/core"
)


func init()  { //设置flag
	flag.StringVar(&core.ReqMethod,"X","GET","-X GET")
	flag.StringVar(&core.TargetUrl,"url","","-u example.com")
	flag.Var(&core.ReqHeaders, "H", `-H "x-forward-for:127.0.0.1"`)
	flag.StringVar(&core.ReqData,"D","","-D \"a=1&b=2\"")
	flag.IntVar(&core.Timeout,"timeout",5,"-timeout 2")
	flag.Var(&core.ReqCookies,"C","-C a=1;b=2")
	flag.Parse()
}


func main()  {
	result := core.GetVulTypeListAndPara(core.TargetUrl,core.ReqData)
	//fmt.Println(result)
	core.PrintVulInfo(result)

}
