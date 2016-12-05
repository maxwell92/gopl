package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "metadata.name"
	str2 := "spec.metadata.name"
	str3 := "spec.template.spec.containers[0].metadata.name"
	str4 := "metadata.namespace"
	str5 := "spec.template.spec.containers[0].resources.limits[\"cpu\"]"

	strs := make([]string, 0)
	strs = append(strs, str1)
	strs = append(strs, str2)
	strs = append(strs, str3)
	strs = append(strs, str4)
	strs = append(strs, str5)

	for _, s := range strs {
		/*
				if strings.Contains(s, "metadata.name") {
					fmt.Printf("name: %s\n", s)
				} else if strings.Contains(s, "metadata.namespace") {
					fmt.Printf("namespace: %s\n", s)
				}
			}
		*/
		if strings.Contains(s, "\"cpu\"") {
			fmt.Println(s)
		}

		if strings.Contains(s, "limits[\"cpu\"]") {
			fmt.Println(s)
		}
	}
}
