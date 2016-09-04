package main

import (
	"fmt"
)

func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	//*buf = append(*buf, b[bp:]...)
	*buf = append(*buf, b[bp:]...)
}

func main() {
	//buf := []byte()
	buf := new([]byte)
	fmt.Println(buf)

	i := 2016
	wid := 10
	itoa(buf, i, wid)
	fmt.Println(string(*buf))
	*buf = append(*buf, '/')
	i = 2018
	wid = -10
	itoa(buf, i, wid)
	fmt.Println(string(*buf))

}
