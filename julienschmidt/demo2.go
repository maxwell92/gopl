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

type Validator struct{}
func (v *Validator)ValidateSession(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("This is Middleware of Validator")
		h.ServeHTTP(w, r)
	})
}

func main() {
	var c Controller
	r := httprouter.New() 	
	r.GET("/hello", c.Hello)	

	var v Validator
	http.ListenAndServe(":8080", v.ValidateSession(r))
}
