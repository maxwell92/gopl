package main

import (
	"fmt"
	"testing"
	"components/template"
)

func Test_Main(t *testing.T) {
	md := &template.Markdown{
		Topic: "TestDoc",
		Author: &template.Author{
			Name: "maxwell92",
			GithubLink: "github.com/maxwell92",
		},
		Revise: &template.Revise{
			LastRevised: "2017-04-05",
			Description: "1st Test",
		},
		Purpose: "Test for DocGen",
		Request: &template.Request{
			Method: "GET",
			Url: "",
			Header: &template.Header{
				Header: map[string]string{"Authorization":"$SessionId"},
			},
			Parameters: "\"{\"Name\":\"yce-alpha}\"",
		},
		Table: &template.Table{
			Description: "",
			Design: `|Col1|Col2|
			         |:--:|:--:|
			         |Row1|Row2|`,
		},
		KeyLogic: "",
		Response: `JSON
		          {
		         	\"Code\": 100,
		          }`,
		Comments: "This is a test of DocGen",
	}
	te := template.NewTemplateEngine(md)
	fmt.Println(te.Render())
}
