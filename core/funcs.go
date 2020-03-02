package core

import (
	"net/url"
)

func HandleTargetUrl(target string) string {
	/*lastChar := string(url[len(url)-1])
	if lastChar != "/" {
		url = url + "/"
	}
	return url+Probe*/

	u, _ := url.Parse(target)
	if ReqData == "" {
		if len(u.Query()) != 0 {
			return target
		} else {
			if string(target[len(target)-1]) != "/" {
				target = target+"/"
			}
			return target+Probe
		}
	} else {
		return target
	}

}














