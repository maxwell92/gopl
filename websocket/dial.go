package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
//	"time"
)

func main() {
	origin := "http://localhost/"
	//	url := "ws://localhost:12345/echo"
	url := "ws://localhost:8080/api/v2/datacenters/75/organizations/yce/pods/yce-855162003-f469n/streamedlog"

	// we can use Dial() or DialConfig() to create a valid WebSocket Connection
	// ws, err := websocket.Dial(url, "", origin)

	cfg, err := websocket.NewConfig(url, origin)
	if err != nil {
		log.Fatal(err)
	}
        cfg.Header["Userid"] = []string{"3"}
        cfg.Header["Orgid"] = []string{"1"}
	ws, err := websocket.DialConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}
	for {
		var msg = make([]byte, 512)
		var n int
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Received: %s.", msg[:n])
		fmt.Printf("%s", msg[:n])
	}
	/*
		t, err := time.Parse(time.RFC3339, "2017-10-27T18:08:00Z")
		if err != nil {
			log.Fatal(err)
		}
		ws.SetDeadline(t)
		tm := time.NewTicker(time.Duration(1) * time.Second)
		for _ = range tm.C {
			/*
	                data := "hello, world" + time.Now().String()
			if _, err := ws.Write([]byte(data)); err != nil {
				log.Fatal(err)
			}
	*/
	/*
			var msg = make([]byte, 512)
			var n int
			if n, err = ws.Read(msg); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Received: %s.\n", msg[:n])
		}
	*/
	/*
		for _ = range tm.C {
			fmt.Printf("%s\n", time.Now().String())
		}
	*/
}
