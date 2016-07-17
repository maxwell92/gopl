package main
import (
    "fmt"
    "os" 
    "strings"
)

func main() {
    //fmt.Println(strings.Join(os.Args[0], strings.Join(os.Args[1:], " ")))
    //os.Args[0] is a string not a []string, so strings.Join() can't work

    fmt.Println(os.Args[0] + strings.Join(os.Args[1:], " "))

    //if this go file is not compiled, the os.Args[0] printed will be somethings like
    // /var/folders/1m/d_7wdylj59j7pwv4fq0q_9mr0000gn/T/go-build766749188/command-line-arguments/_obj/exe/ex1a 
    //but if it is compiled with go build, everythings will be okay!
}

