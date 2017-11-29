package main
import (
	"fmt"
	"encoding/json"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name,omitempty"`
}

func main() {
	t := &User{
		Id: 1,
		// Name: "",
	}
	data, _ := json.Marshal(t)	
	u := new(User)
	_ = json.Unmarshal(data, u)
	fmt.Printf("%d, %s, %p\n", u.Id, u.Name, &u.Name)
}
