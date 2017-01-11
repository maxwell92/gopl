package error

import (
	"fmt"
	"testing"
)

func Test_NewYceError(*testing.T) {
	e := NewYceError(0, "OK")
	fmt.Printf("%v\n", e)
}

func Test_EncodeJson_DecodeJson(t *testing.T) {
	ye := NewYceError(0, "OK")

	// Encode
	json, _ := ye.EncodeJson()
	fmt.Printf("%s\n", json)

	// Decode
	e := new(YceError)
	err := e.DecodeJson(json)

	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", e)
}
