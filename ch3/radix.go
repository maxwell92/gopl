package main
import "fmt"

func main() {
    o := 0666
    x := int64(0xdeadbeef)

    fmt.Printf("%d %[1]o %#[1]o\n", o)
    fmt.Printf("%d %[2]o %#[2]o\n", o, x)

    fmt.Printf("%d %[1]x %#[1]x\n", x)
    fmt.Printf("%d %[2]x %#[2]x\n", x)

    ascii := 'a'
    newline := '\n'
    fmt.Printf("%d %[1]c %[1]q\n", ascii)
    fmt.Printf("%d %[1]c %[1]q\n", newline)

}
