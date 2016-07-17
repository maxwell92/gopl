package main
import (
    "fmt"
    "os" 
)

func main() {
   /* var s string
    *for _, v := range os.Args[1:] {
    *    s += v
    *    s += " "
    *}
    *fmt.Println(s)
    */

    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

