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

        nbytes, err := io.Copy(ioutil.Discard, resp.Body) //ioutil.Discard referencing to the answer
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
        }

        fmt.Fprintf(os.Stdout, "%7d\n", nbytes)

    }

}
