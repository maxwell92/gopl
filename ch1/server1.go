package main
import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello world URL.Path = %q\n", r.URL.Path)
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":10000", nil))
}
