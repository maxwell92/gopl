package main
import "fmt"

func main() {
    var u uint8 = 255
    fmt.Println(u, u + 1, u * u)
    var i int8 = 127
    fmt.Println(i, i + 1, i * i)
   
    fmt.Println((i + 1) * (i + 1), (u + 1) * (u + 1)) //0 0
    fmt.Println((i + 1) * i, (u + 1) * u) //-128 0

// overflow
/*    i = 255
    fmt.Println(i)
*/
}
