package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"net/http"
	"sync"
	"sync/atomic"

)

type API struct{
	Message string `json:"message"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	name := urlParams["user"]	
	HelloMessage := "Hello " + name

	message := API{HelloMessage}
	output, _ := json.Marshal(message)
	fmt.Fprintf(w, "%s", string(output))
}

func main() {
	gorillaRoute := mux.NewRouter()	
	gorillaRoute.HandleFunc("/api/{user:[0-9]+}", Hello)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)
}
