package main

import (
	"fmt"
	"strings"
)

func GetDpNameByPodName(podName string) string {
	frag := strings.Split(podName, "-")
	length := len(frag)

	if length > 2 {
		return strings.Join(frag[:length-2], "-")
	} else {
		return podName
	}
}

func GetDpNameBySvcName(svcName string) string {
	frag := strings.Split(svcName, "-")
	length := len(frag)

	if length > 1 {
		return strings.Join(frag[:length-1], "-")
	} else {
		return svcName
	}
}

func main() {
	var dpName1 string
	// svcName := "abcv-notive-bus-svc"
	svcName := "sbc-svc"
	svcFrag := strings.Split(svcName, "-")
	svcLen := len(svcFrag)
	/*
		for i := 0; i < svcLen-1; i++ {
			if i == svcLen-2 {
				dpName1 = dpName1 + svcFrag[i]
				break
			}
			dpName1 = dpName1 + svcFrag[i] + "-"

		}
	*/
	if svcLen > 1 {
		dpName1 = strings.Join(svcFrag[:svcLen-1], "-")
	} else {
		dpName1 = svcName
	}
	fmt.Printf("dpName1: %s\n", dpName1)

	var dpName2 string
	// gpodName := "abcv-notive-bus-10212312351-185821"
	podName := "sbc-pod"
	podFrag := strings.Split(podName, "-")
	podLen := len(podFrag)
	/*
		for i := 0; i < podLen-2; i++ {
			if i == podLen-3 {
				dpName2 = dpName2 + podFrag[i]
				break
			}
			dpName2 = dpName2 + podFrag[i] + "-"

		}
	*/
	if podLen > 1 {
		dpName2 = strings.Join(podFrag[:podLen-2], "-")
	} else {
		dpName2 = podName
	}
	fmt.Printf("dpName2: %s\n", dpName2)
}
