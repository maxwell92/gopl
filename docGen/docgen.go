package main

import (
	"fmt"
	"components/template"
)

func main() {
	//TODO: any content type should implement the Content interface with one method Input
	md := template.NewMarkdown()
	fmt.Printf("Main: %p\n", md)
	md.Input()
	te := template.NewTemplateEngine(md)
	fmt.Println(te.Render())
}
