package core

import (
	"fmt"
	"regexp"
	"strings"
)

func DetectVulType(targetUrl string,reqData string) []int { //此时传入的参数是已经携带了probe的

	var targetUrlExtra string
	var reqDataExtra string
	var vulTypeList []int

	r := FetchResponse(targetUrl, reqData)
	//fmt.Println(r)//debug
	matches := ProbeRe.FindAllString(r,-1)
	matchesIndex0 := ProbeRe.FindStringIndex(r)
	probeAdd := [][]string{{`'`, `''`, `''`},{`"`, `""`, `""`},{`\'`, `'\\\\'`, `\\\\''`},
		{`\"`, `"\\\\"`, `\\\\""`},{`'">`,`'"&gt;`,`'">`},{`!<x`, `<x`, `!<x`},}
	//"'"sihuo'"" ''"sihuo'"'  "'"
	if len(matches) > 0 {

		subBody0 := SubStr1(r,matchesIndex0[0])
		openTagIndex := Stripos(subBody0,"<",0)
		closeTagIndex := Stripos(subBody0,">",0)
		if openTagIndex > 0 {
			openTagIndex = openTagIndex
			//fmt.Println(oCI)//debug
		} else {
			openTagIndex = 0xfffff-1
		}
		if closeTagIndex > 0 {
			closeTagIndex = closeTagIndex
			//fmt.Println(cCI)//debug
		} else {
			closeTagIndex = 0xfffff
		}

		//fmt.Println(matches[0]) //debug
		for i := 0; i < len(probeAdd); i++ {
			probeExtra := probeAdd[i][0] + Probe + probeAdd[i][0]
			if reqData == "" {
				targetUrlExtra = strings.ReplaceAll(targetUrl,Probe,probeExtra)
			} else {
				targetUrlExtra = targetUrl
				reqDataExtra = strings.ReplaceAll(reqData,Probe,probeExtra)
			}
			//fmt.Println(targetUrlExtra)//debug
			r = FetchResponse(targetUrlExtra,reqDataExtra)
			//fmt.Println(r) //debug
			probeLeft := probeAdd[i][1] + Probe
			probeRight := Probe + probeAdd[i][2]
			//fmt.Println(probeLeft,probeRight)
			probeLeftOrprobeRight := fmt.Sprintf("%s|%s",probeLeft,probeRight)
			probeLeftRe := regexp.MustCompile(probeLeft)
			probeRightRe := regexp.MustCompile(probeRight)
			probeLeftOrprobeRightRe := regexp.MustCompile(probeLeftOrprobeRight)
			matches = probeLeftOrprobeRightRe.FindAllString(r,-1)
			matchLeft := probeLeftRe.MatchString(r)
			matchRight := probeRightRe.MatchString(r)
			matchesIndex := matchesIndex0
			if len(matches) >0 {
				//fmt.Println(matches)//debug
				var oCI,cCI int
				subBody := SubStr1(r,matchesIndex[0])
				//fmt.Println(probeExtra,temp)debug
				openScriptIndex := Stripos(subBody,"<script",0)
				closeScriptIndex := Stripos(subBody,"</script",0)
				if openScriptIndex > 0 {
					oCI = openScriptIndex
					//fmt.Println(oCI)//debug
				} else {
					oCI = 0xfffff-1
				}
				if closeScriptIndex > 0 {
					cCI = closeScriptIndex
					//fmt.Println(cCI)//debug
				} else {
					cCI = 0xfffff
				}
				if oCI < cCI {
					//html context
					//要分是inside还是outside

						if openTagIndex < closeTagIndex {//outside
							if i == 5 {
								vulTypeList=append(vulTypeList,5)
							}
						} else {
							//inside
							//fmt.Println(probeAdd[i][0])
							if i == 4{
								if matchLeft {

									vulTypeList=append(vulTypeList,10)
								}
								if matchRight {

									vulTypeList=append(vulTypeList,10,11)
								}
							}
						}

				} else {
					//==js context
					scriptProbe := Probe + "</script>"
					scriptUrl := strings.Replace(targetUrl,Probe,scriptProbe,1)
					scriptUrl = strings.Replace(scriptUrl,"<","%3C",1)
					scriptUrl = strings.Replace(scriptUrl,">","%3E",1)
					x := FetchResponse(scriptUrl,reqData)
					scriptRe := regexp.MustCompile(scriptProbe)
					matches := scriptRe.FindAllString(x,-1)
					if len(matches) > 0 {
						if i < 4 {
							vulTypeList=append(vulTypeList,i,4)
						} else {
							vulTypeList=append(vulTypeList,4)
						}
					} else {
						vulTypeList=append(vulTypeList,i)
					}

				}
			}

		}
		return vulTypeList
	} else {
		return []int{}
	}
}
