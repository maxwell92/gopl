package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	str1 := "testing is cool"
	str2 := "testing is really cool"
	// str3 := "testing is cool"
	assert.Equal(t, str1, str2, "str1 and str2 should be equal")
}
