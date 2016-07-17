package main
import (
    "net/http"
    "sync"
    "log"
)

var mu sync.Mutex
var count int

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":10000", nil))
}
//TODO: set cycles through url
func handler(w http.ResponseWriter, r *http.Request) {
    cycles := float64(5)
    lissajous(w, cycles)
}

