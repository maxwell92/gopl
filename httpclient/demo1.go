package main
import (
	"net/http"
	"fmt"
	"io/ioutil"
)
func Get () []byte {
    c := new(http.Client)
    req , err := http.NewRequest("Get","http://master:8080",nil)
    if err != nil {
        panic(err)
    }

    resp,err1 := c.Do(req)
    if err1 != nil{
        panic(err1)
    }

    response, err2 := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err2 != nil {
        fmt.Println(string(response))
    }

    return response
}

func main() {
	fmt.Println(string(Get()))
}
