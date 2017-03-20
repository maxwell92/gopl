package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Stat()...")
	rf, _ := os.Stat("/Users/yp-tc-m-7019/mworks/gopl/filemode/regular-file")
	fmt.Printf("%s, %v, %d\n", rf.Name(), rf.Mode(), rf.Mode())
	sf, _ := os.Stat("/Users/yp-tc-m-7019/mworks/gopl/filemode/symlink-file")
	fmt.Printf("%s, %v, %d\n", sf.Name(), sf.Mode(), sf.Mode())
	hf, _ := os.Stat("/Users/yp-tc-m-7019/mworks/gopl/filemode/hardlink-file")
	fmt.Printf("%s, %v, %d\n", hf.Name(), hf.Mode(), hf.Mode())

	fmt.Println("Lstat()...")
	rf, _ = os.Lstat("/Users/yp-tc-m-7019/mworks/gopl/filemode/regular-file")
	fmt.Printf("%s, %v, %d\n", rf.Name(), rf.Mode(), rf.Mode())
	sf, _ = os.Lstat("/Users/yp-tc-m-7019/mworks/gopl/filemode/symlink-file")
	fmt.Printf("%s, %v, %d\n", sf.Name(), sf.Mode(), sf.Mode())
	hf, _ = os.Lstat("/Users/yp-tc-m-7019/mworks/gopl/filemode/hardlink-file")
	fmt.Printf("%s, %v, %d\n", hf.Name(), hf.Mode(), hf.Mode())

}
