package main
import "fmt"

type Puller interface{
	Pull(src, dst string)
}

type Hunter struct {}
func (h *Hunter) Do(puller Puller) {
	puller.Pull("Beijing", "Xian")
}

type Dog struct{}
func (d *Dog)Pull(src, dst string) {
	fmt.Printf("I'm dog pulling from %s to %s\n", src, dst)
}

type Person struct{}
func (p *Person)Pull(src, dst string) {
	fmt.Printf("I'm person pulling from %s to %s\n", src, dst)
}

func main() {
	somebody := new(Person)

	Jack := new(Hunter)	
	Jack.Do(somebody)
}
