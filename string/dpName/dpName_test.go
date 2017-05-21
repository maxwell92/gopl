package main

import (
	"fmt"
	"testing"
)

func Test_GetDpName(t *testing.T) {
	var podName string
	var svcName string

	fmt.Println("--group1--")
	podName = ""
	fmt.Println(GetDpNameByPodName(podName))
	if GetDpNameByPodName(podName) != "" {
		fmt.Println("err")
		return
	}
	svcName = ""
	fmt.Println(GetDpNameBySvcName(svcName))
	if GetDpNameBySvcName(svcName) != "" {
		fmt.Println("err")
		return
	}

	fmt.Println("--group2--")
	podName = "abc"
	fmt.Println(GetDpNameByPodName(podName))
	if GetDpNameByPodName(podName) != "abc" {
		fmt.Println("err")
		return
	}
	svcName = "abc"
	fmt.Println(GetDpNameBySvcName(svcName))
	if GetDpNameBySvcName(svcName) != "abc" {
		fmt.Println("err")
		return
	}

	fmt.Println("--group3--")
	podName = "abc-1231111"
	fmt.Println(GetDpNameByPodName(podName))
	if GetDpNameByPodName(podName) != "abc-1231111" {
		fmt.Println("err")
		return
	}
	svcName = "abc-svc"
	fmt.Println(GetDpNameBySvcName(svcName))
	if GetDpNameBySvcName(svcName) != "abc" {
		fmt.Println("err")
		return
	}

	fmt.Println("--group4--")
	podName = "abc-1231231-12313"
	fmt.Println(GetDpNameByPodName(podName))
	if GetDpNameByPodName(podName) != "abc" {
		fmt.Println("err")
		return
	}
	svcName = "abc-def-svc"
	fmt.Println(GetDpNameBySvcName(svcName))
	if GetDpNameBySvcName(svcName) != "abc-def" {
		fmt.Println("err")
		return
	}

	fmt.Println("--group5--")
	podName = "abc-notice-1231231-1231231"
	fmt.Println(GetDpNameByPodName(podName))
	if GetDpNameByPodName(podName) != "abc-notice" {
		fmt.Println("err")
		return
	}
	svcName = "abc-123123-abc-svc"
	fmt.Println(GetDpNameBySvcName(svcName))
	if GetDpNameBySvcName(svcName) != "abc-123123-abc" {
		fmt.Println("err")
		return
	}
}

/*
func Test_GetDpNameBySvcName(t *testing.T) {
	var svcName string

	svcName = ""
	fmt.Println(GetDpNameBySvcName(svcName))

	svcName = "abc"
	fmt.Println(GetDpNameBySvcName(svcName))

	svcName = "abc-svc"
	fmt.Println(GetDpNameBySvcName(svcName))

	svcName = "abc-def-svc"
	fmt.Println(GetDpNameBySvcName(svcName))

	svcName = "abc-def-efg-svc"
	fmt.Println(GetDpNameBySvcName(svcName))
}
*/
