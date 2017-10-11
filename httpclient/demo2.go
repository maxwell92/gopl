package main

import (
	"fmt"
	"strings"
	"net/http"
	"net/http/httptest"
	// "io/ioutil"
)

func main() {
	/*
		cli := &http.Client{}
	*/
	req := httptest.NewRequest("GET", "http://localhost:8080/api/v2/deployments", strings.NewReader("")) 
	req.RequestURI = ""
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("hello\n")
		fmt.Fprintf(w, "hello\n")
	})
	/*
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "hello\n")
	}))
	defer server.Close()
	*/
	
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	/*
	resp, err := cli.Do(req)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}
	data, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Printf("err1: %s\n", err1)
		return
	}
	fmt.Println(string(data))
	*/
}
