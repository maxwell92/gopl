package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
)

func main() {
	cli := &http.Client{}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	// res, err := http.Get(ts.URL + "/api")
	req := httptest.NewRequest("GET", "http://api/v2/deployments", strings.NewReader(""))
	req.RequestURI = ""
	res, err:= cli.Do(req)	
	if err != nil {
		log.Fatal(err)
	}

	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}
