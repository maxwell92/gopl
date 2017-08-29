package main
import (
	"fmt"
	"encoding/json"
)

type Car struct {
	Brand string `json:"brand"`
	Price int32 `json:"price"`
	Owner string
}

func main() {
	benz := &Car{
		Brand: "Benz",
		Price: 1000000,
		Owner: "Dark",
	}

	data, marshalErr := json.Marshal(benz) // []byte
	if marshalErr != nil {
		fmt.Println(marshalErr)
		return
	}

	fmt.Println(string(data))

	mycar := new(Car)
	unmarshalErr := json.Unmarshal(data, mycar)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
		return
	}

	fmt.Printf("%s, %d, %s\n", mycar.Brand, mycar.Price, mycar.Owner)
}

