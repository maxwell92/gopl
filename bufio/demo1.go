package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"io"
	"bufio"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./dat") // dat []byte
	check(err)
	// fmt.Println(string(dat)) //Hello world
				 //yp-tc..
	// fmt.Printf("%s", string(dat)) // Hello worldyp-tc..
	fmt.Println(dat, string(dat)) // []byte: [ , , ] , Hello world
	// byte , ASCII:
	// byte -> 8 bit -> 0~255
	// ASCII 0 ~ 255

	f, err := os.Open("./dat")
	check(err)
	fmt.Println(f.Fd())

	// STDIN STDOUT STDERR -> 0 1 2
	// xx > /dev/null
	// xx 2> /dev/null
	// xx 2>&1 /dev/null

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 5)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}
