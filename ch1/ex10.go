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


    f, err := os.OpenFile("./ex10result", os.O_RDWR | os.O_APPEND, 0666)
    if err != nil {
        fmt.Println("file open error")
    }


    for i := 0; i < 10; i++ {
        go fetch(os.Args[1], ch)
        n, err := io.WriteString(f, <-ch)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(n)
        time.Sleep(time.Second * 1)
    }


    defer f.Close()
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

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
