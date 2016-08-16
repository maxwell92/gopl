package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	pool := x509.NewCertPool()
	
	caCertPath := "ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)	
		return
	}

	pool.AppendCertsFromPEM(caCrt)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://localhost:8081")

	if err != nil {
		panic(err)	
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.Status)

}
