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
	podName = "yyk"
	fmt.Println(GetDpNameByPodName(podName))
	if GetDpNameByPodName(podName) != "yyk" {
		fmt.Println("err")
		return
	}
	svcName = "yyk"
	fmt.Println(GetDpNameBySvcName(svcName))
	if GetDpNameBySvcName(svcName) != "yyk" {
		fmt.Println("err")
		return
	}

	fmt.Println("--group3--")
	podName = "yyk-notice"
	fmt.Println(GetDpNameByPodName(podName))
	if GetDpNameByPodName(podName) != "yyk-notice" {
		fmt.Println("err")
		return
	}
	svcName = "yyk-svc"
	fmt.Println(GetDpNameBySvcName(svcName))
	if GetDpNameBySvcName(svcName) != "yyk" {
		fmt.Println("err")
		return
	}

	fmt.Println("--group4--")
	podName = "yyk-notice-service"
	fmt.Println(GetDpNameByPodName(podName))
	if GetDpNameByPodName(podName) != "yyk" {
		fmt.Println("err")
		return
	}
	svcName = "yyk-notice-svc"
	fmt.Println(GetDpNameBySvcName(svcName))
	if GetDpNameBySvcName(svcName) != "yyk-notice" {
		fmt.Println("err")
		return
	}

	fmt.Println("--group5--")
	podName = "yyk-notice-1231231-1231231"
	fmt.Println(GetDpNameByPodName(podName))
	if GetDpNameByPodName(podName) != "yyk-notice" {
		fmt.Println("err")
		return
	}
	svcName = "yyk-notice-service-svc"
	fmt.Println(GetDpNameBySvcName(svcName))
	if GetDpNameBySvcName(svcName) != "yyk-notice-service" {
		fmt.Println("err")
		return
	}

	fmt.Println("--group6--")
	podName = "yyk-no-ser-hess-1765227074-3yuzk"
	fmt.Println(GetDpNameByPodName(podName))
	if GetDpNameByPodName(podName) != "yyk-no-ser-hess" {
		fmt.Println("err")
		return
	}
	svcName = "yyk-no-ser-hess-svc"
	fmt.Println(GetDpNameBySvcName(svcName))
	if GetDpNameBySvcName(svcName) != "yyk-no-ser-hess" {
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
