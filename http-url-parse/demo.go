package main
import (
	"net/http"
	"fmt"
	"strings"
)

func main() {
	http.HandleFunc("/api/v1/:namespace/deployments", func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("requested url: %s\n", r.URL.String())
		fmt.Printf("requested path: %s\n", r.URL.Path)

		var namespace string	
		for idx, p := range strings.Split("/api/v1/:namespace/deployments", "/") {
			if strings.Contains(p, ":") && strings.Split(p, ":")[1] == "namespace" {
				namespace = strings.Split(r.URL.String(), "/")[idx]				
				break
			}
		}
		fmt.Printf("namespace: %s\n", namespace)	

	})
	http.ListenAndServe(":8083", nil)
}
