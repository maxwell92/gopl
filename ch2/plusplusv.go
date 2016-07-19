package main
import (
	"os"
	"fmt"
)

func main() {
	//for i := 0; i < len(os.Args); ++i {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i)	
	}
}
