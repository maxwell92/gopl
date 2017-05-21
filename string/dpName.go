package main

import (
	"fmt"
	"strings"
)

func main() {
	var dpName1 string
	svcName := "abcv-notive-bus-svc"
	svcFrag := strings.Split(svcName, "-")
	svcLen := len(svcFrag)
	for i := 0; i < svcLen-1; i++ {
		if i == svcLen-2 {
			dpName1 = dpName1 + svcFrag[i]
			break
		}
		dpName1 = dpName1 + svcFrag[i] + "-"

	}
	fmt.Printf("dpName1: %s\n", dpName1)

	var dpName2 string
	podName := "abcv-notive-bus-10212312351-185821"
	podFrag := strings.Split(podName, "-")
	podLen := len(podFrag)
	for i := 0; i < podLen-2; i++ {
		if i == podLen-3 {
			dpName2 = dpName2 + podFrag[i]
			break
		}
		dpName2 = dpName2 + podFrag[i] + "-"

	}
	fmt.Printf("dpName2: %s\n", dpName2)
}
