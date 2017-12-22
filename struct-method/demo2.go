package main
import (
	"fmt"
	"encoding/json"
)
type User struct {
	id int
	Name string
	Title string `json:"title"`
}	

func main() {
	u := &User{
		id: 1,
		Name: "Jack",
		Title: "Senior Developer",
	}

	data, _ := json.Marshal(u)
	fmt.Println(string(data))
}
