package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/hello", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		fmt.Fprintf(w, "Hello World\n")
	})
	http.ListenAndServe(":8082", router)
}
