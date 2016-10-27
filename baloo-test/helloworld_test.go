package main

import (
	"gopkg.in/h2non/baloo.v0"
	"testing"
)

var test = baloo.New("http://httpbin.org")

func TestSay(t *testing.T) {
	test.Get("/get").
		SetHeader("Foo", "Bar").
		Expect(t).
		Status(200).
		Header("Server", "nginx").
		Type("json").
		JSON(map[string]string{"bar": "foo"}).
		Done()
}
