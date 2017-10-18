package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func main() {
	router := httprouter.New()
	router.GET("/hello", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, "Hello\n")
		time.Sleep(5 * time.Second)
		fmt.Fprintf(w, "World\n")
	})
	http.ListenAndServe(":8082", router)
}
