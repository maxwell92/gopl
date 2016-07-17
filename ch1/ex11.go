package main
import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "time"
    "os"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch)
    }

    /*for {
         select {
            case  re:= <-ch: {
                fmt.Println(re)
            }
        }
    }*/
    for range os.Args[1:] {
        fmt.Print(<-ch)
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

//TODO: add a function to cancel fetch() goroutine when some website cann't be accessed.
func fetch(url string, ch chan<- string) {

    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }


    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2f  %7d   %s\n", secs, nbytes, url)
}
