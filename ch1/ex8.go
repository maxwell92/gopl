package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "io"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
        if strings.HasPrefix(url, "http://") == false {
            url = "http://" + url
        }
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
        }


        nbytes, err := io.Copy(ioutil.Discard, resp.Body) //ioutil.Discard referencing to the answer
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
        }

        fmt.Fprintf(os.Stdout, "%7d\n", nbytes)

    }

}
