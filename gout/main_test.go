package main
import (
	"fmt"
	"testing"
)

func Test_Add(t *testing.T) {
	if 3 == add(1, 2) {
		fmt.Println("ok")
	}else {
		fmt.Println("error")
	}
}
