package core

import "github.com/gookit/color"

func PrintVulInfo(result map[string][]int)  {
	flag := 0
	for k, v := range result {
		if len(v) > 0 {
			for _, value := range v {
				color.Red.Printf(`
Context: %s
VulPara: %s
VerboseInfo: %s

`,VulContext[value],k,VulVerboseInfo[value])
			}
		} else {
			flag++
		}
	}
	if len(result) == flag {
		color.Green.Println("Aho, no xss found")
	}

}
