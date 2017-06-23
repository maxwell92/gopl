package main

import (
	"testing"
)

func Benchmark_Add(b *testing.B) {
	hello := "hello"
	space := " "
	world := "world"

	for i := 0; i < b.N; i++ {
		Add(hello, space, world)
	}
}

func Benchmark_Sprintf(b *testing.B) {
	hello := "hello"
	space := " "
	world := "world"

	for i := 0; i < b.N; i++ {
		Sprintf(hello, space, world)
	}
}
