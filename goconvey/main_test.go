package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAverage(t *testing.T) {
	Convey("Given some integer in a array", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})

		})
	})
}

/*
func TestAverage(t *testing.T) {
	Convey("Given some integer in an arary", t)
}
*/
