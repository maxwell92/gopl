package main
import (
	"encoding/json"
)

const data = `{"name": "liyao.miao"}`

type User struct {
	// Name string `json:"name"`
	Name string
}

func main( ){
	u := &User{}
	err := json.Unmarshal([]byte(data), u)
	if err != nil {
		panic(err)
	}

	println(u.Name)
}
