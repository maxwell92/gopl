package main
import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    r := http.Request{}
    r.Body = ioutil.NopCloser(bytes.NewBuffer([]byte("test")))

/*    s, _ := ioutil.ReadAll(r.Body)
    fmt.Println(string(s))
*/
/*    s, _ = ioutil.ReadAll(r.Body)
    fmt.Println(string(s))
*/

    content, _ := ioutil.ReadAll(r.Body)
    r.Body = ioutil.NopCloser(bytes.NewBuffer(content))
    again, _ := ioutil.ReadAll(r.Body)
    fmt.Println(string(again))
}

func ReadNotDrain(r *http.Request) (content []byte, err error) {
    content, err = ioutil.ReadAll(r.Body)
    r.Body = ioutil.NopCloser(bytes.NewBuffer(content))
    return 
} 
    

