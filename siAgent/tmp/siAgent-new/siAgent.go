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
	Id     int `json:"id,omitempty"`
	Image  string `json:"image"`
	Tag    string `json:"tag"`
	Status string `json:"status,omitempty"`
	Expire int    `json:"expire"`
	Op     string `json:"op"`
	CreatedAt string `json:"createdAt"`
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
var writeJobQueue chan *Job
var runFlag bool

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

				time.Sleep(time.Duration(1) * time.Millisecond)
				//job.c.To("syncImage").Emit("sync", job.Encode())
				writeJobQueue <- job
				runFlag = true
				err := cmd.Run()
				runFlag = false 
				if err != nil {
					job.Status = FAILED
					log.Printf("%s: Running --> Killed\n", job.Image+":"+job.Tag)
					writeJobQueue <- job
					//job.c.To("syncImage").Emit("sync", job.Encode())
				} else {
					job.Status = SUCC
					log.Printf("%s: Running --> Finish\n", job.Image+":"+job.Tag)
					writeJobQueue <- job
					//job.c.To("syncImage").Emit("sync", job.Encode())
				}
			}
		}
	}
}

func emit() {
	for {
		select {
			case job := <-writeJobQueue: {
				job.c.To("syncImage").Emit("sync", job.Encode())
				time.Sleep(time.Duration(1) * time.Millisecond)
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
	writeJobLen :=  waitJobLen 
	writeJobQueue = make(chan *Job, writeJobLen)

	return op
}

func (op OperationAPI) Get() {
	op.Render("client.html", clientPage{"Client Page", op.HostString()})
}

func SetupWebSocket(op *OperationAPI) {
	iris.Config.Websocket.Endpoint = "/sync"
	counter := 0

	var syncImage = "syncImage"
	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {
		c.Join(syncImage)
		log.Printf("Received new Connection\n")

		c.On("sync", func(JobSTR string) {
			job := new(Job)
			JobJSON := []byte(JobSTR)
			err := json.Unmarshal(JobJSON, job)
			if err != nil {
				log.Printf("WS Unmarshal JSON Error: err=%s\n", err)
				os.Exit(1)
			}
			job.Status = WAITING
			counter++
			job.Id = counter
			job.c = c
			waitJobQueue <- job
			writeJobQueue <- job	
			//c.To(syncImage).Emit("sync", job.Encode())
		})
		c.OnDisconnect(func() {
			log.Printf("Disconnect one Connection\n")
		})

		go func(c iris.WebsocketConnection) {
			k := time.NewTimer(time.Duration(15) * time.Second)
			for {
				select {
					case <-k.C: {
						log.Printf("[timer] wait: %d, write: %d, runFlag: %v\n", len(waitJobQueue), len(writeJobQueue), runFlag)
						if len(waitJobQueue) == 0 && len(writeJobQueue) == 0 && !runFlag {
							c.Disconnect()			
						} else {
							k.Reset(time.Duration(15) * time.Second)	
						}		
					}	
				}	
			}
		}(c)
	})

	iris.Static("/js", "./static/js", 1)
	iris.API("/", *op)
	iris.Listen(":8081")
}

func main() {
	UsageAndExit()
	op := Init()
	go sync(op.cfg.script, op.cfg.expire)
	go emit()
	SetupWebSocket(op)
}
