package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Object struct {
	Name string
}

func testAssert(obj1, obj2 interface{}, t *testing.T) {
	assert.Equal(t, obj1, obj2, "obj1 and obj2 should be equal")
}

func Test1(t *testing.T) {
	//str1 := "testing is cool"
	//str2 := "testing is really cool"
	// str3 := "testing is cool"
	//assert.Equal(t, str1, str2, "str1 and str2 should be equal")

	obj1 := &Object{
		Name: "Jack",
	}

	obj2 := &Object{
		Name: "Robin",
	}

	testAssert(obj1, obj2, t)
}
