package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type Controller struct{}

func (c *Controller)Hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "Hello\n")
}

func main() {
	var c Controller
	r := httprouter.New() 	
	r.GET("/hello", c.Hello)	
	r.ServeFiles("/static/*filepath", http.Dir("./frontend"))
	http.ListenAndServe(":8080", r)
}
