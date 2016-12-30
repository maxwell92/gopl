package main

import (
	"fmt"
	"github.com/kataras/iris"
	"log"
	"os"
	"os/exec"
	"strconv"
)

var count int

type Image struct {
	Name string `json:"name"`
}

type OperationAPI struct {
	*iris.Context
	scriptPath string
	Image      *Image `json:"image"`
}

var usage = `Usage: siAgent <filepath>
`

var imageQueue chan string

func sync(script string) {
	for {
		select {
		case image := <-imageQueue:
			{
				log.Printf("Sync Start: %s\n", image)
				cmd := exec.Command("/bin/bash", "-C", script, image)
				err := cmd.Run()
				if err != nil {
					log.Println("Sync Error: err=%s\n", err)
				}
				log.Printf("Sync over: %s\n", image)

			}
		}
	}

}

func (op OperationAPI) Post() {
	log.Println("Received sync request")
	op.Image = new(Image)
	err := op.ReadJSON(op.Image)
	if err != nil {
		log.Printf("ReadJSON Error: err=%s\n", err)
	} else {
		count++
		cnt := strconv.Itoa(count)
		op.Image.Name += cnt
		log.Println(op.Image.Name)
		imageQueue <- op.Image.Name
	}

}

func UsageAndExit() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}

}

func Init() *OperationAPI {
	op := new(OperationAPI)
	op.scriptPath = os.Args[1]

	f, err := os.Open(op.scriptPath)
	defer f.Close()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return op
}

func main() {
	UsageAndExit()
	op := Init()
	imageQueue = make(chan string, 1)
	count = 0
	go sync(op.scriptPath)

	iris.API("/sync", *op)
	iris.Listen(":8081")
}
