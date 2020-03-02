package core

import (
	"fmt"
	"net/url"
	"strings"
	
)

func GetVulTypeListAndPara(targetUrl string,reqData string) (map[string][]int) {
	var VulType []int

	result := make(map[string][]int,30)


		if reqData == "" {

			u, err := url.Parse(targetUrl)

			if err != nil {
				fmt.Println(err)
			}

			queryValue := u.Query()

			if len(queryValue) >0 {

				for k, v := range queryValue {
					queryValue.Set(k,v[0]+Probe)
					u.RawQuery = queryValue.Encode()
					//fmt.Println(u.RequestURI())//debug
					VulType = DetectVulType(u.Scheme+"://"+u.Host+u.RequestURI(), reqData)
					result[k] = VulType
					queryValue.Set(k,v[0])
				}

			} else {//无参数case
				ustring := fmt.Sprintf("%s://%s%s",u.Scheme,u.Host,u.RequestURI())
				if string(ustring[len(ustring)-1]) != "/" {
					VulType = DetectVulType(ustring+"/"+Probe, reqData)


				} else {

					VulType = DetectVulType(ustring+Probe, reqData)
				}
				result["/"] = VulType
			}
		} else {
			reqData := reqData
			splitPostData := strings.Split(reqData,"&")
			for key, v := range splitPostData {
				splitPostData[key] += Probe
				postData := Implode("&",splitPostData)
				//fmt.Println(postData)
				VulType = DetectVulType(targetUrl,postData)
				result[strings.Split(v,"=")[0]] = VulType

				splitPostData[key] = v

			}
		}
		return result

}