package main

import (
    "fmt"
    //"io"
    "net/http"
    "time"

    "golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
    t := time.NewTicker(time.Duration(1) * time.Second)
    for _ = range t.C {
       fmt.Fprintf(ws, "%s\n", time.Now().String())
    }
    // io.Copy(ws, ws)
}

// This example demonstrates a trivial echo server.
func main() {
    http.Handle("/echo", websocket.Handler(EchoServer))
    fmt.Println("listen at :12345")
    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}
