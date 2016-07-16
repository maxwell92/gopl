package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "io"
)

func main() {
    for _, url := range os.Args[1:] {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
        }

        //b, err := ioutil.ReadAll(resp.Body)

        io.Copy(os.Stdout, resp)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
        }


    }

}
