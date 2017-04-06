package template

import (
	"fmt"
	"github.com/mlycore/tools/placeholder"
	"time"
	"bufio"
	"os"
	"strings"
)

type Markdown struct {
	Topic string
	Author *Author
	Revise *Revise
	Purpose string
	Request *Request
	Table *Table
	KeyLogic string
	Response string
	Comments string
}

/*
func (md *Markdown) hint(hint string) {
	fmt.Printf("%s\n", hint)
}

func (md *Markdown) input(value string) {
	consoleReader := bufio.NewReader(os.Stdin)
	input, err := consoleReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	value = input
	fmt.Printf("input: %s", value)
}
*/

func NewMarkdown() *Markdown {
	return &Markdown{
		Author: &Author{},
		Revise: &Revise{},
		Request: &Request{Header: &Header{}},
		Table: &Table{},
	}
}

/*
func (md *Markdown) Input() {
	fmt.Printf("Input: %p\n", md)
	md.hint("Topic:")
	md.input(md.Topic)
	fmt.Printf("Input: %s\n", md.Topic)

	md.hint("AuthorName:")
	md.input(md.Author.Name)
	md.hint("AuthorGithubLink:")
	md.input(md.Author.GithubLink)


	md.Revise.LastRevised = time.Now().Format(TIMEFORMAT)
	md.hint("RevisedDescription:")
	md.input(md.Revise.Description)

	md.hint("Purpose:")
	md.input(md.Purpose)

	md.hint("RequestMethod:")
	md.input(md.Request.Method)
	md.hint("RequestUrl:")
	md.input(md.Request.Url)
	md.hint("RequestHeader:")
	md.Request.Header.input()

	md.hint("RequestParameters:")
	md.input(md.Request.Parameters)

	md.hint("TableDescription:")
	md.input(md.Table.Description)
	md.hint("TableDesign:")
	md.input(md.Table.Design)

	md.hint("KeyLogic:")
	md.input(md.KeyLogic)

	md.hint("Response:")
	md.input(md.Response)

	md.hint("Comments:")
	md.input(md.Comments)
}
*/


func (md *Markdown) Input() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Println("Topic:")
	md.Topic, _ = consoleReader.ReadString('\n')
	fmt.Printf("Input: %s\n", md.Topic)

	fmt.Println("AuthorName:")
	md.Author.Name , _ = consoleReader.ReadString('\n')

	fmt.Println("AuthorGithubLink:")
	md.Author.GithubLink, _ = consoleReader.ReadString('\n')


	md.Revise.LastRevised = time.Now().Format(TIMEFORMAT)
	fmt.Println("RevisedDescription:")
	md.Revise.Description, _ = consoleReader.ReadString('\n')

	fmt.Println("Purpose:")
	md.Purpose, _ = consoleReader.ReadString('\n')

	fmt.Println("RequestMethod:")
	md.Request.Method, _ = consoleReader.ReadString('\n')
	fmt.Println("RequestUrl:")
	md.Request.Url, _ = consoleReader.ReadString('\n')
	fmt.Println("RequestHeader:")
	md.Request.Header.input()

	fmt.Println("RequestParameters:")
	md.Request.Parameters, _ = consoleReader.ReadString('\n')

	fmt.Println("TableDescription:")
	md.Table.Description, _ = consoleReader.ReadString('\n')
	fmt.Println("TableDesign:")
	md.Table.Design, _ = consoleReader.ReadString('\n')

	fmt.Println("KeyLogic:")
	md.KeyLogic, _ = consoleReader.ReadString('\n')

	fmt.Println("Response:")
	md.Response, _ = consoleReader.ReadString('\n')

	fmt.Println("Comments:")
	md.Comments, _ = consoleReader.ReadString('\n')
}

func (md *Markdown) Validate() bool {
	fmt.Printf("Validate: %p\n", md)
	fmt.Printf(
		"Validate: \n" +
		"Topic: %s\n" +
		"AuthorName: %s\n" +
		"AuthorGithubLink: %s\n" +
		"LastRevised: %s\n" +
		"RevisedDescription: %s\n" +
		"Purpose: %s\n" +
		"RequestMethod: %s\n" +
		"RequestUrl: %s\n" +
		"RequestHeader: %s\n" +
		"RequestParameters: %s\n" +
		"TableDescription: %s\n" +
		"TableDesign: %s\n" +
		"KeyLogic: %s\n" +
		"Response: %s\n" +
		"Comments: %s\n",
		md.Topic, md.Author.Name, md.Author.GithubLink, md.Revise.LastRevised, md.Revise.Description,
		md.Purpose, md.Request.Method, md.Request.Url, md.Request.Header, md.Request.Parameters,
		md.Table.Description, md.Table.Design, md.KeyLogic, md.Response, md.Comments)
	md.Topic = strings.Trim(md.Topic, "\n")
	md.Author.Name = strings.Trim(md.Author.Name, "\n")
	md.Author.GithubLink = strings.Trim(md.Author.GithubLink, "\n")
	md.Revise.LastRevised = strings.Trim(md.Revise.LastRevised, "\n")
	md.Revise.Description = strings.Trim(md.Revise.Description, "\n")
	md.Purpose = strings.Trim(md.Purpose, "\n")
	md.Request.Method = strings.Trim(md.Request.Method, "\n")
	md.Request.Url = strings.Trim(md.Request.Url, "\n")
	md.Request.Parameters = strings.Trim(md.Request.Parameters, "\n")
	md.Table.Description = strings.Trim(md.Table.Description, "\n")
	md.Table.Design = strings.Trim(md.Table.Design, "\n")
	md.KeyLogic = strings.Trim(md.KeyLogic, "\n")
	md.Response = strings.Trim(md.Response, "\n")
	md.Comments = strings.Trim(md.Comments, "\n")

	return true
}

func (md *Markdown) Substitute() string {
	fmt.Printf("Substitute: %p\n", md)
	ph := placeholder.NewPlaceHolder(MARKDOWN)
	return ph.Replace(
		"<Topic>", md.Topic,
		"<AuthorName>", md.Author.Name,
		"<AuthorGithubLink>", md.Author.GithubLink,
		"<LastRevised>", md.Revise.LastRevised,
		"<RevisedDescription>", md.Revise.Description,
		"<Purpose>", md.Purpose,
		"<RequestMethod>", md.Request.Method,
		"<RequestUrl>", md.Request.Url,
		"<RequestHeader>", md.Request.Header.String(),
		"<RequestParameters>", md.Request.Parameters,
		"<TableDescription>", md.Table.Description,
		"<TableDesign>", md.Table.Design,
		"<KeyLogic>", md.KeyLogic,
		"<Response>", md.Response,
		"<Comments>", md.Comments,
	)
}

type Author struct {
	Name string
	GithubLink string
}

type Revise struct {
	LastRevised string
	Description string
}

type Request struct {
	Method string
	Url string
	Header *Header
	Parameters string
}

type Header struct {
	Header map[string]string
}

func (h *Header) String() string {
	var str string
	for k, v := range h.Header {
		s := fmt.Sprintf("%s:%s ", k, v)
		str += s
	}
	str += "\n"
	return str
}

func (h *Header) input() {
	h.Header = make(map[string]string)

	consoleReader := bufio.NewReader(os.Stdin)
	kvStr, err := consoleReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	kvGroup := strings.Split(kvStr, " ")
	if len(kvGroup) < 1 {
		fmt.Println("no kv")
		return
	}
	fmt.Println(len(kvGroup))

	for _, g := range kvGroup {
		kv := strings.Split(g, ":")
		if len(kv) == 1 {
			kv[0] = strings.Trim(kv[0], "\n")
			h.Header[kv[0]] = kv[0]
			return
		}
		h.Header[kv[0]] = kv[1]
	}
}

type Table struct {
	Description string
	Design string
}
