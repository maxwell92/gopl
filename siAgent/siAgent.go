package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Job struct {
	Image  string `json:"image"`
	Tag    string `json:"tag"`
	Status string `json:"status,omitempty"`
	Expire int    `json:"expire"`
	c      iris.WebsocketConnection
}

const (
	SUCC    = "SUCCEED"
	FAILED  = "FAILED"
	EXPIRED = "EXPIRED"
	RUNNING = "RUNNING"
	WAITING = "WAITING"
)

type OperationConfig struct {
	script    string
	waitQueue string
	expire    string
}

type OperationAPI struct {
	*iris.Context
	cfg *OperationConfig
	Job *Job `json:"job"`
}

type clientPage struct {
	Title string
	Host  string
}

var usage = `Usage: siAgent <filepath>
`

var waitJobQueue chan *Job
var finJobQueue chan *Job

func (j Job) Encode() string {
	jobBYTE, err := json.Marshal(j)
	if err != nil {
		log.Printf("Job Encoding Error: err=%s\n", err)
		os.Exit(1)
	}
	jobSTR := string(jobBYTE)
	return jobSTR
}

func sync(script, expire string) {
	var flag bool
	for {
		select {
		case job := <-waitJobQueue:
			{
				job.Expire, _ = strconv.Atoi(expire)
				job.Status = RUNNING
				log.Printf("%s: Waiting --> Running\n", job.Image+":"+job.Tag)
				cmd := exec.Command("/bin/bash", "-C", script, job.Image, job.Tag)
				if cmd == nil {
					log.Printf("Get Command Error\n")
					os.Exit(1)
				}

				go job.killer(cmd)

				job.c.To("syncImage").Emit("sync", job.Encode())
				flag = true
				fmt.Println(flag)
				err := cmd.Run()
				flag = false
				fmt.Println(flag)
				if err != nil {
					job.Status = FAILED
					log.Printf("%s: Running --> Killed\n", job.Image+":"+job.Tag)
					job.c.To("syncImage").Emit("sync", job.Encode())
					//finJobQueue <- job
				} else {
					job.Status = SUCC
					log.Printf("%s: Running --> Finish\n", job.Image+":"+job.Tag)
					job.c.To("syncImage").Emit("sync", job.Encode())
					//finJobQueue <- job
				}
			}
		}
	}
}

func (j Job) killer(c *exec.Cmd) {
	k := time.NewTimer(time.Duration(j.Expire) * time.Second)
	<-k.C
	c.Process.Kill()
}

func UsageAndExit() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}

}

func Init() *OperationAPI {
	op := new(OperationAPI)
	op.cfg = new(OperationConfig)

	cfg, err := os.Open(os.Args[1])
	defer cfg.Close()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	reader := bufio.NewReader(cfg)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		str := scanner.Text()

		comment := strings.Split(str, "#")[0] == ""
		pair := len(strings.Split(str, " ")) == 2

		if !comment && pair {
			content := strings.Split(str, " ")

			switch content[0] {
			case "script":
				op.cfg.script = content[1]
				break
			case "waitQueue":
				op.cfg.waitQueue = content[1]
				break
			case "expire":
				op.cfg.expire = content[1]
				break
			default:
				op.cfg.expire = ""
				op.cfg.script = ""
				op.cfg.waitQueue = ""
			}
		} else {
			continue
		}

	}

	f, err1 := os.Open(op.cfg.script)
	defer f.Close()
	if err1 != nil {
		log.Println(err)
		os.Exit(1)
	}

	waitJobLen, _ := strconv.Atoi(op.cfg.waitQueue)
	waitJobQueue = make(chan *Job, waitJobLen)
	finJobQueue = make(chan *Job, 1)

	return op
}

func (op OperationAPI) Get() {
	op.Render("client.html", clientPage{"Client Page", op.HostString()})
}

func SetupWebSocket(op *OperationAPI) {
	iris.Config.Websocket.Endpoint = "/sync"

	var syncImage = "syncImage"
	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {
		c.Join(syncImage)

		c.On("sync", func(JobSTR string) {
			job := new(Job)
			JobJSON := []byte(JobSTR)
			err := json.Unmarshal(JobJSON, job)
			if err != nil {
				log.Printf("WS Unmarshal JSON Error: err=%s\n", err)
				os.Exit(1)
			}
			job.Status = WAITING
			job.c = c
			c.To(syncImage).Emit("sync", job.Encode())
			waitJobQueue <- job
		})
		/*
			go func(c iris.WebsocketConnection) {
				for {
					select {
					case job := <-finJobQueue:
						c.To(syncImage).Emit("sync", job.Encode())
					}
				}
			}(c)
		*/
		c.OnDisconnect(func() {
		})
	})

	iris.Static("/js", "./static/js", 1)
	iris.API("/", *op)
	iris.Listen(":8081")
}

func main() {
	UsageAndExit()
	op := Init()
	go sync(op.cfg.script, op.cfg.expire)
	SetupWebSocket(op)
}
